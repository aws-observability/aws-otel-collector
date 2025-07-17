// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAwsAuth AWS authentication credentials used for accessing AWS services such as S3.
// If omitted, the system’s default credentials are used (for example, the IAM role and environment variables).
type ObservabilityPipelineAwsAuth struct {
	// The Amazon Resource Name (ARN) of the role to assume.
	AssumeRole *string `json:"assume_role,omitempty"`
	// A unique identifier for cross-account role assumption.
	ExternalId *string `json:"external_id,omitempty"`
	// A session identifier used for logging and tracing the assumed role session.
	SessionName *string `json:"session_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAwsAuth instantiates a new ObservabilityPipelineAwsAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAwsAuth() *ObservabilityPipelineAwsAuth {
	this := ObservabilityPipelineAwsAuth{}
	return &this
}

// NewObservabilityPipelineAwsAuthWithDefaults instantiates a new ObservabilityPipelineAwsAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAwsAuthWithDefaults() *ObservabilityPipelineAwsAuth {
	this := ObservabilityPipelineAwsAuth{}
	return &this
}

// GetAssumeRole returns the AssumeRole field value if set, zero value otherwise.
func (o *ObservabilityPipelineAwsAuth) GetAssumeRole() string {
	if o == nil || o.AssumeRole == nil {
		var ret string
		return ret
	}
	return *o.AssumeRole
}

// GetAssumeRoleOk returns a tuple with the AssumeRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAwsAuth) GetAssumeRoleOk() (*string, bool) {
	if o == nil || o.AssumeRole == nil {
		return nil, false
	}
	return o.AssumeRole, true
}

// HasAssumeRole returns a boolean if a field has been set.
func (o *ObservabilityPipelineAwsAuth) HasAssumeRole() bool {
	return o != nil && o.AssumeRole != nil
}

// SetAssumeRole gets a reference to the given string and assigns it to the AssumeRole field.
func (o *ObservabilityPipelineAwsAuth) SetAssumeRole(v string) {
	o.AssumeRole = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ObservabilityPipelineAwsAuth) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAwsAuth) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAwsAuth) HasExternalId() bool {
	return o != nil && o.ExternalId != nil
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ObservabilityPipelineAwsAuth) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *ObservabilityPipelineAwsAuth) GetSessionName() string {
	if o == nil || o.SessionName == nil {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAwsAuth) GetSessionNameOk() (*string, bool) {
	if o == nil || o.SessionName == nil {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *ObservabilityPipelineAwsAuth) HasSessionName() bool {
	return o != nil && o.SessionName != nil
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *ObservabilityPipelineAwsAuth) SetSessionName(v string) {
	o.SessionName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAwsAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssumeRole != nil {
		toSerialize["assume_role"] = o.AssumeRole
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.SessionName != nil {
		toSerialize["session_name"] = o.SessionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAwsAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssumeRole  *string `json:"assume_role,omitempty"`
		ExternalId  *string `json:"external_id,omitempty"`
		SessionName *string `json:"session_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assume_role", "external_id", "session_name"})
	} else {
		return err
	}
	o.AssumeRole = all.AssumeRole
	o.ExternalId = all.ExternalId
	o.SessionName = all.SessionName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
