// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFluentBitSource The `fluent_bit` source ingests logs from Fluent Bit.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineFluentBitSource struct {
	// Name of the environment variable or secret that holds the listen address for the Fluent Bit receiver.
	AddressKey *string `json:"address_key,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. The value should always be `fluent_bit`.
	Type ObservabilityPipelineFluentBitSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineFluentBitSource instantiates a new ObservabilityPipelineFluentBitSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineFluentBitSource(id string, typeVar ObservabilityPipelineFluentBitSourceType) *ObservabilityPipelineFluentBitSource {
	this := ObservabilityPipelineFluentBitSource{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineFluentBitSourceWithDefaults instantiates a new ObservabilityPipelineFluentBitSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineFluentBitSourceWithDefaults() *ObservabilityPipelineFluentBitSource {
	this := ObservabilityPipelineFluentBitSource{}
	var typeVar ObservabilityPipelineFluentBitSourceType = OBSERVABILITYPIPELINEFLUENTBITSOURCETYPE_FLUENT_BIT
	this.Type = typeVar
	return &this
}

// GetAddressKey returns the AddressKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineFluentBitSource) GetAddressKey() string {
	if o == nil || o.AddressKey == nil {
		var ret string
		return ret
	}
	return *o.AddressKey
}

// GetAddressKeyOk returns a tuple with the AddressKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFluentBitSource) GetAddressKeyOk() (*string, bool) {
	if o == nil || o.AddressKey == nil {
		return nil, false
	}
	return o.AddressKey, true
}

// HasAddressKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineFluentBitSource) HasAddressKey() bool {
	return o != nil && o.AddressKey != nil
}

// SetAddressKey gets a reference to the given string and assigns it to the AddressKey field.
func (o *ObservabilityPipelineFluentBitSource) SetAddressKey(v string) {
	o.AddressKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineFluentBitSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFluentBitSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineFluentBitSource) SetId(v string) {
	o.Id = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineFluentBitSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFluentBitSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineFluentBitSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineFluentBitSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineFluentBitSource) GetType() ObservabilityPipelineFluentBitSourceType {
	if o == nil {
		var ret ObservabilityPipelineFluentBitSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineFluentBitSource) GetTypeOk() (*ObservabilityPipelineFluentBitSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineFluentBitSource) SetType(v ObservabilityPipelineFluentBitSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineFluentBitSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddressKey != nil {
		toSerialize["address_key"] = o.AddressKey
	}
	toSerialize["id"] = o.Id
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineFluentBitSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddressKey *string                                   `json:"address_key,omitempty"`
		Id         *string                                   `json:"id"`
		Tls        *ObservabilityPipelineTls                 `json:"tls,omitempty"`
		Type       *ObservabilityPipelineFluentBitSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"address_key", "id", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AddressKey = all.AddressKey
	o.Id = *all.Id
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
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
