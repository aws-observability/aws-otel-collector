// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationJiraSyncDueDate Due date synchronization configuration for Jira integration.
type IntegrationJiraSyncDueDate struct {
	// The Jira field identifier used to store the due date.
	JiraFieldId *string `json:"jira_field_id,omitempty"`
	// The type of synchronization to apply for the due date field.
	SyncType *string `json:"sync_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationJiraSyncDueDate instantiates a new IntegrationJiraSyncDueDate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationJiraSyncDueDate() *IntegrationJiraSyncDueDate {
	this := IntegrationJiraSyncDueDate{}
	return &this
}

// NewIntegrationJiraSyncDueDateWithDefaults instantiates a new IntegrationJiraSyncDueDate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationJiraSyncDueDateWithDefaults() *IntegrationJiraSyncDueDate {
	this := IntegrationJiraSyncDueDate{}
	return &this
}

// GetJiraFieldId returns the JiraFieldId field value if set, zero value otherwise.
func (o *IntegrationJiraSyncDueDate) GetJiraFieldId() string {
	if o == nil || o.JiraFieldId == nil {
		var ret string
		return ret
	}
	return *o.JiraFieldId
}

// GetJiraFieldIdOk returns a tuple with the JiraFieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncDueDate) GetJiraFieldIdOk() (*string, bool) {
	if o == nil || o.JiraFieldId == nil {
		return nil, false
	}
	return o.JiraFieldId, true
}

// HasJiraFieldId returns a boolean if a field has been set.
func (o *IntegrationJiraSyncDueDate) HasJiraFieldId() bool {
	return o != nil && o.JiraFieldId != nil
}

// SetJiraFieldId gets a reference to the given string and assigns it to the JiraFieldId field.
func (o *IntegrationJiraSyncDueDate) SetJiraFieldId(v string) {
	o.JiraFieldId = &v
}

// GetSyncType returns the SyncType field value if set, zero value otherwise.
func (o *IntegrationJiraSyncDueDate) GetSyncType() string {
	if o == nil || o.SyncType == nil {
		var ret string
		return ret
	}
	return *o.SyncType
}

// GetSyncTypeOk returns a tuple with the SyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncDueDate) GetSyncTypeOk() (*string, bool) {
	if o == nil || o.SyncType == nil {
		return nil, false
	}
	return o.SyncType, true
}

// HasSyncType returns a boolean if a field has been set.
func (o *IntegrationJiraSyncDueDate) HasSyncType() bool {
	return o != nil && o.SyncType != nil
}

// SetSyncType gets a reference to the given string and assigns it to the SyncType field.
func (o *IntegrationJiraSyncDueDate) SetSyncType(v string) {
	o.SyncType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationJiraSyncDueDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.JiraFieldId != nil {
		toSerialize["jira_field_id"] = o.JiraFieldId
	}
	if o.SyncType != nil {
		toSerialize["sync_type"] = o.SyncType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationJiraSyncDueDate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JiraFieldId *string `json:"jira_field_id,omitempty"`
		SyncType    *string `json:"sync_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"jira_field_id", "sync_type"})
	} else {
		return err
	}
	o.JiraFieldId = all.JiraFieldId
	o.SyncType = all.SyncType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
