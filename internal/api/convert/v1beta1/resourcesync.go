package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./resourcesync_conv.gen.go
// goverter:name ResourceSyncConverterImpl
type ResourceSyncConverter interface {
	ToDomain(apiv1beta1.ResourceSync) domain.ResourceSync
	FromDomain(*domain.ResourceSync) *apiv1beta1.ResourceSync

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | ResourceSyncListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.ResourceSync]) *apiv1beta1.ResourceSyncList

	ListParamsToDomain(apiv1beta1.ListResourceSyncsParams) domain.ResourceListParams
}
