// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems The definition of `UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems` object.
type UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems struct {
	// The `items` `input_column`.
	InputColumn string `json:"input_column"`
	// The `items` `output_key`.
	OutputKey string `json:"output_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems instantiates a new UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems(inputColumn string, outputKey string) *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems {
	this := UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems{}
	this.InputColumn = inputColumn
	this.OutputKey = outputKey
	return &this
}

// NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItemsWithDefaults instantiates a new UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItemsWithDefaults() *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems {
	this := UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems{}
	return &this
}

// GetInputColumn returns the InputColumn field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) GetInputColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InputColumn
}

// GetInputColumnOk returns a tuple with the InputColumn field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) GetInputColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputColumn, true
}

// SetInputColumn sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) SetInputColumn(v string) {
	o.InputColumn = v
}

// GetOutputKey returns the OutputKey field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) GetOutputKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OutputKey
}

// GetOutputKeyOk returns a tuple with the OutputKey field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) GetOutputKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputKey, true
}

// SetOutputKey sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) SetOutputKey(v string) {
	o.OutputKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["input_column"] = o.InputColumn
	toSerialize["output_key"] = o.OutputKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InputColumn *string `json:"input_column"`
		OutputKey   *string `json:"output_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InputColumn == nil {
		return fmt.Errorf("required field input_column missing")
	}
	if all.OutputKey == nil {
		return fmt.Errorf("required field output_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"input_column", "output_key"})
	} else {
		return err
	}
	o.InputColumn = *all.InputColumn
	o.OutputKey = *all.OutputKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
