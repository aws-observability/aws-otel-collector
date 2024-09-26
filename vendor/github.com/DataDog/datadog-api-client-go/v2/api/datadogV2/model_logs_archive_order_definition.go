// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveOrderDefinition The definition of an archive order.
type LogsArchiveOrderDefinition struct {
	// The attributes associated with the archive order.
	Attributes LogsArchiveOrderAttributes `json:"attributes"`
	// Type of the archive order definition.
	Type LogsArchiveOrderDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveOrderDefinition instantiates a new LogsArchiveOrderDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveOrderDefinition(attributes LogsArchiveOrderAttributes, typeVar LogsArchiveOrderDefinitionType) *LogsArchiveOrderDefinition {
	this := LogsArchiveOrderDefinition{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewLogsArchiveOrderDefinitionWithDefaults instantiates a new LogsArchiveOrderDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveOrderDefinitionWithDefaults() *LogsArchiveOrderDefinition {
	this := LogsArchiveOrderDefinition{}
	var typeVar LogsArchiveOrderDefinitionType = LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *LogsArchiveOrderDefinition) GetAttributes() LogsArchiveOrderAttributes {
	if o == nil {
		var ret LogsArchiveOrderAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveOrderDefinition) GetAttributesOk() (*LogsArchiveOrderAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *LogsArchiveOrderDefinition) SetAttributes(v LogsArchiveOrderAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *LogsArchiveOrderDefinition) GetType() LogsArchiveOrderDefinitionType {
	if o == nil {
		var ret LogsArchiveOrderDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveOrderDefinition) GetTypeOk() (*LogsArchiveOrderDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArchiveOrderDefinition) SetType(v LogsArchiveOrderDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveOrderDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveOrderDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *LogsArchiveOrderAttributes     `json:"attributes"`
		Type       *LogsArchiveOrderDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
