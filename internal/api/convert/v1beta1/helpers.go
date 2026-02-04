package v1beta1

import apiv1beta1 "github.com/flightctl/flightctl/api/core/v1beta1"

func APIVersion() string                { return apiv1beta1.APIGroup + "/v1beta1" }
func DeviceListKind() string            { return apiv1beta1.DeviceListKind }
func FleetListKind() string             { return apiv1beta1.FleetListKind }
func RepositoryListKind() string        { return apiv1beta1.RepositoryListKind }
func EnrollmentRequestListKind() string { return apiv1beta1.EnrollmentRequestListKind }
func CertificateSigningRequestListKind() string {
	return apiv1beta1.CertificateSigningRequestListKind
}
func ResourceSyncListKind() string    { return apiv1beta1.ResourceSyncListKind }
func TemplateVersionListKind() string { return apiv1beta1.TemplateVersionListKind }
func EventListKind() string           { return apiv1beta1.EventListKind }
func OrganizationListKind() string    { return apiv1beta1.OrganizationListKind }
func AuthProviderListKind() string    { return apiv1beta1.AuthProviderListKind }
