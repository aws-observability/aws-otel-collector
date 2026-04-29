// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestOptionsScheduling Object containing timeframes and timezone used for advanced scheduling.
type SyntheticsTestOptionsScheduling struct {
	// Array containing objects describing the scheduling pattern to apply to each day.
	Timeframes []SyntheticsTestOptionsSchedulingTimeframe `json:"timeframes"`
	// Timezone in which the timeframe is based.
	Timezone string `json:"timezone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestOptionsScheduling instantiates a new SyntheticsTestOptionsScheduling object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestOptionsScheduling(timeframes []SyntheticsTestOptionsSchedulingTimeframe, timezone string) *SyntheticsTestOptionsScheduling {
	this := SyntheticsTestOptionsScheduling{}
	this.Timeframes = timeframes
	this.Timezone = timezone
	return &this
}

// NewSyntheticsTestOptionsSchedulingWithDefaults instantiates a new SyntheticsTestOptionsScheduling object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestOptionsSchedulingWithDefaults() *SyntheticsTestOptionsScheduling {
	this := SyntheticsTestOptionsScheduling{}
	return &this
}

// GetTimeframes returns the Timeframes field value.
func (o *SyntheticsTestOptionsScheduling) GetTimeframes() []SyntheticsTestOptionsSchedulingTimeframe {
	if o == nil {
		var ret []SyntheticsTestOptionsSchedulingTimeframe
		return ret
	}
	return o.Timeframes
}

// GetTimeframesOk returns a tuple with the Timeframes field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsScheduling) GetTimeframesOk() (*[]SyntheticsTestOptionsSchedulingTimeframe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframes, true
}

// SetTimeframes sets field value.
func (o *SyntheticsTestOptionsScheduling) SetTimeframes(v []SyntheticsTestOptionsSchedulingTimeframe) {
	o.Timeframes = v
}

// GetTimezone returns the Timezone field value.
func (o *SyntheticsTestOptionsScheduling) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsScheduling) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *SyntheticsTestOptionsScheduling) SetTimezone(v string) {
	o.Timezone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestOptionsScheduling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["timeframes"] = o.Timeframes
	toSerialize["timezone"] = o.Timezone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestOptionsScheduling) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Timeframes *[]SyntheticsTestOptionsSchedulingTimeframe `json:"timeframes"`
		Timezone   *string                                     `json:"timezone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Timeframes == nil {
		return fmt.Errorf("required field timeframes missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"timeframes", "timezone"})
	} else {
		return err
	}
	o.Timeframes = *all.Timeframes
	o.Timezone = *all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
