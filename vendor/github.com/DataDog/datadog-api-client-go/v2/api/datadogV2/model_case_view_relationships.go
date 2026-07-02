// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseViewRelationships Related resources for the case view, including the creator, last modifier, and associated project.
type CaseViewRelationships struct {
	// Relationship to user.
	CreatedBy NullableNullableUserRelationship `json:"created_by,omitempty"`
	// Relationship to user.
	ModifiedBy NullableNullableUserRelationship `json:"modified_by,omitempty"`
	// Relationship to project.
	Project *ProjectRelationship `json:"project,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseViewRelationships instantiates a new CaseViewRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseViewRelationships() *CaseViewRelationships {
	this := CaseViewRelationships{}
	return &this
}

// NewCaseViewRelationshipsWithDefaults instantiates a new CaseViewRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseViewRelationshipsWithDefaults() *CaseViewRelationships {
	this := CaseViewRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseViewRelationships) GetCreatedBy() NullableUserRelationship {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseViewRelationships) GetCreatedByOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CaseViewRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy.IsSet()
}

// SetCreatedBy gets a reference to the given NullableNullableUserRelationship and assigns it to the CreatedBy field.
func (o *CaseViewRelationships) SetCreatedBy(v NullableUserRelationship) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil.
func (o *CaseViewRelationships) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil.
func (o *CaseViewRelationships) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaseViewRelationships) GetModifiedBy() NullableUserRelationship {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret NullableUserRelationship
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CaseViewRelationships) GetModifiedByOk() (*NullableUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *CaseViewRelationships) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy.IsSet()
}

// SetModifiedBy gets a reference to the given NullableNullableUserRelationship and assigns it to the ModifiedBy field.
func (o *CaseViewRelationships) SetModifiedBy(v NullableUserRelationship) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil.
func (o *CaseViewRelationships) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil.
func (o *CaseViewRelationships) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CaseViewRelationships) GetProject() ProjectRelationship {
	if o == nil || o.Project == nil {
		var ret ProjectRelationship
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewRelationships) GetProjectOk() (*ProjectRelationship, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CaseViewRelationships) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given ProjectRelationship and assigns it to the Project field.
func (o *CaseViewRelationships) SetProject(v ProjectRelationship) {
	o.Project = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseViewRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *CaseViewRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy  NullableNullableUserRelationship `json:"created_by,omitempty"`
		ModifiedBy NullableNullableUserRelationship `json:"modified_by,omitempty"`
		Project    *ProjectRelationship             `json:"project,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "modified_by", "project"})
	} else {
		return err
	}

	hasInvalidField := false
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
