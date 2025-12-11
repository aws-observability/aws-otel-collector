// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoResponseDataAttributes
type FacetInfoResponseDataAttributes struct {
	//
	Result *FacetInfoResponseDataAttributesResult `json:"result,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetInfoResponseDataAttributes instantiates a new FacetInfoResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetInfoResponseDataAttributes() *FacetInfoResponseDataAttributes {
	this := FacetInfoResponseDataAttributes{}
	return &this
}

// NewFacetInfoResponseDataAttributesWithDefaults instantiates a new FacetInfoResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetInfoResponseDataAttributesWithDefaults() *FacetInfoResponseDataAttributes {
	this := FacetInfoResponseDataAttributes{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributes) GetResult() FacetInfoResponseDataAttributesResult {
	if o == nil || o.Result == nil {
		var ret FacetInfoResponseDataAttributesResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributes) GetResultOk() (*FacetInfoResponseDataAttributesResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributes) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given FacetInfoResponseDataAttributesResult and assigns it to the Result field.
func (o *FacetInfoResponseDataAttributes) SetResult(v FacetInfoResponseDataAttributesResult) {
	o.Result = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetInfoResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FacetInfoResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Result *FacetInfoResponseDataAttributesResult `json:"result,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"result"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
