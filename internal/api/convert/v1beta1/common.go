package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./common_conv.gen.go
// goverter:name CommonConverterImpl
type CommonConverter interface {
	PatchRequestToDomain(apiv1beta1.PatchRequest) domain.PatchRequest
	StatusFromDomain(domain.Status) apiv1beta1.Status
	LabelListFromDomain(*domain.LabelList) *apiv1beta1.LabelList

	// ListLabelsParamsToDomain converts API list params to domain ResourceListParams
	// Kind is extracted in the transport layer (source field, no target equivalent)
	// ListLabelsParams has no Continue, so ignore that target field
	// goverter:ignore Continue
	ListLabelsParamsToDomain(apiv1beta1.ListLabelsParams) domain.ResourceListParams
}
