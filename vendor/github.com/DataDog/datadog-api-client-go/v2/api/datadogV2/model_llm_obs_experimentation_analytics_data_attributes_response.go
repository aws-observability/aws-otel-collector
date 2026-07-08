// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationAnalyticsDataAttributesResponse Attributes of an analytics response.
type LLMObsExperimentationAnalyticsDataAttributesResponse struct {
	// Total number of events matched by the query before grouping.
	HitCount int64 `json:"hit_count"`
	// Analytics query result containing all buckets.
	Result LLMObsExperimentationAnalyticsResult `json:"result"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationAnalyticsDataAttributesResponse instantiates a new LLMObsExperimentationAnalyticsDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationAnalyticsDataAttributesResponse(hitCount int64, result LLMObsExperimentationAnalyticsResult) *LLMObsExperimentationAnalyticsDataAttributesResponse {
	this := LLMObsExperimentationAnalyticsDataAttributesResponse{}
	this.HitCount = hitCount
	this.Result = result
	return &this
}

// NewLLMObsExperimentationAnalyticsDataAttributesResponseWithDefaults instantiates a new LLMObsExperimentationAnalyticsDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationAnalyticsDataAttributesResponseWithDefaults() *LLMObsExperimentationAnalyticsDataAttributesResponse {
	this := LLMObsExperimentationAnalyticsDataAttributesResponse{}
	return &this
}

// GetHitCount returns the HitCount field value.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) GetHitCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.HitCount
}

// GetHitCountOk returns a tuple with the HitCount field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) GetHitCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HitCount, true
}

// SetHitCount sets field value.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) SetHitCount(v int64) {
	o.HitCount = v
}

// GetResult returns the Result field value.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) GetResult() LLMObsExperimentationAnalyticsResult {
	if o == nil {
		var ret LLMObsExperimentationAnalyticsResult
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) GetResultOk() (*LLMObsExperimentationAnalyticsResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) SetResult(v LLMObsExperimentationAnalyticsResult) {
	o.Result = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationAnalyticsDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hit_count"] = o.HitCount
	toSerialize["result"] = o.Result

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationAnalyticsDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HitCount *int64                                `json:"hit_count"`
		Result   *LLMObsExperimentationAnalyticsResult `json:"result"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HitCount == nil {
		return fmt.Errorf("required field hit_count missing")
	}
	if all.Result == nil {
		return fmt.Errorf("required field result missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hit_count", "result"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HitCount = *all.HitCount
	if all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = *all.Result

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
