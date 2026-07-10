// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRuleRevisionResponseAttributes Response attributes of an AI custom rule revision.
type AiCustomRuleRevisionResponseAttributes struct {
	// Rule category
	Category CustomRuleRevisionAttributesCategory `json:"category"`
	// Checksum of the revision content.
	Checksum string `json:"checksum"`
	// Base64-encoded AI model content for this revision.
	Content string `json:"content"`
	// The creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The identifier of the user who created the revision.
	CreatedBy string `json:"created_by"`
	// The associated CWE identifier.
	Cwe datadog.NullableString `json:"cwe"`
	// Base64-encoded full description.
	Description string `json:"description"`
	// Directory patterns this rule applies to.
	Directories []string `json:"directories"`
	// The execution mode for an AI rule revision.
	ExecutionMode AiCustomRuleRevisionExecutionMode `json:"execution_mode"`
	// File glob patterns this rule applies to.
	Globs []string `json:"globs"`
	// Whether this is a default Datadog rule.
	IsDefault bool `json:"is_default"`
	// Whether this revision is published.
	IsPublished bool `json:"is_published"`
	// Whether this revision is for testing only.
	IsTesting bool `json:"is_testing"`
	// Rule severity
	Severity CustomRuleRevisionAttributesSeverity `json:"severity"`
	// Base64-encoded short description.
	ShortDescription string `json:"short_description"`
	// The version identifier for this revision.
	VersionId int64 `json:"version_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiCustomRuleRevisionResponseAttributes instantiates a new AiCustomRuleRevisionResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiCustomRuleRevisionResponseAttributes(category CustomRuleRevisionAttributesCategory, checksum string, content string, createdAt time.Time, createdBy string, cwe datadog.NullableString, description string, directories []string, executionMode AiCustomRuleRevisionExecutionMode, globs []string, isDefault bool, isPublished bool, isTesting bool, severity CustomRuleRevisionAttributesSeverity, shortDescription string, versionId int64) *AiCustomRuleRevisionResponseAttributes {
	this := AiCustomRuleRevisionResponseAttributes{}
	this.Category = category
	this.Checksum = checksum
	this.Content = content
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Cwe = cwe
	this.Description = description
	this.Directories = directories
	this.ExecutionMode = executionMode
	this.Globs = globs
	this.IsDefault = isDefault
	this.IsPublished = isPublished
	this.IsTesting = isTesting
	this.Severity = severity
	this.ShortDescription = shortDescription
	this.VersionId = versionId
	return &this
}

// NewAiCustomRuleRevisionResponseAttributesWithDefaults instantiates a new AiCustomRuleRevisionResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiCustomRuleRevisionResponseAttributesWithDefaults() *AiCustomRuleRevisionResponseAttributes {
	this := AiCustomRuleRevisionResponseAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetCategory() CustomRuleRevisionAttributesCategory {
	if o == nil {
		var ret CustomRuleRevisionAttributesCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetCategoryOk() (*CustomRuleRevisionAttributesCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetCategory(v CustomRuleRevisionAttributesCategory) {
	o.Category = v
}

// GetChecksum returns the Checksum field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetChecksum(v string) {
	o.Checksum = v
}

// GetContent returns the Content field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetContent(v string) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCwe returns the Cwe field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AiCustomRuleRevisionResponseAttributes) GetCwe() string {
	if o == nil || o.Cwe.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cwe.Get()
}

// GetCweOk returns a tuple with the Cwe field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiCustomRuleRevisionResponseAttributes) GetCweOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cwe.Get(), o.Cwe.IsSet()
}

// SetCwe sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetCwe(v string) {
	o.Cwe.Set(&v)
}

// GetDescription returns the Description field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDirectories returns the Directories field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetDirectories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Directories
}

// GetDirectoriesOk returns a tuple with the Directories field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetDirectoriesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Directories, true
}

// SetDirectories sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetDirectories(v []string) {
	o.Directories = v
}

// GetExecutionMode returns the ExecutionMode field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetExecutionMode() AiCustomRuleRevisionExecutionMode {
	if o == nil {
		var ret AiCustomRuleRevisionExecutionMode
		return ret
	}
	return o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetExecutionModeOk() (*AiCustomRuleRevisionExecutionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionMode, true
}

// SetExecutionMode sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetExecutionMode(v AiCustomRuleRevisionExecutionMode) {
	o.ExecutionMode = v
}

// GetGlobs returns the Globs field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetGlobs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Globs
}

// GetGlobsOk returns a tuple with the Globs field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetGlobsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Globs, true
}

// SetGlobs sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetGlobs(v []string) {
	o.Globs = v
}

// GetIsDefault returns the IsDefault field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetIsPublished returns the IsPublished field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetIsTesting returns the IsTesting field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetIsTesting() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetIsTestingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTesting, true
}

// SetIsTesting sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetIsTesting(v bool) {
	o.IsTesting = v
}

// GetSeverity returns the Severity field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetSeverity() CustomRuleRevisionAttributesSeverity {
	if o == nil {
		var ret CustomRuleRevisionAttributesSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetSeverityOk() (*CustomRuleRevisionAttributesSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetSeverity(v CustomRuleRevisionAttributesSeverity) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// GetVersionId returns the VersionId field value.
func (o *AiCustomRuleRevisionResponseAttributes) GetVersionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleRevisionResponseAttributes) GetVersionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value.
func (o *AiCustomRuleRevisionResponseAttributes) SetVersionId(v int64) {
	o.VersionId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiCustomRuleRevisionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["checksum"] = o.Checksum
	toSerialize["content"] = o.Content
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["cwe"] = o.Cwe.Get()
	toSerialize["description"] = o.Description
	toSerialize["directories"] = o.Directories
	toSerialize["execution_mode"] = o.ExecutionMode
	toSerialize["globs"] = o.Globs
	toSerialize["is_default"] = o.IsDefault
	toSerialize["is_published"] = o.IsPublished
	toSerialize["is_testing"] = o.IsTesting
	toSerialize["severity"] = o.Severity
	toSerialize["short_description"] = o.ShortDescription
	toSerialize["version_id"] = o.VersionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiCustomRuleRevisionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category         *CustomRuleRevisionAttributesCategory `json:"category"`
		Checksum         *string                               `json:"checksum"`
		Content          *string                               `json:"content"`
		CreatedAt        *time.Time                            `json:"created_at"`
		CreatedBy        *string                               `json:"created_by"`
		Cwe              datadog.NullableString                `json:"cwe"`
		Description      *string                               `json:"description"`
		Directories      *[]string                             `json:"directories"`
		ExecutionMode    *AiCustomRuleRevisionExecutionMode    `json:"execution_mode"`
		Globs            *[]string                             `json:"globs"`
		IsDefault        *bool                                 `json:"is_default"`
		IsPublished      *bool                                 `json:"is_published"`
		IsTesting        *bool                                 `json:"is_testing"`
		Severity         *CustomRuleRevisionAttributesSeverity `json:"severity"`
		ShortDescription *string                               `json:"short_description"`
		VersionId        *int64                                `json:"version_id"`
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
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if !all.Cwe.IsSet() {
		return fmt.Errorf("required field cwe missing")
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
	if all.IsDefault == nil {
		return fmt.Errorf("required field is_default missing")
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
	if all.VersionId == nil {
		return fmt.Errorf("required field version_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "checksum", "content", "created_at", "created_by", "cwe", "description", "directories", "execution_mode", "globs", "is_default", "is_published", "is_testing", "severity", "short_description", "version_id"})
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
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Cwe = all.Cwe
	o.Description = *all.Description
	o.Directories = *all.Directories
	if !all.ExecutionMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ExecutionMode = *all.ExecutionMode
	}
	o.Globs = *all.Globs
	o.IsDefault = *all.IsDefault
	o.IsPublished = *all.IsPublished
	o.IsTesting = *all.IsTesting
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.ShortDescription = *all.ShortDescription
	o.VersionId = *all.VersionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
