package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./event_conv.gen.go
// goverter:name EventConverterImpl
// goverter:skipCopySameType
type EventConverter interface {
	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | EventListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Event]) *apiv1beta1.EventList

	ListParamsToDomain(apiv1beta1.ListEventsParams) domain.ListEventsParams
}
