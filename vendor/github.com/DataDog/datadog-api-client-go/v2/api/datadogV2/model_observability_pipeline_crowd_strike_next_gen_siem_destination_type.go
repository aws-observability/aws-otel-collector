// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType The destination type. The value should always be `crowdstrike_next_gen_siem`.
type ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType string

// List of ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType.
const (
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONTYPE_CROWDSTRIKE_NEXT_GEN_SIEM ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType = "crowdstrike_next_gen_siem"
)

var allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationTypeEnumValues = []ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType{
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONTYPE_CROWDSTRIKE_NEXT_GEN_SIEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType) GetAllowedValues() []ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType {
	return allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType(value)
	return nil
}

// NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationTypeFromValue(v string) (*ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType, error) {
	ev := ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType: valid values are %v", v, allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType value.
func (v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType) Ptr() *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType {
	return &v
}
