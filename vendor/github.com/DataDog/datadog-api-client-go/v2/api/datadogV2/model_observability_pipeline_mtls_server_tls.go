// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMtlsServerTls Configuration for enabling TLS encryption between the pipeline component and external connecting clients.
type ObservabilityPipelineMtlsServerTls struct {
	// Path to the Certificate Authority (CA) file used to validate connecting clients' TLS certificates.
	CaFile *string `json:"ca_file,omitempty"`
	// Path to the TLS server certificate file used to used to identify the pipeline component to connecting clients.
	CrtFile string `json:"crt_file"`
	// Path to the private key file associated with the TLS server certificate.
	KeyFile *string `json:"key_file,omitempty"`
	// Name of the environment variable or secret that holds the passphrase for the private key file.
	KeyPassKey *string `json:"key_pass_key,omitempty"`
	// When `true`, requires client connections to present a valid certificate, enabling mutual TLS authentication.
	VerifyCertificate *bool `json:"verify_certificate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineMtlsServerTls instantiates a new ObservabilityPipelineMtlsServerTls object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineMtlsServerTls(crtFile string) *ObservabilityPipelineMtlsServerTls {
	this := ObservabilityPipelineMtlsServerTls{}
	this.CrtFile = crtFile
	return &this
}

// NewObservabilityPipelineMtlsServerTlsWithDefaults instantiates a new ObservabilityPipelineMtlsServerTls object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineMtlsServerTlsWithDefaults() *ObservabilityPipelineMtlsServerTls {
	this := ObservabilityPipelineMtlsServerTls{}
	return &this
}

// GetCaFile returns the CaFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTls) GetCaFile() string {
	if o == nil || o.CaFile == nil {
		var ret string
		return ret
	}
	return *o.CaFile
}

// GetCaFileOk returns a tuple with the CaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTls) GetCaFileOk() (*string, bool) {
	if o == nil || o.CaFile == nil {
		return nil, false
	}
	return o.CaFile, true
}

// HasCaFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTls) HasCaFile() bool {
	return o != nil && o.CaFile != nil
}

// SetCaFile gets a reference to the given string and assigns it to the CaFile field.
func (o *ObservabilityPipelineMtlsServerTls) SetCaFile(v string) {
	o.CaFile = &v
}

// GetCrtFile returns the CrtFile field value.
func (o *ObservabilityPipelineMtlsServerTls) GetCrtFile() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CrtFile
}

// GetCrtFileOk returns a tuple with the CrtFile field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTls) GetCrtFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrtFile, true
}

// SetCrtFile sets field value.
func (o *ObservabilityPipelineMtlsServerTls) SetCrtFile(v string) {
	o.CrtFile = v
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTls) GetKeyFile() string {
	if o == nil || o.KeyFile == nil {
		var ret string
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTls) GetKeyFileOk() (*string, bool) {
	if o == nil || o.KeyFile == nil {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTls) HasKeyFile() bool {
	return o != nil && o.KeyFile != nil
}

// SetKeyFile gets a reference to the given string and assigns it to the KeyFile field.
func (o *ObservabilityPipelineMtlsServerTls) SetKeyFile(v string) {
	o.KeyFile = &v
}

// GetKeyPassKey returns the KeyPassKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTls) GetKeyPassKey() string {
	if o == nil || o.KeyPassKey == nil {
		var ret string
		return ret
	}
	return *o.KeyPassKey
}

// GetKeyPassKeyOk returns a tuple with the KeyPassKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTls) GetKeyPassKeyOk() (*string, bool) {
	if o == nil || o.KeyPassKey == nil {
		return nil, false
	}
	return o.KeyPassKey, true
}

// HasKeyPassKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTls) HasKeyPassKey() bool {
	return o != nil && o.KeyPassKey != nil
}

// SetKeyPassKey gets a reference to the given string and assigns it to the KeyPassKey field.
func (o *ObservabilityPipelineMtlsServerTls) SetKeyPassKey(v string) {
	o.KeyPassKey = &v
}

// GetVerifyCertificate returns the VerifyCertificate field value if set, zero value otherwise.
func (o *ObservabilityPipelineMtlsServerTls) GetVerifyCertificate() bool {
	if o == nil || o.VerifyCertificate == nil {
		var ret bool
		return ret
	}
	return *o.VerifyCertificate
}

// GetVerifyCertificateOk returns a tuple with the VerifyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineMtlsServerTls) GetVerifyCertificateOk() (*bool, bool) {
	if o == nil || o.VerifyCertificate == nil {
		return nil, false
	}
	return o.VerifyCertificate, true
}

// HasVerifyCertificate returns a boolean if a field has been set.
func (o *ObservabilityPipelineMtlsServerTls) HasVerifyCertificate() bool {
	return o != nil && o.VerifyCertificate != nil
}

// SetVerifyCertificate gets a reference to the given bool and assigns it to the VerifyCertificate field.
func (o *ObservabilityPipelineMtlsServerTls) SetVerifyCertificate(v bool) {
	o.VerifyCertificate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineMtlsServerTls) MarshalJSON() ([]byte, error) {
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
	if o.VerifyCertificate != nil {
		toSerialize["verify_certificate"] = o.VerifyCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineMtlsServerTls) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaFile            *string `json:"ca_file,omitempty"`
		CrtFile           *string `json:"crt_file"`
		KeyFile           *string `json:"key_file,omitempty"`
		KeyPassKey        *string `json:"key_pass_key,omitempty"`
		VerifyCertificate *bool   `json:"verify_certificate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CrtFile == nil {
		return fmt.Errorf("required field crt_file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ca_file", "crt_file", "key_file", "key_pass_key", "verify_certificate"})
	} else {
		return err
	}
	o.CaFile = all.CaFile
	o.CrtFile = *all.CrtFile
	o.KeyFile = all.KeyFile
	o.KeyPassKey = all.KeyPassKey
	o.VerifyCertificate = all.VerifyCertificate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
