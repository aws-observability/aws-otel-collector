// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMappingResponseDataAttributesAttributesItems
type GetMappingResponseDataAttributesAttributesItems struct {
	//
	Attribute *string `json:"attribute,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	DisplayName *string `json:"display_name,omitempty"`
	//
	Groups []string `json:"groups,omitempty"`
	//
	IsCustom *bool `json:"is_custom,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMappingResponseDataAttributesAttributesItems instantiates a new GetMappingResponseDataAttributesAttributesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMappingResponseDataAttributesAttributesItems() *GetMappingResponseDataAttributesAttributesItems {
	this := GetMappingResponseDataAttributesAttributesItems{}
	return &this
}

// NewGetMappingResponseDataAttributesAttributesItemsWithDefaults instantiates a new GetMappingResponseDataAttributesAttributesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMappingResponseDataAttributesAttributesItemsWithDefaults() *GetMappingResponseDataAttributesAttributesItems {
	this := GetMappingResponseDataAttributesAttributesItems{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributesAttributesItems) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *GetMappingResponseDataAttributesAttributesItems) SetAttribute(v string) {
	o.Attribute = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributesAttributesItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetMappingResponseDataAttributesAttributesItems) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributesAttributesItems) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GetMappingResponseDataAttributesAttributesItems) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributesAttributesItems) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *GetMappingResponseDataAttributesAttributesItems) SetGroups(v []string) {
	o.Groups = v
}

// GetIsCustom returns the IsCustom field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributesAttributesItems) GetIsCustom() bool {
	if o == nil || o.IsCustom == nil {
		var ret bool
		return ret
	}
	return *o.IsCustom
}

// GetIsCustomOk returns a tuple with the IsCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) GetIsCustomOk() (*bool, bool) {
	if o == nil || o.IsCustom == nil {
		return nil, false
	}
	return o.IsCustom, true
}

// HasIsCustom returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) HasIsCustom() bool {
	return o != nil && o.IsCustom != nil
}

// SetIsCustom gets a reference to the given bool and assigns it to the IsCustom field.
func (o *GetMappingResponseDataAttributesAttributesItems) SetIsCustom(v bool) {
	o.IsCustom = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetMappingResponseDataAttributesAttributesItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetMappingResponseDataAttributesAttributesItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetMappingResponseDataAttributesAttributesItems) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMappingResponseDataAttributesAttributesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.IsCustom != nil {
		toSerialize["is_custom"] = o.IsCustom
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
func (o *GetMappingResponseDataAttributesAttributesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute   *string  `json:"attribute,omitempty"`
		Description *string  `json:"description,omitempty"`
		DisplayName *string  `json:"display_name,omitempty"`
		Groups      []string `json:"groups,omitempty"`
		IsCustom    *bool    `json:"is_custom,omitempty"`
		Type        *string  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "description", "display_name", "groups", "is_custom", "type"})
	} else {
		return err
	}
	o.Attribute = all.Attribute
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.Groups = all.Groups
	o.IsCustom = all.IsCustom
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
