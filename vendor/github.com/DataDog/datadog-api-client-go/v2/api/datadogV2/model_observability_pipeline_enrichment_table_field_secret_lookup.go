// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFieldSecretLookup Looks up a value stored as a pipeline secret.
type ObservabilityPipelineEnrichmentTableFieldSecretLookup struct {
	// The name of the secret containing the lookup key value.
	Secret string `json:"secret"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableFieldSecretLookup instantiates a new ObservabilityPipelineEnrichmentTableFieldSecretLookup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableFieldSecretLookup(secret string) *ObservabilityPipelineEnrichmentTableFieldSecretLookup {
	this := ObservabilityPipelineEnrichmentTableFieldSecretLookup{}
	this.Secret = secret
	return &this
}

// NewObservabilityPipelineEnrichmentTableFieldSecretLookupWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableFieldSecretLookup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableFieldSecretLookupWithDefaults() *ObservabilityPipelineEnrichmentTableFieldSecretLookup {
	this := ObservabilityPipelineEnrichmentTableFieldSecretLookup{}
	return &this
}

// GetSecret returns the Secret field value.
func (o *ObservabilityPipelineEnrichmentTableFieldSecretLookup) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFieldSecretLookup) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value.
func (o *ObservabilityPipelineEnrichmentTableFieldSecretLookup) SetSecret(v string) {
	o.Secret = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableFieldSecretLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableFieldSecretLookup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Secret *string `json:"secret"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Secret == nil {
		return fmt.Errorf("required field secret missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"secret"})
	} else {
		return err
	}
	o.Secret = *all.Secret

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
