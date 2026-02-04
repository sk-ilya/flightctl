package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./fleet_conv.gen.go
// goverter:name FleetConverterImpl
// goverter:skipCopySameType
type FleetConverter interface {
	ToDomain(apiv1beta1.Fleet) domain.Fleet
	FromDomain(*domain.Fleet) *apiv1beta1.Fleet

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | FleetListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Fleet]) *apiv1beta1.FleetList

	ListParamsToDomain(apiv1beta1.ListFleetsParams) domain.ListFleetsParams
	GetParamsToDomain(apiv1beta1.GetFleetParams) domain.GetFleetParams
}
