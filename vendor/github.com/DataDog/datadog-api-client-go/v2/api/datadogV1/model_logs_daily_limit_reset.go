// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsDailyLimitReset Object containing options to override the default daily limit reset time.
type LogsDailyLimitReset struct {
	// String in `HH:00` format representing the time of day the daily limit should be reset. The hours must be between 00 and 23 (inclusive).
	ResetTime *string `json:"reset_time,omitempty"`
	// String in `(-|+)HH:00` format representing the UTC offset to apply to the given reset time. The hours must be between -12 and +14 (inclusive).
	ResetUtcOffset *string `json:"reset_utc_offset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsDailyLimitReset instantiates a new LogsDailyLimitReset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsDailyLimitReset() *LogsDailyLimitReset {
	this := LogsDailyLimitReset{}
	return &this
}

// NewLogsDailyLimitResetWithDefaults instantiates a new LogsDailyLimitReset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsDailyLimitResetWithDefaults() *LogsDailyLimitReset {
	this := LogsDailyLimitReset{}
	return &this
}

// GetResetTime returns the ResetTime field value if set, zero value otherwise.
func (o *LogsDailyLimitReset) GetResetTime() string {
	if o == nil || o.ResetTime == nil {
		var ret string
		return ret
	}
	return *o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsDailyLimitReset) GetResetTimeOk() (*string, bool) {
	if o == nil || o.ResetTime == nil {
		return nil, false
	}
	return o.ResetTime, true
}

// HasResetTime returns a boolean if a field has been set.
func (o *LogsDailyLimitReset) HasResetTime() bool {
	return o != nil && o.ResetTime != nil
}

// SetResetTime gets a reference to the given string and assigns it to the ResetTime field.
func (o *LogsDailyLimitReset) SetResetTime(v string) {
	o.ResetTime = &v
}

// GetResetUtcOffset returns the ResetUtcOffset field value if set, zero value otherwise.
func (o *LogsDailyLimitReset) GetResetUtcOffset() string {
	if o == nil || o.ResetUtcOffset == nil {
		var ret string
		return ret
	}
	return *o.ResetUtcOffset
}

// GetResetUtcOffsetOk returns a tuple with the ResetUtcOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsDailyLimitReset) GetResetUtcOffsetOk() (*string, bool) {
	if o == nil || o.ResetUtcOffset == nil {
		return nil, false
	}
	return o.ResetUtcOffset, true
}

// HasResetUtcOffset returns a boolean if a field has been set.
func (o *LogsDailyLimitReset) HasResetUtcOffset() bool {
	return o != nil && o.ResetUtcOffset != nil
}

// SetResetUtcOffset gets a reference to the given string and assigns it to the ResetUtcOffset field.
func (o *LogsDailyLimitReset) SetResetUtcOffset(v string) {
	o.ResetUtcOffset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsDailyLimitReset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ResetTime != nil {
		toSerialize["reset_time"] = o.ResetTime
	}
	if o.ResetUtcOffset != nil {
		toSerialize["reset_utc_offset"] = o.ResetUtcOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsDailyLimitReset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResetTime      *string `json:"reset_time,omitempty"`
		ResetUtcOffset *string `json:"reset_utc_offset,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reset_time", "reset_utc_offset"})
	} else {
		return err
	}
	o.ResetTime = all.ResetTime
	o.ResetUtcOffset = all.ResetUtcOffset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
