// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerStandardPatternAttributes Attributes of the Sensitive Data Scanner standard pattern.
type SensitiveDataScannerStandardPatternAttributes struct {
	// Description of the standard pattern.
	Description *string `json:"description,omitempty"`
	// List of included keywords.
	IncludedKeywords []string `json:"included_keywords,omitempty"`
	// Name of the standard pattern.
	Name *string `json:"name,omitempty"`
	// (Deprecated) Regex to match, optionally documented for older standard rules. Refer to the `description` field to understand what the rule does.
	// Deprecated
	Pattern *string `json:"pattern,omitempty"`
	// Integer from 1 (high) to 5 (low) indicating standard pattern issue severity.
	Priority *int64 `json:"priority,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerStandardPatternAttributes instantiates a new SensitiveDataScannerStandardPatternAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerStandardPatternAttributes() *SensitiveDataScannerStandardPatternAttributes {
	this := SensitiveDataScannerStandardPatternAttributes{}
	return &this
}

// NewSensitiveDataScannerStandardPatternAttributesWithDefaults instantiates a new SensitiveDataScannerStandardPatternAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerStandardPatternAttributesWithDefaults() *SensitiveDataScannerStandardPatternAttributes {
	this := SensitiveDataScannerStandardPatternAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SensitiveDataScannerStandardPatternAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SensitiveDataScannerStandardPatternAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetIncludedKeywords returns the IncludedKeywords field value if set, zero value otherwise.
func (o *SensitiveDataScannerStandardPatternAttributes) GetIncludedKeywords() []string {
	if o == nil || o.IncludedKeywords == nil {
		var ret []string
		return ret
	}
	return o.IncludedKeywords
}

// GetIncludedKeywordsOk returns a tuple with the IncludedKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) GetIncludedKeywordsOk() (*[]string, bool) {
	if o == nil || o.IncludedKeywords == nil {
		return nil, false
	}
	return &o.IncludedKeywords, true
}

// HasIncludedKeywords returns a boolean if a field has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) HasIncludedKeywords() bool {
	return o != nil && o.IncludedKeywords != nil
}

// SetIncludedKeywords gets a reference to the given []string and assigns it to the IncludedKeywords field.
func (o *SensitiveDataScannerStandardPatternAttributes) SetIncludedKeywords(v []string) {
	o.IncludedKeywords = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensitiveDataScannerStandardPatternAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensitiveDataScannerStandardPatternAttributes) SetName(v string) {
	o.Name = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
// Deprecated
func (o *SensitiveDataScannerStandardPatternAttributes) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SensitiveDataScannerStandardPatternAttributes) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
// Deprecated
func (o *SensitiveDataScannerStandardPatternAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SensitiveDataScannerStandardPatternAttributes) GetPriority() int64 {
	if o == nil || o.Priority == nil {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *SensitiveDataScannerStandardPatternAttributes) SetPriority(v int64) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SensitiveDataScannerStandardPatternAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SensitiveDataScannerStandardPatternAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SensitiveDataScannerStandardPatternAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerStandardPatternAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IncludedKeywords != nil {
		toSerialize["included_keywords"] = o.IncludedKeywords
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerStandardPatternAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string  `json:"description,omitempty"`
		IncludedKeywords []string `json:"included_keywords,omitempty"`
		Name             *string  `json:"name,omitempty"`
		Pattern          *string  `json:"pattern,omitempty"`
		Priority         *int64   `json:"priority,omitempty"`
		Tags             []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "included_keywords", "name", "pattern", "priority", "tags"})
	} else {
		return err
	}
	o.Description = all.Description
	o.IncludedKeywords = all.IncludedKeywords
	o.Name = all.Name
	o.Pattern = all.Pattern
	o.Priority = all.Priority
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
