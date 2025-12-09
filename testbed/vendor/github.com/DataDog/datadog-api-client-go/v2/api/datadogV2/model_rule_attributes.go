// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleAttributes Details of a rule.
type RuleAttributes struct {
	// The scorecard name to which this rule must belong.
	// Deprecated
	Category *string `json:"category,omitempty"`
	// Creation time of the rule outcome.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Defines if the rule is a custom rule.
	Custom *bool `json:"custom,omitempty"`
	// Explanation of the rule.
	Description *string `json:"description,omitempty"`
	// If enabled, the rule is calculated as part of the score.
	Enabled *bool `json:"enabled,omitempty"`
	// The maturity level of the rule (1, 2, or 3).
	Level *int32 `json:"level,omitempty"`
	// Time of the last rule outcome modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Owner of the rule.
	Owner *string `json:"owner,omitempty"`
	// The scorecard name to which this rule must belong.
	ScorecardName *string `json:"scorecard_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleAttributes instantiates a new RuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleAttributes() *RuleAttributes {
	this := RuleAttributes{}
	return &this
}

// NewRuleAttributesWithDefaults instantiates a new RuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleAttributesWithDefaults() *RuleAttributes {
	this := RuleAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
// Deprecated
func (o *RuleAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RuleAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *RuleAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
// Deprecated
func (o *RuleAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RuleAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RuleAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *RuleAttributes) GetCustom() bool {
	if o == nil || o.Custom == nil {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetCustomOk() (*bool, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *RuleAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *RuleAttributes) SetCustom(v bool) {
	o.Custom = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RuleAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RuleAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RuleAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RuleAttributes) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RuleAttributes) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *RuleAttributes) SetLevel(v int32) {
	o.Level = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *RuleAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *RuleAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *RuleAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *RuleAttributes) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *RuleAttributes) HasOwner() bool {
	return o != nil && o.Owner != nil
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *RuleAttributes) SetOwner(v string) {
	o.Owner = &v
}

// GetScorecardName returns the ScorecardName field value if set, zero value otherwise.
func (o *RuleAttributes) GetScorecardName() string {
	if o == nil || o.ScorecardName == nil {
		var ret string
		return ret
	}
	return *o.ScorecardName
}

// GetScorecardNameOk returns a tuple with the ScorecardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributes) GetScorecardNameOk() (*string, bool) {
	if o == nil || o.ScorecardName == nil {
		return nil, false
	}
	return o.ScorecardName, true
}

// HasScorecardName returns a boolean if a field has been set.
func (o *RuleAttributes) HasScorecardName() bool {
	return o != nil && o.ScorecardName != nil
}

// SetScorecardName gets a reference to the given string and assigns it to the ScorecardName field.
func (o *RuleAttributes) SetScorecardName(v string) {
	o.ScorecardName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ScorecardName != nil {
		toSerialize["scorecard_name"] = o.ScorecardName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category      *string    `json:"category,omitempty"`
		CreatedAt     *time.Time `json:"created_at,omitempty"`
		Custom        *bool      `json:"custom,omitempty"`
		Description   *string    `json:"description,omitempty"`
		Enabled       *bool      `json:"enabled,omitempty"`
		Level         *int32     `json:"level,omitempty"`
		ModifiedAt    *time.Time `json:"modified_at,omitempty"`
		Name          *string    `json:"name,omitempty"`
		Owner         *string    `json:"owner,omitempty"`
		ScorecardName *string    `json:"scorecard_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "created_at", "custom", "description", "enabled", "level", "modified_at", "name", "owner", "scorecard_name"})
	} else {
		return err
	}
	o.Category = all.Category
	o.CreatedAt = all.CreatedAt
	o.Custom = all.Custom
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.Level = all.Level
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.Owner = all.Owner
	o.ScorecardName = all.ScorecardName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
