// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems
type GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems struct {
	//
	Arguments []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems `json:"arguments,omitempty"`
	//
	Category *string `json:"category,omitempty"`
	//
	Checksum *string `json:"checksum,omitempty"`
	//
	Code *string `json:"code,omitempty"`
	//
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//
	CreatedBy *string `json:"created_by,omitempty"`
	//
	Cve *string `json:"cve,omitempty"`
	//
	Cwe *string `json:"cwe,omitempty"`
	//
	Data GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData `json:"data"`
	//
	Description *string `json:"description,omitempty"`
	//
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	//
	EntityChecked *string `json:"entity_checked,omitempty"`
	//
	IsPublished *bool `json:"is_published,omitempty"`
	//
	IsTesting *bool `json:"is_testing,omitempty"`
	//
	Language *string `json:"language,omitempty"`
	//
	LastUpdatedAt *time.Time `json:"last_updated_at,omitempty"`
	//
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Regex *string `json:"regex,omitempty"`
	//
	Severity *string `json:"severity,omitempty"`
	//
	ShortDescription *string `json:"short_description,omitempty"`
	//
	ShouldUseAiFix *bool `json:"should_use_ai_fix,omitempty"`
	//
	Tests []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems `json:"tests,omitempty"`
	//
	TreeSitterQuery *string `json:"tree_sitter_query,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems(data GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData) *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems{}
	this.Data = data
	return &this
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsWithDefaults instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsWithDefaults() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetArguments() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems {
	if o == nil || o.Arguments == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetArgumentsOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasArguments() bool {
	return o != nil && o.Arguments != nil
}

// SetArguments gets a reference to the given []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems and assigns it to the Arguments field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetArguments(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) {
	o.Arguments = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCategory(v string) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasChecksum() bool {
	return o != nil && o.Checksum != nil
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetChecksum(v string) {
	o.Checksum = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCode(v string) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCve() string {
	if o == nil || o.Cve == nil {
		var ret string
		return ret
	}
	return *o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCveOk() (*string, bool) {
	if o == nil || o.Cve == nil {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCve() bool {
	return o != nil && o.Cve != nil
}

// SetCve gets a reference to the given string and assigns it to the Cve field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCve(v string) {
	o.Cve = &v
}

// GetCwe returns the Cwe field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCwe() string {
	if o == nil || o.Cwe == nil {
		var ret string
		return ret
	}
	return *o.Cwe
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCweOk() (*string, bool) {
	if o == nil || o.Cwe == nil {
		return nil, false
	}
	return o.Cwe, true
}

// HasCwe returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCwe() bool {
	return o != nil && o.Cwe != nil
}

// SetCwe gets a reference to the given string and assigns it to the Cwe field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCwe(v string) {
	o.Cwe = &v
}

// GetData returns the Data field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetData() GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData {
	if o == nil {
		var ret GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDataOk() (*GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetData(v GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData) {
	o.Data = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasDocumentationUrl() bool {
	return o != nil && o.DocumentationUrl != nil
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetEntityChecked returns the EntityChecked field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetEntityChecked() string {
	if o == nil || o.EntityChecked == nil {
		var ret string
		return ret
	}
	return *o.EntityChecked
}

// GetEntityCheckedOk returns a tuple with the EntityChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetEntityCheckedOk() (*string, bool) {
	if o == nil || o.EntityChecked == nil {
		return nil, false
	}
	return o.EntityChecked, true
}

// HasEntityChecked returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasEntityChecked() bool {
	return o != nil && o.EntityChecked != nil
}

// SetEntityChecked gets a reference to the given string and assigns it to the EntityChecked field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetEntityChecked(v string) {
	o.EntityChecked = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsPublished() bool {
	if o == nil || o.IsPublished == nil {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsPublishedOk() (*bool, bool) {
	if o == nil || o.IsPublished == nil {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasIsPublished() bool {
	return o != nil && o.IsPublished != nil
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetIsTesting returns the IsTesting field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsTesting() bool {
	if o == nil || o.IsTesting == nil {
		var ret bool
		return ret
	}
	return *o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsTestingOk() (*bool, bool) {
	if o == nil || o.IsTesting == nil {
		return nil, false
	}
	return o.IsTesting, true
}

// HasIsTesting returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasIsTesting() bool {
	return o != nil && o.IsTesting != nil
}

// SetIsTesting gets a reference to the given bool and assigns it to the IsTesting field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetIsTesting(v bool) {
	o.IsTesting = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetLanguage(v string) {
	o.Language = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedAt() time.Time {
	if o == nil || o.LastUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedAt == nil {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasLastUpdatedAt() bool {
	return o != nil && o.LastUpdatedAt != nil
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetName(v string) {
	o.Name = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasRegex() bool {
	return o != nil && o.Regex != nil
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetRegex(v string) {
	o.Regex = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetSeverity(v string) {
	o.Severity = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasShortDescription() bool {
	return o != nil && o.ShortDescription != nil
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetShouldUseAiFix returns the ShouldUseAiFix field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShouldUseAiFix() bool {
	if o == nil || o.ShouldUseAiFix == nil {
		var ret bool
		return ret
	}
	return *o.ShouldUseAiFix
}

// GetShouldUseAiFixOk returns a tuple with the ShouldUseAiFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShouldUseAiFixOk() (*bool, bool) {
	if o == nil || o.ShouldUseAiFix == nil {
		return nil, false
	}
	return o.ShouldUseAiFix, true
}

// HasShouldUseAiFix returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasShouldUseAiFix() bool {
	return o != nil && o.ShouldUseAiFix != nil
}

// SetShouldUseAiFix gets a reference to the given bool and assigns it to the ShouldUseAiFix field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetShouldUseAiFix(v bool) {
	o.ShouldUseAiFix = &v
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTests() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems {
	if o == nil || o.Tests == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTestsOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems, bool) {
	if o == nil || o.Tests == nil {
		return nil, false
	}
	return &o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasTests() bool {
	return o != nil && o.Tests != nil
}

// SetTests gets a reference to the given []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems and assigns it to the Tests field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetTests(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) {
	o.Tests = v
}

// GetTreeSitterQuery returns the TreeSitterQuery field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTreeSitterQuery() string {
	if o == nil || o.TreeSitterQuery == nil {
		var ret string
		return ret
	}
	return *o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil || o.TreeSitterQuery == nil {
		return nil, false
	}
	return o.TreeSitterQuery, true
}

// HasTreeSitterQuery returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasTreeSitterQuery() bool {
	return o != nil && o.TreeSitterQuery != nil
}

// SetTreeSitterQuery gets a reference to the given string and assigns it to the TreeSitterQuery field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Cve != nil {
		toSerialize["cve"] = o.Cve
	}
	if o.Cwe != nil {
		toSerialize["cwe"] = o.Cwe
	}
	toSerialize["data"] = o.Data
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentationUrl != nil {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if o.EntityChecked != nil {
		toSerialize["entity_checked"] = o.EntityChecked
	}
	if o.IsPublished != nil {
		toSerialize["is_published"] = o.IsPublished
	}
	if o.IsTesting != nil {
		toSerialize["is_testing"] = o.IsTesting
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.LastUpdatedAt != nil {
		if o.LastUpdatedAt.Nanosecond() == 0 {
			toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}
	if o.ShouldUseAiFix != nil {
		toSerialize["should_use_ai_fix"] = o.ShouldUseAiFix
	}
	if o.Tests != nil {
		toSerialize["tests"] = o.Tests
	}
	if o.TreeSitterQuery != nil {
		toSerialize["tree_sitter_query"] = o.TreeSitterQuery
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments        []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems `json:"arguments,omitempty"`
		Category         *string                                                                          `json:"category,omitempty"`
		Checksum         *string                                                                          `json:"checksum,omitempty"`
		Code             *string                                                                          `json:"code,omitempty"`
		CreatedAt        *time.Time                                                                       `json:"created_at,omitempty"`
		CreatedBy        *string                                                                          `json:"created_by,omitempty"`
		Cve              *string                                                                          `json:"cve,omitempty"`
		Cwe              *string                                                                          `json:"cwe,omitempty"`
		Data             *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsData            `json:"data"`
		Description      *string                                                                          `json:"description,omitempty"`
		DocumentationUrl *string                                                                          `json:"documentation_url,omitempty"`
		EntityChecked    *string                                                                          `json:"entity_checked,omitempty"`
		IsPublished      *bool                                                                            `json:"is_published,omitempty"`
		IsTesting        *bool                                                                            `json:"is_testing,omitempty"`
		Language         *string                                                                          `json:"language,omitempty"`
		LastUpdatedAt    *time.Time                                                                       `json:"last_updated_at,omitempty"`
		LastUpdatedBy    *string                                                                          `json:"last_updated_by,omitempty"`
		Name             *string                                                                          `json:"name,omitempty"`
		Regex            *string                                                                          `json:"regex,omitempty"`
		Severity         *string                                                                          `json:"severity,omitempty"`
		ShortDescription *string                                                                          `json:"short_description,omitempty"`
		ShouldUseAiFix   *bool                                                                            `json:"should_use_ai_fix,omitempty"`
		Tests            []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems     `json:"tests,omitempty"`
		TreeSitterQuery  *string                                                                          `json:"tree_sitter_query,omitempty"`
		Type             *string                                                                          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "category", "checksum", "code", "created_at", "created_by", "cve", "cwe", "data", "description", "documentation_url", "entity_checked", "is_published", "is_testing", "language", "last_updated_at", "last_updated_by", "name", "regex", "severity", "short_description", "should_use_ai_fix", "tests", "tree_sitter_query", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Arguments = all.Arguments
	o.Category = all.Category
	o.Checksum = all.Checksum
	o.Code = all.Code
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Cve = all.Cve
	o.Cwe = all.Cwe
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Description = all.Description
	o.DocumentationUrl = all.DocumentationUrl
	o.EntityChecked = all.EntityChecked
	o.IsPublished = all.IsPublished
	o.IsTesting = all.IsTesting
	o.Language = all.Language
	o.LastUpdatedAt = all.LastUpdatedAt
	o.LastUpdatedBy = all.LastUpdatedBy
	o.Name = all.Name
	o.Regex = all.Regex
	o.Severity = all.Severity
	o.ShortDescription = all.ShortDescription
	o.ShouldUseAiFix = all.ShouldUseAiFix
	o.Tests = all.Tests
	o.TreeSitterQuery = all.TreeSitterQuery
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
