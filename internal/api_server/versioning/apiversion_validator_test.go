package versioning

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flightctl/flightctl/internal/domain"
	"github.com/flightctl/flightctl/internal/transport"
)

type testHandler struct{}

func (testHandler) SetResponse(w http.ResponseWriter, _ any, s domain.Status) {
	transport.WriteJSONResponse(w, nil, s, int(s.Code))
}

func TestValidateBodyAPIVersion(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expected       Version
		wantStatusCode int
	}{
		{
			name:           "GET request is skipped",
			method:         http.MethodGet,
			body:           `{"apiVersion":"v1alpha1"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "DELETE request is skipped",
			method:         http.MethodDelete,
			body:           `{"apiVersion":"v1alpha1"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "HEAD request is skipped",
			method:         http.MethodHead,
			body:           `{"apiVersion":"v1alpha1"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "OPTIONS request is skipped",
			method:         http.MethodOptions,
			body:           `{"apiVersion":"v1alpha1"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "empty body is skipped",
			method:         http.MethodPost,
			body:           "",
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "JSON array body is skipped",
			method:         http.MethodPatch,
			body:           `[{"op":"replace","path":"/metadata/labels/key","value":"val"}]`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "body without apiVersion field is skipped",
			method:         http.MethodPost,
			body:           `{"kind":"Device","metadata":{}}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "malformed JSON is skipped",
			method:         http.MethodPost,
			body:           `{not valid json`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "bare v1beta1 match",
			method:         http.MethodPost,
			body:           `{"apiVersion":"v1beta1","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "prefixed v1beta1 match",
			method:         http.MethodPost,
			body:           `{"apiVersion":"flightctl.io/v1beta1","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "bare v1alpha1 match",
			method:         http.MethodPut,
			body:           `{"apiVersion":"v1alpha1","kind":"Catalog"}`,
			expected:       V1Alpha1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "prefixed v1alpha1 match",
			method:         http.MethodPut,
			body:           `{"apiVersion":"flightctl.io/v1alpha1","kind":"Catalog"}`,
			expected:       V1Alpha1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "bare mismatch returns 400",
			method:         http.MethodPost,
			body:           `{"apiVersion":"v1alpha1","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "prefixed mismatch returns 400",
			method:         http.MethodPost,
			body:           `{"apiVersion":"flightctl.io/v1alpha1","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "different prefix still matches after extracting version",
			method:         http.MethodPost,
			body:           `{"apiVersion":"api.flightctl.io/v1beta1","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "invalid version returns 400",
			method:         http.MethodPost,
			body:           `{"apiVersion":"v2","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "invalid prefixed version returns 400",
			method:         http.MethodPost,
			body:           `{"apiVersion":"flightctl.io/v99","kind":"Device"}`,
			expected:       V1Beta1,
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlerCalled := false
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlerCalled = true
				w.WriteHeader(http.StatusOK)
			})

			mw := ValidateBodyAPIVersion(tt.expected, testHandler{})
			wrapped := mw(handler)

			var body io.Reader
			if tt.body != "" {
				body = strings.NewReader(tt.body)
			}
			req := httptest.NewRequest(tt.method, "/test", body)

			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("status code = %v, want %v", rec.Code, tt.wantStatusCode)
			}

			if tt.wantStatusCode == http.StatusOK && !handlerCalled {
				t.Error("expected handler to be called, but it was not")
			}
			if tt.wantStatusCode == http.StatusBadRequest && handlerCalled {
				t.Error("expected handler NOT to be called on 400, but it was")
			}

			if tt.wantStatusCode == http.StatusBadRequest {
				var status domain.Status
				if err := json.NewDecoder(rec.Body).Decode(&status); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}
				if status.Code != http.StatusBadRequest {
					t.Errorf("error status code = %v, want %v", status.Code, http.StatusBadRequest)
				}
				if status.Message == "" {
					t.Error("error message should not be empty")
				}
			}
		})
	}
}

func TestValidateBodyAPIVersion_bodyPassthrough(t *testing.T) {
	originalBody := `{"apiVersion":"v1beta1","kind":"Device","metadata":{"name":"test"}}`

	var capturedBody string
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read body in handler: %v", err)
		}
		capturedBody = string(data)
		w.WriteHeader(http.StatusOK)
	})

	mw := ValidateBodyAPIVersion(V1Beta1, testHandler{})
	wrapped := mw(handler)

	req := httptest.NewRequest(http.MethodPost, "/devices", strings.NewReader(originalBody))

	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status code = %v, want %v", rec.Code, http.StatusOK)
	}

	if capturedBody != originalBody {
		t.Errorf("downstream body = %q, want %q", capturedBody, originalBody)
	}
}
