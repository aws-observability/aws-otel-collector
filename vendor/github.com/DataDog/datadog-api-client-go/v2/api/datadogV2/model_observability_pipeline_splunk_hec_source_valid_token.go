// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecSourceValidToken An accepted HEC token used to authenticate incoming Splunk HEC requests.
type ObservabilityPipelineSplunkHecSourceValidToken struct {
	// Indicates whether this token is currently accepted. Disabled tokens are rejected without
	// being removed from the configuration.
	Enabled *bool `json:"enabled,omitempty"`
	// An optional metadata field that is attached to every event authenticated by the
	// associated token. Both `key` and `value` must match `^[A-Za-z0-9_]+$`.
	FieldToAdd *ObservabilityPipelineSourceValidTokenFieldToAdd `json:"field_to_add,omitempty"`
	// Name of the environment variable or secret that holds the expected HEC token value.
	TokenKey string `json:"token_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSplunkHecSourceValidToken instantiates a new ObservabilityPipelineSplunkHecSourceValidToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSplunkHecSourceValidToken(tokenKey string) *ObservabilityPipelineSplunkHecSourceValidToken {
	this := ObservabilityPipelineSplunkHecSourceValidToken{}
	var enabled bool = true
	this.Enabled = &enabled
	this.TokenKey = tokenKey
	return &this
}

// NewObservabilityPipelineSplunkHecSourceValidTokenWithDefaults instantiates a new ObservabilityPipelineSplunkHecSourceValidToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSplunkHecSourceValidTokenWithDefaults() *ObservabilityPipelineSplunkHecSourceValidToken {
	this := ObservabilityPipelineSplunkHecSourceValidToken{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFieldToAdd returns the FieldToAdd field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) GetFieldToAdd() ObservabilityPipelineSourceValidTokenFieldToAdd {
	if o == nil || o.FieldToAdd == nil {
		var ret ObservabilityPipelineSourceValidTokenFieldToAdd
		return ret
	}
	return *o.FieldToAdd
}

// GetFieldToAddOk returns a tuple with the FieldToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) GetFieldToAddOk() (*ObservabilityPipelineSourceValidTokenFieldToAdd, bool) {
	if o == nil || o.FieldToAdd == nil {
		return nil, false
	}
	return o.FieldToAdd, true
}

// HasFieldToAdd returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) HasFieldToAdd() bool {
	return o != nil && o.FieldToAdd != nil
}

// SetFieldToAdd gets a reference to the given ObservabilityPipelineSourceValidTokenFieldToAdd and assigns it to the FieldToAdd field.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) SetFieldToAdd(v ObservabilityPipelineSourceValidTokenFieldToAdd) {
	o.FieldToAdd = &v
}

// GetTokenKey returns the TokenKey field value.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) GetTokenKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) GetTokenKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenKey, true
}

// SetTokenKey sets field value.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) SetTokenKey(v string) {
	o.TokenKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSplunkHecSourceValidToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FieldToAdd != nil {
		toSerialize["field_to_add"] = o.FieldToAdd
	}
	toSerialize["token_key"] = o.TokenKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSplunkHecSourceValidToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled    *bool                                            `json:"enabled,omitempty"`
		FieldToAdd *ObservabilityPipelineSourceValidTokenFieldToAdd `json:"field_to_add,omitempty"`
		TokenKey   *string                                          `json:"token_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TokenKey == nil {
		return fmt.Errorf("required field token_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "field_to_add", "token_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.FieldToAdd != nil && all.FieldToAdd.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FieldToAdd = all.FieldToAdd
	o.TokenKey = *all.TokenKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
