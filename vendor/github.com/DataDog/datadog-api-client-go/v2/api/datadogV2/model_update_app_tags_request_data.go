// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppTagsRequestData Data for replacing an app's tags.
type UpdateAppTagsRequestData struct {
	// Attributes for replacing an app's tags.
	Attributes *UpdateAppTagsRequestDataAttributes `json:"attributes,omitempty"`
	// The tags resource type.
	Type *AppTagsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppTagsRequestData instantiates a new UpdateAppTagsRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppTagsRequestData() *UpdateAppTagsRequestData {
	this := UpdateAppTagsRequestData{}
	var typeVar AppTagsType = APPTAGSTYPE_TAGS
	this.Type = &typeVar
	return &this
}

// NewUpdateAppTagsRequestDataWithDefaults instantiates a new UpdateAppTagsRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppTagsRequestDataWithDefaults() *UpdateAppTagsRequestData {
	this := UpdateAppTagsRequestData{}
	var typeVar AppTagsType = APPTAGSTYPE_TAGS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateAppTagsRequestData) GetAttributes() UpdateAppTagsRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UpdateAppTagsRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppTagsRequestData) GetAttributesOk() (*UpdateAppTagsRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateAppTagsRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UpdateAppTagsRequestDataAttributes and assigns it to the Attributes field.
func (o *UpdateAppTagsRequestData) SetAttributes(v UpdateAppTagsRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateAppTagsRequestData) GetType() AppTagsType {
	if o == nil || o.Type == nil {
		var ret AppTagsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppTagsRequestData) GetTypeOk() (*AppTagsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateAppTagsRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AppTagsType and assigns it to the Type field.
func (o *UpdateAppTagsRequestData) SetType(v AppTagsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppTagsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppTagsRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpdateAppTagsRequestDataAttributes `json:"attributes,omitempty"`
		Type       *AppTagsType                        `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
