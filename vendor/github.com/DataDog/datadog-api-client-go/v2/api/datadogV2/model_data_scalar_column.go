// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataScalarColumn A column containing the numerical results for a formula or query.
type DataScalarColumn struct {
	// Metadata for the resulting numerical values.
	Meta *ScalarMeta `json:"meta,omitempty"`
	// The name referencing the formula or query for this column.
	Name *string `json:"name,omitempty"`
	// The type of column present for numbers.
	Type *ScalarColumnTypeNumber `json:"type,omitempty"`
	// The array of numerical values for one formula or query.
	Values []*float64 `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataScalarColumn instantiates a new DataScalarColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataScalarColumn() *DataScalarColumn {
	this := DataScalarColumn{}
	var typeVar ScalarColumnTypeNumber = SCALARCOLUMNTYPENUMBER_NUMBER
	this.Type = &typeVar
	return &this
}

// NewDataScalarColumnWithDefaults instantiates a new DataScalarColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataScalarColumnWithDefaults() *DataScalarColumn {
	this := DataScalarColumn{}
	var typeVar ScalarColumnTypeNumber = SCALARCOLUMNTYPENUMBER_NUMBER
	this.Type = &typeVar
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DataScalarColumn) GetMeta() ScalarMeta {
	if o == nil || o.Meta == nil {
		var ret ScalarMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataScalarColumn) GetMetaOk() (*ScalarMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DataScalarColumn) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ScalarMeta and assigns it to the Meta field.
func (o *DataScalarColumn) SetMeta(v ScalarMeta) {
	o.Meta = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataScalarColumn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataScalarColumn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataScalarColumn) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataScalarColumn) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DataScalarColumn) GetType() ScalarColumnTypeNumber {
	if o == nil || o.Type == nil {
		var ret ScalarColumnTypeNumber
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataScalarColumn) GetTypeOk() (*ScalarColumnTypeNumber, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DataScalarColumn) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ScalarColumnTypeNumber and assigns it to the Type field.
func (o *DataScalarColumn) SetType(v ScalarColumnTypeNumber) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DataScalarColumn) GetValues() []*float64 {
	if o == nil || o.Values == nil {
		var ret []*float64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataScalarColumn) GetValuesOk() (*[]*float64, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DataScalarColumn) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []*float64 and assigns it to the Values field.
func (o *DataScalarColumn) SetValues(v []*float64) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataScalarColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
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
func (o *DataScalarColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta   *ScalarMeta             `json:"meta,omitempty"`
		Name   *string                 `json:"name,omitempty"`
		Type   *ScalarColumnTypeNumber `json:"type,omitempty"`
		Values []*float64              `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta", "name", "type", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
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
