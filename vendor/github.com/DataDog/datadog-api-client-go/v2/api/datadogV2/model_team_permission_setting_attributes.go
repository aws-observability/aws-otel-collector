// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamPermissionSettingAttributes Team permission setting attributes
type TeamPermissionSettingAttributes struct {
	// The identifier for the action
	Action *TeamPermissionSettingSerializerAction `json:"action,omitempty"`
	// Whether or not the permission setting is editable by the current user
	Editable *bool `json:"editable,omitempty"`
	// Possible values for action
	Options []TeamPermissionSettingValue `json:"options,omitempty"`
	// The team permission name
	Title *string `json:"title,omitempty"`
	// What type of user is allowed to perform the specified action
	Value *TeamPermissionSettingValue `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamPermissionSettingAttributes instantiates a new TeamPermissionSettingAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamPermissionSettingAttributes() *TeamPermissionSettingAttributes {
	this := TeamPermissionSettingAttributes{}
	return &this
}

// NewTeamPermissionSettingAttributesWithDefaults instantiates a new TeamPermissionSettingAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamPermissionSettingAttributesWithDefaults() *TeamPermissionSettingAttributes {
	this := TeamPermissionSettingAttributes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TeamPermissionSettingAttributes) GetAction() TeamPermissionSettingSerializerAction {
	if o == nil || o.Action == nil {
		var ret TeamPermissionSettingSerializerAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingAttributes) GetActionOk() (*TeamPermissionSettingSerializerAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TeamPermissionSettingAttributes) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given TeamPermissionSettingSerializerAction and assigns it to the Action field.
func (o *TeamPermissionSettingAttributes) SetAction(v TeamPermissionSettingSerializerAction) {
	o.Action = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *TeamPermissionSettingAttributes) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingAttributes) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *TeamPermissionSettingAttributes) HasEditable() bool {
	return o != nil && o.Editable != nil
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *TeamPermissionSettingAttributes) SetEditable(v bool) {
	o.Editable = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TeamPermissionSettingAttributes) GetOptions() []TeamPermissionSettingValue {
	if o == nil || o.Options == nil {
		var ret []TeamPermissionSettingValue
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingAttributes) GetOptionsOk() (*[]TeamPermissionSettingValue, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TeamPermissionSettingAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given []TeamPermissionSettingValue and assigns it to the Options field.
func (o *TeamPermissionSettingAttributes) SetOptions(v []TeamPermissionSettingValue) {
	o.Options = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TeamPermissionSettingAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TeamPermissionSettingAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TeamPermissionSettingAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TeamPermissionSettingAttributes) GetValue() TeamPermissionSettingValue {
	if o == nil || o.Value == nil {
		var ret TeamPermissionSettingValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPermissionSettingAttributes) GetValueOk() (*TeamPermissionSettingValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TeamPermissionSettingAttributes) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given TeamPermissionSettingValue and assigns it to the Value field.
func (o *TeamPermissionSettingAttributes) SetValue(v TeamPermissionSettingValue) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamPermissionSettingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamPermissionSettingAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action   *TeamPermissionSettingSerializerAction `json:"action,omitempty"`
		Editable *bool                                  `json:"editable,omitempty"`
		Options  []TeamPermissionSettingValue           `json:"options,omitempty"`
		Title    *string                                `json:"title,omitempty"`
		Value    *TeamPermissionSettingValue            `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "editable", "options", "title", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action != nil && !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = all.Action
	}
	o.Editable = all.Editable
	o.Options = all.Options
	o.Title = all.Title
	if all.Value != nil && !all.Value.IsValid() {
		hasInvalidField = true
	} else {
		o.Value = all.Value
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
