// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityMonitorOptions Monitor configuration options for data quality queries.
type MonitorFormulaAndFunctionDataQualityMonitorOptions struct {
	// Crontab expression to override the default schedule.
	CrontabOverride *string `json:"crontab_override,omitempty"`
	// Custom SQL query for the monitor.
	CustomSql *string `json:"custom_sql,omitempty"`
	// Custom WHERE clause for the query.
	CustomWhere *string `json:"custom_where,omitempty"`
	// Columns to group results by.
	GroupByColumns []string `json:"group_by_columns,omitempty"`
	// Override for the model type used in anomaly detection.
	ModelTypeOverride *MonitorFormulaAndFunctionDataQualityModelTypeOverride `json:"model_type_override,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualityMonitorOptions instantiates a new MonitorFormulaAndFunctionDataQualityMonitorOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualityMonitorOptions() *MonitorFormulaAndFunctionDataQualityMonitorOptions {
	this := MonitorFormulaAndFunctionDataQualityMonitorOptions{}
	return &this
}

// NewMonitorFormulaAndFunctionDataQualityMonitorOptionsWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualityMonitorOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualityMonitorOptionsWithDefaults() *MonitorFormulaAndFunctionDataQualityMonitorOptions {
	this := MonitorFormulaAndFunctionDataQualityMonitorOptions{}
	return &this
}

// GetCrontabOverride returns the CrontabOverride field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetCrontabOverride() string {
	if o == nil || o.CrontabOverride == nil {
		var ret string
		return ret
	}
	return *o.CrontabOverride
}

// GetCrontabOverrideOk returns a tuple with the CrontabOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetCrontabOverrideOk() (*string, bool) {
	if o == nil || o.CrontabOverride == nil {
		return nil, false
	}
	return o.CrontabOverride, true
}

// HasCrontabOverride returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) HasCrontabOverride() bool {
	return o != nil && o.CrontabOverride != nil
}

// SetCrontabOverride gets a reference to the given string and assigns it to the CrontabOverride field.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) SetCrontabOverride(v string) {
	o.CrontabOverride = &v
}

// GetCustomSql returns the CustomSql field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetCustomSql() string {
	if o == nil || o.CustomSql == nil {
		var ret string
		return ret
	}
	return *o.CustomSql
}

// GetCustomSqlOk returns a tuple with the CustomSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetCustomSqlOk() (*string, bool) {
	if o == nil || o.CustomSql == nil {
		return nil, false
	}
	return o.CustomSql, true
}

// HasCustomSql returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) HasCustomSql() bool {
	return o != nil && o.CustomSql != nil
}

// SetCustomSql gets a reference to the given string and assigns it to the CustomSql field.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) SetCustomSql(v string) {
	o.CustomSql = &v
}

// GetCustomWhere returns the CustomWhere field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetCustomWhere() string {
	if o == nil || o.CustomWhere == nil {
		var ret string
		return ret
	}
	return *o.CustomWhere
}

// GetCustomWhereOk returns a tuple with the CustomWhere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetCustomWhereOk() (*string, bool) {
	if o == nil || o.CustomWhere == nil {
		return nil, false
	}
	return o.CustomWhere, true
}

// HasCustomWhere returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) HasCustomWhere() bool {
	return o != nil && o.CustomWhere != nil
}

// SetCustomWhere gets a reference to the given string and assigns it to the CustomWhere field.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) SetCustomWhere(v string) {
	o.CustomWhere = &v
}

// GetGroupByColumns returns the GroupByColumns field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetGroupByColumns() []string {
	if o == nil || o.GroupByColumns == nil {
		var ret []string
		return ret
	}
	return o.GroupByColumns
}

// GetGroupByColumnsOk returns a tuple with the GroupByColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetGroupByColumnsOk() (*[]string, bool) {
	if o == nil || o.GroupByColumns == nil {
		return nil, false
	}
	return &o.GroupByColumns, true
}

// HasGroupByColumns returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) HasGroupByColumns() bool {
	return o != nil && o.GroupByColumns != nil
}

// SetGroupByColumns gets a reference to the given []string and assigns it to the GroupByColumns field.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) SetGroupByColumns(v []string) {
	o.GroupByColumns = v
}

// GetModelTypeOverride returns the ModelTypeOverride field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetModelTypeOverride() MonitorFormulaAndFunctionDataQualityModelTypeOverride {
	if o == nil || o.ModelTypeOverride == nil {
		var ret MonitorFormulaAndFunctionDataQualityModelTypeOverride
		return ret
	}
	return *o.ModelTypeOverride
}

// GetModelTypeOverrideOk returns a tuple with the ModelTypeOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) GetModelTypeOverrideOk() (*MonitorFormulaAndFunctionDataQualityModelTypeOverride, bool) {
	if o == nil || o.ModelTypeOverride == nil {
		return nil, false
	}
	return o.ModelTypeOverride, true
}

// HasModelTypeOverride returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) HasModelTypeOverride() bool {
	return o != nil && o.ModelTypeOverride != nil
}

// SetModelTypeOverride gets a reference to the given MonitorFormulaAndFunctionDataQualityModelTypeOverride and assigns it to the ModelTypeOverride field.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) SetModelTypeOverride(v MonitorFormulaAndFunctionDataQualityModelTypeOverride) {
	o.ModelTypeOverride = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualityMonitorOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrontabOverride != nil {
		toSerialize["crontab_override"] = o.CrontabOverride
	}
	if o.CustomSql != nil {
		toSerialize["custom_sql"] = o.CustomSql
	}
	if o.CustomWhere != nil {
		toSerialize["custom_where"] = o.CustomWhere
	}
	if o.GroupByColumns != nil {
		toSerialize["group_by_columns"] = o.GroupByColumns
	}
	if o.ModelTypeOverride != nil {
		toSerialize["model_type_override"] = o.ModelTypeOverride
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualityMonitorOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrontabOverride   *string                                                `json:"crontab_override,omitempty"`
		CustomSql         *string                                                `json:"custom_sql,omitempty"`
		CustomWhere       *string                                                `json:"custom_where,omitempty"`
		GroupByColumns    []string                                               `json:"group_by_columns,omitempty"`
		ModelTypeOverride *MonitorFormulaAndFunctionDataQualityModelTypeOverride `json:"model_type_override,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"crontab_override", "custom_sql", "custom_where", "group_by_columns", "model_type_override"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CrontabOverride = all.CrontabOverride
	o.CustomSql = all.CustomSql
	o.CustomWhere = all.CustomWhere
	o.GroupByColumns = all.GroupByColumns
	if all.ModelTypeOverride != nil && !all.ModelTypeOverride.IsValid() {
		hasInvalidField = true
	} else {
		o.ModelTypeOverride = all.ModelTypeOverride
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
