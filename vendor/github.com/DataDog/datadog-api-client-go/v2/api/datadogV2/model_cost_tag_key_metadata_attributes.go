// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeyMetadataAttributes Attributes of a Cloud Cost Management tag key metadata entry.
type CostTagKeyMetadataAttributes struct {
	// Number of unique tag values observed for this tag key, keyed by cloud account ID.
	CardinalityByAccount map[string]int64 `json:"cardinality_by_account"`
	// Total cost (in the report currency) of cost line items that carry this tag key for the requested period.
	CostCovered float64 `json:"cost_covered"`
	// The day this row corresponds to, in `YYYY-MM-DD` format. Present only when `filter[daily]=true`; omitted for the monthly roll-up returned by default.
	Date *string `json:"date,omitempty"`
	// The Cloud Cost Management metric this row aggregates, for example `aws.cost.net.amortized`.
	Metric string `json:"metric"`
	// Number of cost rows that carry this tag key over the requested period.
	RowCount int64 `json:"row_count"`
	// Origins where this tag key was observed (for example, `aws-user-defined`).
	TagSources []string `json:"tag_sources"`
	// A sample of the most frequent tag values observed for this tag key, keyed by cloud account ID.
	TopValuesByAccount map[string][]string `json:"top_values_by_account"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagKeyMetadataAttributes instantiates a new CostTagKeyMetadataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagKeyMetadataAttributes(cardinalityByAccount map[string]int64, costCovered float64, metric string, rowCount int64, tagSources []string, topValuesByAccount map[string][]string) *CostTagKeyMetadataAttributes {
	this := CostTagKeyMetadataAttributes{}
	this.CardinalityByAccount = cardinalityByAccount
	this.CostCovered = costCovered
	this.Metric = metric
	this.RowCount = rowCount
	this.TagSources = tagSources
	this.TopValuesByAccount = topValuesByAccount
	return &this
}

// NewCostTagKeyMetadataAttributesWithDefaults instantiates a new CostTagKeyMetadataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagKeyMetadataAttributesWithDefaults() *CostTagKeyMetadataAttributes {
	this := CostTagKeyMetadataAttributes{}
	return &this
}

// GetCardinalityByAccount returns the CardinalityByAccount field value.
func (o *CostTagKeyMetadataAttributes) GetCardinalityByAccount() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.CardinalityByAccount
}

// GetCardinalityByAccountOk returns a tuple with the CardinalityByAccount field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetCardinalityByAccountOk() (*map[string]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardinalityByAccount, true
}

// SetCardinalityByAccount sets field value.
func (o *CostTagKeyMetadataAttributes) SetCardinalityByAccount(v map[string]int64) {
	o.CardinalityByAccount = v
}

// GetCostCovered returns the CostCovered field value.
func (o *CostTagKeyMetadataAttributes) GetCostCovered() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CostCovered
}

// GetCostCoveredOk returns a tuple with the CostCovered field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetCostCoveredOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostCovered, true
}

// SetCostCovered sets field value.
func (o *CostTagKeyMetadataAttributes) SetCostCovered(v float64) {
	o.CostCovered = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CostTagKeyMetadataAttributes) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CostTagKeyMetadataAttributes) HasDate() bool {
	return o != nil && o.Date != nil
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *CostTagKeyMetadataAttributes) SetDate(v string) {
	o.Date = &v
}

// GetMetric returns the Metric field value.
func (o *CostTagKeyMetadataAttributes) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *CostTagKeyMetadataAttributes) SetMetric(v string) {
	o.Metric = v
}

// GetRowCount returns the RowCount field value.
func (o *CostTagKeyMetadataAttributes) GetRowCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetRowCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowCount, true
}

// SetRowCount sets field value.
func (o *CostTagKeyMetadataAttributes) SetRowCount(v int64) {
	o.RowCount = v
}

// GetTagSources returns the TagSources field value.
func (o *CostTagKeyMetadataAttributes) GetTagSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagSources
}

// GetTagSourcesOk returns a tuple with the TagSources field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetTagSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagSources, true
}

// SetTagSources sets field value.
func (o *CostTagKeyMetadataAttributes) SetTagSources(v []string) {
	o.TagSources = v
}

// GetTopValuesByAccount returns the TopValuesByAccount field value.
func (o *CostTagKeyMetadataAttributes) GetTopValuesByAccount() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.TopValuesByAccount
}

// GetTopValuesByAccountOk returns a tuple with the TopValuesByAccount field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyMetadataAttributes) GetTopValuesByAccountOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopValuesByAccount, true
}

// SetTopValuesByAccount sets field value.
func (o *CostTagKeyMetadataAttributes) SetTopValuesByAccount(v map[string][]string) {
	o.TopValuesByAccount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagKeyMetadataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cardinality_by_account"] = o.CardinalityByAccount
	toSerialize["cost_covered"] = o.CostCovered
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	toSerialize["metric"] = o.Metric
	toSerialize["row_count"] = o.RowCount
	toSerialize["tag_sources"] = o.TagSources
	toSerialize["top_values_by_account"] = o.TopValuesByAccount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagKeyMetadataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CardinalityByAccount *map[string]int64    `json:"cardinality_by_account"`
		CostCovered          *float64             `json:"cost_covered"`
		Date                 *string              `json:"date,omitempty"`
		Metric               *string              `json:"metric"`
		RowCount             *int64               `json:"row_count"`
		TagSources           *[]string            `json:"tag_sources"`
		TopValuesByAccount   *map[string][]string `json:"top_values_by_account"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CardinalityByAccount == nil {
		return fmt.Errorf("required field cardinality_by_account missing")
	}
	if all.CostCovered == nil {
		return fmt.Errorf("required field cost_covered missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	if all.RowCount == nil {
		return fmt.Errorf("required field row_count missing")
	}
	if all.TagSources == nil {
		return fmt.Errorf("required field tag_sources missing")
	}
	if all.TopValuesByAccount == nil {
		return fmt.Errorf("required field top_values_by_account missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cardinality_by_account", "cost_covered", "date", "metric", "row_count", "tag_sources", "top_values_by_account"})
	} else {
		return err
	}
	o.CardinalityByAccount = *all.CardinalityByAccount
	o.CostCovered = *all.CostCovered
	o.Date = all.Date
	o.Metric = *all.Metric
	o.RowCount = *all.RowCount
	o.TagSources = *all.TagSources
	o.TopValuesByAccount = *all.TopValuesByAccount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
