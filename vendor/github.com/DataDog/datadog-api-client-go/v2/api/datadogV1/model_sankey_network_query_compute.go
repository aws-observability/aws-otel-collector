// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyNetworkQueryCompute Compute aggregation for network queries.
type SankeyNetworkQueryCompute struct {
	// The type of aggregation that can be performed on events-based queries.
	Aggregation EventsAggregation `json:"aggregation"`
	// Metric to aggregate.
	Metric string `json:"metric"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyNetworkQueryCompute instantiates a new SankeyNetworkQueryCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyNetworkQueryCompute(aggregation EventsAggregation, metric string) *SankeyNetworkQueryCompute {
	this := SankeyNetworkQueryCompute{}
	this.Aggregation = aggregation
	this.Metric = metric
	return &this
}

// NewSankeyNetworkQueryComputeWithDefaults instantiates a new SankeyNetworkQueryCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyNetworkQueryComputeWithDefaults() *SankeyNetworkQueryCompute {
	this := SankeyNetworkQueryCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *SankeyNetworkQueryCompute) GetAggregation() EventsAggregation {
	if o == nil {
		var ret EventsAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQueryCompute) GetAggregationOk() (*EventsAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *SankeyNetworkQueryCompute) SetAggregation(v EventsAggregation) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value.
func (o *SankeyNetworkQueryCompute) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQueryCompute) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *SankeyNetworkQueryCompute) SetMetric(v string) {
	o.Metric = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyNetworkQueryCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["metric"] = o.Metric
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyNetworkQueryCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *EventsAggregation `json:"aggregation"`
		Metric      *string            `json:"metric"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	o.Aggregation = *all.Aggregation
	o.Metric = *all.Metric

	return nil
}
