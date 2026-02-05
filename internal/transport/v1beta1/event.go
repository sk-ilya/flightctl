package transportv1beta1

import (
	"net/http"

	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
	"github.com/flightctl/flightctl/internal/transport"
)

// (GET /api/v1/events)
func (h *TransportHandler) ListEvents(w http.ResponseWriter, r *http.Request, params apiv1beta1.ListEventsParams) {
	domainParams := h.converter.Event().ListParamsToDomain(params)

	// Extract order from API params and convert to domain type
	var order *domain.SortOrder
	if params.Order != nil {
		o := domain.SortOrder(map[apiv1beta1.ListEventsParamsOrder]string{
			apiv1beta1.Asc:  string(domain.SortAsc),
			apiv1beta1.Desc: string(domain.SortDesc),
		}[*params.Order])
		order = &o
	}

	body, status := h.serviceHandler.ListEvents(r.Context(), transport.OrgIDFromContext(r.Context()), domainParams, order)
	apiResult := h.converter.Event().ListFromDomain(body)
	transport.SetResponse(w, apiResult, status)
}
