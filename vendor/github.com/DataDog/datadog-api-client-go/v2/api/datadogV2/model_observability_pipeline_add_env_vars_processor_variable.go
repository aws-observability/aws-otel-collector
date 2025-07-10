// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddEnvVarsProcessorVariable Defines a mapping between an environment variable and a log field.
type ObservabilityPipelineAddEnvVarsProcessorVariable struct {
	// The target field in the log event.
	Field string `json:"field"`
	// The name of the environment variable to read.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAddEnvVarsProcessorVariable instantiates a new ObservabilityPipelineAddEnvVarsProcessorVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAddEnvVarsProcessorVariable(field string, name string) *ObservabilityPipelineAddEnvVarsProcessorVariable {
	this := ObservabilityPipelineAddEnvVarsProcessorVariable{}
	this.Field = field
	this.Name = name
	return &this
}

// NewObservabilityPipelineAddEnvVarsProcessorVariableWithDefaults instantiates a new ObservabilityPipelineAddEnvVarsProcessorVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAddEnvVarsProcessorVariableWithDefaults() *ObservabilityPipelineAddEnvVarsProcessorVariable {
	this := ObservabilityPipelineAddEnvVarsProcessorVariable{}
	return &this
}

// GetField returns the Field field value.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) SetField(v string) {
	o.Field = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAddEnvVarsProcessorVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAddEnvVarsProcessorVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field *string `json:"field"`
		Name  *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "name"})
	} else {
		return err
	}
	o.Field = *all.Field
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
