// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3DestinationStorageClass S3 storage class.
type ObservabilityPipelineAmazonS3DestinationStorageClass string

// List of ObservabilityPipelineAmazonS3DestinationStorageClass.
const (
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_STANDARD            ObservabilityPipelineAmazonS3DestinationStorageClass = "STANDARD"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_REDUCED_REDUNDANCY  ObservabilityPipelineAmazonS3DestinationStorageClass = "REDUCED_REDUNDANCY"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_INTELLIGENT_TIERING ObservabilityPipelineAmazonS3DestinationStorageClass = "INTELLIGENT_TIERING"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_STANDARD_IA         ObservabilityPipelineAmazonS3DestinationStorageClass = "STANDARD_IA"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_EXPRESS_ONEZONE     ObservabilityPipelineAmazonS3DestinationStorageClass = "EXPRESS_ONEZONE"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_ONEZONE_IA          ObservabilityPipelineAmazonS3DestinationStorageClass = "ONEZONE_IA"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_GLACIER             ObservabilityPipelineAmazonS3DestinationStorageClass = "GLACIER"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_GLACIER_IR          ObservabilityPipelineAmazonS3DestinationStorageClass = "GLACIER_IR"
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_DEEP_ARCHIVE        ObservabilityPipelineAmazonS3DestinationStorageClass = "DEEP_ARCHIVE"
)

var allowedObservabilityPipelineAmazonS3DestinationStorageClassEnumValues = []ObservabilityPipelineAmazonS3DestinationStorageClass{
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_STANDARD,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_REDUCED_REDUNDANCY,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_INTELLIGENT_TIERING,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_STANDARD_IA,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_EXPRESS_ONEZONE,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_ONEZONE_IA,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_GLACIER,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_GLACIER_IR,
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONSTORAGECLASS_DEEP_ARCHIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3DestinationStorageClass) GetAllowedValues() []ObservabilityPipelineAmazonS3DestinationStorageClass {
	return allowedObservabilityPipelineAmazonS3DestinationStorageClassEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3DestinationStorageClass) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3DestinationStorageClass(value)
	return nil
}

// NewObservabilityPipelineAmazonS3DestinationStorageClassFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3DestinationStorageClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3DestinationStorageClassFromValue(v string) (*ObservabilityPipelineAmazonS3DestinationStorageClass, error) {
	ev := ObservabilityPipelineAmazonS3DestinationStorageClass(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3DestinationStorageClass: valid values are %v", v, allowedObservabilityPipelineAmazonS3DestinationStorageClassEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3DestinationStorageClass) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3DestinationStorageClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3DestinationStorageClass value.
func (v ObservabilityPipelineAmazonS3DestinationStorageClass) Ptr() *ObservabilityPipelineAmazonS3DestinationStorageClass {
	return &v
}
