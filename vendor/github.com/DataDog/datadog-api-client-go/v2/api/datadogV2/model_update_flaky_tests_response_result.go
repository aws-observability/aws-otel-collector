// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsResponseResult Result of updating a single flaky test state.
type UpdateFlakyTestsResponseResult struct {
	// Error message if the update failed.
	Error *string `json:"error,omitempty"`
	// The ID of the flaky test from the request. This is the same ID returned by the Search flaky tests endpoint and corresponds to the test_fingerprint_fqn field in test run events.
	Id string `json:"id"`
	// `True` if the update was successful, `False` if there were any errors.
	Success bool `json:"success"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFlakyTestsResponseResult instantiates a new UpdateFlakyTestsResponseResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFlakyTestsResponseResult(id string, success bool) *UpdateFlakyTestsResponseResult {
	this := UpdateFlakyTestsResponseResult{}
	this.Id = id
	this.Success = success
	return &this
}

// NewUpdateFlakyTestsResponseResultWithDefaults instantiates a new UpdateFlakyTestsResponseResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFlakyTestsResponseResultWithDefaults() *UpdateFlakyTestsResponseResult {
	this := UpdateFlakyTestsResponseResult{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UpdateFlakyTestsResponseResult) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsResponseResult) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UpdateFlakyTestsResponseResult) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *UpdateFlakyTestsResponseResult) SetError(v string) {
	o.Error = &v
}

// GetId returns the Id field value.
func (o *UpdateFlakyTestsResponseResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsResponseResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UpdateFlakyTestsResponseResult) SetId(v string) {
	o.Id = v
}

// GetSuccess returns the Success field value.
func (o *UpdateFlakyTestsResponseResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsResponseResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value.
func (o *UpdateFlakyTestsResponseResult) SetSuccess(v bool) {
	o.Success = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFlakyTestsResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	toSerialize["id"] = o.Id
	toSerialize["success"] = o.Success

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFlakyTestsResponseResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error   *string `json:"error,omitempty"`
		Id      *string `json:"id"`
		Success *bool   `json:"success"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Success == nil {
		return fmt.Errorf("required field success missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "id", "success"})
	} else {
		return err
	}
	o.Error = all.Error
	o.Id = *all.Id
	o.Success = *all.Success

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
