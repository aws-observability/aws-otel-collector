// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringCriticalAssetAttributes The attributes of the critical asset.
type SecurityMonitoringCriticalAssetAttributes struct {
	// ID of user who created the critical asset.
	CreationAuthorId *int64 `json:"creation_author_id,omitempty"`
	// A Unix millisecond timestamp given the creation date of the critical asset.
	CreationDate *int64 `json:"creation_date,omitempty"`
	// A user.
	Creator *SecurityMonitoringUser `json:"creator,omitempty"`
	// Whether the critical asset is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The query for the critical asset. It uses the same syntax as the queries to search signals in the Signals Explorer.
	Query *string `json:"query,omitempty"`
	// The rule query of the critical asset, with the same syntax as the search bar for detection rules. This determines which rules this critical asset will apply to.
	RuleQuery *string `json:"rule_query,omitempty"`
	// Severity associated with this critical asset. Either an explicit severity can be set, or the severity can be increased or decreased, or the severity can be left unchanged (no-op).
	Severity *SecurityMonitoringCriticalAssetSeverity `json:"severity,omitempty"`
	// List of tags associated with the critical asset.
	Tags []string `json:"tags,omitempty"`
	// ID of user who updated the critical asset.
	UpdateAuthorId *int64 `json:"update_author_id,omitempty"`
	// A Unix millisecond timestamp given the update date of the critical asset.
	UpdateDate *int64 `json:"update_date,omitempty"`
	// A user.
	Updater *SecurityMonitoringUser `json:"updater,omitempty"`
	// The version of the critical asset; it starts at 1, and is incremented at each update.
	Version *int32 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringCriticalAssetAttributes instantiates a new SecurityMonitoringCriticalAssetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringCriticalAssetAttributes() *SecurityMonitoringCriticalAssetAttributes {
	this := SecurityMonitoringCriticalAssetAttributes{}
	return &this
}

// NewSecurityMonitoringCriticalAssetAttributesWithDefaults instantiates a new SecurityMonitoringCriticalAssetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringCriticalAssetAttributesWithDefaults() *SecurityMonitoringCriticalAssetAttributes {
	this := SecurityMonitoringCriticalAssetAttributes{}
	return &this
}

// GetCreationAuthorId returns the CreationAuthorId field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetCreationAuthorId() int64 {
	if o == nil || o.CreationAuthorId == nil {
		var ret int64
		return ret
	}
	return *o.CreationAuthorId
}

// GetCreationAuthorIdOk returns a tuple with the CreationAuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetCreationAuthorIdOk() (*int64, bool) {
	if o == nil || o.CreationAuthorId == nil {
		return nil, false
	}
	return o.CreationAuthorId, true
}

// HasCreationAuthorId returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasCreationAuthorId() bool {
	return o != nil && o.CreationAuthorId != nil
}

// SetCreationAuthorId gets a reference to the given int64 and assigns it to the CreationAuthorId field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetCreationAuthorId(v int64) {
	o.CreationAuthorId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetCreationDate() int64 {
	if o == nil || o.CreationDate == nil {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetCreationDateOk() (*int64, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasCreationDate() bool {
	return o != nil && o.CreationDate != nil
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetCreator() SecurityMonitoringUser {
	if o == nil || o.Creator == nil {
		var ret SecurityMonitoringUser
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetCreatorOk() (*SecurityMonitoringUser, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given SecurityMonitoringUser and assigns it to the Creator field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetCreator(v SecurityMonitoringUser) {
	o.Creator = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRuleQuery returns the RuleQuery field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetRuleQuery() string {
	if o == nil || o.RuleQuery == nil {
		var ret string
		return ret
	}
	return *o.RuleQuery
}

// GetRuleQueryOk returns a tuple with the RuleQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetRuleQueryOk() (*string, bool) {
	if o == nil || o.RuleQuery == nil {
		return nil, false
	}
	return o.RuleQuery, true
}

// HasRuleQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasRuleQuery() bool {
	return o != nil && o.RuleQuery != nil
}

// SetRuleQuery gets a reference to the given string and assigns it to the RuleQuery field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetRuleQuery(v string) {
	o.RuleQuery = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetSeverity() SecurityMonitoringCriticalAssetSeverity {
	if o == nil || o.Severity == nil {
		var ret SecurityMonitoringCriticalAssetSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetSeverityOk() (*SecurityMonitoringCriticalAssetSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given SecurityMonitoringCriticalAssetSeverity and assigns it to the Severity field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetSeverity(v SecurityMonitoringCriticalAssetSeverity) {
	o.Severity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdateAuthorId returns the UpdateAuthorId field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetUpdateAuthorId() int64 {
	if o == nil || o.UpdateAuthorId == nil {
		var ret int64
		return ret
	}
	return *o.UpdateAuthorId
}

// GetUpdateAuthorIdOk returns a tuple with the UpdateAuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetUpdateAuthorIdOk() (*int64, bool) {
	if o == nil || o.UpdateAuthorId == nil {
		return nil, false
	}
	return o.UpdateAuthorId, true
}

// HasUpdateAuthorId returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasUpdateAuthorId() bool {
	return o != nil && o.UpdateAuthorId != nil
}

// SetUpdateAuthorId gets a reference to the given int64 and assigns it to the UpdateAuthorId field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetUpdateAuthorId(v int64) {
	o.UpdateAuthorId = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetUpdateDate() int64 {
	if o == nil || o.UpdateDate == nil {
		var ret int64
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetUpdateDateOk() (*int64, bool) {
	if o == nil || o.UpdateDate == nil {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasUpdateDate() bool {
	return o != nil && o.UpdateDate != nil
}

// SetUpdateDate gets a reference to the given int64 and assigns it to the UpdateDate field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetUpdateDate(v int64) {
	o.UpdateDate = &v
}

// GetUpdater returns the Updater field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetUpdater() SecurityMonitoringUser {
	if o == nil || o.Updater == nil {
		var ret SecurityMonitoringUser
		return ret
	}
	return *o.Updater
}

// GetUpdaterOk returns a tuple with the Updater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetUpdaterOk() (*SecurityMonitoringUser, bool) {
	if o == nil || o.Updater == nil {
		return nil, false
	}
	return o.Updater, true
}

// HasUpdater returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasUpdater() bool {
	return o != nil && o.Updater != nil
}

// SetUpdater gets a reference to the given SecurityMonitoringUser and assigns it to the Updater field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetUpdater(v SecurityMonitoringUser) {
	o.Updater = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringCriticalAssetAttributes) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringCriticalAssetAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecurityMonitoringCriticalAssetAttributes) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringCriticalAssetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreationAuthorId != nil {
		toSerialize["creation_author_id"] = o.CreationAuthorId
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RuleQuery != nil {
		toSerialize["rule_query"] = o.RuleQuery
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdateAuthorId != nil {
		toSerialize["update_author_id"] = o.UpdateAuthorId
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
func (o *SecurityMonitoringCriticalAssetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreationAuthorId *int64                                   `json:"creation_author_id,omitempty"`
		CreationDate     *int64                                   `json:"creation_date,omitempty"`
		Creator          *SecurityMonitoringUser                  `json:"creator,omitempty"`
		Enabled          *bool                                    `json:"enabled,omitempty"`
		Query            *string                                  `json:"query,omitempty"`
		RuleQuery        *string                                  `json:"rule_query,omitempty"`
		Severity         *SecurityMonitoringCriticalAssetSeverity `json:"severity,omitempty"`
		Tags             []string                                 `json:"tags,omitempty"`
		UpdateAuthorId   *int64                                   `json:"update_author_id,omitempty"`
		UpdateDate       *int64                                   `json:"update_date,omitempty"`
		Updater          *SecurityMonitoringUser                  `json:"updater,omitempty"`
		Version          *int32                                   `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"creation_author_id", "creation_date", "creator", "enabled", "query", "rule_query", "severity", "tags", "update_author_id", "update_date", "updater", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreationAuthorId = all.CreationAuthorId
	o.CreationDate = all.CreationDate
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Enabled = all.Enabled
	o.Query = all.Query
	o.RuleQuery = all.RuleQuery
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.Tags = all.Tags
	o.UpdateAuthorId = all.UpdateAuthorId
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
