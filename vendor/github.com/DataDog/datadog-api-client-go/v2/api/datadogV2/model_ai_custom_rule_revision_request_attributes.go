// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRuleRevisionRequestAttributes Attributes for creating an AI custom rule revision.
type AiCustomRuleRevisionRequestAttributes struct {
	// Rule category
	Category CustomRuleRevisionAttributesCategory `json:"category"`
	// Base64-encoded AI model content for this revision.
	Content string `json:"content"`
	// The associated CWE identifier.
	Cwe datadog.NullableString `json:"cwe,omitempty"`
	// Base64-encoded full description.
	Description string `json:"description"`
	// Directory patterns this rule applies to.
	Directories []string `json:"directories"`
	// The execution mode for an AI rule revision.
	ExecutionMode AiCustomRuleRevisionExecutionMode `json:"execution_mode"`
	// File glob patterns this rule applies to.
	Globs []string `json:"globs"`
	// Whether this revision is published.
	IsPublished bool `json:"is_published"`
	// Whether this revision is for testing only.
	IsTesting bool `json:"is_testing"`
	// Rule severity
	Severity CustomRuleRevisionAttributesSeverity `json:"severity"`
	// Base64-encoded short description.
	ShortDescription string `json:"short_description"`
	// The version identifier for this revision.
	VersionId *int64 `json:"version_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiCustomRuleRevisionRequestAttributes instantiates a new AiCustomRuleRevisionRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiCustomRuleRevisionRequestAttributes(category CustomRuleRevisionAttributesCategory, content string, description string, directories []string, executionMode AiCustomRuleRevisionExecutionMode, globs []string, isPublished bool, isTesting bool, severity CustomRuleRevisionAttributesSeverity, shortDescription string) *AiCustomRuleRevisionRequestAttributes {
	this := AiCustomRuleRevisionRequestAttributes{}
	this.Category = category
	this.Content = content
	this.Description = description
	this.Directories = directories
	this.ExecutionMode = executionMode
	this.Globs = globs
	this.IsPublished = isPublished
	this.IsTesting = isTesting
	this.Severity = severity
	this.ShortDescription = shortDescription
	return &this
}

// NewAiCustomRuleRevisionRequestAttributesWithDefaults instantiates a new AiCustomRuleRevisionRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiCustomRuleRevisionRequestAttributesWithDefaults() *AiCustomRuleRevisionRequestAttributes {
	this := AiCustomRuleRevisionRequestAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetCategory() CustomRuleRevisionAttributesCategory {
	if o == nil {
		var ret CustomRuleRevisionAttributesCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetCategoryOk() (*CustomRuleRevisionAttributesCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetCategory(v CustomRuleRevisionAttributesCategory) {
	o.Category = v
}

// GetContent returns the Content field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetContent(v string) {
	o.Content = v
}

// GetCwe returns the Cwe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiCustomRuleRevisionRequestAttributes) GetCwe() string {
	if o == nil || o.Cwe.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cwe.Get()
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiCustomRuleRevisionRequestAttributes) GetCweOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cwe.Get(), o.Cwe.IsSet()
}

// HasCwe returns a boolean if a field has been set.
func (o *AiCustomRuleRevisionRequestAttributes) HasCwe() bool {
	return o != nil && o.Cwe.IsSet()
}

// SetCwe gets a reference to the given datadog.NullableString and assigns it to the Cwe field.
func (o *AiCustomRuleRevisionRequestAttributes) SetCwe(v string) {
	o.Cwe.Set(&v)
}

// SetCweNil sets the value for Cwe to be an explicit nil.
func (o *AiCustomRuleRevisionRequestAttributes) SetCweNil() {
	o.Cwe.Set(nil)
}

// UnsetCwe ensures that no value is present for Cwe, not even an explicit nil.
func (o *AiCustomRuleRevisionRequestAttributes) UnsetCwe() {
	o.Cwe.Unset()
}

// GetDescription returns the Description field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDirectories returns the Directories field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetDirectories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Directories
}

// GetDirectoriesOk returns a tuple with the Directories field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetDirectoriesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Directories, true
}

// SetDirectories sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetDirectories(v []string) {
	o.Directories = v
}

// GetExecutionMode returns the ExecutionMode field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetExecutionMode() AiCustomRuleRevisionExecutionMode {
	if o == nil {
		var ret AiCustomRuleRevisionExecutionMode
		return ret
	}
	return o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetExecutionModeOk() (*AiCustomRuleRevisionExecutionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionMode, true
}

// SetExecutionMode sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetExecutionMode(v AiCustomRuleRevisionExecutionMode) {
	o.ExecutionMode = v
}

// GetGlobs returns the Globs field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetGlobs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Globs
}

// GetGlobsOk returns a tuple with the Globs field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetGlobsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Globs, true
}

// SetGlobs sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetGlobs(v []string) {
	o.Globs = v
}

// GetIsPublished returns the IsPublished field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetIsTesting returns the IsTesting field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetIsTesting() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetIsTestingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTesting, true
}

// SetIsTesting sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetIsTesting(v bool) {
	o.IsTesting = v
}

// GetSeverity returns the Severity field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetSeverity() CustomRuleRevisionAttributesSeverity {
	if o == nil {
		var ret CustomRuleRevisionAttributesSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetSeverityOk() (*CustomRuleRevisionAttributesSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetSeverity(v CustomRuleRevisionAttributesSeverity) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *AiCustomRuleRevisionRequestAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *AiCustomRuleRevisionRequestAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *AiCustomRuleRevisionRequestAttributes) GetVersionId() int64 {
	if o == nil || o.VersionId == nil {
		var ret int64
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionRequestAttributes) GetVersionIdOk() (*int64, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *AiCustomRuleRevisionRequestAttributes) HasVersionId() bool {
	return o != nil && o.VersionId != nil
}

// SetVersionId gets a reference to the given int64 and assigns it to the VersionId field.
func (o *AiCustomRuleRevisionRequestAttributes) SetVersionId(v int64) {
	o.VersionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiCustomRuleRevisionRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["content"] = o.Content
	if o.Cwe.IsSet() {
		toSerialize["cwe"] = o.Cwe.Get()
	}
	toSerialize["description"] = o.Description
	toSerialize["directories"] = o.Directories
	toSerialize["execution_mode"] = o.ExecutionMode
	toSerialize["globs"] = o.Globs
	toSerialize["is_published"] = o.IsPublished
	toSerialize["is_testing"] = o.IsTesting
	toSerialize["severity"] = o.Severity
	toSerialize["short_description"] = o.ShortDescription
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiCustomRuleRevisionRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category         *CustomRuleRevisionAttributesCategory `json:"category"`
		Content          *string                               `json:"content"`
		Cwe              datadog.NullableString                `json:"cwe,omitempty"`
		Description      *string                               `json:"description"`
		Directories      *[]string                             `json:"directories"`
		ExecutionMode    *AiCustomRuleRevisionExecutionMode    `json:"execution_mode"`
		Globs            *[]string                             `json:"globs"`
		IsPublished      *bool                                 `json:"is_published"`
		IsTesting        *bool                                 `json:"is_testing"`
		Severity         *CustomRuleRevisionAttributesSeverity `json:"severity"`
		ShortDescription *string                               `json:"short_description"`
		VersionId        *int64                                `json:"version_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
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
	if all.Globs == nil {
		return fmt.Errorf("required field globs missing")
	}
	if all.IsPublished == nil {
		return fmt.Errorf("required field is_published missing")
	}
	if all.IsTesting == nil {
		return fmt.Errorf("required field is_testing missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "content", "cwe", "description", "directories", "execution_mode", "globs", "is_published", "is_testing", "severity", "short_description", "version_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Content = *all.Content
	o.Cwe = all.Cwe
	o.Description = *all.Description
	o.Directories = *all.Directories
	if !all.ExecutionMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ExecutionMode = *all.ExecutionMode
	}
	o.Globs = *all.Globs
	o.IsPublished = *all.IsPublished
	o.IsTesting = *all.IsTesting
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.ShortDescription = *all.ShortDescription
	o.VersionId = all.VersionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
