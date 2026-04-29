// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTls Configuration for enabling TLS encryption between the pipeline component and external services.
type ObservabilityPipelineTls struct {
	// Path to the Certificate Authority (CA) file used to validate the server’s TLS certificate.
	CaFile *string `json:"ca_file,omitempty"`
	// Path to the TLS client certificate file used to authenticate the pipeline component with upstream or downstream services.
	CrtFile string `json:"crt_file"`
	// Path to the private key file associated with the TLS client certificate. Used for mutual TLS authentication.
	KeyFile *string `json:"key_file,omitempty"`
	// Name of the environment variable or secret that holds the passphrase for the private key file.
	KeyPassKey *string `json:"key_pass_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineTls instantiates a new ObservabilityPipelineTls object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineTls(crtFile string) *ObservabilityPipelineTls {
	this := ObservabilityPipelineTls{}
	this.CrtFile = crtFile
	return &this
}

// NewObservabilityPipelineTlsWithDefaults instantiates a new ObservabilityPipelineTls object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineTlsWithDefaults() *ObservabilityPipelineTls {
	this := ObservabilityPipelineTls{}
	return &this
}

// GetCaFile returns the CaFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineTls) GetCaFile() string {
	if o == nil || o.CaFile == nil {
		var ret string
		return ret
	}
	return *o.CaFile
}

// GetCaFileOk returns a tuple with the CaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTls) GetCaFileOk() (*string, bool) {
	if o == nil || o.CaFile == nil {
		return nil, false
	}
	return o.CaFile, true
}

// HasCaFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineTls) HasCaFile() bool {
	return o != nil && o.CaFile != nil
}

// SetCaFile gets a reference to the given string and assigns it to the CaFile field.
func (o *ObservabilityPipelineTls) SetCaFile(v string) {
	o.CaFile = &v
}

// GetCrtFile returns the CrtFile field value.
func (o *ObservabilityPipelineTls) GetCrtFile() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CrtFile
}

// GetCrtFileOk returns a tuple with the CrtFile field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTls) GetCrtFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrtFile, true
}

// SetCrtFile sets field value.
func (o *ObservabilityPipelineTls) SetCrtFile(v string) {
	o.CrtFile = v
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineTls) GetKeyFile() string {
	if o == nil || o.KeyFile == nil {
		var ret string
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTls) GetKeyFileOk() (*string, bool) {
	if o == nil || o.KeyFile == nil {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineTls) HasKeyFile() bool {
	return o != nil && o.KeyFile != nil
}

// SetKeyFile gets a reference to the given string and assigns it to the KeyFile field.
func (o *ObservabilityPipelineTls) SetKeyFile(v string) {
	o.KeyFile = &v
}

// GetKeyPassKey returns the KeyPassKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineTls) GetKeyPassKey() string {
	if o == nil || o.KeyPassKey == nil {
		var ret string
		return ret
	}
	return *o.KeyPassKey
}

// GetKeyPassKeyOk returns a tuple with the KeyPassKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineTls) GetKeyPassKeyOk() (*string, bool) {
	if o == nil || o.KeyPassKey == nil {
		return nil, false
	}
	return o.KeyPassKey, true
}

// HasKeyPassKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineTls) HasKeyPassKey() bool {
	return o != nil && o.KeyPassKey != nil
}

// SetKeyPassKey gets a reference to the given string and assigns it to the KeyPassKey field.
func (o *ObservabilityPipelineTls) SetKeyPassKey(v string) {
	o.KeyPassKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineTls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CaFile != nil {
		toSerialize["ca_file"] = o.CaFile
	}
	toSerialize["crt_file"] = o.CrtFile
	if o.KeyFile != nil {
		toSerialize["key_file"] = o.KeyFile
	}
	if o.KeyPassKey != nil {
		toSerialize["key_pass_key"] = o.KeyPassKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineTls) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaFile     *string `json:"ca_file,omitempty"`
		CrtFile    *string `json:"crt_file"`
		KeyFile    *string `json:"key_file,omitempty"`
		KeyPassKey *string `json:"key_pass_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CrtFile == nil {
		return fmt.Errorf("required field crt_file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ca_file", "crt_file", "key_file", "key_pass_key"})
	} else {
		return err
	}
	o.CaFile = all.CaFile
	o.CrtFile = *all.CrtFile
	o.KeyFile = all.KeyFile
	o.KeyPassKey = all.KeyPassKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
