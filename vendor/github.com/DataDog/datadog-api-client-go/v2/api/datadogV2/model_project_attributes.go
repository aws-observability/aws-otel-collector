// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectAttributes Project attributes.
type ProjectAttributes struct {
	// Project columns configuration.
	ColumnsConfig *ProjectColumnsConfig `json:"columns_config,omitempty"`
	// List of enabled custom case type IDs.
	EnabledCustomCaseTypes []string `json:"enabled_custom_case_types,omitempty"`
	// The project's key.
	Key *string `json:"key,omitempty"`
	// Project's name.
	Name *string `json:"name,omitempty"`
	// Whether the project is restricted.
	Restricted *bool `json:"restricted,omitempty"`
	// Project settings.
	Settings *ProjectSettings `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectAttributes instantiates a new ProjectAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectAttributes() *ProjectAttributes {
	this := ProjectAttributes{}
	return &this
}

// NewProjectAttributesWithDefaults instantiates a new ProjectAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectAttributesWithDefaults() *ProjectAttributes {
	this := ProjectAttributes{}
	return &this
}

// GetColumnsConfig returns the ColumnsConfig field value if set, zero value otherwise.
func (o *ProjectAttributes) GetColumnsConfig() ProjectColumnsConfig {
	if o == nil || o.ColumnsConfig == nil {
		var ret ProjectColumnsConfig
		return ret
	}
	return *o.ColumnsConfig
}

// GetColumnsConfigOk returns a tuple with the ColumnsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAttributes) GetColumnsConfigOk() (*ProjectColumnsConfig, bool) {
	if o == nil || o.ColumnsConfig == nil {
		return nil, false
	}
	return o.ColumnsConfig, true
}

// HasColumnsConfig returns a boolean if a field has been set.
func (o *ProjectAttributes) HasColumnsConfig() bool {
	return o != nil && o.ColumnsConfig != nil
}

// SetColumnsConfig gets a reference to the given ProjectColumnsConfig and assigns it to the ColumnsConfig field.
func (o *ProjectAttributes) SetColumnsConfig(v ProjectColumnsConfig) {
	o.ColumnsConfig = &v
}

// GetEnabledCustomCaseTypes returns the EnabledCustomCaseTypes field value if set, zero value otherwise.
func (o *ProjectAttributes) GetEnabledCustomCaseTypes() []string {
	if o == nil || o.EnabledCustomCaseTypes == nil {
		var ret []string
		return ret
	}
	return o.EnabledCustomCaseTypes
}

// GetEnabledCustomCaseTypesOk returns a tuple with the EnabledCustomCaseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAttributes) GetEnabledCustomCaseTypesOk() (*[]string, bool) {
	if o == nil || o.EnabledCustomCaseTypes == nil {
		return nil, false
	}
	return &o.EnabledCustomCaseTypes, true
}

// HasEnabledCustomCaseTypes returns a boolean if a field has been set.
func (o *ProjectAttributes) HasEnabledCustomCaseTypes() bool {
	return o != nil && o.EnabledCustomCaseTypes != nil
}

// SetEnabledCustomCaseTypes gets a reference to the given []string and assigns it to the EnabledCustomCaseTypes field.
func (o *ProjectAttributes) SetEnabledCustomCaseTypes(v []string) {
	o.EnabledCustomCaseTypes = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProjectAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ProjectAttributes) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectAttributes) SetName(v string) {
	o.Name = &v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *ProjectAttributes) GetRestricted() bool {
	if o == nil || o.Restricted == nil {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAttributes) GetRestrictedOk() (*bool, bool) {
	if o == nil || o.Restricted == nil {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *ProjectAttributes) HasRestricted() bool {
	return o != nil && o.Restricted != nil
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *ProjectAttributes) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ProjectAttributes) GetSettings() ProjectSettings {
	if o == nil || o.Settings == nil {
		var ret ProjectSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAttributes) GetSettingsOk() (*ProjectSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProjectAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given ProjectSettings and assigns it to the Settings field.
func (o *ProjectAttributes) SetSettings(v ProjectSettings) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnsConfig != nil {
		toSerialize["columns_config"] = o.ColumnsConfig
	}
	if o.EnabledCustomCaseTypes != nil {
		toSerialize["enabled_custom_case_types"] = o.EnabledCustomCaseTypes
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Restricted != nil {
		toSerialize["restricted"] = o.Restricted
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnsConfig          *ProjectColumnsConfig `json:"columns_config,omitempty"`
		EnabledCustomCaseTypes []string              `json:"enabled_custom_case_types,omitempty"`
		Key                    *string               `json:"key,omitempty"`
		Name                   *string               `json:"name,omitempty"`
		Restricted             *bool                 `json:"restricted,omitempty"`
		Settings               *ProjectSettings      `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns_config", "enabled_custom_case_types", "key", "name", "restricted", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ColumnsConfig != nil && all.ColumnsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ColumnsConfig = all.ColumnsConfig
	o.EnabledCustomCaseTypes = all.EnabledCustomCaseTypes
	o.Key = all.Key
	o.Name = all.Name
	o.Restricted = all.Restricted
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
