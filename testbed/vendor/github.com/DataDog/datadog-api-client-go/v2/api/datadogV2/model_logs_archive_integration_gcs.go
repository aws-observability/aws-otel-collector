// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveIntegrationGCS The GCS archive's integration destination.
type LogsArchiveIntegrationGCS struct {
	// A client email.
	ClientEmail string `json:"client_email"`
	// A project ID.
	ProjectId *string `json:"project_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveIntegrationGCS instantiates a new LogsArchiveIntegrationGCS object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveIntegrationGCS(clientEmail string) *LogsArchiveIntegrationGCS {
	this := LogsArchiveIntegrationGCS{}
	this.ClientEmail = clientEmail
	return &this
}

// NewLogsArchiveIntegrationGCSWithDefaults instantiates a new LogsArchiveIntegrationGCS object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveIntegrationGCSWithDefaults() *LogsArchiveIntegrationGCS {
	this := LogsArchiveIntegrationGCS{}
	return &this
}

// GetClientEmail returns the ClientEmail field value.
func (o *LogsArchiveIntegrationGCS) GetClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationGCS) GetClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEmail, true
}

// SetClientEmail sets field value.
func (o *LogsArchiveIntegrationGCS) SetClientEmail(v string) {
	o.ClientEmail = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *LogsArchiveIntegrationGCS) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationGCS) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *LogsArchiveIntegrationGCS) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *LogsArchiveIntegrationGCS) SetProjectId(v string) {
	o.ProjectId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveIntegrationGCS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_email"] = o.ClientEmail
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveIntegrationGCS) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientEmail *string `json:"client_email"`
		ProjectId   *string `json:"project_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientEmail == nil {
		return fmt.Errorf("required field client_email missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_email", "project_id"})
	} else {
		return err
	}
	o.ClientEmail = *all.ClientEmail
	o.ProjectId = all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
