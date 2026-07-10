// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsRunsResponseAttributes Attributes of an LLM Observability patterns runs response.
type LLMObsPatternsRunsResponseAttributes struct {
	// List of patterns runs.
	Runs []LLMObsPatternsRunSummary `json:"runs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsRunsResponseAttributes instantiates a new LLMObsPatternsRunsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsRunsResponseAttributes(runs []LLMObsPatternsRunSummary) *LLMObsPatternsRunsResponseAttributes {
	this := LLMObsPatternsRunsResponseAttributes{}
	this.Runs = runs
	return &this
}

// NewLLMObsPatternsRunsResponseAttributesWithDefaults instantiates a new LLMObsPatternsRunsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsRunsResponseAttributesWithDefaults() *LLMObsPatternsRunsResponseAttributes {
	this := LLMObsPatternsRunsResponseAttributes{}
	return &this
}

// GetRuns returns the Runs field value.
func (o *LLMObsPatternsRunsResponseAttributes) GetRuns() []LLMObsPatternsRunSummary {
	if o == nil {
		var ret []LLMObsPatternsRunSummary
		return ret
	}
	return o.Runs
}

// GetRunsOk returns a tuple with the Runs field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsRunsResponseAttributes) GetRunsOk() (*[]LLMObsPatternsRunSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runs, true
}

// SetRuns sets field value.
func (o *LLMObsPatternsRunsResponseAttributes) SetRuns(v []LLMObsPatternsRunSummary) {
	o.Runs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsRunsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["runs"] = o.Runs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsRunsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Runs *[]LLMObsPatternsRunSummary `json:"runs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Runs == nil {
		return fmt.Errorf("required field runs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"runs"})
	} else {
		return err
	}
	o.Runs = *all.Runs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
