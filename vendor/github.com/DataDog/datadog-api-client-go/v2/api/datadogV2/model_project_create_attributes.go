// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectCreateAttributes Project creation attributes.
type ProjectCreateAttributes struct {
	// List of enabled custom case type IDs.
	EnabledCustomCaseTypes []string `json:"enabled_custom_case_types,omitempty"`
	// Project's key. Cannot be "CASE".
	Key string `json:"key"`
	// Project name.
	Name string `json:"name"`
	// Team UUID to associate with the project.
	TeamUuid *string `json:"team_uuid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectCreateAttributes instantiates a new ProjectCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectCreateAttributes(key string, name string) *ProjectCreateAttributes {
	this := ProjectCreateAttributes{}
	this.Key = key
	this.Name = name
	return &this
}

// NewProjectCreateAttributesWithDefaults instantiates a new ProjectCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectCreateAttributesWithDefaults() *ProjectCreateAttributes {
	this := ProjectCreateAttributes{}
	return &this
}

// GetEnabledCustomCaseTypes returns the EnabledCustomCaseTypes field value if set, zero value otherwise.
func (o *ProjectCreateAttributes) GetEnabledCustomCaseTypes() []string {
	if o == nil || o.EnabledCustomCaseTypes == nil {
		var ret []string
		return ret
	}
	return o.EnabledCustomCaseTypes
}

// GetEnabledCustomCaseTypesOk returns a tuple with the EnabledCustomCaseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateAttributes) GetEnabledCustomCaseTypesOk() (*[]string, bool) {
	if o == nil || o.EnabledCustomCaseTypes == nil {
		return nil, false
	}
	return &o.EnabledCustomCaseTypes, true
}

// HasEnabledCustomCaseTypes returns a boolean if a field has been set.
func (o *ProjectCreateAttributes) HasEnabledCustomCaseTypes() bool {
	return o != nil && o.EnabledCustomCaseTypes != nil
}

// SetEnabledCustomCaseTypes gets a reference to the given []string and assigns it to the EnabledCustomCaseTypes field.
func (o *ProjectCreateAttributes) SetEnabledCustomCaseTypes(v []string) {
	o.EnabledCustomCaseTypes = v
}

// GetKey returns the Key field value.
func (o *ProjectCreateAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProjectCreateAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ProjectCreateAttributes) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value.
func (o *ProjectCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProjectCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetTeamUuid returns the TeamUuid field value if set, zero value otherwise.
func (o *ProjectCreateAttributes) GetTeamUuid() string {
	if o == nil || o.TeamUuid == nil {
		var ret string
		return ret
	}
	return *o.TeamUuid
}

// GetTeamUuidOk returns a tuple with the TeamUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateAttributes) GetTeamUuidOk() (*string, bool) {
	if o == nil || o.TeamUuid == nil {
		return nil, false
	}
	return o.TeamUuid, true
}

// HasTeamUuid returns a boolean if a field has been set.
func (o *ProjectCreateAttributes) HasTeamUuid() bool {
	return o != nil && o.TeamUuid != nil
}

// SetTeamUuid gets a reference to the given string and assigns it to the TeamUuid field.
func (o *ProjectCreateAttributes) SetTeamUuid(v string) {
	o.TeamUuid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnabledCustomCaseTypes != nil {
		toSerialize["enabled_custom_case_types"] = o.EnabledCustomCaseTypes
	}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if o.TeamUuid != nil {
		toSerialize["team_uuid"] = o.TeamUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnabledCustomCaseTypes []string `json:"enabled_custom_case_types,omitempty"`
		Key                    *string  `json:"key"`
		Name                   *string  `json:"name"`
		TeamUuid               *string  `json:"team_uuid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled_custom_case_types", "key", "name", "team_uuid"})
	} else {
		return err
	}
	o.EnabledCustomCaseTypes = all.EnabledCustomCaseTypes
	o.Key = *all.Key
	o.Name = *all.Name
	o.TeamUuid = all.TeamUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
