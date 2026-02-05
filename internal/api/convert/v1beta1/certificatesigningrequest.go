package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./certificatesigningrequest_conv.gen.go
// goverter:name CertificateSigningRequestConverterImpl
type CertificateSigningRequestConverter interface {
	ToDomain(apiv1beta1.CertificateSigningRequest) domain.CertificateSigningRequest
	FromDomain(*domain.CertificateSigningRequest) *apiv1beta1.CertificateSigningRequest

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | CertificateSigningRequestListKind
	// goverter:map Pagination Metadata
	ListFromDomain(*domain.ResourceList[domain.CertificateSigningRequest]) *apiv1beta1.CertificateSigningRequestList

	ListParamsToDomain(apiv1beta1.ListCertificateSigningRequestsParams) domain.ResourceListParams
}
