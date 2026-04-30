// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueTemplateUpdateRequestAttributes Attributes for updating a Jira issue template
type JiraIssueTemplateUpdateRequestAttributes struct {
	// Custom fields for the Jira issue template
	Fields map[string]interface{} `json:"fields,omitempty"`
	// The name of the issue template
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueTemplateUpdateRequestAttributes instantiates a new JiraIssueTemplateUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueTemplateUpdateRequestAttributes() *JiraIssueTemplateUpdateRequestAttributes {
	this := JiraIssueTemplateUpdateRequestAttributes{}
	return &this
}

// NewJiraIssueTemplateUpdateRequestAttributesWithDefaults instantiates a new JiraIssueTemplateUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueTemplateUpdateRequestAttributesWithDefaults() *JiraIssueTemplateUpdateRequestAttributes {
	this := JiraIssueTemplateUpdateRequestAttributes{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *JiraIssueTemplateUpdateRequestAttributes) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateUpdateRequestAttributes) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *JiraIssueTemplateUpdateRequestAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *JiraIssueTemplateUpdateRequestAttributes) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JiraIssueTemplateUpdateRequestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateUpdateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JiraIssueTemplateUpdateRequestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JiraIssueTemplateUpdateRequestAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueTemplateUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueTemplateUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields map[string]interface{} `json:"fields,omitempty"`
		Name   *string                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "name"})
	} else {
		return err
	}
	o.Fields = all.Fields
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
