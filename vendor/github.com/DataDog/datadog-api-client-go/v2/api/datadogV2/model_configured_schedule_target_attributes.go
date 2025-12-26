// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfiguredScheduleTargetAttributes Attributes for a configured schedule target, including position.
type ConfiguredScheduleTargetAttributes struct {
	// Specifies the position of a schedule target (example `previous`, `current`, or `next`).
	Position ScheduleTargetPosition `json:"position"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfiguredScheduleTargetAttributes instantiates a new ConfiguredScheduleTargetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfiguredScheduleTargetAttributes(position ScheduleTargetPosition) *ConfiguredScheduleTargetAttributes {
	this := ConfiguredScheduleTargetAttributes{}
	this.Position = position
	return &this
}

// NewConfiguredScheduleTargetAttributesWithDefaults instantiates a new ConfiguredScheduleTargetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfiguredScheduleTargetAttributesWithDefaults() *ConfiguredScheduleTargetAttributes {
	this := ConfiguredScheduleTargetAttributes{}
	return &this
}

// GetPosition returns the Position field value.
func (o *ConfiguredScheduleTargetAttributes) GetPosition() ScheduleTargetPosition {
	if o == nil {
		var ret ScheduleTargetPosition
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *ConfiguredScheduleTargetAttributes) GetPositionOk() (*ScheduleTargetPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *ConfiguredScheduleTargetAttributes) SetPosition(v ScheduleTargetPosition) {
	o.Position = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfiguredScheduleTargetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["position"] = o.Position

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfiguredScheduleTargetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Position *ScheduleTargetPosition `json:"position"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Position == nil {
		return fmt.Errorf("required field position missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"position"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Position.IsValid() {
		hasInvalidField = true
	} else {
		o.Position = *all.Position
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
