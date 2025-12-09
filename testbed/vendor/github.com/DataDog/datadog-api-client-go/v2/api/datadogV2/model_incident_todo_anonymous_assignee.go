// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTodoAnonymousAssignee Anonymous assignee entity.
type IncidentTodoAnonymousAssignee struct {
	// URL for assignee's icon.
	Icon string `json:"icon"`
	// Anonymous assignee's ID.
	Id string `json:"id"`
	// Assignee's name.
	Name string `json:"name"`
	// The source of the anonymous assignee.
	Source IncidentTodoAnonymousAssigneeSource `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTodoAnonymousAssignee instantiates a new IncidentTodoAnonymousAssignee object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTodoAnonymousAssignee(icon string, id string, name string, source IncidentTodoAnonymousAssigneeSource) *IncidentTodoAnonymousAssignee {
	this := IncidentTodoAnonymousAssignee{}
	this.Icon = icon
	this.Id = id
	this.Name = name
	this.Source = source
	return &this
}

// NewIncidentTodoAnonymousAssigneeWithDefaults instantiates a new IncidentTodoAnonymousAssignee object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTodoAnonymousAssigneeWithDefaults() *IncidentTodoAnonymousAssignee {
	this := IncidentTodoAnonymousAssignee{}
	var source IncidentTodoAnonymousAssigneeSource = INCIDENTTODOANONYMOUSASSIGNEESOURCE_SLACK
	this.Source = source
	return &this
}

// GetIcon returns the Icon field value.
func (o *IncidentTodoAnonymousAssignee) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoAnonymousAssignee) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value.
func (o *IncidentTodoAnonymousAssignee) SetIcon(v string) {
	o.Icon = v
}

// GetId returns the Id field value.
func (o *IncidentTodoAnonymousAssignee) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoAnonymousAssignee) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentTodoAnonymousAssignee) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *IncidentTodoAnonymousAssignee) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoAnonymousAssignee) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentTodoAnonymousAssignee) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value.
func (o *IncidentTodoAnonymousAssignee) GetSource() IncidentTodoAnonymousAssigneeSource {
	if o == nil {
		var ret IncidentTodoAnonymousAssigneeSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoAnonymousAssignee) GetSourceOk() (*IncidentTodoAnonymousAssigneeSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *IncidentTodoAnonymousAssignee) SetSource(v IncidentTodoAnonymousAssigneeSource) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTodoAnonymousAssignee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["icon"] = o.Icon
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTodoAnonymousAssignee) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Icon   *string                              `json:"icon"`
		Id     *string                              `json:"id"`
		Name   *string                              `json:"name"`
		Source *IncidentTodoAnonymousAssigneeSource `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Icon == nil {
		return fmt.Errorf("required field icon missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"icon", "id", "name", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Icon = *all.Icon
	o.Id = *all.Id
	o.Name = *all.Name
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
