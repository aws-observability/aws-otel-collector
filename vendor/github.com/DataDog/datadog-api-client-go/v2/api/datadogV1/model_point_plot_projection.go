// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotProjection Projection configuration for the point plot widget.
type PointPlotProjection struct {
	// List of dimension mappings for the projection.
	Dimensions []PointPlotProjectionDimension `json:"dimensions"`
	// Additional columns to include in the projection.
	ExtraColumns []string `json:"extra_columns,omitempty"`
	// Type of the projection.
	Type PointPlotProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPointPlotProjection instantiates a new PointPlotProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotProjection(dimensions []PointPlotProjectionDimension, typeVar PointPlotProjectionType) *PointPlotProjection {
	this := PointPlotProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewPointPlotProjectionWithDefaults instantiates a new PointPlotProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotProjectionWithDefaults() *PointPlotProjection {
	this := PointPlotProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *PointPlotProjection) GetDimensions() []PointPlotProjectionDimension {
	if o == nil {
		var ret []PointPlotProjectionDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *PointPlotProjection) GetDimensionsOk() (*[]PointPlotProjectionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *PointPlotProjection) SetDimensions(v []PointPlotProjectionDimension) {
	o.Dimensions = v
}

// GetExtraColumns returns the ExtraColumns field value if set, zero value otherwise.
func (o *PointPlotProjection) GetExtraColumns() []string {
	if o == nil || o.ExtraColumns == nil {
		var ret []string
		return ret
	}
	return o.ExtraColumns
}

// GetExtraColumnsOk returns a tuple with the ExtraColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotProjection) GetExtraColumnsOk() (*[]string, bool) {
	if o == nil || o.ExtraColumns == nil {
		return nil, false
	}
	return &o.ExtraColumns, true
}

// HasExtraColumns returns a boolean if a field has been set.
func (o *PointPlotProjection) HasExtraColumns() bool {
	return o != nil && o.ExtraColumns != nil
}

// SetExtraColumns gets a reference to the given []string and assigns it to the ExtraColumns field.
func (o *PointPlotProjection) SetExtraColumns(v []string) {
	o.ExtraColumns = v
}

// GetType returns the Type field value.
func (o *PointPlotProjection) GetType() PointPlotProjectionType {
	if o == nil {
		var ret PointPlotProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PointPlotProjection) GetTypeOk() (*PointPlotProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PointPlotProjection) SetType(v PointPlotProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	if o.ExtraColumns != nil {
		toSerialize["extra_columns"] = o.ExtraColumns
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PointPlotProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions   *[]PointPlotProjectionDimension `json:"dimensions"`
		ExtraColumns []string                        `json:"extra_columns,omitempty"`
		Type         *PointPlotProjectionType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dimensions == nil {
		return fmt.Errorf("required field dimensions missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dimensions", "extra_columns", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Dimensions = *all.Dimensions
	o.ExtraColumns = all.ExtraColumns
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
