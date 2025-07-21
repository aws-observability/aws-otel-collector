// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRenameFieldsProcessor The `rename_fields` processor changes field names.
type ObservabilityPipelineRenameFieldsProcessor struct {
	// A list of rename rules specifying which fields to rename in the event, what to rename them to, and whether to preserve the original fields.
	Fields []ObservabilityPipelineRenameFieldsProcessorField `json:"fields"`
	// A unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The processor type. The value should always be `rename_fields`.
	Type ObservabilityPipelineRenameFieldsProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineRenameFieldsProcessor instantiates a new ObservabilityPipelineRenameFieldsProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineRenameFieldsProcessor(fields []ObservabilityPipelineRenameFieldsProcessorField, id string, include string, inputs []string, typeVar ObservabilityPipelineRenameFieldsProcessorType) *ObservabilityPipelineRenameFieldsProcessor {
	this := ObservabilityPipelineRenameFieldsProcessor{}
	this.Fields = fields
	this.Id = id
	this.Include = include
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineRenameFieldsProcessorWithDefaults instantiates a new ObservabilityPipelineRenameFieldsProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineRenameFieldsProcessorWithDefaults() *ObservabilityPipelineRenameFieldsProcessor {
	this := ObservabilityPipelineRenameFieldsProcessor{}
	var typeVar ObservabilityPipelineRenameFieldsProcessorType = OBSERVABILITYPIPELINERENAMEFIELDSPROCESSORTYPE_RENAME_FIELDS
	this.Type = typeVar
	return &this
}

// GetFields returns the Fields field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetFields() []ObservabilityPipelineRenameFieldsProcessorField {
	if o == nil {
		var ret []ObservabilityPipelineRenameFieldsProcessorField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetFieldsOk() (*[]ObservabilityPipelineRenameFieldsProcessorField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) SetFields(v []ObservabilityPipelineRenameFieldsProcessorField) {
	o.Fields = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) SetInclude(v string) {
	o.Include = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetType() ObservabilityPipelineRenameFieldsProcessorType {
	if o == nil {
		var ret ObservabilityPipelineRenameFieldsProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessor) GetTypeOk() (*ObservabilityPipelineRenameFieldsProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessor) SetType(v ObservabilityPipelineRenameFieldsProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineRenameFieldsProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineRenameFieldsProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields  *[]ObservabilityPipelineRenameFieldsProcessorField `json:"fields"`
		Id      *string                                            `json:"id"`
		Include *string                                            `json:"include"`
		Inputs  *[]string                                          `json:"inputs"`
		Type    *ObservabilityPipelineRenameFieldsProcessorType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "id", "include", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Fields = *all.Fields
	o.Id = *all.Id
	o.Include = *all.Include
	o.Inputs = *all.Inputs
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
