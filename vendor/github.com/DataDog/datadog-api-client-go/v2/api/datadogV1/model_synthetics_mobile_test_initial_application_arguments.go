// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestInitialApplicationArguments Initial application arguments for a mobile test.
type SyntheticsMobileTestInitialApplicationArguments struct {
	// Name of the property.
	PropertyNames *SyntheticsMobileTestInitialApplicationArgumentsPropertyNames `json:"propertyNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestInitialApplicationArguments instantiates a new SyntheticsMobileTestInitialApplicationArguments object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestInitialApplicationArguments() *SyntheticsMobileTestInitialApplicationArguments {
	this := SyntheticsMobileTestInitialApplicationArguments{}
	return &this
}

// NewSyntheticsMobileTestInitialApplicationArgumentsWithDefaults instantiates a new SyntheticsMobileTestInitialApplicationArguments object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestInitialApplicationArgumentsWithDefaults() *SyntheticsMobileTestInitialApplicationArguments {
	this := SyntheticsMobileTestInitialApplicationArguments{}
	return &this
}

// GetPropertyNames returns the PropertyNames field value if set, zero value otherwise.
func (o *SyntheticsMobileTestInitialApplicationArguments) GetPropertyNames() SyntheticsMobileTestInitialApplicationArgumentsPropertyNames {
	if o == nil || o.PropertyNames == nil {
		var ret SyntheticsMobileTestInitialApplicationArgumentsPropertyNames
		return ret
	}
	return *o.PropertyNames
}

// GetPropertyNamesOk returns a tuple with the PropertyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestInitialApplicationArguments) GetPropertyNamesOk() (*SyntheticsMobileTestInitialApplicationArgumentsPropertyNames, bool) {
	if o == nil || o.PropertyNames == nil {
		return nil, false
	}
	return o.PropertyNames, true
}

// HasPropertyNames returns a boolean if a field has been set.
func (o *SyntheticsMobileTestInitialApplicationArguments) HasPropertyNames() bool {
	return o != nil && o.PropertyNames != nil
}

// SetPropertyNames gets a reference to the given SyntheticsMobileTestInitialApplicationArgumentsPropertyNames and assigns it to the PropertyNames field.
func (o *SyntheticsMobileTestInitialApplicationArguments) SetPropertyNames(v SyntheticsMobileTestInitialApplicationArgumentsPropertyNames) {
	o.PropertyNames = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestInitialApplicationArguments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PropertyNames != nil {
		toSerialize["propertyNames"] = o.PropertyNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestInitialApplicationArguments) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PropertyNames *SyntheticsMobileTestInitialApplicationArgumentsPropertyNames `json:"propertyNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"propertyNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.PropertyNames != nil && all.PropertyNames.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PropertyNames = all.PropertyNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
