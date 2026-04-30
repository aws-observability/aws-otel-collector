// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmDependencyStatsQuery A query for APM dependency statistics between services, such as call latency and error rates.
type ApmDependencyStatsQuery struct {
	// A data source for APM dependency statistics queries.
	DataSource ApmDependencyStatsDataSource `json:"data_source"`
	// The environment to query.
	Env string `json:"env"`
	// Determines whether stats for upstream or downstream dependencies should be queried.
	IsUpstream *bool `json:"is_upstream,omitempty"`
	// The variable name for use in formulas.
	Name string `json:"name"`
	// The APM operation name.
	OperationName string `json:"operation_name"`
	// The name of the second primary tag used within APM; required when `primary_tag_value` is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog.
	PrimaryTagName *string `json:"primary_tag_name,omitempty"`
	// Filter APM data by the second primary tag. `primary_tag_name` must also be specified.
	PrimaryTagValue *string `json:"primary_tag_value,omitempty"`
	// The resource name to filter by.
	ResourceName string `json:"resource_name"`
	// The service name to filter by.
	Service string `json:"service"`
	// The APM dependency statistic to query.
	Stat ApmDependencyStatName `json:"stat"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApmDependencyStatsQuery instantiates a new ApmDependencyStatsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApmDependencyStatsQuery(dataSource ApmDependencyStatsDataSource, env string, name string, operationName string, resourceName string, service string, stat ApmDependencyStatName) *ApmDependencyStatsQuery {
	this := ApmDependencyStatsQuery{}
	this.DataSource = dataSource
	this.Env = env
	this.Name = name
	this.OperationName = operationName
	this.ResourceName = resourceName
	this.Service = service
	this.Stat = stat
	return &this
}

// NewApmDependencyStatsQueryWithDefaults instantiates a new ApmDependencyStatsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApmDependencyStatsQueryWithDefaults() *ApmDependencyStatsQuery {
	this := ApmDependencyStatsQuery{}
	var dataSource ApmDependencyStatsDataSource = APMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS
	this.DataSource = dataSource
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ApmDependencyStatsQuery) GetDataSource() ApmDependencyStatsDataSource {
	if o == nil {
		var ret ApmDependencyStatsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetDataSourceOk() (*ApmDependencyStatsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ApmDependencyStatsQuery) SetDataSource(v ApmDependencyStatsDataSource) {
	o.DataSource = v
}

// GetEnv returns the Env field value.
func (o *ApmDependencyStatsQuery) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ApmDependencyStatsQuery) SetEnv(v string) {
	o.Env = v
}

// GetIsUpstream returns the IsUpstream field value if set, zero value otherwise.
func (o *ApmDependencyStatsQuery) GetIsUpstream() bool {
	if o == nil || o.IsUpstream == nil {
		var ret bool
		return ret
	}
	return *o.IsUpstream
}

// GetIsUpstreamOk returns a tuple with the IsUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetIsUpstreamOk() (*bool, bool) {
	if o == nil || o.IsUpstream == nil {
		return nil, false
	}
	return o.IsUpstream, true
}

// HasIsUpstream returns a boolean if a field has been set.
func (o *ApmDependencyStatsQuery) HasIsUpstream() bool {
	return o != nil && o.IsUpstream != nil
}

// SetIsUpstream gets a reference to the given bool and assigns it to the IsUpstream field.
func (o *ApmDependencyStatsQuery) SetIsUpstream(v bool) {
	o.IsUpstream = &v
}

// GetName returns the Name field value.
func (o *ApmDependencyStatsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApmDependencyStatsQuery) SetName(v string) {
	o.Name = v
}

// GetOperationName returns the OperationName field value.
func (o *ApmDependencyStatsQuery) GetOperationName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetOperationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationName, true
}

// SetOperationName sets field value.
func (o *ApmDependencyStatsQuery) SetOperationName(v string) {
	o.OperationName = v
}

// GetPrimaryTagName returns the PrimaryTagName field value if set, zero value otherwise.
func (o *ApmDependencyStatsQuery) GetPrimaryTagName() string {
	if o == nil || o.PrimaryTagName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagName
}

// GetPrimaryTagNameOk returns a tuple with the PrimaryTagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetPrimaryTagNameOk() (*string, bool) {
	if o == nil || o.PrimaryTagName == nil {
		return nil, false
	}
	return o.PrimaryTagName, true
}

// HasPrimaryTagName returns a boolean if a field has been set.
func (o *ApmDependencyStatsQuery) HasPrimaryTagName() bool {
	return o != nil && o.PrimaryTagName != nil
}

// SetPrimaryTagName gets a reference to the given string and assigns it to the PrimaryTagName field.
func (o *ApmDependencyStatsQuery) SetPrimaryTagName(v string) {
	o.PrimaryTagName = &v
}

// GetPrimaryTagValue returns the PrimaryTagValue field value if set, zero value otherwise.
func (o *ApmDependencyStatsQuery) GetPrimaryTagValue() string {
	if o == nil || o.PrimaryTagValue == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagValue
}

// GetPrimaryTagValueOk returns a tuple with the PrimaryTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetPrimaryTagValueOk() (*string, bool) {
	if o == nil || o.PrimaryTagValue == nil {
		return nil, false
	}
	return o.PrimaryTagValue, true
}

// HasPrimaryTagValue returns a boolean if a field has been set.
func (o *ApmDependencyStatsQuery) HasPrimaryTagValue() bool {
	return o != nil && o.PrimaryTagValue != nil
}

// SetPrimaryTagValue gets a reference to the given string and assigns it to the PrimaryTagValue field.
func (o *ApmDependencyStatsQuery) SetPrimaryTagValue(v string) {
	o.PrimaryTagValue = &v
}

// GetResourceName returns the ResourceName field value.
func (o *ApmDependencyStatsQuery) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value.
func (o *ApmDependencyStatsQuery) SetResourceName(v string) {
	o.ResourceName = v
}

// GetService returns the Service field value.
func (o *ApmDependencyStatsQuery) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ApmDependencyStatsQuery) SetService(v string) {
	o.Service = v
}

// GetStat returns the Stat field value.
func (o *ApmDependencyStatsQuery) GetStat() ApmDependencyStatName {
	if o == nil {
		var ret ApmDependencyStatName
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *ApmDependencyStatsQuery) GetStatOk() (*ApmDependencyStatName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *ApmDependencyStatsQuery) SetStat(v ApmDependencyStatName) {
	o.Stat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApmDependencyStatsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["env"] = o.Env
	if o.IsUpstream != nil {
		toSerialize["is_upstream"] = o.IsUpstream
	}
	toSerialize["name"] = o.Name
	toSerialize["operation_name"] = o.OperationName
	if o.PrimaryTagName != nil {
		toSerialize["primary_tag_name"] = o.PrimaryTagName
	}
	if o.PrimaryTagValue != nil {
		toSerialize["primary_tag_value"] = o.PrimaryTagValue
	}
	toSerialize["resource_name"] = o.ResourceName
	toSerialize["service"] = o.Service
	toSerialize["stat"] = o.Stat

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApmDependencyStatsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *ApmDependencyStatsDataSource `json:"data_source"`
		Env             *string                       `json:"env"`
		IsUpstream      *bool                         `json:"is_upstream,omitempty"`
		Name            *string                       `json:"name"`
		OperationName   *string                       `json:"operation_name"`
		PrimaryTagName  *string                       `json:"primary_tag_name,omitempty"`
		PrimaryTagValue *string                       `json:"primary_tag_value,omitempty"`
		ResourceName    *string                       `json:"resource_name"`
		Service         *string                       `json:"service"`
		Stat            *ApmDependencyStatName        `json:"stat"`
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
	if all.OperationName == nil {
		return fmt.Errorf("required field operation_name missing")
	}
	if all.ResourceName == nil {
		return fmt.Errorf("required field resource_name missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Stat == nil {
		return fmt.Errorf("required field stat missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "env", "is_upstream", "name", "operation_name", "primary_tag_name", "primary_tag_value", "resource_name", "service", "stat"})
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
	o.IsUpstream = all.IsUpstream
	o.Name = *all.Name
	o.OperationName = *all.OperationName
	o.PrimaryTagName = all.PrimaryTagName
	o.PrimaryTagValue = all.PrimaryTagValue
	o.ResourceName = *all.ResourceName
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
