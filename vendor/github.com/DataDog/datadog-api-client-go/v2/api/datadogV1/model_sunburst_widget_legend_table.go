// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstWidgetLegendTable Configuration of table-based legend.
type SunburstWidgetLegendTable struct {
	// Whether or not to show a table legend.
	Type SunburstWidgetLegendTableType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSunburstWidgetLegendTable instantiates a new SunburstWidgetLegendTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSunburstWidgetLegendTable(typeVar SunburstWidgetLegendTableType) *SunburstWidgetLegendTable {
	this := SunburstWidgetLegendTable{}
	this.Type = typeVar
	return &this
}

// NewSunburstWidgetLegendTableWithDefaults instantiates a new SunburstWidgetLegendTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSunburstWidgetLegendTableWithDefaults() *SunburstWidgetLegendTable {
	this := SunburstWidgetLegendTable{}
	return &this
}

// GetType returns the Type field value.
func (o *SunburstWidgetLegendTable) GetType() SunburstWidgetLegendTableType {
	if o == nil {
		var ret SunburstWidgetLegendTableType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SunburstWidgetLegendTable) GetTypeOk() (*SunburstWidgetLegendTableType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SunburstWidgetLegendTable) SetType(v SunburstWidgetLegendTableType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SunburstWidgetLegendTable) MarshalJSON() ([]byte, error) {
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
func (o *SunburstWidgetLegendTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *SunburstWidgetLegendTableType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
