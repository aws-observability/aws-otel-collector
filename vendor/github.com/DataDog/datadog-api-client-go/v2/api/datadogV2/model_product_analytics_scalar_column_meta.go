// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsScalarColumnMeta Metadata associated with a scalar response column, including optional unit information.
type ProductAnalyticsScalarColumnMeta struct {
	// Unit definitions for the column values, if applicable.
	Unit []ProductAnalyticsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsScalarColumnMeta instantiates a new ProductAnalyticsScalarColumnMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsScalarColumnMeta() *ProductAnalyticsScalarColumnMeta {
	this := ProductAnalyticsScalarColumnMeta{}
	return &this
}

// NewProductAnalyticsScalarColumnMetaWithDefaults instantiates a new ProductAnalyticsScalarColumnMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsScalarColumnMetaWithDefaults() *ProductAnalyticsScalarColumnMeta {
	this := ProductAnalyticsScalarColumnMeta{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsScalarColumnMeta) GetUnit() []ProductAnalyticsUnit {
	if o == nil {
		var ret []ProductAnalyticsUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsScalarColumnMeta) GetUnitOk() (*[]ProductAnalyticsUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return &o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProductAnalyticsScalarColumnMeta) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given []ProductAnalyticsUnit and assigns it to the Unit field.
func (o *ProductAnalyticsScalarColumnMeta) SetUnit(v []ProductAnalyticsUnit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsScalarColumnMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsScalarColumnMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Unit []ProductAnalyticsUnit `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"unit"})
	} else {
		return err
	}
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
