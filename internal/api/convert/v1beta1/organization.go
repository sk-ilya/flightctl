package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./organization_conv.gen.go
// goverter:name OrganizationConverterImpl
type OrganizationConverter interface {
	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | OrganizationListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Organization]) *apiv1beta1.OrganizationList

	// ListOrganizationsParams only has FieldSelector
	// goverter:ignore Continue
	// goverter:ignore LabelSelector
	// goverter:ignore Limit
	ListParamsToDomain(apiv1beta1.ListOrganizationsParams) domain.ResourceListParams
}
