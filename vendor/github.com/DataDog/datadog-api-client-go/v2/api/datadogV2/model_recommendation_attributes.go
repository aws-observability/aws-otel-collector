// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationAttributes Attributes of the SPA Recommendation resource. Contains recommendations for both driver and executor components.
type RecommendationAttributes struct {
	// The confidence level of the recommendation, expressed as a value between 0.0 (low confidence) and 1.0 (high confidence).
	ConfidenceLevel *float64 `json:"confidence_level,omitempty"`
	// Resource recommendation for a single Spark component (driver or executor). Contains estimation data used to patch Spark job specs.
	Driver ComponentRecommendation `json:"driver"`
	// Resource recommendation for a single Spark component (driver or executor). Contains estimation data used to patch Spark job specs.
	Executor ComponentRecommendation `json:"executor"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationAttributes instantiates a new RecommendationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationAttributes(driver ComponentRecommendation, executor ComponentRecommendation) *RecommendationAttributes {
	this := RecommendationAttributes{}
	this.Driver = driver
	this.Executor = executor
	return &this
}

// NewRecommendationAttributesWithDefaults instantiates a new RecommendationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationAttributesWithDefaults() *RecommendationAttributes {
	this := RecommendationAttributes{}
	return &this
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *RecommendationAttributes) GetConfidenceLevel() float64 {
	if o == nil || o.ConfidenceLevel == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationAttributes) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil || o.ConfidenceLevel == nil {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *RecommendationAttributes) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel != nil
}

// SetConfidenceLevel gets a reference to the given float64 and assigns it to the ConfidenceLevel field.
func (o *RecommendationAttributes) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel = &v
}

// GetDriver returns the Driver field value.
func (o *RecommendationAttributes) GetDriver() ComponentRecommendation {
	if o == nil {
		var ret ComponentRecommendation
		return ret
	}
	return o.Driver
}

// GetDriverOk returns a tuple with the Driver field value
// and a boolean to check if the value has been set.
func (o *RecommendationAttributes) GetDriverOk() (*ComponentRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Driver, true
}

// SetDriver sets field value.
func (o *RecommendationAttributes) SetDriver(v ComponentRecommendation) {
	o.Driver = v
}

// GetExecutor returns the Executor field value.
func (o *RecommendationAttributes) GetExecutor() ComponentRecommendation {
	if o == nil {
		var ret ComponentRecommendation
		return ret
	}
	return o.Executor
}

// GetExecutorOk returns a tuple with the Executor field value
// and a boolean to check if the value has been set.
func (o *RecommendationAttributes) GetExecutorOk() (*ComponentRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Executor, true
}

// SetExecutor sets field value.
func (o *RecommendationAttributes) SetExecutor(v ComponentRecommendation) {
	o.Executor = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfidenceLevel != nil {
		toSerialize["confidence_level"] = o.ConfidenceLevel
	}
	toSerialize["driver"] = o.Driver
	toSerialize["executor"] = o.Executor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfidenceLevel *float64                 `json:"confidence_level,omitempty"`
		Driver          *ComponentRecommendation `json:"driver"`
		Executor        *ComponentRecommendation `json:"executor"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Driver == nil {
		return fmt.Errorf("required field driver missing")
	}
	if all.Executor == nil {
		return fmt.Errorf("required field executor missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"confidence_level", "driver", "executor"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfidenceLevel = all.ConfidenceLevel
	if all.Driver.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Driver = *all.Driver
	if all.Executor.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Executor = *all.Executor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
