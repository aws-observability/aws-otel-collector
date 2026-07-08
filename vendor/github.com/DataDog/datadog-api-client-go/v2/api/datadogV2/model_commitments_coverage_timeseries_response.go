// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageTimeseriesResponse Response containing timeseries coverage metrics for cloud commitment programs.
type CommitmentsCoverageTimeseriesResponse struct {
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	Cost CommitmentsTimeseriesMetric `json:"cost"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	Hours CommitmentsTimeseriesMetric `json:"hours"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsCoverageTimeseriesResponse instantiates a new CommitmentsCoverageTimeseriesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsCoverageTimeseriesResponse(cost CommitmentsTimeseriesMetric, hours CommitmentsTimeseriesMetric) *CommitmentsCoverageTimeseriesResponse {
	this := CommitmentsCoverageTimeseriesResponse{}
	this.Cost = cost
	this.Hours = hours
	return &this
}

// NewCommitmentsCoverageTimeseriesResponseWithDefaults instantiates a new CommitmentsCoverageTimeseriesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsCoverageTimeseriesResponseWithDefaults() *CommitmentsCoverageTimeseriesResponse {
	this := CommitmentsCoverageTimeseriesResponse{}
	return &this
}

// GetCost returns the Cost field value.
func (o *CommitmentsCoverageTimeseriesResponse) GetCost() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesResponse) GetCostOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value.
func (o *CommitmentsCoverageTimeseriesResponse) SetCost(v CommitmentsTimeseriesMetric) {
	o.Cost = v
}

// GetHours returns the Hours field value.
func (o *CommitmentsCoverageTimeseriesResponse) GetHours() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesResponse) GetHoursOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value.
func (o *CommitmentsCoverageTimeseriesResponse) SetHours(v CommitmentsTimeseriesMetric) {
	o.Hours = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsCoverageTimeseriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cost"] = o.Cost
	toSerialize["hours"] = o.Hours

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsCoverageTimeseriesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cost  *CommitmentsTimeseriesMetric `json:"cost"`
		Hours *CommitmentsTimeseriesMetric `json:"hours"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cost == nil {
		return fmt.Errorf("required field cost missing")
	}
	if all.Hours == nil {
		return fmt.Errorf("required field hours missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cost", "hours"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cost = *all.Cost
	if all.Hours.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Hours = *all.Hours

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
