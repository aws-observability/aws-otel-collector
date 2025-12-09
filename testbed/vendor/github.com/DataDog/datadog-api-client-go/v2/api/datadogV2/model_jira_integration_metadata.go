// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIntegrationMetadata Incident integration metadata for the Jira integration.
type JiraIntegrationMetadata struct {
	// Array of Jira issues in this integration metadata.
	Issues []JiraIntegrationMetadataIssuesItem `json:"issues"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIntegrationMetadata instantiates a new JiraIntegrationMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIntegrationMetadata(issues []JiraIntegrationMetadataIssuesItem) *JiraIntegrationMetadata {
	this := JiraIntegrationMetadata{}
	this.Issues = issues
	return &this
}

// NewJiraIntegrationMetadataWithDefaults instantiates a new JiraIntegrationMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIntegrationMetadataWithDefaults() *JiraIntegrationMetadata {
	this := JiraIntegrationMetadata{}
	return &this
}

// GetIssues returns the Issues field value.
func (o *JiraIntegrationMetadata) GetIssues() []JiraIntegrationMetadataIssuesItem {
	if o == nil {
		var ret []JiraIntegrationMetadataIssuesItem
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value
// and a boolean to check if the value has been set.
func (o *JiraIntegrationMetadata) GetIssuesOk() (*[]JiraIntegrationMetadataIssuesItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issues, true
}

// SetIssues sets field value.
func (o *JiraIntegrationMetadata) SetIssues(v []JiraIntegrationMetadataIssuesItem) {
	o.Issues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIntegrationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["issues"] = o.Issues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIntegrationMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Issues *[]JiraIntegrationMetadataIssuesItem `json:"issues"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Issues == nil {
		return fmt.Errorf("required field issues missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"issues"})
	} else {
		return err
	}
	o.Issues = *all.Issues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
