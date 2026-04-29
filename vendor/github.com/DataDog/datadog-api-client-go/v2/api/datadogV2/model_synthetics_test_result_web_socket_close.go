// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultWebSocketClose WebSocket close frame information for WebSocket test responses.
type SyntheticsTestResultWebSocketClose struct {
	// Reason string received in the close frame.
	Reason *string `json:"reason,omitempty"`
	// Status code received in the close frame.
	StatusCode *int64 `json:"status_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultWebSocketClose instantiates a new SyntheticsTestResultWebSocketClose object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultWebSocketClose() *SyntheticsTestResultWebSocketClose {
	this := SyntheticsTestResultWebSocketClose{}
	return &this
}

// NewSyntheticsTestResultWebSocketCloseWithDefaults instantiates a new SyntheticsTestResultWebSocketClose object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultWebSocketCloseWithDefaults() *SyntheticsTestResultWebSocketClose {
	this := SyntheticsTestResultWebSocketClose{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SyntheticsTestResultWebSocketClose) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultWebSocketClose) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SyntheticsTestResultWebSocketClose) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SyntheticsTestResultWebSocketClose) SetReason(v string) {
	o.Reason = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SyntheticsTestResultWebSocketClose) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultWebSocketClose) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SyntheticsTestResultWebSocketClose) HasStatusCode() bool {
	return o != nil && o.StatusCode != nil
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *SyntheticsTestResultWebSocketClose) SetStatusCode(v int64) {
	o.StatusCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultWebSocketClose) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultWebSocketClose) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason     *string `json:"reason,omitempty"`
		StatusCode *int64  `json:"status_code,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reason", "status_code"})
	} else {
		return err
	}
	o.Reason = all.Reason
	o.StatusCode = all.StatusCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
