// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetConfigurationLayer Configuration information organized by layers.
type FleetConfigurationLayer struct {
	// The final compiled configuration.
	CompiledConfiguration *string `json:"compiled_configuration,omitempty"`
	// Configuration from environment variables.
	EnvConfiguration *string `json:"env_configuration,omitempty"`
	// Configuration from files.
	FileConfiguration *string `json:"file_configuration,omitempty"`
	// Parsed configuration output.
	ParsedConfiguration *string `json:"parsed_configuration,omitempty"`
	// Remote configuration settings.
	RemoteConfiguration *string `json:"remote_configuration,omitempty"`
	// Runtime configuration.
	RuntimeConfiguration *string `json:"runtime_configuration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetConfigurationLayer instantiates a new FleetConfigurationLayer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetConfigurationLayer() *FleetConfigurationLayer {
	this := FleetConfigurationLayer{}
	return &this
}

// NewFleetConfigurationLayerWithDefaults instantiates a new FleetConfigurationLayer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetConfigurationLayerWithDefaults() *FleetConfigurationLayer {
	this := FleetConfigurationLayer{}
	return &this
}

// GetCompiledConfiguration returns the CompiledConfiguration field value if set, zero value otherwise.
func (o *FleetConfigurationLayer) GetCompiledConfiguration() string {
	if o == nil || o.CompiledConfiguration == nil {
		var ret string
		return ret
	}
	return *o.CompiledConfiguration
}

// GetCompiledConfigurationOk returns a tuple with the CompiledConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationLayer) GetCompiledConfigurationOk() (*string, bool) {
	if o == nil || o.CompiledConfiguration == nil {
		return nil, false
	}
	return o.CompiledConfiguration, true
}

// HasCompiledConfiguration returns a boolean if a field has been set.
func (o *FleetConfigurationLayer) HasCompiledConfiguration() bool {
	return o != nil && o.CompiledConfiguration != nil
}

// SetCompiledConfiguration gets a reference to the given string and assigns it to the CompiledConfiguration field.
func (o *FleetConfigurationLayer) SetCompiledConfiguration(v string) {
	o.CompiledConfiguration = &v
}

// GetEnvConfiguration returns the EnvConfiguration field value if set, zero value otherwise.
func (o *FleetConfigurationLayer) GetEnvConfiguration() string {
	if o == nil || o.EnvConfiguration == nil {
		var ret string
		return ret
	}
	return *o.EnvConfiguration
}

// GetEnvConfigurationOk returns a tuple with the EnvConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationLayer) GetEnvConfigurationOk() (*string, bool) {
	if o == nil || o.EnvConfiguration == nil {
		return nil, false
	}
	return o.EnvConfiguration, true
}

// HasEnvConfiguration returns a boolean if a field has been set.
func (o *FleetConfigurationLayer) HasEnvConfiguration() bool {
	return o != nil && o.EnvConfiguration != nil
}

// SetEnvConfiguration gets a reference to the given string and assigns it to the EnvConfiguration field.
func (o *FleetConfigurationLayer) SetEnvConfiguration(v string) {
	o.EnvConfiguration = &v
}

// GetFileConfiguration returns the FileConfiguration field value if set, zero value otherwise.
func (o *FleetConfigurationLayer) GetFileConfiguration() string {
	if o == nil || o.FileConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FileConfiguration
}

// GetFileConfigurationOk returns a tuple with the FileConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationLayer) GetFileConfigurationOk() (*string, bool) {
	if o == nil || o.FileConfiguration == nil {
		return nil, false
	}
	return o.FileConfiguration, true
}

// HasFileConfiguration returns a boolean if a field has been set.
func (o *FleetConfigurationLayer) HasFileConfiguration() bool {
	return o != nil && o.FileConfiguration != nil
}

// SetFileConfiguration gets a reference to the given string and assigns it to the FileConfiguration field.
func (o *FleetConfigurationLayer) SetFileConfiguration(v string) {
	o.FileConfiguration = &v
}

// GetParsedConfiguration returns the ParsedConfiguration field value if set, zero value otherwise.
func (o *FleetConfigurationLayer) GetParsedConfiguration() string {
	if o == nil || o.ParsedConfiguration == nil {
		var ret string
		return ret
	}
	return *o.ParsedConfiguration
}

// GetParsedConfigurationOk returns a tuple with the ParsedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationLayer) GetParsedConfigurationOk() (*string, bool) {
	if o == nil || o.ParsedConfiguration == nil {
		return nil, false
	}
	return o.ParsedConfiguration, true
}

// HasParsedConfiguration returns a boolean if a field has been set.
func (o *FleetConfigurationLayer) HasParsedConfiguration() bool {
	return o != nil && o.ParsedConfiguration != nil
}

// SetParsedConfiguration gets a reference to the given string and assigns it to the ParsedConfiguration field.
func (o *FleetConfigurationLayer) SetParsedConfiguration(v string) {
	o.ParsedConfiguration = &v
}

// GetRemoteConfiguration returns the RemoteConfiguration field value if set, zero value otherwise.
func (o *FleetConfigurationLayer) GetRemoteConfiguration() string {
	if o == nil || o.RemoteConfiguration == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfiguration
}

// GetRemoteConfigurationOk returns a tuple with the RemoteConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationLayer) GetRemoteConfigurationOk() (*string, bool) {
	if o == nil || o.RemoteConfiguration == nil {
		return nil, false
	}
	return o.RemoteConfiguration, true
}

// HasRemoteConfiguration returns a boolean if a field has been set.
func (o *FleetConfigurationLayer) HasRemoteConfiguration() bool {
	return o != nil && o.RemoteConfiguration != nil
}

// SetRemoteConfiguration gets a reference to the given string and assigns it to the RemoteConfiguration field.
func (o *FleetConfigurationLayer) SetRemoteConfiguration(v string) {
	o.RemoteConfiguration = &v
}

// GetRuntimeConfiguration returns the RuntimeConfiguration field value if set, zero value otherwise.
func (o *FleetConfigurationLayer) GetRuntimeConfiguration() string {
	if o == nil || o.RuntimeConfiguration == nil {
		var ret string
		return ret
	}
	return *o.RuntimeConfiguration
}

// GetRuntimeConfigurationOk returns a tuple with the RuntimeConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationLayer) GetRuntimeConfigurationOk() (*string, bool) {
	if o == nil || o.RuntimeConfiguration == nil {
		return nil, false
	}
	return o.RuntimeConfiguration, true
}

// HasRuntimeConfiguration returns a boolean if a field has been set.
func (o *FleetConfigurationLayer) HasRuntimeConfiguration() bool {
	return o != nil && o.RuntimeConfiguration != nil
}

// SetRuntimeConfiguration gets a reference to the given string and assigns it to the RuntimeConfiguration field.
func (o *FleetConfigurationLayer) SetRuntimeConfiguration(v string) {
	o.RuntimeConfiguration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetConfigurationLayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompiledConfiguration != nil {
		toSerialize["compiled_configuration"] = o.CompiledConfiguration
	}
	if o.EnvConfiguration != nil {
		toSerialize["env_configuration"] = o.EnvConfiguration
	}
	if o.FileConfiguration != nil {
		toSerialize["file_configuration"] = o.FileConfiguration
	}
	if o.ParsedConfiguration != nil {
		toSerialize["parsed_configuration"] = o.ParsedConfiguration
	}
	if o.RemoteConfiguration != nil {
		toSerialize["remote_configuration"] = o.RemoteConfiguration
	}
	if o.RuntimeConfiguration != nil {
		toSerialize["runtime_configuration"] = o.RuntimeConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetConfigurationLayer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompiledConfiguration *string `json:"compiled_configuration,omitempty"`
		EnvConfiguration      *string `json:"env_configuration,omitempty"`
		FileConfiguration     *string `json:"file_configuration,omitempty"`
		ParsedConfiguration   *string `json:"parsed_configuration,omitempty"`
		RemoteConfiguration   *string `json:"remote_configuration,omitempty"`
		RuntimeConfiguration  *string `json:"runtime_configuration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compiled_configuration", "env_configuration", "file_configuration", "parsed_configuration", "remote_configuration", "runtime_configuration"})
	} else {
		return err
	}
	o.CompiledConfiguration = all.CompiledConfiguration
	o.EnvConfiguration = all.EnvConfiguration
	o.FileConfiguration = all.FileConfiguration
	o.ParsedConfiguration = all.ParsedConfiguration
	o.RemoteConfiguration = all.RemoteConfiguration
	o.RuntimeConfiguration = all.RuntimeConfiguration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
