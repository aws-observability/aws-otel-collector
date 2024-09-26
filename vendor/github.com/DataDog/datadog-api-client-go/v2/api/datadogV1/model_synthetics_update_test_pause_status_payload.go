// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsUpdateTestPauseStatusPayload Object to start or pause an existing Synthetic test.
type SyntheticsUpdateTestPauseStatusPayload struct {
	// Define whether you want to start (`live`) or pause (`paused`) a
	// Synthetic test.
	NewStatus *SyntheticsTestPauseStatus `json:"new_status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsUpdateTestPauseStatusPayload instantiates a new SyntheticsUpdateTestPauseStatusPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsUpdateTestPauseStatusPayload() *SyntheticsUpdateTestPauseStatusPayload {
	this := SyntheticsUpdateTestPauseStatusPayload{}
	return &this
}

// NewSyntheticsUpdateTestPauseStatusPayloadWithDefaults instantiates a new SyntheticsUpdateTestPauseStatusPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsUpdateTestPauseStatusPayloadWithDefaults() *SyntheticsUpdateTestPauseStatusPayload {
	this := SyntheticsUpdateTestPauseStatusPayload{}
	return &this
}

// GetNewStatus returns the NewStatus field value if set, zero value otherwise.
func (o *SyntheticsUpdateTestPauseStatusPayload) GetNewStatus() SyntheticsTestPauseStatus {
	if o == nil || o.NewStatus == nil {
		var ret SyntheticsTestPauseStatus
		return ret
	}
	return *o.NewStatus
}

// GetNewStatusOk returns a tuple with the NewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsUpdateTestPauseStatusPayload) GetNewStatusOk() (*SyntheticsTestPauseStatus, bool) {
	if o == nil || o.NewStatus == nil {
		return nil, false
	}
	return o.NewStatus, true
}

// HasNewStatus returns a boolean if a field has been set.
func (o *SyntheticsUpdateTestPauseStatusPayload) HasNewStatus() bool {
	return o != nil && o.NewStatus != nil
}

// SetNewStatus gets a reference to the given SyntheticsTestPauseStatus and assigns it to the NewStatus field.
func (o *SyntheticsUpdateTestPauseStatusPayload) SetNewStatus(v SyntheticsTestPauseStatus) {
	o.NewStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsUpdateTestPauseStatusPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NewStatus != nil {
		toSerialize["new_status"] = o.NewStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsUpdateTestPauseStatusPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NewStatus *SyntheticsTestPauseStatus `json:"new_status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"new_status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.NewStatus != nil && !all.NewStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.NewStatus = all.NewStatus
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
