// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionAttributes The attributes of the suppression rule.
type SecurityMonitoringSuppressionAttributes struct {
	// A Unix millisecond timestamp given the creation date of the suppression rule.
	CreationDate *int64 `json:"creation_date,omitempty"`
	// A user.
	Creator *SecurityMonitoringUser `json:"creator,omitempty"`
	// A description for the suppression rule.
	Description *string `json:"description,omitempty"`
	// Whether the suppression rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// A Unix millisecond timestamp giving an expiration date for the suppression rule. After this date, it won't suppress signals anymore.
	ExpirationDate *int64 `json:"expiration_date,omitempty"`
	// The name of the suppression rule.
	Name *string `json:"name,omitempty"`
	// The rule query of the suppression rule, with the same syntax as the search bar for detection rules.
	RuleQuery *string `json:"rule_query,omitempty"`
	// The suppression query of the suppression rule. If a signal matches this query, it is suppressed and not triggered. Same syntax as the queries to search signals in the signal explorer.
	SuppressionQuery *string `json:"suppression_query,omitempty"`
	// A Unix millisecond timestamp given the update date of the suppression rule.
	UpdateDate *int64 `json:"update_date,omitempty"`
	// A user.
	Updater *SecurityMonitoringUser `json:"updater,omitempty"`
	// The version of the suppression rule; it starts at 1, and is incremented at each update.
	Version *int32 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringSuppressionAttributes instantiates a new SecurityMonitoringSuppressionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSuppressionAttributes() *SecurityMonitoringSuppressionAttributes {
	this := SecurityMonitoringSuppressionAttributes{}
	return &this
}

// NewSecurityMonitoringSuppressionAttributesWithDefaults instantiates a new SecurityMonitoringSuppressionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSuppressionAttributesWithDefaults() *SecurityMonitoringSuppressionAttributes {
	this := SecurityMonitoringSuppressionAttributes{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetCreationDate() int64 {
	if o == nil || o.CreationDate == nil {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetCreationDateOk() (*int64, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasCreationDate() bool {
	return o != nil && o.CreationDate != nil
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *SecurityMonitoringSuppressionAttributes) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetCreator() SecurityMonitoringUser {
	if o == nil || o.Creator == nil {
		var ret SecurityMonitoringUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetCreatorOk() (*SecurityMonitoringUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given SecurityMonitoringUser and assigns it to the Creator field.
func (o *SecurityMonitoringSuppressionAttributes) SetCreator(v SecurityMonitoringUser) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringSuppressionAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringSuppressionAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetExpirationDate() int64 {
	if o == nil || o.ExpirationDate == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetExpirationDateOk() (*int64, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *SecurityMonitoringSuppressionAttributes) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringSuppressionAttributes) SetName(v string) {
	o.Name = &v
}

// GetRuleQuery returns the RuleQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetRuleQuery() string {
	if o == nil || o.RuleQuery == nil {
		var ret string
		return ret
	}
	return *o.RuleQuery
}

// GetRuleQueryOk returns a tuple with the RuleQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetRuleQueryOk() (*string, bool) {
	if o == nil || o.RuleQuery == nil {
		return nil, false
	}
	return o.RuleQuery, true
}

// HasRuleQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasRuleQuery() bool {
	return o != nil && o.RuleQuery != nil
}

// SetRuleQuery gets a reference to the given string and assigns it to the RuleQuery field.
func (o *SecurityMonitoringSuppressionAttributes) SetRuleQuery(v string) {
	o.RuleQuery = &v
}

// GetSuppressionQuery returns the SuppressionQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetSuppressionQuery() string {
	if o == nil || o.SuppressionQuery == nil {
		var ret string
		return ret
	}
	return *o.SuppressionQuery
}

// GetSuppressionQueryOk returns a tuple with the SuppressionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetSuppressionQueryOk() (*string, bool) {
	if o == nil || o.SuppressionQuery == nil {
		return nil, false
	}
	return o.SuppressionQuery, true
}

// HasSuppressionQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasSuppressionQuery() bool {
	return o != nil && o.SuppressionQuery != nil
}

// SetSuppressionQuery gets a reference to the given string and assigns it to the SuppressionQuery field.
func (o *SecurityMonitoringSuppressionAttributes) SetSuppressionQuery(v string) {
	o.SuppressionQuery = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetUpdateDate() int64 {
	if o == nil || o.UpdateDate == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetUpdateDateOk() (*int64, bool) {
	if o == nil || o.UpdateDate == nil {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasUpdateDate() bool {
	return o != nil && o.UpdateDate != nil
}

// SetUpdateDate gets a reference to the given int64 and assigns it to the UpdateDate field.
func (o *SecurityMonitoringSuppressionAttributes) SetUpdateDate(v int64) {
	o.UpdateDate = &v
}

// GetUpdater returns the Updater field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetUpdater() SecurityMonitoringUser {
	if o == nil || o.Updater == nil {
		var ret SecurityMonitoringUser
		return ret
	}
	return *o.Updater
}

// GetUpdaterOk returns a tuple with the Updater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetUpdaterOk() (*SecurityMonitoringUser, bool) {
	if o == nil || o.Updater == nil {
		return nil, false
	}
	return o.Updater, true
}

// HasUpdater returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasUpdater() bool {
	return o != nil && o.Updater != nil
}

// SetUpdater gets a reference to the given SecurityMonitoringUser and assigns it to the Updater field.
func (o *SecurityMonitoringSuppressionAttributes) SetUpdater(v SecurityMonitoringUser) {
	o.Updater = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringSuppressionAttributes) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionAttributes) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringSuppressionAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecurityMonitoringSuppressionAttributes) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSuppressionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RuleQuery != nil {
		toSerialize["rule_query"] = o.RuleQuery
	}
	if o.SuppressionQuery != nil {
		toSerialize["suppression_query"] = o.SuppressionQuery
	}
	if o.UpdateDate != nil {
		toSerialize["update_date"] = o.UpdateDate
	}
	if o.Updater != nil {
		toSerialize["updater"] = o.Updater
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSuppressionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreationDate     *int64                  `json:"creation_date,omitempty"`
		Creator          *SecurityMonitoringUser `json:"creator,omitempty"`
		Description      *string                 `json:"description,omitempty"`
		Enabled          *bool                   `json:"enabled,omitempty"`
		ExpirationDate   *int64                  `json:"expiration_date,omitempty"`
		Name             *string                 `json:"name,omitempty"`
		RuleQuery        *string                 `json:"rule_query,omitempty"`
		SuppressionQuery *string                 `json:"suppression_query,omitempty"`
		UpdateDate       *int64                  `json:"update_date,omitempty"`
		Updater          *SecurityMonitoringUser `json:"updater,omitempty"`
		Version          *int32                  `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"creation_date", "creator", "description", "enabled", "expiration_date", "name", "rule_query", "suppression_query", "update_date", "updater", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreationDate = all.CreationDate
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.ExpirationDate = all.ExpirationDate
	o.Name = all.Name
	o.RuleQuery = all.RuleQuery
	o.SuppressionQuery = all.SuppressionQuery
	o.UpdateDate = all.UpdateDate
	if all.Updater != nil && all.Updater.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Updater = all.Updater
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
