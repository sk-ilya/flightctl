package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./auth_conv.gen.go
// goverter:name AuthConverterImpl
// goverter:skipCopySameType
type AuthConverter interface {
	TokenRequestToDomain(*apiv1beta1.TokenRequest) *domain.TokenRequest
	TokenResponseFromDomain(*domain.TokenResponse) *apiv1beta1.TokenResponse
}
