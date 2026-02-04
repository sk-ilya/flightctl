package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./organization_conv.gen.go
// goverter:name OrganizationConverterImpl
// goverter:skipCopySameType
type OrganizationConverter interface {
	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | OrganizationListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Organization]) *apiv1beta1.OrganizationList

	ListParamsToDomain(apiv1beta1.ListOrganizationsParams) domain.ListOrganizationsParams
}
