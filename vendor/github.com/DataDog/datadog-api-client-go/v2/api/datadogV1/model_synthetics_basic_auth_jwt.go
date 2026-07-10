// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthJWT Object to handle JWT authentication when performing the test.
type SyntheticsBasicAuthJWT struct {
	// Standard JWT claims to automatically inject.
	AddClaims *SyntheticsBasicAuthJWTAddClaims `json:"addClaims,omitempty"`
	// Algorithm to use for the JWT authentication.
	Algorithm SyntheticsBasicAuthJWTAlgorithm `json:"algorithm"`
	// Token time-to-live in seconds.
	ExpiresIn *int64 `json:"expiresIn,omitempty"`
	// Custom JWT header as a JSON string.
	Header *string `json:"header,omitempty"`
	// JWT claims as a JSON string.
	Payload string `json:"payload"`
	// Signing key for the JWT authentication. Use the shared secret for `HS256`
	// or the private key (PEM format) for `RS256` and `ES256`.
	Secret string `json:"secret"`
	// Prefix added before the token in the `Authorization` header. Defaults to `Bearer`.
	TokenPrefix *string `json:"tokenPrefix,omitempty"`
	// The type of authentication to use when performing the test.
	Type SyntheticsBasicAuthJWTType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsBasicAuthJWT instantiates a new SyntheticsBasicAuthJWT object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsBasicAuthJWT(algorithm SyntheticsBasicAuthJWTAlgorithm, payload string, secret string, typeVar SyntheticsBasicAuthJWTType) *SyntheticsBasicAuthJWT {
	this := SyntheticsBasicAuthJWT{}
	this.Algorithm = algorithm
	this.Payload = payload
	this.Secret = secret
	this.Type = typeVar
	return &this
}

// NewSyntheticsBasicAuthJWTWithDefaults instantiates a new SyntheticsBasicAuthJWT object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsBasicAuthJWTWithDefaults() *SyntheticsBasicAuthJWT {
	this := SyntheticsBasicAuthJWT{}
	var typeVar SyntheticsBasicAuthJWTType = SYNTHETICSBASICAUTHJWTTYPE_JWT
	this.Type = typeVar
	return &this
}

// GetAddClaims returns the AddClaims field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthJWT) GetAddClaims() SyntheticsBasicAuthJWTAddClaims {
	if o == nil || o.AddClaims == nil {
		var ret SyntheticsBasicAuthJWTAddClaims
		return ret
	}
	return *o.AddClaims
}

// GetAddClaimsOk returns a tuple with the AddClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetAddClaimsOk() (*SyntheticsBasicAuthJWTAddClaims, bool) {
	if o == nil || o.AddClaims == nil {
		return nil, false
	}
	return o.AddClaims, true
}

// HasAddClaims returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthJWT) HasAddClaims() bool {
	return o != nil && o.AddClaims != nil
}

// SetAddClaims gets a reference to the given SyntheticsBasicAuthJWTAddClaims and assigns it to the AddClaims field.
func (o *SyntheticsBasicAuthJWT) SetAddClaims(v SyntheticsBasicAuthJWTAddClaims) {
	o.AddClaims = &v
}

// GetAlgorithm returns the Algorithm field value.
func (o *SyntheticsBasicAuthJWT) GetAlgorithm() SyntheticsBasicAuthJWTAlgorithm {
	if o == nil {
		var ret SyntheticsBasicAuthJWTAlgorithm
		return ret
	}
	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetAlgorithmOk() (*SyntheticsBasicAuthJWTAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value.
func (o *SyntheticsBasicAuthJWT) SetAlgorithm(v SyntheticsBasicAuthJWTAlgorithm) {
	o.Algorithm = v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthJWT) GetExpiresIn() int64 {
	if o == nil || o.ExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetExpiresInOk() (*int64, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthJWT) HasExpiresIn() bool {
	return o != nil && o.ExpiresIn != nil
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *SyntheticsBasicAuthJWT) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthJWT) GetHeader() string {
	if o == nil || o.Header == nil {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetHeaderOk() (*string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthJWT) HasHeader() bool {
	return o != nil && o.Header != nil
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *SyntheticsBasicAuthJWT) SetHeader(v string) {
	o.Header = &v
}

// GetPayload returns the Payload field value.
func (o *SyntheticsBasicAuthJWT) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value.
func (o *SyntheticsBasicAuthJWT) SetPayload(v string) {
	o.Payload = v
}

// GetSecret returns the Secret field value.
func (o *SyntheticsBasicAuthJWT) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value.
func (o *SyntheticsBasicAuthJWT) SetSecret(v string) {
	o.Secret = v
}

// GetTokenPrefix returns the TokenPrefix field value if set, zero value otherwise.
func (o *SyntheticsBasicAuthJWT) GetTokenPrefix() string {
	if o == nil || o.TokenPrefix == nil {
		var ret string
		return ret
	}
	return *o.TokenPrefix
}

// GetTokenPrefixOk returns a tuple with the TokenPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetTokenPrefixOk() (*string, bool) {
	if o == nil || o.TokenPrefix == nil {
		return nil, false
	}
	return o.TokenPrefix, true
}

// HasTokenPrefix returns a boolean if a field has been set.
func (o *SyntheticsBasicAuthJWT) HasTokenPrefix() bool {
	return o != nil && o.TokenPrefix != nil
}

// SetTokenPrefix gets a reference to the given string and assigns it to the TokenPrefix field.
func (o *SyntheticsBasicAuthJWT) SetTokenPrefix(v string) {
	o.TokenPrefix = &v
}

// GetType returns the Type field value.
func (o *SyntheticsBasicAuthJWT) GetType() SyntheticsBasicAuthJWTType {
	if o == nil {
		var ret SyntheticsBasicAuthJWTType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuthJWT) GetTypeOk() (*SyntheticsBasicAuthJWTType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsBasicAuthJWT) SetType(v SyntheticsBasicAuthJWTType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsBasicAuthJWT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddClaims != nil {
		toSerialize["addClaims"] = o.AddClaims
	}
	toSerialize["algorithm"] = o.Algorithm
	if o.ExpiresIn != nil {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	toSerialize["payload"] = o.Payload
	toSerialize["secret"] = o.Secret
	if o.TokenPrefix != nil {
		toSerialize["tokenPrefix"] = o.TokenPrefix
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsBasicAuthJWT) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddClaims   *SyntheticsBasicAuthJWTAddClaims `json:"addClaims,omitempty"`
		Algorithm   *SyntheticsBasicAuthJWTAlgorithm `json:"algorithm"`
		ExpiresIn   *int64                           `json:"expiresIn,omitempty"`
		Header      *string                          `json:"header,omitempty"`
		Payload     *string                          `json:"payload"`
		Secret      *string                          `json:"secret"`
		TokenPrefix *string                          `json:"tokenPrefix,omitempty"`
		Type        *SyntheticsBasicAuthJWTType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Algorithm == nil {
		return fmt.Errorf("required field algorithm missing")
	}
	if all.Payload == nil {
		return fmt.Errorf("required field payload missing")
	}
	if all.Secret == nil {
		return fmt.Errorf("required field secret missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"addClaims", "algorithm", "expiresIn", "header", "payload", "secret", "tokenPrefix", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AddClaims != nil && all.AddClaims.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AddClaims = all.AddClaims
	if !all.Algorithm.IsValid() {
		hasInvalidField = true
	} else {
		o.Algorithm = *all.Algorithm
	}
	o.ExpiresIn = all.ExpiresIn
	o.Header = all.Header
	o.Payload = *all.Payload
	o.Secret = *all.Secret
	o.TokenPrefix = all.TokenPrefix
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
