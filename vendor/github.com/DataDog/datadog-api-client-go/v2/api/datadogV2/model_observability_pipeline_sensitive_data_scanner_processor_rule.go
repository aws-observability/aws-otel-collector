// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorRule Defines a rule for detecting sensitive data, including matching pattern, scope, and the action to take.
type ObservabilityPipelineSensitiveDataScannerProcessorRule struct {
	// Configuration for keywords used to reinforce sensitive data pattern detection.
	KeywordOptions *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions `json:"keyword_options,omitempty"`
	// A name identifying the rule.
	Name string `json:"name"`
	// Defines what action to take when sensitive data is matched.
	OnMatch ObservabilityPipelineSensitiveDataScannerProcessorAction `json:"on_match"`
	// Pattern detection configuration for identifying sensitive data using either a custom regex or a library reference.
	Pattern ObservabilityPipelineSensitiveDataScannerProcessorPattern `json:"pattern"`
	// Determines which parts of the log the pattern-matching rule should be applied to.
	Scope ObservabilityPipelineSensitiveDataScannerProcessorScope `json:"scope"`
	// Tags assigned to this rule for filtering and classification.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSensitiveDataScannerProcessorRule instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSensitiveDataScannerProcessorRule(name string, onMatch ObservabilityPipelineSensitiveDataScannerProcessorAction, pattern ObservabilityPipelineSensitiveDataScannerProcessorPattern, scope ObservabilityPipelineSensitiveDataScannerProcessorScope, tags []string) *ObservabilityPipelineSensitiveDataScannerProcessorRule {
	this := ObservabilityPipelineSensitiveDataScannerProcessorRule{}
	this.Name = name
	this.OnMatch = onMatch
	this.Pattern = pattern
	this.Scope = scope
	this.Tags = tags
	return &this
}

// NewObservabilityPipelineSensitiveDataScannerProcessorRuleWithDefaults instantiates a new ObservabilityPipelineSensitiveDataScannerProcessorRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSensitiveDataScannerProcessorRuleWithDefaults() *ObservabilityPipelineSensitiveDataScannerProcessorRule {
	this := ObservabilityPipelineSensitiveDataScannerProcessorRule{}
	return &this
}

// GetKeywordOptions returns the KeywordOptions field value if set, zero value otherwise.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetKeywordOptions() ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions {
	if o == nil || o.KeywordOptions == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions
		return ret
	}
	return *o.KeywordOptions
}

// GetKeywordOptionsOk returns a tuple with the KeywordOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetKeywordOptionsOk() (*ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions, bool) {
	if o == nil || o.KeywordOptions == nil {
		return nil, false
	}
	return o.KeywordOptions, true
}

// HasKeywordOptions returns a boolean if a field has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) HasKeywordOptions() bool {
	return o != nil && o.KeywordOptions != nil
}

// SetKeywordOptions gets a reference to the given ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions and assigns it to the KeywordOptions field.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) SetKeywordOptions(v ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions) {
	o.KeywordOptions = &v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) SetName(v string) {
	o.Name = v
}

// GetOnMatch returns the OnMatch field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetOnMatch() ObservabilityPipelineSensitiveDataScannerProcessorAction {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorAction
		return ret
	}
	return o.OnMatch
}

// GetOnMatchOk returns a tuple with the OnMatch field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetOnMatchOk() (*ObservabilityPipelineSensitiveDataScannerProcessorAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnMatch, true
}

// SetOnMatch sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) SetOnMatch(v ObservabilityPipelineSensitiveDataScannerProcessorAction) {
	o.OnMatch = v
}

// GetPattern returns the Pattern field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetPattern() ObservabilityPipelineSensitiveDataScannerProcessorPattern {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorPattern
		return ret
	}
	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetPatternOk() (*ObservabilityPipelineSensitiveDataScannerProcessorPattern, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) SetPattern(v ObservabilityPipelineSensitiveDataScannerProcessorPattern) {
	o.Pattern = v
}

// GetScope returns the Scope field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetScope() ObservabilityPipelineSensitiveDataScannerProcessorScope {
	if o == nil {
		var ret ObservabilityPipelineSensitiveDataScannerProcessorScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetScopeOk() (*ObservabilityPipelineSensitiveDataScannerProcessorScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) SetScope(v ObservabilityPipelineSensitiveDataScannerProcessorScope) {
	o.Scope = v
}

// GetTags returns the Tags field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSensitiveDataScannerProcessorRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.KeywordOptions != nil {
		toSerialize["keyword_options"] = o.KeywordOptions
	}
	toSerialize["name"] = o.Name
	toSerialize["on_match"] = o.OnMatch
	toSerialize["pattern"] = o.Pattern
	toSerialize["scope"] = o.Scope
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSensitiveDataScannerProcessorRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KeywordOptions *ObservabilityPipelineSensitiveDataScannerProcessorKeywordOptions `json:"keyword_options,omitempty"`
		Name           *string                                                           `json:"name"`
		OnMatch        *ObservabilityPipelineSensitiveDataScannerProcessorAction         `json:"on_match"`
		Pattern        *ObservabilityPipelineSensitiveDataScannerProcessorPattern        `json:"pattern"`
		Scope          *ObservabilityPipelineSensitiveDataScannerProcessorScope          `json:"scope"`
		Tags           *[]string                                                         `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OnMatch == nil {
		return fmt.Errorf("required field on_match missing")
	}
	if all.Pattern == nil {
		return fmt.Errorf("required field pattern missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"keyword_options", "name", "on_match", "pattern", "scope", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.KeywordOptions != nil && all.KeywordOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.KeywordOptions = all.KeywordOptions
	o.Name = *all.Name
	o.OnMatch = *all.OnMatch
	o.Pattern = *all.Pattern
	o.Scope = *all.Scope
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
