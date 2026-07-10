// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertAndPublishFormVersionDataAttributes The attributes for upserting and publishing a form version.
type UpsertAndPublishFormVersionDataAttributes struct {
	// A JSON Schema definition that describes the form's data fields.
	DataDefinition FormDataDefinition `json:"data_definition"`
	// UI configuration for rendering form fields, including widget overrides, field ordering, and themes.
	UiDefinition FormUiDefinition `json:"ui_definition"`
	// Concurrency control parameters for the upsert and publish operation.
	UpsertParams UpsertAndPublishFormVersionUpsertParams `json:"upsert_params"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertAndPublishFormVersionDataAttributes instantiates a new UpsertAndPublishFormVersionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertAndPublishFormVersionDataAttributes(dataDefinition FormDataDefinition, uiDefinition FormUiDefinition, upsertParams UpsertAndPublishFormVersionUpsertParams) *UpsertAndPublishFormVersionDataAttributes {
	this := UpsertAndPublishFormVersionDataAttributes{}
	this.DataDefinition = dataDefinition
	this.UiDefinition = uiDefinition
	this.UpsertParams = upsertParams
	return &this
}

// NewUpsertAndPublishFormVersionDataAttributesWithDefaults instantiates a new UpsertAndPublishFormVersionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertAndPublishFormVersionDataAttributesWithDefaults() *UpsertAndPublishFormVersionDataAttributes {
	this := UpsertAndPublishFormVersionDataAttributes{}
	return &this
}

// GetDataDefinition returns the DataDefinition field value.
func (o *UpsertAndPublishFormVersionDataAttributes) GetDataDefinition() FormDataDefinition {
	if o == nil {
		var ret FormDataDefinition
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *UpsertAndPublishFormVersionDataAttributes) GetDataDefinitionOk() (*FormDataDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *UpsertAndPublishFormVersionDataAttributes) SetDataDefinition(v FormDataDefinition) {
	o.DataDefinition = v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *UpsertAndPublishFormVersionDataAttributes) GetUiDefinition() FormUiDefinition {
	if o == nil {
		var ret FormUiDefinition
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *UpsertAndPublishFormVersionDataAttributes) GetUiDefinitionOk() (*FormUiDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *UpsertAndPublishFormVersionDataAttributes) SetUiDefinition(v FormUiDefinition) {
	o.UiDefinition = v
}

// GetUpsertParams returns the UpsertParams field value.
func (o *UpsertAndPublishFormVersionDataAttributes) GetUpsertParams() UpsertAndPublishFormVersionUpsertParams {
	if o == nil {
		var ret UpsertAndPublishFormVersionUpsertParams
		return ret
	}
	return o.UpsertParams
}

// GetUpsertParamsOk returns a tuple with the UpsertParams field value
// and a boolean to check if the value has been set.
func (o *UpsertAndPublishFormVersionDataAttributes) GetUpsertParamsOk() (*UpsertAndPublishFormVersionUpsertParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpsertParams, true
}

// SetUpsertParams sets field value.
func (o *UpsertAndPublishFormVersionDataAttributes) SetUpsertParams(v UpsertAndPublishFormVersionUpsertParams) {
	o.UpsertParams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertAndPublishFormVersionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_definition"] = o.DataDefinition
	toSerialize["ui_definition"] = o.UiDefinition
	toSerialize["upsert_params"] = o.UpsertParams

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertAndPublishFormVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataDefinition *FormDataDefinition                      `json:"data_definition"`
		UiDefinition   *FormUiDefinition                        `json:"ui_definition"`
		UpsertParams   *UpsertAndPublishFormVersionUpsertParams `json:"upsert_params"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	if all.UpsertParams == nil {
		return fmt.Errorf("required field upsert_params missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_definition", "ui_definition", "upsert_params"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DataDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataDefinition = *all.DataDefinition
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
