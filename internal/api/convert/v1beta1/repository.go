package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./repository_conv.gen.go
// goverter:name RepositoryConverterImpl
type RepositoryConverter interface {
	ToDomain(apiv1beta1.Repository) domain.Repository
	FromDomain(*domain.Repository) *apiv1beta1.Repository

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | RepositoryListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.Repository]) *apiv1beta1.RepositoryList

	ListParamsToDomain(apiv1beta1.ListRepositoriesParams) domain.ResourceListParams
}
