// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreshserviceIntegration The definition of the `FreshserviceIntegration` object.
type FreshserviceIntegration struct {
	// The definition of the `FreshserviceCredentials` object.
	Credentials FreshserviceCredentials `json:"credentials"`
	// The definition of the `FreshserviceIntegrationType` object.
	Type FreshserviceIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFreshserviceIntegration instantiates a new FreshserviceIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFreshserviceIntegration(credentials FreshserviceCredentials, typeVar FreshserviceIntegrationType) *FreshserviceIntegration {
	this := FreshserviceIntegration{}
	this.Credentials = credentials
	this.Type = typeVar
	return &this
}

// NewFreshserviceIntegrationWithDefaults instantiates a new FreshserviceIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFreshserviceIntegrationWithDefaults() *FreshserviceIntegration {
	this := FreshserviceIntegration{}
	return &this
}

// GetCredentials returns the Credentials field value.
func (o *FreshserviceIntegration) GetCredentials() FreshserviceCredentials {
	if o == nil {
		var ret FreshserviceCredentials
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *FreshserviceIntegration) GetCredentialsOk() (*FreshserviceCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value.
func (o *FreshserviceIntegration) SetCredentials(v FreshserviceCredentials) {
	o.Credentials = v
}

// GetType returns the Type field value.
func (o *FreshserviceIntegration) GetType() FreshserviceIntegrationType {
	if o == nil {
		var ret FreshserviceIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreshserviceIntegration) GetTypeOk() (*FreshserviceIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FreshserviceIntegration) SetType(v FreshserviceIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FreshserviceIntegration) MarshalJSON() ([]byte, error) {
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
func (o *FreshserviceIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *FreshserviceCredentials     `json:"credentials"`
		Type        *FreshserviceIntegrationType `json:"type"`
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
