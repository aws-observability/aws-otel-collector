// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Schedule Top-level container for a schedule object, including both the `data` payload and any related `included` resources (such as teams, layers, or members).
type Schedule struct {
	// Represents the primary data object for a schedule, linking attributes and relationships.
	Data *ScheduleData `json:"data,omitempty"`
	// Any additional resources related to this schedule, such as teams and layers.
	Included []ScheduleDataIncludedItem `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSchedule instantiates a new Schedule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSchedule() *Schedule {
	this := Schedule{}
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Schedule) GetData() ScheduleData {
	if o == nil || o.Data == nil {
		var ret ScheduleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDataOk() (*ScheduleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Schedule) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given ScheduleData and assigns it to the Data field.
func (o *Schedule) SetData(v ScheduleData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *Schedule) GetIncluded() []ScheduleDataIncludedItem {
	if o == nil || o.Included == nil {
		var ret []ScheduleDataIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetIncludedOk() (*[]ScheduleDataIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *Schedule) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []ScheduleDataIncludedItem and assigns it to the Included field.
func (o *Schedule) SetIncluded(v []ScheduleDataIncludedItem) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Schedule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *ScheduleData              `json:"data,omitempty"`
		Included []ScheduleDataIncludedItem `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
