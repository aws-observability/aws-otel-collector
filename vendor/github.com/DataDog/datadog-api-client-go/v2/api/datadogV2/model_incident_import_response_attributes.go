// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportResponseAttributes The incident's attributes from an import response.
type IncidentImportResponseAttributes struct {
	// Timestamp when the incident was archived.
	Archived datadog.NullableTime `json:"archived,omitempty"`
	// The incident case ID.
	CaseId datadog.NullableInt64 `json:"case_id,omitempty"`
	// Timestamp when the incident was created.
	Created *time.Time `json:"created,omitempty"`
	// UUID of the user who created the incident.
	CreatedByUuid datadog.NullableString `json:"created_by_uuid,omitempty"`
	// A unique key used to ensure idempotent incident creation.
	CreationIdempotencyKey datadog.NullableString `json:"creation_idempotency_key,omitempty"`
	// Timestamp when customers were no longer impacted by the incident.
	CustomerImpactEnd datadog.NullableTime `json:"customer_impact_end,omitempty"`
	// A summary of the impact customers experienced during the incident.
	CustomerImpactScope datadog.NullableString `json:"customer_impact_scope,omitempty"`
	// Timestamp when customers began to be impacted by the incident.
	CustomerImpactStart datadog.NullableTime `json:"customer_impact_start,omitempty"`
	// Timestamp when the incident was declared.
	Declared datadog.NullableTime `json:"declared,omitempty"`
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
	// UUID of the user who last modified the incident.
	LastModifiedByUuid datadog.NullableString `json:"last_modified_by_uuid,omitempty"`
	// Timestamp when the incident was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Incident's non Datadog creator.
	NonDatadogCreator NullableIncidentNonDatadogCreator `json:"non_datadog_creator,omitempty"`
	// Notification handles that are notified of the incident during update.
	NotificationHandles []IncidentNotificationHandle `json:"notification_handles,omitempty"`
	// The monotonically increasing integer ID for the incident.
	PublicId *int64 `json:"public_id,omitempty"`
	// Timestamp when the incident's state was last changed from active or stable to resolved or completed.
	Resolved datadog.NullableTime `json:"resolved,omitempty"`
	// The incident severity.
	Severity *IncidentSeverity `json:"severity,omitempty"`
	// The state of the incident.
	State datadog.NullableString `json:"state,omitempty"`
	// The title of the incident that summarizes what happened.
	Title string `json:"title"`
	// The incident visibility status.
	Visibility datadog.NullableString `json:"visibility,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImportResponseAttributes instantiates a new IncidentImportResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImportResponseAttributes(title string) *IncidentImportResponseAttributes {
	this := IncidentImportResponseAttributes{}
	this.Title = title
	return &this
}

// NewIncidentImportResponseAttributesWithDefaults instantiates a new IncidentImportResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImportResponseAttributesWithDefaults() *IncidentImportResponseAttributes {
	this := IncidentImportResponseAttributes{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetArchived() time.Time {
	if o == nil || o.Archived.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Archived.Get()
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetArchivedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Archived.Get(), o.Archived.IsSet()
}

// HasArchived returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasArchived() bool {
	return o != nil && o.Archived.IsSet()
}

// SetArchived gets a reference to the given datadog.NullableTime and assigns it to the Archived field.
func (o *IncidentImportResponseAttributes) SetArchived(v time.Time) {
	o.Archived.Set(&v)
}

// SetArchivedNil sets the value for Archived to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetArchivedNil() {
	o.Archived.Set(nil)
}

// UnsetArchived ensures that no value is present for Archived, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetArchived() {
	o.Archived.Unset()
}

// GetCaseId returns the CaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetCaseId() int64 {
	if o == nil || o.CaseId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CaseId.Get()
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetCaseIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseId.Get(), o.CaseId.IsSet()
}

// HasCaseId returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCaseId() bool {
	return o != nil && o.CaseId.IsSet()
}

// SetCaseId gets a reference to the given datadog.NullableInt64 and assigns it to the CaseId field.
func (o *IncidentImportResponseAttributes) SetCaseId(v int64) {
	o.CaseId.Set(&v)
}

// SetCaseIdNil sets the value for CaseId to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetCaseIdNil() {
	o.CaseId.Set(nil)
}

// UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetCaseId() {
	o.CaseId.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IncidentImportResponseAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedByUuid returns the CreatedByUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetCreatedByUuid() string {
	if o == nil || o.CreatedByUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUuid.Get()
}

// GetCreatedByUuidOk returns a tuple with the CreatedByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetCreatedByUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedByUuid.Get(), o.CreatedByUuid.IsSet()
}

// HasCreatedByUuid returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCreatedByUuid() bool {
	return o != nil && o.CreatedByUuid.IsSet()
}

// SetCreatedByUuid gets a reference to the given datadog.NullableString and assigns it to the CreatedByUuid field.
func (o *IncidentImportResponseAttributes) SetCreatedByUuid(v string) {
	o.CreatedByUuid.Set(&v)
}

// SetCreatedByUuidNil sets the value for CreatedByUuid to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetCreatedByUuidNil() {
	o.CreatedByUuid.Set(nil)
}

// UnsetCreatedByUuid ensures that no value is present for CreatedByUuid, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetCreatedByUuid() {
	o.CreatedByUuid.Unset()
}

// GetCreationIdempotencyKey returns the CreationIdempotencyKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetCreationIdempotencyKey() string {
	if o == nil || o.CreationIdempotencyKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreationIdempotencyKey.Get()
}

// GetCreationIdempotencyKeyOk returns a tuple with the CreationIdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetCreationIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationIdempotencyKey.Get(), o.CreationIdempotencyKey.IsSet()
}

// HasCreationIdempotencyKey returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCreationIdempotencyKey() bool {
	return o != nil && o.CreationIdempotencyKey.IsSet()
}

// SetCreationIdempotencyKey gets a reference to the given datadog.NullableString and assigns it to the CreationIdempotencyKey field.
func (o *IncidentImportResponseAttributes) SetCreationIdempotencyKey(v string) {
	o.CreationIdempotencyKey.Set(&v)
}

// SetCreationIdempotencyKeyNil sets the value for CreationIdempotencyKey to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetCreationIdempotencyKeyNil() {
	o.CreationIdempotencyKey.Set(nil)
}

// UnsetCreationIdempotencyKey ensures that no value is present for CreationIdempotencyKey, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetCreationIdempotencyKey() {
	o.CreationIdempotencyKey.Unset()
}

// GetCustomerImpactEnd returns the CustomerImpactEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetCustomerImpactEnd() time.Time {
	if o == nil || o.CustomerImpactEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactEnd.Get()
}

// GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetCustomerImpactEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactEnd.Get(), o.CustomerImpactEnd.IsSet()
}

// HasCustomerImpactEnd returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCustomerImpactEnd() bool {
	return o != nil && o.CustomerImpactEnd.IsSet()
}

// SetCustomerImpactEnd gets a reference to the given datadog.NullableTime and assigns it to the CustomerImpactEnd field.
func (o *IncidentImportResponseAttributes) SetCustomerImpactEnd(v time.Time) {
	o.CustomerImpactEnd.Set(&v)
}

// SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetCustomerImpactEndNil() {
	o.CustomerImpactEnd.Set(nil)
}

// UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetCustomerImpactEnd() {
	o.CustomerImpactEnd.Unset()
}

// GetCustomerImpactScope returns the CustomerImpactScope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetCustomerImpactScope() string {
	if o == nil || o.CustomerImpactScope.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerImpactScope.Get()
}

// GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetCustomerImpactScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactScope.Get(), o.CustomerImpactScope.IsSet()
}

// HasCustomerImpactScope returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCustomerImpactScope() bool {
	return o != nil && o.CustomerImpactScope.IsSet()
}

// SetCustomerImpactScope gets a reference to the given datadog.NullableString and assigns it to the CustomerImpactScope field.
func (o *IncidentImportResponseAttributes) SetCustomerImpactScope(v string) {
	o.CustomerImpactScope.Set(&v)
}

// SetCustomerImpactScopeNil sets the value for CustomerImpactScope to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetCustomerImpactScopeNil() {
	o.CustomerImpactScope.Set(nil)
}

// UnsetCustomerImpactScope ensures that no value is present for CustomerImpactScope, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetCustomerImpactScope() {
	o.CustomerImpactScope.Unset()
}

// GetCustomerImpactStart returns the CustomerImpactStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetCustomerImpactStart() time.Time {
	if o == nil || o.CustomerImpactStart.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CustomerImpactStart.Get()
}

// GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetCustomerImpactStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerImpactStart.Get(), o.CustomerImpactStart.IsSet()
}

// HasCustomerImpactStart returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasCustomerImpactStart() bool {
	return o != nil && o.CustomerImpactStart.IsSet()
}

// SetCustomerImpactStart gets a reference to the given datadog.NullableTime and assigns it to the CustomerImpactStart field.
func (o *IncidentImportResponseAttributes) SetCustomerImpactStart(v time.Time) {
	o.CustomerImpactStart.Set(&v)
}

// SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetCustomerImpactStartNil() {
	o.CustomerImpactStart.Set(nil)
}

// UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetCustomerImpactStart() {
	o.CustomerImpactStart.Unset()
}

// GetDeclared returns the Declared field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetDeclared() time.Time {
	if o == nil || o.Declared.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Declared.Get()
}

// GetDeclaredOk returns a tuple with the Declared field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetDeclaredOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Declared.Get(), o.Declared.IsSet()
}

// HasDeclared returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasDeclared() bool {
	return o != nil && o.Declared.IsSet()
}

// SetDeclared gets a reference to the given datadog.NullableTime and assigns it to the Declared field.
func (o *IncidentImportResponseAttributes) SetDeclared(v time.Time) {
	o.Declared.Set(&v)
}

// SetDeclaredNil sets the value for Declared to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetDeclaredNil() {
	o.Declared.Set(nil)
}

// UnsetDeclared ensures that no value is present for Declared, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetDeclared() {
	o.Declared.Unset()
}

// GetDeclaredByUuid returns the DeclaredByUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetDeclaredByUuid() string {
	if o == nil || o.DeclaredByUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeclaredByUuid.Get()
}

// GetDeclaredByUuidOk returns a tuple with the DeclaredByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetDeclaredByUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeclaredByUuid.Get(), o.DeclaredByUuid.IsSet()
}

// HasDeclaredByUuid returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasDeclaredByUuid() bool {
	return o != nil && o.DeclaredByUuid.IsSet()
}

// SetDeclaredByUuid gets a reference to the given datadog.NullableString and assigns it to the DeclaredByUuid field.
func (o *IncidentImportResponseAttributes) SetDeclaredByUuid(v string) {
	o.DeclaredByUuid.Set(&v)
}

// SetDeclaredByUuidNil sets the value for DeclaredByUuid to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetDeclaredByUuidNil() {
	o.DeclaredByUuid.Set(nil)
}

// UnsetDeclaredByUuid ensures that no value is present for DeclaredByUuid, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetDeclaredByUuid() {
	o.DeclaredByUuid.Unset()
}

// GetDetected returns the Detected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetDetected() time.Time {
	if o == nil || o.Detected.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Detected.Get()
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetDetectedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detected.Get(), o.Detected.IsSet()
}

// HasDetected returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasDetected() bool {
	return o != nil && o.Detected.IsSet()
}

// SetDetected gets a reference to the given datadog.NullableTime and assigns it to the Detected field.
func (o *IncidentImportResponseAttributes) SetDetected(v time.Time) {
	o.Detected.Set(&v)
}

// SetDetectedNil sets the value for Detected to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetDetectedNil() {
	o.Detected.Set(nil)
}

// UnsetDetected ensures that no value is present for Detected, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetDetected() {
	o.Detected.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetFields() map[string]IncidentFieldAttributes {
	if o == nil || o.Fields == nil {
		var ret map[string]IncidentFieldAttributes
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetFieldsOk() (*map[string]IncidentFieldAttributes, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]IncidentFieldAttributes and assigns it to the Fields field.
func (o *IncidentImportResponseAttributes) SetFields(v map[string]IncidentFieldAttributes) {
	o.Fields = v
}

// GetIncidentTypeUuid returns the IncidentTypeUuid field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetIncidentTypeUuid() string {
	if o == nil || o.IncidentTypeUuid == nil {
		var ret string
		return ret
	}
	return *o.IncidentTypeUuid
}

// GetIncidentTypeUuidOk returns a tuple with the IncidentTypeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetIncidentTypeUuidOk() (*string, bool) {
	if o == nil || o.IncidentTypeUuid == nil {
		return nil, false
	}
	return o.IncidentTypeUuid, true
}

// HasIncidentTypeUuid returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasIncidentTypeUuid() bool {
	return o != nil && o.IncidentTypeUuid != nil
}

// SetIncidentTypeUuid gets a reference to the given string and assigns it to the IncidentTypeUuid field.
func (o *IncidentImportResponseAttributes) SetIncidentTypeUuid(v string) {
	o.IncidentTypeUuid = &v
}

// GetIsTest returns the IsTest field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetIsTest() bool {
	if o == nil || o.IsTest == nil {
		var ret bool
		return ret
	}
	return *o.IsTest
}

// GetIsTestOk returns a tuple with the IsTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetIsTestOk() (*bool, bool) {
	if o == nil || o.IsTest == nil {
		return nil, false
	}
	return o.IsTest, true
}

// HasIsTest returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasIsTest() bool {
	return o != nil && o.IsTest != nil
}

// SetIsTest gets a reference to the given bool and assigns it to the IsTest field.
func (o *IncidentImportResponseAttributes) SetIsTest(v bool) {
	o.IsTest = &v
}

// GetLastModifiedByUuid returns the LastModifiedByUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetLastModifiedByUuid() string {
	if o == nil || o.LastModifiedByUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedByUuid.Get()
}

// GetLastModifiedByUuidOk returns a tuple with the LastModifiedByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetLastModifiedByUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedByUuid.Get(), o.LastModifiedByUuid.IsSet()
}

// HasLastModifiedByUuid returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasLastModifiedByUuid() bool {
	return o != nil && o.LastModifiedByUuid.IsSet()
}

// SetLastModifiedByUuid gets a reference to the given datadog.NullableString and assigns it to the LastModifiedByUuid field.
func (o *IncidentImportResponseAttributes) SetLastModifiedByUuid(v string) {
	o.LastModifiedByUuid.Set(&v)
}

// SetLastModifiedByUuidNil sets the value for LastModifiedByUuid to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetLastModifiedByUuidNil() {
	o.LastModifiedByUuid.Set(nil)
}

// UnsetLastModifiedByUuid ensures that no value is present for LastModifiedByUuid, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetLastModifiedByUuid() {
	o.LastModifiedByUuid.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentImportResponseAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetNonDatadogCreator returns the NonDatadogCreator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetNonDatadogCreator() IncidentNonDatadogCreator {
	if o == nil || o.NonDatadogCreator.Get() == nil {
		var ret IncidentNonDatadogCreator
		return ret
	}
	return *o.NonDatadogCreator.Get()
}

// GetNonDatadogCreatorOk returns a tuple with the NonDatadogCreator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetNonDatadogCreatorOk() (*IncidentNonDatadogCreator, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonDatadogCreator.Get(), o.NonDatadogCreator.IsSet()
}

// HasNonDatadogCreator returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasNonDatadogCreator() bool {
	return o != nil && o.NonDatadogCreator.IsSet()
}

// SetNonDatadogCreator gets a reference to the given NullableIncidentNonDatadogCreator and assigns it to the NonDatadogCreator field.
func (o *IncidentImportResponseAttributes) SetNonDatadogCreator(v IncidentNonDatadogCreator) {
	o.NonDatadogCreator.Set(&v)
}

// SetNonDatadogCreatorNil sets the value for NonDatadogCreator to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetNonDatadogCreatorNil() {
	o.NonDatadogCreator.Set(nil)
}

// UnsetNonDatadogCreator ensures that no value is present for NonDatadogCreator, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetNonDatadogCreator() {
	o.NonDatadogCreator.Unset()
}

// GetNotificationHandles returns the NotificationHandles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetNotificationHandles() []IncidentNotificationHandle {
	if o == nil {
		var ret []IncidentNotificationHandle
		return ret
	}
	return o.NotificationHandles
}

// GetNotificationHandlesOk returns a tuple with the NotificationHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetNotificationHandlesOk() (*[]IncidentNotificationHandle, bool) {
	if o == nil || o.NotificationHandles == nil {
		return nil, false
	}
	return &o.NotificationHandles, true
}

// HasNotificationHandles returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasNotificationHandles() bool {
	return o != nil && o.NotificationHandles != nil
}

// SetNotificationHandles gets a reference to the given []IncidentNotificationHandle and assigns it to the NotificationHandles field.
func (o *IncidentImportResponseAttributes) SetNotificationHandles(v []IncidentNotificationHandle) {
	o.NotificationHandles = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetPublicId() int64 {
	if o == nil || o.PublicId == nil {
		var ret int64
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetPublicIdOk() (*int64, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given int64 and assigns it to the PublicId field.
func (o *IncidentImportResponseAttributes) SetPublicId(v int64) {
	o.PublicId = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetResolved() time.Time {
	if o == nil || o.Resolved.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasResolved() bool {
	return o != nil && o.Resolved.IsSet()
}

// SetResolved gets a reference to the given datadog.NullableTime and assigns it to the Resolved field.
func (o *IncidentImportResponseAttributes) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetResolved() {
	o.Resolved.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IncidentImportResponseAttributes) GetSeverity() IncidentSeverity {
	if o == nil || o.Severity == nil {
		var ret IncidentSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetSeverityOk() (*IncidentSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given IncidentSeverity and assigns it to the Severity field.
func (o *IncidentImportResponseAttributes) SetSeverity(v IncidentSeverity) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given datadog.NullableString and assigns it to the State field.
func (o *IncidentImportResponseAttributes) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetState() {
	o.State.Unset()
}

// GetTitle returns the Title field value.
func (o *IncidentImportResponseAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentImportResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentImportResponseAttributes) SetTitle(v string) {
	o.Title = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImportResponseAttributes) GetVisibility() string {
	if o == nil || o.Visibility.Get() == nil {
		var ret string
		return ret
	}
	return *o.Visibility.Get()
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImportResponseAttributes) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visibility.Get(), o.Visibility.IsSet()
}

// HasVisibility returns a boolean if a field has been set.
func (o *IncidentImportResponseAttributes) HasVisibility() bool {
	return o != nil && o.Visibility.IsSet()
}

// SetVisibility gets a reference to the given datadog.NullableString and assigns it to the Visibility field.
func (o *IncidentImportResponseAttributes) SetVisibility(v string) {
	o.Visibility.Set(&v)
}

// SetVisibilityNil sets the value for Visibility to be an explicit nil.
func (o *IncidentImportResponseAttributes) SetVisibilityNil() {
	o.Visibility.Set(nil)
}

// UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil.
func (o *IncidentImportResponseAttributes) UnsetVisibility() {
	o.Visibility.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImportResponseAttributes) MarshalJSON() ([]byte, error) {
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
	if o.CreatedByUuid.IsSet() {
		toSerialize["created_by_uuid"] = o.CreatedByUuid.Get()
	}
	if o.CreationIdempotencyKey.IsSet() {
		toSerialize["creation_idempotency_key"] = o.CreationIdempotencyKey.Get()
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
	if o.Declared.IsSet() {
		toSerialize["declared"] = o.Declared.Get()
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
	if o.LastModifiedByUuid.IsSet() {
		toSerialize["last_modified_by_uuid"] = o.LastModifiedByUuid.Get()
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
func (o *IncidentImportResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Archived               datadog.NullableTime               `json:"archived,omitempty"`
		CaseId                 datadog.NullableInt64              `json:"case_id,omitempty"`
		Created                *time.Time                         `json:"created,omitempty"`
		CreatedByUuid          datadog.NullableString             `json:"created_by_uuid,omitempty"`
		CreationIdempotencyKey datadog.NullableString             `json:"creation_idempotency_key,omitempty"`
		CustomerImpactEnd      datadog.NullableTime               `json:"customer_impact_end,omitempty"`
		CustomerImpactScope    datadog.NullableString             `json:"customer_impact_scope,omitempty"`
		CustomerImpactStart    datadog.NullableTime               `json:"customer_impact_start,omitempty"`
		Declared               datadog.NullableTime               `json:"declared,omitempty"`
		DeclaredByUuid         datadog.NullableString             `json:"declared_by_uuid,omitempty"`
		Detected               datadog.NullableTime               `json:"detected,omitempty"`
		Fields                 map[string]IncidentFieldAttributes `json:"fields,omitempty"`
		IncidentTypeUuid       *string                            `json:"incident_type_uuid,omitempty"`
		IsTest                 *bool                              `json:"is_test,omitempty"`
		LastModifiedByUuid     datadog.NullableString             `json:"last_modified_by_uuid,omitempty"`
		Modified               *time.Time                         `json:"modified,omitempty"`
		NonDatadogCreator      NullableIncidentNonDatadogCreator  `json:"non_datadog_creator,omitempty"`
		NotificationHandles    []IncidentNotificationHandle       `json:"notification_handles,omitempty"`
		PublicId               *int64                             `json:"public_id,omitempty"`
		Resolved               datadog.NullableTime               `json:"resolved,omitempty"`
		Severity               *IncidentSeverity                  `json:"severity,omitempty"`
		State                  datadog.NullableString             `json:"state,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"archived", "case_id", "created", "created_by_uuid", "creation_idempotency_key", "customer_impact_end", "customer_impact_scope", "customer_impact_start", "declared", "declared_by_uuid", "detected", "fields", "incident_type_uuid", "is_test", "last_modified_by_uuid", "modified", "non_datadog_creator", "notification_handles", "public_id", "resolved", "severity", "state", "title", "visibility"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Archived = all.Archived
	o.CaseId = all.CaseId
	o.Created = all.Created
	o.CreatedByUuid = all.CreatedByUuid
	o.CreationIdempotencyKey = all.CreationIdempotencyKey
	o.CustomerImpactEnd = all.CustomerImpactEnd
	o.CustomerImpactScope = all.CustomerImpactScope
	o.CustomerImpactStart = all.CustomerImpactStart
	o.Declared = all.Declared
	o.DeclaredByUuid = all.DeclaredByUuid
	o.Detected = all.Detected
	o.Fields = all.Fields
	o.IncidentTypeUuid = all.IncidentTypeUuid
	o.IsTest = all.IsTest
	o.LastModifiedByUuid = all.LastModifiedByUuid
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
