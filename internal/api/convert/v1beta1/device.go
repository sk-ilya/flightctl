package v1beta1

import (
	apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"
	"github.com/flightctl/flightctl/internal/domain"
)

// goverter:converter
// goverter:output:file ./device_conv.gen.go
// goverter:name DeviceConverterImpl
type DeviceConverter interface {
	ToDomain(apiv1beta1.Device) domain.Device
	FromDomain(*domain.Device) *apiv1beta1.Device

	// goverter:map . ApiVersion | APIVersion
	// goverter:map . Kind | DeviceListKind
	// goverter:map ResourceList.Pagination Metadata
	// goverter:autoMap ResourceList
	ListFromDomain(*domain.DeviceList) *apiv1beta1.DeviceList

	DecommissionToDomain(apiv1beta1.DeviceDecommission) domain.DeviceDecommission
	ResumeRequestToDomain(apiv1beta1.DeviceResumeRequest) domain.DeviceResumeRequest
	ResumeResponseFromDomain(domain.DeviceResumeResponse) apiv1beta1.DeviceResumeResponse
	LastSeenFromDomain(*domain.DeviceLastSeen) *apiv1beta1.DeviceLastSeen
	GetRenderedParamsToDomain(apiv1beta1.GetRenderedDeviceParams) domain.GetRenderedDeviceParams

	// ListParamsToDomain converts API list params to domain ResourceListParams
	// SummaryOnly is extracted in the transport layer (source field, no target equivalent)
	ListParamsToDomain(apiv1beta1.ListDevicesParams) domain.ResourceListParams
}
