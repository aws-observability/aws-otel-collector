// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsRequestDataType
type ResolveVulnerableSymbolsRequestDataType string

// List of ResolveVulnerableSymbolsRequestDataType.
const (
	RESOLVEVULNERABLESYMBOLSREQUESTDATATYPE_RESOLVE_VULNERABLE_SYMBOLS_REQUEST ResolveVulnerableSymbolsRequestDataType = "resolve-vulnerable-symbols-request"
)

var allowedResolveVulnerableSymbolsRequestDataTypeEnumValues = []ResolveVulnerableSymbolsRequestDataType{
	RESOLVEVULNERABLESYMBOLSREQUESTDATATYPE_RESOLVE_VULNERABLE_SYMBOLS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ResolveVulnerableSymbolsRequestDataType) GetAllowedValues() []ResolveVulnerableSymbolsRequestDataType {
	return allowedResolveVulnerableSymbolsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ResolveVulnerableSymbolsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ResolveVulnerableSymbolsRequestDataType(value)
	return nil
}

// NewResolveVulnerableSymbolsRequestDataTypeFromValue returns a pointer to a valid ResolveVulnerableSymbolsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewResolveVulnerableSymbolsRequestDataTypeFromValue(v string) (*ResolveVulnerableSymbolsRequestDataType, error) {
	ev := ResolveVulnerableSymbolsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ResolveVulnerableSymbolsRequestDataType: valid values are %v", v, allowedResolveVulnerableSymbolsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ResolveVulnerableSymbolsRequestDataType) IsValid() bool {
	for _, existing := range allowedResolveVulnerableSymbolsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResolveVulnerableSymbolsRequestDataType value.
func (v ResolveVulnerableSymbolsRequestDataType) Ptr() *ResolveVulnerableSymbolsRequestDataType {
	return &v
}
