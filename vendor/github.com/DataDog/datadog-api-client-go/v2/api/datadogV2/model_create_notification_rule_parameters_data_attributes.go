// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateNotificationRuleParametersDataAttributes Attributes of the notification rule create request.
type CreateNotificationRuleParametersDataAttributes struct {
	// Field used to enable or disable the rule.
	Enabled *bool `json:"enabled,omitempty"`
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
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateNotificationRuleParametersDataAttributes instantiates a new CreateNotificationRuleParametersDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateNotificationRuleParametersDataAttributes(name string, selectors Selectors, targets []string) *CreateNotificationRuleParametersDataAttributes {
	this := CreateNotificationRuleParametersDataAttributes{}
	this.Name = name
	this.Selectors = selectors
	this.Targets = targets
	return &this
}

// NewCreateNotificationRuleParametersDataAttributesWithDefaults instantiates a new CreateNotificationRuleParametersDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateNotificationRuleParametersDataAttributesWithDefaults() *CreateNotificationRuleParametersDataAttributes {
	this := CreateNotificationRuleParametersDataAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateNotificationRuleParametersDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRuleParametersDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateNotificationRuleParametersDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateNotificationRuleParametersDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value.
func (o *CreateNotificationRuleParametersDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRuleParametersDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateNotificationRuleParametersDataAttributes) SetName(v string) {
	o.Name = v
}

// GetSelectors returns the Selectors field value.
func (o *CreateNotificationRuleParametersDataAttributes) GetSelectors() Selectors {
	if o == nil {
		var ret Selectors
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRuleParametersDataAttributes) GetSelectorsOk() (*Selectors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selectors, true
}

// SetSelectors sets field value.
func (o *CreateNotificationRuleParametersDataAttributes) SetSelectors(v Selectors) {
	o.Selectors = v
}

// GetTargets returns the Targets field value.
func (o *CreateNotificationRuleParametersDataAttributes) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRuleParametersDataAttributes) GetTargetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *CreateNotificationRuleParametersDataAttributes) SetTargets(v []string) {
	o.Targets = v
}

// GetTimeAggregation returns the TimeAggregation field value if set, zero value otherwise.
func (o *CreateNotificationRuleParametersDataAttributes) GetTimeAggregation() int64 {
	if o == nil || o.TimeAggregation == nil {
		var ret int64
		return ret
	}
	return *o.TimeAggregation
}

// GetTimeAggregationOk returns a tuple with the TimeAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRuleParametersDataAttributes) GetTimeAggregationOk() (*int64, bool) {
	if o == nil || o.TimeAggregation == nil {
		return nil, false
	}
	return o.TimeAggregation, true
}

// HasTimeAggregation returns a boolean if a field has been set.
func (o *CreateNotificationRuleParametersDataAttributes) HasTimeAggregation() bool {
	return o != nil && o.TimeAggregation != nil
}

// SetTimeAggregation gets a reference to the given int64 and assigns it to the TimeAggregation field.
func (o *CreateNotificationRuleParametersDataAttributes) SetTimeAggregation(v int64) {
	o.TimeAggregation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateNotificationRuleParametersDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["name"] = o.Name
	toSerialize["selectors"] = o.Selectors
	toSerialize["targets"] = o.Targets
	if o.TimeAggregation != nil {
		toSerialize["time_aggregation"] = o.TimeAggregation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateNotificationRuleParametersDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled         *bool      `json:"enabled,omitempty"`
		Name            *string    `json:"name"`
		Selectors       *Selectors `json:"selectors"`
		Targets         *[]string  `json:"targets"`
		TimeAggregation *int64     `json:"time_aggregation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "name", "selectors", "targets", "time_aggregation"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Name = *all.Name
	if all.Selectors.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Selectors = *all.Selectors
	o.Targets = *all.Targets
	o.TimeAggregation = all.TimeAggregation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
