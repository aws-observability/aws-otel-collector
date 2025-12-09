// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaRemapper The schema remapper maps source log fields to their correct fields.
type LogsSchemaRemapper struct {
	// Name of the logs schema remapper.
	Name string `json:"name"`
	// Override or not the target element if already set.
	OverrideOnConflict *bool `json:"override_on_conflict,omitempty"`
	// Remove or preserve the remapped source element.
	PreserveSource *bool `json:"preserve_source,omitempty"`
	// Array of source attributes.
	Sources []string `json:"sources"`
	// Target field to map log source field to.
	Target string `json:"target"`
	// If the `target_type` of the remapper is `attribute`, try to cast the value to a new specific type.
	// If the cast is not possible, the original type is kept. `string`, `integer`, or `double` are the possible types.
	// If the `target_type` is `tag`, this parameter may not be specified.
	TargetFormat *TargetFormatType `json:"target_format,omitempty"`
	// Type of logs schema remapper.
	Type LogsSchemaRemapperType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsSchemaRemapper instantiates a new LogsSchemaRemapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsSchemaRemapper(name string, sources []string, target string, typeVar LogsSchemaRemapperType) *LogsSchemaRemapper {
	this := LogsSchemaRemapper{}
	this.Name = name
	var overrideOnConflict bool = false
	this.OverrideOnConflict = &overrideOnConflict
	var preserveSource bool = false
	this.PreserveSource = &preserveSource
	this.Sources = sources
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsSchemaRemapperWithDefaults instantiates a new LogsSchemaRemapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsSchemaRemapperWithDefaults() *LogsSchemaRemapper {
	this := LogsSchemaRemapper{}
	var overrideOnConflict bool = false
	this.OverrideOnConflict = &overrideOnConflict
	var preserveSource bool = false
	this.PreserveSource = &preserveSource
	return &this
}

// GetName returns the Name field value.
func (o *LogsSchemaRemapper) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LogsSchemaRemapper) SetName(v string) {
	o.Name = v
}

// GetOverrideOnConflict returns the OverrideOnConflict field value if set, zero value otherwise.
func (o *LogsSchemaRemapper) GetOverrideOnConflict() bool {
	if o == nil || o.OverrideOnConflict == nil {
		var ret bool
		return ret
	}
	return *o.OverrideOnConflict
}

// GetOverrideOnConflictOk returns a tuple with the OverrideOnConflict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetOverrideOnConflictOk() (*bool, bool) {
	if o == nil || o.OverrideOnConflict == nil {
		return nil, false
	}
	return o.OverrideOnConflict, true
}

// HasOverrideOnConflict returns a boolean if a field has been set.
func (o *LogsSchemaRemapper) HasOverrideOnConflict() bool {
	return o != nil && o.OverrideOnConflict != nil
}

// SetOverrideOnConflict gets a reference to the given bool and assigns it to the OverrideOnConflict field.
func (o *LogsSchemaRemapper) SetOverrideOnConflict(v bool) {
	o.OverrideOnConflict = &v
}

// GetPreserveSource returns the PreserveSource field value if set, zero value otherwise.
func (o *LogsSchemaRemapper) GetPreserveSource() bool {
	if o == nil || o.PreserveSource == nil {
		var ret bool
		return ret
	}
	return *o.PreserveSource
}

// GetPreserveSourceOk returns a tuple with the PreserveSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetPreserveSourceOk() (*bool, bool) {
	if o == nil || o.PreserveSource == nil {
		return nil, false
	}
	return o.PreserveSource, true
}

// HasPreserveSource returns a boolean if a field has been set.
func (o *LogsSchemaRemapper) HasPreserveSource() bool {
	return o != nil && o.PreserveSource != nil
}

// SetPreserveSource gets a reference to the given bool and assigns it to the PreserveSource field.
func (o *LogsSchemaRemapper) SetPreserveSource(v bool) {
	o.PreserveSource = &v
}

// GetSources returns the Sources field value.
func (o *LogsSchemaRemapper) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *LogsSchemaRemapper) SetSources(v []string) {
	o.Sources = v
}

// GetTarget returns the Target field value.
func (o *LogsSchemaRemapper) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsSchemaRemapper) SetTarget(v string) {
	o.Target = v
}

// GetTargetFormat returns the TargetFormat field value if set, zero value otherwise.
func (o *LogsSchemaRemapper) GetTargetFormat() TargetFormatType {
	if o == nil || o.TargetFormat == nil {
		var ret TargetFormatType
		return ret
	}
	return *o.TargetFormat
}

// GetTargetFormatOk returns a tuple with the TargetFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetTargetFormatOk() (*TargetFormatType, bool) {
	if o == nil || o.TargetFormat == nil {
		return nil, false
	}
	return o.TargetFormat, true
}

// HasTargetFormat returns a boolean if a field has been set.
func (o *LogsSchemaRemapper) HasTargetFormat() bool {
	return o != nil && o.TargetFormat != nil
}

// SetTargetFormat gets a reference to the given TargetFormatType and assigns it to the TargetFormat field.
func (o *LogsSchemaRemapper) SetTargetFormat(v TargetFormatType) {
	o.TargetFormat = &v
}

// GetType returns the Type field value.
func (o *LogsSchemaRemapper) GetType() LogsSchemaRemapperType {
	if o == nil {
		var ret LogsSchemaRemapperType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaRemapper) GetTypeOk() (*LogsSchemaRemapperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsSchemaRemapper) SetType(v LogsSchemaRemapperType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsSchemaRemapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.OverrideOnConflict != nil {
		toSerialize["override_on_conflict"] = o.OverrideOnConflict
	}
	if o.PreserveSource != nil {
		toSerialize["preserve_source"] = o.PreserveSource
	}
	toSerialize["sources"] = o.Sources
	toSerialize["target"] = o.Target
	if o.TargetFormat != nil {
		toSerialize["target_format"] = o.TargetFormat
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsSchemaRemapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name               *string                 `json:"name"`
		OverrideOnConflict *bool                   `json:"override_on_conflict,omitempty"`
		PreserveSource     *bool                   `json:"preserve_source,omitempty"`
		Sources            *[]string               `json:"sources"`
		Target             *string                 `json:"target"`
		TargetFormat       *TargetFormatType       `json:"target_format,omitempty"`
		Type               *LogsSchemaRemapperType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "override_on_conflict", "preserve_source", "sources", "target", "target_format", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.OverrideOnConflict = all.OverrideOnConflict
	o.PreserveSource = all.PreserveSource
	o.Sources = *all.Sources
	o.Target = *all.Target
	if all.TargetFormat != nil && !all.TargetFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.TargetFormat = all.TargetFormat
	}
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
