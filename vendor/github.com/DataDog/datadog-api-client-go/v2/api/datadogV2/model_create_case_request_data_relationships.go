// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateCaseRequestDataRelationships Relationships of the case to create.
type CreateCaseRequestDataRelationships struct {
	// A list of security findings.
	Findings Findings `json:"findings"`
	// Case management project.
	Project CaseManagementProject `json:"project"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateCaseRequestDataRelationships instantiates a new CreateCaseRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateCaseRequestDataRelationships(findings Findings, project CaseManagementProject) *CreateCaseRequestDataRelationships {
	this := CreateCaseRequestDataRelationships{}
	this.Findings = findings
	this.Project = project
	return &this
}

// NewCreateCaseRequestDataRelationshipsWithDefaults instantiates a new CreateCaseRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateCaseRequestDataRelationshipsWithDefaults() *CreateCaseRequestDataRelationships {
	this := CreateCaseRequestDataRelationships{}
	return &this
}

// GetFindings returns the Findings field value.
func (o *CreateCaseRequestDataRelationships) GetFindings() Findings {
	if o == nil {
		var ret Findings
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *CreateCaseRequestDataRelationships) GetFindingsOk() (*Findings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *CreateCaseRequestDataRelationships) SetFindings(v Findings) {
	o.Findings = v
}

// GetProject returns the Project field value.
func (o *CreateCaseRequestDataRelationships) GetProject() CaseManagementProject {
	if o == nil {
		var ret CaseManagementProject
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *CreateCaseRequestDataRelationships) GetProjectOk() (*CaseManagementProject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value.
func (o *CreateCaseRequestDataRelationships) SetProject(v CaseManagementProject) {
	o.Project = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateCaseRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["findings"] = o.Findings
	toSerialize["project"] = o.Project

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateCaseRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Findings *Findings              `json:"findings"`
		Project  *CaseManagementProject `json:"project"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Findings == nil {
		return fmt.Errorf("required field findings missing")
	}
	if all.Project == nil {
		return fmt.Errorf("required field project missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"findings", "project"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Findings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Findings = *all.Findings
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
