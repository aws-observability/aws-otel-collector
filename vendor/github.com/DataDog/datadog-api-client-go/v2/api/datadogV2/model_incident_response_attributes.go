// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponseAttributes The incident's attributes from a response.
type IncidentResponseAttributes struct {
	// Timestamp of when the incident was archived.
	Archived datadog.NullableTime `json:"archived,omitempty"`
	// The incident case id.
	CaseId datadog.NullableInt64 `json:"case_id,omitempty"`
	// Timestamp when the incident was created.
	Created *time.Time `json:"created,omitempty"`
	// Length of the incident's customer impact in seconds.
	// Equals the difference between `customer_impact_start` and `customer_impact_end`.
	CustomerImpactDuration *int64 `json:"customer_impact_duration,omitempty"`
	// Timestamp when customers were no longer impacted by the incident.
	CustomerImpactEnd datadog.NullableTime `json:"customer_impact_end,omitempty"`
	// A summary of the impact customers experienced during the incident.
	CustomerImpactScope datadog.NullableString `json:"customer_impact_scope,omitempty"`
	// Timestamp when customers began being impacted by the incident.
	CustomerImpactStart datadog.NullableTime `json:"customer_impact_start,omitempty"`
	// A flag indicating whether the incident caused customer impact.
	CustomerImpacted *bool `json:"customer_impacted,omitempty"`
	// Timestamp when the incident was declared.
	Declared *time.Time `json:"declared,omitempty"`
	// Incident's non Datadog creator.
	DeclaredBy NullableIncidentNonDatadogCreator `json:"declared_by,omitempty"`
	// UUID of the user who declared the incident.
	DeclaredByUuid datadog.NullableString `json:"declared_by_uuid,omitempty"`
	// Timestamp when the incident was detected.
	Detected datadog.NullableTime `json:"detected,omitempty"`
	// A condensed view of the user-defined fields attached to incidents.
	Fields map[string]IncidentFieldAttributes `json:"fields,omitempty"`
	// A unique identifier that represents an incident type.
	IncidentTypeUuid *string `json:"incident_type_uuid,omitempty"`
	// A flag indicating whether the incident is a test incident.
	IsTest *bool `json:"is_test,omitempty"`
	// Timestamp when the incident was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Incident's non Datadog creator.
	NonDatadogCreator NullableIncidentNonDatadogCreator `json:"non_datadog_creator,omitempty"`
	// Notification handles that will be notified of the incident during update.
	NotificationHandles []IncidentNotificationHandle `json:"notification_handles,omitempty"`
	// The monotonically increasing integer ID for the incident.
	PublicId *int64 `json:"public_id,omitempty"`
	// Timestamp when the incident's state was last changed from active or stable to resolved or completed.
	Resolved datadog.NullableTime `json:"resolved,omitempty"`
	// The incident severity.
	Severity *IncidentSeverity `json:"severity,omitempty"`
	// The state incident.
	State datadog.NullableString `json:"state,omitempty"`
	// The amount of time in seconds to detect the incident.
	// Equals the difference between `customer_impact_start` and `detected`.
	TimeToDetect *int64 `json:"time_to_detect,omitempty"`
	// The amount of time in seconds to call incident after detection. Equals the difference of `detected` and `created`.
	TimeToInternalResponse *int64 `json:"time_to_internal_response,omitempty"`
	// The amount of time in seconds to resolve customer impact after detecting the issue. Equals the difference between `customer_impact_end` and `detected`.
	TimeToRepair *int64 `json:"time_to_repair,omitempty"`
	// The amount of time in seconds to resolve the incident after it was created. Equals the difference between `created` and `resolved`.
	TimeToResolve *int64 `json:"time_to_resolve,omitempty"`
	// The title of the incident, which summarizes what happened.
	Title string `json:"title"`
	// The incident visibility status.
	Visibility datadog.NullableString `json:"visibility,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponseAttributes instantiates a new IncidentResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponseAttributes(title string) *IncidentResponseAttributes {
	this := IncidentResponseAttributes{}
	this.Title = title
	return &this
}

// NewIncidentResponseAttributesWithDefaults instantiates a new IncidentResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponseAttributesWithDefaults() *IncidentResponseAttributes {
	this := IncidentResponseAttributes{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetArchived() time.Time {
	if o == nil || o.Archived.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Archived.Get()
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetArchivedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Archived.Get(), o.Archived.IsSet()
}

// HasArchived returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasArchived() bool {
	return o != nil && o.Archived.IsSet()
}

// SetArchived gets a reference to the given datadog.NullableTime and assigns it to the Archived field.
func (o *IncidentResponseAttributes) SetArchived(v time.Time) {
	o.Archived.Set(&v)
}

// SetArchivedNil sets the value for Archived to be an explicit nil.
func (o *IncidentResponseAttributes) SetArchivedNil() {
	o.Archived.Set(nil)
}

// UnsetArchived ensures that no value is present for Archived, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetArchived() {
	o.Archived.Unset()
}

// GetCaseId returns the CaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetCaseId() int64 {
	if o == nil || o.CaseId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CaseId.Get()
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetCaseIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseId.Get(), o.CaseId.IsSet()
}

// HasCaseId returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCaseId() bool {
	return o != nil && o.CaseId.IsSet()
}

// SetCaseId gets a reference to the given datadog.NullableInt64 and assigns it to the CaseId field.
func (o *IncidentResponseAttributes) SetCaseId(v int64) {
	o.CaseId.Set(&v)
}

// SetCaseIdNil sets the value for CaseId to be an explicit nil.
func (o *IncidentResponseAttributes) SetCaseIdNil() {
	o.CaseId.Set(nil)
}

// UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetCaseId() {
	o.CaseId.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IncidentResponseAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCustomerImpactDuration returns the CustomerImpactDuration field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetCustomerImpactDuration() int64 {
	if o == nil || o.CustomerImpactDuration == nil {
		var ret int64
		return ret
	}
	return *o.CustomerImpactDuration
}

// GetCustomerImpactDurationOk returns a tuple with the CustomerImpactDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetCustomerImpactDurationOk() (*int64, bool) {
	if o == nil || o.CustomerImpactDuration == nil {
		return nil, false
	}
	return o.CustomerImpactDuration, true
}

// HasCustomerImpactDuration returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCustomerImpactDuration() bool {
	return o != nil && o.CustomerImpactDuration != nil
}

// SetCustomerImpactDuration gets a reference to the given int64 and assigns it to the CustomerImpactDuration field.
func (o *IncidentResponseAttributes) SetCustomerImpactDuration(v int64) {
	o.CustomerImpactDuration = &v
}

// GetCustomerImpactEnd returns the CustomerImpactEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetCustomerImpactEnd() time.Time {
	if o == nil || o.CustomerImpactEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactEnd.Get()
}

// GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetCustomerImpactEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactEnd.Get(), o.CustomerImpactEnd.IsSet()
}

// HasCustomerImpactEnd returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCustomerImpactEnd() bool {
	return o != nil && o.CustomerImpactEnd.IsSet()
}

// SetCustomerImpactEnd gets a reference to the given datadog.NullableTime and assigns it to the CustomerImpactEnd field.
func (o *IncidentResponseAttributes) SetCustomerImpactEnd(v time.Time) {
	o.CustomerImpactEnd.Set(&v)
}

// SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil.
func (o *IncidentResponseAttributes) SetCustomerImpactEndNil() {
	o.CustomerImpactEnd.Set(nil)
}

// UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetCustomerImpactEnd() {
	o.CustomerImpactEnd.Unset()
}

// GetCustomerImpactScope returns the CustomerImpactScope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetCustomerImpactScope() string {
	if o == nil || o.CustomerImpactScope.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerImpactScope.Get()
}

// GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetCustomerImpactScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactScope.Get(), o.CustomerImpactScope.IsSet()
}

// HasCustomerImpactScope returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCustomerImpactScope() bool {
	return o != nil && o.CustomerImpactScope.IsSet()
}

// SetCustomerImpactScope gets a reference to the given datadog.NullableString and assigns it to the CustomerImpactScope field.
func (o *IncidentResponseAttributes) SetCustomerImpactScope(v string) {
	o.CustomerImpactScope.Set(&v)
}

// SetCustomerImpactScopeNil sets the value for CustomerImpactScope to be an explicit nil.
func (o *IncidentResponseAttributes) SetCustomerImpactScopeNil() {
	o.CustomerImpactScope.Set(nil)
}

// UnsetCustomerImpactScope ensures that no value is present for CustomerImpactScope, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetCustomerImpactScope() {
	o.CustomerImpactScope.Unset()
}

// GetCustomerImpactStart returns the CustomerImpactStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetCustomerImpactStart() time.Time {
	if o == nil || o.CustomerImpactStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactStart.Get()
}

// GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetCustomerImpactStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactStart.Get(), o.CustomerImpactStart.IsSet()
}

// HasCustomerImpactStart returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCustomerImpactStart() bool {
	return o != nil && o.CustomerImpactStart.IsSet()
}

// SetCustomerImpactStart gets a reference to the given datadog.NullableTime and assigns it to the CustomerImpactStart field.
func (o *IncidentResponseAttributes) SetCustomerImpactStart(v time.Time) {
	o.CustomerImpactStart.Set(&v)
}

// SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil.
func (o *IncidentResponseAttributes) SetCustomerImpactStartNil() {
	o.CustomerImpactStart.Set(nil)
}

// UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetCustomerImpactStart() {
	o.CustomerImpactStart.Unset()
}

// GetCustomerImpacted returns the CustomerImpacted field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetCustomerImpacted() bool {
	if o == nil || o.CustomerImpacted == nil {
		var ret bool
		return ret
	}
	return *o.CustomerImpacted
}

// GetCustomerImpactedOk returns a tuple with the CustomerImpacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetCustomerImpactedOk() (*bool, bool) {
	if o == nil || o.CustomerImpacted == nil {
		return nil, false
	}
	return o.CustomerImpacted, true
}

// HasCustomerImpacted returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasCustomerImpacted() bool {
	return o != nil && o.CustomerImpacted != nil
}

// SetCustomerImpacted gets a reference to the given bool and assigns it to the CustomerImpacted field.
func (o *IncidentResponseAttributes) SetCustomerImpacted(v bool) {
	o.CustomerImpacted = &v
}

// GetDeclared returns the Declared field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetDeclared() time.Time {
	if o == nil || o.Declared == nil {
		var ret time.Time
		return ret
	}
	return *o.Declared
}

// GetDeclaredOk returns a tuple with the Declared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetDeclaredOk() (*time.Time, bool) {
	if o == nil || o.Declared == nil {
		return nil, false
	}
	return o.Declared, true
}

// HasDeclared returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasDeclared() bool {
	return o != nil && o.Declared != nil
}

// SetDeclared gets a reference to the given time.Time and assigns it to the Declared field.
func (o *IncidentResponseAttributes) SetDeclared(v time.Time) {
	o.Declared = &v
}

// GetDeclaredBy returns the DeclaredBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetDeclaredBy() IncidentNonDatadogCreator {
	if o == nil || o.DeclaredBy.Get() == nil {
		var ret IncidentNonDatadogCreator
		return ret
	}
	return *o.DeclaredBy.Get()
}

// GetDeclaredByOk returns a tuple with the DeclaredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetDeclaredByOk() (*IncidentNonDatadogCreator, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeclaredBy.Get(), o.DeclaredBy.IsSet()
}

// HasDeclaredBy returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasDeclaredBy() bool {
	return o != nil && o.DeclaredBy.IsSet()
}

// SetDeclaredBy gets a reference to the given NullableIncidentNonDatadogCreator and assigns it to the DeclaredBy field.
func (o *IncidentResponseAttributes) SetDeclaredBy(v IncidentNonDatadogCreator) {
	o.DeclaredBy.Set(&v)
}

// SetDeclaredByNil sets the value for DeclaredBy to be an explicit nil.
func (o *IncidentResponseAttributes) SetDeclaredByNil() {
	o.DeclaredBy.Set(nil)
}

// UnsetDeclaredBy ensures that no value is present for DeclaredBy, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetDeclaredBy() {
	o.DeclaredBy.Unset()
}

// GetDeclaredByUuid returns the DeclaredByUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetDeclaredByUuid() string {
	if o == nil || o.DeclaredByUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeclaredByUuid.Get()
}

// GetDeclaredByUuidOk returns a tuple with the DeclaredByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetDeclaredByUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeclaredByUuid.Get(), o.DeclaredByUuid.IsSet()
}

// HasDeclaredByUuid returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasDeclaredByUuid() bool {
	return o != nil && o.DeclaredByUuid.IsSet()
}

// SetDeclaredByUuid gets a reference to the given datadog.NullableString and assigns it to the DeclaredByUuid field.
func (o *IncidentResponseAttributes) SetDeclaredByUuid(v string) {
	o.DeclaredByUuid.Set(&v)
}

// SetDeclaredByUuidNil sets the value for DeclaredByUuid to be an explicit nil.
func (o *IncidentResponseAttributes) SetDeclaredByUuidNil() {
	o.DeclaredByUuid.Set(nil)
}

// UnsetDeclaredByUuid ensures that no value is present for DeclaredByUuid, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetDeclaredByUuid() {
	o.DeclaredByUuid.Unset()
}

// GetDetected returns the Detected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetDetected() time.Time {
	if o == nil || o.Detected.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Detected.Get()
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetDetectedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detected.Get(), o.Detected.IsSet()
}

// HasDetected returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasDetected() bool {
	return o != nil && o.Detected.IsSet()
}

// SetDetected gets a reference to the given datadog.NullableTime and assigns it to the Detected field.
func (o *IncidentResponseAttributes) SetDetected(v time.Time) {
	o.Detected.Set(&v)
}

// SetDetectedNil sets the value for Detected to be an explicit nil.
func (o *IncidentResponseAttributes) SetDetectedNil() {
	o.Detected.Set(nil)
}

// UnsetDetected ensures that no value is present for Detected, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetDetected() {
	o.Detected.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetFields() map[string]IncidentFieldAttributes {
	if o == nil || o.Fields == nil {
		var ret map[string]IncidentFieldAttributes
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]IncidentFieldAttributes and assigns it to the Fields field.
func (o *IncidentResponseAttributes) SetFields(v map[string]IncidentFieldAttributes) {
	o.Fields = v
}

// GetIncidentTypeUuid returns the IncidentTypeUuid field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetIncidentTypeUuid() string {
	if o == nil || o.IncidentTypeUuid == nil {
		var ret string
		return ret
	}
	return *o.IncidentTypeUuid
}

// GetIncidentTypeUuidOk returns a tuple with the IncidentTypeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetIncidentTypeUuidOk() (*string, bool) {
	if o == nil || o.IncidentTypeUuid == nil {
		return nil, false
	}
	return o.IncidentTypeUuid, true
}

// HasIncidentTypeUuid returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasIncidentTypeUuid() bool {
	return o != nil && o.IncidentTypeUuid != nil
}

// SetIncidentTypeUuid gets a reference to the given string and assigns it to the IncidentTypeUuid field.
func (o *IncidentResponseAttributes) SetIncidentTypeUuid(v string) {
	o.IncidentTypeUuid = &v
}

// GetIsTest returns the IsTest field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetIsTest() bool {
	if o == nil || o.IsTest == nil {
		var ret bool
		return ret
	}
	return *o.IsTest
}

// GetIsTestOk returns a tuple with the IsTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetIsTestOk() (*bool, bool) {
	if o == nil || o.IsTest == nil {
		return nil, false
	}
	return o.IsTest, true
}

// HasIsTest returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasIsTest() bool {
	return o != nil && o.IsTest != nil
}

// SetIsTest gets a reference to the given bool and assigns it to the IsTest field.
func (o *IncidentResponseAttributes) SetIsTest(v bool) {
	o.IsTest = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentResponseAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetNonDatadogCreator returns the NonDatadogCreator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetNonDatadogCreator() IncidentNonDatadogCreator {
	if o == nil || o.NonDatadogCreator.Get() == nil {
		var ret IncidentNonDatadogCreator
		return ret
	}
	return *o.NonDatadogCreator.Get()
}

// GetNonDatadogCreatorOk returns a tuple with the NonDatadogCreator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetNonDatadogCreatorOk() (*IncidentNonDatadogCreator, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonDatadogCreator.Get(), o.NonDatadogCreator.IsSet()
}

// HasNonDatadogCreator returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasNonDatadogCreator() bool {
	return o != nil && o.NonDatadogCreator.IsSet()
}

// SetNonDatadogCreator gets a reference to the given NullableIncidentNonDatadogCreator and assigns it to the NonDatadogCreator field.
func (o *IncidentResponseAttributes) SetNonDatadogCreator(v IncidentNonDatadogCreator) {
	o.NonDatadogCreator.Set(&v)
}

// SetNonDatadogCreatorNil sets the value for NonDatadogCreator to be an explicit nil.
func (o *IncidentResponseAttributes) SetNonDatadogCreatorNil() {
	o.NonDatadogCreator.Set(nil)
}

// UnsetNonDatadogCreator ensures that no value is present for NonDatadogCreator, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetNonDatadogCreator() {
	o.NonDatadogCreator.Unset()
}

// GetNotificationHandles returns the NotificationHandles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetNotificationHandles() []IncidentNotificationHandle {
	if o == nil {
		var ret []IncidentNotificationHandle
		return ret
	}
	return o.NotificationHandles
}

// GetNotificationHandlesOk returns a tuple with the NotificationHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool) {
	if o == nil || o.NotificationHandles == nil {
		return nil, false
	}
	return &o.NotificationHandles, true
}

// HasNotificationHandles returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasNotificationHandles() bool {
	return o != nil && o.NotificationHandles != nil
}

// SetNotificationHandles gets a reference to the given []IncidentNotificationHandle and assigns it to the NotificationHandles field.
func (o *IncidentResponseAttributes) SetNotificationHandles(v []IncidentNotificationHandle) {
	o.NotificationHandles = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetPublicId() int64 {
	if o == nil || o.PublicId == nil {
		var ret int64
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetPublicIdOk() (*int64, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given int64 and assigns it to the PublicId field.
func (o *IncidentResponseAttributes) SetPublicId(v int64) {
	o.PublicId = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasResolved() bool {
	return o != nil && o.Resolved.IsSet()
}

// SetResolved gets a reference to the given datadog.NullableTime and assigns it to the Resolved field.
func (o *IncidentResponseAttributes) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil.
func (o *IncidentResponseAttributes) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetResolved() {
	o.Resolved.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetSeverity() IncidentSeverity {
	if o == nil || o.Severity == nil {
		var ret IncidentSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetSeverityOk() (*IncidentSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given IncidentSeverity and assigns it to the Severity field.
func (o *IncidentResponseAttributes) SetSeverity(v IncidentSeverity) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given datadog.NullableString and assigns it to the State field.
func (o *IncidentResponseAttributes) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *IncidentResponseAttributes) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetState() {
	o.State.Unset()
}

// GetTimeToDetect returns the TimeToDetect field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetTimeToDetect() int64 {
	if o == nil || o.TimeToDetect == nil {
		var ret int64
		return ret
	}
	return *o.TimeToDetect
}

// GetTimeToDetectOk returns a tuple with the TimeToDetect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetTimeToDetectOk() (*int64, bool) {
	if o == nil || o.TimeToDetect == nil {
		return nil, false
	}
	return o.TimeToDetect, true
}

// HasTimeToDetect returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasTimeToDetect() bool {
	return o != nil && o.TimeToDetect != nil
}

// SetTimeToDetect gets a reference to the given int64 and assigns it to the TimeToDetect field.
func (o *IncidentResponseAttributes) SetTimeToDetect(v int64) {
	o.TimeToDetect = &v
}

// GetTimeToInternalResponse returns the TimeToInternalResponse field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetTimeToInternalResponse() int64 {
	if o == nil || o.TimeToInternalResponse == nil {
		var ret int64
		return ret
	}
	return *o.TimeToInternalResponse
}

// GetTimeToInternalResponseOk returns a tuple with the TimeToInternalResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetTimeToInternalResponseOk() (*int64, bool) {
	if o == nil || o.TimeToInternalResponse == nil {
		return nil, false
	}
	return o.TimeToInternalResponse, true
}

// HasTimeToInternalResponse returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasTimeToInternalResponse() bool {
	return o != nil && o.TimeToInternalResponse != nil
}

// SetTimeToInternalResponse gets a reference to the given int64 and assigns it to the TimeToInternalResponse field.
func (o *IncidentResponseAttributes) SetTimeToInternalResponse(v int64) {
	o.TimeToInternalResponse = &v
}

// GetTimeToRepair returns the TimeToRepair field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetTimeToRepair() int64 {
	if o == nil || o.TimeToRepair == nil {
		var ret int64
		return ret
	}
	return *o.TimeToRepair
}

// GetTimeToRepairOk returns a tuple with the TimeToRepair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetTimeToRepairOk() (*int64, bool) {
	if o == nil || o.TimeToRepair == nil {
		return nil, false
	}
	return o.TimeToRepair, true
}

// HasTimeToRepair returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasTimeToRepair() bool {
	return o != nil && o.TimeToRepair != nil
}

// SetTimeToRepair gets a reference to the given int64 and assigns it to the TimeToRepair field.
func (o *IncidentResponseAttributes) SetTimeToRepair(v int64) {
	o.TimeToRepair = &v
}

// GetTimeToResolve returns the TimeToResolve field value if set, zero value otherwise.
func (o *IncidentResponseAttributes) GetTimeToResolve() int64 {
	if o == nil || o.TimeToResolve == nil {
		var ret int64
		return ret
	}
	return *o.TimeToResolve
}

// GetTimeToResolveOk returns a tuple with the TimeToResolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetTimeToResolveOk() (*int64, bool) {
	if o == nil || o.TimeToResolve == nil {
		return nil, false
	}
	return o.TimeToResolve, true
}

// HasTimeToResolve returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasTimeToResolve() bool {
	return o != nil && o.TimeToResolve != nil
}

// SetTimeToResolve gets a reference to the given int64 and assigns it to the TimeToResolve field.
func (o *IncidentResponseAttributes) SetTimeToResolve(v int64) {
	o.TimeToResolve = &v
}

// GetTitle returns the Title field value.
func (o *IncidentResponseAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentResponseAttributes) SetTitle(v string) {
	o.Title = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponseAttributes) GetVisibility() string {
	if o == nil || o.Visibility.Get() == nil {
		var ret string
		return ret
	}
	return *o.Visibility.Get()
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponseAttributes) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visibility.Get(), o.Visibility.IsSet()
}

// HasVisibility returns a boolean if a field has been set.
func (o *IncidentResponseAttributes) HasVisibility() bool {
	return o != nil && o.Visibility.IsSet()
}

// SetVisibility gets a reference to the given datadog.NullableString and assigns it to the Visibility field.
func (o *IncidentResponseAttributes) SetVisibility(v string) {
	o.Visibility.Set(&v)
}

// SetVisibilityNil sets the value for Visibility to be an explicit nil.
func (o *IncidentResponseAttributes) SetVisibilityNil() {
	o.Visibility.Set(nil)
}

// UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil.
func (o *IncidentResponseAttributes) UnsetVisibility() {
	o.Visibility.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Archived.IsSet() {
		toSerialize["archived"] = o.Archived.Get()
	}
	if o.CaseId.IsSet() {
		toSerialize["case_id"] = o.CaseId.Get()
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CustomerImpactDuration != nil {
		toSerialize["customer_impact_duration"] = o.CustomerImpactDuration
	}
	if o.CustomerImpactEnd.IsSet() {
		toSerialize["customer_impact_end"] = o.CustomerImpactEnd.Get()
	}
	if o.CustomerImpactScope.IsSet() {
		toSerialize["customer_impact_scope"] = o.CustomerImpactScope.Get()
	}
	if o.CustomerImpactStart.IsSet() {
		toSerialize["customer_impact_start"] = o.CustomerImpactStart.Get()
	}
	if o.CustomerImpacted != nil {
		toSerialize["customer_impacted"] = o.CustomerImpacted
	}
	if o.Declared != nil {
		if o.Declared.Nanosecond() == 0 {
			toSerialize["declared"] = o.Declared.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["declared"] = o.Declared.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DeclaredBy.IsSet() {
		toSerialize["declared_by"] = o.DeclaredBy.Get()
	}
	if o.DeclaredByUuid.IsSet() {
		toSerialize["declared_by_uuid"] = o.DeclaredByUuid.Get()
	}
	if o.Detected.IsSet() {
		toSerialize["detected"] = o.Detected.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.IncidentTypeUuid != nil {
		toSerialize["incident_type_uuid"] = o.IncidentTypeUuid
	}
	if o.IsTest != nil {
		toSerialize["is_test"] = o.IsTest
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.NonDatadogCreator.IsSet() {
		toSerialize["non_datadog_creator"] = o.NonDatadogCreator.Get()
	}
	if o.NotificationHandles != nil {
		toSerialize["notification_handles"] = o.NotificationHandles
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.TimeToDetect != nil {
		toSerialize["time_to_detect"] = o.TimeToDetect
	}
	if o.TimeToInternalResponse != nil {
		toSerialize["time_to_internal_response"] = o.TimeToInternalResponse
	}
	if o.TimeToRepair != nil {
		toSerialize["time_to_repair"] = o.TimeToRepair
	}
	if o.TimeToResolve != nil {
		toSerialize["time_to_resolve"] = o.TimeToResolve
	}
	toSerialize["title"] = o.Title
	if o.Visibility.IsSet() {
		toSerialize["visibility"] = o.Visibility.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Archived               datadog.NullableTime               `json:"archived,omitempty"`
		CaseId                 datadog.NullableInt64              `json:"case_id,omitempty"`
		Created                *time.Time                         `json:"created,omitempty"`
		CustomerImpactDuration *int64                             `json:"customer_impact_duration,omitempty"`
		CustomerImpactEnd      datadog.NullableTime               `json:"customer_impact_end,omitempty"`
		CustomerImpactScope    datadog.NullableString             `json:"customer_impact_scope,omitempty"`
		CustomerImpactStart    datadog.NullableTime               `json:"customer_impact_start,omitempty"`
		CustomerImpacted       *bool                              `json:"customer_impacted,omitempty"`
		Declared               *time.Time                         `json:"declared,omitempty"`
		DeclaredBy             NullableIncidentNonDatadogCreator  `json:"declared_by,omitempty"`
		DeclaredByUuid         datadog.NullableString             `json:"declared_by_uuid,omitempty"`
		Detected               datadog.NullableTime               `json:"detected,omitempty"`
		Fields                 map[string]IncidentFieldAttributes `json:"fields,omitempty"`
		IncidentTypeUuid       *string                            `json:"incident_type_uuid,omitempty"`
		IsTest                 *bool                              `json:"is_test,omitempty"`
		Modified               *time.Time                         `json:"modified,omitempty"`
		NonDatadogCreator      NullableIncidentNonDatadogCreator  `json:"non_datadog_creator,omitempty"`
		NotificationHandles    []IncidentNotificationHandle       `json:"notification_handles,omitempty"`
		PublicId               *int64                             `json:"public_id,omitempty"`
		Resolved               datadog.NullableTime               `json:"resolved,omitempty"`
		Severity               *IncidentSeverity                  `json:"severity,omitempty"`
		State                  datadog.NullableString             `json:"state,omitempty"`
		TimeToDetect           *int64                             `json:"time_to_detect,omitempty"`
		TimeToInternalResponse *int64                             `json:"time_to_internal_response,omitempty"`
		TimeToRepair           *int64                             `json:"time_to_repair,omitempty"`
		TimeToResolve          *int64                             `json:"time_to_resolve,omitempty"`
		Title                  *string                            `json:"title"`
		Visibility             datadog.NullableString             `json:"visibility,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived", "case_id", "created", "customer_impact_duration", "customer_impact_end", "customer_impact_scope", "customer_impact_start", "customer_impacted", "declared", "declared_by", "declared_by_uuid", "detected", "fields", "incident_type_uuid", "is_test", "modified", "non_datadog_creator", "notification_handles", "public_id", "resolved", "severity", "state", "time_to_detect", "time_to_internal_response", "time_to_repair", "time_to_resolve", "title", "visibility"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Archived = all.Archived
	o.CaseId = all.CaseId
	o.Created = all.Created
	o.CustomerImpactDuration = all.CustomerImpactDuration
	o.CustomerImpactEnd = all.CustomerImpactEnd
	o.CustomerImpactScope = all.CustomerImpactScope
	o.CustomerImpactStart = all.CustomerImpactStart
	o.CustomerImpacted = all.CustomerImpacted
	o.Declared = all.Declared
	o.DeclaredBy = all.DeclaredBy
	o.DeclaredByUuid = all.DeclaredByUuid
	o.Detected = all.Detected
	o.Fields = all.Fields
	o.IncidentTypeUuid = all.IncidentTypeUuid
	o.IsTest = all.IsTest
	o.Modified = all.Modified
	o.NonDatadogCreator = all.NonDatadogCreator
	o.NotificationHandles = all.NotificationHandles
	o.PublicId = all.PublicId
	o.Resolved = all.Resolved
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.State = all.State
	o.TimeToDetect = all.TimeToDetect
	o.TimeToInternalResponse = all.TimeToInternalResponse
	o.TimeToRepair = all.TimeToRepair
	o.TimeToResolve = all.TimeToResolve
	o.Title = *all.Title
	o.Visibility = all.Visibility

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
