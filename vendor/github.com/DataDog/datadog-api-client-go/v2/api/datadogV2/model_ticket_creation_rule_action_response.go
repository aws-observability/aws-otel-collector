// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TicketCreationRuleActionResponse The action to take when the ticket creation rule matches a finding.
type TicketCreationRuleActionResponse struct {
	// The UUID of the default assignee for created tickets.
	AssigneeId *uuid.UUID `json:"assignee_id,omitempty"`
	// The reason the rule was automatically disabled by the system due to a ticketing integration error.
	AutoDisabledReason *string `json:"auto_disabled_reason,omitempty"`
	// Custom fields of the Jira issue to create. For the list of available fields, see [Jira documentation](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-createmeta-projectidorkey-issuetypes-issuetypeid-get).
	Fields interface{} `json:"fields,omitempty"`
	// The maximum number of tickets the rule may create per day. If exceeded, one final ticket will be created, explaining the limit was hit and link back to the responsible rule.
	MaxTicketsPerDay int64 `json:"max_tickets_per_day"`
	// The UUID of the case management project.
	ProjectId uuid.UUID `json:"project_id"`
	// The ticketing system to create tickets in.
	Target TicketCreationTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTicketCreationRuleActionResponse instantiates a new TicketCreationRuleActionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTicketCreationRuleActionResponse(maxTicketsPerDay int64, projectId uuid.UUID, target TicketCreationTarget) *TicketCreationRuleActionResponse {
	this := TicketCreationRuleActionResponse{}
	this.MaxTicketsPerDay = maxTicketsPerDay
	this.ProjectId = projectId
	this.Target = target
	return &this
}

// NewTicketCreationRuleActionResponseWithDefaults instantiates a new TicketCreationRuleActionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTicketCreationRuleActionResponseWithDefaults() *TicketCreationRuleActionResponse {
	this := TicketCreationRuleActionResponse{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *TicketCreationRuleActionResponse) GetAssigneeId() uuid.UUID {
	if o == nil || o.AssigneeId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCreationRuleActionResponse) GetAssigneeIdOk() (*uuid.UUID, bool) {
	if o == nil || o.AssigneeId == nil {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *TicketCreationRuleActionResponse) HasAssigneeId() bool {
	return o != nil && o.AssigneeId != nil
}

// SetAssigneeId gets a reference to the given uuid.UUID and assigns it to the AssigneeId field.
func (o *TicketCreationRuleActionResponse) SetAssigneeId(v uuid.UUID) {
	o.AssigneeId = &v
}

// GetAutoDisabledReason returns the AutoDisabledReason field value if set, zero value otherwise.
func (o *TicketCreationRuleActionResponse) GetAutoDisabledReason() string {
	if o == nil || o.AutoDisabledReason == nil {
		var ret string
		return ret
	}
	return *o.AutoDisabledReason
}

// GetAutoDisabledReasonOk returns a tuple with the AutoDisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCreationRuleActionResponse) GetAutoDisabledReasonOk() (*string, bool) {
	if o == nil || o.AutoDisabledReason == nil {
		return nil, false
	}
	return o.AutoDisabledReason, true
}

// HasAutoDisabledReason returns a boolean if a field has been set.
func (o *TicketCreationRuleActionResponse) HasAutoDisabledReason() bool {
	return o != nil && o.AutoDisabledReason != nil
}

// SetAutoDisabledReason gets a reference to the given string and assigns it to the AutoDisabledReason field.
func (o *TicketCreationRuleActionResponse) SetAutoDisabledReason(v string) {
	o.AutoDisabledReason = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TicketCreationRuleActionResponse) GetFields() interface{} {
	if o == nil || o.Fields == nil {
		var ret interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCreationRuleActionResponse) GetFieldsOk() (*interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TicketCreationRuleActionResponse) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given interface{} and assigns it to the Fields field.
func (o *TicketCreationRuleActionResponse) SetFields(v interface{}) {
	o.Fields = v
}

// GetMaxTicketsPerDay returns the MaxTicketsPerDay field value.
func (o *TicketCreationRuleActionResponse) GetMaxTicketsPerDay() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxTicketsPerDay
}

// GetMaxTicketsPerDayOk returns a tuple with the MaxTicketsPerDay field value
// and a boolean to check if the value has been set.
func (o *TicketCreationRuleActionResponse) GetMaxTicketsPerDayOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTicketsPerDay, true
}

// SetMaxTicketsPerDay sets field value.
func (o *TicketCreationRuleActionResponse) SetMaxTicketsPerDay(v int64) {
	o.MaxTicketsPerDay = v
}

// GetProjectId returns the ProjectId field value.
func (o *TicketCreationRuleActionResponse) GetProjectId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TicketCreationRuleActionResponse) GetProjectIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *TicketCreationRuleActionResponse) SetProjectId(v uuid.UUID) {
	o.ProjectId = v
}

// GetTarget returns the Target field value.
func (o *TicketCreationRuleActionResponse) GetTarget() TicketCreationTarget {
	if o == nil {
		var ret TicketCreationTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *TicketCreationRuleActionResponse) GetTargetOk() (*TicketCreationTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *TicketCreationRuleActionResponse) SetTarget(v TicketCreationTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TicketCreationRuleActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeId != nil {
		toSerialize["assignee_id"] = o.AssigneeId
	}
	if o.AutoDisabledReason != nil {
		toSerialize["auto_disabled_reason"] = o.AutoDisabledReason
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["max_tickets_per_day"] = o.MaxTicketsPerDay
	toSerialize["project_id"] = o.ProjectId
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TicketCreationRuleActionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeId         *uuid.UUID            `json:"assignee_id,omitempty"`
		AutoDisabledReason *string               `json:"auto_disabled_reason,omitempty"`
		Fields             interface{}           `json:"fields,omitempty"`
		MaxTicketsPerDay   *int64                `json:"max_tickets_per_day"`
		ProjectId          *uuid.UUID            `json:"project_id"`
		Target             *TicketCreationTarget `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MaxTicketsPerDay == nil {
		return fmt.Errorf("required field max_tickets_per_day missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee_id", "auto_disabled_reason", "fields", "max_tickets_per_day", "project_id", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssigneeId = all.AssigneeId
	o.AutoDisabledReason = all.AutoDisabledReason
	o.Fields = all.Fields
	o.MaxTicketsPerDay = *all.MaxTicketsPerDay
	o.ProjectId = *all.ProjectId
	if !all.Target.IsValid() {
		hasInvalidField = true
	} else {
		o.Target = *all.Target
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
