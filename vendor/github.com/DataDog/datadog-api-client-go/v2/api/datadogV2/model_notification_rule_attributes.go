// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRuleAttributes Attributes of the notification rule.
type NotificationRuleAttributes struct {
	// Date as Unix timestamp in milliseconds.
	CreatedAt int64 `json:"created_at"`
	// User creating or modifying a rule.
	CreatedBy RuleUser `json:"created_by"`
	// Field used to enable or disable the rule.
	Enabled bool `json:"enabled"`
	// Date as Unix timestamp in milliseconds.
	ModifiedAt int64 `json:"modified_at"`
	// User creating or modifying a rule.
	ModifiedBy RuleUser `json:"modified_by"`
	// Name of the notification rule.
	Name string `json:"name"`
	// Selectors are used to filter security issues for which notifications should be generated.
	// Users can specify rule severities, rule types, a query to filter security issues on tags and attributes, and the trigger source.
	// Only the trigger_source field is required.
	Selectors Selectors `json:"selectors"`
	// List of recipients to notify when a notification rule is triggered. Many different target types are supported,
	// such as email addresses, Slack channels, and PagerDuty services.
	// The appropriate integrations need to be properly configured to send notifications to the specified targets.
	Targets []string `json:"targets"`
	// Time aggregation period (in seconds) is used to aggregate the results of the notification rule evaluation.
	// Results are aggregated over a selected time frame using a rolling window, which updates with each new evaluation.
	// Notifications are only sent for new issues discovered during the window.
	// Time aggregation is only available for vulnerability-based notification rules. When omitted or set to 0, no aggregation
	// is done.
	TimeAggregation *int64 `json:"time_aggregation,omitempty"`
	// Version of the notification rule. It is updated when the rule is modified.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationRuleAttributes instantiates a new NotificationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationRuleAttributes(createdAt int64, createdBy RuleUser, enabled bool, modifiedAt int64, modifiedBy RuleUser, name string, selectors Selectors, targets []string, version int64) *NotificationRuleAttributes {
	this := NotificationRuleAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.Selectors = selectors
	this.Targets = targets
	this.Version = version
	return &this
}

// NewNotificationRuleAttributesWithDefaults instantiates a new NotificationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationRuleAttributesWithDefaults() *NotificationRuleAttributes {
	this := NotificationRuleAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *NotificationRuleAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *NotificationRuleAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *NotificationRuleAttributes) GetCreatedBy() RuleUser {
	if o == nil {
		var ret RuleUser
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetCreatedByOk() (*RuleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *NotificationRuleAttributes) SetCreatedBy(v RuleUser) {
	o.CreatedBy = v
}

// GetEnabled returns the Enabled field value.
func (o *NotificationRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *NotificationRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *NotificationRuleAttributes) GetModifiedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetModifiedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *NotificationRuleAttributes) SetModifiedAt(v int64) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *NotificationRuleAttributes) GetModifiedBy() RuleUser {
	if o == nil {
		var ret RuleUser
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetModifiedByOk() (*RuleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *NotificationRuleAttributes) SetModifiedBy(v RuleUser) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *NotificationRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NotificationRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetSelectors returns the Selectors field value.
func (o *NotificationRuleAttributes) GetSelectors() Selectors {
	if o == nil {
		var ret Selectors
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetSelectorsOk() (*Selectors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selectors, true
}

// SetSelectors sets field value.
func (o *NotificationRuleAttributes) SetSelectors(v Selectors) {
	o.Selectors = v
}

// GetTargets returns the Targets field value.
func (o *NotificationRuleAttributes) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetTargetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *NotificationRuleAttributes) SetTargets(v []string) {
	o.Targets = v
}

// GetTimeAggregation returns the TimeAggregation field value if set, zero value otherwise.
func (o *NotificationRuleAttributes) GetTimeAggregation() int64 {
	if o == nil || o.TimeAggregation == nil {
		var ret int64
		return ret
	}
	return *o.TimeAggregation
}

// GetTimeAggregationOk returns a tuple with the TimeAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetTimeAggregationOk() (*int64, bool) {
	if o == nil || o.TimeAggregation == nil {
		return nil, false
	}
	return o.TimeAggregation, true
}

// HasTimeAggregation returns a boolean if a field has been set.
func (o *NotificationRuleAttributes) HasTimeAggregation() bool {
	return o != nil && o.TimeAggregation != nil
}

// SetTimeAggregation gets a reference to the given int64 and assigns it to the TimeAggregation field.
func (o *NotificationRuleAttributes) SetTimeAggregation(v int64) {
	o.TimeAggregation = &v
}

// GetVersion returns the Version field value.
func (o *NotificationRuleAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *NotificationRuleAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["enabled"] = o.Enabled
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["selectors"] = o.Selectors
	toSerialize["targets"] = o.Targets
	if o.TimeAggregation != nil {
		toSerialize["time_aggregation"] = o.TimeAggregation
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *int64     `json:"created_at"`
		CreatedBy       *RuleUser  `json:"created_by"`
		Enabled         *bool      `json:"enabled"`
		ModifiedAt      *int64     `json:"modified_at"`
		ModifiedBy      *RuleUser  `json:"modified_by"`
		Name            *string    `json:"name"`
		Selectors       *Selectors `json:"selectors"`
		Targets         *[]string  `json:"targets"`
		TimeAggregation *int64     `json:"time_aggregation,omitempty"`
		Version         *int64     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Selectors == nil {
		return fmt.Errorf("required field selectors missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "enabled", "modified_at", "modified_by", "name", "selectors", "targets", "time_aggregation", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	o.Enabled = *all.Enabled
	o.ModifiedAt = *all.ModifiedAt
	if all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	if all.Selectors.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Selectors = *all.Selectors
	o.Targets = *all.Targets
	o.TimeAggregation = all.TimeAggregation
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
