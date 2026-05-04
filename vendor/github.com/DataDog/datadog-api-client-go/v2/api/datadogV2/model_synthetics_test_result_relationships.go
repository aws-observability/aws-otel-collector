// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultRelationships Relationships for a Synthetic test result.
type SyntheticsTestResultRelationships struct {
	// Relationship to the Synthetic test.
	Test *SyntheticsTestResultRelationshipTest `json:"test,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultRelationships instantiates a new SyntheticsTestResultRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultRelationships() *SyntheticsTestResultRelationships {
	this := SyntheticsTestResultRelationships{}
	return &this
}

// NewSyntheticsTestResultRelationshipsWithDefaults instantiates a new SyntheticsTestResultRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultRelationshipsWithDefaults() *SyntheticsTestResultRelationships {
	this := SyntheticsTestResultRelationships{}
	return &this
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *SyntheticsTestResultRelationships) GetTest() SyntheticsTestResultRelationshipTest {
	if o == nil || o.Test == nil {
		var ret SyntheticsTestResultRelationshipTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRelationships) GetTestOk() (*SyntheticsTestResultRelationshipTest, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *SyntheticsTestResultRelationships) HasTest() bool {
	return o != nil && o.Test != nil
}

// SetTest gets a reference to the given SyntheticsTestResultRelationshipTest and assigns it to the Test field.
func (o *SyntheticsTestResultRelationships) SetTest(v SyntheticsTestResultRelationshipTest) {
	o.Test = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Test *SyntheticsTestResultRelationshipTest `json:"test,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"test"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Test != nil && all.Test.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Test = all.Test

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
