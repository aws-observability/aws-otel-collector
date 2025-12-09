// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseCreateRelationships Relationships formed with the case on creation
type CaseCreateRelationships struct {
	// Relationship to user.
	Assignee NullableNullableUserRelationship `json:"assignee,omitempty"`
	// Relationship to project
	Project ProjectRelationship `json:"project"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseCreateRelationships instantiates a new CaseCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseCreateRelationships(project ProjectRelationship) *CaseCreateRelationships {
	this := CaseCreateRelationships{}
	this.Project = project
	return &this
}

// NewCaseCreateRelationshipsWithDefaults instantiates a new CaseCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseCreateRelationshipsWithDefaults() *CaseCreateRelationships {
	this := CaseCreateRelationships{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseCreateRelationships) GetAssignee() NullableUserRelationship {
	if o == nil || o.Assignee.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseCreateRelationships) GetAssigneeOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// HasAssignee returns a boolean if a field has been set.
func (o *CaseCreateRelationships) HasAssignee() bool {
	return o != nil && o.Assignee.IsSet()
}

// SetAssignee gets a reference to the given NullableNullableUserRelationship and assigns it to the Assignee field.
func (o *CaseCreateRelationships) SetAssignee(v NullableUserRelationship) {
	o.Assignee.Set(&v)
}

// SetAssigneeNil sets the value for Assignee to be an explicit nil.
func (o *CaseCreateRelationships) SetAssigneeNil() {
	o.Assignee.Set(nil)
}

// UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil.
func (o *CaseCreateRelationships) UnsetAssignee() {
	o.Assignee.Unset()
}

// GetProject returns the Project field value.
func (o *CaseCreateRelationships) GetProject() ProjectRelationship {
	if o == nil {
		var ret ProjectRelationship
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *CaseCreateRelationships) GetProjectOk() (*ProjectRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value.
func (o *CaseCreateRelationships) SetProject(v ProjectRelationship) {
	o.Project = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignee.IsSet() {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	toSerialize["project"] = o.Project

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignee NullableNullableUserRelationship `json:"assignee,omitempty"`
		Project  *ProjectRelationship             `json:"project"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Project == nil {
		return fmt.Errorf("required field project missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee", "project"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assignee = all.Assignee
	if all.Project.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Project = *all.Project

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
