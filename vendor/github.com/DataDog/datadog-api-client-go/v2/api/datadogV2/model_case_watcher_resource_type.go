// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseWatcherResourceType JSON:API resource type for case watchers.
type CaseWatcherResourceType string

// List of CaseWatcherResourceType.
const (
	CASEWATCHERRESOURCETYPE_WATCHER CaseWatcherResourceType = "watcher"
)

var allowedCaseWatcherResourceTypeEnumValues = []CaseWatcherResourceType{
	CASEWATCHERRESOURCETYPE_WATCHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseWatcherResourceType) GetAllowedValues() []CaseWatcherResourceType {
	return allowedCaseWatcherResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseWatcherResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseWatcherResourceType(value)
	return nil
}

// NewCaseWatcherResourceTypeFromValue returns a pointer to a valid CaseWatcherResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseWatcherResourceTypeFromValue(v string) (*CaseWatcherResourceType, error) {
	ev := CaseWatcherResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseWatcherResourceType: valid values are %v", v, allowedCaseWatcherResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseWatcherResourceType) IsValid() bool {
	for _, existing := range allowedCaseWatcherResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseWatcherResourceType value.
func (v CaseWatcherResourceType) Ptr() *CaseWatcherResourceType {
	return &v
}
