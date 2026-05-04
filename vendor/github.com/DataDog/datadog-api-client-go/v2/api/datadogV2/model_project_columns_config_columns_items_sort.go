// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectColumnsConfigColumnsItemsSort Sort configuration for a project board column.
type ProjectColumnsConfigColumnsItemsSort struct {
	// Whether to sort in ascending order.
	Ascending *bool `json:"ascending,omitempty"`
	// The sort priority order for this column.
	Priority *int64 `json:"priority,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectColumnsConfigColumnsItemsSort instantiates a new ProjectColumnsConfigColumnsItemsSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectColumnsConfigColumnsItemsSort() *ProjectColumnsConfigColumnsItemsSort {
	this := ProjectColumnsConfigColumnsItemsSort{}
	return &this
}

// NewProjectColumnsConfigColumnsItemsSortWithDefaults instantiates a new ProjectColumnsConfigColumnsItemsSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectColumnsConfigColumnsItemsSortWithDefaults() *ProjectColumnsConfigColumnsItemsSort {
	this := ProjectColumnsConfigColumnsItemsSort{}
	return &this
}

// GetAscending returns the Ascending field value if set, zero value otherwise.
func (o *ProjectColumnsConfigColumnsItemsSort) GetAscending() bool {
	if o == nil || o.Ascending == nil {
		var ret bool
		return ret
	}
	return *o.Ascending
}

// GetAscendingOk returns a tuple with the Ascending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectColumnsConfigColumnsItemsSort) GetAscendingOk() (*bool, bool) {
	if o == nil || o.Ascending == nil {
		return nil, false
	}
	return o.Ascending, true
}

// HasAscending returns a boolean if a field has been set.
func (o *ProjectColumnsConfigColumnsItemsSort) HasAscending() bool {
	return o != nil && o.Ascending != nil
}

// SetAscending gets a reference to the given bool and assigns it to the Ascending field.
func (o *ProjectColumnsConfigColumnsItemsSort) SetAscending(v bool) {
	o.Ascending = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ProjectColumnsConfigColumnsItemsSort) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectColumnsConfigColumnsItemsSort) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ProjectColumnsConfigColumnsItemsSort) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *ProjectColumnsConfigColumnsItemsSort) SetPriority(v int64) {
	o.Priority = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectColumnsConfigColumnsItemsSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ascending != nil {
		toSerialize["ascending"] = o.Ascending
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectColumnsConfigColumnsItemsSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ascending *bool  `json:"ascending,omitempty"`
		Priority  *int64 `json:"priority,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ascending", "priority"})
	} else {
		return err
	}
	o.Ascending = all.Ascending
	o.Priority = all.Priority

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
