// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevisionInputAttributes Input attributes for creating or updating a custom rule revision.
type CustomRuleRevisionInputAttributes struct {
	// Rule arguments
	Arguments []Argument `json:"arguments"`
	// Rule category
	Category CustomRuleRevisionAttributesCategory `json:"category"`
	// Rule code
	Code string `json:"code"`
	// Revision creation message
	CreationMessage string `json:"creation_message"`
	// Associated CVE
	Cve datadog.NullableString `json:"cve"`
	// Associated CWE
	Cwe datadog.NullableString `json:"cwe"`
	// Full description
	Description string `json:"description"`
	// Documentation URL
	DocumentationUrl datadog.NullableString `json:"documentation_url"`
	// Whether the revision is published
	IsPublished bool `json:"is_published"`
	// Whether this is a testing revision
	IsTesting bool `json:"is_testing"`
	// Programming language
	Language Language `json:"language"`
	// Rule severity
	Severity CustomRuleRevisionAttributesSeverity `json:"severity"`
	// Short description
	ShortDescription string `json:"short_description"`
	// Whether to use AI for fixes
	ShouldUseAiFix bool `json:"should_use_ai_fix"`
	// Rule tags
	Tags []string `json:"tags"`
	// Rule tests
	Tests []CustomRuleRevisionTest `json:"tests"`
	// Tree-sitter query
	TreeSitterQuery string `json:"tree_sitter_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRuleRevisionInputAttributes instantiates a new CustomRuleRevisionInputAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRuleRevisionInputAttributes(arguments []Argument, category CustomRuleRevisionAttributesCategory, code string, creationMessage string, cve datadog.NullableString, cwe datadog.NullableString, description string, documentationUrl datadog.NullableString, isPublished bool, isTesting bool, language Language, severity CustomRuleRevisionAttributesSeverity, shortDescription string, shouldUseAiFix bool, tags []string, tests []CustomRuleRevisionTest, treeSitterQuery string) *CustomRuleRevisionInputAttributes {
	this := CustomRuleRevisionInputAttributes{}
	this.Arguments = arguments
	this.Category = category
	this.Code = code
	this.CreationMessage = creationMessage
	this.Cve = cve
	this.Cwe = cwe
	this.Description = description
	this.DocumentationUrl = documentationUrl
	this.IsPublished = isPublished
	this.IsTesting = isTesting
	this.Language = language
	this.Severity = severity
	this.ShortDescription = shortDescription
	this.ShouldUseAiFix = shouldUseAiFix
	this.Tags = tags
	this.Tests = tests
	this.TreeSitterQuery = treeSitterQuery
	return &this
}

// NewCustomRuleRevisionInputAttributesWithDefaults instantiates a new CustomRuleRevisionInputAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleRevisionInputAttributesWithDefaults() *CustomRuleRevisionInputAttributes {
	this := CustomRuleRevisionInputAttributes{}
	return &this
}

// GetArguments returns the Arguments field value.
func (o *CustomRuleRevisionInputAttributes) GetArguments() []Argument {
	if o == nil {
		var ret []Argument
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetArgumentsOk() (*[]Argument, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value.
func (o *CustomRuleRevisionInputAttributes) SetArguments(v []Argument) {
	o.Arguments = v
}

// GetCategory returns the Category field value.
func (o *CustomRuleRevisionInputAttributes) GetCategory() CustomRuleRevisionAttributesCategory {
	if o == nil {
		var ret CustomRuleRevisionAttributesCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetCategoryOk() (*CustomRuleRevisionAttributesCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *CustomRuleRevisionInputAttributes) SetCategory(v CustomRuleRevisionAttributesCategory) {
	o.Category = v
}

// GetCode returns the Code field value.
func (o *CustomRuleRevisionInputAttributes) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *CustomRuleRevisionInputAttributes) SetCode(v string) {
	o.Code = v
}

// GetCreationMessage returns the CreationMessage field value.
func (o *CustomRuleRevisionInputAttributes) GetCreationMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreationMessage
}

// GetCreationMessageOk returns a tuple with the CreationMessage field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetCreationMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationMessage, true
}

// SetCreationMessage sets field value.
func (o *CustomRuleRevisionInputAttributes) SetCreationMessage(v string) {
	o.CreationMessage = v
}

// GetCve returns the Cve field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CustomRuleRevisionInputAttributes) GetCve() string {
	if o == nil || o.Cve.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cve.Get()
}

// GetCveOk returns a tuple with the Cve field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInputAttributes) GetCveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cve.Get(), o.Cve.IsSet()
}

// SetCve sets field value.
func (o *CustomRuleRevisionInputAttributes) SetCve(v string) {
	o.Cve.Set(&v)
}

// GetCwe returns the Cwe field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CustomRuleRevisionInputAttributes) GetCwe() string {
	if o == nil || o.Cwe.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cwe.Get()
}

// GetCweOk returns a tuple with the Cwe field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInputAttributes) GetCweOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cwe.Get(), o.Cwe.IsSet()
}

// SetCwe sets field value.
func (o *CustomRuleRevisionInputAttributes) SetCwe(v string) {
	o.Cwe.Set(&v)
}

// GetDescription returns the Description field value.
func (o *CustomRuleRevisionInputAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CustomRuleRevisionInputAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDocumentationUrl returns the DocumentationUrl field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *CustomRuleRevisionInputAttributes) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl.Get()
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInputAttributes) GetDocumentationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentationUrl.Get(), o.DocumentationUrl.IsSet()
}

// SetDocumentationUrl sets field value.
func (o *CustomRuleRevisionInputAttributes) SetDocumentationUrl(v string) {
	o.DocumentationUrl.Set(&v)
}

// GetIsPublished returns the IsPublished field value.
func (o *CustomRuleRevisionInputAttributes) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value.
func (o *CustomRuleRevisionInputAttributes) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetIsTesting returns the IsTesting field value.
func (o *CustomRuleRevisionInputAttributes) GetIsTesting() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetIsTestingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTesting, true
}

// SetIsTesting sets field value.
func (o *CustomRuleRevisionInputAttributes) SetIsTesting(v bool) {
	o.IsTesting = v
}

// GetLanguage returns the Language field value.
func (o *CustomRuleRevisionInputAttributes) GetLanguage() Language {
	if o == nil {
		var ret Language
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetLanguageOk() (*Language, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *CustomRuleRevisionInputAttributes) SetLanguage(v Language) {
	o.Language = v
}

// GetSeverity returns the Severity field value.
func (o *CustomRuleRevisionInputAttributes) GetSeverity() CustomRuleRevisionAttributesSeverity {
	if o == nil {
		var ret CustomRuleRevisionAttributesSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetSeverityOk() (*CustomRuleRevisionAttributesSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *CustomRuleRevisionInputAttributes) SetSeverity(v CustomRuleRevisionAttributesSeverity) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *CustomRuleRevisionInputAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *CustomRuleRevisionInputAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// GetShouldUseAiFix returns the ShouldUseAiFix field value.
func (o *CustomRuleRevisionInputAttributes) GetShouldUseAiFix() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ShouldUseAiFix
}

// GetShouldUseAiFixOk returns a tuple with the ShouldUseAiFix field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetShouldUseAiFixOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldUseAiFix, true
}

// SetShouldUseAiFix sets field value.
func (o *CustomRuleRevisionInputAttributes) SetShouldUseAiFix(v bool) {
	o.ShouldUseAiFix = v
}

// GetTags returns the Tags field value.
func (o *CustomRuleRevisionInputAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *CustomRuleRevisionInputAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTests returns the Tests field value.
func (o *CustomRuleRevisionInputAttributes) GetTests() []CustomRuleRevisionTest {
	if o == nil {
		var ret []CustomRuleRevisionTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetTestsOk() (*[]CustomRuleRevisionTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tests, true
}

// SetTests sets field value.
func (o *CustomRuleRevisionInputAttributes) SetTests(v []CustomRuleRevisionTest) {
	o.Tests = v
}

// GetTreeSitterQuery returns the TreeSitterQuery field value.
func (o *CustomRuleRevisionInputAttributes) GetTreeSitterQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInputAttributes) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TreeSitterQuery, true
}

// SetTreeSitterQuery sets field value.
func (o *CustomRuleRevisionInputAttributes) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRuleRevisionInputAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arguments"] = o.Arguments
	toSerialize["category"] = o.Category
	toSerialize["code"] = o.Code
	toSerialize["creation_message"] = o.CreationMessage
	toSerialize["cve"] = o.Cve.Get()
	toSerialize["cwe"] = o.Cwe.Get()
	toSerialize["description"] = o.Description
	toSerialize["documentation_url"] = o.DocumentationUrl.Get()
	toSerialize["is_published"] = o.IsPublished
	toSerialize["is_testing"] = o.IsTesting
	toSerialize["language"] = o.Language
	toSerialize["severity"] = o.Severity
	toSerialize["short_description"] = o.ShortDescription
	toSerialize["should_use_ai_fix"] = o.ShouldUseAiFix
	toSerialize["tags"] = o.Tags
	toSerialize["tests"] = o.Tests
	toSerialize["tree_sitter_query"] = o.TreeSitterQuery

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRuleRevisionInputAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments        *[]Argument                           `json:"arguments"`
		Category         *CustomRuleRevisionAttributesCategory `json:"category"`
		Code             *string                               `json:"code"`
		CreationMessage  *string                               `json:"creation_message"`
		Cve              datadog.NullableString                `json:"cve"`
		Cwe              datadog.NullableString                `json:"cwe"`
		Description      *string                               `json:"description"`
		DocumentationUrl datadog.NullableString                `json:"documentation_url"`
		IsPublished      *bool                                 `json:"is_published"`
		IsTesting        *bool                                 `json:"is_testing"`
		Language         *Language                             `json:"language"`
		Severity         *CustomRuleRevisionAttributesSeverity `json:"severity"`
		ShortDescription *string                               `json:"short_description"`
		ShouldUseAiFix   *bool                                 `json:"should_use_ai_fix"`
		Tags             *[]string                             `json:"tags"`
		Tests            *[]CustomRuleRevisionTest             `json:"tests"`
		TreeSitterQuery  *string                               `json:"tree_sitter_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Arguments == nil {
		return fmt.Errorf("required field arguments missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.CreationMessage == nil {
		return fmt.Errorf("required field creation_message missing")
	}
	if !all.Cve.IsSet() {
		return fmt.Errorf("required field cve missing")
	}
	if !all.Cwe.IsSet() {
		return fmt.Errorf("required field cwe missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if !all.DocumentationUrl.IsSet() {
		return fmt.Errorf("required field documentation_url missing")
	}
	if all.IsPublished == nil {
		return fmt.Errorf("required field is_published missing")
	}
	if all.IsTesting == nil {
		return fmt.Errorf("required field is_testing missing")
	}
	if all.Language == nil {
		return fmt.Errorf("required field language missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	if all.ShouldUseAiFix == nil {
		return fmt.Errorf("required field should_use_ai_fix missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Tests == nil {
		return fmt.Errorf("required field tests missing")
	}
	if all.TreeSitterQuery == nil {
		return fmt.Errorf("required field tree_sitter_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "category", "code", "creation_message", "cve", "cwe", "description", "documentation_url", "is_published", "is_testing", "language", "severity", "short_description", "should_use_ai_fix", "tags", "tests", "tree_sitter_query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Arguments = *all.Arguments
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Code = *all.Code
	o.CreationMessage = *all.CreationMessage
	o.Cve = all.Cve
	o.Cwe = all.Cwe
	o.Description = *all.Description
	o.DocumentationUrl = all.DocumentationUrl
	o.IsPublished = *all.IsPublished
	o.IsTesting = *all.IsTesting
	if !all.Language.IsValid() {
		hasInvalidField = true
	} else {
		o.Language = *all.Language
	}
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.ShortDescription = *all.ShortDescription
	o.ShouldUseAiFix = *all.ShouldUseAiFix
	o.Tags = *all.Tags
	o.Tests = *all.Tests
	o.TreeSitterQuery = *all.TreeSitterQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
