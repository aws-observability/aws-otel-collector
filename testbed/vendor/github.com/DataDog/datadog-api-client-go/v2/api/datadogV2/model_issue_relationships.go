// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueRelationships Relationship between the issue and an assignee, case and/or teams.
type IssueRelationships struct {
	// Relationship between the issue and assignee.
	Assignee *IssueAssigneeRelationship `json:"assignee,omitempty"`
	// Relationship between the issue and case.
	Case *IssueCaseRelationship `json:"case,omitempty"`
	// Relationship between the issue and teams.
	TeamOwners *IssueTeamOwnersRelationship `json:"team_owners,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueRelationships instantiates a new IssueRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueRelationships() *IssueRelationships {
	this := IssueRelationships{}
	return &this
}

// NewIssueRelationshipsWithDefaults instantiates a new IssueRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueRelationshipsWithDefaults() *IssueRelationships {
	this := IssueRelationships{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *IssueRelationships) GetAssignee() IssueAssigneeRelationship {
	if o == nil || o.Assignee == nil {
		var ret IssueAssigneeRelationship
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetAssigneeOk() (*IssueAssigneeRelationship, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *IssueRelationships) HasAssignee() bool {
	return o != nil && o.Assignee != nil
}

// SetAssignee gets a reference to the given IssueAssigneeRelationship and assigns it to the Assignee field.
func (o *IssueRelationships) SetAssignee(v IssueAssigneeRelationship) {
	o.Assignee = &v
}

// GetCase returns the Case field value if set, zero value otherwise.
func (o *IssueRelationships) GetCase() IssueCaseRelationship {
	if o == nil || o.Case == nil {
		var ret IssueCaseRelationship
		return ret
	}
	return *o.Case
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetCaseOk() (*IssueCaseRelationship, bool) {
	if o == nil || o.Case == nil {
		return nil, false
	}
	return o.Case, true
}

// HasCase returns a boolean if a field has been set.
func (o *IssueRelationships) HasCase() bool {
	return o != nil && o.Case != nil
}

// SetCase gets a reference to the given IssueCaseRelationship and assigns it to the Case field.
func (o *IssueRelationships) SetCase(v IssueCaseRelationship) {
	o.Case = &v
}

// GetTeamOwners returns the TeamOwners field value if set, zero value otherwise.
func (o *IssueRelationships) GetTeamOwners() IssueTeamOwnersRelationship {
	if o == nil || o.TeamOwners == nil {
		var ret IssueTeamOwnersRelationship
		return ret
	}
	return *o.TeamOwners
}

// GetTeamOwnersOk returns a tuple with the TeamOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetTeamOwnersOk() (*IssueTeamOwnersRelationship, bool) {
	if o == nil || o.TeamOwners == nil {
		return nil, false
	}
	return o.TeamOwners, true
}

// HasTeamOwners returns a boolean if a field has been set.
func (o *IssueRelationships) HasTeamOwners() bool {
	return o != nil && o.TeamOwners != nil
}

// SetTeamOwners gets a reference to the given IssueTeamOwnersRelationship and assigns it to the TeamOwners field.
func (o *IssueRelationships) SetTeamOwners(v IssueTeamOwnersRelationship) {
	o.TeamOwners = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Case != nil {
		toSerialize["case"] = o.Case
	}
	if o.TeamOwners != nil {
		toSerialize["team_owners"] = o.TeamOwners
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignee   *IssueAssigneeRelationship   `json:"assignee,omitempty"`
		Case       *IssueCaseRelationship       `json:"case,omitempty"`
		TeamOwners *IssueTeamOwnersRelationship `json:"team_owners,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee", "case", "team_owners"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assignee != nil && all.Assignee.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Assignee = all.Assignee
	if all.Case != nil && all.Case.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Case = all.Case
	if all.TeamOwners != nil && all.TeamOwners.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TeamOwners = all.TeamOwners

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
