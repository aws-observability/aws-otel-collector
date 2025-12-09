// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsUptime Object containing the uptime information.
type SyntheticsUptime struct {
	// An array of error objects returned while querying the history data for the service level objective.
	Errors []SLOHistoryResponseErrorWithType `json:"errors,omitempty"`
	// The location name
	Group *string `json:"group,omitempty"`
	// The state transition history for the monitor, represented as an array of
	// pairs. Each pair is an array where the first element is the transition timestamp
	// in Unix epoch format (integer) and the second element is the state (integer).
	// For the state, an integer value of `0` indicates uptime, `1` indicates downtime,
	// and `2` indicates no data.
	History [][]float64 `json:"history,omitempty"`
	// The number of decimal places to which the SLI value is accurate for the given from-to timestamps.
	SpanPrecision *float64 `json:"span_precision,omitempty"`
	// The overall uptime.
	Uptime *float64 `json:"uptime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsUptime instantiates a new SyntheticsUptime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsUptime() *SyntheticsUptime {
	this := SyntheticsUptime{}
	return &this
}

// NewSyntheticsUptimeWithDefaults instantiates a new SyntheticsUptime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsUptimeWithDefaults() *SyntheticsUptime {
	this := SyntheticsUptime{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyntheticsUptime) GetErrors() []SLOHistoryResponseErrorWithType {
	if o == nil {
		var ret []SLOHistoryResponseErrorWithType
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SyntheticsUptime) GetErrorsOk() (*[]SLOHistoryResponseErrorWithType, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SyntheticsUptime) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []SLOHistoryResponseErrorWithType and assigns it to the Errors field.
func (o *SyntheticsUptime) SetErrors(v []SLOHistoryResponseErrorWithType) {
	o.Errors = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SyntheticsUptime) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsUptime) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SyntheticsUptime) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *SyntheticsUptime) SetGroup(v string) {
	o.Group = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *SyntheticsUptime) GetHistory() [][]float64 {
	if o == nil || o.History == nil {
		var ret [][]float64
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsUptime) GetHistoryOk() (*[][]float64, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return &o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *SyntheticsUptime) HasHistory() bool {
	return o != nil && o.History != nil
}

// SetHistory gets a reference to the given [][]float64 and assigns it to the History field.
func (o *SyntheticsUptime) SetHistory(v [][]float64) {
	o.History = v
}

// GetSpanPrecision returns the SpanPrecision field value if set, zero value otherwise.
func (o *SyntheticsUptime) GetSpanPrecision() float64 {
	if o == nil || o.SpanPrecision == nil {
		var ret float64
		return ret
	}
	return *o.SpanPrecision
}

// GetSpanPrecisionOk returns a tuple with the SpanPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsUptime) GetSpanPrecisionOk() (*float64, bool) {
	if o == nil || o.SpanPrecision == nil {
		return nil, false
	}
	return o.SpanPrecision, true
}

// HasSpanPrecision returns a boolean if a field has been set.
func (o *SyntheticsUptime) HasSpanPrecision() bool {
	return o != nil && o.SpanPrecision != nil
}

// SetSpanPrecision gets a reference to the given float64 and assigns it to the SpanPrecision field.
func (o *SyntheticsUptime) SetSpanPrecision(v float64) {
	o.SpanPrecision = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *SyntheticsUptime) GetUptime() float64 {
	if o == nil || o.Uptime == nil {
		var ret float64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsUptime) GetUptimeOk() (*float64, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *SyntheticsUptime) HasUptime() bool {
	return o != nil && o.Uptime != nil
}

// SetUptime gets a reference to the given float64 and assigns it to the Uptime field.
func (o *SyntheticsUptime) SetUptime(v float64) {
	o.Uptime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsUptime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	if o.SpanPrecision != nil {
		toSerialize["span_precision"] = o.SpanPrecision
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsUptime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors        []SLOHistoryResponseErrorWithType `json:"errors,omitempty"`
		Group         *string                           `json:"group,omitempty"`
		History       [][]float64                       `json:"history,omitempty"`
		SpanPrecision *float64                          `json:"span_precision,omitempty"`
		Uptime        *float64                          `json:"uptime,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors", "group", "history", "span_precision", "uptime"})
	} else {
		return err
	}
	o.Errors = all.Errors
	o.Group = all.Group
	o.History = all.History
	o.SpanPrecision = all.SpanPrecision
	o.Uptime = all.Uptime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
