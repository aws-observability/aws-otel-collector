// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScriptDataAttributes The definition of `ScriptDataAttributes` object.
type ScriptDataAttributes struct {
	// The `attributes` `name`.
	Name *string `json:"name,omitempty"`
	// The `attributes` `src`.
	Src *string `json:"src,omitempty"`
	// The `attributes` `type`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScriptDataAttributes instantiates a new ScriptDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScriptDataAttributes() *ScriptDataAttributes {
	this := ScriptDataAttributes{}
	return &this
}

// NewScriptDataAttributesWithDefaults instantiates a new ScriptDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScriptDataAttributesWithDefaults() *ScriptDataAttributes {
	this := ScriptDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScriptDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScriptDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScriptDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *ScriptDataAttributes) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptDataAttributes) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *ScriptDataAttributes) HasSrc() bool {
	return o != nil && o.Src != nil
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *ScriptDataAttributes) SetSrc(v string) {
	o.Src = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScriptDataAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScriptDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScriptDataAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScriptDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Src != nil {
		toSerialize["src"] = o.Src
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
func (o *ScriptDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name,omitempty"`
		Src  *string `json:"src,omitempty"`
		Type *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "src", "type"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Src = all.Src
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
