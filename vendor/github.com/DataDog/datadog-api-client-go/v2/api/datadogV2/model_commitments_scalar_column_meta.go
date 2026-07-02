// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsScalarColumnMeta Metadata for a scalar column, including unit information.
type CommitmentsScalarColumnMeta struct {
	// Unit metadata for a numeric metric.
	Unit CommitmentsUnit `json:"unit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsScalarColumnMeta instantiates a new CommitmentsScalarColumnMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsScalarColumnMeta(unit CommitmentsUnit) *CommitmentsScalarColumnMeta {
	this := CommitmentsScalarColumnMeta{}
	this.Unit = unit
	return &this
}

// NewCommitmentsScalarColumnMetaWithDefaults instantiates a new CommitmentsScalarColumnMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsScalarColumnMetaWithDefaults() *CommitmentsScalarColumnMeta {
	this := CommitmentsScalarColumnMeta{}
	return &this
}

// GetUnit returns the Unit field value.
func (o *CommitmentsScalarColumnMeta) GetUnit() CommitmentsUnit {
	if o == nil {
		var ret CommitmentsUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *CommitmentsScalarColumnMeta) GetUnitOk() (*CommitmentsUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *CommitmentsScalarColumnMeta) SetUnit(v CommitmentsUnit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsScalarColumnMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["unit"] = o.Unit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsScalarColumnMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Unit *CommitmentsUnit `json:"unit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"unit"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Unit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Unit = *all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
