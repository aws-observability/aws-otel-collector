// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateFormDataAttributes The attributes for creating a form.
type CreateFormDataAttributes struct {
	// Whether the form accepts anonymous submissions.
	Anonymous *bool `json:"anonymous,omitempty"`
	// A JSON Schema definition that describes the form's data fields.
	DataDefinition FormDataDefinition `json:"data_definition"`
	// The description of the form.
	Description *string `json:"description,omitempty"`
	// Whether the form is an IDP survey.
	IdpSurvey *bool `json:"idp_survey,omitempty"`
	// The name of the form.
	Name string `json:"name"`
	// Whether each user can only submit one response.
	SingleResponse *bool `json:"single_response,omitempty"`
	// UI configuration for rendering form fields, including widget overrides, field ordering, and themes.
	UiDefinition FormUiDefinition `json:"ui_definition"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateFormDataAttributes instantiates a new CreateFormDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateFormDataAttributes(dataDefinition FormDataDefinition, name string, uiDefinition FormUiDefinition) *CreateFormDataAttributes {
	this := CreateFormDataAttributes{}
	var anonymous bool = false
	this.Anonymous = &anonymous
	this.DataDefinition = dataDefinition
	var idpSurvey bool = false
	this.IdpSurvey = &idpSurvey
	this.Name = name
	var singleResponse bool = false
	this.SingleResponse = &singleResponse
	this.UiDefinition = uiDefinition
	return &this
}

// NewCreateFormDataAttributesWithDefaults instantiates a new CreateFormDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateFormDataAttributesWithDefaults() *CreateFormDataAttributes {
	this := CreateFormDataAttributes{}
	var anonymous bool = false
	this.Anonymous = &anonymous
	var idpSurvey bool = false
	this.IdpSurvey = &idpSurvey
	var singleResponse bool = false
	this.SingleResponse = &singleResponse
	return &this
}

// GetAnonymous returns the Anonymous field value if set, zero value otherwise.
func (o *CreateFormDataAttributes) GetAnonymous() bool {
	if o == nil || o.Anonymous == nil {
		var ret bool
		return ret
	}
	return *o.Anonymous
}

// GetAnonymousOk returns a tuple with the Anonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetAnonymousOk() (*bool, bool) {
	if o == nil || o.Anonymous == nil {
		return nil, false
	}
	return o.Anonymous, true
}

// HasAnonymous returns a boolean if a field has been set.
func (o *CreateFormDataAttributes) HasAnonymous() bool {
	return o != nil && o.Anonymous != nil
}

// SetAnonymous gets a reference to the given bool and assigns it to the Anonymous field.
func (o *CreateFormDataAttributes) SetAnonymous(v bool) {
	o.Anonymous = &v
}

// GetDataDefinition returns the DataDefinition field value.
func (o *CreateFormDataAttributes) GetDataDefinition() FormDataDefinition {
	if o == nil {
		var ret FormDataDefinition
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetDataDefinitionOk() (*FormDataDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *CreateFormDataAttributes) SetDataDefinition(v FormDataDefinition) {
	o.DataDefinition = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFormDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFormDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFormDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetIdpSurvey returns the IdpSurvey field value if set, zero value otherwise.
func (o *CreateFormDataAttributes) GetIdpSurvey() bool {
	if o == nil || o.IdpSurvey == nil {
		var ret bool
		return ret
	}
	return *o.IdpSurvey
}

// GetIdpSurveyOk returns a tuple with the IdpSurvey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetIdpSurveyOk() (*bool, bool) {
	if o == nil || o.IdpSurvey == nil {
		return nil, false
	}
	return o.IdpSurvey, true
}

// HasIdpSurvey returns a boolean if a field has been set.
func (o *CreateFormDataAttributes) HasIdpSurvey() bool {
	return o != nil && o.IdpSurvey != nil
}

// SetIdpSurvey gets a reference to the given bool and assigns it to the IdpSurvey field.
func (o *CreateFormDataAttributes) SetIdpSurvey(v bool) {
	o.IdpSurvey = &v
}

// GetName returns the Name field value.
func (o *CreateFormDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateFormDataAttributes) SetName(v string) {
	o.Name = v
}

// GetSingleResponse returns the SingleResponse field value if set, zero value otherwise.
func (o *CreateFormDataAttributes) GetSingleResponse() bool {
	if o == nil || o.SingleResponse == nil {
		var ret bool
		return ret
	}
	return *o.SingleResponse
}

// GetSingleResponseOk returns a tuple with the SingleResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetSingleResponseOk() (*bool, bool) {
	if o == nil || o.SingleResponse == nil {
		return nil, false
	}
	return o.SingleResponse, true
}

// HasSingleResponse returns a boolean if a field has been set.
func (o *CreateFormDataAttributes) HasSingleResponse() bool {
	return o != nil && o.SingleResponse != nil
}

// SetSingleResponse gets a reference to the given bool and assigns it to the SingleResponse field.
func (o *CreateFormDataAttributes) SetSingleResponse(v bool) {
	o.SingleResponse = &v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *CreateFormDataAttributes) GetUiDefinition() FormUiDefinition {
	if o == nil {
		var ret FormUiDefinition
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *CreateFormDataAttributes) GetUiDefinitionOk() (*FormUiDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *CreateFormDataAttributes) SetUiDefinition(v FormUiDefinition) {
	o.UiDefinition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateFormDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Anonymous != nil {
		toSerialize["anonymous"] = o.Anonymous
	}
	toSerialize["data_definition"] = o.DataDefinition
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IdpSurvey != nil {
		toSerialize["idp_survey"] = o.IdpSurvey
	}
	toSerialize["name"] = o.Name
	if o.SingleResponse != nil {
		toSerialize["single_response"] = o.SingleResponse
	}
	toSerialize["ui_definition"] = o.UiDefinition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateFormDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Anonymous      *bool               `json:"anonymous,omitempty"`
		DataDefinition *FormDataDefinition `json:"data_definition"`
		Description    *string             `json:"description,omitempty"`
		IdpSurvey      *bool               `json:"idp_survey,omitempty"`
		Name           *string             `json:"name"`
		SingleResponse *bool               `json:"single_response,omitempty"`
		UiDefinition   *FormUiDefinition   `json:"ui_definition"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anonymous", "data_definition", "description", "idp_survey", "name", "single_response", "ui_definition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Anonymous = all.Anonymous
	if all.DataDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataDefinition = *all.DataDefinition
	o.Description = all.Description
	o.IdpSurvey = all.IdpSurvey
	o.Name = *all.Name
	o.SingleResponse = all.SingleResponse
	if all.UiDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UiDefinition = *all.UiDefinition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
