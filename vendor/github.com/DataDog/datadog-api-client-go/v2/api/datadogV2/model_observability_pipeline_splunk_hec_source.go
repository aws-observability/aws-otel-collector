// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecSource The `splunk_hec` source implements the Splunk HTTP Event Collector (HEC) API.
type ObservabilityPipelineSplunkHecSource struct {
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. Always `splunk_hec`.
	Type ObservabilityPipelineSplunkHecSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSplunkHecSource instantiates a new ObservabilityPipelineSplunkHecSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSplunkHecSource(id string, typeVar ObservabilityPipelineSplunkHecSourceType) *ObservabilityPipelineSplunkHecSource {
	this := ObservabilityPipelineSplunkHecSource{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSplunkHecSourceWithDefaults instantiates a new ObservabilityPipelineSplunkHecSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSplunkHecSourceWithDefaults() *ObservabilityPipelineSplunkHecSource {
	this := ObservabilityPipelineSplunkHecSource{}
	var typeVar ObservabilityPipelineSplunkHecSourceType = OBSERVABILITYPIPELINESPLUNKHECSOURCETYPE_SPLUNK_HEC
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSplunkHecSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSplunkHecSource) SetId(v string) {
	o.Id = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineSplunkHecSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSplunkHecSource) GetType() ObservabilityPipelineSplunkHecSourceType {
	if o == nil {
		var ret ObservabilityPipelineSplunkHecSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecSource) GetTypeOk() (*ObservabilityPipelineSplunkHecSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSplunkHecSource) SetType(v ObservabilityPipelineSplunkHecSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSplunkHecSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *ObservabilityPipelineSplunkHecSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                   `json:"id"`
		Tls  *ObservabilityPipelineTls                 `json:"tls,omitempty"`
		Type *ObservabilityPipelineSplunkHecSourceType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
