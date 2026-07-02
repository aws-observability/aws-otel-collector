// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAnomaly A single detected Cloud Cost Management anomaly.
type CostAnomaly struct {
	// Actual cost incurred during the anomaly window.
	ActualCost float64 `json:"actual_cost"`
	// Anomalous cost change relative to the expected baseline.
	AnomalousCostChange float64 `json:"anomalous_cost_change"`
	// Anomaly end timestamp in Unix milliseconds.
	AnomalyEnd int64 `json:"anomaly_end"`
	// Anomaly start timestamp in Unix milliseconds.
	AnomalyStart int64 `json:"anomaly_start"`
	// Map of correlated tag keys to the list of correlated tag values.
	CorrelatedTags map[string][]string `json:"correlated_tags"`
	// Map of cost dimension keys to their values for the anomaly grouping.
	Dimensions map[string]string `json:"dimensions"`
	// Resolution metadata for an anomaly that has been dismissed.
	Dismissal *CostAnomalyDismissal `json:"dismissal,omitempty"`
	// Maximum cost observed during the anomaly window.
	MaxCost float64 `json:"max_cost"`
	// Cloud or SaaS provider associated with the anomaly (for example `aws`, `gcp`, `azure`).
	Provider string `json:"provider"`
	// The metrics query that detected the anomaly.
	Query string `json:"query"`
	// The unique identifier of the anomaly.
	Uuid string `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostAnomaly instantiates a new CostAnomaly object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostAnomaly(actualCost float64, anomalousCostChange float64, anomalyEnd int64, anomalyStart int64, correlatedTags map[string][]string, dimensions map[string]string, maxCost float64, provider string, query string, uuid string) *CostAnomaly {
	this := CostAnomaly{}
	this.ActualCost = actualCost
	this.AnomalousCostChange = anomalousCostChange
	this.AnomalyEnd = anomalyEnd
	this.AnomalyStart = anomalyStart
	this.CorrelatedTags = correlatedTags
	this.Dimensions = dimensions
	this.MaxCost = maxCost
	this.Provider = provider
	this.Query = query
	this.Uuid = uuid
	return &this
}

// NewCostAnomalyWithDefaults instantiates a new CostAnomaly object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostAnomalyWithDefaults() *CostAnomaly {
	this := CostAnomaly{}
	return &this
}

// GetActualCost returns the ActualCost field value.
func (o *CostAnomaly) GetActualCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ActualCost
}

// GetActualCostOk returns a tuple with the ActualCost field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetActualCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualCost, true
}

// SetActualCost sets field value.
func (o *CostAnomaly) SetActualCost(v float64) {
	o.ActualCost = v
}

// GetAnomalousCostChange returns the AnomalousCostChange field value.
func (o *CostAnomaly) GetAnomalousCostChange() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AnomalousCostChange
}

// GetAnomalousCostChangeOk returns a tuple with the AnomalousCostChange field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetAnomalousCostChangeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnomalousCostChange, true
}

// SetAnomalousCostChange sets field value.
func (o *CostAnomaly) SetAnomalousCostChange(v float64) {
	o.AnomalousCostChange = v
}

// GetAnomalyEnd returns the AnomalyEnd field value.
func (o *CostAnomaly) GetAnomalyEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AnomalyEnd
}

// GetAnomalyEndOk returns a tuple with the AnomalyEnd field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetAnomalyEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnomalyEnd, true
}

// SetAnomalyEnd sets field value.
func (o *CostAnomaly) SetAnomalyEnd(v int64) {
	o.AnomalyEnd = v
}

// GetAnomalyStart returns the AnomalyStart field value.
func (o *CostAnomaly) GetAnomalyStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AnomalyStart
}

// GetAnomalyStartOk returns a tuple with the AnomalyStart field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetAnomalyStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnomalyStart, true
}

// SetAnomalyStart sets field value.
func (o *CostAnomaly) SetAnomalyStart(v int64) {
	o.AnomalyStart = v
}

// GetCorrelatedTags returns the CorrelatedTags field value.
// If the value is explicit nil, the zero value for map[string][]string will be returned.
func (o *CostAnomaly) GetCorrelatedTags() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.CorrelatedTags
}

// GetCorrelatedTagsOk returns a tuple with the CorrelatedTags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CostAnomaly) GetCorrelatedTagsOk() (*map[string][]string, bool) {
	if o == nil || o.CorrelatedTags == nil {
		return nil, false
	}
	return &o.CorrelatedTags, true
}

// SetCorrelatedTags sets field value.
func (o *CostAnomaly) SetCorrelatedTags(v map[string][]string) {
	o.CorrelatedTags = v
}

// GetDimensions returns the Dimensions field value.
func (o *CostAnomaly) GetDimensions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetDimensionsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *CostAnomaly) SetDimensions(v map[string]string) {
	o.Dimensions = v
}

// GetDismissal returns the Dismissal field value if set, zero value otherwise.
func (o *CostAnomaly) GetDismissal() CostAnomalyDismissal {
	if o == nil || o.Dismissal == nil {
		var ret CostAnomalyDismissal
		return ret
	}
	return *o.Dismissal
}

// GetDismissalOk returns a tuple with the Dismissal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetDismissalOk() (*CostAnomalyDismissal, bool) {
	if o == nil || o.Dismissal == nil {
		return nil, false
	}
	return o.Dismissal, true
}

// HasDismissal returns a boolean if a field has been set.
func (o *CostAnomaly) HasDismissal() bool {
	return o != nil && o.Dismissal != nil
}

// SetDismissal gets a reference to the given CostAnomalyDismissal and assigns it to the Dismissal field.
func (o *CostAnomaly) SetDismissal(v CostAnomalyDismissal) {
	o.Dismissal = &v
}

// GetMaxCost returns the MaxCost field value.
func (o *CostAnomaly) GetMaxCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxCost
}

// GetMaxCostOk returns a tuple with the MaxCost field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetMaxCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCost, true
}

// SetMaxCost sets field value.
func (o *CostAnomaly) SetMaxCost(v float64) {
	o.MaxCost = v
}

// GetProvider returns the Provider field value.
func (o *CostAnomaly) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *CostAnomaly) SetProvider(v string) {
	o.Provider = v
}

// GetQuery returns the Query field value.
func (o *CostAnomaly) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *CostAnomaly) SetQuery(v string) {
	o.Query = v
}

// GetUuid returns the Uuid field value.
func (o *CostAnomaly) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CostAnomaly) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *CostAnomaly) SetUuid(v string) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostAnomaly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actual_cost"] = o.ActualCost
	toSerialize["anomalous_cost_change"] = o.AnomalousCostChange
	toSerialize["anomaly_end"] = o.AnomalyEnd
	toSerialize["anomaly_start"] = o.AnomalyStart
	if o.CorrelatedTags != nil {
		toSerialize["correlated_tags"] = o.CorrelatedTags
	}
	toSerialize["dimensions"] = o.Dimensions
	if o.Dismissal != nil {
		toSerialize["dismissal"] = o.Dismissal
	}
	toSerialize["max_cost"] = o.MaxCost
	toSerialize["provider"] = o.Provider
	toSerialize["query"] = o.Query
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostAnomaly) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActualCost          *float64              `json:"actual_cost"`
		AnomalousCostChange *float64              `json:"anomalous_cost_change"`
		AnomalyEnd          *int64                `json:"anomaly_end"`
		AnomalyStart        *int64                `json:"anomaly_start"`
		CorrelatedTags      map[string][]string   `json:"correlated_tags"`
		Dimensions          *map[string]string    `json:"dimensions"`
		Dismissal           *CostAnomalyDismissal `json:"dismissal,omitempty"`
		MaxCost             *float64              `json:"max_cost"`
		Provider            *string               `json:"provider"`
		Query               *string               `json:"query"`
		Uuid                *string               `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActualCost == nil {
		return fmt.Errorf("required field actual_cost missing")
	}
	if all.AnomalousCostChange == nil {
		return fmt.Errorf("required field anomalous_cost_change missing")
	}
	if all.AnomalyEnd == nil {
		return fmt.Errorf("required field anomaly_end missing")
	}
	if all.AnomalyStart == nil {
		return fmt.Errorf("required field anomaly_start missing")
	}
	if all.CorrelatedTags == nil {
		return fmt.Errorf("required field correlated_tags missing")
	}
	if all.Dimensions == nil {
		return fmt.Errorf("required field dimensions missing")
	}
	if all.MaxCost == nil {
		return fmt.Errorf("required field max_cost missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actual_cost", "anomalous_cost_change", "anomaly_end", "anomaly_start", "correlated_tags", "dimensions", "dismissal", "max_cost", "provider", "query", "uuid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActualCost = *all.ActualCost
	o.AnomalousCostChange = *all.AnomalousCostChange
	o.AnomalyEnd = *all.AnomalyEnd
	o.AnomalyStart = *all.AnomalyStart
	o.CorrelatedTags = all.CorrelatedTags
	o.Dimensions = *all.Dimensions
	if all.Dismissal != nil && all.Dismissal.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dismissal = all.Dismissal
	o.MaxCost = *all.MaxCost
	o.Provider = *all.Provider
	o.Query = *all.Query
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
