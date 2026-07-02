// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiPromptResponseAttributes Response attributes of an AI prompt.
type AiPromptResponseAttributes struct {
	// Rule category
	Category CustomRuleRevisionAttributesCategory `json:"category"`
	// Checksum of the prompt content.
	Checksum string `json:"checksum"`
	// Base64-encoded AI prompt content.
	Content string `json:"content"`
	// The CWE identifier associated with this prompt.
	Cwe *string `json:"cwe,omitempty"`
	// Base64-encoded full description.
	Description string `json:"description"`
	// Directory patterns this prompt applies to.
	Directories []string `json:"directories"`
	// The execution mode for an AI rule revision.
	ExecutionMode AiCustomRuleRevisionExecutionMode `json:"execution_mode"`
	// Keywords used to search for relevant files.
	FileSearchKeywords []string `json:"file_search_keywords"`
	// File glob patterns this prompt applies to.
	Globs []string `json:"globs"`
	// Whether this is a default Datadog prompt.
	IsDefault bool `json:"is_default"`
	// Whether this prompt is for testing only.
	IsTesting bool `json:"is_testing"`
	// Programming language
	Language *Language `json:"language,omitempty"`
	// Keywords to exclude from results.
	ResultKeywordsExclude []string `json:"result_keywords_exclude"`
	// The version of the rule this prompt is associated with.
	RuleVersion string `json:"rule_version"`
	// Rule severity
	Severity CustomRuleRevisionAttributesSeverity `json:"severity"`
	// Base64-encoded short description.
	ShortDescription string `json:"short_description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiPromptResponseAttributes instantiates a new AiPromptResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiPromptResponseAttributes(category CustomRuleRevisionAttributesCategory, checksum string, content string, description string, directories []string, executionMode AiCustomRuleRevisionExecutionMode, fileSearchKeywords []string, globs []string, isDefault bool, isTesting bool, resultKeywordsExclude []string, ruleVersion string, severity CustomRuleRevisionAttributesSeverity, shortDescription string) *AiPromptResponseAttributes {
	this := AiPromptResponseAttributes{}
	this.Category = category
	this.Checksum = checksum
	this.Content = content
	this.Description = description
	this.Directories = directories
	this.ExecutionMode = executionMode
	this.FileSearchKeywords = fileSearchKeywords
	this.Globs = globs
	this.IsDefault = isDefault
	this.IsTesting = isTesting
	this.ResultKeywordsExclude = resultKeywordsExclude
	this.RuleVersion = ruleVersion
	this.Severity = severity
	this.ShortDescription = shortDescription
	return &this
}

// NewAiPromptResponseAttributesWithDefaults instantiates a new AiPromptResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiPromptResponseAttributesWithDefaults() *AiPromptResponseAttributes {
	this := AiPromptResponseAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AiPromptResponseAttributes) GetCategory() CustomRuleRevisionAttributesCategory {
	if o == nil {
		var ret CustomRuleRevisionAttributesCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetCategoryOk() (*CustomRuleRevisionAttributesCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AiPromptResponseAttributes) SetCategory(v CustomRuleRevisionAttributesCategory) {
	o.Category = v
}

// GetChecksum returns the Checksum field value.
func (o *AiPromptResponseAttributes) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *AiPromptResponseAttributes) SetChecksum(v string) {
	o.Checksum = v
}

// GetContent returns the Content field value.
func (o *AiPromptResponseAttributes) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *AiPromptResponseAttributes) SetContent(v string) {
	o.Content = v
}

// GetCwe returns the Cwe field value if set, zero value otherwise.
func (o *AiPromptResponseAttributes) GetCwe() string {
	if o == nil || o.Cwe == nil {
		var ret string
		return ret
	}
	return *o.Cwe
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetCweOk() (*string, bool) {
	if o == nil || o.Cwe == nil {
		return nil, false
	}
	return o.Cwe, true
}

// HasCwe returns a boolean if a field has been set.
func (o *AiPromptResponseAttributes) HasCwe() bool {
	return o != nil && o.Cwe != nil
}

// SetCwe gets a reference to the given string and assigns it to the Cwe field.
func (o *AiPromptResponseAttributes) SetCwe(v string) {
	o.Cwe = &v
}

// GetDescription returns the Description field value.
func (o *AiPromptResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AiPromptResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDirectories returns the Directories field value.
func (o *AiPromptResponseAttributes) GetDirectories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Directories
}

// GetDirectoriesOk returns a tuple with the Directories field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetDirectoriesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Directories, true
}

// SetDirectories sets field value.
func (o *AiPromptResponseAttributes) SetDirectories(v []string) {
	o.Directories = v
}

// GetExecutionMode returns the ExecutionMode field value.
func (o *AiPromptResponseAttributes) GetExecutionMode() AiCustomRuleRevisionExecutionMode {
	if o == nil {
		var ret AiCustomRuleRevisionExecutionMode
		return ret
	}
	return o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetExecutionModeOk() (*AiCustomRuleRevisionExecutionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionMode, true
}

// SetExecutionMode sets field value.
func (o *AiPromptResponseAttributes) SetExecutionMode(v AiCustomRuleRevisionExecutionMode) {
	o.ExecutionMode = v
}

// GetFileSearchKeywords returns the FileSearchKeywords field value.
func (o *AiPromptResponseAttributes) GetFileSearchKeywords() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FileSearchKeywords
}

// GetFileSearchKeywordsOk returns a tuple with the FileSearchKeywords field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetFileSearchKeywordsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSearchKeywords, true
}

// SetFileSearchKeywords sets field value.
func (o *AiPromptResponseAttributes) SetFileSearchKeywords(v []string) {
	o.FileSearchKeywords = v
}

// GetGlobs returns the Globs field value.
func (o *AiPromptResponseAttributes) GetGlobs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Globs
}

// GetGlobsOk returns a tuple with the Globs field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetGlobsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Globs, true
}

// SetGlobs sets field value.
func (o *AiPromptResponseAttributes) SetGlobs(v []string) {
	o.Globs = v
}

// GetIsDefault returns the IsDefault field value.
func (o *AiPromptResponseAttributes) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value.
func (o *AiPromptResponseAttributes) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetIsTesting returns the IsTesting field value.
func (o *AiPromptResponseAttributes) GetIsTesting() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetIsTestingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTesting, true
}

// SetIsTesting sets field value.
func (o *AiPromptResponseAttributes) SetIsTesting(v bool) {
	o.IsTesting = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AiPromptResponseAttributes) GetLanguage() Language {
	if o == nil || o.Language == nil {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetLanguageOk() (*Language, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AiPromptResponseAttributes) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *AiPromptResponseAttributes) SetLanguage(v Language) {
	o.Language = &v
}

// GetResultKeywordsExclude returns the ResultKeywordsExclude field value.
func (o *AiPromptResponseAttributes) GetResultKeywordsExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResultKeywordsExclude
}

// GetResultKeywordsExcludeOk returns a tuple with the ResultKeywordsExclude field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetResultKeywordsExcludeOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultKeywordsExclude, true
}

// SetResultKeywordsExclude sets field value.
func (o *AiPromptResponseAttributes) SetResultKeywordsExclude(v []string) {
	o.ResultKeywordsExclude = v
}

// GetRuleVersion returns the RuleVersion field value.
func (o *AiPromptResponseAttributes) GetRuleVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleVersion
}

// GetRuleVersionOk returns a tuple with the RuleVersion field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetRuleVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleVersion, true
}

// SetRuleVersion sets field value.
func (o *AiPromptResponseAttributes) SetRuleVersion(v string) {
	o.RuleVersion = v
}

// GetSeverity returns the Severity field value.
func (o *AiPromptResponseAttributes) GetSeverity() CustomRuleRevisionAttributesSeverity {
	if o == nil {
		var ret CustomRuleRevisionAttributesSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetSeverityOk() (*CustomRuleRevisionAttributesSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AiPromptResponseAttributes) SetSeverity(v CustomRuleRevisionAttributesSeverity) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *AiPromptResponseAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *AiPromptResponseAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *AiPromptResponseAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiPromptResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["checksum"] = o.Checksum
	toSerialize["content"] = o.Content
	if o.Cwe != nil {
		toSerialize["cwe"] = o.Cwe
	}
	toSerialize["description"] = o.Description
	toSerialize["directories"] = o.Directories
	toSerialize["execution_mode"] = o.ExecutionMode
	toSerialize["file_search_keywords"] = o.FileSearchKeywords
	toSerialize["globs"] = o.Globs
	toSerialize["is_default"] = o.IsDefault
	toSerialize["is_testing"] = o.IsTesting
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	toSerialize["result_keywords_exclude"] = o.ResultKeywordsExclude
	toSerialize["rule_version"] = o.RuleVersion
	toSerialize["severity"] = o.Severity
	toSerialize["short_description"] = o.ShortDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiPromptResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category              *CustomRuleRevisionAttributesCategory `json:"category"`
		Checksum              *string                               `json:"checksum"`
		Content               *string                               `json:"content"`
		Cwe                   *string                               `json:"cwe,omitempty"`
		Description           *string                               `json:"description"`
		Directories           *[]string                             `json:"directories"`
		ExecutionMode         *AiCustomRuleRevisionExecutionMode    `json:"execution_mode"`
		FileSearchKeywords    *[]string                             `json:"file_search_keywords"`
		Globs                 *[]string                             `json:"globs"`
		IsDefault             *bool                                 `json:"is_default"`
		IsTesting             *bool                                 `json:"is_testing"`
		Language              *Language                             `json:"language,omitempty"`
		ResultKeywordsExclude *[]string                             `json:"result_keywords_exclude"`
		RuleVersion           *string                               `json:"rule_version"`
		Severity              *CustomRuleRevisionAttributesSeverity `json:"severity"`
		ShortDescription      *string                               `json:"short_description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Checksum == nil {
		return fmt.Errorf("required field checksum missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Directories == nil {
		return fmt.Errorf("required field directories missing")
	}
	if all.ExecutionMode == nil {
		return fmt.Errorf("required field execution_mode missing")
	}
	if all.FileSearchKeywords == nil {
		return fmt.Errorf("required field file_search_keywords missing")
	}
	if all.Globs == nil {
		return fmt.Errorf("required field globs missing")
	}
	if all.IsDefault == nil {
		return fmt.Errorf("required field is_default missing")
	}
	if all.IsTesting == nil {
		return fmt.Errorf("required field is_testing missing")
	}
	if all.ResultKeywordsExclude == nil {
		return fmt.Errorf("required field result_keywords_exclude missing")
	}
	if all.RuleVersion == nil {
		return fmt.Errorf("required field rule_version missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "checksum", "content", "cwe", "description", "directories", "execution_mode", "file_search_keywords", "globs", "is_default", "is_testing", "language", "result_keywords_exclude", "rule_version", "severity", "short_description"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Checksum = *all.Checksum
	o.Content = *all.Content
	o.Cwe = all.Cwe
	o.Description = *all.Description
	o.Directories = *all.Directories
	if !all.ExecutionMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ExecutionMode = *all.ExecutionMode
	}
	o.FileSearchKeywords = *all.FileSearchKeywords
	o.Globs = *all.Globs
	o.IsDefault = *all.IsDefault
	o.IsTesting = *all.IsTesting
	if all.Language != nil && !all.Language.IsValid() {
		hasInvalidField = true
	} else {
		o.Language = all.Language
	}
	o.ResultKeywordsExclude = *all.ResultKeywordsExclude
	o.RuleVersion = *all.RuleVersion
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.ShortDescription = *all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
