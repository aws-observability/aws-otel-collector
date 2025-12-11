// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoResponseDataAttributesResult
type FacetInfoResponseDataAttributesResult struct {
	//
	Range *FacetInfoResponseDataAttributesResultRange `json:"range,omitempty"`
	//
	Values []FacetInfoResponseDataAttributesResultValuesItems `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetInfoResponseDataAttributesResult instantiates a new FacetInfoResponseDataAttributesResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetInfoResponseDataAttributesResult() *FacetInfoResponseDataAttributesResult {
	this := FacetInfoResponseDataAttributesResult{}
	return &this
}

// NewFacetInfoResponseDataAttributesResultWithDefaults instantiates a new FacetInfoResponseDataAttributesResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetInfoResponseDataAttributesResultWithDefaults() *FacetInfoResponseDataAttributesResult {
	this := FacetInfoResponseDataAttributesResult{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributesResult) GetRange() FacetInfoResponseDataAttributesResultRange {
	if o == nil || o.Range == nil {
		var ret FacetInfoResponseDataAttributesResultRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributesResult) GetRangeOk() (*FacetInfoResponseDataAttributesResultRange, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResult) HasRange() bool {
	return o != nil && o.Range != nil
}

// SetRange gets a reference to the given FacetInfoResponseDataAttributesResultRange and assigns it to the Range field.
func (o *FacetInfoResponseDataAttributesResult) SetRange(v FacetInfoResponseDataAttributesResultRange) {
	o.Range = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributesResult) GetValues() []FacetInfoResponseDataAttributesResultValuesItems {
	if o == nil || o.Values == nil {
		var ret []FacetInfoResponseDataAttributesResultValuesItems
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributesResult) GetValuesOk() (*[]FacetInfoResponseDataAttributesResultValuesItems, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResult) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []FacetInfoResponseDataAttributesResultValuesItems and assigns it to the Values field.
func (o *FacetInfoResponseDataAttributesResult) SetValues(v []FacetInfoResponseDataAttributesResultValuesItems) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetInfoResponseDataAttributesResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
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
func (o *FacetInfoResponseDataAttributesResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Range  *FacetInfoResponseDataAttributesResultRange        `json:"range,omitempty"`
		Values []FacetInfoResponseDataAttributesResultValuesItems `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"range", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Range != nil && all.Range.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Range = all.Range
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
