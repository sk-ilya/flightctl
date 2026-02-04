package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./enrollmentrequest_conv.gen.go
// goverter:name EnrollmentRequestConverterImpl
// goverter:skipCopySameType
type EnrollmentRequestConverter interface {
	ToDomain(apiv1beta1.EnrollmentRequest) domain.EnrollmentRequest
	FromDomain(*domain.EnrollmentRequest) *apiv1beta1.EnrollmentRequest

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | EnrollmentRequestListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.EnrollmentRequest]) *apiv1beta1.EnrollmentRequestList

	ApprovalToDomain(apiv1beta1.EnrollmentRequestApproval) domain.EnrollmentRequestApproval
	ApprovalStatusFromDomain(*domain.EnrollmentRequestApprovalStatus) *apiv1beta1.EnrollmentRequestApprovalStatus

	ConfigFromDomain(*domain.EnrollmentConfig) *apiv1beta1.EnrollmentConfig

	ListParamsToDomain(apiv1beta1.ListEnrollmentRequestsParams) domain.ListEnrollmentRequestsParams
	GetConfigParamsToDomain(apiv1beta1.GetEnrollmentConfigParams) domain.GetEnrollmentConfigParams
}
