// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationCredentialsValidateAttributes The credentials to validate against the external entity source.
type SecurityMonitoringIntegrationCredentialsValidateAttributes struct {
	// The domain associated with the external entity source.
	Domain string `json:"domain"`
	// The type of external source that provides entities to Cloud SIEM.
	IntegrationType SecurityMonitoringIntegrationType `json:"integration_type"`
	// The secrets used to authenticate against the external entity source. The accepted keys depend on the source type (for example, `admin_email` for Google Workspace).
	Secrets map[string]interface{} `json:"secrets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationCredentialsValidateAttributes instantiates a new SecurityMonitoringIntegrationCredentialsValidateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationCredentialsValidateAttributes(domain string, integrationType SecurityMonitoringIntegrationType, secrets map[string]interface{}) *SecurityMonitoringIntegrationCredentialsValidateAttributes {
	this := SecurityMonitoringIntegrationCredentialsValidateAttributes{}
	this.Domain = domain
	this.IntegrationType = integrationType
	this.Secrets = secrets
	return &this
}

// NewSecurityMonitoringIntegrationCredentialsValidateAttributesWithDefaults instantiates a new SecurityMonitoringIntegrationCredentialsValidateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationCredentialsValidateAttributesWithDefaults() *SecurityMonitoringIntegrationCredentialsValidateAttributes {
	this := SecurityMonitoringIntegrationCredentialsValidateAttributes{}
	return &this
}

// GetDomain returns the Domain field value.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetIntegrationType() SecurityMonitoringIntegrationType {
	if o == nil {
		var ret SecurityMonitoringIntegrationType
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) SetIntegrationType(v SecurityMonitoringIntegrationType) {
	o.IntegrationType = v
}

// GetSecrets returns the Secrets field value.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetSecrets() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) GetSecretsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) SetSecrets(v map[string]interface{}) {
	o.Secrets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationCredentialsValidateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["domain"] = o.Domain
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["secrets"] = o.Secrets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationCredentialsValidateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domain          *string                            `json:"domain"`
		IntegrationType *SecurityMonitoringIntegrationType `json:"integration_type"`
		Secrets         *map[string]interface{}            `json:"secrets"`
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
	if all.Secrets == nil {
		return fmt.Errorf("required field secrets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain", "integration_type", "secrets"})
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
	o.Secrets = *all.Secrets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
