// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueTemplateDataRelationships Relationships of a Jira issue template
type JiraIssueTemplateDataRelationships struct {
	// Relationship to a Jira account
	JiraAccount JiraAccountRelationship `json:"jira-account"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueTemplateDataRelationships instantiates a new JiraIssueTemplateDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueTemplateDataRelationships(jiraAccount JiraAccountRelationship) *JiraIssueTemplateDataRelationships {
	this := JiraIssueTemplateDataRelationships{}
	this.JiraAccount = jiraAccount
	return &this
}

// NewJiraIssueTemplateDataRelationshipsWithDefaults instantiates a new JiraIssueTemplateDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueTemplateDataRelationshipsWithDefaults() *JiraIssueTemplateDataRelationships {
	this := JiraIssueTemplateDataRelationships{}
	return &this
}

// GetJiraAccount returns the JiraAccount field value.
func (o *JiraIssueTemplateDataRelationships) GetJiraAccount() JiraAccountRelationship {
	if o == nil {
		var ret JiraAccountRelationship
		return ret
	}
	return o.JiraAccount
}

// GetJiraAccountOk returns a tuple with the JiraAccount field value
// and a boolean to check if the value has been set.
func (o *JiraIssueTemplateDataRelationships) GetJiraAccountOk() (*JiraAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraAccount, true
}

// SetJiraAccount sets field value.
func (o *JiraIssueTemplateDataRelationships) SetJiraAccount(v JiraAccountRelationship) {
	o.JiraAccount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueTemplateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["jira-account"] = o.JiraAccount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueTemplateDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JiraAccount *JiraAccountRelationship `json:"jira-account"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JiraAccount == nil {
		return fmt.Errorf("required field jira-account missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"jira-account"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.JiraAccount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JiraAccount = *all.JiraAccount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
