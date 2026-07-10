// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayMapCategorySubProcessor A category sub-processor for use inside an array-map processor.
// Unlike the top-level category processor, `is_enabled` is not supported.
type LogsArrayMapCategorySubProcessor struct {
	// Array of filters to match against a log and the corresponding value to assign.
	Categories []LogsCategoryProcessorCategory `json:"categories"`
	// Name of the sub-processor.
	Name *string `json:"name,omitempty"`
	// Target attribute path for the category value.
	Target string `json:"target"`
	// Type of logs category processor.
	Type LogsCategoryProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArrayMapCategorySubProcessor instantiates a new LogsArrayMapCategorySubProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArrayMapCategorySubProcessor(categories []LogsCategoryProcessorCategory, target string, typeVar LogsCategoryProcessorType) *LogsArrayMapCategorySubProcessor {
	this := LogsArrayMapCategorySubProcessor{}
	this.Categories = categories
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsArrayMapCategorySubProcessorWithDefaults instantiates a new LogsArrayMapCategorySubProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArrayMapCategorySubProcessorWithDefaults() *LogsArrayMapCategorySubProcessor {
	this := LogsArrayMapCategorySubProcessor{}
	var typeVar LogsCategoryProcessorType = LOGSCATEGORYPROCESSORTYPE_CATEGORY_PROCESSOR
	this.Type = typeVar
	return &this
}

// GetCategories returns the Categories field value.
func (o *LogsArrayMapCategorySubProcessor) GetCategories() []LogsCategoryProcessorCategory {
	if o == nil {
		var ret []LogsCategoryProcessorCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapCategorySubProcessor) GetCategoriesOk() (*[]LogsCategoryProcessorCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value.
func (o *LogsArrayMapCategorySubProcessor) SetCategories(v []LogsCategoryProcessorCategory) {
	o.Categories = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsArrayMapCategorySubProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArrayMapCategorySubProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsArrayMapCategorySubProcessor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsArrayMapCategorySubProcessor) SetName(v string) {
	o.Name = &v
}

// GetTarget returns the Target field value.
func (o *LogsArrayMapCategorySubProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapCategorySubProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsArrayMapCategorySubProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsArrayMapCategorySubProcessor) GetType() LogsCategoryProcessorType {
	if o == nil {
		var ret LogsCategoryProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArrayMapCategorySubProcessor) GetTypeOk() (*LogsCategoryProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArrayMapCategorySubProcessor) SetType(v LogsCategoryProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArrayMapCategorySubProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["categories"] = o.Categories
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArrayMapCategorySubProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Categories *[]LogsCategoryProcessorCategory `json:"categories"`
		Name       *string                          `json:"name,omitempty"`
		Target     *string                          `json:"target"`
		Type       *LogsCategoryProcessorType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Categories == nil {
		return fmt.Errorf("required field categories missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"categories", "name", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Categories = *all.Categories
	o.Name = all.Name
	o.Target = *all.Target
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
