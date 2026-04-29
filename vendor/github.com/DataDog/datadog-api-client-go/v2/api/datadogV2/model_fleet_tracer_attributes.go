// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetTracerAttributes Attributes of a fleet tracer representing a service instance reporting telemetry.
type FleetTracerAttributes struct {
	// The environment the tracer is reporting from.
	Env *string `json:"env,omitempty"`
	// The hostname where the tracer is running.
	Hostname *string `json:"hostname,omitempty"`
	// The programming language of the traced application.
	Language *string `json:"language,omitempty"`
	// The version of the programming language runtime.
	LanguageVersion *string `json:"language_version,omitempty"`
	// The remote configuration status of the tracer.
	RemoteConfigStatus *string `json:"remote_config_status,omitempty"`
	// Runtime identifiers for the tracer instances.
	RuntimeIds []string `json:"runtime_ids,omitempty"`
	// The telemetry-derived service name reported by the tracer.
	Service *string `json:"service,omitempty"`
	// The service hostname reported by the tracer.
	ServiceHostname *string `json:"service_hostname,omitempty"`
	// The version of the traced service.
	ServiceVersion *string `json:"service_version,omitempty"`
	// The version of the Datadog tracer library.
	TracerVersion *string `json:"tracer_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetTracerAttributes instantiates a new FleetTracerAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetTracerAttributes() *FleetTracerAttributes {
	this := FleetTracerAttributes{}
	return &this
}

// NewFleetTracerAttributesWithDefaults instantiates a new FleetTracerAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetTracerAttributesWithDefaults() *FleetTracerAttributes {
	this := FleetTracerAttributes{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *FleetTracerAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetTracerAttributes) SetHostname(v string) {
	o.Hostname = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *FleetTracerAttributes) SetLanguage(v string) {
	o.Language = &v
}

// GetLanguageVersion returns the LanguageVersion field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetLanguageVersion() string {
	if o == nil || o.LanguageVersion == nil {
		var ret string
		return ret
	}
	return *o.LanguageVersion
}

// GetLanguageVersionOk returns a tuple with the LanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetLanguageVersionOk() (*string, bool) {
	if o == nil || o.LanguageVersion == nil {
		return nil, false
	}
	return o.LanguageVersion, true
}

// HasLanguageVersion returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasLanguageVersion() bool {
	return o != nil && o.LanguageVersion != nil
}

// SetLanguageVersion gets a reference to the given string and assigns it to the LanguageVersion field.
func (o *FleetTracerAttributes) SetLanguageVersion(v string) {
	o.LanguageVersion = &v
}

// GetRemoteConfigStatus returns the RemoteConfigStatus field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetRemoteConfigStatus() string {
	if o == nil || o.RemoteConfigStatus == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigStatus
}

// GetRemoteConfigStatusOk returns a tuple with the RemoteConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetRemoteConfigStatusOk() (*string, bool) {
	if o == nil || o.RemoteConfigStatus == nil {
		return nil, false
	}
	return o.RemoteConfigStatus, true
}

// HasRemoteConfigStatus returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasRemoteConfigStatus() bool {
	return o != nil && o.RemoteConfigStatus != nil
}

// SetRemoteConfigStatus gets a reference to the given string and assigns it to the RemoteConfigStatus field.
func (o *FleetTracerAttributes) SetRemoteConfigStatus(v string) {
	o.RemoteConfigStatus = &v
}

// GetRuntimeIds returns the RuntimeIds field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetRuntimeIds() []string {
	if o == nil || o.RuntimeIds == nil {
		var ret []string
		return ret
	}
	return o.RuntimeIds
}

// GetRuntimeIdsOk returns a tuple with the RuntimeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetRuntimeIdsOk() (*[]string, bool) {
	if o == nil || o.RuntimeIds == nil {
		return nil, false
	}
	return &o.RuntimeIds, true
}

// HasRuntimeIds returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasRuntimeIds() bool {
	return o != nil && o.RuntimeIds != nil
}

// SetRuntimeIds gets a reference to the given []string and assigns it to the RuntimeIds field.
func (o *FleetTracerAttributes) SetRuntimeIds(v []string) {
	o.RuntimeIds = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FleetTracerAttributes) SetService(v string) {
	o.Service = &v
}

// GetServiceHostname returns the ServiceHostname field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetServiceHostname() string {
	if o == nil || o.ServiceHostname == nil {
		var ret string
		return ret
	}
	return *o.ServiceHostname
}

// GetServiceHostnameOk returns a tuple with the ServiceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetServiceHostnameOk() (*string, bool) {
	if o == nil || o.ServiceHostname == nil {
		return nil, false
	}
	return o.ServiceHostname, true
}

// HasServiceHostname returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasServiceHostname() bool {
	return o != nil && o.ServiceHostname != nil
}

// SetServiceHostname gets a reference to the given string and assigns it to the ServiceHostname field.
func (o *FleetTracerAttributes) SetServiceHostname(v string) {
	o.ServiceHostname = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasServiceVersion() bool {
	return o != nil && o.ServiceVersion != nil
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *FleetTracerAttributes) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetTracerVersion returns the TracerVersion field value if set, zero value otherwise.
func (o *FleetTracerAttributes) GetTracerVersion() string {
	if o == nil || o.TracerVersion == nil {
		var ret string
		return ret
	}
	return *o.TracerVersion
}

// GetTracerVersionOk returns a tuple with the TracerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetTracerAttributes) GetTracerVersionOk() (*string, bool) {
	if o == nil || o.TracerVersion == nil {
		return nil, false
	}
	return o.TracerVersion, true
}

// HasTracerVersion returns a boolean if a field has been set.
func (o *FleetTracerAttributes) HasTracerVersion() bool {
	return o != nil && o.TracerVersion != nil
}

// SetTracerVersion gets a reference to the given string and assigns it to the TracerVersion field.
func (o *FleetTracerAttributes) SetTracerVersion(v string) {
	o.TracerVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetTracerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.LanguageVersion != nil {
		toSerialize["language_version"] = o.LanguageVersion
	}
	if o.RemoteConfigStatus != nil {
		toSerialize["remote_config_status"] = o.RemoteConfigStatus
	}
	if o.RuntimeIds != nil {
		toSerialize["runtime_ids"] = o.RuntimeIds
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServiceHostname != nil {
		toSerialize["service_hostname"] = o.ServiceHostname
	}
	if o.ServiceVersion != nil {
		toSerialize["service_version"] = o.ServiceVersion
	}
	if o.TracerVersion != nil {
		toSerialize["tracer_version"] = o.TracerVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetTracerAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env                *string  `json:"env,omitempty"`
		Hostname           *string  `json:"hostname,omitempty"`
		Language           *string  `json:"language,omitempty"`
		LanguageVersion    *string  `json:"language_version,omitempty"`
		RemoteConfigStatus *string  `json:"remote_config_status,omitempty"`
		RuntimeIds         []string `json:"runtime_ids,omitempty"`
		Service            *string  `json:"service,omitempty"`
		ServiceHostname    *string  `json:"service_hostname,omitempty"`
		ServiceVersion     *string  `json:"service_version,omitempty"`
		TracerVersion      *string  `json:"tracer_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "hostname", "language", "language_version", "remote_config_status", "runtime_ids", "service", "service_hostname", "service_version", "tracer_version"})
	} else {
		return err
	}
	o.Env = all.Env
	o.Hostname = all.Hostname
	o.Language = all.Language
	o.LanguageVersion = all.LanguageVersion
	o.RemoteConfigStatus = all.RemoteConfigStatus
	o.RuntimeIds = all.RuntimeIds
	o.Service = all.Service
	o.ServiceHostname = all.ServiceHostname
	o.ServiceVersion = all.ServiceVersion
	o.TracerVersion = all.TracerVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
