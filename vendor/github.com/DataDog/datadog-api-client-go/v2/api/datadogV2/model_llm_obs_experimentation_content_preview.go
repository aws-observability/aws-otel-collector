// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationContentPreview Options to control content preview truncation.
type LLMObsExperimentationContentPreview struct {
	// Maximum number of characters to include in content previews.
	Limit *int64 `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationContentPreview instantiates a new LLMObsExperimentationContentPreview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationContentPreview() *LLMObsExperimentationContentPreview {
	this := LLMObsExperimentationContentPreview{}
	return &this
}

// NewLLMObsExperimentationContentPreviewWithDefaults instantiates a new LLMObsExperimentationContentPreview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationContentPreviewWithDefaults() *LLMObsExperimentationContentPreview {
	this := LLMObsExperimentationContentPreview{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *LLMObsExperimentationContentPreview) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationContentPreview) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *LLMObsExperimentationContentPreview) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *LLMObsExperimentationContentPreview) SetLimit(v int64) {
	o.Limit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationContentPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationContentPreview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit *int64 `json:"limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit"})
	} else {
		return err
	}
	o.Limit = all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
