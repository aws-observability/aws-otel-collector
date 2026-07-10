// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestinationBatchEncoding Batch encoding configuration for the ClickHouse destination.
// Required when `format` is `arrow_stream`. The `codec` field must be set to `arrow_stream`.
type ObservabilityPipelineClickhouseDestinationBatchEncoding struct {
	// When `true`, null values are allowed for non-nullable fields in the ClickHouse schema.
	// When `false` (default), missing values for non-nullable columns cause encoding errors.
	AllowNullableFields *bool `json:"allow_nullable_fields,omitempty"`
	// The codec used for batch encoding. Only `arrow_stream` is supported.
	Codec ObservabilityPipelineClickhouseDestinationBatchEncodingCodec `json:"codec"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineClickhouseDestinationBatchEncoding instantiates a new ObservabilityPipelineClickhouseDestinationBatchEncoding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineClickhouseDestinationBatchEncoding(codec ObservabilityPipelineClickhouseDestinationBatchEncodingCodec) *ObservabilityPipelineClickhouseDestinationBatchEncoding {
	this := ObservabilityPipelineClickhouseDestinationBatchEncoding{}
	this.Codec = codec
	return &this
}

// NewObservabilityPipelineClickhouseDestinationBatchEncodingWithDefaults instantiates a new ObservabilityPipelineClickhouseDestinationBatchEncoding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineClickhouseDestinationBatchEncodingWithDefaults() *ObservabilityPipelineClickhouseDestinationBatchEncoding {
	this := ObservabilityPipelineClickhouseDestinationBatchEncoding{}
	return &this
}

// GetAllowNullableFields returns the AllowNullableFields field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) GetAllowNullableFields() bool {
	if o == nil || o.AllowNullableFields == nil {
		var ret bool
		return ret
	}
	return *o.AllowNullableFields
}

// GetAllowNullableFieldsOk returns a tuple with the AllowNullableFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) GetAllowNullableFieldsOk() (*bool, bool) {
	if o == nil || o.AllowNullableFields == nil {
		return nil, false
	}
	return o.AllowNullableFields, true
}

// HasAllowNullableFields returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) HasAllowNullableFields() bool {
	return o != nil && o.AllowNullableFields != nil
}

// SetAllowNullableFields gets a reference to the given bool and assigns it to the AllowNullableFields field.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) SetAllowNullableFields(v bool) {
	o.AllowNullableFields = &v
}

// GetCodec returns the Codec field value.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) GetCodec() ObservabilityPipelineClickhouseDestinationBatchEncodingCodec {
	if o == nil {
		var ret ObservabilityPipelineClickhouseDestinationBatchEncodingCodec
		return ret
	}
	return o.Codec
}

// GetCodecOk returns a tuple with the Codec field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) GetCodecOk() (*ObservabilityPipelineClickhouseDestinationBatchEncodingCodec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codec, true
}

// SetCodec sets field value.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) SetCodec(v ObservabilityPipelineClickhouseDestinationBatchEncodingCodec) {
	o.Codec = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineClickhouseDestinationBatchEncoding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowNullableFields != nil {
		toSerialize["allow_nullable_fields"] = o.AllowNullableFields
	}
	toSerialize["codec"] = o.Codec

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineClickhouseDestinationBatchEncoding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowNullableFields *bool                                                         `json:"allow_nullable_fields,omitempty"`
		Codec               *ObservabilityPipelineClickhouseDestinationBatchEncodingCodec `json:"codec"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Codec == nil {
		return fmt.Errorf("required field codec missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_nullable_fields", "codec"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowNullableFields = all.AllowNullableFields
	if !all.Codec.IsValid() {
		hasInvalidField = true
	} else {
		o.Codec = *all.Codec
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
