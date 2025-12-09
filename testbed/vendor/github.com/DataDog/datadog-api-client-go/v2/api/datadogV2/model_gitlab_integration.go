// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitlabIntegration The definition of the `GitlabIntegration` object.
type GitlabIntegration struct {
	// The definition of the `GitlabCredentials` object.
	Credentials GitlabCredentials `json:"credentials"`
	// The definition of the `GitlabIntegrationType` object.
	Type GitlabIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitlabIntegration instantiates a new GitlabIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitlabIntegration(credentials GitlabCredentials, typeVar GitlabIntegrationType) *GitlabIntegration {
	this := GitlabIntegration{}
	this.Credentials = credentials
	this.Type = typeVar
	return &this
}

// NewGitlabIntegrationWithDefaults instantiates a new GitlabIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitlabIntegrationWithDefaults() *GitlabIntegration {
	this := GitlabIntegration{}
	return &this
}

// GetCredentials returns the Credentials field value.
func (o *GitlabIntegration) GetCredentials() GitlabCredentials {
	if o == nil {
		var ret GitlabCredentials
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *GitlabIntegration) GetCredentialsOk() (*GitlabCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value.
func (o *GitlabIntegration) SetCredentials(v GitlabCredentials) {
	o.Credentials = v
}

// GetType returns the Type field value.
func (o *GitlabIntegration) GetType() GitlabIntegrationType {
	if o == nil {
		var ret GitlabIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GitlabIntegration) GetTypeOk() (*GitlabIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GitlabIntegration) SetType(v GitlabIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitlabIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["credentials"] = o.Credentials
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitlabIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *GitlabCredentials     `json:"credentials"`
		Type        *GitlabIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Credentials == nil {
		return fmt.Errorf("required field credentials missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"credentials", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Credentials = *all.Credentials
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
