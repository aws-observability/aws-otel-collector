// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataAttributes Provides core properties of a schedule object such as its name and time zone.
type ScheduleDataAttributes struct {
	// A short name for the schedule.
	Name *string `json:"name,omitempty"`
	// The time zone in which this schedule operates.
	TimeZone *string `json:"time_zone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleDataAttributes instantiates a new ScheduleDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleDataAttributes() *ScheduleDataAttributes {
	this := ScheduleDataAttributes{}
	return &this
}

// NewScheduleDataAttributesWithDefaults instantiates a new ScheduleDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleDataAttributesWithDefaults() *ScheduleDataAttributes {
	this := ScheduleDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScheduleDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScheduleDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScheduleDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ScheduleDataAttributes) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDataAttributes) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ScheduleDataAttributes) HasTimeZone() bool {
	return o != nil && o.TimeZone != nil
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ScheduleDataAttributes) SetTimeZone(v string) {
	o.TimeZone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TimeZone != nil {
		toSerialize["time_zone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name,omitempty"`
		TimeZone *string `json:"time_zone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "time_zone"})
	} else {
		return err
	}
	o.Name = all.Name
	o.TimeZone = all.TimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
