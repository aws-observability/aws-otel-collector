// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigUpdateAttributes Fields to update on the entity context sync configuration. All fields are optional.
type SecurityMonitoringIntegrationConfigUpdateAttributes struct {
	// The new domain associated with the external entity source.
	Domain *string `json:"domain,omitempty"`
	// Whether the entity context sync should be enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The type of external source that provides entities to Cloud SIEM.
	IntegrationType *SecurityMonitoringIntegrationType `json:"integration_type,omitempty"`
	// The new display name for the entity context sync configuration.
	Name *string `json:"name,omitempty"`
	// The secrets used to authenticate against the external entity source. The accepted keys depend on the source type (for example, `admin_email` for Google Workspace).
	Secrets map[string]interface{} `json:"secrets,omitempty"`
	// Free-form, non-sensitive settings for the entity context sync. The accepted keys depend on the source type.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigUpdateAttributes instantiates a new SecurityMonitoringIntegrationConfigUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigUpdateAttributes() *SecurityMonitoringIntegrationConfigUpdateAttributes {
	this := SecurityMonitoringIntegrationConfigUpdateAttributes{}
	return &this
}

// NewSecurityMonitoringIntegrationConfigUpdateAttributesWithDefaults instantiates a new SecurityMonitoringIntegrationConfigUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigUpdateAttributesWithDefaults() *SecurityMonitoringIntegrationConfigUpdateAttributes {
	this := SecurityMonitoringIntegrationConfigUpdateAttributes{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) SetDomain(v string) {
	o.Domain = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetIntegrationType() SecurityMonitoringIntegrationType {
	if o == nil || o.IntegrationType == nil {
		var ret SecurityMonitoringIntegrationType
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationType, bool) {
	if o == nil || o.IntegrationType == nil {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) HasIntegrationType() bool {
	return o != nil && o.IntegrationType != nil
}

// SetIntegrationType gets a reference to the given SecurityMonitoringIntegrationType and assigns it to the IntegrationType field.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) SetIntegrationType(v SecurityMonitoringIntegrationType) {
	o.IntegrationType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetSecrets() map[string]interface{} {
	if o == nil || o.Secrets == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetSecretsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) HasSecrets() bool {
	return o != nil && o.Secrets != nil
}

// SetSecrets gets a reference to the given map[string]interface{} and assigns it to the Secrets field.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) SetSecrets(v map[string]interface{}) {
	o.Secrets = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.IntegrationType != nil {
		toSerialize["integration_type"] = o.IntegrationType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain          *string                            `json:"domain,omitempty"`
		Enabled         *bool                              `json:"enabled,omitempty"`
		IntegrationType *SecurityMonitoringIntegrationType `json:"integration_type,omitempty"`
		Name            *string                            `json:"name,omitempty"`
		Secrets         map[string]interface{}             `json:"secrets,omitempty"`
		Settings        map[string]interface{}             `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "enabled", "integration_type", "name", "secrets", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Domain = all.Domain
	o.Enabled = all.Enabled
	if all.IntegrationType != nil && !all.IntegrationType.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationType = all.IntegrationType
	}
	o.Name = all.Name
	o.Secrets = all.Secrets
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
