// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAggregateResponseAttributes Attributes of the aggregation result, including the total count across all groups and the per-group breakdowns.
type CaseAggregateResponseAttributes struct {
	// Aggregated groups.
	Groups []CaseAggregateGroup `json:"groups"`
	// Total count of aggregated cases.
	Total float64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseAggregateResponseAttributes instantiates a new CaseAggregateResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseAggregateResponseAttributes(groups []CaseAggregateGroup, total float64) *CaseAggregateResponseAttributes {
	this := CaseAggregateResponseAttributes{}
	this.Groups = groups
	this.Total = total
	return &this
}

// NewCaseAggregateResponseAttributesWithDefaults instantiates a new CaseAggregateResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseAggregateResponseAttributesWithDefaults() *CaseAggregateResponseAttributes {
	this := CaseAggregateResponseAttributes{}
	return &this
}

// GetGroups returns the Groups field value.
func (o *CaseAggregateResponseAttributes) GetGroups() []CaseAggregateGroup {
	if o == nil {
		var ret []CaseAggregateGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateResponseAttributes) GetGroupsOk() (*[]CaseAggregateGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value.
func (o *CaseAggregateResponseAttributes) SetGroups(v []CaseAggregateGroup) {
	o.Groups = v
}

// GetTotal returns the Total field value.
func (o *CaseAggregateResponseAttributes) GetTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateResponseAttributes) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *CaseAggregateResponseAttributes) SetTotal(v float64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseAggregateResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["groups"] = o.Groups
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseAggregateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Groups *[]CaseAggregateGroup `json:"groups"`
		Total  *float64              `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Groups == nil {
		return fmt.Errorf("required field groups missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"groups", "total"})
	} else {
		return err
	}
	o.Groups = *all.Groups
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
