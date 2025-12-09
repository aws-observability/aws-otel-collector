// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions Configuration for keywords used to reinforce sensitive data pattern detection.
type ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions struct {
	// A list of keywords to match near the sensitive pattern.
	Keywords []string `json:"keywords"`
	// Maximum number of tokens between a keyword and a sensitive value match.
	Proximity int64 `json:"proximity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions(keywords []string, proximity int64) *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions{}
	this.Keywords = keywords
	this.Proximity = proximity
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorKeywordOptionsWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorKeywordOptionsWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions{}
	return &this
}

// GetKeywords returns the Keywords field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) GetKeywords() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) GetKeywordsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) SetKeywords(v []string) {
	o.Keywords = v
}

// GetProximity returns the Proximity field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) GetProximity() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Proximity
}

// GetProximityOk returns a tuple with the Proximity field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) GetProximityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proximity, true
}

// SetProximity sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) SetProximity(v int64) {
	o.Proximity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["keywords"] = o.Keywords
	toSerialize["proximity"] = o.Proximity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Keywords  *[]string `json:"keywords"`
		Proximity *int64    `json:"proximity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Keywords == nil {
		return fmt.Errorf("required field keywords missing")
	}
	if all.Proximity == nil {
		return fmt.Errorf("required field proximity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"keywords", "proximity"})
	} else {
		return err
	}
	o.Keywords = *all.Keywords
	o.Proximity = *all.Proximity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
