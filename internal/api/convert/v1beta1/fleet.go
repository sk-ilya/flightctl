package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./fleet_conv.gen.go
// goverter:name FleetConverterImpl
type FleetConverter interface {
	ToDomain(apiv1beta1.Fleet) domain.Fleet
	FromDomain(*domain.Fleet) *apiv1beta1.Fleet

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | FleetListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Fleet]) *apiv1beta1.FleetList

	GetParamsToDomain(apiv1beta1.GetFleetParams) domain.GetFleetParams

	// ListParamsToDomain converts API list params to domain ResourceListParams
	// AddDevicesSummary is extracted in the transport layer (source field, no target equivalent)
	ListParamsToDomain(apiv1beta1.ListFleetsParams) domain.ResourceListParams
}
