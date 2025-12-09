// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetSchedulePatchAttributes Attributes for partially updating a schedule. All fields are optional.
type FleetSchedulePatchAttributes struct {
	// Human-readable name for the schedule.
	Name *string `json:"name,omitempty"`
	// Query used to filter and select target hosts for scheduled deployments. Uses the Datadog query syntax.
	Query *string `json:"query,omitempty"`
	// Defines the recurrence pattern for the schedule. Specifies when deployments should be
	// automatically triggered based on maintenance windows.
	Rule *FleetScheduleRecurrenceRule `json:"rule,omitempty"`
	// The status of the schedule.
	// - `active`: The schedule is active and will create deployments according to its recurrence rule.
	// - `inactive`: The schedule is inactive and will not create any deployments.
	Status *FleetScheduleStatus `json:"status,omitempty"`
	// Number of major versions behind the latest to target for upgrades.
	// - 0: Always upgrade to the latest version
	// - 1: Upgrade to latest minus 1 major version
	// - 2: Upgrade to latest minus 2 major versions
	// Maximum value is 2.
	VersionToLatest *int64 `json:"version_to_latest,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetSchedulePatchAttributes instantiates a new FleetSchedulePatchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetSchedulePatchAttributes() *FleetSchedulePatchAttributes {
	this := FleetSchedulePatchAttributes{}
	return &this
}

// NewFleetSchedulePatchAttributesWithDefaults instantiates a new FleetSchedulePatchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetSchedulePatchAttributesWithDefaults() *FleetSchedulePatchAttributes {
	this := FleetSchedulePatchAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetSchedulePatchAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetSchedulePatchAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetSchedulePatchAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetSchedulePatchAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetSchedulePatchAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetSchedulePatchAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetSchedulePatchAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetSchedulePatchAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FleetSchedulePatchAttributes) GetRule() FleetScheduleRecurrenceRule {
	if o == nil || o.Rule == nil {
		var ret FleetScheduleRecurrenceRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetSchedulePatchAttributes) GetRuleOk() (*FleetScheduleRecurrenceRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FleetSchedulePatchAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given FleetScheduleRecurrenceRule and assigns it to the Rule field.
func (o *FleetSchedulePatchAttributes) SetRule(v FleetScheduleRecurrenceRule) {
	o.Rule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetSchedulePatchAttributes) GetStatus() FleetScheduleStatus {
	if o == nil || o.Status == nil {
		var ret FleetScheduleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetSchedulePatchAttributes) GetStatusOk() (*FleetScheduleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetSchedulePatchAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FleetScheduleStatus and assigns it to the Status field.
func (o *FleetSchedulePatchAttributes) SetStatus(v FleetScheduleStatus) {
	o.Status = &v
}

// GetVersionToLatest returns the VersionToLatest field value if set, zero value otherwise.
func (o *FleetSchedulePatchAttributes) GetVersionToLatest() int64 {
	if o == nil || o.VersionToLatest == nil {
		var ret int64
		return ret
	}
	return *o.VersionToLatest
}

// GetVersionToLatestOk returns a tuple with the VersionToLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetSchedulePatchAttributes) GetVersionToLatestOk() (*int64, bool) {
	if o == nil || o.VersionToLatest == nil {
		return nil, false
	}
	return o.VersionToLatest, true
}

// HasVersionToLatest returns a boolean if a field has been set.
func (o *FleetSchedulePatchAttributes) HasVersionToLatest() bool {
	return o != nil && o.VersionToLatest != nil
}

// SetVersionToLatest gets a reference to the given int64 and assigns it to the VersionToLatest field.
func (o *FleetSchedulePatchAttributes) SetVersionToLatest(v int64) {
	o.VersionToLatest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetSchedulePatchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
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
func (o *FleetSchedulePatchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string                      `json:"name,omitempty"`
		Query           *string                      `json:"query,omitempty"`
		Rule            *FleetScheduleRecurrenceRule `json:"rule,omitempty"`
		Status          *FleetScheduleStatus         `json:"status,omitempty"`
		VersionToLatest *int64                       `json:"version_to_latest,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "query", "rule", "status", "version_to_latest"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Query = all.Query
	if all.Rule != nil && all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = all.Rule
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
