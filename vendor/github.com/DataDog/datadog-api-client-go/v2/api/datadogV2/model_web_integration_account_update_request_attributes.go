// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebIntegrationAccountUpdateRequestAttributes Attributes object for updating a web integration account.
type WebIntegrationAccountUpdateRequestAttributes struct {
	// A human-readable name for the account.
	Name *string `json:"name,omitempty"`
	// Integration-specific secrets. The shape of this object varies by integration. Secrets
	// are write-only and never returned by the API.
	Secrets map[string]interface{} `json:"secrets,omitempty"`
	// Integration-specific settings. The shape of this object varies by integration.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebIntegrationAccountUpdateRequestAttributes instantiates a new WebIntegrationAccountUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebIntegrationAccountUpdateRequestAttributes() *WebIntegrationAccountUpdateRequestAttributes {
	this := WebIntegrationAccountUpdateRequestAttributes{}
	return &this
}

// NewWebIntegrationAccountUpdateRequestAttributesWithDefaults instantiates a new WebIntegrationAccountUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebIntegrationAccountUpdateRequestAttributesWithDefaults() *WebIntegrationAccountUpdateRequestAttributes {
	this := WebIntegrationAccountUpdateRequestAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebIntegrationAccountUpdateRequestAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountUpdateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebIntegrationAccountUpdateRequestAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebIntegrationAccountUpdateRequestAttributes) SetName(v string) {
	o.Name = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *WebIntegrationAccountUpdateRequestAttributes) GetSecrets() map[string]interface{} {
	if o == nil || o.Secrets == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountUpdateRequestAttributes) GetSecretsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *WebIntegrationAccountUpdateRequestAttributes) HasSecrets() bool {
	return o != nil && o.Secrets != nil
}

// SetSecrets gets a reference to the given map[string]interface{} and assigns it to the Secrets field.
func (o *WebIntegrationAccountUpdateRequestAttributes) SetSecrets(v map[string]interface{}) {
	o.Secrets = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *WebIntegrationAccountUpdateRequestAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebIntegrationAccountUpdateRequestAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *WebIntegrationAccountUpdateRequestAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *WebIntegrationAccountUpdateRequestAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebIntegrationAccountUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *WebIntegrationAccountUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string                `json:"name,omitempty"`
		Secrets  map[string]interface{} `json:"secrets,omitempty"`
		Settings map[string]interface{} `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "secrets", "settings"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Secrets = all.Secrets
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
