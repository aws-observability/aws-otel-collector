// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerSuppressions Object describing the suppressions for a rule. There are three types of suppressions, `starts_with`, `ends_with`, and `exact_match`.
// Suppressed matches are not obfuscated, counted in metrics, or displayed in the Findings page.
type SensitiveDataScannerSuppressions struct {
	// List of strings to use for suppression of matches ending with these strings.
	EndsWith []string `json:"ends_with,omitempty"`
	// List of strings to use for suppression of matches exactly matching these strings.
	ExactMatch []string `json:"exact_match,omitempty"`
	// List of strings to use for suppression of matches starting with these strings.
	StartsWith []string `json:"starts_with,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerSuppressions instantiates a new SensitiveDataScannerSuppressions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerSuppressions() *SensitiveDataScannerSuppressions {
	this := SensitiveDataScannerSuppressions{}
	return &this
}

// NewSensitiveDataScannerSuppressionsWithDefaults instantiates a new SensitiveDataScannerSuppressions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerSuppressionsWithDefaults() *SensitiveDataScannerSuppressions {
	this := SensitiveDataScannerSuppressions{}
	return &this
}

// GetEndsWith returns the EndsWith field value if set, zero value otherwise.
func (o *SensitiveDataScannerSuppressions) GetEndsWith() []string {
	if o == nil || o.EndsWith == nil {
		var ret []string
		return ret
	}
	return o.EndsWith
}

// GetEndsWithOk returns a tuple with the EndsWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerSuppressions) GetEndsWithOk() (*[]string, bool) {
	if o == nil || o.EndsWith == nil {
		return nil, false
	}
	return &o.EndsWith, true
}

// HasEndsWith returns a boolean if a field has been set.
func (o *SensitiveDataScannerSuppressions) HasEndsWith() bool {
	return o != nil && o.EndsWith != nil
}

// SetEndsWith gets a reference to the given []string and assigns it to the EndsWith field.
func (o *SensitiveDataScannerSuppressions) SetEndsWith(v []string) {
	o.EndsWith = v
}

// GetExactMatch returns the ExactMatch field value if set, zero value otherwise.
func (o *SensitiveDataScannerSuppressions) GetExactMatch() []string {
	if o == nil || o.ExactMatch == nil {
		var ret []string
		return ret
	}
	return o.ExactMatch
}

// GetExactMatchOk returns a tuple with the ExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerSuppressions) GetExactMatchOk() (*[]string, bool) {
	if o == nil || o.ExactMatch == nil {
		return nil, false
	}
	return &o.ExactMatch, true
}

// HasExactMatch returns a boolean if a field has been set.
func (o *SensitiveDataScannerSuppressions) HasExactMatch() bool {
	return o != nil && o.ExactMatch != nil
}

// SetExactMatch gets a reference to the given []string and assigns it to the ExactMatch field.
func (o *SensitiveDataScannerSuppressions) SetExactMatch(v []string) {
	o.ExactMatch = v
}

// GetStartsWith returns the StartsWith field value if set, zero value otherwise.
func (o *SensitiveDataScannerSuppressions) GetStartsWith() []string {
	if o == nil || o.StartsWith == nil {
		var ret []string
		return ret
	}
	return o.StartsWith
}

// GetStartsWithOk returns a tuple with the StartsWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerSuppressions) GetStartsWithOk() (*[]string, bool) {
	if o == nil || o.StartsWith == nil {
		return nil, false
	}
	return &o.StartsWith, true
}

// HasStartsWith returns a boolean if a field has been set.
func (o *SensitiveDataScannerSuppressions) HasStartsWith() bool {
	return o != nil && o.StartsWith != nil
}

// SetStartsWith gets a reference to the given []string and assigns it to the StartsWith field.
func (o *SensitiveDataScannerSuppressions) SetStartsWith(v []string) {
	o.StartsWith = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerSuppressions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EndsWith != nil {
		toSerialize["ends_with"] = o.EndsWith
	}
	if o.ExactMatch != nil {
		toSerialize["exact_match"] = o.ExactMatch
	}
	if o.StartsWith != nil {
		toSerialize["starts_with"] = o.StartsWith
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerSuppressions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndsWith   []string `json:"ends_with,omitempty"`
		ExactMatch []string `json:"exact_match,omitempty"`
		StartsWith []string `json:"starts_with,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ends_with", "exact_match", "starts_with"})
	} else {
		return err
	}
	o.EndsWith = all.EndsWith
	o.ExactMatch = all.ExactMatch
	o.StartsWith = all.StartsWith

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
