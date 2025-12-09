// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItems
type GetMultipleRulesetsResponseDataAttributesRulesetsItems struct {
	//
	Data GetMultipleRulesetsResponseDataAttributesRulesetsItemsData `json:"data"`
	//
	Description *string `json:"description,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Rules []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems `json:"rules,omitempty"`
	//
	ShortDescription *string `json:"short_description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItems instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItems(data GetMultipleRulesetsResponseDataAttributesRulesetsItemsData) *GetMultipleRulesetsResponseDataAttributesRulesetsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItems{}
	this.Data = data
	return &this
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsWithDefaults instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsWithDefaults() *GetMultipleRulesetsResponseDataAttributesRulesetsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItems{}
	return &this
}

// GetData returns the Data field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetData() GetMultipleRulesetsResponseDataAttributesRulesetsItemsData {
	if o == nil {
		var ret GetMultipleRulesetsResponseDataAttributesRulesetsItemsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetDataOk() (*GetMultipleRulesetsResponseDataAttributesRulesetsItemsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetData(v GetMultipleRulesetsResponseDataAttributesRulesetsItemsData) {
	o.Data = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetRules() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	if o == nil || o.Rules == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetRulesOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems and assigns it to the Rules field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetRules(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) {
	o.Rules = v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) HasShortDescription() bool {
	return o != nil && o.ShortDescription != nil
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseDataAttributesRulesetsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data             *GetMultipleRulesetsResponseDataAttributesRulesetsItemsData        `json:"data"`
		Description      *string                                                            `json:"description,omitempty"`
		Name             *string                                                            `json:"name,omitempty"`
		Rules            []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems `json:"rules,omitempty"`
		ShortDescription *string                                                            `json:"short_description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "description", "name", "rules", "short_description"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Description = all.Description
	o.Name = all.Name
	o.Rules = all.Rules
	o.ShortDescription = all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
