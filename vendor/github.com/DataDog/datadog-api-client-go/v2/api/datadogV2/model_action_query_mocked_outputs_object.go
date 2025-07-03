// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQueryMockedOutputsObject The mocked outputs of the action query.
type ActionQueryMockedOutputsObject struct {
	// Whether to enable the mocked outputs for testing.
	Enabled ActionQueryMockedOutputsEnabled `json:"enabled"`
	// The mocked outputs of the action query, serialized as JSON.
	Outputs *string `json:"outputs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionQueryMockedOutputsObject instantiates a new ActionQueryMockedOutputsObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionQueryMockedOutputsObject(enabled ActionQueryMockedOutputsEnabled) *ActionQueryMockedOutputsObject {
	this := ActionQueryMockedOutputsObject{}
	this.Enabled = enabled
	return &this
}

// NewActionQueryMockedOutputsObjectWithDefaults instantiates a new ActionQueryMockedOutputsObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionQueryMockedOutputsObjectWithDefaults() *ActionQueryMockedOutputsObject {
	this := ActionQueryMockedOutputsObject{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *ActionQueryMockedOutputsObject) GetEnabled() ActionQueryMockedOutputsEnabled {
	if o == nil {
		var ret ActionQueryMockedOutputsEnabled
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ActionQueryMockedOutputsObject) GetEnabledOk() (*ActionQueryMockedOutputsEnabled, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ActionQueryMockedOutputsObject) SetEnabled(v ActionQueryMockedOutputsEnabled) {
	o.Enabled = v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *ActionQueryMockedOutputsObject) GetOutputs() string {
	if o == nil || o.Outputs == nil {
		var ret string
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryMockedOutputsObject) GetOutputsOk() (*string, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *ActionQueryMockedOutputsObject) HasOutputs() bool {
	return o != nil && o.Outputs != nil
}

// SetOutputs gets a reference to the given string and assigns it to the Outputs field.
func (o *ActionQueryMockedOutputsObject) SetOutputs(v string) {
	o.Outputs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionQueryMockedOutputsObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionQueryMockedOutputsObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *ActionQueryMockedOutputsEnabled `json:"enabled"`
		Outputs *string                          `json:"outputs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "outputs"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Outputs = all.Outputs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
