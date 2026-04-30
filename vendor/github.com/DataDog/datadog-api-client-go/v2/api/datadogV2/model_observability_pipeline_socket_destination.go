// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestination The `socket` destination sends logs over TCP or UDP to a remote server.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineSocketDestination struct {
	// Name of the environment variable or secret that holds the socket address (host:port).
	AddressKey *string `json:"address_key,omitempty"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// Encoding format for log events.
	Encoding ObservabilityPipelineSocketDestinationEncoding `json:"encoding"`
	// Framing method configuration.
	Framing ObservabilityPipelineSocketDestinationFraming `json:"framing"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Protocol used to send logs.
	Mode ObservabilityPipelineSocketDestinationMode `json:"mode"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. The value should always be `socket`.
	Type ObservabilityPipelineSocketDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSocketDestination instantiates a new ObservabilityPipelineSocketDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSocketDestination(encoding ObservabilityPipelineSocketDestinationEncoding, framing ObservabilityPipelineSocketDestinationFraming, id string, inputs []string, mode ObservabilityPipelineSocketDestinationMode, typeVar ObservabilityPipelineSocketDestinationType) *ObservabilityPipelineSocketDestination {
	this := ObservabilityPipelineSocketDestination{}
	this.Encoding = encoding
	this.Framing = framing
	this.Id = id
	this.Inputs = inputs
	this.Mode = mode
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSocketDestinationWithDefaults instantiates a new ObservabilityPipelineSocketDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSocketDestinationWithDefaults() *ObservabilityPipelineSocketDestination {
	this := ObservabilityPipelineSocketDestination{}
	var typeVar ObservabilityPipelineSocketDestinationType = OBSERVABILITYPIPELINESOCKETDESTINATIONTYPE_SOCKET
	this.Type = typeVar
	return &this
}

// GetAddressKey returns the AddressKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineSocketDestination) GetAddressKey() string {
	if o == nil || o.AddressKey == nil {
		var ret string
		return ret
	}
	return *o.AddressKey
}

// GetAddressKeyOk returns a tuple with the AddressKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetAddressKeyOk() (*string, bool) {
	if o == nil || o.AddressKey == nil {
		return nil, false
	}
	return o.AddressKey, true
}

// HasAddressKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineSocketDestination) HasAddressKey() bool {
	return o != nil && o.AddressKey != nil
}

// SetAddressKey gets a reference to the given string and assigns it to the AddressKey field.
func (o *ObservabilityPipelineSocketDestination) SetAddressKey(v string) {
	o.AddressKey = &v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineSocketDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineSocketDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineSocketDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineSocketDestination) GetEncoding() ObservabilityPipelineSocketDestinationEncoding {
	if o == nil {
		var ret ObservabilityPipelineSocketDestinationEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetEncodingOk() (*ObservabilityPipelineSocketDestinationEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineSocketDestination) SetEncoding(v ObservabilityPipelineSocketDestinationEncoding) {
	o.Encoding = v
}

// GetFraming returns the Framing field value.
func (o *ObservabilityPipelineSocketDestination) GetFraming() ObservabilityPipelineSocketDestinationFraming {
	if o == nil {
		var ret ObservabilityPipelineSocketDestinationFraming
		return ret
	}
	return o.Framing
}

// GetFramingOk returns a tuple with the Framing field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetFramingOk() (*ObservabilityPipelineSocketDestinationFraming, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Framing, true
}

// SetFraming sets field value.
func (o *ObservabilityPipelineSocketDestination) SetFraming(v ObservabilityPipelineSocketDestinationFraming) {
	o.Framing = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSocketDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSocketDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSocketDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSocketDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineSocketDestination) GetMode() ObservabilityPipelineSocketDestinationMode {
	if o == nil {
		var ret ObservabilityPipelineSocketDestinationMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetModeOk() (*ObservabilityPipelineSocketDestinationMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineSocketDestination) SetMode(v ObservabilityPipelineSocketDestinationMode) {
	o.Mode = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineSocketDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineSocketDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineSocketDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSocketDestination) GetType() ObservabilityPipelineSocketDestinationType {
	if o == nil {
		var ret ObservabilityPipelineSocketDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSocketDestination) GetTypeOk() (*ObservabilityPipelineSocketDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSocketDestination) SetType(v ObservabilityPipelineSocketDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSocketDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddressKey != nil {
		toSerialize["address_key"] = o.AddressKey
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	toSerialize["encoding"] = o.Encoding
	toSerialize["framing"] = o.Framing
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["mode"] = o.Mode
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
func (o *ObservabilityPipelineSocketDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddressKey *string                                         `json:"address_key,omitempty"`
		Buffer     *ObservabilityPipelineBufferOptions             `json:"buffer,omitempty"`
		Encoding   *ObservabilityPipelineSocketDestinationEncoding `json:"encoding"`
		Framing    *ObservabilityPipelineSocketDestinationFraming  `json:"framing"`
		Id         *string                                         `json:"id"`
		Inputs     *[]string                                       `json:"inputs"`
		Mode       *ObservabilityPipelineSocketDestinationMode     `json:"mode"`
		Tls        *ObservabilityPipelineTls                       `json:"tls,omitempty"`
		Type       *ObservabilityPipelineSocketDestinationType     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Encoding == nil {
		return fmt.Errorf("required field encoding missing")
	}
	if all.Framing == nil {
		return fmt.Errorf("required field framing missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"address_key", "buffer", "encoding", "framing", "id", "inputs", "mode", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AddressKey = all.AddressKey
	o.Buffer = all.Buffer
	if !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = *all.Encoding
	}
	o.Framing = *all.Framing
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
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
