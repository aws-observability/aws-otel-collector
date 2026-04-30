// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedUpdateRequestAttributes Attributes for updating a secure embed shared dashboard. All fields are optional.
type SecureEmbedUpdateRequestAttributes struct {
	// Default time range configuration for the secure embed.
	GlobalTime *SecureEmbedGlobalTime `json:"global_time,omitempty"`
	// Updated time selectability.
	GlobalTimeSelectable *bool `json:"global_time_selectable,omitempty"`
	// Updated template variables.
	SelectableTemplateVars []SecureEmbedSelectableTemplateVariable `json:"selectable_template_vars,omitempty"`
	// The status of the secure embed share. Active means the shared dashboard is available. Paused means it is not.
	Status *SecureEmbedStatus `json:"status,omitempty"`
	// Updated title.
	Title *string `json:"title,omitempty"`
	// Display settings for the secure embed shared dashboard.
	ViewingPreferences *SecureEmbedViewingPreferences `json:"viewing_preferences,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecureEmbedUpdateRequestAttributes instantiates a new SecureEmbedUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecureEmbedUpdateRequestAttributes() *SecureEmbedUpdateRequestAttributes {
	this := SecureEmbedUpdateRequestAttributes{}
	return &this
}

// NewSecureEmbedUpdateRequestAttributesWithDefaults instantiates a new SecureEmbedUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecureEmbedUpdateRequestAttributesWithDefaults() *SecureEmbedUpdateRequestAttributes {
	this := SecureEmbedUpdateRequestAttributes{}
	return &this
}

// GetGlobalTime returns the GlobalTime field value if set, zero value otherwise.
func (o *SecureEmbedUpdateRequestAttributes) GetGlobalTime() SecureEmbedGlobalTime {
	if o == nil || o.GlobalTime == nil {
		var ret SecureEmbedGlobalTime
		return ret
	}
	return *o.GlobalTime
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateRequestAttributes) GetGlobalTimeOk() (*SecureEmbedGlobalTime, bool) {
	if o == nil || o.GlobalTime == nil {
		return nil, false
	}
	return o.GlobalTime, true
}

// HasGlobalTime returns a boolean if a field has been set.
func (o *SecureEmbedUpdateRequestAttributes) HasGlobalTime() bool {
	return o != nil && o.GlobalTime != nil
}

// SetGlobalTime gets a reference to the given SecureEmbedGlobalTime and assigns it to the GlobalTime field.
func (o *SecureEmbedUpdateRequestAttributes) SetGlobalTime(v SecureEmbedGlobalTime) {
	o.GlobalTime = &v
}

// GetGlobalTimeSelectable returns the GlobalTimeSelectable field value if set, zero value otherwise.
func (o *SecureEmbedUpdateRequestAttributes) GetGlobalTimeSelectable() bool {
	if o == nil || o.GlobalTimeSelectable == nil {
		var ret bool
		return ret
	}
	return *o.GlobalTimeSelectable
}

// GetGlobalTimeSelectableOk returns a tuple with the GlobalTimeSelectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateRequestAttributes) GetGlobalTimeSelectableOk() (*bool, bool) {
	if o == nil || o.GlobalTimeSelectable == nil {
		return nil, false
	}
	return o.GlobalTimeSelectable, true
}

// HasGlobalTimeSelectable returns a boolean if a field has been set.
func (o *SecureEmbedUpdateRequestAttributes) HasGlobalTimeSelectable() bool {
	return o != nil && o.GlobalTimeSelectable != nil
}

// SetGlobalTimeSelectable gets a reference to the given bool and assigns it to the GlobalTimeSelectable field.
func (o *SecureEmbedUpdateRequestAttributes) SetGlobalTimeSelectable(v bool) {
	o.GlobalTimeSelectable = &v
}

// GetSelectableTemplateVars returns the SelectableTemplateVars field value if set, zero value otherwise.
func (o *SecureEmbedUpdateRequestAttributes) GetSelectableTemplateVars() []SecureEmbedSelectableTemplateVariable {
	if o == nil || o.SelectableTemplateVars == nil {
		var ret []SecureEmbedSelectableTemplateVariable
		return ret
	}
	return o.SelectableTemplateVars
}

// GetSelectableTemplateVarsOk returns a tuple with the SelectableTemplateVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateRequestAttributes) GetSelectableTemplateVarsOk() (*[]SecureEmbedSelectableTemplateVariable, bool) {
	if o == nil || o.SelectableTemplateVars == nil {
		return nil, false
	}
	return &o.SelectableTemplateVars, true
}

// HasSelectableTemplateVars returns a boolean if a field has been set.
func (o *SecureEmbedUpdateRequestAttributes) HasSelectableTemplateVars() bool {
	return o != nil && o.SelectableTemplateVars != nil
}

// SetSelectableTemplateVars gets a reference to the given []SecureEmbedSelectableTemplateVariable and assigns it to the SelectableTemplateVars field.
func (o *SecureEmbedUpdateRequestAttributes) SetSelectableTemplateVars(v []SecureEmbedSelectableTemplateVariable) {
	o.SelectableTemplateVars = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecureEmbedUpdateRequestAttributes) GetStatus() SecureEmbedStatus {
	if o == nil || o.Status == nil {
		var ret SecureEmbedStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateRequestAttributes) GetStatusOk() (*SecureEmbedStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecureEmbedUpdateRequestAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SecureEmbedStatus and assigns it to the Status field.
func (o *SecureEmbedUpdateRequestAttributes) SetStatus(v SecureEmbedStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SecureEmbedUpdateRequestAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateRequestAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SecureEmbedUpdateRequestAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SecureEmbedUpdateRequestAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetViewingPreferences returns the ViewingPreferences field value if set, zero value otherwise.
func (o *SecureEmbedUpdateRequestAttributes) GetViewingPreferences() SecureEmbedViewingPreferences {
	if o == nil || o.ViewingPreferences == nil {
		var ret SecureEmbedViewingPreferences
		return ret
	}
	return *o.ViewingPreferences
}

// GetViewingPreferencesOk returns a tuple with the ViewingPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedUpdateRequestAttributes) GetViewingPreferencesOk() (*SecureEmbedViewingPreferences, bool) {
	if o == nil || o.ViewingPreferences == nil {
		return nil, false
	}
	return o.ViewingPreferences, true
}

// HasViewingPreferences returns a boolean if a field has been set.
func (o *SecureEmbedUpdateRequestAttributes) HasViewingPreferences() bool {
	return o != nil && o.ViewingPreferences != nil
}

// SetViewingPreferences gets a reference to the given SecureEmbedViewingPreferences and assigns it to the ViewingPreferences field.
func (o *SecureEmbedUpdateRequestAttributes) SetViewingPreferences(v SecureEmbedViewingPreferences) {
	o.ViewingPreferences = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecureEmbedUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GlobalTime != nil {
		toSerialize["global_time"] = o.GlobalTime
	}
	if o.GlobalTimeSelectable != nil {
		toSerialize["global_time_selectable"] = o.GlobalTimeSelectable
	}
	if o.SelectableTemplateVars != nil {
		toSerialize["selectable_template_vars"] = o.SelectableTemplateVars
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ViewingPreferences != nil {
		toSerialize["viewing_preferences"] = o.ViewingPreferences
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecureEmbedUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GlobalTime             *SecureEmbedGlobalTime                  `json:"global_time,omitempty"`
		GlobalTimeSelectable   *bool                                   `json:"global_time_selectable,omitempty"`
		SelectableTemplateVars []SecureEmbedSelectableTemplateVariable `json:"selectable_template_vars,omitempty"`
		Status                 *SecureEmbedStatus                      `json:"status,omitempty"`
		Title                  *string                                 `json:"title,omitempty"`
		ViewingPreferences     *SecureEmbedViewingPreferences          `json:"viewing_preferences,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"global_time", "global_time_selectable", "selectable_template_vars", "status", "title", "viewing_preferences"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GlobalTime != nil && all.GlobalTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GlobalTime = all.GlobalTime
	o.GlobalTimeSelectable = all.GlobalTimeSelectable
	o.SelectableTemplateVars = all.SelectableTemplateVars
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title
	if all.ViewingPreferences != nil && all.ViewingPreferences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ViewingPreferences = all.ViewingPreferences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
