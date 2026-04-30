// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedCreateRequestAttributes Attributes for creating a secure embed shared dashboard.
type SecureEmbedCreateRequestAttributes struct {
	// Default time range configuration for the secure embed.
	GlobalTime SecureEmbedGlobalTime `json:"global_time"`
	// Whether viewers can change the time range.
	GlobalTimeSelectable bool `json:"global_time_selectable"`
	// Template variables viewers can modify.
	SelectableTemplateVars []SecureEmbedSelectableTemplateVariable `json:"selectable_template_vars"`
	// The status of the secure embed share. Active means the shared dashboard is available. Paused means it is not.
	Status SecureEmbedStatus `json:"status"`
	// Display title for the shared dashboard.
	Title string `json:"title"`
	// Display settings for the secure embed shared dashboard.
	ViewingPreferences SecureEmbedViewingPreferences `json:"viewing_preferences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecureEmbedCreateRequestAttributes instantiates a new SecureEmbedCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecureEmbedCreateRequestAttributes(globalTime SecureEmbedGlobalTime, globalTimeSelectable bool, selectableTemplateVars []SecureEmbedSelectableTemplateVariable, status SecureEmbedStatus, title string, viewingPreferences SecureEmbedViewingPreferences) *SecureEmbedCreateRequestAttributes {
	this := SecureEmbedCreateRequestAttributes{}
	this.GlobalTime = globalTime
	this.GlobalTimeSelectable = globalTimeSelectable
	this.SelectableTemplateVars = selectableTemplateVars
	this.Status = status
	this.Title = title
	this.ViewingPreferences = viewingPreferences
	return &this
}

// NewSecureEmbedCreateRequestAttributesWithDefaults instantiates a new SecureEmbedCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecureEmbedCreateRequestAttributesWithDefaults() *SecureEmbedCreateRequestAttributes {
	this := SecureEmbedCreateRequestAttributes{}
	return &this
}

// GetGlobalTime returns the GlobalTime field value.
func (o *SecureEmbedCreateRequestAttributes) GetGlobalTime() SecureEmbedGlobalTime {
	if o == nil {
		var ret SecureEmbedGlobalTime
		return ret
	}
	return o.GlobalTime
}

// GetGlobalTimeOk returns a tuple with the GlobalTime field value
// and a boolean to check if the value has been set.
func (o *SecureEmbedCreateRequestAttributes) GetGlobalTimeOk() (*SecureEmbedGlobalTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalTime, true
}

// SetGlobalTime sets field value.
func (o *SecureEmbedCreateRequestAttributes) SetGlobalTime(v SecureEmbedGlobalTime) {
	o.GlobalTime = v
}

// GetGlobalTimeSelectable returns the GlobalTimeSelectable field value.
func (o *SecureEmbedCreateRequestAttributes) GetGlobalTimeSelectable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.GlobalTimeSelectable
}

// GetGlobalTimeSelectableOk returns a tuple with the GlobalTimeSelectable field value
// and a boolean to check if the value has been set.
func (o *SecureEmbedCreateRequestAttributes) GetGlobalTimeSelectableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalTimeSelectable, true
}

// SetGlobalTimeSelectable sets field value.
func (o *SecureEmbedCreateRequestAttributes) SetGlobalTimeSelectable(v bool) {
	o.GlobalTimeSelectable = v
}

// GetSelectableTemplateVars returns the SelectableTemplateVars field value.
func (o *SecureEmbedCreateRequestAttributes) GetSelectableTemplateVars() []SecureEmbedSelectableTemplateVariable {
	if o == nil {
		var ret []SecureEmbedSelectableTemplateVariable
		return ret
	}
	return o.SelectableTemplateVars
}

// GetSelectableTemplateVarsOk returns a tuple with the SelectableTemplateVars field value
// and a boolean to check if the value has been set.
func (o *SecureEmbedCreateRequestAttributes) GetSelectableTemplateVarsOk() (*[]SecureEmbedSelectableTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectableTemplateVars, true
}

// SetSelectableTemplateVars sets field value.
func (o *SecureEmbedCreateRequestAttributes) SetSelectableTemplateVars(v []SecureEmbedSelectableTemplateVariable) {
	o.SelectableTemplateVars = v
}

// GetStatus returns the Status field value.
func (o *SecureEmbedCreateRequestAttributes) GetStatus() SecureEmbedStatus {
	if o == nil {
		var ret SecureEmbedStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecureEmbedCreateRequestAttributes) GetStatusOk() (*SecureEmbedStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecureEmbedCreateRequestAttributes) SetStatus(v SecureEmbedStatus) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *SecureEmbedCreateRequestAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SecureEmbedCreateRequestAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *SecureEmbedCreateRequestAttributes) SetTitle(v string) {
	o.Title = v
}

// GetViewingPreferences returns the ViewingPreferences field value.
func (o *SecureEmbedCreateRequestAttributes) GetViewingPreferences() SecureEmbedViewingPreferences {
	if o == nil {
		var ret SecureEmbedViewingPreferences
		return ret
	}
	return o.ViewingPreferences
}

// GetViewingPreferencesOk returns a tuple with the ViewingPreferences field value
// and a boolean to check if the value has been set.
func (o *SecureEmbedCreateRequestAttributes) GetViewingPreferencesOk() (*SecureEmbedViewingPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewingPreferences, true
}

// SetViewingPreferences sets field value.
func (o *SecureEmbedCreateRequestAttributes) SetViewingPreferences(v SecureEmbedViewingPreferences) {
	o.ViewingPreferences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecureEmbedCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["global_time"] = o.GlobalTime
	toSerialize["global_time_selectable"] = o.GlobalTimeSelectable
	toSerialize["selectable_template_vars"] = o.SelectableTemplateVars
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["viewing_preferences"] = o.ViewingPreferences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecureEmbedCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GlobalTime             *SecureEmbedGlobalTime                   `json:"global_time"`
		GlobalTimeSelectable   *bool                                    `json:"global_time_selectable"`
		SelectableTemplateVars *[]SecureEmbedSelectableTemplateVariable `json:"selectable_template_vars"`
		Status                 *SecureEmbedStatus                       `json:"status"`
		Title                  *string                                  `json:"title"`
		ViewingPreferences     *SecureEmbedViewingPreferences           `json:"viewing_preferences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GlobalTime == nil {
		return fmt.Errorf("required field global_time missing")
	}
	if all.GlobalTimeSelectable == nil {
		return fmt.Errorf("required field global_time_selectable missing")
	}
	if all.SelectableTemplateVars == nil {
		return fmt.Errorf("required field selectable_template_vars missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.ViewingPreferences == nil {
		return fmt.Errorf("required field viewing_preferences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"global_time", "global_time_selectable", "selectable_template_vars", "status", "title", "viewing_preferences"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GlobalTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GlobalTime = *all.GlobalTime
	o.GlobalTimeSelectable = *all.GlobalTimeSelectable
	o.SelectableTemplateVars = *all.SelectableTemplateVars
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Title = *all.Title
	if all.ViewingPreferences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ViewingPreferences = *all.ViewingPreferences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
