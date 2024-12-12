// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsGlobalVariableRequest Details of the global variable to create.
type SyntheticsGlobalVariableRequest struct {
	// Attributes of the global variable.
	Attributes *SyntheticsGlobalVariableAttributes `json:"attributes,omitempty"`
	// Description of the global variable.
	Description string `json:"description"`
	// Unique identifier of the global variable.
	Id *string `json:"id,omitempty"`
	// Determines if the global variable is a FIDO variable.
	IsFido *bool `json:"is_fido,omitempty"`
	// Determines if the global variable is a TOTP/MFA variable.
	IsTotp *bool `json:"is_totp,omitempty"`
	// Name of the global variable. Unique across Synthetic global variables.
	Name string `json:"name"`
	// Parser options to use for retrieving a Synthetic global variable from a Synthetic test. Used in conjunction with `parse_test_public_id`.
	ParseTestOptions *SyntheticsGlobalVariableParseTestOptions `json:"parse_test_options,omitempty"`
	// A Synthetic test ID to use as a test to generate the variable value.
	ParseTestPublicId *string `json:"parse_test_public_id,omitempty"`
	// Tags of the global variable.
	Tags []string `json:"tags"`
	// Value of the global variable.
	Value *SyntheticsGlobalVariableValue `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsGlobalVariableRequest instantiates a new SyntheticsGlobalVariableRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsGlobalVariableRequest(description string, name string, tags []string) *SyntheticsGlobalVariableRequest {
	this := SyntheticsGlobalVariableRequest{}
	this.Description = description
	this.Name = name
	this.Tags = tags
	return &this
}

// NewSyntheticsGlobalVariableRequestWithDefaults instantiates a new SyntheticsGlobalVariableRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsGlobalVariableRequestWithDefaults() *SyntheticsGlobalVariableRequest {
	this := SyntheticsGlobalVariableRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetAttributes() SyntheticsGlobalVariableAttributes {
	if o == nil || o.Attributes == nil {
		var ret SyntheticsGlobalVariableAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetAttributesOk() (*SyntheticsGlobalVariableAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SyntheticsGlobalVariableAttributes and assigns it to the Attributes field.
func (o *SyntheticsGlobalVariableRequest) SetAttributes(v SyntheticsGlobalVariableAttributes) {
	o.Attributes = &v
}

// GetDescription returns the Description field value.
func (o *SyntheticsGlobalVariableRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SyntheticsGlobalVariableRequest) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsGlobalVariableRequest) SetId(v string) {
	o.Id = &v
}

// GetIsFido returns the IsFido field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetIsFido() bool {
	if o == nil || o.IsFido == nil {
		var ret bool
		return ret
	}
	return *o.IsFido
}

// GetIsFidoOk returns a tuple with the IsFido field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetIsFidoOk() (*bool, bool) {
	if o == nil || o.IsFido == nil {
		return nil, false
	}
	return o.IsFido, true
}

// HasIsFido returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasIsFido() bool {
	return o != nil && o.IsFido != nil
}

// SetIsFido gets a reference to the given bool and assigns it to the IsFido field.
func (o *SyntheticsGlobalVariableRequest) SetIsFido(v bool) {
	o.IsFido = &v
}

// GetIsTotp returns the IsTotp field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetIsTotp() bool {
	if o == nil || o.IsTotp == nil {
		var ret bool
		return ret
	}
	return *o.IsTotp
}

// GetIsTotpOk returns a tuple with the IsTotp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetIsTotpOk() (*bool, bool) {
	if o == nil || o.IsTotp == nil {
		return nil, false
	}
	return o.IsTotp, true
}

// HasIsTotp returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasIsTotp() bool {
	return o != nil && o.IsTotp != nil
}

// SetIsTotp gets a reference to the given bool and assigns it to the IsTotp field.
func (o *SyntheticsGlobalVariableRequest) SetIsTotp(v bool) {
	o.IsTotp = &v
}

// GetName returns the Name field value.
func (o *SyntheticsGlobalVariableRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsGlobalVariableRequest) SetName(v string) {
	o.Name = v
}

// GetParseTestOptions returns the ParseTestOptions field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetParseTestOptions() SyntheticsGlobalVariableParseTestOptions {
	if o == nil || o.ParseTestOptions == nil {
		var ret SyntheticsGlobalVariableParseTestOptions
		return ret
	}
	return *o.ParseTestOptions
}

// GetParseTestOptionsOk returns a tuple with the ParseTestOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetParseTestOptionsOk() (*SyntheticsGlobalVariableParseTestOptions, bool) {
	if o == nil || o.ParseTestOptions == nil {
		return nil, false
	}
	return o.ParseTestOptions, true
}

// HasParseTestOptions returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasParseTestOptions() bool {
	return o != nil && o.ParseTestOptions != nil
}

// SetParseTestOptions gets a reference to the given SyntheticsGlobalVariableParseTestOptions and assigns it to the ParseTestOptions field.
func (o *SyntheticsGlobalVariableRequest) SetParseTestOptions(v SyntheticsGlobalVariableParseTestOptions) {
	o.ParseTestOptions = &v
}

// GetParseTestPublicId returns the ParseTestPublicId field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetParseTestPublicId() string {
	if o == nil || o.ParseTestPublicId == nil {
		var ret string
		return ret
	}
	return *o.ParseTestPublicId
}

// GetParseTestPublicIdOk returns a tuple with the ParseTestPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetParseTestPublicIdOk() (*string, bool) {
	if o == nil || o.ParseTestPublicId == nil {
		return nil, false
	}
	return o.ParseTestPublicId, true
}

// HasParseTestPublicId returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasParseTestPublicId() bool {
	return o != nil && o.ParseTestPublicId != nil
}

// SetParseTestPublicId gets a reference to the given string and assigns it to the ParseTestPublicId field.
func (o *SyntheticsGlobalVariableRequest) SetParseTestPublicId(v string) {
	o.ParseTestPublicId = &v
}

// GetTags returns the Tags field value.
func (o *SyntheticsGlobalVariableRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *SyntheticsGlobalVariableRequest) SetTags(v []string) {
	o.Tags = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsGlobalVariableRequest) GetValue() SyntheticsGlobalVariableValue {
	if o == nil || o.Value == nil {
		var ret SyntheticsGlobalVariableValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGlobalVariableRequest) GetValueOk() (*SyntheticsGlobalVariableValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsGlobalVariableRequest) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given SyntheticsGlobalVariableValue and assigns it to the Value field.
func (o *SyntheticsGlobalVariableRequest) SetValue(v SyntheticsGlobalVariableValue) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsGlobalVariableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["description"] = o.Description
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsFido != nil {
		toSerialize["is_fido"] = o.IsFido
	}
	if o.IsTotp != nil {
		toSerialize["is_totp"] = o.IsTotp
	}
	toSerialize["name"] = o.Name
	if o.ParseTestOptions != nil {
		toSerialize["parse_test_options"] = o.ParseTestOptions
	}
	if o.ParseTestPublicId != nil {
		toSerialize["parse_test_public_id"] = o.ParseTestPublicId
	}
	toSerialize["tags"] = o.Tags
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsGlobalVariableRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes        *SyntheticsGlobalVariableAttributes       `json:"attributes,omitempty"`
		Description       *string                                   `json:"description"`
		Id                *string                                   `json:"id,omitempty"`
		IsFido            *bool                                     `json:"is_fido,omitempty"`
		IsTotp            *bool                                     `json:"is_totp,omitempty"`
		Name              *string                                   `json:"name"`
		ParseTestOptions  *SyntheticsGlobalVariableParseTestOptions `json:"parse_test_options,omitempty"`
		ParseTestPublicId *string                                   `json:"parse_test_public_id,omitempty"`
		Tags              *[]string                                 `json:"tags"`
		Value             *SyntheticsGlobalVariableValue            `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "description", "id", "is_fido", "is_totp", "name", "parse_test_options", "parse_test_public_id", "tags", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Description = *all.Description
	o.Id = all.Id
	o.IsFido = all.IsFido
	o.IsTotp = all.IsTotp
	o.Name = *all.Name
	if all.ParseTestOptions != nil && all.ParseTestOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ParseTestOptions = all.ParseTestOptions
	o.ParseTestPublicId = all.ParseTestPublicId
	o.Tags = *all.Tags
	if all.Value != nil && all.Value.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
