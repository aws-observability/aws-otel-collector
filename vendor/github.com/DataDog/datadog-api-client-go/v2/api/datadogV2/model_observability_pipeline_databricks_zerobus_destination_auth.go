// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatabricksZerobusDestinationAuth OAuth credentials for authenticating with the Databricks Zerobus ingestion API.
type ObservabilityPipelineDatabricksZerobusDestinationAuth struct {
	// Your service principal application ID (UUID).
	ClientId string `json:"client_id"`
	// Name of the environment variable or secret that holds the OAuth client secret used to authenticate with the Databricks ingestion endpoint.
	ClientSecretKey *string `json:"client_secret_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDatabricksZerobusDestinationAuth instantiates a new ObservabilityPipelineDatabricksZerobusDestinationAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDatabricksZerobusDestinationAuth(clientId string) *ObservabilityPipelineDatabricksZerobusDestinationAuth {
	this := ObservabilityPipelineDatabricksZerobusDestinationAuth{}
	this.ClientId = clientId
	return &this
}

// NewObservabilityPipelineDatabricksZerobusDestinationAuthWithDefaults instantiates a new ObservabilityPipelineDatabricksZerobusDestinationAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDatabricksZerobusDestinationAuthWithDefaults() *ObservabilityPipelineDatabricksZerobusDestinationAuth {
	this := ObservabilityPipelineDatabricksZerobusDestinationAuth{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecretKey returns the ClientSecretKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) GetClientSecretKey() string {
	if o == nil || o.ClientSecretKey == nil {
		var ret string
		return ret
	}
	return *o.ClientSecretKey
}

// GetClientSecretKeyOk returns a tuple with the ClientSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) GetClientSecretKeyOk() (*string, bool) {
	if o == nil || o.ClientSecretKey == nil {
		return nil, false
	}
	return o.ClientSecretKey, true
}

// HasClientSecretKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) HasClientSecretKey() bool {
	return o != nil && o.ClientSecretKey != nil
}

// SetClientSecretKey gets a reference to the given string and assigns it to the ClientSecretKey field.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) SetClientSecretKey(v string) {
	o.ClientSecretKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDatabricksZerobusDestinationAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	if o.ClientSecretKey != nil {
		toSerialize["client_secret_key"] = o.ClientSecretKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDatabricksZerobusDestinationAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId        *string `json:"client_id"`
		ClientSecretKey *string `json:"client_secret_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "client_secret_key"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.ClientSecretKey = all.ClientSecretKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
