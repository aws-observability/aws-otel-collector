// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsSavingsTimeseriesResponse Response containing timeseries savings metrics for cloud commitment programs.
type CommitmentsSavingsTimeseriesResponse struct {
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	ActualCost CommitmentsTimeseriesMetric `json:"actual_cost"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	EffectiveSavingsRate CommitmentsTimeseriesMetric `json:"effective_savings_rate"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	OnDemandEquivalentCost CommitmentsTimeseriesMetric `json:"on_demand_equivalent_cost"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	RealizedSavings CommitmentsTimeseriesMetric `json:"realized_savings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsSavingsTimeseriesResponse instantiates a new CommitmentsSavingsTimeseriesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsSavingsTimeseriesResponse(actualCost CommitmentsTimeseriesMetric, effectiveSavingsRate CommitmentsTimeseriesMetric, onDemandEquivalentCost CommitmentsTimeseriesMetric, realizedSavings CommitmentsTimeseriesMetric) *CommitmentsSavingsTimeseriesResponse {
	this := CommitmentsSavingsTimeseriesResponse{}
	this.ActualCost = actualCost
	this.EffectiveSavingsRate = effectiveSavingsRate
	this.OnDemandEquivalentCost = onDemandEquivalentCost
	this.RealizedSavings = realizedSavings
	return &this
}

// NewCommitmentsSavingsTimeseriesResponseWithDefaults instantiates a new CommitmentsSavingsTimeseriesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsSavingsTimeseriesResponseWithDefaults() *CommitmentsSavingsTimeseriesResponse {
	this := CommitmentsSavingsTimeseriesResponse{}
	return &this
}

// GetActualCost returns the ActualCost field value.
func (o *CommitmentsSavingsTimeseriesResponse) GetActualCost() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.ActualCost
}

// GetActualCostOk returns a tuple with the ActualCost field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponse) GetActualCostOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualCost, true
}

// SetActualCost sets field value.
func (o *CommitmentsSavingsTimeseriesResponse) SetActualCost(v CommitmentsTimeseriesMetric) {
	o.ActualCost = v
}

// GetEffectiveSavingsRate returns the EffectiveSavingsRate field value.
func (o *CommitmentsSavingsTimeseriesResponse) GetEffectiveSavingsRate() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.EffectiveSavingsRate
}

// GetEffectiveSavingsRateOk returns a tuple with the EffectiveSavingsRate field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponse) GetEffectiveSavingsRateOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveSavingsRate, true
}

// SetEffectiveSavingsRate sets field value.
func (o *CommitmentsSavingsTimeseriesResponse) SetEffectiveSavingsRate(v CommitmentsTimeseriesMetric) {
	o.EffectiveSavingsRate = v
}

// GetOnDemandEquivalentCost returns the OnDemandEquivalentCost field value.
func (o *CommitmentsSavingsTimeseriesResponse) GetOnDemandEquivalentCost() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.OnDemandEquivalentCost
}

// GetOnDemandEquivalentCostOk returns a tuple with the OnDemandEquivalentCost field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponse) GetOnDemandEquivalentCostOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnDemandEquivalentCost, true
}

// SetOnDemandEquivalentCost sets field value.
func (o *CommitmentsSavingsTimeseriesResponse) SetOnDemandEquivalentCost(v CommitmentsTimeseriesMetric) {
	o.OnDemandEquivalentCost = v
}

// GetRealizedSavings returns the RealizedSavings field value.
func (o *CommitmentsSavingsTimeseriesResponse) GetRealizedSavings() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.RealizedSavings
}

// GetRealizedSavingsOk returns a tuple with the RealizedSavings field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponse) GetRealizedSavingsOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealizedSavings, true
}

// SetRealizedSavings sets field value.
func (o *CommitmentsSavingsTimeseriesResponse) SetRealizedSavings(v CommitmentsTimeseriesMetric) {
	o.RealizedSavings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsSavingsTimeseriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actual_cost"] = o.ActualCost
	toSerialize["effective_savings_rate"] = o.EffectiveSavingsRate
	toSerialize["on_demand_equivalent_cost"] = o.OnDemandEquivalentCost
	toSerialize["realized_savings"] = o.RealizedSavings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsSavingsTimeseriesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActualCost             *CommitmentsTimeseriesMetric `json:"actual_cost"`
		EffectiveSavingsRate   *CommitmentsTimeseriesMetric `json:"effective_savings_rate"`
		OnDemandEquivalentCost *CommitmentsTimeseriesMetric `json:"on_demand_equivalent_cost"`
		RealizedSavings        *CommitmentsTimeseriesMetric `json:"realized_savings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActualCost == nil {
		return fmt.Errorf("required field actual_cost missing")
	}
	if all.EffectiveSavingsRate == nil {
		return fmt.Errorf("required field effective_savings_rate missing")
	}
	if all.OnDemandEquivalentCost == nil {
		return fmt.Errorf("required field on_demand_equivalent_cost missing")
	}
	if all.RealizedSavings == nil {
		return fmt.Errorf("required field realized_savings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actual_cost", "effective_savings_rate", "on_demand_equivalent_cost", "realized_savings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ActualCost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActualCost = *all.ActualCost
	if all.EffectiveSavingsRate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EffectiveSavingsRate = *all.EffectiveSavingsRate
	if all.OnDemandEquivalentCost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OnDemandEquivalentCost = *all.OnDemandEquivalentCost
	if all.RealizedSavings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RealizedSavings = *all.RealizedSavings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
