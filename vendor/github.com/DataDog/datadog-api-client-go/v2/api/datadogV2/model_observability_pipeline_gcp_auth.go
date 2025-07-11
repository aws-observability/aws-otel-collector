// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGcpAuth GCP credentials used to authenticate with Google Cloud Storage.
type ObservabilityPipelineGcpAuth struct {
	// Path to the GCP service account key file.
	CredentialsFile string `json:"credentials_file"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGcpAuth instantiates a new ObservabilityPipelineGcpAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGcpAuth(credentialsFile string) *ObservabilityPipelineGcpAuth {
	this := ObservabilityPipelineGcpAuth{}
	this.CredentialsFile = credentialsFile
	return &this
}

// NewObservabilityPipelineGcpAuthWithDefaults instantiates a new ObservabilityPipelineGcpAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGcpAuthWithDefaults() *ObservabilityPipelineGcpAuth {
	this := ObservabilityPipelineGcpAuth{}
	return &this
}

// GetCredentialsFile returns the CredentialsFile field value.
func (o *ObservabilityPipelineGcpAuth) GetCredentialsFile() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CredentialsFile
}

// GetCredentialsFileOk returns a tuple with the CredentialsFile field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGcpAuth) GetCredentialsFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialsFile, true
}

// SetCredentialsFile sets field value.
func (o *ObservabilityPipelineGcpAuth) SetCredentialsFile(v string) {
	o.CredentialsFile = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGcpAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["credentials_file"] = o.CredentialsFile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGcpAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CredentialsFile *string `json:"credentials_file"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CredentialsFile == nil {
		return fmt.Errorf("required field credentials_file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"credentials_file"})
	} else {
		return err
	}
	o.CredentialsFile = *all.CredentialsFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
