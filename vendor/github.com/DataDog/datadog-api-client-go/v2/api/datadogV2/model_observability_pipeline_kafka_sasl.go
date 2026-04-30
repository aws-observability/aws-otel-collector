// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaSasl Specifies the SASL mechanism for authenticating with a Kafka cluster.
type ObservabilityPipelineKafkaSasl struct {
	// SASL mechanism used for Kafka authentication.
	Mechanism *ObservabilityPipelineKafkaSaslMechanism `json:"mechanism,omitempty"`
	// Name of the environment variable or secret that holds the SASL password.
	PasswordKey *string `json:"password_key,omitempty"`
	// Name of the environment variable or secret that holds the SASL username.
	UsernameKey *string `json:"username_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineKafkaSasl instantiates a new ObservabilityPipelineKafkaSasl object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineKafkaSasl() *ObservabilityPipelineKafkaSasl {
	this := ObservabilityPipelineKafkaSasl{}
	return &this
}

// NewObservabilityPipelineKafkaSaslWithDefaults instantiates a new ObservabilityPipelineKafkaSasl object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineKafkaSaslWithDefaults() *ObservabilityPipelineKafkaSasl {
	this := ObservabilityPipelineKafkaSasl{}
	return &this
}

// GetMechanism returns the Mechanism field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSasl) GetMechanism() ObservabilityPipelineKafkaSaslMechanism {
	if o == nil || o.Mechanism == nil {
		var ret ObservabilityPipelineKafkaSaslMechanism
		return ret
	}
	return *o.Mechanism
}

// GetMechanismOk returns a tuple with the Mechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSasl) GetMechanismOk() (*ObservabilityPipelineKafkaSaslMechanism, bool) {
	if o == nil || o.Mechanism == nil {
		return nil, false
	}
	return o.Mechanism, true
}

// HasMechanism returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSasl) HasMechanism() bool {
	return o != nil && o.Mechanism != nil
}

// SetMechanism gets a reference to the given ObservabilityPipelineKafkaSaslMechanism and assigns it to the Mechanism field.
func (o *ObservabilityPipelineKafkaSasl) SetMechanism(v ObservabilityPipelineKafkaSaslMechanism) {
	o.Mechanism = &v
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSasl) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSasl) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSasl) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelineKafkaSasl) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSasl) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSasl) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSasl) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelineKafkaSasl) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineKafkaSasl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Mechanism != nil {
		toSerialize["mechanism"] = o.Mechanism
	}
	if o.PasswordKey != nil {
		toSerialize["password_key"] = o.PasswordKey
	}
	if o.UsernameKey != nil {
		toSerialize["username_key"] = o.UsernameKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineKafkaSasl) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mechanism   *ObservabilityPipelineKafkaSaslMechanism `json:"mechanism,omitempty"`
		PasswordKey *string                                  `json:"password_key,omitempty"`
		UsernameKey *string                                  `json:"username_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mechanism", "password_key", "username_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Mechanism != nil && !all.Mechanism.IsValid() {
		hasInvalidField = true
	} else {
		o.Mechanism = all.Mechanism
	}
	o.PasswordKey = all.PasswordKey
	o.UsernameKey = all.UsernameKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
