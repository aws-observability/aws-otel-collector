// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertFormVersionDataAttributes The attributes for creating or updating a form version.
type UpsertFormVersionDataAttributes struct {
	// A JSON Schema definition that describes the form's data fields.
	DataDefinition FormDataDefinition `json:"data_definition"`
	// The state of a form version.
	State FormVersionState `json:"state"`
	// UI configuration for rendering form fields, including widget overrides, field ordering, and themes.
	UiDefinition FormUiDefinition `json:"ui_definition"`
	// Concurrency control parameters for the form version upsert operation.
	UpsertParams UpsertFormVersionUpsertParams `json:"upsert_params"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertFormVersionDataAttributes instantiates a new UpsertFormVersionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertFormVersionDataAttributes(dataDefinition FormDataDefinition, state FormVersionState, uiDefinition FormUiDefinition, upsertParams UpsertFormVersionUpsertParams) *UpsertFormVersionDataAttributes {
	this := UpsertFormVersionDataAttributes{}
	this.DataDefinition = dataDefinition
	this.State = state
	this.UiDefinition = uiDefinition
	this.UpsertParams = upsertParams
	return &this
}

// NewUpsertFormVersionDataAttributesWithDefaults instantiates a new UpsertFormVersionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertFormVersionDataAttributesWithDefaults() *UpsertFormVersionDataAttributes {
	this := UpsertFormVersionDataAttributes{}
	return &this
}

// GetDataDefinition returns the DataDefinition field value.
func (o *UpsertFormVersionDataAttributes) GetDataDefinition() FormDataDefinition {
	if o == nil {
		var ret FormDataDefinition
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionDataAttributes) GetDataDefinitionOk() (*FormDataDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *UpsertFormVersionDataAttributes) SetDataDefinition(v FormDataDefinition) {
	o.DataDefinition = v
}

// GetState returns the State field value.
func (o *UpsertFormVersionDataAttributes) GetState() FormVersionState {
	if o == nil {
		var ret FormVersionState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionDataAttributes) GetStateOk() (*FormVersionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *UpsertFormVersionDataAttributes) SetState(v FormVersionState) {
	o.State = v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *UpsertFormVersionDataAttributes) GetUiDefinition() FormUiDefinition {
	if o == nil {
		var ret FormUiDefinition
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionDataAttributes) GetUiDefinitionOk() (*FormUiDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *UpsertFormVersionDataAttributes) SetUiDefinition(v FormUiDefinition) {
	o.UiDefinition = v
}

// GetUpsertParams returns the UpsertParams field value.
func (o *UpsertFormVersionDataAttributes) GetUpsertParams() UpsertFormVersionUpsertParams {
	if o == nil {
		var ret UpsertFormVersionUpsertParams
		return ret
	}
	return o.UpsertParams
}

// GetUpsertParamsOk returns a tuple with the UpsertParams field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionDataAttributes) GetUpsertParamsOk() (*UpsertFormVersionUpsertParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpsertParams, true
}

// SetUpsertParams sets field value.
func (o *UpsertFormVersionDataAttributes) SetUpsertParams(v UpsertFormVersionUpsertParams) {
	o.UpsertParams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertFormVersionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_definition"] = o.DataDefinition
	toSerialize["state"] = o.State
	toSerialize["ui_definition"] = o.UiDefinition
	toSerialize["upsert_params"] = o.UpsertParams

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertFormVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataDefinition *FormDataDefinition            `json:"data_definition"`
		State          *FormVersionState              `json:"state"`
		UiDefinition   *FormUiDefinition              `json:"ui_definition"`
		UpsertParams   *UpsertFormVersionUpsertParams `json:"upsert_params"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	if all.UpsertParams == nil {
		return fmt.Errorf("required field upsert_params missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_definition", "state", "ui_definition", "upsert_params"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DataDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataDefinition = *all.DataDefinition
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if all.UiDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UiDefinition = *all.UiDefinition
	if all.UpsertParams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpsertParams = *all.UpsertParams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
