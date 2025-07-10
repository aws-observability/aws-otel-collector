// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchNotificationRuleParametersDataAttributes Attributes of the notification rule patch request. It is required to update the version of the rule when patching it.
type PatchNotificationRuleParametersDataAttributes struct {
	// Field used to enable or disable the rule.
	Enabled *bool `json:"enabled,omitempty"`
	// Name of the notification rule.
	Name *string `json:"name,omitempty"`
	// Selectors are used to filter security issues for which notifications should be generated.
	// Users can specify rule severities, rule types, a query to filter security issues on tags and attributes, and the trigger source.
	// Only the trigger_source field is required.
	Selectors *Selectors `json:"selectors,omitempty"`
	// List of recipients to notify when a notification rule is triggered. Many different target types are supported,
	// such as email addresses, Slack channels, and PagerDuty services.
	// The appropriate integrations need to be properly configured to send notifications to the specified targets.
	Targets []string `json:"targets,omitempty"`
	// Time aggregation period (in seconds) is used to aggregate the results of the notification rule evaluation.
	// Results are aggregated over a selected time frame using a rolling window, which updates with each new evaluation.
	// Notifications are only sent for new issues discovered during the window.
	// Time aggregation is only available for vulnerability-based notification rules. When omitted or set to 0, no aggregation
	// is done.
	TimeAggregation *int64 `json:"time_aggregation,omitempty"`
	// Version of the notification rule. It is updated when the rule is modified.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchNotificationRuleParametersDataAttributes instantiates a new PatchNotificationRuleParametersDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchNotificationRuleParametersDataAttributes() *PatchNotificationRuleParametersDataAttributes {
	this := PatchNotificationRuleParametersDataAttributes{}
	return &this
}

// NewPatchNotificationRuleParametersDataAttributesWithDefaults instantiates a new PatchNotificationRuleParametersDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchNotificationRuleParametersDataAttributesWithDefaults() *PatchNotificationRuleParametersDataAttributes {
	this := PatchNotificationRuleParametersDataAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchNotificationRuleParametersDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNotificationRuleParametersDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchNotificationRuleParametersDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchNotificationRuleParametersDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchNotificationRuleParametersDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNotificationRuleParametersDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchNotificationRuleParametersDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchNotificationRuleParametersDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *PatchNotificationRuleParametersDataAttributes) GetSelectors() Selectors {
	if o == nil || o.Selectors == nil {
		var ret Selectors
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNotificationRuleParametersDataAttributes) GetSelectorsOk() (*Selectors, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *PatchNotificationRuleParametersDataAttributes) HasSelectors() bool {
	return o != nil && o.Selectors != nil
}

// SetSelectors gets a reference to the given Selectors and assigns it to the Selectors field.
func (o *PatchNotificationRuleParametersDataAttributes) SetSelectors(v Selectors) {
	o.Selectors = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *PatchNotificationRuleParametersDataAttributes) GetTargets() []string {
	if o == nil || o.Targets == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNotificationRuleParametersDataAttributes) GetTargetsOk() (*[]string, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *PatchNotificationRuleParametersDataAttributes) HasTargets() bool {
	return o != nil && o.Targets != nil
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *PatchNotificationRuleParametersDataAttributes) SetTargets(v []string) {
	o.Targets = v
}

// GetTimeAggregation returns the TimeAggregation field value if set, zero value otherwise.
func (o *PatchNotificationRuleParametersDataAttributes) GetTimeAggregation() int64 {
	if o == nil || o.TimeAggregation == nil {
		var ret int64
		return ret
	}
	return *o.TimeAggregation
}

// GetTimeAggregationOk returns a tuple with the TimeAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNotificationRuleParametersDataAttributes) GetTimeAggregationOk() (*int64, bool) {
	if o == nil || o.TimeAggregation == nil {
		return nil, false
	}
	return o.TimeAggregation, true
}

// HasTimeAggregation returns a boolean if a field has been set.
func (o *PatchNotificationRuleParametersDataAttributes) HasTimeAggregation() bool {
	return o != nil && o.TimeAggregation != nil
}

// SetTimeAggregation gets a reference to the given int64 and assigns it to the TimeAggregation field.
func (o *PatchNotificationRuleParametersDataAttributes) SetTimeAggregation(v int64) {
	o.TimeAggregation = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PatchNotificationRuleParametersDataAttributes) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchNotificationRuleParametersDataAttributes) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PatchNotificationRuleParametersDataAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *PatchNotificationRuleParametersDataAttributes) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchNotificationRuleParametersDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.TimeAggregation != nil {
		toSerialize["time_aggregation"] = o.TimeAggregation
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
func (o *PatchNotificationRuleParametersDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled         *bool      `json:"enabled,omitempty"`
		Name            *string    `json:"name,omitempty"`
		Selectors       *Selectors `json:"selectors,omitempty"`
		Targets         []string   `json:"targets,omitempty"`
		TimeAggregation *int64     `json:"time_aggregation,omitempty"`
		Version         *int64     `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "name", "selectors", "targets", "time_aggregation", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Name = all.Name
	if all.Selectors != nil && all.Selectors.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Selectors = all.Selectors
	o.Targets = all.Targets
	o.TimeAggregation = all.TimeAggregation
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
