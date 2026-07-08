// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationAnalyticsGroupBy A field to group analytics results by.
type LLMObsExperimentationAnalyticsGroupBy struct {
	// Field name to group by.
	Field string `json:"field"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationAnalyticsGroupBy instantiates a new LLMObsExperimentationAnalyticsGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationAnalyticsGroupBy(field string) *LLMObsExperimentationAnalyticsGroupBy {
	this := LLMObsExperimentationAnalyticsGroupBy{}
	this.Field = field
	return &this
}

// NewLLMObsExperimentationAnalyticsGroupByWithDefaults instantiates a new LLMObsExperimentationAnalyticsGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationAnalyticsGroupByWithDefaults() *LLMObsExperimentationAnalyticsGroupBy {
	this := LLMObsExperimentationAnalyticsGroupBy{}
	return &this
}

// GetField returns the Field field value.
func (o *LLMObsExperimentationAnalyticsGroupBy) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsGroupBy) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *LLMObsExperimentationAnalyticsGroupBy) SetField(v string) {
	o.Field = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationAnalyticsGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationAnalyticsGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field *string `json:"field"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field"})
	} else {
		return err
	}
	o.Field = *all.Field

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
