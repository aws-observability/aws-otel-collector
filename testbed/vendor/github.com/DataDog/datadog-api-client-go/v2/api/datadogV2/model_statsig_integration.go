// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatsigIntegration The definition of the `StatsigIntegration` object.
type StatsigIntegration struct {
	// The definition of the `StatsigCredentials` object.
	Credentials StatsigCredentials `json:"credentials"`
	// The definition of the `StatsigIntegrationType` object.
	Type StatsigIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatsigIntegration instantiates a new StatsigIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatsigIntegration(credentials StatsigCredentials, typeVar StatsigIntegrationType) *StatsigIntegration {
	this := StatsigIntegration{}
	this.Credentials = credentials
	this.Type = typeVar
	return &this
}

// NewStatsigIntegrationWithDefaults instantiates a new StatsigIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatsigIntegrationWithDefaults() *StatsigIntegration {
	this := StatsigIntegration{}
	return &this
}

// GetCredentials returns the Credentials field value.
func (o *StatsigIntegration) GetCredentials() StatsigCredentials {
	if o == nil {
		var ret StatsigCredentials
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *StatsigIntegration) GetCredentialsOk() (*StatsigCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value.
func (o *StatsigIntegration) SetCredentials(v StatsigCredentials) {
	o.Credentials = v
}

// GetType returns the Type field value.
func (o *StatsigIntegration) GetType() StatsigIntegrationType {
	if o == nil {
		var ret StatsigIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StatsigIntegration) GetTypeOk() (*StatsigIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StatsigIntegration) SetType(v StatsigIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatsigIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["credentials"] = o.Credentials
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatsigIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *StatsigCredentials     `json:"credentials"`
		Type        *StatsigIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Credentials == nil {
		return fmt.Errorf("required field credentials missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"credentials", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Credentials = *all.Credentials
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
