// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAggregateGroupBy Configuration for grouping aggregated results by one or more case fields.
type CaseAggregateGroupBy struct {
	// Fields to group by.
	Groups []string `json:"groups"`
	// Maximum number of groups to return.
	Limit int32 `json:"limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseAggregateGroupBy instantiates a new CaseAggregateGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseAggregateGroupBy(groups []string, limit int32) *CaseAggregateGroupBy {
	this := CaseAggregateGroupBy{}
	this.Groups = groups
	this.Limit = limit
	return &this
}

// NewCaseAggregateGroupByWithDefaults instantiates a new CaseAggregateGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseAggregateGroupByWithDefaults() *CaseAggregateGroupBy {
	this := CaseAggregateGroupBy{}
	return &this
}

// GetGroups returns the Groups field value.
func (o *CaseAggregateGroupBy) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateGroupBy) GetGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value.
func (o *CaseAggregateGroupBy) SetGroups(v []string) {
	o.Groups = v
}

// GetLimit returns the Limit field value.
func (o *CaseAggregateGroupBy) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateGroupBy) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *CaseAggregateGroupBy) SetLimit(v int32) {
	o.Limit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseAggregateGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["groups"] = o.Groups
	toSerialize["limit"] = o.Limit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseAggregateGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Groups *[]string `json:"groups"`
		Limit  *int32    `json:"limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Groups == nil {
		return fmt.Errorf("required field groups missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"groups", "limit"})
	} else {
		return err
	}
	o.Groups = *all.Groups
	o.Limit = *all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
