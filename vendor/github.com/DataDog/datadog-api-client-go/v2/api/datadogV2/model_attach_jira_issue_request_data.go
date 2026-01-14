// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachJiraIssueRequestData Data of the Jira issue to attach security findings to.
type AttachJiraIssueRequestData struct {
	// Attributes of the Jira issue to attach security findings to.
	Attributes *AttachJiraIssueRequestDataAttributes `json:"attributes,omitempty"`
	// Relationships of the Jira issue to attach security findings to.
	Relationships *AttachJiraIssueRequestDataRelationships `json:"relationships,omitempty"`
	// Jira issues resource type.
	Type JiraIssuesDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachJiraIssueRequestData instantiates a new AttachJiraIssueRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachJiraIssueRequestData(typeVar JiraIssuesDataType) *AttachJiraIssueRequestData {
	this := AttachJiraIssueRequestData{}
	this.Type = typeVar
	return &this
}

// NewAttachJiraIssueRequestDataWithDefaults instantiates a new AttachJiraIssueRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachJiraIssueRequestDataWithDefaults() *AttachJiraIssueRequestData {
	this := AttachJiraIssueRequestData{}
	var typeVar JiraIssuesDataType = JIRAISSUESDATATYPE_JIRA_ISSUES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AttachJiraIssueRequestData) GetAttributes() AttachJiraIssueRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret AttachJiraIssueRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachJiraIssueRequestData) GetAttributesOk() (*AttachJiraIssueRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AttachJiraIssueRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AttachJiraIssueRequestDataAttributes and assigns it to the Attributes field.
func (o *AttachJiraIssueRequestData) SetAttributes(v AttachJiraIssueRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AttachJiraIssueRequestData) GetRelationships() AttachJiraIssueRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AttachJiraIssueRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachJiraIssueRequestData) GetRelationshipsOk() (*AttachJiraIssueRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AttachJiraIssueRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given AttachJiraIssueRequestDataRelationships and assigns it to the Relationships field.
func (o *AttachJiraIssueRequestData) SetRelationships(v AttachJiraIssueRequestDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *AttachJiraIssueRequestData) GetType() JiraIssuesDataType {
	if o == nil {
		var ret JiraIssuesDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachJiraIssueRequestData) GetTypeOk() (*JiraIssuesDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AttachJiraIssueRequestData) SetType(v JiraIssuesDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachJiraIssueRequestData) MarshalJSON() ([]byte, error) {
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
func (o *AttachJiraIssueRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *AttachJiraIssueRequestDataAttributes    `json:"attributes,omitempty"`
		Relationships *AttachJiraIssueRequestDataRelationships `json:"relationships,omitempty"`
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
