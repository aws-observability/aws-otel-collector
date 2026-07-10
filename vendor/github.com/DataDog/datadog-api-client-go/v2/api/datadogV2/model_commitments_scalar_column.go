// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsScalarColumn A column in a scalar response. When type is "group", values contains arrays of strings. When type is "number", values contains numeric values.
type CommitmentsScalarColumn struct {
	// Metadata for a scalar column, including unit information.
	Meta *CommitmentsScalarColumnMeta `json:"meta,omitempty"`
	// The column name.
	Name string `json:"name"`
	// The column type. "group" for dimension columns, "number" for metric columns.
	Type CommitmentsScalarColumnType `json:"type"`
	// Values for a scalar column. Arrays of strings for group columns, numbers for value columns.
	Values []interface{} `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsScalarColumn instantiates a new CommitmentsScalarColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsScalarColumn(name string, typeVar CommitmentsScalarColumnType, values []interface{}) *CommitmentsScalarColumn {
	this := CommitmentsScalarColumn{}
	this.Name = name
	this.Type = typeVar
	this.Values = values
	return &this
}

// NewCommitmentsScalarColumnWithDefaults instantiates a new CommitmentsScalarColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsScalarColumnWithDefaults() *CommitmentsScalarColumn {
	this := CommitmentsScalarColumn{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsScalarColumn) GetMeta() CommitmentsScalarColumnMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsScalarColumnMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsScalarColumn) GetMetaOk() (*CommitmentsScalarColumnMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsScalarColumn) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsScalarColumnMeta and assigns it to the Meta field.
func (o *CommitmentsScalarColumn) SetMeta(v CommitmentsScalarColumnMeta) {
	o.Meta = &v
}

// GetName returns the Name field value.
func (o *CommitmentsScalarColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommitmentsScalarColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CommitmentsScalarColumn) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *CommitmentsScalarColumn) GetType() CommitmentsScalarColumnType {
	if o == nil {
		var ret CommitmentsScalarColumnType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CommitmentsScalarColumn) GetTypeOk() (*CommitmentsScalarColumnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CommitmentsScalarColumn) SetType(v CommitmentsScalarColumnType) {
	o.Type = v
}

// GetValues returns the Values field value.
func (o *CommitmentsScalarColumn) GetValues() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CommitmentsScalarColumn) GetValuesOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CommitmentsScalarColumn) SetValues(v []interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsScalarColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsScalarColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta   *CommitmentsScalarColumnMeta `json:"meta,omitempty"`
		Name   *string                      `json:"name"`
		Type   *CommitmentsScalarColumnType `json:"type"`
		Values *[]interface{}               `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta", "name", "type", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
