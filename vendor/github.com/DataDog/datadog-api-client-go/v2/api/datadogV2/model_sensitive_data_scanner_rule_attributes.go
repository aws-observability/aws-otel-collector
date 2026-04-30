// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerRuleAttributes Attributes of the Sensitive Data Scanner rule.
type SensitiveDataScannerRuleAttributes struct {
	// Description of the rule.
	Description *string `json:"description,omitempty"`
	// Attributes excluded from the scan. If namespaces is provided, it has to be a sub-path of the namespaces array.
	ExcludedNamespaces []string `json:"excluded_namespaces,omitempty"`
	// Object defining a set of keywords and a number of characters that help reduce noise.
	// You can provide a list of keywords you would like to check within a defined proximity of the matching pattern.
	// If any of the keywords are found within the proximity check, the match is kept.
	// If none are found, the match is discarded.
	IncludedKeywordConfiguration *SensitiveDataScannerIncludedKeywordConfiguration `json:"included_keyword_configuration,omitempty"`
	// Whether or not the rule is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Attributes included in the scan. If namespaces is empty or missing, all attributes except excluded_namespaces are scanned.
	// If both are missing the whole event is scanned.
	Namespaces []string `json:"namespaces,omitempty"`
	// Not included if there is a relationship to a standard pattern.
	Pattern *string `json:"pattern,omitempty"`
	// Integer from 1 (high) to 5 (low) indicating rule issue severity.
	Priority *int64 `json:"priority,omitempty"`
	// Object describing the suppressions for a rule. There are three types of suppressions, `starts_with`, `ends_with`, and `exact_match`.
	// Suppressed matches are not obfuscated, counted in metrics, or displayed in the Findings page.
	Suppressions *SensitiveDataScannerSuppressions `json:"suppressions,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	// Object describing how the scanned event will be replaced.
	TextReplacement *SensitiveDataScannerTextReplacement `json:"text_replacement,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerRuleAttributes instantiates a new SensitiveDataScannerRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerRuleAttributes() *SensitiveDataScannerRuleAttributes {
	this := SensitiveDataScannerRuleAttributes{}
	return &this
}

// NewSensitiveDataScannerRuleAttributesWithDefaults instantiates a new SensitiveDataScannerRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerRuleAttributesWithDefaults() *SensitiveDataScannerRuleAttributes {
	this := SensitiveDataScannerRuleAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SensitiveDataScannerRuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExcludedNamespaces returns the ExcludedNamespaces field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetExcludedNamespaces() []string {
	if o == nil || o.ExcludedNamespaces == nil {
		var ret []string
		return ret
	}
	return o.ExcludedNamespaces
}

// GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetExcludedNamespacesOk() (*[]string, bool) {
	if o == nil || o.ExcludedNamespaces == nil {
		return nil, false
	}
	return &o.ExcludedNamespaces, true
}

// HasExcludedNamespaces returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasExcludedNamespaces() bool {
	return o != nil && o.ExcludedNamespaces != nil
}

// SetExcludedNamespaces gets a reference to the given []string and assigns it to the ExcludedNamespaces field.
func (o *SensitiveDataScannerRuleAttributes) SetExcludedNamespaces(v []string) {
	o.ExcludedNamespaces = v
}

// GetIncludedKeywordConfiguration returns the IncludedKeywordConfiguration field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetIncludedKeywordConfiguration() SensitiveDataScannerIncludedKeywordConfiguration {
	if o == nil || o.IncludedKeywordConfiguration == nil {
		var ret SensitiveDataScannerIncludedKeywordConfiguration
		return ret
	}
	return *o.IncludedKeywordConfiguration
}

// GetIncludedKeywordConfigurationOk returns a tuple with the IncludedKeywordConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetIncludedKeywordConfigurationOk() (*SensitiveDataScannerIncludedKeywordConfiguration, bool) {
	if o == nil || o.IncludedKeywordConfiguration == nil {
		return nil, false
	}
	return o.IncludedKeywordConfiguration, true
}

// HasIncludedKeywordConfiguration returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasIncludedKeywordConfiguration() bool {
	return o != nil && o.IncludedKeywordConfiguration != nil
}

// SetIncludedKeywordConfiguration gets a reference to the given SensitiveDataScannerIncludedKeywordConfiguration and assigns it to the IncludedKeywordConfiguration field.
func (o *SensitiveDataScannerRuleAttributes) SetIncludedKeywordConfiguration(v SensitiveDataScannerIncludedKeywordConfiguration) {
	o.IncludedKeywordConfiguration = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SensitiveDataScannerRuleAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensitiveDataScannerRuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasNamespaces() bool {
	return o != nil && o.Namespaces != nil
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *SensitiveDataScannerRuleAttributes) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SensitiveDataScannerRuleAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *SensitiveDataScannerRuleAttributes) SetPriority(v int64) {
	o.Priority = &v
}

// GetSuppressions returns the Suppressions field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetSuppressions() SensitiveDataScannerSuppressions {
	if o == nil || o.Suppressions == nil {
		var ret SensitiveDataScannerSuppressions
		return ret
	}
	return *o.Suppressions
}

// GetSuppressionsOk returns a tuple with the Suppressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetSuppressionsOk() (*SensitiveDataScannerSuppressions, bool) {
	if o == nil || o.Suppressions == nil {
		return nil, false
	}
	return o.Suppressions, true
}

// HasSuppressions returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasSuppressions() bool {
	return o != nil && o.Suppressions != nil
}

// SetSuppressions gets a reference to the given SensitiveDataScannerSuppressions and assigns it to the Suppressions field.
func (o *SensitiveDataScannerRuleAttributes) SetSuppressions(v SensitiveDataScannerSuppressions) {
	o.Suppressions = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SensitiveDataScannerRuleAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTextReplacement returns the TextReplacement field value if set, zero value otherwise.
func (o *SensitiveDataScannerRuleAttributes) GetTextReplacement() SensitiveDataScannerTextReplacement {
	if o == nil || o.TextReplacement == nil {
		var ret SensitiveDataScannerTextReplacement
		return ret
	}
	return *o.TextReplacement
}

// GetTextReplacementOk returns a tuple with the TextReplacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerRuleAttributes) GetTextReplacementOk() (*SensitiveDataScannerTextReplacement, bool) {
	if o == nil || o.TextReplacement == nil {
		return nil, false
	}
	return o.TextReplacement, true
}

// HasTextReplacement returns a boolean if a field has been set.
func (o *SensitiveDataScannerRuleAttributes) HasTextReplacement() bool {
	return o != nil && o.TextReplacement != nil
}

// SetTextReplacement gets a reference to the given SensitiveDataScannerTextReplacement and assigns it to the TextReplacement field.
func (o *SensitiveDataScannerRuleAttributes) SetTextReplacement(v SensitiveDataScannerTextReplacement) {
	o.TextReplacement = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExcludedNamespaces != nil {
		toSerialize["excluded_namespaces"] = o.ExcludedNamespaces
	}
	if o.IncludedKeywordConfiguration != nil {
		toSerialize["included_keyword_configuration"] = o.IncludedKeywordConfiguration
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Suppressions != nil {
		toSerialize["suppressions"] = o.Suppressions
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TextReplacement != nil {
		toSerialize["text_replacement"] = o.TextReplacement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description                  *string                                           `json:"description,omitempty"`
		ExcludedNamespaces           []string                                          `json:"excluded_namespaces,omitempty"`
		IncludedKeywordConfiguration *SensitiveDataScannerIncludedKeywordConfiguration `json:"included_keyword_configuration,omitempty"`
		IsEnabled                    *bool                                             `json:"is_enabled,omitempty"`
		Name                         *string                                           `json:"name,omitempty"`
		Namespaces                   []string                                          `json:"namespaces,omitempty"`
		Pattern                      *string                                           `json:"pattern,omitempty"`
		Priority                     *int64                                            `json:"priority,omitempty"`
		Suppressions                 *SensitiveDataScannerSuppressions                 `json:"suppressions,omitempty"`
		Tags                         []string                                          `json:"tags,omitempty"`
		TextReplacement              *SensitiveDataScannerTextReplacement              `json:"text_replacement,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "excluded_namespaces", "included_keyword_configuration", "is_enabled", "name", "namespaces", "pattern", "priority", "suppressions", "tags", "text_replacement"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.ExcludedNamespaces = all.ExcludedNamespaces
	if all.IncludedKeywordConfiguration != nil && all.IncludedKeywordConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IncludedKeywordConfiguration = all.IncludedKeywordConfiguration
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.Namespaces = all.Namespaces
	o.Pattern = all.Pattern
	o.Priority = all.Priority
	if all.Suppressions != nil && all.Suppressions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Suppressions = all.Suppressions
	o.Tags = all.Tags
	if all.TextReplacement != nil && all.TextReplacement.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TextReplacement = all.TextReplacement

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
