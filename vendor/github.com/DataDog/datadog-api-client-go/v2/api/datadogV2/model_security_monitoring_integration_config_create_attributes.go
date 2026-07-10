// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigCreateAttributes The attributes of the entity context sync configuration to create.
type SecurityMonitoringIntegrationConfigCreateAttributes struct {
	// The domain associated with the external entity source.
	Domain string `json:"domain"`
	// The type of external source that provides entities to Cloud SIEM.
	IntegrationType SecurityMonitoringIntegrationType `json:"integration_type"`
	// The display name for the entity context sync configuration.
	Name string `json:"name"`
	// The secrets used to authenticate against the external entity source. The accepted keys depend on the source type (for example, `admin_email` for Google Workspace).
	Secrets map[string]interface{} `json:"secrets"`
	// Free-form, non-sensitive settings for the entity context sync. The accepted keys depend on the source type.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigCreateAttributes instantiates a new SecurityMonitoringIntegrationConfigCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigCreateAttributes(domain string, integrationType SecurityMonitoringIntegrationType, name string, secrets map[string]interface{}) *SecurityMonitoringIntegrationConfigCreateAttributes {
	this := SecurityMonitoringIntegrationConfigCreateAttributes{}
	this.Domain = domain
	this.IntegrationType = integrationType
	this.Name = name
	this.Secrets = secrets
	return &this
}

// NewSecurityMonitoringIntegrationConfigCreateAttributesWithDefaults instantiates a new SecurityMonitoringIntegrationConfigCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigCreateAttributesWithDefaults() *SecurityMonitoringIntegrationConfigCreateAttributes {
	this := SecurityMonitoringIntegrationConfigCreateAttributes{}
	return &this
}

// GetDomain returns the Domain field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetIntegrationType() SecurityMonitoringIntegrationType {
	if o == nil {
		var ret SecurityMonitoringIntegrationType
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) SetIntegrationType(v SecurityMonitoringIntegrationType) {
	o.IntegrationType = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSecrets returns the Secrets field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetSecrets() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetSecretsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) SetSecrets(v map[string]interface{}) {
	o.Secrets = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["domain"] = o.Domain
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["name"] = o.Name
	toSerialize["secrets"] = o.Secrets
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain          *string                            `json:"domain"`
		IntegrationType *SecurityMonitoringIntegrationType `json:"integration_type"`
		Name            *string                            `json:"name"`
		Secrets         *map[string]interface{}            `json:"secrets"`
		Settings        map[string]interface{}             `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Secrets == nil {
		return fmt.Errorf("required field secrets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "integration_type", "name", "secrets", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Domain = *all.Domain
	if !all.IntegrationType.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationType = *all.IntegrationType
	}
	o.Name = *all.Name
	o.Secrets = *all.Secrets
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
