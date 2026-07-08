// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseCountGroup A facet group containing counts broken down by the distinct values of a case field (for example, status or priority).
type CaseCountGroup struct {
	// The name of the field being grouped on (for example, `status` or `priority`).
	Group string `json:"group"`
	// Values within this group.
	GroupValues []CaseCountGroupValue `json:"group_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseCountGroup instantiates a new CaseCountGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseCountGroup(group string, groupValues []CaseCountGroupValue) *CaseCountGroup {
	this := CaseCountGroup{}
	this.Group = group
	this.GroupValues = groupValues
	return &this
}

// NewCaseCountGroupWithDefaults instantiates a new CaseCountGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseCountGroupWithDefaults() *CaseCountGroup {
	this := CaseCountGroup{}
	return &this
}

// GetGroup returns the Group field value.
func (o *CaseCountGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *CaseCountGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *CaseCountGroup) SetGroup(v string) {
	o.Group = v
}

// GetGroupValues returns the GroupValues field value.
func (o *CaseCountGroup) GetGroupValues() []CaseCountGroupValue {
	if o == nil {
		var ret []CaseCountGroupValue
		return ret
	}
	return o.GroupValues
}

// GetGroupValuesOk returns a tuple with the GroupValues field value
// and a boolean to check if the value has been set.
func (o *CaseCountGroup) GetGroupValuesOk() (*[]CaseCountGroupValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupValues, true
}

// SetGroupValues sets field value.
func (o *CaseCountGroup) SetGroupValues(v []CaseCountGroupValue) {
	o.GroupValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseCountGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group"] = o.Group
	toSerialize["group_values"] = o.GroupValues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseCountGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group       *string                `json:"group"`
		GroupValues *[]CaseCountGroupValue `json:"group_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.GroupValues == nil {
		return fmt.Errorf("required field group_values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group", "group_values"})
	} else {
		return err
	}
	o.Group = *all.Group
	o.GroupValues = *all.GroupValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
