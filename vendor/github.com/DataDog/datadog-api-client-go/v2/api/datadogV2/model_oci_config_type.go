// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OCIConfigType Type of OCI config.
type OCIConfigType string

// List of OCIConfigType.
const (
	OCICONFIGTYPE_OCI_CONFIG OCIConfigType = "oci_config"
)

var allowedOCIConfigTypeEnumValues = []OCIConfigType{
	OCICONFIGTYPE_OCI_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OCIConfigType) GetAllowedValues() []OCIConfigType {
	return allowedOCIConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OCIConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OCIConfigType(value)
	return nil
}

// NewOCIConfigTypeFromValue returns a pointer to a valid OCIConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOCIConfigTypeFromValue(v string) (*OCIConfigType, error) {
	ev := OCIConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OCIConfigType: valid values are %v", v, allowedOCIConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OCIConfigType) IsValid() bool {
	for _, existing := range allowedOCIConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OCIConfigType value.
func (v OCIConfigType) Ptr() *OCIConfigType {
	return &v
}
