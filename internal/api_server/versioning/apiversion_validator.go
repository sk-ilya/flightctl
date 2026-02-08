package versioning

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/flightctl/flightctl/internal/domain"
)

// StatusResponder can write a domain.Status as an HTTP error response.
// Transport handlers implement this interface via their SetResponse method.
type StatusResponder interface {
	SetResponse(http.ResponseWriter, any, domain.Status)
}

// ValidateBodyAPIVersion returns middleware that checks the body's apiVersion
// field matches the expected version. It accepts both bare ("v1beta1") and
// prefixed ("flightctl.io/v1beta1") formats by extracting the version part
// after the last slash. Unrecognized versions are rejected with 400.
func ValidateBodyAPIVersion(expected Version, responder StatusResponder) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if !validateBodyAPIVersion(w, req, expected, responder) {
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func validateBodyAPIVersion(w http.ResponseWriter, req *http.Request, expected Version, responder StatusResponder) bool {
	switch req.Method {
	case http.MethodGet, http.MethodDelete, http.MethodHead, http.MethodOptions:
		return true
	}

	if req.Body == nil || req.Body == http.NoBody {
		return true
	}

	body, err := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil || len(body) == 0 {
		return true
	}

	var envelope struct {
		ApiVersion string `json:"apiVersion"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil || envelope.ApiVersion == "" {
		return true
	}

	// Extract the version part after the last slash (e.g. "flightctl.io/v1beta1" -> "v1beta1").
	// If there is no slash, the value is used as-is.
	version := envelope.ApiVersion
	if i := strings.LastIndex(version, "/"); i >= 0 {
		version = version[i+1:]
	}

	bodyVersion := Version(version)
	if !bodyVersion.IsValid() {
		responder.SetResponse(w, nil, domain.StatusBadRequest(fmt.Sprintf(
			"apiVersion %q is not a valid API version.",
			envelope.ApiVersion,
		)))
		return false
	}

	if bodyVersion != expected {
		responder.SetResponse(w, nil, domain.StatusBadRequest(fmt.Sprintf(
			"apiVersion %q does not match the requested API version %q.",
			envelope.ApiVersion, string(expected),
		)))
		return false
	}

	return true
}
