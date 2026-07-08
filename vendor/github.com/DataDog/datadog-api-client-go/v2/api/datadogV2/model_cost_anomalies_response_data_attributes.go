// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAnomaliesResponseDataAttributes Cost anomaly results and aggregated totals for the queried window.
type CostAnomaliesResponseDataAttributes struct {
	// The list of cost anomalies that match the request.
	Anomalies []CostAnomaly `json:"anomalies"`
	// Average daily anomalous cost change across the queried window.
	AvgDailyAnomalousCost float64 `json:"avg_daily_anomalous_cost"`
	// Total actual cost spent across the queried window for the matching providers.
	TotalActualCost float64 `json:"total_actual_cost"`
	// Sum of the anomalous cost change across all returned anomalies.
	TotalAnomalousCost float64 `json:"total_anomalous_cost"`
	// Total number of anomalies that match the request.
	TotalCount int64 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostAnomaliesResponseDataAttributes instantiates a new CostAnomaliesResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostAnomaliesResponseDataAttributes(anomalies []CostAnomaly, avgDailyAnomalousCost float64, totalActualCost float64, totalAnomalousCost float64, totalCount int64) *CostAnomaliesResponseDataAttributes {
	this := CostAnomaliesResponseDataAttributes{}
	this.Anomalies = anomalies
	this.AvgDailyAnomalousCost = avgDailyAnomalousCost
	this.TotalActualCost = totalActualCost
	this.TotalAnomalousCost = totalAnomalousCost
	this.TotalCount = totalCount
	return &this
}

// NewCostAnomaliesResponseDataAttributesWithDefaults instantiates a new CostAnomaliesResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostAnomaliesResponseDataAttributesWithDefaults() *CostAnomaliesResponseDataAttributes {
	this := CostAnomaliesResponseDataAttributes{}
	return &this
}

// GetAnomalies returns the Anomalies field value.
func (o *CostAnomaliesResponseDataAttributes) GetAnomalies() []CostAnomaly {
	if o == nil {
		var ret []CostAnomaly
		return ret
	}
	return o.Anomalies
}

// GetAnomaliesOk returns a tuple with the Anomalies field value
// and a boolean to check if the value has been set.
func (o *CostAnomaliesResponseDataAttributes) GetAnomaliesOk() (*[]CostAnomaly, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anomalies, true
}

// SetAnomalies sets field value.
func (o *CostAnomaliesResponseDataAttributes) SetAnomalies(v []CostAnomaly) {
	o.Anomalies = v
}

// GetAvgDailyAnomalousCost returns the AvgDailyAnomalousCost field value.
func (o *CostAnomaliesResponseDataAttributes) GetAvgDailyAnomalousCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgDailyAnomalousCost
}

// GetAvgDailyAnomalousCostOk returns a tuple with the AvgDailyAnomalousCost field value
// and a boolean to check if the value has been set.
func (o *CostAnomaliesResponseDataAttributes) GetAvgDailyAnomalousCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDailyAnomalousCost, true
}

// SetAvgDailyAnomalousCost sets field value.
func (o *CostAnomaliesResponseDataAttributes) SetAvgDailyAnomalousCost(v float64) {
	o.AvgDailyAnomalousCost = v
}

// GetTotalActualCost returns the TotalActualCost field value.
func (o *CostAnomaliesResponseDataAttributes) GetTotalActualCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalActualCost
}

// GetTotalActualCostOk returns a tuple with the TotalActualCost field value
// and a boolean to check if the value has been set.
func (o *CostAnomaliesResponseDataAttributes) GetTotalActualCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalActualCost, true
}

// SetTotalActualCost sets field value.
func (o *CostAnomaliesResponseDataAttributes) SetTotalActualCost(v float64) {
	o.TotalActualCost = v
}

// GetTotalAnomalousCost returns the TotalAnomalousCost field value.
func (o *CostAnomaliesResponseDataAttributes) GetTotalAnomalousCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalAnomalousCost
}

// GetTotalAnomalousCostOk returns a tuple with the TotalAnomalousCost field value
// and a boolean to check if the value has been set.
func (o *CostAnomaliesResponseDataAttributes) GetTotalAnomalousCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAnomalousCost, true
}

// SetTotalAnomalousCost sets field value.
func (o *CostAnomaliesResponseDataAttributes) SetTotalAnomalousCost(v float64) {
	o.TotalAnomalousCost = v
}

// GetTotalCount returns the TotalCount field value.
func (o *CostAnomaliesResponseDataAttributes) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *CostAnomaliesResponseDataAttributes) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *CostAnomaliesResponseDataAttributes) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostAnomaliesResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["anomalies"] = o.Anomalies
	toSerialize["avg_daily_anomalous_cost"] = o.AvgDailyAnomalousCost
	toSerialize["total_actual_cost"] = o.TotalActualCost
	toSerialize["total_anomalous_cost"] = o.TotalAnomalousCost
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostAnomaliesResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Anomalies             *[]CostAnomaly `json:"anomalies"`
		AvgDailyAnomalousCost *float64       `json:"avg_daily_anomalous_cost"`
		TotalActualCost       *float64       `json:"total_actual_cost"`
		TotalAnomalousCost    *float64       `json:"total_anomalous_cost"`
		TotalCount            *int64         `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Anomalies == nil {
		return fmt.Errorf("required field anomalies missing")
	}
	if all.AvgDailyAnomalousCost == nil {
		return fmt.Errorf("required field avg_daily_anomalous_cost missing")
	}
	if all.TotalActualCost == nil {
		return fmt.Errorf("required field total_actual_cost missing")
	}
	if all.TotalAnomalousCost == nil {
		return fmt.Errorf("required field total_anomalous_cost missing")
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anomalies", "avg_daily_anomalous_cost", "total_actual_cost", "total_anomalous_cost", "total_count"})
	} else {
		return err
	}
	o.Anomalies = *all.Anomalies
	o.AvgDailyAnomalousCost = *all.AvgDailyAnomalousCost
	o.TotalActualCost = *all.TotalActualCost
	o.TotalAnomalousCost = *all.TotalAnomalousCost
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
