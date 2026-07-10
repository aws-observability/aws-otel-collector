// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScorecardScoreAttributes Attributes of a scorecard score.
type ScorecardScoreAttributes struct {
	// Dimension to group scores by.
	Aggregation *ScorecardScoresAggregation `json:"aggregation,omitempty"`
	// The denominator used to compute the score ratio.
	Denominator *int64 `json:"denominator,omitempty"`
	// The maturity level of the associated rule.
	Level *int64 `json:"level,omitempty"`
	// The numerator used to compute the score ratio.
	Numerator *int64 `json:"numerator,omitempty"`
	// The computed score ratio (numerator/denominator), from 0 to 1.
	Score *float64 `json:"score,omitempty"`
	// The total number of entities evaluated.
	TotalEntities *int64 `json:"total_entities,omitempty"`
	// The number of rules that failed.
	TotalFail *int64 `json:"total_fail,omitempty"`
	// The number of rules with no data.
	TotalNoData *int64 `json:"total_no_data,omitempty"`
	// The number of rules that passed.
	TotalPass *int64 `json:"total_pass,omitempty"`
	// The number of rules that were skipped.
	TotalSkip *int64 `json:"total_skip,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScorecardScoreAttributes instantiates a new ScorecardScoreAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScorecardScoreAttributes() *ScorecardScoreAttributes {
	this := ScorecardScoreAttributes{}
	return &this
}

// NewScorecardScoreAttributesWithDefaults instantiates a new ScorecardScoreAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScorecardScoreAttributesWithDefaults() *ScorecardScoreAttributes {
	this := ScorecardScoreAttributes{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetAggregation() ScorecardScoresAggregation {
	if o == nil || o.Aggregation == nil {
		var ret ScorecardScoresAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetAggregationOk() (*ScorecardScoresAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given ScorecardScoresAggregation and assigns it to the Aggregation field.
func (o *ScorecardScoreAttributes) SetAggregation(v ScorecardScoresAggregation) {
	o.Aggregation = &v
}

// GetDenominator returns the Denominator field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetDenominator() int64 {
	if o == nil || o.Denominator == nil {
		var ret int64
		return ret
	}
	return *o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetDenominatorOk() (*int64, bool) {
	if o == nil || o.Denominator == nil {
		return nil, false
	}
	return o.Denominator, true
}

// HasDenominator returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasDenominator() bool {
	return o != nil && o.Denominator != nil
}

// SetDenominator gets a reference to the given int64 and assigns it to the Denominator field.
func (o *ScorecardScoreAttributes) SetDenominator(v int64) {
	o.Denominator = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetLevel() int64 {
	if o == nil || o.Level == nil {
		var ret int64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetLevelOk() (*int64, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given int64 and assigns it to the Level field.
func (o *ScorecardScoreAttributes) SetLevel(v int64) {
	o.Level = &v
}

// GetNumerator returns the Numerator field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetNumerator() int64 {
	if o == nil || o.Numerator == nil {
		var ret int64
		return ret
	}
	return *o.Numerator
}

// GetNumeratorOk returns a tuple with the Numerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetNumeratorOk() (*int64, bool) {
	if o == nil || o.Numerator == nil {
		return nil, false
	}
	return o.Numerator, true
}

// HasNumerator returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasNumerator() bool {
	return o != nil && o.Numerator != nil
}

// SetNumerator gets a reference to the given int64 and assigns it to the Numerator field.
func (o *ScorecardScoreAttributes) SetNumerator(v int64) {
	o.Numerator = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetScore() float64 {
	if o == nil || o.Score == nil {
		var ret float64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetScoreOk() (*float64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given float64 and assigns it to the Score field.
func (o *ScorecardScoreAttributes) SetScore(v float64) {
	o.Score = &v
}

// GetTotalEntities returns the TotalEntities field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetTotalEntities() int64 {
	if o == nil || o.TotalEntities == nil {
		var ret int64
		return ret
	}
	return *o.TotalEntities
}

// GetTotalEntitiesOk returns a tuple with the TotalEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetTotalEntitiesOk() (*int64, bool) {
	if o == nil || o.TotalEntities == nil {
		return nil, false
	}
	return o.TotalEntities, true
}

// HasTotalEntities returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasTotalEntities() bool {
	return o != nil && o.TotalEntities != nil
}

// SetTotalEntities gets a reference to the given int64 and assigns it to the TotalEntities field.
func (o *ScorecardScoreAttributes) SetTotalEntities(v int64) {
	o.TotalEntities = &v
}

// GetTotalFail returns the TotalFail field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetTotalFail() int64 {
	if o == nil || o.TotalFail == nil {
		var ret int64
		return ret
	}
	return *o.TotalFail
}

// GetTotalFailOk returns a tuple with the TotalFail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetTotalFailOk() (*int64, bool) {
	if o == nil || o.TotalFail == nil {
		return nil, false
	}
	return o.TotalFail, true
}

// HasTotalFail returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasTotalFail() bool {
	return o != nil && o.TotalFail != nil
}

// SetTotalFail gets a reference to the given int64 and assigns it to the TotalFail field.
func (o *ScorecardScoreAttributes) SetTotalFail(v int64) {
	o.TotalFail = &v
}

// GetTotalNoData returns the TotalNoData field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetTotalNoData() int64 {
	if o == nil || o.TotalNoData == nil {
		var ret int64
		return ret
	}
	return *o.TotalNoData
}

// GetTotalNoDataOk returns a tuple with the TotalNoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetTotalNoDataOk() (*int64, bool) {
	if o == nil || o.TotalNoData == nil {
		return nil, false
	}
	return o.TotalNoData, true
}

// HasTotalNoData returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasTotalNoData() bool {
	return o != nil && o.TotalNoData != nil
}

// SetTotalNoData gets a reference to the given int64 and assigns it to the TotalNoData field.
func (o *ScorecardScoreAttributes) SetTotalNoData(v int64) {
	o.TotalNoData = &v
}

// GetTotalPass returns the TotalPass field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetTotalPass() int64 {
	if o == nil || o.TotalPass == nil {
		var ret int64
		return ret
	}
	return *o.TotalPass
}

// GetTotalPassOk returns a tuple with the TotalPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetTotalPassOk() (*int64, bool) {
	if o == nil || o.TotalPass == nil {
		return nil, false
	}
	return o.TotalPass, true
}

// HasTotalPass returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasTotalPass() bool {
	return o != nil && o.TotalPass != nil
}

// SetTotalPass gets a reference to the given int64 and assigns it to the TotalPass field.
func (o *ScorecardScoreAttributes) SetTotalPass(v int64) {
	o.TotalPass = &v
}

// GetTotalSkip returns the TotalSkip field value if set, zero value otherwise.
func (o *ScorecardScoreAttributes) GetTotalSkip() int64 {
	if o == nil || o.TotalSkip == nil {
		var ret int64
		return ret
	}
	return *o.TotalSkip
}

// GetTotalSkipOk returns a tuple with the TotalSkip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScorecardScoreAttributes) GetTotalSkipOk() (*int64, bool) {
	if o == nil || o.TotalSkip == nil {
		return nil, false
	}
	return o.TotalSkip, true
}

// HasTotalSkip returns a boolean if a field has been set.
func (o *ScorecardScoreAttributes) HasTotalSkip() bool {
	return o != nil && o.TotalSkip != nil
}

// SetTotalSkip gets a reference to the given int64 and assigns it to the TotalSkip field.
func (o *ScorecardScoreAttributes) SetTotalSkip(v int64) {
	o.TotalSkip = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScorecardScoreAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Denominator != nil {
		toSerialize["denominator"] = o.Denominator
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Numerator != nil {
		toSerialize["numerator"] = o.Numerator
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.TotalEntities != nil {
		toSerialize["total_entities"] = o.TotalEntities
	}
	if o.TotalFail != nil {
		toSerialize["total_fail"] = o.TotalFail
	}
	if o.TotalNoData != nil {
		toSerialize["total_no_data"] = o.TotalNoData
	}
	if o.TotalPass != nil {
		toSerialize["total_pass"] = o.TotalPass
	}
	if o.TotalSkip != nil {
		toSerialize["total_skip"] = o.TotalSkip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScorecardScoreAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation   *ScorecardScoresAggregation `json:"aggregation,omitempty"`
		Denominator   *int64                      `json:"denominator,omitempty"`
		Level         *int64                      `json:"level,omitempty"`
		Numerator     *int64                      `json:"numerator,omitempty"`
		Score         *float64                    `json:"score,omitempty"`
		TotalEntities *int64                      `json:"total_entities,omitempty"`
		TotalFail     *int64                      `json:"total_fail,omitempty"`
		TotalNoData   *int64                      `json:"total_no_data,omitempty"`
		TotalPass     *int64                      `json:"total_pass,omitempty"`
		TotalSkip     *int64                      `json:"total_skip,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "denominator", "level", "numerator", "score", "total_entities", "total_fail", "total_no_data", "total_pass", "total_skip"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregation != nil && !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	o.Denominator = all.Denominator
	o.Level = all.Level
	o.Numerator = all.Numerator
	o.Score = all.Score
	o.TotalEntities = all.TotalEntities
	o.TotalFail = all.TotalFail
	o.TotalNoData = all.TotalNoData
	o.TotalPass = all.TotalPass
	o.TotalSkip = all.TotalSkip

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
