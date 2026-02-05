package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./authprovider_conv.gen.go
// goverter:name AuthProviderConverterImpl
type AuthProviderConverter interface {
	ToDomain(apiv1beta1.AuthProvider) domain.AuthProvider
	FromDomain(*domain.AuthProvider) *apiv1beta1.AuthProvider

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | AuthProviderListKind
	ListFromDomain(*domain.AuthProviderList) *apiv1beta1.AuthProviderList

	ListParamsToDomain(apiv1beta1.ListAuthProvidersParams) domain.ResourceListParams
}
