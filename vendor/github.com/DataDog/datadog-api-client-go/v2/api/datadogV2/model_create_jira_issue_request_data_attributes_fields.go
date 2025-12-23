// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateJiraIssueRequestDataAttributesFields Custom fields of the Jira issue to create. For the list of available fields, see [Jira documentation](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-createmeta-projectidorkey-issuetypes-issuetypeid-get).
type CreateJiraIssueRequestDataAttributesFields struct {
	//
	Fields interface{} `json:"fields,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateJiraIssueRequestDataAttributesFields instantiates a new CreateJiraIssueRequestDataAttributesFields object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateJiraIssueRequestDataAttributesFields() *CreateJiraIssueRequestDataAttributesFields {
	this := CreateJiraIssueRequestDataAttributesFields{}
	return &this
}

// NewCreateJiraIssueRequestDataAttributesFieldsWithDefaults instantiates a new CreateJiraIssueRequestDataAttributesFields object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateJiraIssueRequestDataAttributesFieldsWithDefaults() *CreateJiraIssueRequestDataAttributesFields {
	this := CreateJiraIssueRequestDataAttributesFields{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestDataAttributesFields) GetFields() interface{} {
	if o == nil || o.Fields == nil {
		var ret interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestDataAttributesFields) GetFieldsOk() (*interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestDataAttributesFields) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given interface{} and assigns it to the Fields field.
func (o *CreateJiraIssueRequestDataAttributesFields) SetFields(v interface{}) {
	o.Fields = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateJiraIssueRequestDataAttributesFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateJiraIssueRequestDataAttributesFields) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields interface{} `json:"fields,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields"})
	} else {
		return err
	}
	o.Fields = all.Fields

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
