// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDataAttributes Defines the pipeline’s name and its components (sources, processors, and destinations).
type ObservabilityPipelineDataAttributes struct {
	// Specifies the pipeline's configuration, including its sources, processors, and destinations.
	Config ObservabilityPipelineConfig `json:"config"`
	// Name of the pipeline.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDataAttributes instantiates a new ObservabilityPipelineDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDataAttributes(config ObservabilityPipelineConfig, name string) *ObservabilityPipelineDataAttributes {
	this := ObservabilityPipelineDataAttributes{}
	this.Config = config
	this.Name = name
	return &this
}

// NewObservabilityPipelineDataAttributesWithDefaults instantiates a new ObservabilityPipelineDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDataAttributesWithDefaults() *ObservabilityPipelineDataAttributes {
	this := ObservabilityPipelineDataAttributes{}
	return &this
}

// GetConfig returns the Config field value.
func (o *ObservabilityPipelineDataAttributes) GetConfig() ObservabilityPipelineConfig {
	if o == nil {
		var ret ObservabilityPipelineConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDataAttributes) GetConfigOk() (*ObservabilityPipelineConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *ObservabilityPipelineDataAttributes) SetConfig(v ObservabilityPipelineConfig) {
	o.Config = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineDataAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config *ObservabilityPipelineConfig `json:"config"`
		Name   *string                      `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = *all.Config
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
