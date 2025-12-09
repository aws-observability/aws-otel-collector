// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentRecommendation Resource recommendation for a single Spark component (driver or executor). Contains estimation data used to patch Spark job specs.
type ComponentRecommendation struct {
	// Recommended resource values for a Spark driver or executor, derived from recent real usage metrics. Used by SPA to propose more efficient pod sizing.
	Estimation Estimation `json:"estimation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentRecommendation instantiates a new ComponentRecommendation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentRecommendation(estimation Estimation) *ComponentRecommendation {
	this := ComponentRecommendation{}
	this.Estimation = estimation
	return &this
}

// NewComponentRecommendationWithDefaults instantiates a new ComponentRecommendation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentRecommendationWithDefaults() *ComponentRecommendation {
	this := ComponentRecommendation{}
	return &this
}

// GetEstimation returns the Estimation field value.
func (o *ComponentRecommendation) GetEstimation() Estimation {
	if o == nil {
		var ret Estimation
		return ret
	}
	return o.Estimation
}

// GetEstimationOk returns a tuple with the Estimation field value
// and a boolean to check if the value has been set.
func (o *ComponentRecommendation) GetEstimationOk() (*Estimation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Estimation, true
}

// SetEstimation sets field value.
func (o *ComponentRecommendation) SetEstimation(v Estimation) {
	o.Estimation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["estimation"] = o.Estimation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentRecommendation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Estimation *Estimation `json:"estimation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Estimation == nil {
		return fmt.Errorf("required field estimation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"estimation"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Estimation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Estimation = *all.Estimation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
