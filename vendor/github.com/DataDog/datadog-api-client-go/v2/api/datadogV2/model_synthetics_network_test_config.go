// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTestConfig Configuration object for a Network Path test.
type SyntheticsNetworkTestConfig struct {
	// Array of assertions used for the test.
	Assertions []SyntheticsNetworkAssertion `json:"assertions,omitempty"`
	// Object describing the request for a Network Path test.
	Request *SyntheticsNetworkTestRequest `json:"request,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsNetworkTestConfig instantiates a new SyntheticsNetworkTestConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsNetworkTestConfig() *SyntheticsNetworkTestConfig {
	this := SyntheticsNetworkTestConfig{}
	return &this
}

// NewSyntheticsNetworkTestConfigWithDefaults instantiates a new SyntheticsNetworkTestConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsNetworkTestConfigWithDefaults() *SyntheticsNetworkTestConfig {
	this := SyntheticsNetworkTestConfig{}
	return &this
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestConfig) GetAssertions() []SyntheticsNetworkAssertion {
	if o == nil || o.Assertions == nil {
		var ret []SyntheticsNetworkAssertion
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestConfig) GetAssertionsOk() (*[]SyntheticsNetworkAssertion, bool) {
	if o == nil || o.Assertions == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestConfig) HasAssertions() bool {
	return o != nil && o.Assertions != nil
}

// SetAssertions gets a reference to the given []SyntheticsNetworkAssertion and assigns it to the Assertions field.
func (o *SyntheticsNetworkTestConfig) SetAssertions(v []SyntheticsNetworkAssertion) {
	o.Assertions = v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestConfig) GetRequest() SyntheticsNetworkTestRequest {
	if o == nil || o.Request == nil {
		var ret SyntheticsNetworkTestRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestConfig) GetRequestOk() (*SyntheticsNetworkTestRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestConfig) HasRequest() bool {
	return o != nil && o.Request != nil
}

// SetRequest gets a reference to the given SyntheticsNetworkTestRequest and assigns it to the Request field.
func (o *SyntheticsNetworkTestConfig) SetRequest(v SyntheticsNetworkTestRequest) {
	o.Request = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsNetworkTestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assertions != nil {
		toSerialize["assertions"] = o.Assertions
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsNetworkTestConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assertions []SyntheticsNetworkAssertion  `json:"assertions,omitempty"`
		Request    *SyntheticsNetworkTestRequest `json:"request,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assertions", "request"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assertions = all.Assertions
	if all.Request != nil && all.Request.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Request = all.Request

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
