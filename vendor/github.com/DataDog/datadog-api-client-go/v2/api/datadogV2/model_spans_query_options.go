// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansQueryOptions Global query options that are used during the query.
// Note: You should only supply timezone or time offset but not both otherwise the query will fail.
type SpansQueryOptions struct {
	// The time offset (in seconds) to apply to the query.
	TimeOffset *int64 `json:"timeOffset,omitempty"`
	// The timezone can be specified as GMT, UTC, an offset from UTC (like UTC+1), or as a Timezone Database identifier (like America/New_York).
	Timezone *string `json:"timezone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansQueryOptions instantiates a new SpansQueryOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansQueryOptions() *SpansQueryOptions {
	this := SpansQueryOptions{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// NewSpansQueryOptionsWithDefaults instantiates a new SpansQueryOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansQueryOptionsWithDefaults() *SpansQueryOptions {
	this := SpansQueryOptions{}
	var timezone string = "UTC"
	this.Timezone = &timezone
	return &this
}

// GetTimeOffset returns the TimeOffset field value if set, zero value otherwise.
func (o *SpansQueryOptions) GetTimeOffset() int64 {
	if o == nil || o.TimeOffset == nil {
		var ret int64
		return ret
	}
	return *o.TimeOffset
}

// GetTimeOffsetOk returns a tuple with the TimeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansQueryOptions) GetTimeOffsetOk() (*int64, bool) {
	if o == nil || o.TimeOffset == nil {
		return nil, false
	}
	return o.TimeOffset, true
}

// HasTimeOffset returns a boolean if a field has been set.
func (o *SpansQueryOptions) HasTimeOffset() bool {
	return o != nil && o.TimeOffset != nil
}

// SetTimeOffset gets a reference to the given int64 and assigns it to the TimeOffset field.
func (o *SpansQueryOptions) SetTimeOffset(v int64) {
	o.TimeOffset = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SpansQueryOptions) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansQueryOptions) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SpansQueryOptions) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *SpansQueryOptions) SetTimezone(v string) {
	o.Timezone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansQueryOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TimeOffset != nil {
		toSerialize["timeOffset"] = o.TimeOffset
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansQueryOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimeOffset *int64  `json:"timeOffset,omitempty"`
		Timezone   *string `json:"timezone,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"timeOffset", "timezone"})
	} else {
		return err
	}
	o.TimeOffset = all.TimeOffset
	o.Timezone = all.Timezone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
