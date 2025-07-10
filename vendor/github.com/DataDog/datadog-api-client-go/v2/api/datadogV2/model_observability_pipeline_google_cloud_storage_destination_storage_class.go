// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestinationStorageClass Storage class used for objects stored in GCS.
type ObservabilityPipelineGoogleCloudStorageDestinationStorageClass string

// List of ObservabilityPipelineGoogleCloudStorageDestinationStorageClass.
const (
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_STANDARD ObservabilityPipelineGoogleCloudStorageDestinationStorageClass = "STANDARD"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_NEARLINE ObservabilityPipelineGoogleCloudStorageDestinationStorageClass = "NEARLINE"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_COLDLINE ObservabilityPipelineGoogleCloudStorageDestinationStorageClass = "COLDLINE"
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_ARCHIVE  ObservabilityPipelineGoogleCloudStorageDestinationStorageClass = "ARCHIVE"
)

var allowedObservabilityPipelineGoogleCloudStorageDestinationStorageClassEnumValues = []ObservabilityPipelineGoogleCloudStorageDestinationStorageClass{
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_STANDARD,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_NEARLINE,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_COLDLINE,
	OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONSTORAGECLASS_ARCHIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationStorageClass) GetAllowedValues() []ObservabilityPipelineGoogleCloudStorageDestinationStorageClass {
	return allowedObservabilityPipelineGoogleCloudStorageDestinationStorageClassEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGoogleCloudStorageDestinationStorageClass) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGoogleCloudStorageDestinationStorageClass(value)
	return nil
}

// NewObservabilityPipelineGoogleCloudStorageDestinationStorageClassFromValue returns a pointer to a valid ObservabilityPipelineGoogleCloudStorageDestinationStorageClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGoogleCloudStorageDestinationStorageClassFromValue(v string) (*ObservabilityPipelineGoogleCloudStorageDestinationStorageClass, error) {
	ev := ObservabilityPipelineGoogleCloudStorageDestinationStorageClass(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGoogleCloudStorageDestinationStorageClass: valid values are %v", v, allowedObservabilityPipelineGoogleCloudStorageDestinationStorageClassEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGoogleCloudStorageDestinationStorageClass) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGoogleCloudStorageDestinationStorageClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGoogleCloudStorageDestinationStorageClass value.
func (v ObservabilityPipelineGoogleCloudStorageDestinationStorageClass) Ptr() *ObservabilityPipelineGoogleCloudStorageDestinationStorageClass {
	return &v
}
