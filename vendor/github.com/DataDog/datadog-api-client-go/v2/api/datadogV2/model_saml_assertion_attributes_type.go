// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLAssertionAttributesType SAML assertion attributes resource type.
type SAMLAssertionAttributesType string

// List of SAMLAssertionAttributesType.
const (
	SAMLASSERTIONATTRIBUTESTYPE_SAML_ASSERTION_ATTRIBUTES SAMLAssertionAttributesType = "saml_assertion_attributes"
)

var allowedSAMLAssertionAttributesTypeEnumValues = []SAMLAssertionAttributesType{
	SAMLASSERTIONATTRIBUTESTYPE_SAML_ASSERTION_ATTRIBUTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SAMLAssertionAttributesType) GetAllowedValues() []SAMLAssertionAttributesType {
	return allowedSAMLAssertionAttributesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SAMLAssertionAttributesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SAMLAssertionAttributesType(value)
	return nil
}

// NewSAMLAssertionAttributesTypeFromValue returns a pointer to a valid SAMLAssertionAttributesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSAMLAssertionAttributesTypeFromValue(v string) (*SAMLAssertionAttributesType, error) {
	ev := SAMLAssertionAttributesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SAMLAssertionAttributesType: valid values are %v", v, allowedSAMLAssertionAttributesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SAMLAssertionAttributesType) IsValid() bool {
	for _, existing := range allowedSAMLAssertionAttributesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SAMLAssertionAttributesType value.
func (v SAMLAssertionAttributesType) Ptr() *SAMLAssertionAttributesType {
	return &v
}
