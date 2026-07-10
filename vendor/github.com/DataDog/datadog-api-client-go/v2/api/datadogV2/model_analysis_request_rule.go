// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisRequestRule A static analysis rule to apply during code analysis.
type AnalysisRequestRule struct {
	// The category of the rule (for example, `BEST_PRACTICES`, `SECURITY`).
	Category string `json:"category"`
	// A checksum of the rule definition.
	Checksum string `json:"checksum"`
	// The base64-encoded rule implementation code.
	Code string `json:"code"`
	// The code entity type checked by the rule, applicable when rule type is `AST_CHECK`.
	EntityChecked datadog.NullableString `json:"entity_checked,omitempty"`
	// The unique identifier of the rule.
	Id string `json:"id"`
	// The programming language this rule targets.
	Language string `json:"language"`
	// A base64-encoded regex pattern used by the rule, applicable when rule type is `REGEX`.
	Regex datadog.NullableString `json:"regex,omitempty"`
	// The severity of findings from this rule (for example, `ERROR`, `WARNING`).
	Severity string `json:"severity"`
	// The base64-encoded tree-sitter query used by the rule.
	TreeSitterQuery string `json:"tree_sitter_query"`
	// The rule type indicating the detection mechanism (for example, `TREE_SITTER_QUERY`).
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisRequestRule instantiates a new AnalysisRequestRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisRequestRule(category string, checksum string, code string, id string, language string, severity string, treeSitterQuery string, typeVar string) *AnalysisRequestRule {
	this := AnalysisRequestRule{}
	this.Category = category
	this.Checksum = checksum
	this.Code = code
	this.Id = id
	this.Language = language
	this.Severity = severity
	this.TreeSitterQuery = treeSitterQuery
	this.Type = typeVar
	return &this
}

// NewAnalysisRequestRuleWithDefaults instantiates a new AnalysisRequestRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisRequestRuleWithDefaults() *AnalysisRequestRule {
	this := AnalysisRequestRule{}
	return &this
}

// GetCategory returns the Category field value.
func (o *AnalysisRequestRule) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AnalysisRequestRule) SetCategory(v string) {
	o.Category = v
}

// GetChecksum returns the Checksum field value.
func (o *AnalysisRequestRule) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *AnalysisRequestRule) SetChecksum(v string) {
	o.Checksum = v
}

// GetCode returns the Code field value.
func (o *AnalysisRequestRule) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *AnalysisRequestRule) SetCode(v string) {
	o.Code = v
}

// GetEntityChecked returns the EntityChecked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisRequestRule) GetEntityChecked() string {
	if o == nil || o.EntityChecked.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityChecked.Get()
}

// GetEntityCheckedOk returns a tuple with the EntityChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnalysisRequestRule) GetEntityCheckedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityChecked.Get(), o.EntityChecked.IsSet()
}

// HasEntityChecked returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasEntityChecked() bool {
	return o != nil && o.EntityChecked.IsSet()
}

// SetEntityChecked gets a reference to the given datadog.NullableString and assigns it to the EntityChecked field.
func (o *AnalysisRequestRule) SetEntityChecked(v string) {
	o.EntityChecked.Set(&v)
}

// SetEntityCheckedNil sets the value for EntityChecked to be an explicit nil.
func (o *AnalysisRequestRule) SetEntityCheckedNil() {
	o.EntityChecked.Set(nil)
}

// UnsetEntityChecked ensures that no value is present for EntityChecked, not even an explicit nil.
func (o *AnalysisRequestRule) UnsetEntityChecked() {
	o.EntityChecked.Unset()
}

// GetId returns the Id field value.
func (o *AnalysisRequestRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AnalysisRequestRule) SetId(v string) {
	o.Id = v
}

// GetLanguage returns the Language field value.
func (o *AnalysisRequestRule) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *AnalysisRequestRule) SetLanguage(v string) {
	o.Language = v
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisRequestRule) GetRegex() string {
	if o == nil || o.Regex.Get() == nil {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnalysisRequestRule) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasRegex() bool {
	return o != nil && o.Regex.IsSet()
}

// SetRegex gets a reference to the given datadog.NullableString and assigns it to the Regex field.
func (o *AnalysisRequestRule) SetRegex(v string) {
	o.Regex.Set(&v)
}

// SetRegexNil sets the value for Regex to be an explicit nil.
func (o *AnalysisRequestRule) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil.
func (o *AnalysisRequestRule) UnsetRegex() {
	o.Regex.Unset()
}

// GetSeverity returns the Severity field value.
func (o *AnalysisRequestRule) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AnalysisRequestRule) SetSeverity(v string) {
	o.Severity = v
}

// GetTreeSitterQuery returns the TreeSitterQuery field value.
func (o *AnalysisRequestRule) GetTreeSitterQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TreeSitterQuery, true
}

// SetTreeSitterQuery sets field value.
func (o *AnalysisRequestRule) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = v
}

// GetType returns the Type field value.
func (o *AnalysisRequestRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AnalysisRequestRule) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisRequestRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["checksum"] = o.Checksum
	toSerialize["code"] = o.Code
	if o.EntityChecked.IsSet() {
		toSerialize["entity_checked"] = o.EntityChecked.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["language"] = o.Language
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	toSerialize["severity"] = o.Severity
	toSerialize["tree_sitter_query"] = o.TreeSitterQuery
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisRequestRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category        *string                `json:"category"`
		Checksum        *string                `json:"checksum"`
		Code            *string                `json:"code"`
		EntityChecked   datadog.NullableString `json:"entity_checked,omitempty"`
		Id              *string                `json:"id"`
		Language        *string                `json:"language"`
		Regex           datadog.NullableString `json:"regex,omitempty"`
		Severity        *string                `json:"severity"`
		TreeSitterQuery *string                `json:"tree_sitter_query"`
		Type            *string                `json:"type"`
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
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Language == nil {
		return fmt.Errorf("required field language missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.TreeSitterQuery == nil {
		return fmt.Errorf("required field tree_sitter_query missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "checksum", "code", "entity_checked", "id", "language", "regex", "severity", "tree_sitter_query", "type"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.Checksum = *all.Checksum
	o.Code = *all.Code
	o.EntityChecked = all.EntityChecked
	o.Id = *all.Id
	o.Language = *all.Language
	o.Regex = all.Regex
	o.Severity = *all.Severity
	o.TreeSitterQuery = *all.TreeSitterQuery
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
