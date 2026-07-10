// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightAuditCompute The aggregation applied to an audit log query.
type GovernanceInsightAuditCompute struct {
	// The aggregation function to apply.
	Aggregation string `json:"aggregation"`
	// The aggregation time window, in milliseconds.
	Interval int64 `json:"interval"`
	// The metric or attribute to aggregate.
	Metric string `json:"metric"`
	// An optional secondary aggregation applied to the audit query result.
	Rollup *string `json:"rollup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightAuditCompute instantiates a new GovernanceInsightAuditCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightAuditCompute(aggregation string, interval int64, metric string) *GovernanceInsightAuditCompute {
	this := GovernanceInsightAuditCompute{}
	this.Aggregation = aggregation
	this.Interval = interval
	this.Metric = metric
	return &this
}

// NewGovernanceInsightAuditComputeWithDefaults instantiates a new GovernanceInsightAuditCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightAuditComputeWithDefaults() *GovernanceInsightAuditCompute {
	this := GovernanceInsightAuditCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *GovernanceInsightAuditCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *GovernanceInsightAuditCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value.
func (o *GovernanceInsightAuditCompute) GetInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditCompute) GetIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *GovernanceInsightAuditCompute) SetInterval(v int64) {
	o.Interval = v
}

// GetMetric returns the Metric field value.
func (o *GovernanceInsightAuditCompute) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditCompute) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *GovernanceInsightAuditCompute) SetMetric(v string) {
	o.Metric = v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *GovernanceInsightAuditCompute) GetRollup() string {
	if o == nil || o.Rollup == nil {
		var ret string
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditCompute) GetRollupOk() (*string, bool) {
	if o == nil || o.Rollup == nil {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *GovernanceInsightAuditCompute) HasRollup() bool {
	return o != nil && o.Rollup != nil
}

// SetRollup gets a reference to the given string and assigns it to the Rollup field.
func (o *GovernanceInsightAuditCompute) SetRollup(v string) {
	o.Rollup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightAuditCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["interval"] = o.Interval
	toSerialize["metric"] = o.Metric
	if o.Rollup != nil {
		toSerialize["rollup"] = o.Rollup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightAuditCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation"`
		Interval    *int64  `json:"interval"`
		Metric      *string `json:"metric"`
		Rollup      *string `json:"rollup,omitempty"`
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
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "interval", "metric", "rollup"})
	} else {
		return err
	}
	o.Aggregation = *all.Aggregation
	o.Interval = *all.Interval
	o.Metric = *all.Metric
	o.Rollup = all.Rollup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
