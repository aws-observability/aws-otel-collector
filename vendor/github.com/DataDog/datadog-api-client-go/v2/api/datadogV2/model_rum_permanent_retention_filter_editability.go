// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterEditability Indicates which cross-product fields of a permanent RUM retention filter can be updated.
type RumPermanentRetentionFilterEditability struct {
	// Whether the APM trace cross-product configuration of the filter can be updated.
	TraceEditable *bool `json:"trace_editable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumPermanentRetentionFilterEditability instantiates a new RumPermanentRetentionFilterEditability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumPermanentRetentionFilterEditability() *RumPermanentRetentionFilterEditability {
	this := RumPermanentRetentionFilterEditability{}
	return &this
}

// NewRumPermanentRetentionFilterEditabilityWithDefaults instantiates a new RumPermanentRetentionFilterEditability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumPermanentRetentionFilterEditabilityWithDefaults() *RumPermanentRetentionFilterEditability {
	this := RumPermanentRetentionFilterEditability{}
	return &this
}

// GetTraceEditable returns the TraceEditable field value if set, zero value otherwise.
func (o *RumPermanentRetentionFilterEditability) GetTraceEditable() bool {
	if o == nil || o.TraceEditable == nil {
		var ret bool
		return ret
	}
	return *o.TraceEditable
}

// GetTraceEditableOk returns a tuple with the TraceEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentRetentionFilterEditability) GetTraceEditableOk() (*bool, bool) {
	if o == nil || o.TraceEditable == nil {
		return nil, false
	}
	return o.TraceEditable, true
}

// HasTraceEditable returns a boolean if a field has been set.
func (o *RumPermanentRetentionFilterEditability) HasTraceEditable() bool {
	return o != nil && o.TraceEditable != nil
}

// SetTraceEditable gets a reference to the given bool and assigns it to the TraceEditable field.
func (o *RumPermanentRetentionFilterEditability) SetTraceEditable(v bool) {
	o.TraceEditable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumPermanentRetentionFilterEditability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TraceEditable != nil {
		toSerialize["trace_editable"] = o.TraceEditable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumPermanentRetentionFilterEditability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TraceEditable *bool `json:"trace_editable,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"trace_editable"})
	} else {
		return err
	}
	o.TraceEditable = all.TraceEditable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
