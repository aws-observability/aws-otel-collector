// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultTrace Trace identifiers associated with a Synthetic test result.
type SyntheticsTestResultTrace struct {
	// Datadog APM trace identifier.
	Id *string `json:"id,omitempty"`
	// OpenTelemetry trace identifier.
	OtelId *string `json:"otel_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultTrace instantiates a new SyntheticsTestResultTrace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultTrace() *SyntheticsTestResultTrace {
	this := SyntheticsTestResultTrace{}
	return &this
}

// NewSyntheticsTestResultTraceWithDefaults instantiates a new SyntheticsTestResultTrace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultTraceWithDefaults() *SyntheticsTestResultTrace {
	this := SyntheticsTestResultTrace{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultTrace) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTrace) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultTrace) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultTrace) SetId(v string) {
	o.Id = &v
}

// GetOtelId returns the OtelId field value if set, zero value otherwise.
func (o *SyntheticsTestResultTrace) GetOtelId() string {
	if o == nil || o.OtelId == nil {
		var ret string
		return ret
	}
	return *o.OtelId
}

// GetOtelIdOk returns a tuple with the OtelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTrace) GetOtelIdOk() (*string, bool) {
	if o == nil || o.OtelId == nil {
		return nil, false
	}
	return o.OtelId, true
}

// HasOtelId returns a boolean if a field has been set.
func (o *SyntheticsTestResultTrace) HasOtelId() bool {
	return o != nil && o.OtelId != nil
}

// SetOtelId gets a reference to the given string and assigns it to the OtelId field.
func (o *SyntheticsTestResultTrace) SetOtelId(v string) {
	o.OtelId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultTrace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OtelId != nil {
		toSerialize["otel_id"] = o.OtelId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultTrace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string `json:"id,omitempty"`
		OtelId *string `json:"otel_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "otel_id"})
	} else {
		return err
	}
	o.Id = all.Id
	o.OtelId = all.OtelId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
