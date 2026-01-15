// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateJiraIssueRequestData Data of the Jira issue to create.
type CreateJiraIssueRequestData struct {
	// Attributes of the Jira issue to create.
	Attributes *CreateJiraIssueRequestDataAttributes `json:"attributes,omitempty"`
	// Relationships of the Jira issue to create.
	Relationships *CreateJiraIssueRequestDataRelationships `json:"relationships,omitempty"`
	// Jira issues resource type.
	Type JiraIssuesDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateJiraIssueRequestData instantiates a new CreateJiraIssueRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateJiraIssueRequestData(typeVar JiraIssuesDataType) *CreateJiraIssueRequestData {
	this := CreateJiraIssueRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreateJiraIssueRequestDataWithDefaults instantiates a new CreateJiraIssueRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateJiraIssueRequestDataWithDefaults() *CreateJiraIssueRequestData {
	this := CreateJiraIssueRequestData{}
	var typeVar JiraIssuesDataType = JIRAISSUESDATATYPE_JIRA_ISSUES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestData) GetAttributes() CreateJiraIssueRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateJiraIssueRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestData) GetAttributesOk() (*CreateJiraIssueRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateJiraIssueRequestDataAttributes and assigns it to the Attributes field.
func (o *CreateJiraIssueRequestData) SetAttributes(v CreateJiraIssueRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestData) GetRelationships() CreateJiraIssueRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CreateJiraIssueRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestData) GetRelationshipsOk() (*CreateJiraIssueRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given CreateJiraIssueRequestDataRelationships and assigns it to the Relationships field.
func (o *CreateJiraIssueRequestData) SetRelationships(v CreateJiraIssueRequestDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *CreateJiraIssueRequestData) GetType() JiraIssuesDataType {
	if o == nil {
		var ret JiraIssuesDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestData) GetTypeOk() (*JiraIssuesDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateJiraIssueRequestData) SetType(v JiraIssuesDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateJiraIssueRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateJiraIssueRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *CreateJiraIssueRequestDataAttributes    `json:"attributes,omitempty"`
		Relationships *CreateJiraIssueRequestDataRelationships `json:"relationships,omitempty"`
		Type          *JiraIssuesDataType                      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
