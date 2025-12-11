// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesetRespDataAttributes The definition of `RulesetRespDataAttributes` object.
type RulesetRespDataAttributes struct {
	// The definition of `RulesetRespDataAttributesCreated` object.
	Created RulesetRespDataAttributesCreated `json:"created"`
	// The `attributes` `enabled`.
	Enabled bool `json:"enabled"`
	// The `attributes` `last_modified_user_uuid`.
	LastModifiedUserUuid string `json:"last_modified_user_uuid"`
	// The definition of `RulesetRespDataAttributesModified` object.
	Modified RulesetRespDataAttributesModified `json:"modified"`
	// The `attributes` `name`.
	Name string `json:"name"`
	// The `attributes` `position`.
	Position int32 `json:"position"`
	// The `attributes` `processing_status`.
	ProcessingStatus *string `json:"processing_status,omitempty"`
	// The `attributes` `rules`.
	Rules []RulesetRespDataAttributesRulesItems `json:"rules"`
	// The `attributes` `version`.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesetRespDataAttributes instantiates a new RulesetRespDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesetRespDataAttributes(created RulesetRespDataAttributesCreated, enabled bool, lastModifiedUserUuid string, modified RulesetRespDataAttributesModified, name string, position int32, rules []RulesetRespDataAttributesRulesItems, version int64) *RulesetRespDataAttributes {
	this := RulesetRespDataAttributes{}
	this.Created = created
	this.Enabled = enabled
	this.LastModifiedUserUuid = lastModifiedUserUuid
	this.Modified = modified
	this.Name = name
	this.Position = position
	this.Rules = rules
	this.Version = version
	return &this
}

// NewRulesetRespDataAttributesWithDefaults instantiates a new RulesetRespDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesetRespDataAttributesWithDefaults() *RulesetRespDataAttributes {
	this := RulesetRespDataAttributes{}
	return &this
}

// GetCreated returns the Created field value.
func (o *RulesetRespDataAttributes) GetCreated() RulesetRespDataAttributesCreated {
	if o == nil {
		var ret RulesetRespDataAttributesCreated
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetCreatedOk() (*RulesetRespDataAttributesCreated, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *RulesetRespDataAttributes) SetCreated(v RulesetRespDataAttributesCreated) {
	o.Created = v
}

// GetEnabled returns the Enabled field value.
func (o *RulesetRespDataAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RulesetRespDataAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLastModifiedUserUuid returns the LastModifiedUserUuid field value.
func (o *RulesetRespDataAttributes) GetLastModifiedUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastModifiedUserUuid
}

// GetLastModifiedUserUuidOk returns a tuple with the LastModifiedUserUuid field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetLastModifiedUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedUserUuid, true
}

// SetLastModifiedUserUuid sets field value.
func (o *RulesetRespDataAttributes) SetLastModifiedUserUuid(v string) {
	o.LastModifiedUserUuid = v
}

// GetModified returns the Modified field value.
func (o *RulesetRespDataAttributes) GetModified() RulesetRespDataAttributesModified {
	if o == nil {
		var ret RulesetRespDataAttributesModified
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetModifiedOk() (*RulesetRespDataAttributesModified, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *RulesetRespDataAttributes) SetModified(v RulesetRespDataAttributesModified) {
	o.Modified = v
}

// GetName returns the Name field value.
func (o *RulesetRespDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RulesetRespDataAttributes) SetName(v string) {
	o.Name = v
}

// GetPosition returns the Position field value.
func (o *RulesetRespDataAttributes) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *RulesetRespDataAttributes) SetPosition(v int32) {
	o.Position = v
}

// GetProcessingStatus returns the ProcessingStatus field value if set, zero value otherwise.
func (o *RulesetRespDataAttributes) GetProcessingStatus() string {
	if o == nil || o.ProcessingStatus == nil {
		var ret string
		return ret
	}
	return *o.ProcessingStatus
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetProcessingStatusOk() (*string, bool) {
	if o == nil || o.ProcessingStatus == nil {
		return nil, false
	}
	return o.ProcessingStatus, true
}

// HasProcessingStatus returns a boolean if a field has been set.
func (o *RulesetRespDataAttributes) HasProcessingStatus() bool {
	return o != nil && o.ProcessingStatus != nil
}

// SetProcessingStatus gets a reference to the given string and assigns it to the ProcessingStatus field.
func (o *RulesetRespDataAttributes) SetProcessingStatus(v string) {
	o.ProcessingStatus = &v
}

// GetRules returns the Rules field value.
func (o *RulesetRespDataAttributes) GetRules() []RulesetRespDataAttributesRulesItems {
	if o == nil {
		var ret []RulesetRespDataAttributesRulesItems
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetRulesOk() (*[]RulesetRespDataAttributesRulesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *RulesetRespDataAttributes) SetRules(v []RulesetRespDataAttributesRulesItems) {
	o.Rules = v
}

// GetVersion returns the Version field value.
func (o *RulesetRespDataAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *RulesetRespDataAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesetRespDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created"] = o.Created
	toSerialize["enabled"] = o.Enabled
	toSerialize["last_modified_user_uuid"] = o.LastModifiedUserUuid
	toSerialize["modified"] = o.Modified
	toSerialize["name"] = o.Name
	toSerialize["position"] = o.Position
	if o.ProcessingStatus != nil {
		toSerialize["processing_status"] = o.ProcessingStatus
	}
	toSerialize["rules"] = o.Rules
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RulesetRespDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created              *RulesetRespDataAttributesCreated      `json:"created"`
		Enabled              *bool                                  `json:"enabled"`
		LastModifiedUserUuid *string                                `json:"last_modified_user_uuid"`
		Modified             *RulesetRespDataAttributesModified     `json:"modified"`
		Name                 *string                                `json:"name"`
		Position             *int32                                 `json:"position"`
		ProcessingStatus     *string                                `json:"processing_status,omitempty"`
		Rules                *[]RulesetRespDataAttributesRulesItems `json:"rules"`
		Version              *int64                                 `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.LastModifiedUserUuid == nil {
		return fmt.Errorf("required field last_modified_user_uuid missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Position == nil {
		return fmt.Errorf("required field position missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "enabled", "last_modified_user_uuid", "modified", "name", "position", "processing_status", "rules", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Created.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Created = *all.Created
	o.Enabled = *all.Enabled
	o.LastModifiedUserUuid = *all.LastModifiedUserUuid
	if all.Modified.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Modified = *all.Modified
	o.Name = *all.Name
	o.Position = *all.Position
	o.ProcessingStatus = all.ProcessingStatus
	o.Rules = *all.Rules
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
