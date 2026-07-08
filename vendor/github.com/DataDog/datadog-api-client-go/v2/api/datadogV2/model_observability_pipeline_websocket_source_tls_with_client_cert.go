// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceTlsWithClientCert TLS configuration that enables encryption and presents a client certificate for mutual TLS authentication.
type ObservabilityPipelineWebsocketSourceTlsWithClientCert struct {
	// Path to the Certificate Authority (CA) file used to validate the remote server's TLS certificate.
	CaFile *string `json:"ca_file,omitempty"`
	// Path to the TLS client certificate file used to identify this source to the remote server.
	CrtFile string `json:"crt_file"`
	// Path to the private key file associated with the client certificate.
	KeyFile *string `json:"key_file,omitempty"`
	// Name of the environment variable or secret that holds the passphrase for the private key file.
	KeyPassKey *string `json:"key_pass_key,omitempty"`
	// TLS mode. Must be `with_client_cert`.
	Mode ObservabilityPipelineWebsocketSourceTlsWithClientCertMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineWebsocketSourceTlsWithClientCert instantiates a new ObservabilityPipelineWebsocketSourceTlsWithClientCert object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineWebsocketSourceTlsWithClientCert(crtFile string, mode ObservabilityPipelineWebsocketSourceTlsWithClientCertMode) *ObservabilityPipelineWebsocketSourceTlsWithClientCert {
	this := ObservabilityPipelineWebsocketSourceTlsWithClientCert{}
	this.CrtFile = crtFile
	this.Mode = mode
	return &this
}

// NewObservabilityPipelineWebsocketSourceTlsWithClientCertWithDefaults instantiates a new ObservabilityPipelineWebsocketSourceTlsWithClientCert object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineWebsocketSourceTlsWithClientCertWithDefaults() *ObservabilityPipelineWebsocketSourceTlsWithClientCert {
	this := ObservabilityPipelineWebsocketSourceTlsWithClientCert{}
	return &this
}

// GetCaFile returns the CaFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetCaFile() string {
	if o == nil || o.CaFile == nil {
		var ret string
		return ret
	}
	return *o.CaFile
}

// GetCaFileOk returns a tuple with the CaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetCaFileOk() (*string, bool) {
	if o == nil || o.CaFile == nil {
		return nil, false
	}
	return o.CaFile, true
}

// HasCaFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) HasCaFile() bool {
	return o != nil && o.CaFile != nil
}

// SetCaFile gets a reference to the given string and assigns it to the CaFile field.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) SetCaFile(v string) {
	o.CaFile = &v
}

// GetCrtFile returns the CrtFile field value.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetCrtFile() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CrtFile
}

// GetCrtFileOk returns a tuple with the CrtFile field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetCrtFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrtFile, true
}

// SetCrtFile sets field value.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) SetCrtFile(v string) {
	o.CrtFile = v
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetKeyFile() string {
	if o == nil || o.KeyFile == nil {
		var ret string
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetKeyFileOk() (*string, bool) {
	if o == nil || o.KeyFile == nil {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) HasKeyFile() bool {
	return o != nil && o.KeyFile != nil
}

// SetKeyFile gets a reference to the given string and assigns it to the KeyFile field.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) SetKeyFile(v string) {
	o.KeyFile = &v
}

// GetKeyPassKey returns the KeyPassKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetKeyPassKey() string {
	if o == nil || o.KeyPassKey == nil {
		var ret string
		return ret
	}
	return *o.KeyPassKey
}

// GetKeyPassKeyOk returns a tuple with the KeyPassKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetKeyPassKeyOk() (*string, bool) {
	if o == nil || o.KeyPassKey == nil {
		return nil, false
	}
	return o.KeyPassKey, true
}

// HasKeyPassKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) HasKeyPassKey() bool {
	return o != nil && o.KeyPassKey != nil
}

// SetKeyPassKey gets a reference to the given string and assigns it to the KeyPassKey field.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) SetKeyPassKey(v string) {
	o.KeyPassKey = &v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetMode() ObservabilityPipelineWebsocketSourceTlsWithClientCertMode {
	if o == nil {
		var ret ObservabilityPipelineWebsocketSourceTlsWithClientCertMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) GetModeOk() (*ObservabilityPipelineWebsocketSourceTlsWithClientCertMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) SetMode(v ObservabilityPipelineWebsocketSourceTlsWithClientCertMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineWebsocketSourceTlsWithClientCert) MarshalJSON() ([]byte, error) {
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
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineWebsocketSourceTlsWithClientCert) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaFile     *string                                                    `json:"ca_file,omitempty"`
		CrtFile    *string                                                    `json:"crt_file"`
		KeyFile    *string                                                    `json:"key_file,omitempty"`
		KeyPassKey *string                                                    `json:"key_pass_key,omitempty"`
		Mode       *ObservabilityPipelineWebsocketSourceTlsWithClientCertMode `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CrtFile == nil {
		return fmt.Errorf("required field crt_file missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ca_file", "crt_file", "key_file", "key_pass_key", "mode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CaFile = all.CaFile
	o.CrtFile = *all.CrtFile
	o.KeyFile = all.KeyFile
	o.KeyPassKey = all.KeyPassKey
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
