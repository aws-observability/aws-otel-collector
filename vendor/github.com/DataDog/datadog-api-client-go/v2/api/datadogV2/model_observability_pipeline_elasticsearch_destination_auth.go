// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationAuth Authentication settings for the Elasticsearch destination.
// When `strategy` is `basic`, use `username_key` and `password_key` to reference credentials stored in environment variables or secrets.
type ObservabilityPipelineElasticsearchDestinationAuth struct {
	// Name of the environment variable or secret that holds the Elasticsearch password (used when `strategy` is `basic`).
	PasswordKey *string `json:"password_key,omitempty"`
	// The authentication strategy to use.
	Strategy ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy `json:"strategy"`
	// Name of the environment variable or secret that holds the Elasticsearch username (used when `strategy` is `basic`).
	UsernameKey *string `json:"username_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineElasticsearchDestinationAuth instantiates a new ObservabilityPipelineElasticsearchDestinationAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineElasticsearchDestinationAuth(strategy ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) *ObservabilityPipelineElasticsearchDestinationAuth {
	this := ObservabilityPipelineElasticsearchDestinationAuth{}
	this.Strategy = strategy
	return &this
}

// NewObservabilityPipelineElasticsearchDestinationAuthWithDefaults instantiates a new ObservabilityPipelineElasticsearchDestinationAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineElasticsearchDestinationAuthWithDefaults() *ObservabilityPipelineElasticsearchDestinationAuth {
	this := ObservabilityPipelineElasticsearchDestinationAuth{}
	return &this
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetStrategy returns the Strategy field value.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) GetStrategy() ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy {
	if o == nil {
		var ret ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) GetStrategyOk() (*ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) SetStrategy(v ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) {
	o.Strategy = v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineElasticsearchDestinationAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PasswordKey != nil {
		toSerialize["password_key"] = o.PasswordKey
	}
	toSerialize["strategy"] = o.Strategy
	if o.UsernameKey != nil {
		toSerialize["username_key"] = o.UsernameKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineElasticsearchDestinationAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PasswordKey *string                                                       `json:"password_key,omitempty"`
		Strategy    *ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy `json:"strategy"`
		UsernameKey *string                                                       `json:"username_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"password_key", "strategy", "username_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PasswordKey = all.PasswordKey
	if !all.Strategy.IsValid() {
		hasInvalidField = true
	} else {
		o.Strategy = *all.Strategy
	}
	o.UsernameKey = all.UsernameKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
