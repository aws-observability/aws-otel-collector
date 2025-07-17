// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOcsfMappingLibrary Predefined library mappings for common log formats.
type ObservabilityPipelineOcsfMappingLibrary string

// List of ObservabilityPipelineOcsfMappingLibrary.
const (
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_CLOUDTRAIL_ACCOUNT_CHANGE                 ObservabilityPipelineOcsfMappingLibrary = "CloudTrail Account Change"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_CREATEBUCKET              ObservabilityPipelineOcsfMappingLibrary = "GCP Cloud Audit CreateBucket"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_CREATESINK                ObservabilityPipelineOcsfMappingLibrary = "GCP Cloud Audit CreateSink"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_SETIAMPOLICY              ObservabilityPipelineOcsfMappingLibrary = "GCP Cloud Audit SetIamPolicy"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_UPDATESINK                ObservabilityPipelineOcsfMappingLibrary = "GCP Cloud Audit UpdateSink"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GITHUB_AUDIT_LOG_API_ACTIVITY             ObservabilityPipelineOcsfMappingLibrary = "Github Audit Log API Activity"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GOOGLE_WORKSPACE_ADMIN_AUDIT_ADDPRIVILEGE ObservabilityPipelineOcsfMappingLibrary = "Google Workspace Admin Audit addPrivilege"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_MICROSOFT_365_DEFENDER_INCIDENT           ObservabilityPipelineOcsfMappingLibrary = "Microsoft 365 Defender Incident"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_MICROSOFT_365_DEFENDER_USERLOGGEDIN       ObservabilityPipelineOcsfMappingLibrary = "Microsoft 365 Defender UserLoggedIn"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_OKTA_SYSTEM_LOG_AUTHENTICATION            ObservabilityPipelineOcsfMappingLibrary = "Okta System Log Authentication"
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_PALO_ALTO_NETWORKS_FIREWALL_TRAFFIC       ObservabilityPipelineOcsfMappingLibrary = "Palo Alto Networks Firewall Traffic"
)

var allowedObservabilityPipelineOcsfMappingLibraryEnumValues = []ObservabilityPipelineOcsfMappingLibrary{
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_CLOUDTRAIL_ACCOUNT_CHANGE,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_CREATEBUCKET,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_CREATESINK,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_SETIAMPOLICY,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GCP_CLOUD_AUDIT_UPDATESINK,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GITHUB_AUDIT_LOG_API_ACTIVITY,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_GOOGLE_WORKSPACE_ADMIN_AUDIT_ADDPRIVILEGE,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_MICROSOFT_365_DEFENDER_INCIDENT,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_MICROSOFT_365_DEFENDER_USERLOGGEDIN,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_OKTA_SYSTEM_LOG_AUTHENTICATION,
	OBSERVABILITYPIPELINEOCSFMAPPINGLIBRARY_PALO_ALTO_NETWORKS_FIREWALL_TRAFFIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineOcsfMappingLibrary) GetAllowedValues() []ObservabilityPipelineOcsfMappingLibrary {
	return allowedObservabilityPipelineOcsfMappingLibraryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineOcsfMappingLibrary) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineOcsfMappingLibrary(value)
	return nil
}

// NewObservabilityPipelineOcsfMappingLibraryFromValue returns a pointer to a valid ObservabilityPipelineOcsfMappingLibrary
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineOcsfMappingLibraryFromValue(v string) (*ObservabilityPipelineOcsfMappingLibrary, error) {
	ev := ObservabilityPipelineOcsfMappingLibrary(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineOcsfMappingLibrary: valid values are %v", v, allowedObservabilityPipelineOcsfMappingLibraryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineOcsfMappingLibrary) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineOcsfMappingLibraryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineOcsfMappingLibrary value.
func (v ObservabilityPipelineOcsfMappingLibrary) Ptr() *ObservabilityPipelineOcsfMappingLibrary {
	return &v
}
