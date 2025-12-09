// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition
type ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition struct {
	//
	Col *int32 `json:"col,omitempty"`
	//
	Line *int32 `json:"line,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsPosition instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsPosition() *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsPositionWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsPositionWithDefaults() *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition{}
	return &this
}

// GetCol returns the Col field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) GetCol() int32 {
	if o == nil || o.Col == nil {
		var ret int32
		return ret
	}
	return *o.Col
}

// GetColOk returns a tuple with the Col field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) GetColOk() (*int32, bool) {
	if o == nil || o.Col == nil {
		return nil, false
	}
	return o.Col, true
}

// HasCol returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) HasCol() bool {
	return o != nil && o.Col != nil
}

// SetCol gets a reference to the given int32 and assigns it to the Col field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) SetCol(v int32) {
	o.Col = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) GetLine() int32 {
	if o == nil || o.Line == nil {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) GetLineOk() (*int32, bool) {
	if o == nil || o.Line == nil {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) HasLine() bool {
	return o != nil && o.Line != nil
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) SetLine(v int32) {
	o.Line = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Col != nil {
		toSerialize["col"] = o.Col
	}
	if o.Line != nil {
		toSerialize["line"] = o.Line
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Col  *int32 `json:"col,omitempty"`
		Line *int32 `json:"line,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"col", "line"})
	} else {
		return err
	}
	o.Col = all.Col
	o.Line = all.Line

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
