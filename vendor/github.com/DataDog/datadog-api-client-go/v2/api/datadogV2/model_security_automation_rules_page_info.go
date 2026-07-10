// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAutomationRulesPageInfo Pagination information for the list of automation rules.
type SecurityAutomationRulesPageInfo struct {
	// The total number of rules matching the current filter.
	TotalFilteredCount int64 `json:"total_filtered_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAutomationRulesPageInfo instantiates a new SecurityAutomationRulesPageInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAutomationRulesPageInfo(totalFilteredCount int64) *SecurityAutomationRulesPageInfo {
	this := SecurityAutomationRulesPageInfo{}
	this.TotalFilteredCount = totalFilteredCount
	return &this
}

// NewSecurityAutomationRulesPageInfoWithDefaults instantiates a new SecurityAutomationRulesPageInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAutomationRulesPageInfoWithDefaults() *SecurityAutomationRulesPageInfo {
	this := SecurityAutomationRulesPageInfo{}
	return &this
}

// GetTotalFilteredCount returns the TotalFilteredCount field value.
func (o *SecurityAutomationRulesPageInfo) GetTotalFilteredCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value
// and a boolean to check if the value has been set.
func (o *SecurityAutomationRulesPageInfo) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFilteredCount, true
}

// SetTotalFilteredCount sets field value.
func (o *SecurityAutomationRulesPageInfo) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAutomationRulesPageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["total_filtered_count"] = o.TotalFilteredCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAutomationRulesPageInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalFilteredCount *int64 `json:"total_filtered_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalFilteredCount == nil {
		return fmt.Errorf("required field total_filtered_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_filtered_count"})
	} else {
		return err
	}
	o.TotalFilteredCount = *all.TotalFilteredCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
