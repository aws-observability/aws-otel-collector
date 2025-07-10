// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestinationType The destination type. Always `google_cloud_storage`.
type ObservabilityPipelineGoogleCloudStorageDestinationType string

// List of ObservabilityPipelineGoogleCloudStorageDestinationType.
const (
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONTYPE_GOOGLE_CLOUD_STORAGE ObservabilityPipelineGoogleCloudStorageDestinationType = "google_cloud_storage"
)

var allowedObservabilityPipelineGoogleCloudStorageDestinationTypeEnumValues = []ObservabilityPipelineGoogleCloudStorageDestinationType{
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONTYPE_GOOGLE_CLOUD_STORAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationType) GetAllowedValues() []ObservabilityPipelineGoogleCloudStorageDestinationType {
	return allowedObservabilityPipelineGoogleCloudStorageDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGoogleCloudStorageDestinationType(value)
	return nil
}

// NewObservabilityPipelineGoogleCloudStorageDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineGoogleCloudStorageDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGoogleCloudStorageDestinationTypeFromValue(v string) (*ObservabilityPipelineGoogleCloudStorageDestinationType, error) {
	ev := ObservabilityPipelineGoogleCloudStorageDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGoogleCloudStorageDestinationType: valid values are %v", v, allowedObservabilityPipelineGoogleCloudStorageDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGoogleCloudStorageDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGoogleCloudStorageDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGoogleCloudStorageDestinationType value.
func (v ObservabilityPipelineGoogleCloudStorageDestinationType) Ptr() *ObservabilityPipelineGoogleCloudStorageDestinationType {
	return &v
}
