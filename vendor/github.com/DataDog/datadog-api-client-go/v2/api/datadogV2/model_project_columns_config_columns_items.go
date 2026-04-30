// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectColumnsConfigColumnsItems Configuration for a single column in a project board view.
type ProjectColumnsConfigColumnsItems struct {
	// Sort configuration for a project board column.
	Sort *ProjectColumnsConfigColumnsItemsSort `json:"sort,omitempty"`
	// The field used to sort items in this column.
	SortField *string `json:"sort_field,omitempty"`
	// The type of column.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectColumnsConfigColumnsItems instantiates a new ProjectColumnsConfigColumnsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectColumnsConfigColumnsItems() *ProjectColumnsConfigColumnsItems {
	this := ProjectColumnsConfigColumnsItems{}
	return &this
}

// NewProjectColumnsConfigColumnsItemsWithDefaults instantiates a new ProjectColumnsConfigColumnsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectColumnsConfigColumnsItemsWithDefaults() *ProjectColumnsConfigColumnsItems {
	this := ProjectColumnsConfigColumnsItems{}
	return &this
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProjectColumnsConfigColumnsItems) GetSort() ProjectColumnsConfigColumnsItemsSort {
	if o == nil || o.Sort == nil {
		var ret ProjectColumnsConfigColumnsItemsSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectColumnsConfigColumnsItems) GetSortOk() (*ProjectColumnsConfigColumnsItemsSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProjectColumnsConfigColumnsItems) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ProjectColumnsConfigColumnsItemsSort and assigns it to the Sort field.
func (o *ProjectColumnsConfigColumnsItems) SetSort(v ProjectColumnsConfigColumnsItemsSort) {
	o.Sort = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *ProjectColumnsConfigColumnsItems) GetSortField() string {
	if o == nil || o.SortField == nil {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectColumnsConfigColumnsItems) GetSortFieldOk() (*string, bool) {
	if o == nil || o.SortField == nil {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *ProjectColumnsConfigColumnsItems) HasSortField() bool {
	return o != nil && o.SortField != nil
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *ProjectColumnsConfigColumnsItems) SetSortField(v string) {
	o.SortField = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectColumnsConfigColumnsItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectColumnsConfigColumnsItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectColumnsConfigColumnsItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProjectColumnsConfigColumnsItems) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectColumnsConfigColumnsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.SortField != nil {
		toSerialize["sort_field"] = o.SortField
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
func (o *ProjectColumnsConfigColumnsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sort      *ProjectColumnsConfigColumnsItemsSort `json:"sort,omitempty"`
		SortField *string                               `json:"sort_field,omitempty"`
		Type      *string                               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sort", "sort_field", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.SortField = all.SortField
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
