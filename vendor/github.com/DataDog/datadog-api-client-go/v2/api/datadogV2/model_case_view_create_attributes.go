// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseViewCreateAttributes Attributes required to create a case view.
type CaseViewCreateAttributes struct {
	// The name of the view.
	Name string `json:"name"`
	// The identifier of a notification rule linked to this view. When set, users subscribed to the view receive alerts for matching cases.
	NpRuleId *string `json:"np_rule_id,omitempty"`
	// The UUID of the project this view belongs to. Views are scoped to a single project.
	ProjectId string `json:"project_id"`
	// The query used to filter cases in this view.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseViewCreateAttributes instantiates a new CaseViewCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseViewCreateAttributes(name string, projectId string, query string) *CaseViewCreateAttributes {
	this := CaseViewCreateAttributes{}
	this.Name = name
	this.ProjectId = projectId
	this.Query = query
	return &this
}

// NewCaseViewCreateAttributesWithDefaults instantiates a new CaseViewCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseViewCreateAttributesWithDefaults() *CaseViewCreateAttributes {
	this := CaseViewCreateAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *CaseViewCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaseViewCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CaseViewCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetNpRuleId returns the NpRuleId field value if set, zero value otherwise.
func (o *CaseViewCreateAttributes) GetNpRuleId() string {
	if o == nil || o.NpRuleId == nil {
		var ret string
		return ret
	}
	return *o.NpRuleId
}

// GetNpRuleIdOk returns a tuple with the NpRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewCreateAttributes) GetNpRuleIdOk() (*string, bool) {
	if o == nil || o.NpRuleId == nil {
		return nil, false
	}
	return o.NpRuleId, true
}

// HasNpRuleId returns a boolean if a field has been set.
func (o *CaseViewCreateAttributes) HasNpRuleId() bool {
	return o != nil && o.NpRuleId != nil
}

// SetNpRuleId gets a reference to the given string and assigns it to the NpRuleId field.
func (o *CaseViewCreateAttributes) SetNpRuleId(v string) {
	o.NpRuleId = &v
}

// GetProjectId returns the ProjectId field value.
func (o *CaseViewCreateAttributes) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CaseViewCreateAttributes) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *CaseViewCreateAttributes) SetProjectId(v string) {
	o.ProjectId = v
}

// GetQuery returns the Query field value.
func (o *CaseViewCreateAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CaseViewCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *CaseViewCreateAttributes) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseViewCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.NpRuleId != nil {
		toSerialize["np_rule_id"] = o.NpRuleId
	}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseViewCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name"`
		NpRuleId  *string `json:"np_rule_id,omitempty"`
		ProjectId *string `json:"project_id"`
		Query     *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "np_rule_id", "project_id", "query"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.NpRuleId = all.NpRuleId
	o.ProjectId = *all.ProjectId
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
