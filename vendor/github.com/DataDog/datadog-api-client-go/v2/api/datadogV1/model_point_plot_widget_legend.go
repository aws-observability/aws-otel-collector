// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotWidgetLegend Legend configuration for the point plot widget.
type PointPlotWidgetLegend struct {
	// Type of legend to show for the point plot widget.
	Type PointPlotWidgetLegendType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPointPlotWidgetLegend instantiates a new PointPlotWidgetLegend object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotWidgetLegend(typeVar PointPlotWidgetLegendType) *PointPlotWidgetLegend {
	this := PointPlotWidgetLegend{}
	this.Type = typeVar
	return &this
}

// NewPointPlotWidgetLegendWithDefaults instantiates a new PointPlotWidgetLegend object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotWidgetLegendWithDefaults() *PointPlotWidgetLegend {
	this := PointPlotWidgetLegend{}
	return &this
}

// GetType returns the Type field value.
func (o *PointPlotWidgetLegend) GetType() PointPlotWidgetLegendType {
	if o == nil {
		var ret PointPlotWidgetLegendType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetLegend) GetTypeOk() (*PointPlotWidgetLegendType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PointPlotWidgetLegend) SetType(v PointPlotWidgetLegendType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotWidgetLegend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PointPlotWidgetLegend) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *PointPlotWidgetLegendType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type"})
	} else {
		return err
	}

	hasInvalidField := false
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
