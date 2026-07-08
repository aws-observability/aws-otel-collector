// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsProvider Cloud provider for commitment programs.
type CommitmentsProvider string

// List of CommitmentsProvider.
const (
	COMMITMENTSPROVIDER_AWS   CommitmentsProvider = "aws"
	COMMITMENTSPROVIDER_AZURE CommitmentsProvider = "azure"
)

var allowedCommitmentsProviderEnumValues = []CommitmentsProvider{
	COMMITMENTSPROVIDER_AWS,
	COMMITMENTSPROVIDER_AZURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsProvider) GetAllowedValues() []CommitmentsProvider {
	return allowedCommitmentsProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsProvider(value)
	return nil
}

// NewCommitmentsProviderFromValue returns a pointer to a valid CommitmentsProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsProviderFromValue(v string) (*CommitmentsProvider, error) {
	ev := CommitmentsProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsProvider: valid values are %v", v, allowedCommitmentsProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsProvider) IsValid() bool {
	for _, existing := range allowedCommitmentsProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsProvider value.
func (v CommitmentsProvider) Ptr() *CommitmentsProvider {
	return &v
}
