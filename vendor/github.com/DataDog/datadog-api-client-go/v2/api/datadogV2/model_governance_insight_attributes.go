// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightAttributes The attributes of a governance insight.
type GovernanceInsightAttributes struct {
	// An audit log query used to compute an insight value.
	AuditQuery GovernanceInsightAuditQuery `json:"audit_query"`
	// The best practice associated with an insight. Populated with the first active best practice
	// matched to the insight; `null` when no best practice is attached.
	BestPractice GovernanceBestPracticeDefinition `json:"best_practice"`
	// A relative link to the product surface where the insight can be acted upon.
	DeepLink string `json:"deep_link"`
	// A human-readable description of what the insight measures.
	Description string `json:"description"`
	// Human-readable name of the insight.
	DisplayName string `json:"display_name"`
	// An event query used to compute an insight value.
	EventQuery GovernanceInsightEventQuery `json:"event_query"`
	// A metric query used to compute an insight value.
	MetricQuery GovernanceInsightMetricQuery `json:"metric_query"`
	// The value of the insight over the previous comparison window. `null` when values were
	// not requested or could not be computed.
	OldValue datadog.NullableFloat64 `json:"old_value"`
	// A percentage query that computes an insight value as a ratio of two metric queries.
	PercentageQuery GovernanceInsightPercentageQuery `json:"percentage_query"`
	// The product the insight belongs to.
	Product string `json:"product"`
	// Query execution context that allows the frontend to execute insight queries directly.
	QueryConfig *GovernanceInsightQueryConfig `json:"query_config,omitempty"`
	// The relative order in which the insight should be displayed.
	SortOrder *int64 `json:"sort_order,omitempty"`
	// The state of the insight. A `critical` insight receives extra UI treatment to draw
	// attention to it.
	State string `json:"state"`
	// The sub-product the insight belongs to, if any.
	SubProduct string `json:"sub_product"`
	// The time range the insight value is computed over, if applicable.
	TimeRange string `json:"time_range"`
	// The unit that the insight's value is measured in.
	UnitName string `json:"unit_name"`
	// A usage query used to compute an insight value.
	UsageQuery GovernanceInsightUsageQuery `json:"usage_query"`
	// The current value of the insight. `null` when values were not requested or could not be computed.
	Value datadog.NullableFloat64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightAttributes instantiates a new GovernanceInsightAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightAttributes(auditQuery GovernanceInsightAuditQuery, bestPractice GovernanceBestPracticeDefinition, deepLink string, description string, displayName string, eventQuery GovernanceInsightEventQuery, metricQuery GovernanceInsightMetricQuery, oldValue datadog.NullableFloat64, percentageQuery GovernanceInsightPercentageQuery, product string, state string, subProduct string, timeRange string, unitName string, usageQuery GovernanceInsightUsageQuery, value datadog.NullableFloat64) *GovernanceInsightAttributes {
	this := GovernanceInsightAttributes{}
	this.AuditQuery = auditQuery
	this.BestPractice = bestPractice
	this.DeepLink = deepLink
	this.Description = description
	this.DisplayName = displayName
	this.EventQuery = eventQuery
	this.MetricQuery = metricQuery
	this.OldValue = oldValue
	this.PercentageQuery = percentageQuery
	this.Product = product
	this.State = state
	this.SubProduct = subProduct
	this.TimeRange = timeRange
	this.UnitName = unitName
	this.UsageQuery = usageQuery
	this.Value = value
	return &this
}

// NewGovernanceInsightAttributesWithDefaults instantiates a new GovernanceInsightAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightAttributesWithDefaults() *GovernanceInsightAttributes {
	this := GovernanceInsightAttributes{}
	return &this
}

// GetAuditQuery returns the AuditQuery field value.
func (o *GovernanceInsightAttributes) GetAuditQuery() GovernanceInsightAuditQuery {
	if o == nil {
		var ret GovernanceInsightAuditQuery
		return ret
	}
	return o.AuditQuery
}

// GetAuditQueryOk returns a tuple with the AuditQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetAuditQueryOk() (*GovernanceInsightAuditQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditQuery, true
}

// SetAuditQuery sets field value.
func (o *GovernanceInsightAttributes) SetAuditQuery(v GovernanceInsightAuditQuery) {
	o.AuditQuery = v
}

// GetBestPractice returns the BestPractice field value.
func (o *GovernanceInsightAttributes) GetBestPractice() GovernanceBestPracticeDefinition {
	if o == nil {
		var ret GovernanceBestPracticeDefinition
		return ret
	}
	return o.BestPractice
}

// GetBestPracticeOk returns a tuple with the BestPractice field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetBestPracticeOk() (*GovernanceBestPracticeDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BestPractice, true
}

// SetBestPractice sets field value.
func (o *GovernanceInsightAttributes) SetBestPractice(v GovernanceBestPracticeDefinition) {
	o.BestPractice = v
}

// GetDeepLink returns the DeepLink field value.
func (o *GovernanceInsightAttributes) GetDeepLink() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeepLink
}

// GetDeepLinkOk returns a tuple with the DeepLink field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetDeepLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeepLink, true
}

// SetDeepLink sets field value.
func (o *GovernanceInsightAttributes) SetDeepLink(v string) {
	o.DeepLink = v
}

// GetDescription returns the Description field value.
func (o *GovernanceInsightAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceInsightAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *GovernanceInsightAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *GovernanceInsightAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEventQuery returns the EventQuery field value.
func (o *GovernanceInsightAttributes) GetEventQuery() GovernanceInsightEventQuery {
	if o == nil {
		var ret GovernanceInsightEventQuery
		return ret
	}
	return o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetEventQueryOk() (*GovernanceInsightEventQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventQuery, true
}

// SetEventQuery sets field value.
func (o *GovernanceInsightAttributes) SetEventQuery(v GovernanceInsightEventQuery) {
	o.EventQuery = v
}

// GetMetricQuery returns the MetricQuery field value.
func (o *GovernanceInsightAttributes) GetMetricQuery() GovernanceInsightMetricQuery {
	if o == nil {
		var ret GovernanceInsightMetricQuery
		return ret
	}
	return o.MetricQuery
}

// GetMetricQueryOk returns a tuple with the MetricQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetMetricQueryOk() (*GovernanceInsightMetricQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricQuery, true
}

// SetMetricQuery sets field value.
func (o *GovernanceInsightAttributes) SetMetricQuery(v GovernanceInsightMetricQuery) {
	o.MetricQuery = v
}

// GetOldValue returns the OldValue field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *GovernanceInsightAttributes) GetOldValue() float64 {
	if o == nil || o.OldValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceInsightAttributes) GetOldValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// SetOldValue sets field value.
func (o *GovernanceInsightAttributes) SetOldValue(v float64) {
	o.OldValue.Set(&v)
}

// GetPercentageQuery returns the PercentageQuery field value.
func (o *GovernanceInsightAttributes) GetPercentageQuery() GovernanceInsightPercentageQuery {
	if o == nil {
		var ret GovernanceInsightPercentageQuery
		return ret
	}
	return o.PercentageQuery
}

// GetPercentageQueryOk returns a tuple with the PercentageQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetPercentageQueryOk() (*GovernanceInsightPercentageQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentageQuery, true
}

// SetPercentageQuery sets field value.
func (o *GovernanceInsightAttributes) SetPercentageQuery(v GovernanceInsightPercentageQuery) {
	o.PercentageQuery = v
}

// GetProduct returns the Product field value.
func (o *GovernanceInsightAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *GovernanceInsightAttributes) SetProduct(v string) {
	o.Product = v
}

// GetQueryConfig returns the QueryConfig field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetQueryConfig() GovernanceInsightQueryConfig {
	if o == nil || o.QueryConfig == nil {
		var ret GovernanceInsightQueryConfig
		return ret
	}
	return *o.QueryConfig
}

// GetQueryConfigOk returns a tuple with the QueryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetQueryConfigOk() (*GovernanceInsightQueryConfig, bool) {
	if o == nil || o.QueryConfig == nil {
		return nil, false
	}
	return o.QueryConfig, true
}

// HasQueryConfig returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasQueryConfig() bool {
	return o != nil && o.QueryConfig != nil
}

// SetQueryConfig gets a reference to the given GovernanceInsightQueryConfig and assigns it to the QueryConfig field.
func (o *GovernanceInsightAttributes) SetQueryConfig(v GovernanceInsightQueryConfig) {
	o.QueryConfig = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetSortOrder() int64 {
	if o == nil || o.SortOrder == nil {
		var ret int64
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetSortOrderOk() (*int64, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasSortOrder() bool {
	return o != nil && o.SortOrder != nil
}

// SetSortOrder gets a reference to the given int64 and assigns it to the SortOrder field.
func (o *GovernanceInsightAttributes) SetSortOrder(v int64) {
	o.SortOrder = &v
}

// GetState returns the State field value.
func (o *GovernanceInsightAttributes) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *GovernanceInsightAttributes) SetState(v string) {
	o.State = v
}

// GetSubProduct returns the SubProduct field value.
func (o *GovernanceInsightAttributes) GetSubProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubProduct
}

// GetSubProductOk returns a tuple with the SubProduct field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetSubProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubProduct, true
}

// SetSubProduct sets field value.
func (o *GovernanceInsightAttributes) SetSubProduct(v string) {
	o.SubProduct = v
}

// GetTimeRange returns the TimeRange field value.
func (o *GovernanceInsightAttributes) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value.
func (o *GovernanceInsightAttributes) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetUnitName returns the UnitName field value.
func (o *GovernanceInsightAttributes) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value.
func (o *GovernanceInsightAttributes) SetUnitName(v string) {
	o.UnitName = v
}

// GetUsageQuery returns the UsageQuery field value.
func (o *GovernanceInsightAttributes) GetUsageQuery() GovernanceInsightUsageQuery {
	if o == nil {
		var ret GovernanceInsightUsageQuery
		return ret
	}
	return o.UsageQuery
}

// GetUsageQueryOk returns a tuple with the UsageQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetUsageQueryOk() (*GovernanceInsightUsageQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageQuery, true
}

// SetUsageQuery sets field value.
func (o *GovernanceInsightAttributes) SetUsageQuery(v GovernanceInsightUsageQuery) {
	o.UsageQuery = v
}

// GetValue returns the Value field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *GovernanceInsightAttributes) GetValue() float64 {
	if o == nil || o.Value.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceInsightAttributes) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value.
func (o *GovernanceInsightAttributes) SetValue(v float64) {
	o.Value.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["audit_query"] = o.AuditQuery
	toSerialize["best_practice"] = o.BestPractice
	toSerialize["deep_link"] = o.DeepLink
	toSerialize["description"] = o.Description
	toSerialize["display_name"] = o.DisplayName
	toSerialize["event_query"] = o.EventQuery
	toSerialize["metric_query"] = o.MetricQuery
	toSerialize["old_value"] = o.OldValue.Get()
	toSerialize["percentage_query"] = o.PercentageQuery
	toSerialize["product"] = o.Product
	if o.QueryConfig != nil {
		toSerialize["query_config"] = o.QueryConfig
	}
	if o.SortOrder != nil {
		toSerialize["sort_order"] = o.SortOrder
	}
	toSerialize["state"] = o.State
	toSerialize["sub_product"] = o.SubProduct
	toSerialize["time_range"] = o.TimeRange
	toSerialize["unit_name"] = o.UnitName
	toSerialize["usage_query"] = o.UsageQuery
	toSerialize["value"] = o.Value.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuditQuery      *GovernanceInsightAuditQuery      `json:"audit_query"`
		BestPractice    *GovernanceBestPracticeDefinition `json:"best_practice"`
		DeepLink        *string                           `json:"deep_link"`
		Description     *string                           `json:"description"`
		DisplayName     *string                           `json:"display_name"`
		EventQuery      *GovernanceInsightEventQuery      `json:"event_query"`
		MetricQuery     *GovernanceInsightMetricQuery     `json:"metric_query"`
		OldValue        datadog.NullableFloat64           `json:"old_value"`
		PercentageQuery *GovernanceInsightPercentageQuery `json:"percentage_query"`
		Product         *string                           `json:"product"`
		QueryConfig     *GovernanceInsightQueryConfig     `json:"query_config,omitempty"`
		SortOrder       *int64                            `json:"sort_order,omitempty"`
		State           *string                           `json:"state"`
		SubProduct      *string                           `json:"sub_product"`
		TimeRange       *string                           `json:"time_range"`
		UnitName        *string                           `json:"unit_name"`
		UsageQuery      *GovernanceInsightUsageQuery      `json:"usage_query"`
		Value           datadog.NullableFloat64           `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuditQuery == nil {
		return fmt.Errorf("required field audit_query missing")
	}
	if all.BestPractice == nil {
		return fmt.Errorf("required field best_practice missing")
	}
	if all.DeepLink == nil {
		return fmt.Errorf("required field deep_link missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.EventQuery == nil {
		return fmt.Errorf("required field event_query missing")
	}
	if all.MetricQuery == nil {
		return fmt.Errorf("required field metric_query missing")
	}
	if !all.OldValue.IsSet() {
		return fmt.Errorf("required field old_value missing")
	}
	if all.PercentageQuery == nil {
		return fmt.Errorf("required field percentage_query missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.SubProduct == nil {
		return fmt.Errorf("required field sub_product missing")
	}
	if all.TimeRange == nil {
		return fmt.Errorf("required field time_range missing")
	}
	if all.UnitName == nil {
		return fmt.Errorf("required field unit_name missing")
	}
	if all.UsageQuery == nil {
		return fmt.Errorf("required field usage_query missing")
	}
	if !all.Value.IsSet() {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audit_query", "best_practice", "deep_link", "description", "display_name", "event_query", "metric_query", "old_value", "percentage_query", "product", "query_config", "sort_order", "state", "sub_product", "time_range", "unit_name", "usage_query", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuditQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuditQuery = *all.AuditQuery
	if all.BestPractice.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BestPractice = *all.BestPractice
	o.DeepLink = *all.DeepLink
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	if all.EventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EventQuery = *all.EventQuery
	if all.MetricQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricQuery = *all.MetricQuery
	o.OldValue = all.OldValue
	if all.PercentageQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PercentageQuery = *all.PercentageQuery
	o.Product = *all.Product
	if all.QueryConfig != nil && all.QueryConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryConfig = all.QueryConfig
	o.SortOrder = all.SortOrder
	o.State = *all.State
	o.SubProduct = *all.SubProduct
	o.TimeRange = *all.TimeRange
	o.UnitName = *all.UnitName
	if all.UsageQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UsageQuery = *all.UsageQuery
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
