// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ShiftDataAttributes Attributes for an on-call shift.
type ShiftDataAttributes struct {
	// The end time of the shift.
	End *time.Time `json:"end,omitempty"`
	// The start time of the shift.
	Start *time.Time `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewShiftDataAttributes instantiates a new ShiftDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewShiftDataAttributes() *ShiftDataAttributes {
	this := ShiftDataAttributes{}
	return &this
}

// NewShiftDataAttributesWithDefaults instantiates a new ShiftDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewShiftDataAttributesWithDefaults() *ShiftDataAttributes {
	this := ShiftDataAttributes{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ShiftDataAttributes) GetEnd() time.Time {
	if o == nil || o.End == nil {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShiftDataAttributes) GetEndOk() (*time.Time, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ShiftDataAttributes) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *ShiftDataAttributes) SetEnd(v time.Time) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ShiftDataAttributes) GetStart() time.Time {
	if o == nil || o.Start == nil {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShiftDataAttributes) GetStartOk() (*time.Time, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ShiftDataAttributes) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *ShiftDataAttributes) SetStart(v time.Time) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ShiftDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		if o.End.Nanosecond() == 0 {
			toSerialize["end"] = o.End.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end"] = o.End.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Start != nil {
		if o.Start.Nanosecond() == 0 {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ShiftDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *time.Time `json:"end,omitempty"`
		Start *time.Time `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start"})
	} else {
		return err
	}
	o.End = all.End
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
