// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueCaseRelationships Resources related to a case.
type IssueCaseRelationships struct {
	// Relationship to user.
	Assignee NullableNullableUserRelationship `json:"assignee,omitempty"`
	// Relationship to user.
	CreatedBy NullableNullableUserRelationship `json:"created_by,omitempty"`
	// Relationship to user.
	ModifiedBy NullableNullableUserRelationship `json:"modified_by,omitempty"`
	// Relationship to project
	Project *ProjectRelationship `json:"project,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueCaseRelationships instantiates a new IssueCaseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueCaseRelationships() *IssueCaseRelationships {
	this := IssueCaseRelationships{}
	return &this
}

// NewIssueCaseRelationshipsWithDefaults instantiates a new IssueCaseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueCaseRelationshipsWithDefaults() *IssueCaseRelationships {
	this := IssueCaseRelationships{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueCaseRelationships) GetAssignee() NullableUserRelationship {
	if o == nil || o.Assignee.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IssueCaseRelationships) GetAssigneeOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// HasAssignee returns a boolean if a field has been set.
func (o *IssueCaseRelationships) HasAssignee() bool {
	return o != nil && o.Assignee.IsSet()
}

// SetAssignee gets a reference to the given NullableNullableUserRelationship and assigns it to the Assignee field.
func (o *IssueCaseRelationships) SetAssignee(v NullableUserRelationship) {
	o.Assignee.Set(&v)
}

// SetAssigneeNil sets the value for Assignee to be an explicit nil.
func (o *IssueCaseRelationships) SetAssigneeNil() {
	o.Assignee.Set(nil)
}

// UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil.
func (o *IssueCaseRelationships) UnsetAssignee() {
	o.Assignee.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueCaseRelationships) GetCreatedBy() NullableUserRelationship {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IssueCaseRelationships) GetCreatedByOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IssueCaseRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy.IsSet()
}

// SetCreatedBy gets a reference to the given NullableNullableUserRelationship and assigns it to the CreatedBy field.
func (o *IssueCaseRelationships) SetCreatedBy(v NullableUserRelationship) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil.
func (o *IssueCaseRelationships) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil.
func (o *IssueCaseRelationships) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueCaseRelationships) GetModifiedBy() NullableUserRelationship {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IssueCaseRelationships) GetModifiedByOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *IssueCaseRelationships) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy.IsSet()
}

// SetModifiedBy gets a reference to the given NullableNullableUserRelationship and assigns it to the ModifiedBy field.
func (o *IssueCaseRelationships) SetModifiedBy(v NullableUserRelationship) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil.
func (o *IssueCaseRelationships) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil.
func (o *IssueCaseRelationships) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *IssueCaseRelationships) GetProject() ProjectRelationship {
	if o == nil || o.Project == nil {
		var ret ProjectRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseRelationships) GetProjectOk() (*ProjectRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *IssueCaseRelationships) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given ProjectRelationship and assigns it to the Project field.
func (o *IssueCaseRelationships) SetProject(v ProjectRelationship) {
	o.Project = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueCaseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignee.IsSet() {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modified_by"] = o.ModifiedBy.Get()
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
func (o *IssueCaseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignee   NullableNullableUserRelationship `json:"assignee,omitempty"`
		CreatedBy  NullableNullableUserRelationship `json:"created_by,omitempty"`
		ModifiedBy NullableNullableUserRelationship `json:"modified_by,omitempty"`
		Project    *ProjectRelationship             `json:"project,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee", "created_by", "modified_by", "project"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assignee = all.Assignee
	o.CreatedBy = all.CreatedBy
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
