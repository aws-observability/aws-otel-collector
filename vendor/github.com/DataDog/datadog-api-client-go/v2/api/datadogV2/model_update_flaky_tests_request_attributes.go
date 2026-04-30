// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsRequestAttributes Attributes for updating flaky test states.
type UpdateFlakyTestsRequestAttributes struct {
	// List of flaky tests to update.
	Tests []UpdateFlakyTestsRequestTest `json:"tests"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFlakyTestsRequestAttributes instantiates a new UpdateFlakyTestsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFlakyTestsRequestAttributes(tests []UpdateFlakyTestsRequestTest) *UpdateFlakyTestsRequestAttributes {
	this := UpdateFlakyTestsRequestAttributes{}
	this.Tests = tests
	return &this
}

// NewUpdateFlakyTestsRequestAttributesWithDefaults instantiates a new UpdateFlakyTestsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFlakyTestsRequestAttributesWithDefaults() *UpdateFlakyTestsRequestAttributes {
	this := UpdateFlakyTestsRequestAttributes{}
	return &this
}

// GetTests returns the Tests field value.
func (o *UpdateFlakyTestsRequestAttributes) GetTests() []UpdateFlakyTestsRequestTest {
	if o == nil {
		var ret []UpdateFlakyTestsRequestTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsRequestAttributes) GetTestsOk() (*[]UpdateFlakyTestsRequestTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tests, true
}

// SetTests sets field value.
func (o *UpdateFlakyTestsRequestAttributes) SetTests(v []UpdateFlakyTestsRequestTest) {
	o.Tests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFlakyTestsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tests"] = o.Tests

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFlakyTestsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tests *[]UpdateFlakyTestsRequestTest `json:"tests"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Tests == nil {
		return fmt.Errorf("required field tests missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tests"})
	} else {
		return err
	}
	o.Tests = *all.Tests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
