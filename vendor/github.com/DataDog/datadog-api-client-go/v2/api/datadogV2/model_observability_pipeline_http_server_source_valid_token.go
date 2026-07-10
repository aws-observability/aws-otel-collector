// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSourceValidToken An accepted token used to authenticate incoming HTTP server requests.
type ObservabilityPipelineHttpServerSourceValidToken struct {
	// Indicates whether this token is currently accepted. Disabled tokens are rejected without
	// being removed from the configuration.
	Enabled *bool `json:"enabled,omitempty"`
	// An optional metadata field that is attached to every event authenticated by the
	// associated token. Both `key` and `value` must match `^[A-Za-z0-9_]+$`.
	FieldToAdd *ObservabilityPipelineSourceValidTokenFieldToAdd `json:"field_to_add,omitempty"`
	// Specifies where the worker extracts the token from in the incoming HTTP request.
	// This can be either a built-in location (`path` or `address`) or an HTTP header object.
	PathToToken *ObservabilityPipelineHttpServerSourceValidTokenPathToToken `json:"path_to_token,omitempty"`
	// Name of the environment variable or secret that holds the expected token value.
	TokenKey string `json:"token_key"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineHttpServerSourceValidToken instantiates a new ObservabilityPipelineHttpServerSourceValidToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineHttpServerSourceValidToken(tokenKey string) *ObservabilityPipelineHttpServerSourceValidToken {
	this := ObservabilityPipelineHttpServerSourceValidToken{}
	var enabled bool = true
	this.Enabled = &enabled
	this.TokenKey = tokenKey
	return &this
}

// NewObservabilityPipelineHttpServerSourceValidTokenWithDefaults instantiates a new ObservabilityPipelineHttpServerSourceValidToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineHttpServerSourceValidTokenWithDefaults() *ObservabilityPipelineHttpServerSourceValidToken {
	this := ObservabilityPipelineHttpServerSourceValidToken{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ObservabilityPipelineHttpServerSourceValidToken) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFieldToAdd returns the FieldToAdd field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetFieldToAdd() ObservabilityPipelineSourceValidTokenFieldToAdd {
	if o == nil || o.FieldToAdd == nil {
		var ret ObservabilityPipelineSourceValidTokenFieldToAdd
		return ret
	}
	return *o.FieldToAdd
}

// GetFieldToAddOk returns a tuple with the FieldToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetFieldToAddOk() (*ObservabilityPipelineSourceValidTokenFieldToAdd, bool) {
	if o == nil || o.FieldToAdd == nil {
		return nil, false
	}
	return o.FieldToAdd, true
}

// HasFieldToAdd returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) HasFieldToAdd() bool {
	return o != nil && o.FieldToAdd != nil
}

// SetFieldToAdd gets a reference to the given ObservabilityPipelineSourceValidTokenFieldToAdd and assigns it to the FieldToAdd field.
func (o *ObservabilityPipelineHttpServerSourceValidToken) SetFieldToAdd(v ObservabilityPipelineSourceValidTokenFieldToAdd) {
	o.FieldToAdd = &v
}

// GetPathToToken returns the PathToToken field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetPathToToken() ObservabilityPipelineHttpServerSourceValidTokenPathToToken {
	if o == nil || o.PathToToken == nil {
		var ret ObservabilityPipelineHttpServerSourceValidTokenPathToToken
		return ret
	}
	return *o.PathToToken
}

// GetPathToTokenOk returns a tuple with the PathToToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetPathToTokenOk() (*ObservabilityPipelineHttpServerSourceValidTokenPathToToken, bool) {
	if o == nil || o.PathToToken == nil {
		return nil, false
	}
	return o.PathToToken, true
}

// HasPathToToken returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) HasPathToToken() bool {
	return o != nil && o.PathToToken != nil
}

// SetPathToToken gets a reference to the given ObservabilityPipelineHttpServerSourceValidTokenPathToToken and assigns it to the PathToToken field.
func (o *ObservabilityPipelineHttpServerSourceValidToken) SetPathToToken(v ObservabilityPipelineHttpServerSourceValidTokenPathToToken) {
	o.PathToToken = &v
}

// GetTokenKey returns the TokenKey field value.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetTokenKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSourceValidToken) GetTokenKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenKey, true
}

// SetTokenKey sets field value.
func (o *ObservabilityPipelineHttpServerSourceValidToken) SetTokenKey(v string) {
	o.TokenKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineHttpServerSourceValidToken) MarshalJSON() ([]byte, error) {
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
	if o.PathToToken != nil {
		toSerialize["path_to_token"] = o.PathToToken
	}
	toSerialize["token_key"] = o.TokenKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineHttpServerSourceValidToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled     *bool                                                       `json:"enabled,omitempty"`
		FieldToAdd  *ObservabilityPipelineSourceValidTokenFieldToAdd            `json:"field_to_add,omitempty"`
		PathToToken *ObservabilityPipelineHttpServerSourceValidTokenPathToToken `json:"path_to_token,omitempty"`
		TokenKey    *string                                                     `json:"token_key"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TokenKey == nil {
		return fmt.Errorf("required field token_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "field_to_add", "path_to_token", "token_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.FieldToAdd != nil && all.FieldToAdd.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FieldToAdd = all.FieldToAdd
	o.PathToToken = all.PathToToken
	o.TokenKey = *all.TokenKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
