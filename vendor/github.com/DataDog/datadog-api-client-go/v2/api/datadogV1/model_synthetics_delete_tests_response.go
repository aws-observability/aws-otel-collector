// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDeleteTestsResponse Response object for deleting Synthetic tests.
type SyntheticsDeleteTestsResponse struct {
	// Array of objects containing a deleted Synthetic test ID with
	// the associated deletion timestamp.
	DeletedTests []SyntheticsDeletedTest `json:"deleted_tests,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsDeleteTestsResponse instantiates a new SyntheticsDeleteTestsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsDeleteTestsResponse() *SyntheticsDeleteTestsResponse {
	this := SyntheticsDeleteTestsResponse{}
	return &this
}

// NewSyntheticsDeleteTestsResponseWithDefaults instantiates a new SyntheticsDeleteTestsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsDeleteTestsResponseWithDefaults() *SyntheticsDeleteTestsResponse {
	this := SyntheticsDeleteTestsResponse{}
	return &this
}

// GetDeletedTests returns the DeletedTests field value if set, zero value otherwise.
func (o *SyntheticsDeleteTestsResponse) GetDeletedTests() []SyntheticsDeletedTest {
	if o == nil || o.DeletedTests == nil {
		var ret []SyntheticsDeletedTest
		return ret
	}
	return o.DeletedTests
}

// GetDeletedTestsOk returns a tuple with the DeletedTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDeleteTestsResponse) GetDeletedTestsOk() (*[]SyntheticsDeletedTest, bool) {
	if o == nil || o.DeletedTests == nil {
		return nil, false
	}
	return &o.DeletedTests, true
}

// HasDeletedTests returns a boolean if a field has been set.
func (o *SyntheticsDeleteTestsResponse) HasDeletedTests() bool {
	return o != nil && o.DeletedTests != nil
}

// SetDeletedTests gets a reference to the given []SyntheticsDeletedTest and assigns it to the DeletedTests field.
func (o *SyntheticsDeleteTestsResponse) SetDeletedTests(v []SyntheticsDeletedTest) {
	o.DeletedTests = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsDeleteTestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedTests != nil {
		toSerialize["deleted_tests"] = o.DeletedTests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsDeleteTestsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedTests []SyntheticsDeletedTest `json:"deleted_tests,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted_tests"})
	} else {
		return err
	}
	o.DeletedTests = all.DeletedTests

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
