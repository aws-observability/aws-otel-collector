// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFilterProcessor The `filter` processor allows conditional processing of logs/metrics based on a Datadog search query. Logs/metrics that match the `include` query are passed through; others are discarded.
//
// **Supported pipeline types:** logs, metrics
type ObservabilityPipelineFilterProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs/metrics should pass through the filter. Logs/metrics that match this query continue to downstream components; others are dropped.
	Include string `json:"include"`
	// The processor type. The value should always be `filter`.
	Type ObservabilityPipelineFilterProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineFilterProcessor instantiates a new ObservabilityPipelineFilterProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineFilterProcessor(enabled bool, id string, include string, typeVar ObservabilityPipelineFilterProcessorType) *ObservabilityPipelineFilterProcessor {
	this := ObservabilityPipelineFilterProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineFilterProcessorWithDefaults instantiates a new ObservabilityPipelineFilterProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineFilterProcessorWithDefaults() *ObservabilityPipelineFilterProcessor {
	this := ObservabilityPipelineFilterProcessor{}
	var typeVar ObservabilityPipelineFilterProcessorType = OBSERVABILITYPIPELINEFILTERPROCESSORTYPE_FILTER
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineFilterProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFilterProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineFilterProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineFilterProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineFilterProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFilterProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineFilterProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineFilterProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFilterProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineFilterProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineFilterProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFilterProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineFilterProcessor) SetInclude(v string) {
	o.Include = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineFilterProcessor) GetType() ObservabilityPipelineFilterProcessorType {
	if o == nil {
		var ret ObservabilityPipelineFilterProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFilterProcessor) GetTypeOk() (*ObservabilityPipelineFilterProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineFilterProcessor) SetType(v ObservabilityPipelineFilterProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineFilterProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineFilterProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string                                   `json:"display_name,omitempty"`
		Enabled     *bool                                     `json:"enabled"`
		Id          *string                                   `json:"id"`
		Include     *string                                   `json:"include"`
		Type        *ObservabilityPipelineFilterProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "id", "include", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
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
