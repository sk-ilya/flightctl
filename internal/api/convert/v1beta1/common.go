package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./common_conv.gen.go
// goverter:name CommonConverterImpl
// goverter:skipCopySameType
type CommonConverter interface {
	PatchRequestToDomain(apiv1beta1.PatchRequest) domain.PatchRequest
	StatusFromDomain(domain.Status) apiv1beta1.Status
	LabelListFromDomain(*domain.LabelList) *apiv1beta1.LabelList

	ListLabelsParamsToDomain(apiv1beta1.ListLabelsParams) domain.ListLabelsParams
}
