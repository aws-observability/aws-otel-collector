// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleAttributes Attributes of a schedule in the response.
type FleetScheduleAttributes struct {
	// Unix timestamp (seconds since epoch) when the schedule was created.
	CreatedAtUnix *int64 `json:"created_at_unix,omitempty"`
	// User handle of the person who created the schedule.
	CreatedBy *string `json:"created_by,omitempty"`
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
	// Unix timestamp (seconds since epoch) when the schedule was last updated.
	UpdatedAtUnix *int64 `json:"updated_at_unix,omitempty"`
	// User handle of the person who last updated the schedule.
	UpdatedBy *string `json:"updated_by,omitempty"`
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

// NewFleetScheduleAttributes instantiates a new FleetScheduleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetScheduleAttributes() *FleetScheduleAttributes {
	this := FleetScheduleAttributes{}
	return &this
}

// NewFleetScheduleAttributesWithDefaults instantiates a new FleetScheduleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetScheduleAttributesWithDefaults() *FleetScheduleAttributes {
	this := FleetScheduleAttributes{}
	return &this
}

// GetCreatedAtUnix returns the CreatedAtUnix field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetCreatedAtUnix() int64 {
	if o == nil || o.CreatedAtUnix == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAtUnix
}

// GetCreatedAtUnixOk returns a tuple with the CreatedAtUnix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetCreatedAtUnixOk() (*int64, bool) {
	if o == nil || o.CreatedAtUnix == nil {
		return nil, false
	}
	return o.CreatedAtUnix, true
}

// HasCreatedAtUnix returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasCreatedAtUnix() bool {
	return o != nil && o.CreatedAtUnix != nil
}

// SetCreatedAtUnix gets a reference to the given int64 and assigns it to the CreatedAtUnix field.
func (o *FleetScheduleAttributes) SetCreatedAtUnix(v int64) {
	o.CreatedAtUnix = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *FleetScheduleAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetScheduleAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetScheduleAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetRule() FleetScheduleRecurrenceRule {
	if o == nil || o.Rule == nil {
		var ret FleetScheduleRecurrenceRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetRuleOk() (*FleetScheduleRecurrenceRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given FleetScheduleRecurrenceRule and assigns it to the Rule field.
func (o *FleetScheduleAttributes) SetRule(v FleetScheduleRecurrenceRule) {
	o.Rule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetStatus() FleetScheduleStatus {
	if o == nil || o.Status == nil {
		var ret FleetScheduleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetStatusOk() (*FleetScheduleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FleetScheduleStatus and assigns it to the Status field.
func (o *FleetScheduleAttributes) SetStatus(v FleetScheduleStatus) {
	o.Status = &v
}

// GetUpdatedAtUnix returns the UpdatedAtUnix field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetUpdatedAtUnix() int64 {
	if o == nil || o.UpdatedAtUnix == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAtUnix
}

// GetUpdatedAtUnixOk returns a tuple with the UpdatedAtUnix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetUpdatedAtUnixOk() (*int64, bool) {
	if o == nil || o.UpdatedAtUnix == nil {
		return nil, false
	}
	return o.UpdatedAtUnix, true
}

// HasUpdatedAtUnix returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasUpdatedAtUnix() bool {
	return o != nil && o.UpdatedAtUnix != nil
}

// SetUpdatedAtUnix gets a reference to the given int64 and assigns it to the UpdatedAtUnix field.
func (o *FleetScheduleAttributes) SetUpdatedAtUnix(v int64) {
	o.UpdatedAtUnix = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *FleetScheduleAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetVersionToLatest returns the VersionToLatest field value if set, zero value otherwise.
func (o *FleetScheduleAttributes) GetVersionToLatest() int64 {
	if o == nil || o.VersionToLatest == nil {
		var ret int64
		return ret
	}
	return *o.VersionToLatest
}

// GetVersionToLatestOk returns a tuple with the VersionToLatest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleAttributes) GetVersionToLatestOk() (*int64, bool) {
	if o == nil || o.VersionToLatest == nil {
		return nil, false
	}
	return o.VersionToLatest, true
}

// HasVersionToLatest returns a boolean if a field has been set.
func (o *FleetScheduleAttributes) HasVersionToLatest() bool {
	return o != nil && o.VersionToLatest != nil
}

// SetVersionToLatest gets a reference to the given int64 and assigns it to the VersionToLatest field.
func (o *FleetScheduleAttributes) SetVersionToLatest(v int64) {
	o.VersionToLatest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetScheduleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAtUnix != nil {
		toSerialize["created_at_unix"] = o.CreatedAtUnix
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
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
	if o.UpdatedAtUnix != nil {
		toSerialize["updated_at_unix"] = o.UpdatedAtUnix
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
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
func (o *FleetScheduleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAtUnix   *int64                       `json:"created_at_unix,omitempty"`
		CreatedBy       *string                      `json:"created_by,omitempty"`
		Name            *string                      `json:"name,omitempty"`
		Query           *string                      `json:"query,omitempty"`
		Rule            *FleetScheduleRecurrenceRule `json:"rule,omitempty"`
		Status          *FleetScheduleStatus         `json:"status,omitempty"`
		UpdatedAtUnix   *int64                       `json:"updated_at_unix,omitempty"`
		UpdatedBy       *string                      `json:"updated_by,omitempty"`
		VersionToLatest *int64                       `json:"version_to_latest,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at_unix", "created_by", "name", "query", "rule", "status", "updated_at_unix", "updated_by", "version_to_latest"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAtUnix = all.CreatedAtUnix
	o.CreatedBy = all.CreatedBy
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
	o.UpdatedAtUnix = all.UpdatedAtUnix
	o.UpdatedBy = all.UpdatedBy
	o.VersionToLatest = all.VersionToLatest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
