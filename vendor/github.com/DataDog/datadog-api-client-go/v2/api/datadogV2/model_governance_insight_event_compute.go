// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightEventCompute The aggregation applied to an event query.
type GovernanceInsightEventCompute struct {
	// The aggregation function to apply.
	Aggregation string `json:"aggregation"`
	// The aggregation time window, in milliseconds.
	Interval int64 `json:"interval"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightEventCompute instantiates a new GovernanceInsightEventCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightEventCompute(aggregation string, interval int64) *GovernanceInsightEventCompute {
	this := GovernanceInsightEventCompute{}
	this.Aggregation = aggregation
	this.Interval = interval
	return &this
}

// NewGovernanceInsightEventComputeWithDefaults instantiates a new GovernanceInsightEventCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightEventComputeWithDefaults() *GovernanceInsightEventCompute {
	this := GovernanceInsightEventCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *GovernanceInsightEventCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightEventCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *GovernanceInsightEventCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value.
func (o *GovernanceInsightEventCompute) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightEventCompute) GetIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *GovernanceInsightEventCompute) SetInterval(v int64) {
	o.Interval = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightEventCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["interval"] = o.Interval

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightEventCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation"`
		Interval    *int64  `json:"interval"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Interval == nil {
		return fmt.Errorf("required field interval missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "interval"})
	} else {
		return err
	}
	o.Aggregation = *all.Aggregation
	o.Interval = *all.Interval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
