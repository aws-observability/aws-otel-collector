// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseCreateAttributes Case creation attributes
type CaseCreateAttributes struct {
	// Case custom attributes
	CustomAttributes map[string]CustomAttributeValue `json:"custom_attributes,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
	// Status of the case. Must be one of the existing statuses for the case's type.
	StatusName *string `json:"status_name,omitempty"`
	// Title
	Title string `json:"title"`
	// Case type UUID
	TypeId string `json:"type_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseCreateAttributes instantiates a new CaseCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseCreateAttributes(title string, typeId string) *CaseCreateAttributes {
	this := CaseCreateAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	this.Title = title
	this.TypeId = typeId
	return &this
}

// NewCaseCreateAttributesWithDefaults instantiates a new CaseCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseCreateAttributesWithDefaults() *CaseCreateAttributes {
	this := CaseCreateAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = &priority
	return &this
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise.
func (o *CaseCreateAttributes) GetCustomAttributes() map[string]CustomAttributeValue {
	if o == nil || o.CustomAttributes == nil {
		var ret map[string]CustomAttributeValue
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseCreateAttributes) GetCustomAttributesOk() (*map[string]CustomAttributeValue, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return &o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *CaseCreateAttributes) HasCustomAttributes() bool {
	return o != nil && o.CustomAttributes != nil
}

// SetCustomAttributes gets a reference to the given map[string]CustomAttributeValue and assigns it to the CustomAttributes field.
func (o *CaseCreateAttributes) SetCustomAttributes(v map[string]CustomAttributeValue) {
	o.CustomAttributes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CaseCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CaseCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CaseCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CaseCreateAttributes) GetPriority() CasePriority {
	if o == nil || o.Priority == nil {
		var ret CasePriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseCreateAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CaseCreateAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given CasePriority and assigns it to the Priority field.
func (o *CaseCreateAttributes) SetPriority(v CasePriority) {
	o.Priority = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *CaseCreateAttributes) GetStatusName() string {
	if o == nil || o.StatusName == nil {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseCreateAttributes) GetStatusNameOk() (*string, bool) {
	if o == nil || o.StatusName == nil {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *CaseCreateAttributes) HasStatusName() bool {
	return o != nil && o.StatusName != nil
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *CaseCreateAttributes) SetStatusName(v string) {
	o.StatusName = &v
}

// GetTitle returns the Title field value.
func (o *CaseCreateAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CaseCreateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *CaseCreateAttributes) SetTitle(v string) {
	o.Title = v
}

// GetTypeId returns the TypeId field value.
func (o *CaseCreateAttributes) GetTypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *CaseCreateAttributes) GetTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value.
func (o *CaseCreateAttributes) SetTypeId(v string) {
	o.TypeId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomAttributes != nil {
		toSerialize["custom_attributes"] = o.CustomAttributes
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.StatusName != nil {
		toSerialize["status_name"] = o.StatusName
	}
	toSerialize["title"] = o.Title
	toSerialize["type_id"] = o.TypeId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomAttributes map[string]CustomAttributeValue `json:"custom_attributes,omitempty"`
		Description      *string                         `json:"description,omitempty"`
		Priority         *CasePriority                   `json:"priority,omitempty"`
		StatusName       *string                         `json:"status_name,omitempty"`
		Title            *string                         `json:"title"`
		TypeId           *string                         `json:"type_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.TypeId == nil {
		return fmt.Errorf("required field type_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_attributes", "description", "priority", "status_name", "title", "type_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomAttributes = all.CustomAttributes
	o.Description = all.Description
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	o.StatusName = all.StatusName
	o.Title = *all.Title
	o.TypeId = *all.TypeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
