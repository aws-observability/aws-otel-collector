// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestResponseAttributes Attributes of a change request response.
type ChangeRequestResponseAttributes struct {
	// Timestamp of when the change request was archived.
	ArchivedAt datadog.NullableTime `json:"archived_at,omitempty"`
	// Custom attributes of the change request as key-value pairs.
	Attributes map[string][]string `json:"attributes"`
	// The UUID of the linked incident.
	ChangeRequestLinkedIncidentUuid string `json:"change_request_linked_incident_uuid"`
	// The maintenance window query for the change request.
	ChangeRequestMaintenanceWindowQuery string `json:"change_request_maintenance_window_query"`
	// The plan associated with the change request.
	ChangeRequestPlan string `json:"change_request_plan"`
	// The risk level of the change request.
	ChangeRequestRisk ChangeRequestRiskLevel `json:"change_request_risk"`
	// The type of the change request.
	ChangeRequestType ChangeRequestChangeType `json:"change_request_type"`
	// Timestamp of when the change request was closed.
	ClosedAt datadog.NullableTime `json:"closed_at,omitempty"`
	// Timestamp of when the change request was created.
	CreatedAt time.Time `json:"created_at"`
	// The source from which the change request was created.
	CreationSource string `json:"creation_source"`
	// The description of the change request.
	Description string `json:"description"`
	// The planned end date of the change request.
	EndDate *time.Time `json:"end_date,omitempty"`
	// The human-readable key of the change request.
	Key string `json:"key"`
	// Timestamp of when the change request was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The notebook ID associated with the change request plan.
	PlanNotebookId int64 `json:"plan_notebook_id"`
	// The priority of the change request.
	Priority string `json:"priority"`
	// The project UUID associated with the change request.
	ProjectId string `json:"project_id"`
	// The planned start date of the change request.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The current status of the change request.
	Status string `json:"status"`
	// The title of the change request.
	Title string `json:"title"`
	// The case type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestResponseAttributes instantiates a new ChangeRequestResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestResponseAttributes(attributes map[string][]string, changeRequestLinkedIncidentUuid string, changeRequestMaintenanceWindowQuery string, changeRequestPlan string, changeRequestRisk ChangeRequestRiskLevel, changeRequestType ChangeRequestChangeType, createdAt time.Time, creationSource string, description string, key string, modifiedAt time.Time, planNotebookId int64, priority string, projectId string, status string, title string, typeVar string) *ChangeRequestResponseAttributes {
	this := ChangeRequestResponseAttributes{}
	this.Attributes = attributes
	this.ChangeRequestLinkedIncidentUuid = changeRequestLinkedIncidentUuid
	this.ChangeRequestMaintenanceWindowQuery = changeRequestMaintenanceWindowQuery
	this.ChangeRequestPlan = changeRequestPlan
	this.ChangeRequestRisk = changeRequestRisk
	this.ChangeRequestType = changeRequestType
	this.CreatedAt = createdAt
	this.CreationSource = creationSource
	this.Description = description
	this.Key = key
	this.ModifiedAt = modifiedAt
	this.PlanNotebookId = planNotebookId
	this.Priority = priority
	this.ProjectId = projectId
	this.Status = status
	this.Title = title
	this.Type = typeVar
	return &this
}

// NewChangeRequestResponseAttributesWithDefaults instantiates a new ChangeRequestResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestResponseAttributesWithDefaults() *ChangeRequestResponseAttributes {
	this := ChangeRequestResponseAttributes{}
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeRequestResponseAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt.Get()
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ChangeRequestResponseAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivedAt.Get(), o.ArchivedAt.IsSet()
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *ChangeRequestResponseAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt.IsSet()
}

// SetArchivedAt gets a reference to the given datadog.NullableTime and assigns it to the ArchivedAt field.
func (o *ChangeRequestResponseAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt.Set(&v)
}

// SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil.
func (o *ChangeRequestResponseAttributes) SetArchivedAtNil() {
	o.ArchivedAt.Set(nil)
}

// UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil.
func (o *ChangeRequestResponseAttributes) UnsetArchivedAt() {
	o.ArchivedAt.Unset()
}

// GetAttributes returns the Attributes field value.
func (o *ChangeRequestResponseAttributes) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ChangeRequestResponseAttributes) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetChangeRequestLinkedIncidentUuid returns the ChangeRequestLinkedIncidentUuid field value.
func (o *ChangeRequestResponseAttributes) GetChangeRequestLinkedIncidentUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChangeRequestLinkedIncidentUuid
}

// GetChangeRequestLinkedIncidentUuidOk returns a tuple with the ChangeRequestLinkedIncidentUuid field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetChangeRequestLinkedIncidentUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestLinkedIncidentUuid, true
}

// SetChangeRequestLinkedIncidentUuid sets field value.
func (o *ChangeRequestResponseAttributes) SetChangeRequestLinkedIncidentUuid(v string) {
	o.ChangeRequestLinkedIncidentUuid = v
}

// GetChangeRequestMaintenanceWindowQuery returns the ChangeRequestMaintenanceWindowQuery field value.
func (o *ChangeRequestResponseAttributes) GetChangeRequestMaintenanceWindowQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChangeRequestMaintenanceWindowQuery
}

// GetChangeRequestMaintenanceWindowQueryOk returns a tuple with the ChangeRequestMaintenanceWindowQuery field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetChangeRequestMaintenanceWindowQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestMaintenanceWindowQuery, true
}

// SetChangeRequestMaintenanceWindowQuery sets field value.
func (o *ChangeRequestResponseAttributes) SetChangeRequestMaintenanceWindowQuery(v string) {
	o.ChangeRequestMaintenanceWindowQuery = v
}

// GetChangeRequestPlan returns the ChangeRequestPlan field value.
func (o *ChangeRequestResponseAttributes) GetChangeRequestPlan() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChangeRequestPlan
}

// GetChangeRequestPlanOk returns a tuple with the ChangeRequestPlan field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetChangeRequestPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestPlan, true
}

// SetChangeRequestPlan sets field value.
func (o *ChangeRequestResponseAttributes) SetChangeRequestPlan(v string) {
	o.ChangeRequestPlan = v
}

// GetChangeRequestRisk returns the ChangeRequestRisk field value.
func (o *ChangeRequestResponseAttributes) GetChangeRequestRisk() ChangeRequestRiskLevel {
	if o == nil {
		var ret ChangeRequestRiskLevel
		return ret
	}
	return o.ChangeRequestRisk
}

// GetChangeRequestRiskOk returns a tuple with the ChangeRequestRisk field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetChangeRequestRiskOk() (*ChangeRequestRiskLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestRisk, true
}

// SetChangeRequestRisk sets field value.
func (o *ChangeRequestResponseAttributes) SetChangeRequestRisk(v ChangeRequestRiskLevel) {
	o.ChangeRequestRisk = v
}

// GetChangeRequestType returns the ChangeRequestType field value.
func (o *ChangeRequestResponseAttributes) GetChangeRequestType() ChangeRequestChangeType {
	if o == nil {
		var ret ChangeRequestChangeType
		return ret
	}
	return o.ChangeRequestType
}

// GetChangeRequestTypeOk returns a tuple with the ChangeRequestType field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetChangeRequestTypeOk() (*ChangeRequestChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestType, true
}

// SetChangeRequestType sets field value.
func (o *ChangeRequestResponseAttributes) SetChangeRequestType(v ChangeRequestChangeType) {
	o.ChangeRequestType = v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeRequestResponseAttributes) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt.Get()
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ChangeRequestResponseAttributes) GetClosedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedAt.Get(), o.ClosedAt.IsSet()
}

// HasClosedAt returns a boolean if a field has been set.
func (o *ChangeRequestResponseAttributes) HasClosedAt() bool {
	return o != nil && o.ClosedAt.IsSet()
}

// SetClosedAt gets a reference to the given datadog.NullableTime and assigns it to the ClosedAt field.
func (o *ChangeRequestResponseAttributes) SetClosedAt(v time.Time) {
	o.ClosedAt.Set(&v)
}

// SetClosedAtNil sets the value for ClosedAt to be an explicit nil.
func (o *ChangeRequestResponseAttributes) SetClosedAtNil() {
	o.ClosedAt.Set(nil)
}

// UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil.
func (o *ChangeRequestResponseAttributes) UnsetClosedAt() {
	o.ClosedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ChangeRequestResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ChangeRequestResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreationSource returns the CreationSource field value.
func (o *ChangeRequestResponseAttributes) GetCreationSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreationSource
}

// GetCreationSourceOk returns a tuple with the CreationSource field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetCreationSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationSource, true
}

// SetCreationSource sets field value.
func (o *ChangeRequestResponseAttributes) SetCreationSource(v string) {
	o.CreationSource = v
}

// GetDescription returns the Description field value.
func (o *ChangeRequestResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ChangeRequestResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ChangeRequestResponseAttributes) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ChangeRequestResponseAttributes) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ChangeRequestResponseAttributes) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetKey returns the Key field value.
func (o *ChangeRequestResponseAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ChangeRequestResponseAttributes) SetKey(v string) {
	o.Key = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *ChangeRequestResponseAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *ChangeRequestResponseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetPlanNotebookId returns the PlanNotebookId field value.
func (o *ChangeRequestResponseAttributes) GetPlanNotebookId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PlanNotebookId
}

// GetPlanNotebookIdOk returns a tuple with the PlanNotebookId field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetPlanNotebookIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanNotebookId, true
}

// SetPlanNotebookId sets field value.
func (o *ChangeRequestResponseAttributes) SetPlanNotebookId(v int64) {
	o.PlanNotebookId = v
}

// GetPriority returns the Priority field value.
func (o *ChangeRequestResponseAttributes) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *ChangeRequestResponseAttributes) SetPriority(v string) {
	o.Priority = v
}

// GetProjectId returns the ProjectId field value.
func (o *ChangeRequestResponseAttributes) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *ChangeRequestResponseAttributes) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ChangeRequestResponseAttributes) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ChangeRequestResponseAttributes) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ChangeRequestResponseAttributes) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStatus returns the Status field value.
func (o *ChangeRequestResponseAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ChangeRequestResponseAttributes) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *ChangeRequestResponseAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ChangeRequestResponseAttributes) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value.
func (o *ChangeRequestResponseAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestResponseAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChangeRequestResponseAttributes) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt.IsSet() {
		toSerialize["archived_at"] = o.ArchivedAt.Get()
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["change_request_linked_incident_uuid"] = o.ChangeRequestLinkedIncidentUuid
	toSerialize["change_request_maintenance_window_query"] = o.ChangeRequestMaintenanceWindowQuery
	toSerialize["change_request_plan"] = o.ChangeRequestPlan
	toSerialize["change_request_risk"] = o.ChangeRequestRisk
	toSerialize["change_request_type"] = o.ChangeRequestType
	if o.ClosedAt.IsSet() {
		toSerialize["closed_at"] = o.ClosedAt.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["creation_source"] = o.CreationSource
	toSerialize["description"] = o.Description
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["key"] = o.Key
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["plan_notebook_id"] = o.PlanNotebookId
	toSerialize["priority"] = o.Priority
	toSerialize["project_id"] = o.ProjectId
	if o.StartDate != nil {
		if o.StartDate.Nanosecond() == 0 {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt                          datadog.NullableTime     `json:"archived_at,omitempty"`
		Attributes                          *map[string][]string     `json:"attributes"`
		ChangeRequestLinkedIncidentUuid     *string                  `json:"change_request_linked_incident_uuid"`
		ChangeRequestMaintenanceWindowQuery *string                  `json:"change_request_maintenance_window_query"`
		ChangeRequestPlan                   *string                  `json:"change_request_plan"`
		ChangeRequestRisk                   *ChangeRequestRiskLevel  `json:"change_request_risk"`
		ChangeRequestType                   *ChangeRequestChangeType `json:"change_request_type"`
		ClosedAt                            datadog.NullableTime     `json:"closed_at,omitempty"`
		CreatedAt                           *time.Time               `json:"created_at"`
		CreationSource                      *string                  `json:"creation_source"`
		Description                         *string                  `json:"description"`
		EndDate                             *time.Time               `json:"end_date,omitempty"`
		Key                                 *string                  `json:"key"`
		ModifiedAt                          *time.Time               `json:"modified_at"`
		PlanNotebookId                      *int64                   `json:"plan_notebook_id"`
		Priority                            *string                  `json:"priority"`
		ProjectId                           *string                  `json:"project_id"`
		StartDate                           *time.Time               `json:"start_date,omitempty"`
		Status                              *string                  `json:"status"`
		Title                               *string                  `json:"title"`
		Type                                *string                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.ChangeRequestLinkedIncidentUuid == nil {
		return fmt.Errorf("required field change_request_linked_incident_uuid missing")
	}
	if all.ChangeRequestMaintenanceWindowQuery == nil {
		return fmt.Errorf("required field change_request_maintenance_window_query missing")
	}
	if all.ChangeRequestPlan == nil {
		return fmt.Errorf("required field change_request_plan missing")
	}
	if all.ChangeRequestRisk == nil {
		return fmt.Errorf("required field change_request_risk missing")
	}
	if all.ChangeRequestType == nil {
		return fmt.Errorf("required field change_request_type missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreationSource == nil {
		return fmt.Errorf("required field creation_source missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.PlanNotebookId == nil {
		return fmt.Errorf("required field plan_notebook_id missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "attributes", "change_request_linked_incident_uuid", "change_request_maintenance_window_query", "change_request_plan", "change_request_risk", "change_request_type", "closed_at", "created_at", "creation_source", "description", "end_date", "key", "modified_at", "plan_notebook_id", "priority", "project_id", "start_date", "status", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
	o.Attributes = *all.Attributes
	o.ChangeRequestLinkedIncidentUuid = *all.ChangeRequestLinkedIncidentUuid
	o.ChangeRequestMaintenanceWindowQuery = *all.ChangeRequestMaintenanceWindowQuery
	o.ChangeRequestPlan = *all.ChangeRequestPlan
	if !all.ChangeRequestRisk.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestRisk = *all.ChangeRequestRisk
	}
	if !all.ChangeRequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestType = *all.ChangeRequestType
	}
	o.ClosedAt = all.ClosedAt
	o.CreatedAt = *all.CreatedAt
	o.CreationSource = *all.CreationSource
	o.Description = *all.Description
	o.EndDate = all.EndDate
	o.Key = *all.Key
	o.ModifiedAt = *all.ModifiedAt
	o.PlanNotebookId = *all.PlanNotebookId
	o.Priority = *all.Priority
	o.ProjectId = *all.ProjectId
	o.StartDate = all.StartDate
	o.Status = *all.Status
	o.Title = *all.Title
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
