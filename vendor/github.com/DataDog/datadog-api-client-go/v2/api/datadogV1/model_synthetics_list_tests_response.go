// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsListTestsResponse Object containing an array of Synthetic tests configuration.
type SyntheticsListTestsResponse struct {
	// Array of Synthetic tests configuration.
	Tests []SyntheticsTestDetails `json:"tests,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsListTestsResponse instantiates a new SyntheticsListTestsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsListTestsResponse() *SyntheticsListTestsResponse {
	this := SyntheticsListTestsResponse{}
	return &this
}

// NewSyntheticsListTestsResponseWithDefaults instantiates a new SyntheticsListTestsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsListTestsResponseWithDefaults() *SyntheticsListTestsResponse {
	this := SyntheticsListTestsResponse{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *SyntheticsListTestsResponse) GetTests() []SyntheticsTestDetails {
	if o == nil || o.Tests == nil {
		var ret []SyntheticsTestDetails
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsListTestsResponse) GetTestsOk() (*[]SyntheticsTestDetails, bool) {
	if o == nil || o.Tests == nil {
		return nil, false
	}
	return &o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *SyntheticsListTestsResponse) HasTests() bool {
	return o != nil && o.Tests != nil
}

// SetTests gets a reference to the given []SyntheticsTestDetails and assigns it to the Tests field.
func (o *SyntheticsListTestsResponse) SetTests(v []SyntheticsTestDetails) {
	o.Tests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsListTestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Tests != nil {
		toSerialize["tests"] = o.Tests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsListTestsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tests []SyntheticsTestDetails `json:"tests,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tests"})
	} else {
		return err
	}
	o.Tests = all.Tests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
