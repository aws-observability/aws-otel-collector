// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultHealthCheck Health check information returned from a gRPC health check call.
type SyntheticsTestResultHealthCheck struct {
	// Raw health check message payload.
	Message map[string]string `json:"message,omitempty"`
	// Health check status code.
	Status *int64 `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultHealthCheck instantiates a new SyntheticsTestResultHealthCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultHealthCheck() *SyntheticsTestResultHealthCheck {
	this := SyntheticsTestResultHealthCheck{}
	return &this
}

// NewSyntheticsTestResultHealthCheckWithDefaults instantiates a new SyntheticsTestResultHealthCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultHealthCheckWithDefaults() *SyntheticsTestResultHealthCheck {
	this := SyntheticsTestResultHealthCheck{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestResultHealthCheck) GetMessage() map[string]string {
	if o == nil || o.Message == nil {
		var ret map[string]string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultHealthCheck) GetMessageOk() (*map[string]string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultHealthCheck) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given map[string]string and assigns it to the Message field.
func (o *SyntheticsTestResultHealthCheck) SetMessage(v map[string]string) {
	o.Message = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultHealthCheck) GetStatus() int64 {
	if o == nil || o.Status == nil {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultHealthCheck) GetStatusOk() (*int64, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultHealthCheck) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SyntheticsTestResultHealthCheck) SetStatus(v int64) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultHealthCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message map[string]string `json:"message,omitempty"`
		Status  *int64            `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "status"})
	} else {
		return err
	}
	o.Message = all.Message
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
