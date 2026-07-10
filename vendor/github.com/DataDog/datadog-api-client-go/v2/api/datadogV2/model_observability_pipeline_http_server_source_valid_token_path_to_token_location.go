// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation Built-in token location on the incoming HTTP request.
type ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation string

// List of ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation.
const (
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEVALIDTOKENPATHTOTOKENLOCATION_PATH    ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation = "path"
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEVALIDTOKENPATHTOTOKENLOCATION_ADDRESS ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation = "address"
)

var allowedObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationEnumValues = []ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation{
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEVALIDTOKENPATHTOTOKENLOCATION_PATH,
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEVALIDTOKENPATHTOTOKENLOCATION_ADDRESS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation) GetAllowedValues() []ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation {
	return allowedObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation(value)
	return nil
}

// NewObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationFromValue returns a pointer to a valid ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationFromValue(v string) (*ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation, error) {
	ev := ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation: valid values are %v", v, allowedObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation value.
func (v ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation) Ptr() *ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation {
	return &v
}
