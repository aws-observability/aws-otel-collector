// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingCaseResponseDataRelationships Relationships of the case.
type FindingCaseResponseDataRelationships struct {
	// Relationship to user.
	CreatedBy *RelationshipToUser `json:"created_by,omitempty"`
	// Relationship to user.
	ModifiedBy *RelationshipToUser `json:"modified_by,omitempty"`
	// Case management project.
	Project *CaseManagementProject `json:"project,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFindingCaseResponseDataRelationships instantiates a new FindingCaseResponseDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingCaseResponseDataRelationships() *FindingCaseResponseDataRelationships {
	this := FindingCaseResponseDataRelationships{}
	return &this
}

// NewFindingCaseResponseDataRelationshipsWithDefaults instantiates a new FindingCaseResponseDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingCaseResponseDataRelationshipsWithDefaults() *FindingCaseResponseDataRelationships {
	this := FindingCaseResponseDataRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FindingCaseResponseDataRelationships) GetCreatedBy() RelationshipToUser {
	if o == nil || o.CreatedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataRelationships) GetCreatedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FindingCaseResponseDataRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given RelationshipToUser and assigns it to the CreatedBy field.
func (o *FindingCaseResponseDataRelationships) SetCreatedBy(v RelationshipToUser) {
	o.CreatedBy = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *FindingCaseResponseDataRelationships) GetModifiedBy() RelationshipToUser {
	if o == nil || o.ModifiedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataRelationships) GetModifiedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *FindingCaseResponseDataRelationships) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given RelationshipToUser and assigns it to the ModifiedBy field.
func (o *FindingCaseResponseDataRelationships) SetModifiedBy(v RelationshipToUser) {
	o.ModifiedBy = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *FindingCaseResponseDataRelationships) GetProject() CaseManagementProject {
	if o == nil || o.Project == nil {
		var ret CaseManagementProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingCaseResponseDataRelationships) GetProjectOk() (*CaseManagementProject, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *FindingCaseResponseDataRelationships) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given CaseManagementProject and assigns it to the Project field.
func (o *FindingCaseResponseDataRelationships) SetProject(v CaseManagementProject) {
	o.Project = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingCaseResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingCaseResponseDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy  *RelationshipToUser    `json:"created_by,omitempty"`
		ModifiedBy *RelationshipToUser    `json:"modified_by,omitempty"`
		Project    *CaseManagementProject `json:"project,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "modified_by", "project"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	if all.ModifiedBy != nil && all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = all.ModifiedBy
	if all.Project != nil && all.Project.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Project = all.Project

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
