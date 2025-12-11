// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCustomProcessorRemap Defines a single VRL remap rule with its own filtering and transformation logic.
type ObservabilityPipelineCustomProcessorRemap struct {
	// Whether to drop events that caused errors during processing.
	DropOnError bool `json:"drop_on_error"`
	// Whether this remap rule is enabled.
	Enabled bool `json:"enabled"`
	// A Datadog search query used to filter events for this specific remap rule.
	Include string `json:"include"`
	// A descriptive name for this remap rule.
	Name string `json:"name"`
	// The VRL script source code that defines the processing logic.
	Source string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCustomProcessorRemap instantiates a new ObservabilityPipelineCustomProcessorRemap object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCustomProcessorRemap(dropOnError bool, enabled bool, include string, name string, source string) *ObservabilityPipelineCustomProcessorRemap {
	this := ObservabilityPipelineCustomProcessorRemap{}
	this.DropOnError = dropOnError
	this.Enabled = enabled
	this.Include = include
	this.Name = name
	this.Source = source
	return &this
}

// NewObservabilityPipelineCustomProcessorRemapWithDefaults instantiates a new ObservabilityPipelineCustomProcessorRemap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCustomProcessorRemapWithDefaults() *ObservabilityPipelineCustomProcessorRemap {
	this := ObservabilityPipelineCustomProcessorRemap{}
	return &this
}

// GetDropOnError returns the DropOnError field value.
func (o *ObservabilityPipelineCustomProcessorRemap) GetDropOnError() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DropOnError
}

// GetDropOnErrorOk returns a tuple with the DropOnError field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorRemap) GetDropOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropOnError, true
}

// SetDropOnError sets field value.
func (o *ObservabilityPipelineCustomProcessorRemap) SetDropOnError(v bool) {
	o.DropOnError = v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineCustomProcessorRemap) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorRemap) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineCustomProcessorRemap) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineCustomProcessorRemap) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorRemap) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineCustomProcessorRemap) SetInclude(v string) {
	o.Include = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineCustomProcessorRemap) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorRemap) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineCustomProcessorRemap) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value.
func (o *ObservabilityPipelineCustomProcessorRemap) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCustomProcessorRemap) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ObservabilityPipelineCustomProcessorRemap) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCustomProcessorRemap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["drop_on_error"] = o.DropOnError
	toSerialize["enabled"] = o.Enabled
	toSerialize["include"] = o.Include
	toSerialize["name"] = o.Name
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineCustomProcessorRemap) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DropOnError *bool   `json:"drop_on_error"`
		Enabled     *bool   `json:"enabled"`
		Include     *string `json:"include"`
		Name        *string `json:"name"`
		Source      *string `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DropOnError == nil {
		return fmt.Errorf("required field drop_on_error missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"drop_on_error", "enabled", "include", "name", "source"})
	} else {
		return err
	}
	o.DropOnError = *all.DropOnError
	o.Enabled = *all.Enabled
	o.Include = *all.Include
	o.Name = *all.Name
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
