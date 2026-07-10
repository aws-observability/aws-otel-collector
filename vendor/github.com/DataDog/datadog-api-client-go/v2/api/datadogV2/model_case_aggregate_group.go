// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAggregateGroup A single group within the aggregation results, containing the group key and its associated count values.
type CaseAggregateGroup struct {
	// The value of the field being grouped on (for example, `OPEN` when grouping by status).
	Group string `json:"group"`
	// The count of cases in this group.
	Value []float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseAggregateGroup instantiates a new CaseAggregateGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseAggregateGroup(group string, value []float64) *CaseAggregateGroup {
	this := CaseAggregateGroup{}
	this.Group = group
	this.Value = value
	return &this
}

// NewCaseAggregateGroupWithDefaults instantiates a new CaseAggregateGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseAggregateGroupWithDefaults() *CaseAggregateGroup {
	this := CaseAggregateGroup{}
	return &this
}

// GetGroup returns the Group field value.
func (o *CaseAggregateGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *CaseAggregateGroup) SetGroup(v string) {
	o.Group = v
}

// GetValue returns the Value field value.
func (o *CaseAggregateGroup) GetValue() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateGroup) GetValueOk() (*[]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *CaseAggregateGroup) SetValue(v []float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseAggregateGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group"] = o.Group
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseAggregateGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group *string    `json:"group"`
		Value *[]float64 `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group", "value"})
	} else {
		return err
	}
	o.Group = *all.Group
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
