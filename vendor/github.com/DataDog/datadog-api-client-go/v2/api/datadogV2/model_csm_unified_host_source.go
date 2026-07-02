// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmUnifiedHostSource The source of a unified host entry, indicating whether it was discovered via agent, agentless scanning, or both.
type CsmUnifiedHostSource string

// List of CsmUnifiedHostSource.
const (
	CSMUNIFIEDHOSTSOURCE_AGENT     CsmUnifiedHostSource = "agent"
	CSMUNIFIEDHOSTSOURCE_AGENTLESS CsmUnifiedHostSource = "agentless"
	CSMUNIFIEDHOSTSOURCE_BOTH      CsmUnifiedHostSource = "both"
)

var allowedCsmUnifiedHostSourceEnumValues = []CsmUnifiedHostSource{
	CSMUNIFIEDHOSTSOURCE_AGENT,
	CSMUNIFIEDHOSTSOURCE_AGENTLESS,
	CSMUNIFIEDHOSTSOURCE_BOTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CsmUnifiedHostSource) GetAllowedValues() []CsmUnifiedHostSource {
	return allowedCsmUnifiedHostSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CsmUnifiedHostSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CsmUnifiedHostSource(value)
	return nil
}

// NewCsmUnifiedHostSourceFromValue returns a pointer to a valid CsmUnifiedHostSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCsmUnifiedHostSourceFromValue(v string) (*CsmUnifiedHostSource, error) {
	ev := CsmUnifiedHostSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CsmUnifiedHostSource: valid values are %v", v, allowedCsmUnifiedHostSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CsmUnifiedHostSource) IsValid() bool {
	for _, existing := range allowedCsmUnifiedHostSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CsmUnifiedHostSource value.
func (v CsmUnifiedHostSource) Ptr() *CsmUnifiedHostSource {
	return &v
}
