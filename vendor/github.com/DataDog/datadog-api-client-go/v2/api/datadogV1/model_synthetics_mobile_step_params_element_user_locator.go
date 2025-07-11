// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsElementUserLocator User locator to find the element.
type SyntheticsMobileStepParamsElementUserLocator struct {
	// Whether if the test should fail if the element cannot be found.
	FailTestOnCannotLocate *bool `json:"failTestOnCannotLocate,omitempty"`
	// Values of the user locator.
	Values []SyntheticsMobileStepParamsElementUserLocatorValuesItems `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsElementUserLocator instantiates a new SyntheticsMobileStepParamsElementUserLocator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsElementUserLocator() *SyntheticsMobileStepParamsElementUserLocator {
	this := SyntheticsMobileStepParamsElementUserLocator{}
	return &this
}

// NewSyntheticsMobileStepParamsElementUserLocatorWithDefaults instantiates a new SyntheticsMobileStepParamsElementUserLocator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsElementUserLocatorWithDefaults() *SyntheticsMobileStepParamsElementUserLocator {
	this := SyntheticsMobileStepParamsElementUserLocator{}
	return &this
}

// GetFailTestOnCannotLocate returns the FailTestOnCannotLocate field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElementUserLocator) GetFailTestOnCannotLocate() bool {
	if o == nil || o.FailTestOnCannotLocate == nil {
		var ret bool
		return ret
	}
	return *o.FailTestOnCannotLocate
}

// GetFailTestOnCannotLocateOk returns a tuple with the FailTestOnCannotLocate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElementUserLocator) GetFailTestOnCannotLocateOk() (*bool, bool) {
	if o == nil || o.FailTestOnCannotLocate == nil {
		return nil, false
	}
	return o.FailTestOnCannotLocate, true
}

// HasFailTestOnCannotLocate returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElementUserLocator) HasFailTestOnCannotLocate() bool {
	return o != nil && o.FailTestOnCannotLocate != nil
}

// SetFailTestOnCannotLocate gets a reference to the given bool and assigns it to the FailTestOnCannotLocate field.
func (o *SyntheticsMobileStepParamsElementUserLocator) SetFailTestOnCannotLocate(v bool) {
	o.FailTestOnCannotLocate = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElementUserLocator) GetValues() []SyntheticsMobileStepParamsElementUserLocatorValuesItems {
	if o == nil || o.Values == nil {
		var ret []SyntheticsMobileStepParamsElementUserLocatorValuesItems
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElementUserLocator) GetValuesOk() (*[]SyntheticsMobileStepParamsElementUserLocatorValuesItems, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElementUserLocator) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []SyntheticsMobileStepParamsElementUserLocatorValuesItems and assigns it to the Values field.
func (o *SyntheticsMobileStepParamsElementUserLocator) SetValues(v []SyntheticsMobileStepParamsElementUserLocatorValuesItems) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsElementUserLocator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FailTestOnCannotLocate != nil {
		toSerialize["failTestOnCannotLocate"] = o.FailTestOnCannotLocate
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
func (o *SyntheticsMobileStepParamsElementUserLocator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FailTestOnCannotLocate *bool                                                     `json:"failTestOnCannotLocate,omitempty"`
		Values                 []SyntheticsMobileStepParamsElementUserLocatorValuesItems `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failTestOnCannotLocate", "values"})
	} else {
		return err
	}
	o.FailTestOnCannotLocate = all.FailTestOnCannotLocate
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
