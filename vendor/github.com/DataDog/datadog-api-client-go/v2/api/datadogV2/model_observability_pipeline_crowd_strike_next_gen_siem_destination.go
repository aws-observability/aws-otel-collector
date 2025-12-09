// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCrowdStrikeNextGenSiemDestination The `crowdstrike_next_gen_siem` destination forwards logs to CrowdStrike Next Gen SIEM.
type ObservabilityPipelineCrowdStrikeNextGenSiemDestination struct {
	// Compression configuration for log events.
	Compression *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression `json:"compression,omitempty"`
	// Encoding format for log events.
	Encoding ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding `json:"encoding"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. The value should always be `crowdstrike_next_gen_siem`.
	Type ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCrowdStrikeNextGenSiemDestination instantiates a new ObservabilityPipelineCrowdStrikeNextGenSiemDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCrowdStrikeNextGenSiemDestination(encoding ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding, id string, inputs []string, typeVar ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType) *ObservabilityPipelineCrowdStrikeNextGenSiemDestination {
	this := ObservabilityPipelineCrowdStrikeNextGenSiemDestination{}
	this.Encoding = encoding
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationWithDefaults instantiates a new ObservabilityPipelineCrowdStrikeNextGenSiemDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationWithDefaults() *ObservabilityPipelineCrowdStrikeNextGenSiemDestination {
	this := ObservabilityPipelineCrowdStrikeNextGenSiemDestination{}
	var typeVar ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType = OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONTYPE_CROWDSTRIKE_NEXT_GEN_SIEM
	this.Type = typeVar
	return &this
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetCompression() ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetCompressionOk() (*ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) SetCompression(v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression) {
	o.Compression = &v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetEncoding() ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding {
	if o == nil {
		var ret ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetEncodingOk() (*ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) SetEncoding(v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding) {
	o.Encoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetType() ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType {
	if o == nil {
		var ret ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) GetTypeOk() (*ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) SetType(v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCrowdStrikeNextGenSiemDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	toSerialize["encoding"] = o.Encoding
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
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
func (o *ObservabilityPipelineCrowdStrikeNextGenSiemDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compression *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompression `json:"compression,omitempty"`
		Encoding    *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationEncoding    `json:"encoding"`
		Id          *string                                                            `json:"id"`
		Inputs      *[]string                                                          `json:"inputs"`
		Tls         *ObservabilityPipelineTls                                          `json:"tls,omitempty"`
		Type        *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Encoding == nil {
		return fmt.Errorf("required field encoding missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"compression", "encoding", "id", "inputs", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compression != nil && all.Compression.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compression = all.Compression
	if !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = *all.Encoding
	}
	o.Id = *all.Id
	o.Inputs = *all.Inputs
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
