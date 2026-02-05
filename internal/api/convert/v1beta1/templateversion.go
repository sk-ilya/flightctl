package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./templateversion_conv.gen.go
// goverter:name TemplateVersionConverterImpl
type TemplateVersionConverter interface {
	FromDomain(*domain.TemplateVersion) *apiv1beta1.TemplateVersion

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | TemplateVersionListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.TemplateVersion]) *apiv1beta1.TemplateVersionList

	ListParamsToDomain(apiv1beta1.ListTemplateVersionsParams) domain.ResourceListParams
}
