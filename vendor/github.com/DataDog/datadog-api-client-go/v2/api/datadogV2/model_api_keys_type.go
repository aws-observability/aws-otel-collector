// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeysType API Keys resource type.
type APIKeysType string

// List of APIKeysType.
const (
	APIKEYSTYPE_API_KEYS APIKeysType = "api_keys"
)

var allowedAPIKeysTypeEnumValues = []APIKeysType{
	APIKEYSTYPE_API_KEYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *APIKeysType) GetAllowedValues() []APIKeysType {
	return allowedAPIKeysTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *APIKeysType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = APIKeysType(value)
	return nil
}

// NewAPIKeysTypeFromValue returns a pointer to a valid APIKeysType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAPIKeysTypeFromValue(v string) (*APIKeysType, error) {
	ev := APIKeysType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for APIKeysType: valid values are %v", v, allowedAPIKeysTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v APIKeysType) IsValid() bool {
	for _, existing := range allowedAPIKeysTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to APIKeysType value.
func (v APIKeysType) Ptr() *APIKeysType {
	return &v
}
