// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsResponseDataType
type ResolveVulnerableSymbolsResponseDataType string

// List of ResolveVulnerableSymbolsResponseDataType.
const (
	RESOLVEVULNERABLESYMBOLSRESPONSEDATATYPE_RESOLVE_VULNERABLE_SYMBOLS_RESPONSE ResolveVulnerableSymbolsResponseDataType = "resolve-vulnerable-symbols-response"
)

var allowedResolveVulnerableSymbolsResponseDataTypeEnumValues = []ResolveVulnerableSymbolsResponseDataType{
	RESOLVEVULNERABLESYMBOLSRESPONSEDATATYPE_RESOLVE_VULNERABLE_SYMBOLS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ResolveVulnerableSymbolsResponseDataType) GetAllowedValues() []ResolveVulnerableSymbolsResponseDataType {
	return allowedResolveVulnerableSymbolsResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ResolveVulnerableSymbolsResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ResolveVulnerableSymbolsResponseDataType(value)
	return nil
}

// NewResolveVulnerableSymbolsResponseDataTypeFromValue returns a pointer to a valid ResolveVulnerableSymbolsResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewResolveVulnerableSymbolsResponseDataTypeFromValue(v string) (*ResolveVulnerableSymbolsResponseDataType, error) {
	ev := ResolveVulnerableSymbolsResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ResolveVulnerableSymbolsResponseDataType: valid values are %v", v, allowedResolveVulnerableSymbolsResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ResolveVulnerableSymbolsResponseDataType) IsValid() bool {
	for _, existing := range allowedResolveVulnerableSymbolsResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResolveVulnerableSymbolsResponseDataType value.
func (v ResolveVulnerableSymbolsResponseDataType) Ptr() *ResolveVulnerableSymbolsResponseDataType {
	return &v
}
