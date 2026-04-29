// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestCreateAttributes Attributes for creating a change request.
type ChangeRequestCreateAttributes struct {
	// The UUID of an incident to link to the change request.
	ChangeRequestLinkedIncidentUuid *string `json:"change_request_linked_incident_uuid,omitempty"`
	// The maintenance window query for the change request.
	ChangeRequestMaintenanceWindowQuery *string `json:"change_request_maintenance_window_query,omitempty"`
	// The plan associated with the change request.
	ChangeRequestPlan *string `json:"change_request_plan,omitempty"`
	// The risk level of the change request.
	ChangeRequestRisk *ChangeRequestRiskLevel `json:"change_request_risk,omitempty"`
	// The type of the change request.
	ChangeRequestType *ChangeRequestChangeType `json:"change_request_type,omitempty"`
	// The description of the change request.
	Description *string `json:"description,omitempty"`
	// The planned end date of the change request.
	EndDate *time.Time `json:"end_date,omitempty"`
	// The project UUID to associate with the change request.
	ProjectId *string `json:"project_id,omitempty"`
	// A list of team handles to request decisions from.
	RequestedTeams []string `json:"requested_teams,omitempty"`
	// The planned start date of the change request.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The title of the change request.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestCreateAttributes instantiates a new ChangeRequestCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestCreateAttributes(title string) *ChangeRequestCreateAttributes {
	this := ChangeRequestCreateAttributes{}
	this.Title = title
	return &this
}

// NewChangeRequestCreateAttributesWithDefaults instantiates a new ChangeRequestCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestCreateAttributesWithDefaults() *ChangeRequestCreateAttributes {
	this := ChangeRequestCreateAttributes{}
	return &this
}

// GetChangeRequestLinkedIncidentUuid returns the ChangeRequestLinkedIncidentUuid field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetChangeRequestLinkedIncidentUuid() string {
	if o == nil || o.ChangeRequestLinkedIncidentUuid == nil {
		var ret string
		return ret
	}
	return *o.ChangeRequestLinkedIncidentUuid
}

// GetChangeRequestLinkedIncidentUuidOk returns a tuple with the ChangeRequestLinkedIncidentUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetChangeRequestLinkedIncidentUuidOk() (*string, bool) {
	if o == nil || o.ChangeRequestLinkedIncidentUuid == nil {
		return nil, false
	}
	return o.ChangeRequestLinkedIncidentUuid, true
}

// HasChangeRequestLinkedIncidentUuid returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasChangeRequestLinkedIncidentUuid() bool {
	return o != nil && o.ChangeRequestLinkedIncidentUuid != nil
}

// SetChangeRequestLinkedIncidentUuid gets a reference to the given string and assigns it to the ChangeRequestLinkedIncidentUuid field.
func (o *ChangeRequestCreateAttributes) SetChangeRequestLinkedIncidentUuid(v string) {
	o.ChangeRequestLinkedIncidentUuid = &v
}

// GetChangeRequestMaintenanceWindowQuery returns the ChangeRequestMaintenanceWindowQuery field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetChangeRequestMaintenanceWindowQuery() string {
	if o == nil || o.ChangeRequestMaintenanceWindowQuery == nil {
		var ret string
		return ret
	}
	return *o.ChangeRequestMaintenanceWindowQuery
}

// GetChangeRequestMaintenanceWindowQueryOk returns a tuple with the ChangeRequestMaintenanceWindowQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetChangeRequestMaintenanceWindowQueryOk() (*string, bool) {
	if o == nil || o.ChangeRequestMaintenanceWindowQuery == nil {
		return nil, false
	}
	return o.ChangeRequestMaintenanceWindowQuery, true
}

// HasChangeRequestMaintenanceWindowQuery returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasChangeRequestMaintenanceWindowQuery() bool {
	return o != nil && o.ChangeRequestMaintenanceWindowQuery != nil
}

// SetChangeRequestMaintenanceWindowQuery gets a reference to the given string and assigns it to the ChangeRequestMaintenanceWindowQuery field.
func (o *ChangeRequestCreateAttributes) SetChangeRequestMaintenanceWindowQuery(v string) {
	o.ChangeRequestMaintenanceWindowQuery = &v
}

// GetChangeRequestPlan returns the ChangeRequestPlan field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetChangeRequestPlan() string {
	if o == nil || o.ChangeRequestPlan == nil {
		var ret string
		return ret
	}
	return *o.ChangeRequestPlan
}

// GetChangeRequestPlanOk returns a tuple with the ChangeRequestPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetChangeRequestPlanOk() (*string, bool) {
	if o == nil || o.ChangeRequestPlan == nil {
		return nil, false
	}
	return o.ChangeRequestPlan, true
}

// HasChangeRequestPlan returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasChangeRequestPlan() bool {
	return o != nil && o.ChangeRequestPlan != nil
}

// SetChangeRequestPlan gets a reference to the given string and assigns it to the ChangeRequestPlan field.
func (o *ChangeRequestCreateAttributes) SetChangeRequestPlan(v string) {
	o.ChangeRequestPlan = &v
}

// GetChangeRequestRisk returns the ChangeRequestRisk field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetChangeRequestRisk() ChangeRequestRiskLevel {
	if o == nil || o.ChangeRequestRisk == nil {
		var ret ChangeRequestRiskLevel
		return ret
	}
	return *o.ChangeRequestRisk
}

// GetChangeRequestRiskOk returns a tuple with the ChangeRequestRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetChangeRequestRiskOk() (*ChangeRequestRiskLevel, bool) {
	if o == nil || o.ChangeRequestRisk == nil {
		return nil, false
	}
	return o.ChangeRequestRisk, true
}

// HasChangeRequestRisk returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasChangeRequestRisk() bool {
	return o != nil && o.ChangeRequestRisk != nil
}

// SetChangeRequestRisk gets a reference to the given ChangeRequestRiskLevel and assigns it to the ChangeRequestRisk field.
func (o *ChangeRequestCreateAttributes) SetChangeRequestRisk(v ChangeRequestRiskLevel) {
	o.ChangeRequestRisk = &v
}

// GetChangeRequestType returns the ChangeRequestType field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetChangeRequestType() ChangeRequestChangeType {
	if o == nil || o.ChangeRequestType == nil {
		var ret ChangeRequestChangeType
		return ret
	}
	return *o.ChangeRequestType
}

// GetChangeRequestTypeOk returns a tuple with the ChangeRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetChangeRequestTypeOk() (*ChangeRequestChangeType, bool) {
	if o == nil || o.ChangeRequestType == nil {
		return nil, false
	}
	return o.ChangeRequestType, true
}

// HasChangeRequestType returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasChangeRequestType() bool {
	return o != nil && o.ChangeRequestType != nil
}

// SetChangeRequestType gets a reference to the given ChangeRequestChangeType and assigns it to the ChangeRequestType field.
func (o *ChangeRequestCreateAttributes) SetChangeRequestType(v ChangeRequestChangeType) {
	o.ChangeRequestType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChangeRequestCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ChangeRequestCreateAttributes) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ChangeRequestCreateAttributes) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRequestedTeams returns the RequestedTeams field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetRequestedTeams() []string {
	if o == nil || o.RequestedTeams == nil {
		var ret []string
		return ret
	}
	return o.RequestedTeams
}

// GetRequestedTeamsOk returns a tuple with the RequestedTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetRequestedTeamsOk() (*[]string, bool) {
	if o == nil || o.RequestedTeams == nil {
		return nil, false
	}
	return &o.RequestedTeams, true
}

// HasRequestedTeams returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasRequestedTeams() bool {
	return o != nil && o.RequestedTeams != nil
}

// SetRequestedTeams gets a reference to the given []string and assigns it to the RequestedTeams field.
func (o *ChangeRequestCreateAttributes) SetRequestedTeams(v []string) {
	o.RequestedTeams = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ChangeRequestCreateAttributes) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ChangeRequestCreateAttributes) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ChangeRequestCreateAttributes) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetTitle returns the Title field value.
func (o *ChangeRequestCreateAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ChangeRequestCreateAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeRequestLinkedIncidentUuid != nil {
		toSerialize["change_request_linked_incident_uuid"] = o.ChangeRequestLinkedIncidentUuid
	}
	if o.ChangeRequestMaintenanceWindowQuery != nil {
		toSerialize["change_request_maintenance_window_query"] = o.ChangeRequestMaintenanceWindowQuery
	}
	if o.ChangeRequestPlan != nil {
		toSerialize["change_request_plan"] = o.ChangeRequestPlan
	}
	if o.ChangeRequestRisk != nil {
		toSerialize["change_request_risk"] = o.ChangeRequestRisk
	}
	if o.ChangeRequestType != nil {
		toSerialize["change_request_type"] = o.ChangeRequestType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.RequestedTeams != nil {
		toSerialize["requested_teams"] = o.RequestedTeams
	}
	if o.StartDate != nil {
		if o.StartDate.Nanosecond() == 0 {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestLinkedIncidentUuid     *string                  `json:"change_request_linked_incident_uuid,omitempty"`
		ChangeRequestMaintenanceWindowQuery *string                  `json:"change_request_maintenance_window_query,omitempty"`
		ChangeRequestPlan                   *string                  `json:"change_request_plan,omitempty"`
		ChangeRequestRisk                   *ChangeRequestRiskLevel  `json:"change_request_risk,omitempty"`
		ChangeRequestType                   *ChangeRequestChangeType `json:"change_request_type,omitempty"`
		Description                         *string                  `json:"description,omitempty"`
		EndDate                             *time.Time               `json:"end_date,omitempty"`
		ProjectId                           *string                  `json:"project_id,omitempty"`
		RequestedTeams                      []string                 `json:"requested_teams,omitempty"`
		StartDate                           *time.Time               `json:"start_date,omitempty"`
		Title                               *string                  `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_linked_incident_uuid", "change_request_maintenance_window_query", "change_request_plan", "change_request_risk", "change_request_type", "description", "end_date", "project_id", "requested_teams", "start_date", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChangeRequestLinkedIncidentUuid = all.ChangeRequestLinkedIncidentUuid
	o.ChangeRequestMaintenanceWindowQuery = all.ChangeRequestMaintenanceWindowQuery
	o.ChangeRequestPlan = all.ChangeRequestPlan
	if all.ChangeRequestRisk != nil && !all.ChangeRequestRisk.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestRisk = all.ChangeRequestRisk
	}
	if all.ChangeRequestType != nil && !all.ChangeRequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestType = all.ChangeRequestType
	}
	o.Description = all.Description
	o.EndDate = all.EndDate
	o.ProjectId = all.ProjectId
	o.RequestedTeams = all.RequestedTeams
	o.StartDate = all.StartDate
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
