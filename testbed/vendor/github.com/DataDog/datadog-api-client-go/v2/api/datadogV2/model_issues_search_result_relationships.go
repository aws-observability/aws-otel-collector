// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchResultRelationships Relationships between the search result and other resources.
type IssuesSearchResultRelationships struct {
	// Relationship between the search result and the corresponding issue.
	Issue *IssuesSearchResultIssueRelationship `json:"issue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssuesSearchResultRelationships instantiates a new IssuesSearchResultRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssuesSearchResultRelationships() *IssuesSearchResultRelationships {
	this := IssuesSearchResultRelationships{}
	return &this
}

// NewIssuesSearchResultRelationshipsWithDefaults instantiates a new IssuesSearchResultRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssuesSearchResultRelationshipsWithDefaults() *IssuesSearchResultRelationships {
	this := IssuesSearchResultRelationships{}
	return &this
}

// GetIssue returns the Issue field value if set, zero value otherwise.
func (o *IssuesSearchResultRelationships) GetIssue() IssuesSearchResultIssueRelationship {
	if o == nil || o.Issue == nil {
		var ret IssuesSearchResultIssueRelationship
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchResultRelationships) GetIssueOk() (*IssuesSearchResultIssueRelationship, bool) {
	if o == nil || o.Issue == nil {
		return nil, false
	}
	return o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *IssuesSearchResultRelationships) HasIssue() bool {
	return o != nil && o.Issue != nil
}

// SetIssue gets a reference to the given IssuesSearchResultIssueRelationship and assigns it to the Issue field.
func (o *IssuesSearchResultRelationships) SetIssue(v IssuesSearchResultIssueRelationship) {
	o.Issue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssuesSearchResultRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Issue != nil {
		toSerialize["issue"] = o.Issue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssuesSearchResultRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Issue *IssuesSearchResultIssueRelationship `json:"issue,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"issue"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Issue != nil && all.Issue.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Issue = all.Issue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
