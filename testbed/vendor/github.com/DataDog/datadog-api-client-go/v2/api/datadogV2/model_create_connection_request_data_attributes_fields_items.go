// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateConnectionRequestDataAttributesFieldsItems
type CreateConnectionRequestDataAttributesFieldsItems struct {
	//
	Description *string `json:"description,omitempty"`
	//
	DisplayName *string `json:"display_name,omitempty"`
	//
	Groups []string `json:"groups,omitempty"`
	//
	Id string `json:"id"`
	//
	SourceName string `json:"source_name"`
	//
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateConnectionRequestDataAttributesFieldsItems instantiates a new CreateConnectionRequestDataAttributesFieldsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateConnectionRequestDataAttributesFieldsItems(id string, sourceName string, typeVar string) *CreateConnectionRequestDataAttributesFieldsItems {
	this := CreateConnectionRequestDataAttributesFieldsItems{}
	this.Id = id
	this.SourceName = sourceName
	this.Type = typeVar
	return &this
}

// NewCreateConnectionRequestDataAttributesFieldsItemsWithDefaults instantiates a new CreateConnectionRequestDataAttributesFieldsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateConnectionRequestDataAttributesFieldsItemsWithDefaults() *CreateConnectionRequestDataAttributesFieldsItems {
	this := CreateConnectionRequestDataAttributesFieldsItems{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateConnectionRequestDataAttributesFieldsItems) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateConnectionRequestDataAttributesFieldsItems) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *CreateConnectionRequestDataAttributesFieldsItems) SetGroups(v []string) {
	o.Groups = v
}

// GetId returns the Id field value.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CreateConnectionRequestDataAttributesFieldsItems) SetId(v string) {
	o.Id = v
}

// GetSourceName returns the SourceName field value.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetSourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceName, true
}

// SetSourceName sets field value.
func (o *CreateConnectionRequestDataAttributesFieldsItems) SetSourceName(v string) {
	o.SourceName = v
}

// GetType returns the Type field value.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributesFieldsItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateConnectionRequestDataAttributesFieldsItems) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateConnectionRequestDataAttributesFieldsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	toSerialize["id"] = o.Id
	toSerialize["source_name"] = o.SourceName
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateConnectionRequestDataAttributesFieldsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description,omitempty"`
		DisplayName *string  `json:"display_name,omitempty"`
		Groups      []string `json:"groups,omitempty"`
		Id          *string  `json:"id"`
		SourceName  *string  `json:"source_name"`
		Type        *string  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.SourceName == nil {
		return fmt.Errorf("required field source_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "groups", "id", "source_name", "type"})
	} else {
		return err
	}
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.Groups = all.Groups
	o.Id = *all.Id
	o.SourceName = *all.SourceName
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
