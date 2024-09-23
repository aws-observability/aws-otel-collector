// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupScalarColumn A column containing the tag keys and values in a group.
type GroupScalarColumn struct {
	// The name of the tag key or group.
	Name *string `json:"name,omitempty"`
	// The type of column present for groups.
	Type *ScalarColumnTypeGroup `json:"type,omitempty"`
	// The array of tag values for each group found for the results of the formulas or queries.
	Values [][]string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGroupScalarColumn instantiates a new GroupScalarColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGroupScalarColumn() *GroupScalarColumn {
	this := GroupScalarColumn{}
	var typeVar ScalarColumnTypeGroup = SCALARCOLUMNTYPEGROUP_GROUP
	this.Type = &typeVar
	return &this
}

// NewGroupScalarColumnWithDefaults instantiates a new GroupScalarColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGroupScalarColumnWithDefaults() *GroupScalarColumn {
	this := GroupScalarColumn{}
	var typeVar ScalarColumnTypeGroup = SCALARCOLUMNTYPEGROUP_GROUP
	this.Type = &typeVar
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupScalarColumn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupScalarColumn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupScalarColumn) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupScalarColumn) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupScalarColumn) GetType() ScalarColumnTypeGroup {
	if o == nil || o.Type == nil {
		var ret ScalarColumnTypeGroup
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupScalarColumn) GetTypeOk() (*ScalarColumnTypeGroup, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupScalarColumn) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ScalarColumnTypeGroup and assigns it to the Type field.
func (o *GroupScalarColumn) SetType(v ScalarColumnTypeGroup) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GroupScalarColumn) GetValues() [][]string {
	if o == nil || o.Values == nil {
		var ret [][]string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupScalarColumn) GetValuesOk() (*[][]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GroupScalarColumn) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given [][]string and assigns it to the Values field.
func (o *GroupScalarColumn) SetValues(v [][]string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GroupScalarColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GroupScalarColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string                `json:"name,omitempty"`
		Type   *ScalarColumnTypeGroup `json:"type,omitempty"`
		Values [][]string             `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
