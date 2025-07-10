// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeRestrictions Holds time zone information and a list of time restrictions for a routing rule.
type TimeRestrictions struct {
	// Defines the list of time-based restrictions.
	Restrictions []TimeRestriction `json:"restrictions"`
	// Specifies the time zone applicable to the restrictions.
	TimeZone string `json:"time_zone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeRestrictions instantiates a new TimeRestrictions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeRestrictions(restrictions []TimeRestriction, timeZone string) *TimeRestrictions {
	this := TimeRestrictions{}
	this.Restrictions = restrictions
	this.TimeZone = timeZone
	return &this
}

// NewTimeRestrictionsWithDefaults instantiates a new TimeRestrictions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeRestrictionsWithDefaults() *TimeRestrictions {
	this := TimeRestrictions{}
	return &this
}

// GetRestrictions returns the Restrictions field value.
func (o *TimeRestrictions) GetRestrictions() []TimeRestriction {
	if o == nil {
		var ret []TimeRestriction
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value
// and a boolean to check if the value has been set.
func (o *TimeRestrictions) GetRestrictionsOk() (*[]TimeRestriction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Restrictions, true
}

// SetRestrictions sets field value.
func (o *TimeRestrictions) SetRestrictions(v []TimeRestriction) {
	o.Restrictions = v
}

// GetTimeZone returns the TimeZone field value.
func (o *TimeRestrictions) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *TimeRestrictions) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value.
func (o *TimeRestrictions) SetTimeZone(v string) {
	o.TimeZone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["restrictions"] = o.Restrictions
	toSerialize["time_zone"] = o.TimeZone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeRestrictions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Restrictions *[]TimeRestriction `json:"restrictions"`
		TimeZone     *string            `json:"time_zone"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Restrictions == nil {
		return fmt.Errorf("required field restrictions missing")
	}
	if all.TimeZone == nil {
		return fmt.Errorf("required field time_zone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"restrictions", "time_zone"})
	} else {
		return err
	}
	o.Restrictions = *all.Restrictions
	o.TimeZone = *all.TimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
