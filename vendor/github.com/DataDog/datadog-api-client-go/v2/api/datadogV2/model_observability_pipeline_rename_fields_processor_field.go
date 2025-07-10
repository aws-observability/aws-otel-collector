// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRenameFieldsProcessorField Defines how to rename a field in log events.
type ObservabilityPipelineRenameFieldsProcessorField struct {
	// The field name to assign the renamed value to.
	Destination string `json:"destination"`
	// Indicates whether the original field, that is received from the source, should be kept (`true`) or removed (`false`) after renaming.
	PreserveSource bool `json:"preserve_source"`
	// The original field name in the log event that should be renamed.
	Source string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineRenameFieldsProcessorField instantiates a new ObservabilityPipelineRenameFieldsProcessorField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineRenameFieldsProcessorField(destination string, preserveSource bool, source string) *ObservabilityPipelineRenameFieldsProcessorField {
	this := ObservabilityPipelineRenameFieldsProcessorField{}
	this.Destination = destination
	this.PreserveSource = preserveSource
	this.Source = source
	return &this
}

// NewObservabilityPipelineRenameFieldsProcessorFieldWithDefaults instantiates a new ObservabilityPipelineRenameFieldsProcessorField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineRenameFieldsProcessorFieldWithDefaults() *ObservabilityPipelineRenameFieldsProcessorField {
	this := ObservabilityPipelineRenameFieldsProcessorField{}
	return &this
}

// GetDestination returns the Destination field value.
func (o *ObservabilityPipelineRenameFieldsProcessorField) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessorField) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessorField) SetDestination(v string) {
	o.Destination = v
}

// GetPreserveSource returns the PreserveSource field value.
func (o *ObservabilityPipelineRenameFieldsProcessorField) GetPreserveSource() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PreserveSource
}

// GetPreserveSourceOk returns a tuple with the PreserveSource field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessorField) GetPreserveSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreserveSource, true
}

// SetPreserveSource sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessorField) SetPreserveSource(v bool) {
	o.PreserveSource = v
}

// GetSource returns the Source field value.
func (o *ObservabilityPipelineRenameFieldsProcessorField) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineRenameFieldsProcessorField) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ObservabilityPipelineRenameFieldsProcessorField) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineRenameFieldsProcessorField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destination"] = o.Destination
	toSerialize["preserve_source"] = o.PreserveSource
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineRenameFieldsProcessorField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destination    *string `json:"destination"`
		PreserveSource *bool   `json:"preserve_source"`
		Source         *string `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Destination == nil {
		return fmt.Errorf("required field destination missing")
	}
	if all.PreserveSource == nil {
		return fmt.Errorf("required field preserve_source missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination", "preserve_source", "source"})
	} else {
		return err
	}
	o.Destination = *all.Destination
	o.PreserveSource = *all.PreserveSource
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
