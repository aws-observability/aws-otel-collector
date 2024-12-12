// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitGraphWidgetDefinition The split graph widget allows you to create repeating units of a graph - one for each value in a group (for example: one per service)
type SplitGraphWidgetDefinition struct {
	// Normalize y axes across graphs
	HasUniformYAxes *bool `json:"has_uniform_y_axes,omitempty"`
	// Size of the individual graphs in the split.
	Size SplitGraphVizSize `json:"size"`
	// The original widget we are splitting on.
	SourceWidgetDefinition SplitGraphSourceWidgetDefinition `json:"source_widget_definition"`
	// Encapsulates all user choices about how to split a graph.
	SplitConfig SplitConfig `json:"split_config"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title *string `json:"title,omitempty"`
	// Type of the split graph widget
	Type SplitGraphWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSplitGraphWidgetDefinition instantiates a new SplitGraphWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitGraphWidgetDefinition(size SplitGraphVizSize, sourceWidgetDefinition SplitGraphSourceWidgetDefinition, splitConfig SplitConfig, typeVar SplitGraphWidgetDefinitionType) *SplitGraphWidgetDefinition {
	this := SplitGraphWidgetDefinition{}
	this.Size = size
	this.SourceWidgetDefinition = sourceWidgetDefinition
	this.SplitConfig = splitConfig
	this.Type = typeVar
	return &this
}

// NewSplitGraphWidgetDefinitionWithDefaults instantiates a new SplitGraphWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitGraphWidgetDefinitionWithDefaults() *SplitGraphWidgetDefinition {
	this := SplitGraphWidgetDefinition{}
	var typeVar SplitGraphWidgetDefinitionType = SPLITGRAPHWIDGETDEFINITIONTYPE_SPLIT_GROUP
	this.Type = typeVar
	return &this
}

// GetHasUniformYAxes returns the HasUniformYAxes field value if set, zero value otherwise.
func (o *SplitGraphWidgetDefinition) GetHasUniformYAxes() bool {
	if o == nil || o.HasUniformYAxes == nil {
		var ret bool
		return ret
	}
	return *o.HasUniformYAxes
}

// GetHasUniformYAxesOk returns a tuple with the HasUniformYAxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetHasUniformYAxesOk() (*bool, bool) {
	if o == nil || o.HasUniformYAxes == nil {
		return nil, false
	}
	return o.HasUniformYAxes, true
}

// HasHasUniformYAxes returns a boolean if a field has been set.
func (o *SplitGraphWidgetDefinition) HasHasUniformYAxes() bool {
	return o != nil && o.HasUniformYAxes != nil
}

// SetHasUniformYAxes gets a reference to the given bool and assigns it to the HasUniformYAxes field.
func (o *SplitGraphWidgetDefinition) SetHasUniformYAxes(v bool) {
	o.HasUniformYAxes = &v
}

// GetSize returns the Size field value.
func (o *SplitGraphWidgetDefinition) GetSize() SplitGraphVizSize {
	if o == nil {
		var ret SplitGraphVizSize
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetSizeOk() (*SplitGraphVizSize, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *SplitGraphWidgetDefinition) SetSize(v SplitGraphVizSize) {
	o.Size = v
}

// GetSourceWidgetDefinition returns the SourceWidgetDefinition field value.
func (o *SplitGraphWidgetDefinition) GetSourceWidgetDefinition() SplitGraphSourceWidgetDefinition {
	if o == nil {
		var ret SplitGraphSourceWidgetDefinition
		return ret
	}
	return o.SourceWidgetDefinition
}

// GetSourceWidgetDefinitionOk returns a tuple with the SourceWidgetDefinition field value
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetSourceWidgetDefinitionOk() (*SplitGraphSourceWidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceWidgetDefinition, true
}

// SetSourceWidgetDefinition sets field value.
func (o *SplitGraphWidgetDefinition) SetSourceWidgetDefinition(v SplitGraphSourceWidgetDefinition) {
	o.SourceWidgetDefinition = v
}

// GetSplitConfig returns the SplitConfig field value.
func (o *SplitGraphWidgetDefinition) GetSplitConfig() SplitConfig {
	if o == nil {
		var ret SplitConfig
		return ret
	}
	return o.SplitConfig
}

// GetSplitConfigOk returns a tuple with the SplitConfig field value
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetSplitConfigOk() (*SplitConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplitConfig, true
}

// SetSplitConfig sets field value.
func (o *SplitGraphWidgetDefinition) SetSplitConfig(v SplitConfig) {
	o.SplitConfig = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SplitGraphWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SplitGraphWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *SplitGraphWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SplitGraphWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SplitGraphWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SplitGraphWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *SplitGraphWidgetDefinition) GetType() SplitGraphWidgetDefinitionType {
	if o == nil {
		var ret SplitGraphWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SplitGraphWidgetDefinition) GetTypeOk() (*SplitGraphWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SplitGraphWidgetDefinition) SetType(v SplitGraphWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitGraphWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HasUniformYAxes != nil {
		toSerialize["has_uniform_y_axes"] = o.HasUniformYAxes
	}
	toSerialize["size"] = o.Size
	toSerialize["source_widget_definition"] = o.SourceWidgetDefinition
	toSerialize["split_config"] = o.SplitConfig
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitGraphWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasUniformYAxes        *bool                             `json:"has_uniform_y_axes,omitempty"`
		Size                   *SplitGraphVizSize                `json:"size"`
		SourceWidgetDefinition *SplitGraphSourceWidgetDefinition `json:"source_widget_definition"`
		SplitConfig            *SplitConfig                      `json:"split_config"`
		Time                   *WidgetTime                       `json:"time,omitempty"`
		Title                  *string                           `json:"title,omitempty"`
		Type                   *SplitGraphWidgetDefinitionType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	if all.SourceWidgetDefinition == nil {
		return fmt.Errorf("required field source_widget_definition missing")
	}
	if all.SplitConfig == nil {
		return fmt.Errorf("required field split_config missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_uniform_y_axes", "size", "source_widget_definition", "split_config", "time", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HasUniformYAxes = all.HasUniformYAxes
	if !all.Size.IsValid() {
		hasInvalidField = true
	} else {
		o.Size = *all.Size
	}
	o.SourceWidgetDefinition = *all.SourceWidgetDefinition
	if all.SplitConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SplitConfig = *all.SplitConfig
	o.Time = all.Time
	o.Title = all.Title
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
