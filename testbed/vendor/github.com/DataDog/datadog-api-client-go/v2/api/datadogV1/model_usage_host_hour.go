// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageHostHour Number of hosts/containers recorded for each hour for a given organization.
type UsageHostHour struct {
	// Contains the total number of infrastructure hosts reporting
	// during a given hour that were running the Datadog Agent.
	AgentHostCount datadog.NullableInt64 `json:"agent_host_count,omitempty"`
	// Contains the total number of hosts that reported through Alibaba integration
	// (and were NOT running the Datadog Agent).
	AlibabaHostCount datadog.NullableInt64 `json:"alibaba_host_count,omitempty"`
	// Contains the total number of Azure App Services hosts using APM.
	ApmAzureAppServiceHostCount datadog.NullableInt64 `json:"apm_azure_app_service_host_count,omitempty"`
	// Shows the total number of hosts using APM during the hour,
	// these are counted as billable (except during trial periods).
	ApmHostCount datadog.NullableInt64 `json:"apm_host_count,omitempty"`
	// Contains the total number of hosts that reported through the AWS integration
	// (and were NOT running the Datadog Agent).
	AwsHostCount datadog.NullableInt64 `json:"aws_host_count,omitempty"`
	// Contains the total number of hosts that reported through Azure integration
	// (and were NOT running the Datadog Agent).
	AzureHostCount datadog.NullableInt64 `json:"azure_host_count,omitempty"`
	// Shows the total number of containers reported by the Docker integration during the hour.
	ContainerCount datadog.NullableInt64 `json:"container_count,omitempty"`
	// Contains the total number of hosts that reported through the Google Cloud integration
	// (and were NOT running the Datadog Agent).
	GcpHostCount datadog.NullableInt64 `json:"gcp_host_count,omitempty"`
	// Contains the total number of Heroku dynos reported by the Datadog Agent.
	HerokuHostCount datadog.NullableInt64 `json:"heroku_host_count,omitempty"`
	// Contains the total number of billable infrastructure hosts reporting during a given hour.
	// This is the sum of `agent_host_count`, `aws_host_count`, and `gcp_host_count`.
	HostCount datadog.NullableInt64 `json:"host_count,omitempty"`
	// The hour for the usage.
	Hour datadog.NullableTime `json:"hour,omitempty"`
	// Contains the total number of hosts that reported through the Azure App Services integration
	// (and were NOT running the Datadog Agent).
	InfraAzureAppService datadog.NullableInt64 `json:"infra_azure_app_service,omitempty"`
	// Contains the total number of hosts using APM reported by Datadog exporter for the OpenTelemetry Collector.
	OpentelemetryApmHostCount datadog.NullableInt64 `json:"opentelemetry_apm_host_count,omitempty"`
	// Contains the total number of hosts reported by Datadog exporter for the OpenTelemetry Collector.
	OpentelemetryHostCount datadog.NullableInt64 `json:"opentelemetry_host_count,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// Contains the total number of hosts that reported through vSphere integration
	// (and were NOT running the Datadog Agent).
	VsphereHostCount datadog.NullableInt64 `json:"vsphere_host_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageHostHour instantiates a new UsageHostHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageHostHour() *UsageHostHour {
	this := UsageHostHour{}
	return &this
}

// NewUsageHostHourWithDefaults instantiates a new UsageHostHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageHostHourWithDefaults() *UsageHostHour {
	this := UsageHostHour{}
	return &this
}

// GetAgentHostCount returns the AgentHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetAgentHostCount() int64 {
	if o == nil || o.AgentHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostCount.Get()
}

// GetAgentHostCountOk returns a tuple with the AgentHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetAgentHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentHostCount.Get(), o.AgentHostCount.IsSet()
}

// HasAgentHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasAgentHostCount() bool {
	return o != nil && o.AgentHostCount.IsSet()
}

// SetAgentHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the AgentHostCount field.
func (o *UsageHostHour) SetAgentHostCount(v int64) {
	o.AgentHostCount.Set(&v)
}

// SetAgentHostCountNil sets the value for AgentHostCount to be an explicit nil.
func (o *UsageHostHour) SetAgentHostCountNil() {
	o.AgentHostCount.Set(nil)
}

// UnsetAgentHostCount ensures that no value is present for AgentHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetAgentHostCount() {
	o.AgentHostCount.Unset()
}

// GetAlibabaHostCount returns the AlibabaHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetAlibabaHostCount() int64 {
	if o == nil || o.AlibabaHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AlibabaHostCount.Get()
}

// GetAlibabaHostCountOk returns a tuple with the AlibabaHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetAlibabaHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlibabaHostCount.Get(), o.AlibabaHostCount.IsSet()
}

// HasAlibabaHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasAlibabaHostCount() bool {
	return o != nil && o.AlibabaHostCount.IsSet()
}

// SetAlibabaHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the AlibabaHostCount field.
func (o *UsageHostHour) SetAlibabaHostCount(v int64) {
	o.AlibabaHostCount.Set(&v)
}

// SetAlibabaHostCountNil sets the value for AlibabaHostCount to be an explicit nil.
func (o *UsageHostHour) SetAlibabaHostCountNil() {
	o.AlibabaHostCount.Set(nil)
}

// UnsetAlibabaHostCount ensures that no value is present for AlibabaHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetAlibabaHostCount() {
	o.AlibabaHostCount.Unset()
}

// GetApmAzureAppServiceHostCount returns the ApmAzureAppServiceHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetApmAzureAppServiceHostCount() int64 {
	if o == nil || o.ApmAzureAppServiceHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmAzureAppServiceHostCount.Get()
}

// GetApmAzureAppServiceHostCountOk returns a tuple with the ApmAzureAppServiceHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetApmAzureAppServiceHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmAzureAppServiceHostCount.Get(), o.ApmAzureAppServiceHostCount.IsSet()
}

// HasApmAzureAppServiceHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasApmAzureAppServiceHostCount() bool {
	return o != nil && o.ApmAzureAppServiceHostCount.IsSet()
}

// SetApmAzureAppServiceHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the ApmAzureAppServiceHostCount field.
func (o *UsageHostHour) SetApmAzureAppServiceHostCount(v int64) {
	o.ApmAzureAppServiceHostCount.Set(&v)
}

// SetApmAzureAppServiceHostCountNil sets the value for ApmAzureAppServiceHostCount to be an explicit nil.
func (o *UsageHostHour) SetApmAzureAppServiceHostCountNil() {
	o.ApmAzureAppServiceHostCount.Set(nil)
}

// UnsetApmAzureAppServiceHostCount ensures that no value is present for ApmAzureAppServiceHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetApmAzureAppServiceHostCount() {
	o.ApmAzureAppServiceHostCount.Unset()
}

// GetApmHostCount returns the ApmHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetApmHostCount() int64 {
	if o == nil || o.ApmHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostCount.Get()
}

// GetApmHostCountOk returns a tuple with the ApmHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetApmHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApmHostCount.Get(), o.ApmHostCount.IsSet()
}

// HasApmHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasApmHostCount() bool {
	return o != nil && o.ApmHostCount.IsSet()
}

// SetApmHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the ApmHostCount field.
func (o *UsageHostHour) SetApmHostCount(v int64) {
	o.ApmHostCount.Set(&v)
}

// SetApmHostCountNil sets the value for ApmHostCount to be an explicit nil.
func (o *UsageHostHour) SetApmHostCountNil() {
	o.ApmHostCount.Set(nil)
}

// UnsetApmHostCount ensures that no value is present for ApmHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetApmHostCount() {
	o.ApmHostCount.Unset()
}

// GetAwsHostCount returns the AwsHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetAwsHostCount() int64 {
	if o == nil || o.AwsHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostCount.Get()
}

// GetAwsHostCountOk returns a tuple with the AwsHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetAwsHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsHostCount.Get(), o.AwsHostCount.IsSet()
}

// HasAwsHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasAwsHostCount() bool {
	return o != nil && o.AwsHostCount.IsSet()
}

// SetAwsHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the AwsHostCount field.
func (o *UsageHostHour) SetAwsHostCount(v int64) {
	o.AwsHostCount.Set(&v)
}

// SetAwsHostCountNil sets the value for AwsHostCount to be an explicit nil.
func (o *UsageHostHour) SetAwsHostCountNil() {
	o.AwsHostCount.Set(nil)
}

// UnsetAwsHostCount ensures that no value is present for AwsHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetAwsHostCount() {
	o.AwsHostCount.Unset()
}

// GetAzureHostCount returns the AzureHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetAzureHostCount() int64 {
	if o == nil || o.AzureHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AzureHostCount.Get()
}

// GetAzureHostCountOk returns a tuple with the AzureHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetAzureHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureHostCount.Get(), o.AzureHostCount.IsSet()
}

// HasAzureHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasAzureHostCount() bool {
	return o != nil && o.AzureHostCount.IsSet()
}

// SetAzureHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the AzureHostCount field.
func (o *UsageHostHour) SetAzureHostCount(v int64) {
	o.AzureHostCount.Set(&v)
}

// SetAzureHostCountNil sets the value for AzureHostCount to be an explicit nil.
func (o *UsageHostHour) SetAzureHostCountNil() {
	o.AzureHostCount.Set(nil)
}

// UnsetAzureHostCount ensures that no value is present for AzureHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetAzureHostCount() {
	o.AzureHostCount.Unset()
}

// GetContainerCount returns the ContainerCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetContainerCount() int64 {
	if o == nil || o.ContainerCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerCount.Get()
}

// GetContainerCountOk returns a tuple with the ContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetContainerCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerCount.Get(), o.ContainerCount.IsSet()
}

// HasContainerCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasContainerCount() bool {
	return o != nil && o.ContainerCount.IsSet()
}

// SetContainerCount gets a reference to the given datadog.NullableInt64 and assigns it to the ContainerCount field.
func (o *UsageHostHour) SetContainerCount(v int64) {
	o.ContainerCount.Set(&v)
}

// SetContainerCountNil sets the value for ContainerCount to be an explicit nil.
func (o *UsageHostHour) SetContainerCountNil() {
	o.ContainerCount.Set(nil)
}

// UnsetContainerCount ensures that no value is present for ContainerCount, not even an explicit nil.
func (o *UsageHostHour) UnsetContainerCount() {
	o.ContainerCount.Unset()
}

// GetGcpHostCount returns the GcpHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetGcpHostCount() int64 {
	if o == nil || o.GcpHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostCount.Get()
}

// GetGcpHostCountOk returns a tuple with the GcpHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetGcpHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GcpHostCount.Get(), o.GcpHostCount.IsSet()
}

// HasGcpHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasGcpHostCount() bool {
	return o != nil && o.GcpHostCount.IsSet()
}

// SetGcpHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the GcpHostCount field.
func (o *UsageHostHour) SetGcpHostCount(v int64) {
	o.GcpHostCount.Set(&v)
}

// SetGcpHostCountNil sets the value for GcpHostCount to be an explicit nil.
func (o *UsageHostHour) SetGcpHostCountNil() {
	o.GcpHostCount.Set(nil)
}

// UnsetGcpHostCount ensures that no value is present for GcpHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetGcpHostCount() {
	o.GcpHostCount.Unset()
}

// GetHerokuHostCount returns the HerokuHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetHerokuHostCount() int64 {
	if o == nil || o.HerokuHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HerokuHostCount.Get()
}

// GetHerokuHostCountOk returns a tuple with the HerokuHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetHerokuHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HerokuHostCount.Get(), o.HerokuHostCount.IsSet()
}

// HasHerokuHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasHerokuHostCount() bool {
	return o != nil && o.HerokuHostCount.IsSet()
}

// SetHerokuHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the HerokuHostCount field.
func (o *UsageHostHour) SetHerokuHostCount(v int64) {
	o.HerokuHostCount.Set(&v)
}

// SetHerokuHostCountNil sets the value for HerokuHostCount to be an explicit nil.
func (o *UsageHostHour) SetHerokuHostCountNil() {
	o.HerokuHostCount.Set(nil)
}

// UnsetHerokuHostCount ensures that no value is present for HerokuHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetHerokuHostCount() {
	o.HerokuHostCount.Unset()
}

// GetHostCount returns the HostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetHostCount() int64 {
	if o == nil || o.HostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HostCount.Get()
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostCount.Get(), o.HostCount.IsSet()
}

// HasHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasHostCount() bool {
	return o != nil && o.HostCount.IsSet()
}

// SetHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the HostCount field.
func (o *UsageHostHour) SetHostCount(v int64) {
	o.HostCount.Set(&v)
}

// SetHostCountNil sets the value for HostCount to be an explicit nil.
func (o *UsageHostHour) SetHostCountNil() {
	o.HostCount.Set(nil)
}

// UnsetHostCount ensures that no value is present for HostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetHostCount() {
	o.HostCount.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetHour() time.Time {
	if o == nil || o.Hour.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour.Get()
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetHourOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hour.Get(), o.Hour.IsSet()
}

// HasHour returns a boolean if a field has been set.
func (o *UsageHostHour) HasHour() bool {
	return o != nil && o.Hour.IsSet()
}

// SetHour gets a reference to the given datadog.NullableTime and assigns it to the Hour field.
func (o *UsageHostHour) SetHour(v time.Time) {
	o.Hour.Set(&v)
}

// SetHourNil sets the value for Hour to be an explicit nil.
func (o *UsageHostHour) SetHourNil() {
	o.Hour.Set(nil)
}

// UnsetHour ensures that no value is present for Hour, not even an explicit nil.
func (o *UsageHostHour) UnsetHour() {
	o.Hour.Unset()
}

// GetInfraAzureAppService returns the InfraAzureAppService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetInfraAzureAppService() int64 {
	if o == nil || o.InfraAzureAppService.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InfraAzureAppService.Get()
}

// GetInfraAzureAppServiceOk returns a tuple with the InfraAzureAppService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetInfraAzureAppServiceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraAzureAppService.Get(), o.InfraAzureAppService.IsSet()
}

// HasInfraAzureAppService returns a boolean if a field has been set.
func (o *UsageHostHour) HasInfraAzureAppService() bool {
	return o != nil && o.InfraAzureAppService.IsSet()
}

// SetInfraAzureAppService gets a reference to the given datadog.NullableInt64 and assigns it to the InfraAzureAppService field.
func (o *UsageHostHour) SetInfraAzureAppService(v int64) {
	o.InfraAzureAppService.Set(&v)
}

// SetInfraAzureAppServiceNil sets the value for InfraAzureAppService to be an explicit nil.
func (o *UsageHostHour) SetInfraAzureAppServiceNil() {
	o.InfraAzureAppService.Set(nil)
}

// UnsetInfraAzureAppService ensures that no value is present for InfraAzureAppService, not even an explicit nil.
func (o *UsageHostHour) UnsetInfraAzureAppService() {
	o.InfraAzureAppService.Unset()
}

// GetOpentelemetryApmHostCount returns the OpentelemetryApmHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetOpentelemetryApmHostCount() int64 {
	if o == nil || o.OpentelemetryApmHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryApmHostCount.Get()
}

// GetOpentelemetryApmHostCountOk returns a tuple with the OpentelemetryApmHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetOpentelemetryApmHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryApmHostCount.Get(), o.OpentelemetryApmHostCount.IsSet()
}

// HasOpentelemetryApmHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasOpentelemetryApmHostCount() bool {
	return o != nil && o.OpentelemetryApmHostCount.IsSet()
}

// SetOpentelemetryApmHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryApmHostCount field.
func (o *UsageHostHour) SetOpentelemetryApmHostCount(v int64) {
	o.OpentelemetryApmHostCount.Set(&v)
}

// SetOpentelemetryApmHostCountNil sets the value for OpentelemetryApmHostCount to be an explicit nil.
func (o *UsageHostHour) SetOpentelemetryApmHostCountNil() {
	o.OpentelemetryApmHostCount.Set(nil)
}

// UnsetOpentelemetryApmHostCount ensures that no value is present for OpentelemetryApmHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetOpentelemetryApmHostCount() {
	o.OpentelemetryApmHostCount.Unset()
}

// GetOpentelemetryHostCount returns the OpentelemetryHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetOpentelemetryHostCount() int64 {
	if o == nil || o.OpentelemetryHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryHostCount.Get()
}

// GetOpentelemetryHostCountOk returns a tuple with the OpentelemetryHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetOpentelemetryHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpentelemetryHostCount.Get(), o.OpentelemetryHostCount.IsSet()
}

// HasOpentelemetryHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasOpentelemetryHostCount() bool {
	return o != nil && o.OpentelemetryHostCount.IsSet()
}

// SetOpentelemetryHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the OpentelemetryHostCount field.
func (o *UsageHostHour) SetOpentelemetryHostCount(v int64) {
	o.OpentelemetryHostCount.Set(&v)
}

// SetOpentelemetryHostCountNil sets the value for OpentelemetryHostCount to be an explicit nil.
func (o *UsageHostHour) SetOpentelemetryHostCountNil() {
	o.OpentelemetryHostCount.Set(nil)
}

// UnsetOpentelemetryHostCount ensures that no value is present for OpentelemetryHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetOpentelemetryHostCount() {
	o.OpentelemetryHostCount.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageHostHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageHostHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageHostHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageHostHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageHostHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageHostHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageHostHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageHostHour) SetPublicId(v string) {
	o.PublicId = &v
}

// GetVsphereHostCount returns the VsphereHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageHostHour) GetVsphereHostCount() int64 {
	if o == nil || o.VsphereHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.VsphereHostCount.Get()
}

// GetVsphereHostCountOk returns a tuple with the VsphereHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageHostHour) GetVsphereHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VsphereHostCount.Get(), o.VsphereHostCount.IsSet()
}

// HasVsphereHostCount returns a boolean if a field has been set.
func (o *UsageHostHour) HasVsphereHostCount() bool {
	return o != nil && o.VsphereHostCount.IsSet()
}

// SetVsphereHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the VsphereHostCount field.
func (o *UsageHostHour) SetVsphereHostCount(v int64) {
	o.VsphereHostCount.Set(&v)
}

// SetVsphereHostCountNil sets the value for VsphereHostCount to be an explicit nil.
func (o *UsageHostHour) SetVsphereHostCountNil() {
	o.VsphereHostCount.Set(nil)
}

// UnsetVsphereHostCount ensures that no value is present for VsphereHostCount, not even an explicit nil.
func (o *UsageHostHour) UnsetVsphereHostCount() {
	o.VsphereHostCount.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageHostHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentHostCount.IsSet() {
		toSerialize["agent_host_count"] = o.AgentHostCount.Get()
	}
	if o.AlibabaHostCount.IsSet() {
		toSerialize["alibaba_host_count"] = o.AlibabaHostCount.Get()
	}
	if o.ApmAzureAppServiceHostCount.IsSet() {
		toSerialize["apm_azure_app_service_host_count"] = o.ApmAzureAppServiceHostCount.Get()
	}
	if o.ApmHostCount.IsSet() {
		toSerialize["apm_host_count"] = o.ApmHostCount.Get()
	}
	if o.AwsHostCount.IsSet() {
		toSerialize["aws_host_count"] = o.AwsHostCount.Get()
	}
	if o.AzureHostCount.IsSet() {
		toSerialize["azure_host_count"] = o.AzureHostCount.Get()
	}
	if o.ContainerCount.IsSet() {
		toSerialize["container_count"] = o.ContainerCount.Get()
	}
	if o.GcpHostCount.IsSet() {
		toSerialize["gcp_host_count"] = o.GcpHostCount.Get()
	}
	if o.HerokuHostCount.IsSet() {
		toSerialize["heroku_host_count"] = o.HerokuHostCount.Get()
	}
	if o.HostCount.IsSet() {
		toSerialize["host_count"] = o.HostCount.Get()
	}
	if o.Hour.IsSet() {
		toSerialize["hour"] = o.Hour.Get()
	}
	if o.InfraAzureAppService.IsSet() {
		toSerialize["infra_azure_app_service"] = o.InfraAzureAppService.Get()
	}
	if o.OpentelemetryApmHostCount.IsSet() {
		toSerialize["opentelemetry_apm_host_count"] = o.OpentelemetryApmHostCount.Get()
	}
	if o.OpentelemetryHostCount.IsSet() {
		toSerialize["opentelemetry_host_count"] = o.OpentelemetryHostCount.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.VsphereHostCount.IsSet() {
		toSerialize["vsphere_host_count"] = o.VsphereHostCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageHostHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentHostCount              datadog.NullableInt64 `json:"agent_host_count,omitempty"`
		AlibabaHostCount            datadog.NullableInt64 `json:"alibaba_host_count,omitempty"`
		ApmAzureAppServiceHostCount datadog.NullableInt64 `json:"apm_azure_app_service_host_count,omitempty"`
		ApmHostCount                datadog.NullableInt64 `json:"apm_host_count,omitempty"`
		AwsHostCount                datadog.NullableInt64 `json:"aws_host_count,omitempty"`
		AzureHostCount              datadog.NullableInt64 `json:"azure_host_count,omitempty"`
		ContainerCount              datadog.NullableInt64 `json:"container_count,omitempty"`
		GcpHostCount                datadog.NullableInt64 `json:"gcp_host_count,omitempty"`
		HerokuHostCount             datadog.NullableInt64 `json:"heroku_host_count,omitempty"`
		HostCount                   datadog.NullableInt64 `json:"host_count,omitempty"`
		Hour                        datadog.NullableTime  `json:"hour,omitempty"`
		InfraAzureAppService        datadog.NullableInt64 `json:"infra_azure_app_service,omitempty"`
		OpentelemetryApmHostCount   datadog.NullableInt64 `json:"opentelemetry_apm_host_count,omitempty"`
		OpentelemetryHostCount      datadog.NullableInt64 `json:"opentelemetry_host_count,omitempty"`
		OrgName                     *string               `json:"org_name,omitempty"`
		PublicId                    *string               `json:"public_id,omitempty"`
		VsphereHostCount            datadog.NullableInt64 `json:"vsphere_host_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_host_count", "alibaba_host_count", "apm_azure_app_service_host_count", "apm_host_count", "aws_host_count", "azure_host_count", "container_count", "gcp_host_count", "heroku_host_count", "host_count", "hour", "infra_azure_app_service", "opentelemetry_apm_host_count", "opentelemetry_host_count", "org_name", "public_id", "vsphere_host_count"})
	} else {
		return err
	}
	o.AgentHostCount = all.AgentHostCount
	o.AlibabaHostCount = all.AlibabaHostCount
	o.ApmAzureAppServiceHostCount = all.ApmAzureAppServiceHostCount
	o.ApmHostCount = all.ApmHostCount
	o.AwsHostCount = all.AwsHostCount
	o.AzureHostCount = all.AzureHostCount
	o.ContainerCount = all.ContainerCount
	o.GcpHostCount = all.GcpHostCount
	o.HerokuHostCount = all.HerokuHostCount
	o.HostCount = all.HostCount
	o.Hour = all.Hour
	o.InfraAzureAppService = all.InfraAzureAppService
	o.OpentelemetryApmHostCount = all.OpentelemetryApmHostCount
	o.OpentelemetryHostCount = all.OpentelemetryHostCount
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.VsphereHostCount = all.VsphereHostCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
