// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertFormVersionData The data for creating or updating a form version.
type UpsertFormVersionData struct {
	// The attributes for creating or updating a form version.
	Attributes UpsertFormVersionDataAttributes `json:"attributes"`
	// The resource type for a form version.
	Type FormVersionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertFormVersionData instantiates a new UpsertFormVersionData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertFormVersionData(attributes UpsertFormVersionDataAttributes, typeVar FormVersionType) *UpsertFormVersionData {
	this := UpsertFormVersionData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewUpsertFormVersionDataWithDefaults instantiates a new UpsertFormVersionData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertFormVersionDataWithDefaults() *UpsertFormVersionData {
	this := UpsertFormVersionData{}
	var typeVar FormVersionType = FORMVERSIONTYPE_FORM_VERSIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *UpsertFormVersionData) GetAttributes() UpsertFormVersionDataAttributes {
	if o == nil {
		var ret UpsertFormVersionDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionData) GetAttributesOk() (*UpsertFormVersionDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *UpsertFormVersionData) SetAttributes(v UpsertFormVersionDataAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *UpsertFormVersionData) GetType() FormVersionType {
	if o == nil {
		var ret FormVersionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionData) GetTypeOk() (*FormVersionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpsertFormVersionData) SetType(v FormVersionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertFormVersionData) MarshalJSON() ([]byte, error) {
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
func (o *UpsertFormVersionData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpsertFormVersionDataAttributes `json:"attributes"`
		Type       *FormVersionType                 `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
