// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebIntegrationAccountCreateRequestAttributes Attributes object for creating a web integration account.
type WebIntegrationAccountCreateRequestAttributes struct {
	// A human-readable name for the account. Must be unique among accounts of the same integration.
	Name string `json:"name"`
	// Integration-specific secrets. The shape of this object varies by integration. Secrets
	// are write-only and never returned by the API.
	Secrets map[string]interface{} `json:"secrets"`
	// Integration-specific settings. The shape of this object varies by integration.
	Settings map[string]interface{} `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebIntegrationAccountCreateRequestAttributes instantiates a new WebIntegrationAccountCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebIntegrationAccountCreateRequestAttributes(name string, secrets map[string]interface{}, settings map[string]interface{}) *WebIntegrationAccountCreateRequestAttributes {
	this := WebIntegrationAccountCreateRequestAttributes{}
	this.Name = name
	this.Secrets = secrets
	this.Settings = settings
	return &this
}

// NewWebIntegrationAccountCreateRequestAttributesWithDefaults instantiates a new WebIntegrationAccountCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebIntegrationAccountCreateRequestAttributesWithDefaults() *WebIntegrationAccountCreateRequestAttributes {
	this := WebIntegrationAccountCreateRequestAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *WebIntegrationAccountCreateRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *WebIntegrationAccountCreateRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetSecrets returns the Secrets field value.
func (o *WebIntegrationAccountCreateRequestAttributes) GetSecrets() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountCreateRequestAttributes) GetSecretsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value.
func (o *WebIntegrationAccountCreateRequestAttributes) SetSecrets(v map[string]interface{}) {
	o.Secrets = v
}

// GetSettings returns the Settings field value.
func (o *WebIntegrationAccountCreateRequestAttributes) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountCreateRequestAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *WebIntegrationAccountCreateRequestAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebIntegrationAccountCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["secrets"] = o.Secrets
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebIntegrationAccountCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string                 `json:"name"`
		Secrets  *map[string]interface{} `json:"secrets"`
		Settings *map[string]interface{} `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Secrets == nil {
		return fmt.Errorf("required field secrets missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "secrets", "settings"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Secrets = *all.Secrets
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
