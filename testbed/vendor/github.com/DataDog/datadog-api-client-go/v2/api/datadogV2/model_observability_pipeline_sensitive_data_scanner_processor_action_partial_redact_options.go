// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions Controls how partial redaction is applied, including character count and direction.
type ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions struct {
	// The `ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions` `characters`.
	Characters int64 `json:"characters"`
	// Indicates whether to redact characters from the first or last part of the matched value.
	Direction ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection `json:"direction"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions(characters int64, direction ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection) *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions{}
	this.Characters = characters
	this.Direction = direction
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions {
	this := ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions{}
	return &this
}

// GetCharacters returns the Characters field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) GetCharacters() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Characters
}

// GetCharactersOk returns a tuple with the Characters field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) GetCharactersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Characters, true
}

// SetCharacters sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) SetCharacters(v int64) {
	o.Characters = v
}

// GetDirection returns the Direction field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) GetDirection() ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection
		return ret
	}
	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) GetDirectionOk() (*ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) SetDirection(v ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection) {
	o.Direction = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["characters"] = o.Characters
	toSerialize["direction"] = o.Direction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Characters *int64                                                                                 `json:"characters"`
		Direction  *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactOptionsDirection `json:"direction"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Characters == nil {
		return fmt.Errorf("required field characters missing")
	}
	if all.Direction == nil {
		return fmt.Errorf("required field direction missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"characters", "direction"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Characters = *all.Characters
	if !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = *all.Direction
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
