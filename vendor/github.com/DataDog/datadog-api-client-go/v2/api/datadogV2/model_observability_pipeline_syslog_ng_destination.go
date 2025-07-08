// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSyslogNgDestination The `syslog_ng` destination forwards logs to an external `syslog-ng` server over TCP or UDP using the syslog protocol.
type ObservabilityPipelineSyslogNgDestination struct {
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Optional socket keepalive duration in milliseconds.
	Keepalive *int64 `json:"keepalive,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. The value should always be `syslog_ng`.
	Type ObservabilityPipelineSyslogNgDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSyslogNgDestination instantiates a new ObservabilityPipelineSyslogNgDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSyslogNgDestination(id string, inputs []string, typeVar ObservabilityPipelineSyslogNgDestinationType) *ObservabilityPipelineSyslogNgDestination {
	this := ObservabilityPipelineSyslogNgDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSyslogNgDestinationWithDefaults instantiates a new ObservabilityPipelineSyslogNgDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSyslogNgDestinationWithDefaults() *ObservabilityPipelineSyslogNgDestination {
	this := ObservabilityPipelineSyslogNgDestination{}
	var typeVar ObservabilityPipelineSyslogNgDestinationType = OBSERVABILITYPIPELINESYSLOGNGDESTINATIONTYPE_SYSLOG_NG
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSyslogNgDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSyslogNgDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSyslogNgDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSyslogNgDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetKeepalive returns the Keepalive field value if set, zero value otherwise.
func (o *ObservabilityPipelineSyslogNgDestination) GetKeepalive() int64 {
	if o == nil || o.Keepalive == nil {
		var ret int64
		return ret
	}
	return *o.Keepalive
}

// GetKeepaliveOk returns a tuple with the Keepalive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgDestination) GetKeepaliveOk() (*int64, bool) {
	if o == nil || o.Keepalive == nil {
		return nil, false
	}
	return o.Keepalive, true
}

// HasKeepalive returns a boolean if a field has been set.
func (o *ObservabilityPipelineSyslogNgDestination) HasKeepalive() bool {
	return o != nil && o.Keepalive != nil
}

// SetKeepalive gets a reference to the given int64 and assigns it to the Keepalive field.
func (o *ObservabilityPipelineSyslogNgDestination) SetKeepalive(v int64) {
	o.Keepalive = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineSyslogNgDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineSyslogNgDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineSyslogNgDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSyslogNgDestination) GetType() ObservabilityPipelineSyslogNgDestinationType {
	if o == nil {
		var ret ObservabilityPipelineSyslogNgDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgDestination) GetTypeOk() (*ObservabilityPipelineSyslogNgDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSyslogNgDestination) SetType(v ObservabilityPipelineSyslogNgDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSyslogNgDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.Keepalive != nil {
		toSerialize["keepalive"] = o.Keepalive
	}
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
func (o *ObservabilityPipelineSyslogNgDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string                                       `json:"id"`
		Inputs    *[]string                                     `json:"inputs"`
		Keepalive *int64                                        `json:"keepalive,omitempty"`
		Tls       *ObservabilityPipelineTls                     `json:"tls,omitempty"`
		Type      *ObservabilityPipelineSyslogNgDestinationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "inputs", "keepalive", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.Keepalive = all.Keepalive
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
