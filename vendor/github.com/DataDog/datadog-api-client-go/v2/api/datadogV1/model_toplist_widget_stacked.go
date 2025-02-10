// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistWidgetStacked Top list widget stacked display options.
type ToplistWidgetStacked struct {
	// Top list widget stacked legend behavior.
	Legend *ToplistWidgetLegend `json:"legend,omitempty"`
	// Top list widget stacked display type.
	Type ToplistWidgetStackedType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewToplistWidgetStacked instantiates a new ToplistWidgetStacked object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToplistWidgetStacked(typeVar ToplistWidgetStackedType) *ToplistWidgetStacked {
	this := ToplistWidgetStacked{}
	this.Type = typeVar
	return &this
}

// NewToplistWidgetStackedWithDefaults instantiates a new ToplistWidgetStacked object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewToplistWidgetStackedWithDefaults() *ToplistWidgetStacked {
	this := ToplistWidgetStacked{}
	var typeVar ToplistWidgetStackedType = TOPLISTWIDGETSTACKEDTYPE_STACKED
	this.Type = typeVar
	return &this
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *ToplistWidgetStacked) GetLegend() ToplistWidgetLegend {
	if o == nil || o.Legend == nil {
		var ret ToplistWidgetLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToplistWidgetStacked) GetLegendOk() (*ToplistWidgetLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *ToplistWidgetStacked) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given ToplistWidgetLegend and assigns it to the Legend field.
func (o *ToplistWidgetStacked) SetLegend(v ToplistWidgetLegend) {
	o.Legend = &v
}

// GetType returns the Type field value.
func (o *ToplistWidgetStacked) GetType() ToplistWidgetStackedType {
	if o == nil {
		var ret ToplistWidgetStackedType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ToplistWidgetStacked) GetTypeOk() (*ToplistWidgetStackedType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ToplistWidgetStacked) SetType(v ToplistWidgetStackedType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ToplistWidgetStacked) MarshalJSON() ([]byte, error) {
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
func (o *ToplistWidgetStacked) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Legend *ToplistWidgetLegend      `json:"legend,omitempty"`
		Type   *ToplistWidgetStackedType `json:"type"`
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
