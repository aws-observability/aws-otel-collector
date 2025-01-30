// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerIncludedKeywordConfiguration Object defining a set of keywords and a number of characters that help reduce noise.
// You can provide a list of keywords you would like to check within a defined proximity of the matching pattern.
// If any of the keywords are found within the proximity check, the match is kept.
// If none are found, the match is discarded.
type SensitiveDataScannerIncludedKeywordConfiguration struct {
	// The number of characters behind a match detected by Sensitive Data Scanner to look for the keywords defined.
	// `character_count` should be greater than the maximum length of a keyword defined for a rule.
	CharacterCount int64 `json:"character_count"`
	// Keyword list that will be checked during scanning in order to validate a match.
	// The number of keywords in the list must be less than or equal to 30.
	Keywords []string `json:"keywords"`
	// Should the rule use the underlying standard pattern keyword configuration. If set to `true`, the rule must be tied
	// to a standard pattern. If set to `false`, the specified keywords and `character_count` are applied.
	UseRecommendedKeywords *bool `json:"use_recommended_keywords,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerIncludedKeywordConfiguration instantiates a new SensitiveDataScannerIncludedKeywordConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerIncludedKeywordConfiguration(characterCount int64, keywords []string) *SensitiveDataScannerIncludedKeywordConfiguration {
	this := SensitiveDataScannerIncludedKeywordConfiguration{}
	this.CharacterCount = characterCount
	this.Keywords = keywords
	return &this
}

// NewSensitiveDataScannerIncludedKeywordConfigurationWithDefaults instantiates a new SensitiveDataScannerIncludedKeywordConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerIncludedKeywordConfigurationWithDefaults() *SensitiveDataScannerIncludedKeywordConfiguration {
	this := SensitiveDataScannerIncludedKeywordConfiguration{}
	return &this
}

// GetCharacterCount returns the CharacterCount field value.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) GetCharacterCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CharacterCount
}

// GetCharacterCountOk returns a tuple with the CharacterCount field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) GetCharacterCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CharacterCount, true
}

// SetCharacterCount sets field value.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) SetCharacterCount(v int64) {
	o.CharacterCount = v
}

// GetKeywords returns the Keywords field value.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) GetKeywords() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) GetKeywordsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keywords, true
}

// SetKeywords sets field value.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) SetKeywords(v []string) {
	o.Keywords = v
}

// GetUseRecommendedKeywords returns the UseRecommendedKeywords field value if set, zero value otherwise.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) GetUseRecommendedKeywords() bool {
	if o == nil || o.UseRecommendedKeywords == nil {
		var ret bool
		return ret
	}
	return *o.UseRecommendedKeywords
}

// GetUseRecommendedKeywordsOk returns a tuple with the UseRecommendedKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) GetUseRecommendedKeywordsOk() (*bool, bool) {
	if o == nil || o.UseRecommendedKeywords == nil {
		return nil, false
	}
	return o.UseRecommendedKeywords, true
}

// HasUseRecommendedKeywords returns a boolean if a field has been set.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) HasUseRecommendedKeywords() bool {
	return o != nil && o.UseRecommendedKeywords != nil
}

// SetUseRecommendedKeywords gets a reference to the given bool and assigns it to the UseRecommendedKeywords field.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) SetUseRecommendedKeywords(v bool) {
	o.UseRecommendedKeywords = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerIncludedKeywordConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["character_count"] = o.CharacterCount
	toSerialize["keywords"] = o.Keywords
	if o.UseRecommendedKeywords != nil {
		toSerialize["use_recommended_keywords"] = o.UseRecommendedKeywords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerIncludedKeywordConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CharacterCount         *int64    `json:"character_count"`
		Keywords               *[]string `json:"keywords"`
		UseRecommendedKeywords *bool     `json:"use_recommended_keywords,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CharacterCount == nil {
		return fmt.Errorf("required field character_count missing")
	}
	if all.Keywords == nil {
		return fmt.Errorf("required field keywords missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"character_count", "keywords", "use_recommended_keywords"})
	} else {
		return err
	}
	o.CharacterCount = *all.CharacterCount
	o.Keywords = *all.Keywords
	o.UseRecommendedKeywords = all.UseRecommendedKeywords

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
