// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributes Defines the updatable attributes for a schedule, such as name, time zone, and layers.
type ScheduleUpdateRequestDataAttributes struct {
	// The updated list of layers (rotations) for this schedule.
	Layers []ScheduleUpdateRequestDataAttributesLayersItems `json:"layers"`
	// A short name for the schedule.
	Name string `json:"name"`
	// The time zone used when interpreting rotation times.
	TimeZone string `json:"time_zone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestDataAttributes instantiates a new ScheduleUpdateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestDataAttributes(layers []ScheduleUpdateRequestDataAttributesLayersItems, name string, timeZone string) *ScheduleUpdateRequestDataAttributes {
	this := ScheduleUpdateRequestDataAttributes{}
	this.Layers = layers
	this.Name = name
	this.TimeZone = timeZone
	return &this
}

// NewScheduleUpdateRequestDataAttributesWithDefaults instantiates a new ScheduleUpdateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataAttributesWithDefaults() *ScheduleUpdateRequestDataAttributes {
	this := ScheduleUpdateRequestDataAttributes{}
	return &this
}

// GetLayers returns the Layers field value.
func (o *ScheduleUpdateRequestDataAttributes) GetLayers() []ScheduleUpdateRequestDataAttributesLayersItems {
	if o == nil {
		var ret []ScheduleUpdateRequestDataAttributesLayersItems
		return ret
	}
	return o.Layers
}

// GetLayersOk returns a tuple with the Layers field value
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributes) GetLayersOk() (*[]ScheduleUpdateRequestDataAttributesLayersItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layers, true
}

// SetLayers sets field value.
func (o *ScheduleUpdateRequestDataAttributes) SetLayers(v []ScheduleUpdateRequestDataAttributesLayersItems) {
	o.Layers = v
}

// GetName returns the Name field value.
func (o *ScheduleUpdateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ScheduleUpdateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetTimeZone returns the TimeZone field value.
func (o *ScheduleUpdateRequestDataAttributes) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributes) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value.
func (o *ScheduleUpdateRequestDataAttributes) SetTimeZone(v string) {
	o.TimeZone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["layers"] = o.Layers
	toSerialize["name"] = o.Name
	toSerialize["time_zone"] = o.TimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleUpdateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Layers   *[]ScheduleUpdateRequestDataAttributesLayersItems `json:"layers"`
		Name     *string                                           `json:"name"`
		TimeZone *string                                           `json:"time_zone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Layers == nil {
		return fmt.Errorf("required field layers missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TimeZone == nil {
		return fmt.Errorf("required field time_zone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"layers", "name", "time_zone"})
	} else {
		return err
	}
	o.Layers = *all.Layers
	o.Name = *all.Name
	o.TimeZone = *all.TimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
