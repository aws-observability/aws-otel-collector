// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeConfigUpdateAttributes Attributes that can be updated on a custom attribute configuration. All fields are optional; only provided fields are changed.
type CustomAttributeConfigUpdateAttributes struct {
	// A description explaining the purpose and expected values for this custom attribute.
	Description *string `json:"description,omitempty"`
	// The human-readable label shown in the Case Management UI for this custom attribute.
	DisplayName *string `json:"display_name,omitempty"`
	// An external field identifier to auto-populate this attribute from (used for integrations with external systems).
	MapFrom *string `json:"map_from,omitempty"`
	// The data type of the custom attribute, which determines the allowed values and UI input control.
	Type *CustomAttributeType `json:"type,omitempty"`
	// Type-specific configuration for the custom attribute. For SELECT-type attributes, this contains the list of allowed options.
	TypeData *CustomAttributeTypeData `json:"type_data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomAttributeConfigUpdateAttributes instantiates a new CustomAttributeConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomAttributeConfigUpdateAttributes() *CustomAttributeConfigUpdateAttributes {
	this := CustomAttributeConfigUpdateAttributes{}
	return &this
}

// NewCustomAttributeConfigUpdateAttributesWithDefaults instantiates a new CustomAttributeConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomAttributeConfigUpdateAttributesWithDefaults() *CustomAttributeConfigUpdateAttributes {
	this := CustomAttributeConfigUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomAttributeConfigUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomAttributeConfigUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomAttributeConfigUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CustomAttributeConfigUpdateAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigUpdateAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CustomAttributeConfigUpdateAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CustomAttributeConfigUpdateAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMapFrom returns the MapFrom field value if set, zero value otherwise.
func (o *CustomAttributeConfigUpdateAttributes) GetMapFrom() string {
	if o == nil || o.MapFrom == nil {
		var ret string
		return ret
	}
	return *o.MapFrom
}

// GetMapFromOk returns a tuple with the MapFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigUpdateAttributes) GetMapFromOk() (*string, bool) {
	if o == nil || o.MapFrom == nil {
		return nil, false
	}
	return o.MapFrom, true
}

// HasMapFrom returns a boolean if a field has been set.
func (o *CustomAttributeConfigUpdateAttributes) HasMapFrom() bool {
	return o != nil && o.MapFrom != nil
}

// SetMapFrom gets a reference to the given string and assigns it to the MapFrom field.
func (o *CustomAttributeConfigUpdateAttributes) SetMapFrom(v string) {
	o.MapFrom = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomAttributeConfigUpdateAttributes) GetType() CustomAttributeType {
	if o == nil || o.Type == nil {
		var ret CustomAttributeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigUpdateAttributes) GetTypeOk() (*CustomAttributeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomAttributeConfigUpdateAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CustomAttributeType and assigns it to the Type field.
func (o *CustomAttributeConfigUpdateAttributes) SetType(v CustomAttributeType) {
	o.Type = &v
}

// GetTypeData returns the TypeData field value if set, zero value otherwise.
func (o *CustomAttributeConfigUpdateAttributes) GetTypeData() CustomAttributeTypeData {
	if o == nil || o.TypeData == nil {
		var ret CustomAttributeTypeData
		return ret
	}
	return *o.TypeData
}

// GetTypeDataOk returns a tuple with the TypeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeConfigUpdateAttributes) GetTypeDataOk() (*CustomAttributeTypeData, bool) {
	if o == nil || o.TypeData == nil {
		return nil, false
	}
	return o.TypeData, true
}

// HasTypeData returns a boolean if a field has been set.
func (o *CustomAttributeConfigUpdateAttributes) HasTypeData() bool {
	return o != nil && o.TypeData != nil
}

// SetTypeData gets a reference to the given CustomAttributeTypeData and assigns it to the TypeData field.
func (o *CustomAttributeConfigUpdateAttributes) SetTypeData(v CustomAttributeTypeData) {
	o.TypeData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomAttributeConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.MapFrom != nil {
		toSerialize["map_from"] = o.MapFrom
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TypeData != nil {
		toSerialize["type_data"] = o.TypeData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomAttributeConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                  `json:"description,omitempty"`
		DisplayName *string                  `json:"display_name,omitempty"`
		MapFrom     *string                  `json:"map_from,omitempty"`
		Type        *CustomAttributeType     `json:"type,omitempty"`
		TypeData    *CustomAttributeTypeData `json:"type_data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "map_from", "type", "type_data"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.MapFrom = all.MapFrom
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	if all.TypeData != nil && all.TypeData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TypeData = all.TypeData

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
