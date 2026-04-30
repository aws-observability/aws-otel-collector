// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseXMLProcessor The `parse_xml` processor parses XML from a specified field and extracts it into the event.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineParseXMLProcessor struct {
	// Whether to always use a text key for element content.
	AlwaysUseTextKey *bool `json:"always_use_text_key,omitempty"`
	// The prefix to use for XML attributes in the parsed output.
	AttrPrefix *string `json:"attr_prefix,omitempty"`
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The name of the log field that contains an XML string.
	Field string `json:"field"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// Whether to include XML attributes in the parsed output.
	IncludeAttr *bool `json:"include_attr,omitempty"`
	// Whether to parse boolean values from strings.
	ParseBool *bool `json:"parse_bool,omitempty"`
	// Whether to parse null values.
	ParseNull *bool `json:"parse_null,omitempty"`
	// Whether to parse numeric values from strings.
	ParseNumber *bool `json:"parse_number,omitempty"`
	// The key name to use for text content within XML elements. Must be at least 1 character if specified.
	TextKey *string `json:"text_key,omitempty"`
	// The processor type. The value should always be `parse_xml`.
	Type ObservabilityPipelineParseXMLProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineParseXMLProcessor instantiates a new ObservabilityPipelineParseXMLProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineParseXMLProcessor(enabled bool, field string, id string, include string, typeVar ObservabilityPipelineParseXMLProcessorType) *ObservabilityPipelineParseXMLProcessor {
	this := ObservabilityPipelineParseXMLProcessor{}
	this.Enabled = enabled
	this.Field = field
	this.Id = id
	this.Include = include
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineParseXMLProcessorWithDefaults instantiates a new ObservabilityPipelineParseXMLProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineParseXMLProcessorWithDefaults() *ObservabilityPipelineParseXMLProcessor {
	this := ObservabilityPipelineParseXMLProcessor{}
	var typeVar ObservabilityPipelineParseXMLProcessorType = OBSERVABILITYPIPELINEPARSEXMLPROCESSORTYPE_PARSE_XML
	this.Type = typeVar
	return &this
}

// GetAlwaysUseTextKey returns the AlwaysUseTextKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetAlwaysUseTextKey() bool {
	if o == nil || o.AlwaysUseTextKey == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysUseTextKey
}

// GetAlwaysUseTextKeyOk returns a tuple with the AlwaysUseTextKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetAlwaysUseTextKeyOk() (*bool, bool) {
	if o == nil || o.AlwaysUseTextKey == nil {
		return nil, false
	}
	return o.AlwaysUseTextKey, true
}

// HasAlwaysUseTextKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasAlwaysUseTextKey() bool {
	return o != nil && o.AlwaysUseTextKey != nil
}

// SetAlwaysUseTextKey gets a reference to the given bool and assigns it to the AlwaysUseTextKey field.
func (o *ObservabilityPipelineParseXMLProcessor) SetAlwaysUseTextKey(v bool) {
	o.AlwaysUseTextKey = &v
}

// GetAttrPrefix returns the AttrPrefix field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetAttrPrefix() string {
	if o == nil || o.AttrPrefix == nil {
		var ret string
		return ret
	}
	return *o.AttrPrefix
}

// GetAttrPrefixOk returns a tuple with the AttrPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetAttrPrefixOk() (*string, bool) {
	if o == nil || o.AttrPrefix == nil {
		return nil, false
	}
	return o.AttrPrefix, true
}

// HasAttrPrefix returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasAttrPrefix() bool {
	return o != nil && o.AttrPrefix != nil
}

// SetAttrPrefix gets a reference to the given string and assigns it to the AttrPrefix field.
func (o *ObservabilityPipelineParseXMLProcessor) SetAttrPrefix(v string) {
	o.AttrPrefix = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineParseXMLProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineParseXMLProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineParseXMLProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetField returns the Field field value.
func (o *ObservabilityPipelineParseXMLProcessor) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ObservabilityPipelineParseXMLProcessor) SetField(v string) {
	o.Field = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineParseXMLProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineParseXMLProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineParseXMLProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineParseXMLProcessor) SetInclude(v string) {
	o.Include = v
}

// GetIncludeAttr returns the IncludeAttr field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetIncludeAttr() bool {
	if o == nil || o.IncludeAttr == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAttr
}

// GetIncludeAttrOk returns a tuple with the IncludeAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetIncludeAttrOk() (*bool, bool) {
	if o == nil || o.IncludeAttr == nil {
		return nil, false
	}
	return o.IncludeAttr, true
}

// HasIncludeAttr returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasIncludeAttr() bool {
	return o != nil && o.IncludeAttr != nil
}

// SetIncludeAttr gets a reference to the given bool and assigns it to the IncludeAttr field.
func (o *ObservabilityPipelineParseXMLProcessor) SetIncludeAttr(v bool) {
	o.IncludeAttr = &v
}

// GetParseBool returns the ParseBool field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetParseBool() bool {
	if o == nil || o.ParseBool == nil {
		var ret bool
		return ret
	}
	return *o.ParseBool
}

// GetParseBoolOk returns a tuple with the ParseBool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetParseBoolOk() (*bool, bool) {
	if o == nil || o.ParseBool == nil {
		return nil, false
	}
	return o.ParseBool, true
}

// HasParseBool returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasParseBool() bool {
	return o != nil && o.ParseBool != nil
}

// SetParseBool gets a reference to the given bool and assigns it to the ParseBool field.
func (o *ObservabilityPipelineParseXMLProcessor) SetParseBool(v bool) {
	o.ParseBool = &v
}

// GetParseNull returns the ParseNull field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetParseNull() bool {
	if o == nil || o.ParseNull == nil {
		var ret bool
		return ret
	}
	return *o.ParseNull
}

// GetParseNullOk returns a tuple with the ParseNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetParseNullOk() (*bool, bool) {
	if o == nil || o.ParseNull == nil {
		return nil, false
	}
	return o.ParseNull, true
}

// HasParseNull returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasParseNull() bool {
	return o != nil && o.ParseNull != nil
}

// SetParseNull gets a reference to the given bool and assigns it to the ParseNull field.
func (o *ObservabilityPipelineParseXMLProcessor) SetParseNull(v bool) {
	o.ParseNull = &v
}

// GetParseNumber returns the ParseNumber field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetParseNumber() bool {
	if o == nil || o.ParseNumber == nil {
		var ret bool
		return ret
	}
	return *o.ParseNumber
}

// GetParseNumberOk returns a tuple with the ParseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetParseNumberOk() (*bool, bool) {
	if o == nil || o.ParseNumber == nil {
		return nil, false
	}
	return o.ParseNumber, true
}

// HasParseNumber returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasParseNumber() bool {
	return o != nil && o.ParseNumber != nil
}

// SetParseNumber gets a reference to the given bool and assigns it to the ParseNumber field.
func (o *ObservabilityPipelineParseXMLProcessor) SetParseNumber(v bool) {
	o.ParseNumber = &v
}

// GetTextKey returns the TextKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineParseXMLProcessor) GetTextKey() string {
	if o == nil || o.TextKey == nil {
		var ret string
		return ret
	}
	return *o.TextKey
}

// GetTextKeyOk returns a tuple with the TextKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetTextKeyOk() (*string, bool) {
	if o == nil || o.TextKey == nil {
		return nil, false
	}
	return o.TextKey, true
}

// HasTextKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineParseXMLProcessor) HasTextKey() bool {
	return o != nil && o.TextKey != nil
}

// SetTextKey gets a reference to the given string and assigns it to the TextKey field.
func (o *ObservabilityPipelineParseXMLProcessor) SetTextKey(v string) {
	o.TextKey = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineParseXMLProcessor) GetType() ObservabilityPipelineParseXMLProcessorType {
	if o == nil {
		var ret ObservabilityPipelineParseXMLProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineParseXMLProcessor) GetTypeOk() (*ObservabilityPipelineParseXMLProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineParseXMLProcessor) SetType(v ObservabilityPipelineParseXMLProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineParseXMLProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlwaysUseTextKey != nil {
		toSerialize["always_use_text_key"] = o.AlwaysUseTextKey
	}
	if o.AttrPrefix != nil {
		toSerialize["attr_prefix"] = o.AttrPrefix
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["field"] = o.Field
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	if o.IncludeAttr != nil {
		toSerialize["include_attr"] = o.IncludeAttr
	}
	if o.ParseBool != nil {
		toSerialize["parse_bool"] = o.ParseBool
	}
	if o.ParseNull != nil {
		toSerialize["parse_null"] = o.ParseNull
	}
	if o.ParseNumber != nil {
		toSerialize["parse_number"] = o.ParseNumber
	}
	if o.TextKey != nil {
		toSerialize["text_key"] = o.TextKey
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineParseXMLProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlwaysUseTextKey *bool                                       `json:"always_use_text_key,omitempty"`
		AttrPrefix       *string                                     `json:"attr_prefix,omitempty"`
		DisplayName      *string                                     `json:"display_name,omitempty"`
		Enabled          *bool                                       `json:"enabled"`
		Field            *string                                     `json:"field"`
		Id               *string                                     `json:"id"`
		Include          *string                                     `json:"include"`
		IncludeAttr      *bool                                       `json:"include_attr,omitempty"`
		ParseBool        *bool                                       `json:"parse_bool,omitempty"`
		ParseNull        *bool                                       `json:"parse_null,omitempty"`
		ParseNumber      *bool                                       `json:"parse_number,omitempty"`
		TextKey          *string                                     `json:"text_key,omitempty"`
		Type             *ObservabilityPipelineParseXMLProcessorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"always_use_text_key", "attr_prefix", "display_name", "enabled", "field", "id", "include", "include_attr", "parse_bool", "parse_null", "parse_number", "text_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AlwaysUseTextKey = all.AlwaysUseTextKey
	o.AttrPrefix = all.AttrPrefix
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Field = *all.Field
	o.Id = *all.Id
	o.Include = *all.Include
	o.IncludeAttr = all.IncludeAttr
	o.ParseBool = all.ParseBool
	o.ParseNull = all.ParseNull
	o.ParseNumber = all.ParseNumber
	o.TextKey = all.TextKey
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
