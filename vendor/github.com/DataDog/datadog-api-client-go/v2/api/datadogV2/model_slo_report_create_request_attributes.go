// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloReportCreateRequestAttributes The attributes portion of the SLO report request.
type SloReportCreateRequestAttributes struct {
	// The `from` timestamp for the report in epoch seconds.
	FromTs int64 `json:"from_ts"`
	// The frequency at which report data is to be generated.
	Interval *SLOReportInterval `json:"interval,omitempty"`
	// The query string used to filter SLO results. Some examples of queries include `service:<service-name>` and `slo-name`.
	Query string `json:"query"`
	// The timezone used to determine the start and end of each interval. For example, weekly intervals start at 12am on Sunday in the specified timezone.
	Timezone *string `json:"timezone,omitempty"`
	// The `to` timestamp for the report in epoch seconds.
	ToTs int64 `json:"to_ts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSloReportCreateRequestAttributes instantiates a new SloReportCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSloReportCreateRequestAttributes(fromTs int64, query string, toTs int64) *SloReportCreateRequestAttributes {
	this := SloReportCreateRequestAttributes{}
	this.FromTs = fromTs
	this.Query = query
	this.ToTs = toTs
	return &this
}

// NewSloReportCreateRequestAttributesWithDefaults instantiates a new SloReportCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSloReportCreateRequestAttributesWithDefaults() *SloReportCreateRequestAttributes {
	this := SloReportCreateRequestAttributes{}
	return &this
}

// GetFromTs returns the FromTs field value.
func (o *SloReportCreateRequestAttributes) GetFromTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value
// and a boolean to check if the value has been set.
func (o *SloReportCreateRequestAttributes) GetFromTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTs, true
}

// SetFromTs sets field value.
func (o *SloReportCreateRequestAttributes) SetFromTs(v int64) {
	o.FromTs = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *SloReportCreateRequestAttributes) GetInterval() SLOReportInterval {
	if o == nil || o.Interval == nil {
		var ret SLOReportInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloReportCreateRequestAttributes) GetIntervalOk() (*SLOReportInterval, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *SloReportCreateRequestAttributes) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given SLOReportInterval and assigns it to the Interval field.
func (o *SloReportCreateRequestAttributes) SetInterval(v SLOReportInterval) {
	o.Interval = &v
}

// GetQuery returns the Query field value.
func (o *SloReportCreateRequestAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SloReportCreateRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SloReportCreateRequestAttributes) SetQuery(v string) {
	o.Query = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *SloReportCreateRequestAttributes) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloReportCreateRequestAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *SloReportCreateRequestAttributes) HasTimezone() bool {
	return o != nil && o.Timezone != nil
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *SloReportCreateRequestAttributes) SetTimezone(v string) {
	o.Timezone = &v
}

// GetToTs returns the ToTs field value.
func (o *SloReportCreateRequestAttributes) GetToTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value
// and a boolean to check if the value has been set.
func (o *SloReportCreateRequestAttributes) GetToTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTs, true
}

// SetToTs sets field value.
func (o *SloReportCreateRequestAttributes) SetToTs(v int64) {
	o.ToTs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SloReportCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from_ts"] = o.FromTs
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	toSerialize["query"] = o.Query
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	toSerialize["to_ts"] = o.ToTs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SloReportCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromTs   *int64             `json:"from_ts"`
		Interval *SLOReportInterval `json:"interval,omitempty"`
		Query    *string            `json:"query"`
		Timezone *string            `json:"timezone,omitempty"`
		ToTs     *int64             `json:"to_ts"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FromTs == nil {
		return fmt.Errorf("required field from_ts missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.ToTs == nil {
		return fmt.Errorf("required field to_ts missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from_ts", "interval", "query", "timezone", "to_ts"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FromTs = *all.FromTs
	if all.Interval != nil && !all.Interval.IsValid() {
		hasInvalidField = true
	} else {
		o.Interval = all.Interval
	}
	o.Query = *all.Query
	o.Timezone = all.Timezone
	o.ToTs = *all.ToTs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
