// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleCreateAttributes Attributes for creating a new schedule.
type FleetScheduleCreateAttributes struct {
	// Human-readable name for the schedule.
	Name string `json:"name"`
	// Query used to filter and select target hosts for scheduled deployments. Uses the Datadog query syntax.
	Query string `json:"query"`
	// Defines the recurrence pattern for the schedule. Specifies when deployments should be
	// automatically triggered based on maintenance windows.
	Rule FleetScheduleRecurrenceRule `json:"rule"`
	// The status of the schedule.
	// - `active`: The schedule is active and will create deployments according to its recurrence rule.
	// - `inactive`: The schedule is inactive and will not create any deployments.
	Status *FleetScheduleStatus `json:"status,omitempty"`
	// Number of major versions behind the latest to target for upgrades.
	// - 0: Always upgrade to the latest version (default)
	// - 1: Upgrade to latest minus 1 major version
	// - 2: Upgrade to latest minus 2 major versions
	// Maximum value is 2.
	VersionToLatest *int64 `json:"version_to_latest,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetScheduleCreateAttributes instantiates a new FleetScheduleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetScheduleCreateAttributes(name string, query string, rule FleetScheduleRecurrenceRule) *FleetScheduleCreateAttributes {
	this := FleetScheduleCreateAttributes{}
	this.Name = name
	this.Query = query
	this.Rule = rule
	return &this
}

// NewFleetScheduleCreateAttributesWithDefaults instantiates a new FleetScheduleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetScheduleCreateAttributesWithDefaults() *FleetScheduleCreateAttributes {
	this := FleetScheduleCreateAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *FleetScheduleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FleetScheduleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *FleetScheduleCreateAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FleetScheduleCreateAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRule returns the Rule field value.
func (o *FleetScheduleCreateAttributes) GetRule() FleetScheduleRecurrenceRule {
	if o == nil {
		var ret FleetScheduleRecurrenceRule
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *FleetScheduleCreateAttributes) GetRuleOk() (*FleetScheduleRecurrenceRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *FleetScheduleCreateAttributes) SetRule(v FleetScheduleRecurrenceRule) {
	o.Rule = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetScheduleCreateAttributes) GetStatus() FleetScheduleStatus {
	if o == nil || o.Status == nil {
		var ret FleetScheduleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleCreateAttributes) GetStatusOk() (*FleetScheduleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetScheduleCreateAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FleetScheduleStatus and assigns it to the Status field.
func (o *FleetScheduleCreateAttributes) SetStatus(v FleetScheduleStatus) {
	o.Status = &v
}

// GetVersionToLatest returns the VersionToLatest field value if set, zero value otherwise.
func (o *FleetScheduleCreateAttributes) GetVersionToLatest() int64 {
	if o == nil || o.VersionToLatest == nil {
		var ret int64
		return ret
	}
	return *o.VersionToLatest
}

// GetVersionToLatestOk returns a tuple with the VersionToLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleCreateAttributes) GetVersionToLatestOk() (*int64, bool) {
	if o == nil || o.VersionToLatest == nil {
		return nil, false
	}
	return o.VersionToLatest, true
}

// HasVersionToLatest returns a boolean if a field has been set.
func (o *FleetScheduleCreateAttributes) HasVersionToLatest() bool {
	return o != nil && o.VersionToLatest != nil
}

// SetVersionToLatest gets a reference to the given int64 and assigns it to the VersionToLatest field.
func (o *FleetScheduleCreateAttributes) SetVersionToLatest(v int64) {
	o.VersionToLatest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetScheduleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	toSerialize["rule"] = o.Rule
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.VersionToLatest != nil {
		toSerialize["version_to_latest"] = o.VersionToLatest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetScheduleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string                      `json:"name"`
		Query           *string                      `json:"query"`
		Rule            *FleetScheduleRecurrenceRule `json:"rule"`
		Status          *FleetScheduleStatus         `json:"status,omitempty"`
		VersionToLatest *int64                       `json:"version_to_latest,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "query", "rule", "status", "version_to_latest"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Query = *all.Query
	if all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = *all.Rule
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.VersionToLatest = all.VersionToLatest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
