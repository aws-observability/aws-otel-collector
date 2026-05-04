// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsRequestTest Details of what tests to update and their new attributes.
type UpdateFlakyTestsRequestTest struct {
	// The ID of the flaky test. This is the same ID returned by the Search flaky tests endpoint and corresponds to the test_fingerprint_fqn field in test run events.
	Id string `json:"id"`
	// The new state to set for the flaky test.
	NewState UpdateFlakyTestsRequestTestNewState `json:"new_state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFlakyTestsRequestTest instantiates a new UpdateFlakyTestsRequestTest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFlakyTestsRequestTest(id string, newState UpdateFlakyTestsRequestTestNewState) *UpdateFlakyTestsRequestTest {
	this := UpdateFlakyTestsRequestTest{}
	this.Id = id
	this.NewState = newState
	return &this
}

// NewUpdateFlakyTestsRequestTestWithDefaults instantiates a new UpdateFlakyTestsRequestTest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFlakyTestsRequestTestWithDefaults() *UpdateFlakyTestsRequestTest {
	this := UpdateFlakyTestsRequestTest{}
	return &this
}

// GetId returns the Id field value.
func (o *UpdateFlakyTestsRequestTest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsRequestTest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UpdateFlakyTestsRequestTest) SetId(v string) {
	o.Id = v
}

// GetNewState returns the NewState field value.
func (o *UpdateFlakyTestsRequestTest) GetNewState() UpdateFlakyTestsRequestTestNewState {
	if o == nil {
		var ret UpdateFlakyTestsRequestTestNewState
		return ret
	}
	return o.NewState
}

// GetNewStateOk returns a tuple with the NewState field value
// and a boolean to check if the value has been set.
func (o *UpdateFlakyTestsRequestTest) GetNewStateOk() (*UpdateFlakyTestsRequestTestNewState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewState, true
}

// SetNewState sets field value.
func (o *UpdateFlakyTestsRequestTest) SetNewState(v UpdateFlakyTestsRequestTestNewState) {
	o.NewState = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFlakyTestsRequestTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["new_state"] = o.NewState

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFlakyTestsRequestTest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *string                              `json:"id"`
		NewState *UpdateFlakyTestsRequestTestNewState `json:"new_state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.NewState == nil {
		return fmt.Errorf("required field new_state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "new_state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if !all.NewState.IsValid() {
		hasInvalidField = true
	} else {
		o.NewState = *all.NewState
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
