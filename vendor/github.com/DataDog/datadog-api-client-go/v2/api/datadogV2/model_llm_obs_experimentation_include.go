// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationInclude Additional data to include in the response.
type LLMObsExperimentationInclude struct {
	// When `true`, enrich results with author user data (name and email).
	UserData *bool `json:"user_data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationInclude instantiates a new LLMObsExperimentationInclude object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationInclude() *LLMObsExperimentationInclude {
	this := LLMObsExperimentationInclude{}
	var userData bool = false
	this.UserData = &userData
	return &this
}

// NewLLMObsExperimentationIncludeWithDefaults instantiates a new LLMObsExperimentationInclude object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationIncludeWithDefaults() *LLMObsExperimentationInclude {
	this := LLMObsExperimentationInclude{}
	var userData bool = false
	this.UserData = &userData
	return &this
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *LLMObsExperimentationInclude) GetUserData() bool {
	if o == nil || o.UserData == nil {
		var ret bool
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationInclude) GetUserDataOk() (*bool, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *LLMObsExperimentationInclude) HasUserData() bool {
	return o != nil && o.UserData != nil
}

// SetUserData gets a reference to the given bool and assigns it to the UserData field.
func (o *LLMObsExperimentationInclude) SetUserData(v bool) {
	o.UserData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationInclude) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationInclude) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserData *bool `json:"user_data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"user_data"})
	} else {
		return err
	}
	o.UserData = all.UserData

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
