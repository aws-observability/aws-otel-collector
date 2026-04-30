// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmResourceStatsQuery A query for APM resource statistics such as latency, error rate, and hit count, grouped by resource name.
type ApmResourceStatsQuery struct {
	// A data source for APM resource statistics queries.
	DataSource ApmResourceStatsDataSource `json:"data_source"`
	// The environment to query.
	Env string `json:"env"`
	// Tag keys to group results by.
	GroupBy []string `json:"group_by,omitempty"`
	// The variable name for use in formulas.
	Name string `json:"name"`
	// The APM operation name.
	OperationName *string `json:"operation_name,omitempty"`
	// Name of the second primary tag used within APM. Required when `primary_tag_value` is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog
	PrimaryTagName *string `json:"primary_tag_name,omitempty"`
	// Value of the second primary tag by which to filter APM data. `primary_tag_name` must also be specified.
	PrimaryTagValue *string `json:"primary_tag_value,omitempty"`
	// The resource name to filter by.
	ResourceName *string `json:"resource_name,omitempty"`
	// The service name to filter by.
	Service string `json:"service"`
	// The APM resource statistic to query.
	Stat ApmResourceStatName `json:"stat"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApmResourceStatsQuery instantiates a new ApmResourceStatsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApmResourceStatsQuery(dataSource ApmResourceStatsDataSource, env string, name string, service string, stat ApmResourceStatName) *ApmResourceStatsQuery {
	this := ApmResourceStatsQuery{}
	this.DataSource = dataSource
	this.Env = env
	this.Name = name
	this.Service = service
	this.Stat = stat
	return &this
}

// NewApmResourceStatsQueryWithDefaults instantiates a new ApmResourceStatsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApmResourceStatsQueryWithDefaults() *ApmResourceStatsQuery {
	this := ApmResourceStatsQuery{}
	var dataSource ApmResourceStatsDataSource = APMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS
	this.DataSource = dataSource
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ApmResourceStatsQuery) GetDataSource() ApmResourceStatsDataSource {
	if o == nil {
		var ret ApmResourceStatsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetDataSourceOk() (*ApmResourceStatsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ApmResourceStatsQuery) SetDataSource(v ApmResourceStatsDataSource) {
	o.DataSource = v
}

// GetEnv returns the Env field value.
func (o *ApmResourceStatsQuery) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ApmResourceStatsQuery) SetEnv(v string) {
	o.Env = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ApmResourceStatsQuery) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ApmResourceStatsQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *ApmResourceStatsQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *ApmResourceStatsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApmResourceStatsQuery) SetName(v string) {
	o.Name = v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *ApmResourceStatsQuery) GetOperationName() string {
	if o == nil || o.OperationName == nil {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetOperationNameOk() (*string, bool) {
	if o == nil || o.OperationName == nil {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *ApmResourceStatsQuery) HasOperationName() bool {
	return o != nil && o.OperationName != nil
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *ApmResourceStatsQuery) SetOperationName(v string) {
	o.OperationName = &v
}

// GetPrimaryTagName returns the PrimaryTagName field value if set, zero value otherwise.
func (o *ApmResourceStatsQuery) GetPrimaryTagName() string {
	if o == nil || o.PrimaryTagName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagName
}

// GetPrimaryTagNameOk returns a tuple with the PrimaryTagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetPrimaryTagNameOk() (*string, bool) {
	if o == nil || o.PrimaryTagName == nil {
		return nil, false
	}
	return o.PrimaryTagName, true
}

// HasPrimaryTagName returns a boolean if a field has been set.
func (o *ApmResourceStatsQuery) HasPrimaryTagName() bool {
	return o != nil && o.PrimaryTagName != nil
}

// SetPrimaryTagName gets a reference to the given string and assigns it to the PrimaryTagName field.
func (o *ApmResourceStatsQuery) SetPrimaryTagName(v string) {
	o.PrimaryTagName = &v
}

// GetPrimaryTagValue returns the PrimaryTagValue field value if set, zero value otherwise.
func (o *ApmResourceStatsQuery) GetPrimaryTagValue() string {
	if o == nil || o.PrimaryTagValue == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagValue
}

// GetPrimaryTagValueOk returns a tuple with the PrimaryTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetPrimaryTagValueOk() (*string, bool) {
	if o == nil || o.PrimaryTagValue == nil {
		return nil, false
	}
	return o.PrimaryTagValue, true
}

// HasPrimaryTagValue returns a boolean if a field has been set.
func (o *ApmResourceStatsQuery) HasPrimaryTagValue() bool {
	return o != nil && o.PrimaryTagValue != nil
}

// SetPrimaryTagValue gets a reference to the given string and assigns it to the PrimaryTagValue field.
func (o *ApmResourceStatsQuery) SetPrimaryTagValue(v string) {
	o.PrimaryTagValue = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ApmResourceStatsQuery) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ApmResourceStatsQuery) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ApmResourceStatsQuery) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetService returns the Service field value.
func (o *ApmResourceStatsQuery) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ApmResourceStatsQuery) SetService(v string) {
	o.Service = v
}

// GetStat returns the Stat field value.
func (o *ApmResourceStatsQuery) GetStat() ApmResourceStatName {
	if o == nil {
		var ret ApmResourceStatName
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *ApmResourceStatsQuery) GetStatOk() (*ApmResourceStatName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *ApmResourceStatsQuery) SetStat(v ApmResourceStatName) {
	o.Stat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApmResourceStatsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["env"] = o.Env
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	if o.OperationName != nil {
		toSerialize["operation_name"] = o.OperationName
	}
	if o.PrimaryTagName != nil {
		toSerialize["primary_tag_name"] = o.PrimaryTagName
	}
	if o.PrimaryTagValue != nil {
		toSerialize["primary_tag_value"] = o.PrimaryTagValue
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	toSerialize["service"] = o.Service
	toSerialize["stat"] = o.Stat

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApmResourceStatsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *ApmResourceStatsDataSource `json:"data_source"`
		Env             *string                     `json:"env"`
		GroupBy         []string                    `json:"group_by,omitempty"`
		Name            *string                     `json:"name"`
		OperationName   *string                     `json:"operation_name,omitempty"`
		PrimaryTagName  *string                     `json:"primary_tag_name,omitempty"`
		PrimaryTagValue *string                     `json:"primary_tag_value,omitempty"`
		ResourceName    *string                     `json:"resource_name,omitempty"`
		Service         *string                     `json:"service"`
		Stat            *ApmResourceStatName        `json:"stat"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Stat == nil {
		return fmt.Errorf("required field stat missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "env", "group_by", "name", "operation_name", "primary_tag_name", "primary_tag_value", "resource_name", "service", "stat"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Env = *all.Env
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	o.OperationName = all.OperationName
	o.PrimaryTagName = all.PrimaryTagName
	o.PrimaryTagValue = all.PrimaryTagValue
	o.ResourceName = all.ResourceName
	o.Service = *all.Service
	if !all.Stat.IsValid() {
		hasInvalidField = true
	} else {
		o.Stat = *all.Stat
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
