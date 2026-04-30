// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetStacked Bar chart widget stacked display options.
type BarChartWidgetStacked struct {
	// Bar chart widget stacked legend behavior.
	Legend *BarChartWidgetLegend `json:"legend,omitempty"`
	// Bar chart widget stacked display type.
	Type BarChartWidgetStackedType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBarChartWidgetStacked instantiates a new BarChartWidgetStacked object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartWidgetStacked(typeVar BarChartWidgetStackedType) *BarChartWidgetStacked {
	this := BarChartWidgetStacked{}
	this.Type = typeVar
	return &this
}

// NewBarChartWidgetStackedWithDefaults instantiates a new BarChartWidgetStacked object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartWidgetStackedWithDefaults() *BarChartWidgetStacked {
	this := BarChartWidgetStacked{}
	var typeVar BarChartWidgetStackedType = BARCHARTWIDGETSTACKEDTYPE_STACKED
	this.Type = typeVar
	return &this
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *BarChartWidgetStacked) GetLegend() BarChartWidgetLegend {
	if o == nil || o.Legend == nil {
		var ret BarChartWidgetLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetStacked) GetLegendOk() (*BarChartWidgetLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *BarChartWidgetStacked) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given BarChartWidgetLegend and assigns it to the Legend field.
func (o *BarChartWidgetStacked) SetLegend(v BarChartWidgetLegend) {
	o.Legend = &v
}

// GetType returns the Type field value.
func (o *BarChartWidgetStacked) GetType() BarChartWidgetStackedType {
	if o == nil {
		var ret BarChartWidgetStackedType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BarChartWidgetStacked) GetTypeOk() (*BarChartWidgetStackedType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *BarChartWidgetStacked) SetType(v BarChartWidgetStackedType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartWidgetStacked) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BarChartWidgetStacked) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Legend *BarChartWidgetLegend      `json:"legend,omitempty"`
		Type   *BarChartWidgetStackedType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"legend", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Legend != nil && !all.Legend.IsValid() {
		hasInvalidField = true
	} else {
		o.Legend = all.Legend
	}
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
