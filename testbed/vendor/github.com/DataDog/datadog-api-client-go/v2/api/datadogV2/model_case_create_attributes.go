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
	// Description
	Description *string `json:"description,omitempty"`
	// Case priority
	Priority *CasePriority `json:"priority,omitempty"`
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
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
		Description *string       `json:"description,omitempty"`
		Priority    *CasePriority `json:"priority,omitempty"`
		Title       *string       `json:"title"`
		TypeId      *string       `json:"type_id"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "priority", "title", "type_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
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
