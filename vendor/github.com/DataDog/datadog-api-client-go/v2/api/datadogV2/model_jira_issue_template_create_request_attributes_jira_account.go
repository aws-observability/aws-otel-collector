// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueTemplateCreateRequestAttributesJiraAccount Reference to the Jira account
type JiraIssueTemplateCreateRequestAttributesJiraAccount struct {
	// The ID of the Jira account
	Id uuid.UUID `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueTemplateCreateRequestAttributesJiraAccount instantiates a new JiraIssueTemplateCreateRequestAttributesJiraAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueTemplateCreateRequestAttributesJiraAccount(id uuid.UUID) *JiraIssueTemplateCreateRequestAttributesJiraAccount {
	this := JiraIssueTemplateCreateRequestAttributesJiraAccount{}
	this.Id = id
	return &this
}

// NewJiraIssueTemplateCreateRequestAttributesJiraAccountWithDefaults instantiates a new JiraIssueTemplateCreateRequestAttributesJiraAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueTemplateCreateRequestAttributesJiraAccountWithDefaults() *JiraIssueTemplateCreateRequestAttributesJiraAccount {
	this := JiraIssueTemplateCreateRequestAttributesJiraAccount{}
	return &this
}

// GetId returns the Id field value.
func (o *JiraIssueTemplateCreateRequestAttributesJiraAccount) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateCreateRequestAttributesJiraAccount) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *JiraIssueTemplateCreateRequestAttributesJiraAccount) SetId(v uuid.UUID) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueTemplateCreateRequestAttributesJiraAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueTemplateCreateRequestAttributesJiraAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *uuid.UUID `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id"})
	} else {
		return err
	}
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
