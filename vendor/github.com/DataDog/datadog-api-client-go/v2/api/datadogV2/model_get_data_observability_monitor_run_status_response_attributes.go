// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetDataObservabilityMonitorRunStatusResponseAttributes The attributes of a data observability monitor run status response.
type GetDataObservabilityMonitorRunStatusResponseAttributes struct {
	// Error message describing why the monitor run failed. Only present when status is error.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The status of a data observability monitor run.
	Status DataObservabilityMonitorRunStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetDataObservabilityMonitorRunStatusResponseAttributes instantiates a new GetDataObservabilityMonitorRunStatusResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetDataObservabilityMonitorRunStatusResponseAttributes(status DataObservabilityMonitorRunStatus) *GetDataObservabilityMonitorRunStatusResponseAttributes {
	this := GetDataObservabilityMonitorRunStatusResponseAttributes{}
	this.Status = status
	return &this
}

// NewGetDataObservabilityMonitorRunStatusResponseAttributesWithDefaults instantiates a new GetDataObservabilityMonitorRunStatusResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetDataObservabilityMonitorRunStatusResponseAttributesWithDefaults() *GetDataObservabilityMonitorRunStatusResponseAttributes {
	this := GetDataObservabilityMonitorRunStatusResponseAttributes{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetStatus returns the Status field value.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) GetStatus() DataObservabilityMonitorRunStatus {
	if o == nil {
		var ret DataObservabilityMonitorRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) GetStatusOk() (*DataObservabilityMonitorRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) SetStatus(v DataObservabilityMonitorRunStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetDataObservabilityMonitorRunStatusResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetDataObservabilityMonitorRunStatusResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorMessage *string                            `json:"error_message,omitempty"`
		Status       *DataObservabilityMonitorRunStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_message", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ErrorMessage = all.ErrorMessage
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
