// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsScalarColumn A column in a scalar response.
type ProductAnalyticsScalarColumn struct {
	// Metadata associated with a scalar response column, including optional unit information.
	Meta *ProductAnalyticsScalarColumnMeta `json:"meta,omitempty"`
	// Column name (facet name for group-by, or "query").
	Name *string `json:"name,omitempty"`
	// Column type.
	Type *ProductAnalyticsScalarColumnType `json:"type,omitempty"`
	// Column values.
	Values []interface{} `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsScalarColumn instantiates a new ProductAnalyticsScalarColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsScalarColumn() *ProductAnalyticsScalarColumn {
	this := ProductAnalyticsScalarColumn{}
	return &this
}

// NewProductAnalyticsScalarColumnWithDefaults instantiates a new ProductAnalyticsScalarColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsScalarColumnWithDefaults() *ProductAnalyticsScalarColumn {
	this := ProductAnalyticsScalarColumn{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProductAnalyticsScalarColumn) GetMeta() ProductAnalyticsScalarColumnMeta {
	if o == nil || o.Meta == nil {
		var ret ProductAnalyticsScalarColumnMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsScalarColumn) GetMetaOk() (*ProductAnalyticsScalarColumnMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProductAnalyticsScalarColumn) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ProductAnalyticsScalarColumnMeta and assigns it to the Meta field.
func (o *ProductAnalyticsScalarColumn) SetMeta(v ProductAnalyticsScalarColumnMeta) {
	o.Meta = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsScalarColumn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsScalarColumn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsScalarColumn) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsScalarColumn) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAnalyticsScalarColumn) GetType() ProductAnalyticsScalarColumnType {
	if o == nil || o.Type == nil {
		var ret ProductAnalyticsScalarColumnType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsScalarColumn) GetTypeOk() (*ProductAnalyticsScalarColumnType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAnalyticsScalarColumn) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ProductAnalyticsScalarColumnType and assigns it to the Type field.
func (o *ProductAnalyticsScalarColumn) SetType(v ProductAnalyticsScalarColumnType) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ProductAnalyticsScalarColumn) GetValues() []interface{} {
	if o == nil || o.Values == nil {
		var ret []interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsScalarColumn) GetValuesOk() (*[]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ProductAnalyticsScalarColumn) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []interface{} and assigns it to the Values field.
func (o *ProductAnalyticsScalarColumn) SetValues(v []interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsScalarColumn) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsScalarColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta   *ProductAnalyticsScalarColumnMeta `json:"meta,omitempty"`
		Name   *string                           `json:"name,omitempty"`
		Type   *ProductAnalyticsScalarColumnType `json:"type,omitempty"`
		Values []interface{}                     `json:"values,omitempty"`
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
