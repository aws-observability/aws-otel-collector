// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarInterval Calendar interval definition.
type CalendarInterval struct {
	// Alignment of the interval. Valid values depend on the interval type. For `day`, use hours (for example, `1am`, `2pm`, or `14`). For `week`, use day names (for example, `monday`). For `month`, use day-of-month ordinals (for example, `1st`, `15th`). For `year` or `quarter`, use month names (for example, `january`).
	Alignment *string `json:"alignment,omitempty"`
	// Quantity of the interval.
	Quantity *int64 `json:"quantity,omitempty"`
	// Timezone for the interval.
	Timezone *string `json:"timezone,omitempty"`
	// Type of calendar interval.
	Type CalendarIntervalType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCalendarInterval instantiates a new CalendarInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalendarInterval(typeVar CalendarIntervalType) *CalendarInterval {
	this := CalendarInterval{}
	this.Type = typeVar
	return &this
}

// NewCalendarIntervalWithDefaults instantiates a new CalendarInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalendarIntervalWithDefaults() *CalendarInterval {
	this := CalendarInterval{}
	return &this
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *CalendarInterval) GetAlignment() string {
	if o == nil || o.Alignment == nil {
		var ret string
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetAlignmentOk() (*string, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *CalendarInterval) HasAlignment() bool {
	return o != nil && o.Alignment != nil
}

// SetAlignment gets a reference to the given string and assigns it to the Alignment field.
func (o *CalendarInterval) SetAlignment(v string) {
	o.Alignment = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CalendarInterval) GetQuantity() int64 {
	if o == nil || o.Quantity == nil {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetQuantityOk() (*int64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CalendarInterval) HasQuantity() bool {
	return o != nil && o.Quantity != nil
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *CalendarInterval) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CalendarInterval) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CalendarInterval) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CalendarInterval) SetTimezone(v string) {
	o.Timezone = &v
}

// GetType returns the Type field value.
func (o *CalendarInterval) GetType() CalendarIntervalType {
	if o == nil {
		var ret CalendarIntervalType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CalendarInterval) GetTypeOk() (*CalendarIntervalType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CalendarInterval) SetType(v CalendarIntervalType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalendarInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alignment != nil {
		toSerialize["alignment"] = o.Alignment
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CalendarInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alignment *string               `json:"alignment,omitempty"`
		Quantity  *int64                `json:"quantity,omitempty"`
		Timezone  *string               `json:"timezone,omitempty"`
		Type      *CalendarIntervalType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Alignment = all.Alignment
	o.Quantity = all.Quantity
	o.Timezone = all.Timezone
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
