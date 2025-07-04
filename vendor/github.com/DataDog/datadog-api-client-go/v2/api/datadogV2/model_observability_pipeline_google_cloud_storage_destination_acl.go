// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestinationAcl Access control list setting for objects written to the bucket.
type ObservabilityPipelineGoogleCloudStorageDestinationAcl string

// List of ObservabilityPipelineGoogleCloudStorageDestinationAcl.
const (
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_PRIVATE                            ObservabilityPipelineGoogleCloudStorageDestinationAcl = "private"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_PROJECTNOT_PRIVATE                 ObservabilityPipelineGoogleCloudStorageDestinationAcl = "project-private"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_PUBLICNOT_READ                     ObservabilityPipelineGoogleCloudStorageDestinationAcl = "public-read"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_AUTHENTICATEDNOT_READ              ObservabilityPipelineGoogleCloudStorageDestinationAcl = "authenticated-read"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_BUCKETNOT_OWNERNOT_READ            ObservabilityPipelineGoogleCloudStorageDestinationAcl = "bucket-owner-read"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_BUCKETNOT_OWNERNOT_FULLNOT_CONTROL ObservabilityPipelineGoogleCloudStorageDestinationAcl = "bucket-owner-full-control"
)

var allowedObservabilityPipelineGoogleCloudStorageDestinationAclEnumValues = []ObservabilityPipelineGoogleCloudStorageDestinationAcl{
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_PRIVATE,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_PROJECTNOT_PRIVATE,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_PUBLICNOT_READ,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_AUTHENTICATEDNOT_READ,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_BUCKETNOT_OWNERNOT_READ,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONACL_BUCKETNOT_OWNERNOT_FULLNOT_CONTROL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationAcl) GetAllowedValues() []ObservabilityPipelineGoogleCloudStorageDestinationAcl {
	return allowedObservabilityPipelineGoogleCloudStorageDestinationAclEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationAcl) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGoogleCloudStorageDestinationAcl(value)
	return nil
}

// NewObservabilityPipelineGoogleCloudStorageDestinationAclFromValue returns a pointer to a valid ObservabilityPipelineGoogleCloudStorageDestinationAcl
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGoogleCloudStorageDestinationAclFromValue(v string) (*ObservabilityPipelineGoogleCloudStorageDestinationAcl, error) {
	ev := ObservabilityPipelineGoogleCloudStorageDestinationAcl(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGoogleCloudStorageDestinationAcl: valid values are %v", v, allowedObservabilityPipelineGoogleCloudStorageDestinationAclEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGoogleCloudStorageDestinationAcl) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGoogleCloudStorageDestinationAclEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGoogleCloudStorageDestinationAcl value.
func (v ObservabilityPipelineGoogleCloudStorageDestinationAcl) Ptr() *ObservabilityPipelineGoogleCloudStorageDestinationAcl {
	return &v
}
