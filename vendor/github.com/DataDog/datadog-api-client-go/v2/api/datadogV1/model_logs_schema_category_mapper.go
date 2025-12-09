// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaCategoryMapper Use the Schema Category Mapper to categorize log event into enum fields.
// In the case of OCSF, they can be used to map sibling fields which are composed of an ID and a name.
//
// **Notes**:
//
//   - The syntax of the query is the one of Logs Explorer search bar.
//     The query can be done on any log attribute or tag, whether it is a facet or not.
//     Wildcards can also be used inside your query.
//   - Categories are executed in order and processing stops at the first match.
//     Make sure categories are properly ordered in case a log could match multiple queries.
//   - Sibling fields always have a numerical ID field and a human-readable string name.
//   - A fallback section handles cases where the name or ID value matches a specific value.
//     If the name matches "Other" or the ID matches 99, the value of the sibling name field will be pulled from a source field from the original log.
type LogsSchemaCategoryMapper struct {
	// Array of filters to match or not a log and their
	// corresponding `name` to assign a custom value to the log.
	Categories []LogsSchemaCategoryMapperCategory `json:"categories"`
	// Used to override hardcoded category values with a value pulled from a source attribute on the log.
	Fallback *LogsSchemaCategoryMapperFallback `json:"fallback,omitempty"`
	// Name of the logs schema category mapper.
	Name string `json:"name"`
	// Name of the target attributes which value is defined by the matching category.
	Targets LogsSchemaCategoryMapperTargets `json:"targets"`
	// Type of logs schema category mapper.
	Type LogsSchemaCategoryMapperType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsSchemaCategoryMapper instantiates a new LogsSchemaCategoryMapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsSchemaCategoryMapper(categories []LogsSchemaCategoryMapperCategory, name string, targets LogsSchemaCategoryMapperTargets, typeVar LogsSchemaCategoryMapperType) *LogsSchemaCategoryMapper {
	this := LogsSchemaCategoryMapper{}
	this.Categories = categories
	this.Name = name
	this.Targets = targets
	this.Type = typeVar
	return &this
}

// NewLogsSchemaCategoryMapperWithDefaults instantiates a new LogsSchemaCategoryMapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsSchemaCategoryMapperWithDefaults() *LogsSchemaCategoryMapper {
	this := LogsSchemaCategoryMapper{}
	return &this
}

// GetCategories returns the Categories field value.
func (o *LogsSchemaCategoryMapper) GetCategories() []LogsSchemaCategoryMapperCategory {
	if o == nil {
		var ret []LogsSchemaCategoryMapperCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapper) GetCategoriesOk() (*[]LogsSchemaCategoryMapperCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value.
func (o *LogsSchemaCategoryMapper) SetCategories(v []LogsSchemaCategoryMapperCategory) {
	o.Categories = v
}

// GetFallback returns the Fallback field value if set, zero value otherwise.
func (o *LogsSchemaCategoryMapper) GetFallback() LogsSchemaCategoryMapperFallback {
	if o == nil || o.Fallback == nil {
		var ret LogsSchemaCategoryMapperFallback
		return ret
	}
	return *o.Fallback
}

// GetFallbackOk returns a tuple with the Fallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapper) GetFallbackOk() (*LogsSchemaCategoryMapperFallback, bool) {
	if o == nil || o.Fallback == nil {
		return nil, false
	}
	return o.Fallback, true
}

// HasFallback returns a boolean if a field has been set.
func (o *LogsSchemaCategoryMapper) HasFallback() bool {
	return o != nil && o.Fallback != nil
}

// SetFallback gets a reference to the given LogsSchemaCategoryMapperFallback and assigns it to the Fallback field.
func (o *LogsSchemaCategoryMapper) SetFallback(v LogsSchemaCategoryMapperFallback) {
	o.Fallback = &v
}

// GetName returns the Name field value.
func (o *LogsSchemaCategoryMapper) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapper) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LogsSchemaCategoryMapper) SetName(v string) {
	o.Name = v
}

// GetTargets returns the Targets field value.
func (o *LogsSchemaCategoryMapper) GetTargets() LogsSchemaCategoryMapperTargets {
	if o == nil {
		var ret LogsSchemaCategoryMapperTargets
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapper) GetTargetsOk() (*LogsSchemaCategoryMapperTargets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *LogsSchemaCategoryMapper) SetTargets(v LogsSchemaCategoryMapperTargets) {
	o.Targets = v
}

// GetType returns the Type field value.
func (o *LogsSchemaCategoryMapper) GetType() LogsSchemaCategoryMapperType {
	if o == nil {
		var ret LogsSchemaCategoryMapperType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaCategoryMapper) GetTypeOk() (*LogsSchemaCategoryMapperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsSchemaCategoryMapper) SetType(v LogsSchemaCategoryMapperType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsSchemaCategoryMapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["categories"] = o.Categories
	if o.Fallback != nil {
		toSerialize["fallback"] = o.Fallback
	}
	toSerialize["name"] = o.Name
	toSerialize["targets"] = o.Targets
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsSchemaCategoryMapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Categories *[]LogsSchemaCategoryMapperCategory `json:"categories"`
		Fallback   *LogsSchemaCategoryMapperFallback   `json:"fallback,omitempty"`
		Name       *string                             `json:"name"`
		Targets    *LogsSchemaCategoryMapperTargets    `json:"targets"`
		Type       *LogsSchemaCategoryMapperType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Categories == nil {
		return fmt.Errorf("required field categories missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"categories", "fallback", "name", "targets", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Categories = *all.Categories
	if all.Fallback != nil && all.Fallback.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Fallback = all.Fallback
	o.Name = *all.Name
	if all.Targets.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Targets = *all.Targets
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
