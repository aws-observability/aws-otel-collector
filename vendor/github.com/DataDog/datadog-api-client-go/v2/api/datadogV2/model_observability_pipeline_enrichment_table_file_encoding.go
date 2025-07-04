// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFileEncoding File encoding format.
type ObservabilityPipelineEnrichmentTableFileEncoding struct {
	// The `encoding` `delimiter`.
	Delimiter string `json:"delimiter"`
	// The `encoding` `includes_headers`.
	IncludesHeaders bool `json:"includes_headers"`
	// Specifies the encoding format (e.g., CSV) used for enrichment tables.
	Type ObservabilityPipelineEnrichmentTableFileEncodingType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableFileEncoding instantiates a new ObservabilityPipelineEnrichmentTableFileEncoding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableFileEncoding(delimiter string, includesHeaders bool, typeVar ObservabilityPipelineEnrichmentTableFileEncodingType) *ObservabilityPipelineEnrichmentTableFileEncoding {
	this := ObservabilityPipelineEnrichmentTableFileEncoding{}
	this.Delimiter = delimiter
	this.IncludesHeaders = includesHeaders
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineEnrichmentTableFileEncodingWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableFileEncoding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableFileEncodingWithDefaults() *ObservabilityPipelineEnrichmentTableFileEncoding {
	this := ObservabilityPipelineEnrichmentTableFileEncoding{}
	return &this
}

// GetDelimiter returns the Delimiter field value.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) GetDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) GetDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delimiter, true
}

// SetDelimiter sets field value.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) SetDelimiter(v string) {
	o.Delimiter = v
}

// GetIncludesHeaders returns the IncludesHeaders field value.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) GetIncludesHeaders() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IncludesHeaders
}

// GetIncludesHeadersOk returns a tuple with the IncludesHeaders field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) GetIncludesHeadersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludesHeaders, true
}

// SetIncludesHeaders sets field value.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) SetIncludesHeaders(v bool) {
	o.IncludesHeaders = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) GetType() ObservabilityPipelineEnrichmentTableFileEncodingType {
	if o == nil {
		var ret ObservabilityPipelineEnrichmentTableFileEncodingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) GetTypeOk() (*ObservabilityPipelineEnrichmentTableFileEncodingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) SetType(v ObservabilityPipelineEnrichmentTableFileEncodingType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableFileEncoding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["delimiter"] = o.Delimiter
	toSerialize["includes_headers"] = o.IncludesHeaders
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableFileEncoding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Delimiter       *string                                               `json:"delimiter"`
		IncludesHeaders *bool                                                 `json:"includes_headers"`
		Type            *ObservabilityPipelineEnrichmentTableFileEncodingType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Delimiter == nil {
		return fmt.Errorf("required field delimiter missing")
	}
	if all.IncludesHeaders == nil {
		return fmt.Errorf("required field includes_headers missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delimiter", "includes_headers", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Delimiter = *all.Delimiter
	o.IncludesHeaders = *all.IncludesHeaders
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
