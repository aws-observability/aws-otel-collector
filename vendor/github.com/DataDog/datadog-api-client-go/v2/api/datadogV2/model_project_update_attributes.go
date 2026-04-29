// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectUpdateAttributes Project update attributes.
type ProjectUpdateAttributes struct {
	// Project columns configuration.
	ColumnsConfig *ProjectColumnsConfig `json:"columns_config,omitempty"`
	// List of enabled custom case type IDs.
	EnabledCustomCaseTypes []string `json:"enabled_custom_case_types,omitempty"`
	// Project name.
	Name *string `json:"name,omitempty"`
	// Project settings.
	Settings *ProjectSettings `json:"settings,omitempty"`
	// Team UUID to associate with the project.
	TeamUuid *string `json:"team_uuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectUpdateAttributes instantiates a new ProjectUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectUpdateAttributes() *ProjectUpdateAttributes {
	this := ProjectUpdateAttributes{}
	return &this
}

// NewProjectUpdateAttributesWithDefaults instantiates a new ProjectUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectUpdateAttributesWithDefaults() *ProjectUpdateAttributes {
	this := ProjectUpdateAttributes{}
	return &this
}

// GetColumnsConfig returns the ColumnsConfig field value if set, zero value otherwise.
func (o *ProjectUpdateAttributes) GetColumnsConfig() ProjectColumnsConfig {
	if o == nil || o.ColumnsConfig == nil {
		var ret ProjectColumnsConfig
		return ret
	}
	return *o.ColumnsConfig
}

// GetColumnsConfigOk returns a tuple with the ColumnsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateAttributes) GetColumnsConfigOk() (*ProjectColumnsConfig, bool) {
	if o == nil || o.ColumnsConfig == nil {
		return nil, false
	}
	return o.ColumnsConfig, true
}

// HasColumnsConfig returns a boolean if a field has been set.
func (o *ProjectUpdateAttributes) HasColumnsConfig() bool {
	return o != nil && o.ColumnsConfig != nil
}

// SetColumnsConfig gets a reference to the given ProjectColumnsConfig and assigns it to the ColumnsConfig field.
func (o *ProjectUpdateAttributes) SetColumnsConfig(v ProjectColumnsConfig) {
	o.ColumnsConfig = &v
}

// GetEnabledCustomCaseTypes returns the EnabledCustomCaseTypes field value if set, zero value otherwise.
func (o *ProjectUpdateAttributes) GetEnabledCustomCaseTypes() []string {
	if o == nil || o.EnabledCustomCaseTypes == nil {
		var ret []string
		return ret
	}
	return o.EnabledCustomCaseTypes
}

// GetEnabledCustomCaseTypesOk returns a tuple with the EnabledCustomCaseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateAttributes) GetEnabledCustomCaseTypesOk() (*[]string, bool) {
	if o == nil || o.EnabledCustomCaseTypes == nil {
		return nil, false
	}
	return &o.EnabledCustomCaseTypes, true
}

// HasEnabledCustomCaseTypes returns a boolean if a field has been set.
func (o *ProjectUpdateAttributes) HasEnabledCustomCaseTypes() bool {
	return o != nil && o.EnabledCustomCaseTypes != nil
}

// SetEnabledCustomCaseTypes gets a reference to the given []string and assigns it to the EnabledCustomCaseTypes field.
func (o *ProjectUpdateAttributes) SetEnabledCustomCaseTypes(v []string) {
	o.EnabledCustomCaseTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ProjectUpdateAttributes) GetSettings() ProjectSettings {
	if o == nil || o.Settings == nil {
		var ret ProjectSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateAttributes) GetSettingsOk() (*ProjectSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProjectUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given ProjectSettings and assigns it to the Settings field.
func (o *ProjectUpdateAttributes) SetSettings(v ProjectSettings) {
	o.Settings = &v
}

// GetTeamUuid returns the TeamUuid field value if set, zero value otherwise.
func (o *ProjectUpdateAttributes) GetTeamUuid() string {
	if o == nil || o.TeamUuid == nil {
		var ret string
		return ret
	}
	return *o.TeamUuid
}

// GetTeamUuidOk returns a tuple with the TeamUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateAttributes) GetTeamUuidOk() (*string, bool) {
	if o == nil || o.TeamUuid == nil {
		return nil, false
	}
	return o.TeamUuid, true
}

// HasTeamUuid returns a boolean if a field has been set.
func (o *ProjectUpdateAttributes) HasTeamUuid() bool {
	return o != nil && o.TeamUuid != nil
}

// SetTeamUuid gets a reference to the given string and assigns it to the TeamUuid field.
func (o *ProjectUpdateAttributes) SetTeamUuid(v string) {
	o.TeamUuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.TeamUuid != nil {
		toSerialize["team_uuid"] = o.TeamUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnsConfig          *ProjectColumnsConfig `json:"columns_config,omitempty"`
		EnabledCustomCaseTypes []string              `json:"enabled_custom_case_types,omitempty"`
		Name                   *string               `json:"name,omitempty"`
		Settings               *ProjectSettings      `json:"settings,omitempty"`
		TeamUuid               *string               `json:"team_uuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns_config", "enabled_custom_case_types", "name", "settings", "team_uuid"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ColumnsConfig != nil && all.ColumnsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ColumnsConfig = all.ColumnsConfig
	o.EnabledCustomCaseTypes = all.EnabledCustomCaseTypes
	o.Name = all.Name
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings
	o.TeamUuid = all.TeamUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
