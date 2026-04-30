// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsResponseAttributes Attributes for the update flaky test state response.
type UpdateFlakyTestsResponseAttributes struct {
	// `True` if any errors occurred during the update operations. `False` if all tests succeeded to be updated.
	HasErrors bool `json:"has_errors"`
	// Results of the update operation for each test.
	Results []UpdateFlakyTestsResponseResult `json:"results"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFlakyTestsResponseAttributes instantiates a new UpdateFlakyTestsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFlakyTestsResponseAttributes(hasErrors bool, results []UpdateFlakyTestsResponseResult) *UpdateFlakyTestsResponseAttributes {
	this := UpdateFlakyTestsResponseAttributes{}
	this.HasErrors = hasErrors
	this.Results = results
	return &this
}

// NewUpdateFlakyTestsResponseAttributesWithDefaults instantiates a new UpdateFlakyTestsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFlakyTestsResponseAttributesWithDefaults() *UpdateFlakyTestsResponseAttributes {
	this := UpdateFlakyTestsResponseAttributes{}
	return &this
}

// GetHasErrors returns the HasErrors field value.
func (o *UpdateFlakyTestsResponseAttributes) GetHasErrors() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasErrors
}

// GetHasErrorsOk returns a tuple with the HasErrors field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsResponseAttributes) GetHasErrorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasErrors, true
}

// SetHasErrors sets field value.
func (o *UpdateFlakyTestsResponseAttributes) SetHasErrors(v bool) {
	o.HasErrors = v
}

// GetResults returns the Results field value.
func (o *UpdateFlakyTestsResponseAttributes) GetResults() []UpdateFlakyTestsResponseResult {
	if o == nil {
		var ret []UpdateFlakyTestsResponseResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsResponseAttributes) GetResultsOk() (*[]UpdateFlakyTestsResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value.
func (o *UpdateFlakyTestsResponseAttributes) SetResults(v []UpdateFlakyTestsResponseResult) {
	o.Results = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFlakyTestsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["has_errors"] = o.HasErrors
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFlakyTestsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasErrors *bool                             `json:"has_errors"`
		Results   *[]UpdateFlakyTestsResponseResult `json:"results"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HasErrors == nil {
		return fmt.Errorf("required field has_errors missing")
	}
	if all.Results == nil {
		return fmt.Errorf("required field results missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_errors", "results"})
	} else {
		return err
	}
	o.HasErrors = *all.HasErrors
	o.Results = *all.Results

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
