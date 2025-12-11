// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeConfigAttributesCreate Custom attribute config resource attributes
type CustomAttributeConfigAttributesCreate struct {
	// Custom attribute description.
	Description *string `json:"description,omitempty"`
	// Custom attribute name.
	DisplayName string `json:"display_name"`
	// Whether multiple values can be set
	IsMulti bool `json:"is_multi"`
	// Custom attribute key. This will be the value use to search on this custom attribute
	Key string `json:"key"`
	// Custom attributes type
	Type CustomAttributeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAttributeConfigAttributesCreate instantiates a new CustomAttributeConfigAttributesCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAttributeConfigAttributesCreate(displayName string, isMulti bool, key string, typeVar CustomAttributeType) *CustomAttributeConfigAttributesCreate {
	this := CustomAttributeConfigAttributesCreate{}
	this.DisplayName = displayName
	this.IsMulti = isMulti
	this.Key = key
	this.Type = typeVar
	return &this
}

// NewCustomAttributeConfigAttributesCreateWithDefaults instantiates a new CustomAttributeConfigAttributesCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAttributeConfigAttributesCreateWithDefaults() *CustomAttributeConfigAttributesCreate {
	this := CustomAttributeConfigAttributesCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomAttributeConfigAttributesCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigAttributesCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomAttributeConfigAttributesCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomAttributeConfigAttributesCreate) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value.
func (o *CustomAttributeConfigAttributesCreate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigAttributesCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *CustomAttributeConfigAttributesCreate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIsMulti returns the IsMulti field value.
func (o *CustomAttributeConfigAttributesCreate) GetIsMulti() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsMulti
}

// GetIsMultiOk returns a tuple with the IsMulti field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigAttributesCreate) GetIsMultiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMulti, true
}

// SetIsMulti sets field value.
func (o *CustomAttributeConfigAttributesCreate) SetIsMulti(v bool) {
	o.IsMulti = v
}

// GetKey returns the Key field value.
func (o *CustomAttributeConfigAttributesCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigAttributesCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *CustomAttributeConfigAttributesCreate) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value.
func (o *CustomAttributeConfigAttributesCreate) GetType() CustomAttributeType {
	if o == nil {
		var ret CustomAttributeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigAttributesCreate) GetTypeOk() (*CustomAttributeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomAttributeConfigAttributesCreate) SetType(v CustomAttributeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAttributeConfigAttributesCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["display_name"] = o.DisplayName
	toSerialize["is_multi"] = o.IsMulti
	toSerialize["key"] = o.Key
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomAttributeConfigAttributesCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string              `json:"description,omitempty"`
		DisplayName *string              `json:"display_name"`
		IsMulti     *bool                `json:"is_multi"`
		Key         *string              `json:"key"`
		Type        *CustomAttributeType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.IsMulti == nil {
		return fmt.Errorf("required field is_multi missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "is_multi", "key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.DisplayName = *all.DisplayName
	o.IsMulti = *all.IsMulti
	o.Key = *all.Key
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
