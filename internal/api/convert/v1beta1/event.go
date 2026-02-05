package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./event_conv.gen.go
// goverter:name EventConverterImpl
type EventConverter interface {
	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | EventListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Event]) *apiv1beta1.EventList

	// ListParamsToDomain converts API list params to domain ResourceListParams
	// Order is extracted in the transport layer (source field, no target equivalent)
	// ListEventsParams has no LabelSelector, so ignore that target field
	// goverter:ignore LabelSelector
	ListParamsToDomain(apiv1beta1.ListEventsParams) domain.ResourceListParams
}
