// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureAccount Datadog-Azure integrations configured for your organization.
type AzureAccount struct {
	// Limit the Azure app service plans that are pulled into Datadog using tags.
	// Only app service plans that match one of the defined tags are imported into Datadog.
	AppServicePlanFilters *string `json:"app_service_plan_filters,omitempty"`
	// Silence monitors for expected Azure VM shutdowns.
	Automute *bool `json:"automute,omitempty"`
	// Your Azure web application ID.
	ClientId *string `json:"client_id,omitempty"`
	// Your Azure web application secret key.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Limit the Azure container apps that are pulled into Datadog using tags.
	// Only container apps that match one of the defined tags are imported into Datadog.
	ContainerAppFilters *string `json:"container_app_filters,omitempty"`
	// When enabled, Datadog’s Cloud Security Management product scans resource configurations monitored by this app registration.
	// Note: This requires resource_collection_enabled to be set to true.
	CspmEnabled *bool `json:"cspm_enabled,omitempty"`
	// Enable custom metrics for your organization.
	CustomMetricsEnabled *bool `json:"custom_metrics_enabled,omitempty"`
	// Errors in your configuration.
	Errors []string `json:"errors,omitempty"`
	// Limit the Azure instances that are pulled into Datadog by using tags.
	// Only hosts that match one of the defined tags are imported into Datadog.
	HostFilters *string `json:"host_filters,omitempty"`
	// Enable Azure metrics for your organization.
	MetricsEnabled *bool `json:"metrics_enabled,omitempty"`
	// Enable Azure metrics for your organization for resource providers where no resource provider config is specified.
	MetricsEnabledDefault *bool `json:"metrics_enabled_default,omitempty"`
	// Your New Azure web application ID.
	NewClientId *string `json:"new_client_id,omitempty"`
	// Your New Azure Active Directory ID.
	NewTenantName *string `json:"new_tenant_name,omitempty"`
	// When enabled, Datadog collects metadata and configuration info from cloud resources (compute instances, databases, load balancers, etc.) monitored by this app registration.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// Configuration settings applied to resources from the specified Azure resource providers.
	ResourceProviderConfigs []ResourceProviderConfig `json:"resource_provider_configs,omitempty"`
	// Your Azure Active Directory ID.
	TenantName *string `json:"tenant_name,omitempty"`
	// Enable azure.usage metrics for your organization.
	UsageMetricsEnabled *bool `json:"usage_metrics_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureAccount instantiates a new AzureAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureAccount() *AzureAccount {
	this := AzureAccount{}
	return &this
}

// NewAzureAccountWithDefaults instantiates a new AzureAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureAccountWithDefaults() *AzureAccount {
	this := AzureAccount{}
	return &this
}

// GetAppServicePlanFilters returns the AppServicePlanFilters field value if set, zero value otherwise.
func (o *AzureAccount) GetAppServicePlanFilters() string {
	if o == nil || o.AppServicePlanFilters == nil {
		var ret string
		return ret
	}
	return *o.AppServicePlanFilters
}

// GetAppServicePlanFiltersOk returns a tuple with the AppServicePlanFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetAppServicePlanFiltersOk() (*string, bool) {
	if o == nil || o.AppServicePlanFilters == nil {
		return nil, false
	}
	return o.AppServicePlanFilters, true
}

// HasAppServicePlanFilters returns a boolean if a field has been set.
func (o *AzureAccount) HasAppServicePlanFilters() bool {
	return o != nil && o.AppServicePlanFilters != nil
}

// SetAppServicePlanFilters gets a reference to the given string and assigns it to the AppServicePlanFilters field.
func (o *AzureAccount) SetAppServicePlanFilters(v string) {
	o.AppServicePlanFilters = &v
}

// GetAutomute returns the Automute field value if set, zero value otherwise.
func (o *AzureAccount) GetAutomute() bool {
	if o == nil || o.Automute == nil {
		var ret bool
		return ret
	}
	return *o.Automute
}

// GetAutomuteOk returns a tuple with the Automute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetAutomuteOk() (*bool, bool) {
	if o == nil || o.Automute == nil {
		return nil, false
	}
	return o.Automute, true
}

// HasAutomute returns a boolean if a field has been set.
func (o *AzureAccount) HasAutomute() bool {
	return o != nil && o.Automute != nil
}

// SetAutomute gets a reference to the given bool and assigns it to the Automute field.
func (o *AzureAccount) SetAutomute(v bool) {
	o.Automute = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AzureAccount) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AzureAccount) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AzureAccount) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AzureAccount) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AzureAccount) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AzureAccount) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetContainerAppFilters returns the ContainerAppFilters field value if set, zero value otherwise.
func (o *AzureAccount) GetContainerAppFilters() string {
	if o == nil || o.ContainerAppFilters == nil {
		var ret string
		return ret
	}
	return *o.ContainerAppFilters
}

// GetContainerAppFiltersOk returns a tuple with the ContainerAppFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetContainerAppFiltersOk() (*string, bool) {
	if o == nil || o.ContainerAppFilters == nil {
		return nil, false
	}
	return o.ContainerAppFilters, true
}

// HasContainerAppFilters returns a boolean if a field has been set.
func (o *AzureAccount) HasContainerAppFilters() bool {
	return o != nil && o.ContainerAppFilters != nil
}

// SetContainerAppFilters gets a reference to the given string and assigns it to the ContainerAppFilters field.
func (o *AzureAccount) SetContainerAppFilters(v string) {
	o.ContainerAppFilters = &v
}

// GetCspmEnabled returns the CspmEnabled field value if set, zero value otherwise.
func (o *AzureAccount) GetCspmEnabled() bool {
	if o == nil || o.CspmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CspmEnabled
}

// GetCspmEnabledOk returns a tuple with the CspmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetCspmEnabledOk() (*bool, bool) {
	if o == nil || o.CspmEnabled == nil {
		return nil, false
	}
	return o.CspmEnabled, true
}

// HasCspmEnabled returns a boolean if a field has been set.
func (o *AzureAccount) HasCspmEnabled() bool {
	return o != nil && o.CspmEnabled != nil
}

// SetCspmEnabled gets a reference to the given bool and assigns it to the CspmEnabled field.
func (o *AzureAccount) SetCspmEnabled(v bool) {
	o.CspmEnabled = &v
}

// GetCustomMetricsEnabled returns the CustomMetricsEnabled field value if set, zero value otherwise.
func (o *AzureAccount) GetCustomMetricsEnabled() bool {
	if o == nil || o.CustomMetricsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CustomMetricsEnabled
}

// GetCustomMetricsEnabledOk returns a tuple with the CustomMetricsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetCustomMetricsEnabledOk() (*bool, bool) {
	if o == nil || o.CustomMetricsEnabled == nil {
		return nil, false
	}
	return o.CustomMetricsEnabled, true
}

// HasCustomMetricsEnabled returns a boolean if a field has been set.
func (o *AzureAccount) HasCustomMetricsEnabled() bool {
	return o != nil && o.CustomMetricsEnabled != nil
}

// SetCustomMetricsEnabled gets a reference to the given bool and assigns it to the CustomMetricsEnabled field.
func (o *AzureAccount) SetCustomMetricsEnabled(v bool) {
	o.CustomMetricsEnabled = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AzureAccount) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetErrorsOk() (*[]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AzureAccount) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AzureAccount) SetErrors(v []string) {
	o.Errors = v
}

// GetHostFilters returns the HostFilters field value if set, zero value otherwise.
func (o *AzureAccount) GetHostFilters() string {
	if o == nil || o.HostFilters == nil {
		var ret string
		return ret
	}
	return *o.HostFilters
}

// GetHostFiltersOk returns a tuple with the HostFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetHostFiltersOk() (*string, bool) {
	if o == nil || o.HostFilters == nil {
		return nil, false
	}
	return o.HostFilters, true
}

// HasHostFilters returns a boolean if a field has been set.
func (o *AzureAccount) HasHostFilters() bool {
	return o != nil && o.HostFilters != nil
}

// SetHostFilters gets a reference to the given string and assigns it to the HostFilters field.
func (o *AzureAccount) SetHostFilters(v string) {
	o.HostFilters = &v
}

// GetMetricsEnabled returns the MetricsEnabled field value if set, zero value otherwise.
func (o *AzureAccount) GetMetricsEnabled() bool {
	if o == nil || o.MetricsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MetricsEnabled
}

// GetMetricsEnabledOk returns a tuple with the MetricsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetMetricsEnabledOk() (*bool, bool) {
	if o == nil || o.MetricsEnabled == nil {
		return nil, false
	}
	return o.MetricsEnabled, true
}

// HasMetricsEnabled returns a boolean if a field has been set.
func (o *AzureAccount) HasMetricsEnabled() bool {
	return o != nil && o.MetricsEnabled != nil
}

// SetMetricsEnabled gets a reference to the given bool and assigns it to the MetricsEnabled field.
func (o *AzureAccount) SetMetricsEnabled(v bool) {
	o.MetricsEnabled = &v
}

// GetMetricsEnabledDefault returns the MetricsEnabledDefault field value if set, zero value otherwise.
func (o *AzureAccount) GetMetricsEnabledDefault() bool {
	if o == nil || o.MetricsEnabledDefault == nil {
		var ret bool
		return ret
	}
	return *o.MetricsEnabledDefault
}

// GetMetricsEnabledDefaultOk returns a tuple with the MetricsEnabledDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetMetricsEnabledDefaultOk() (*bool, bool) {
	if o == nil || o.MetricsEnabledDefault == nil {
		return nil, false
	}
	return o.MetricsEnabledDefault, true
}

// HasMetricsEnabledDefault returns a boolean if a field has been set.
func (o *AzureAccount) HasMetricsEnabledDefault() bool {
	return o != nil && o.MetricsEnabledDefault != nil
}

// SetMetricsEnabledDefault gets a reference to the given bool and assigns it to the MetricsEnabledDefault field.
func (o *AzureAccount) SetMetricsEnabledDefault(v bool) {
	o.MetricsEnabledDefault = &v
}

// GetNewClientId returns the NewClientId field value if set, zero value otherwise.
func (o *AzureAccount) GetNewClientId() string {
	if o == nil || o.NewClientId == nil {
		var ret string
		return ret
	}
	return *o.NewClientId
}

// GetNewClientIdOk returns a tuple with the NewClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetNewClientIdOk() (*string, bool) {
	if o == nil || o.NewClientId == nil {
		return nil, false
	}
	return o.NewClientId, true
}

// HasNewClientId returns a boolean if a field has been set.
func (o *AzureAccount) HasNewClientId() bool {
	return o != nil && o.NewClientId != nil
}

// SetNewClientId gets a reference to the given string and assigns it to the NewClientId field.
func (o *AzureAccount) SetNewClientId(v string) {
	o.NewClientId = &v
}

// GetNewTenantName returns the NewTenantName field value if set, zero value otherwise.
func (o *AzureAccount) GetNewTenantName() string {
	if o == nil || o.NewTenantName == nil {
		var ret string
		return ret
	}
	return *o.NewTenantName
}

// GetNewTenantNameOk returns a tuple with the NewTenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetNewTenantNameOk() (*string, bool) {
	if o == nil || o.NewTenantName == nil {
		return nil, false
	}
	return o.NewTenantName, true
}

// HasNewTenantName returns a boolean if a field has been set.
func (o *AzureAccount) HasNewTenantName() bool {
	return o != nil && o.NewTenantName != nil
}

// SetNewTenantName gets a reference to the given string and assigns it to the NewTenantName field.
func (o *AzureAccount) SetNewTenantName(v string) {
	o.NewTenantName = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *AzureAccount) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *AzureAccount) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *AzureAccount) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// GetResourceProviderConfigs returns the ResourceProviderConfigs field value if set, zero value otherwise.
func (o *AzureAccount) GetResourceProviderConfigs() []ResourceProviderConfig {
	if o == nil || o.ResourceProviderConfigs == nil {
		var ret []ResourceProviderConfig
		return ret
	}
	return o.ResourceProviderConfigs
}

// GetResourceProviderConfigsOk returns a tuple with the ResourceProviderConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetResourceProviderConfigsOk() (*[]ResourceProviderConfig, bool) {
	if o == nil || o.ResourceProviderConfigs == nil {
		return nil, false
	}
	return &o.ResourceProviderConfigs, true
}

// HasResourceProviderConfigs returns a boolean if a field has been set.
func (o *AzureAccount) HasResourceProviderConfigs() bool {
	return o != nil && o.ResourceProviderConfigs != nil
}

// SetResourceProviderConfigs gets a reference to the given []ResourceProviderConfig and assigns it to the ResourceProviderConfigs field.
func (o *AzureAccount) SetResourceProviderConfigs(v []ResourceProviderConfig) {
	o.ResourceProviderConfigs = v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *AzureAccount) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *AzureAccount) HasTenantName() bool {
	return o != nil && o.TenantName != nil
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *AzureAccount) SetTenantName(v string) {
	o.TenantName = &v
}

// GetUsageMetricsEnabled returns the UsageMetricsEnabled field value if set, zero value otherwise.
func (o *AzureAccount) GetUsageMetricsEnabled() bool {
	if o == nil || o.UsageMetricsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.UsageMetricsEnabled
}

// GetUsageMetricsEnabledOk returns a tuple with the UsageMetricsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureAccount) GetUsageMetricsEnabledOk() (*bool, bool) {
	if o == nil || o.UsageMetricsEnabled == nil {
		return nil, false
	}
	return o.UsageMetricsEnabled, true
}

// HasUsageMetricsEnabled returns a boolean if a field has been set.
func (o *AzureAccount) HasUsageMetricsEnabled() bool {
	return o != nil && o.UsageMetricsEnabled != nil
}

// SetUsageMetricsEnabled gets a reference to the given bool and assigns it to the UsageMetricsEnabled field.
func (o *AzureAccount) SetUsageMetricsEnabled(v bool) {
	o.UsageMetricsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppServicePlanFilters != nil {
		toSerialize["app_service_plan_filters"] = o.AppServicePlanFilters
	}
	if o.Automute != nil {
		toSerialize["automute"] = o.Automute
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.ContainerAppFilters != nil {
		toSerialize["container_app_filters"] = o.ContainerAppFilters
	}
	if o.CspmEnabled != nil {
		toSerialize["cspm_enabled"] = o.CspmEnabled
	}
	if o.CustomMetricsEnabled != nil {
		toSerialize["custom_metrics_enabled"] = o.CustomMetricsEnabled
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.HostFilters != nil {
		toSerialize["host_filters"] = o.HostFilters
	}
	if o.MetricsEnabled != nil {
		toSerialize["metrics_enabled"] = o.MetricsEnabled
	}
	if o.MetricsEnabledDefault != nil {
		toSerialize["metrics_enabled_default"] = o.MetricsEnabledDefault
	}
	if o.NewClientId != nil {
		toSerialize["new_client_id"] = o.NewClientId
	}
	if o.NewTenantName != nil {
		toSerialize["new_tenant_name"] = o.NewTenantName
	}
	if o.ResourceCollectionEnabled != nil {
		toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	}
	if o.ResourceProviderConfigs != nil {
		toSerialize["resource_provider_configs"] = o.ResourceProviderConfigs
	}
	if o.TenantName != nil {
		toSerialize["tenant_name"] = o.TenantName
	}
	if o.UsageMetricsEnabled != nil {
		toSerialize["usage_metrics_enabled"] = o.UsageMetricsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppServicePlanFilters     *string                  `json:"app_service_plan_filters,omitempty"`
		Automute                  *bool                    `json:"automute,omitempty"`
		ClientId                  *string                  `json:"client_id,omitempty"`
		ClientSecret              *string                  `json:"client_secret,omitempty"`
		ContainerAppFilters       *string                  `json:"container_app_filters,omitempty"`
		CspmEnabled               *bool                    `json:"cspm_enabled,omitempty"`
		CustomMetricsEnabled      *bool                    `json:"custom_metrics_enabled,omitempty"`
		Errors                    []string                 `json:"errors,omitempty"`
		HostFilters               *string                  `json:"host_filters,omitempty"`
		MetricsEnabled            *bool                    `json:"metrics_enabled,omitempty"`
		MetricsEnabledDefault     *bool                    `json:"metrics_enabled_default,omitempty"`
		NewClientId               *string                  `json:"new_client_id,omitempty"`
		NewTenantName             *string                  `json:"new_tenant_name,omitempty"`
		ResourceCollectionEnabled *bool                    `json:"resource_collection_enabled,omitempty"`
		ResourceProviderConfigs   []ResourceProviderConfig `json:"resource_provider_configs,omitempty"`
		TenantName                *string                  `json:"tenant_name,omitempty"`
		UsageMetricsEnabled       *bool                    `json:"usage_metrics_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_service_plan_filters", "automute", "client_id", "client_secret", "container_app_filters", "cspm_enabled", "custom_metrics_enabled", "errors", "host_filters", "metrics_enabled", "metrics_enabled_default", "new_client_id", "new_tenant_name", "resource_collection_enabled", "resource_provider_configs", "tenant_name", "usage_metrics_enabled"})
	} else {
		return err
	}
	o.AppServicePlanFilters = all.AppServicePlanFilters
	o.Automute = all.Automute
	o.ClientId = all.ClientId
	o.ClientSecret = all.ClientSecret
	o.ContainerAppFilters = all.ContainerAppFilters
	o.CspmEnabled = all.CspmEnabled
	o.CustomMetricsEnabled = all.CustomMetricsEnabled
	o.Errors = all.Errors
	o.HostFilters = all.HostFilters
	o.MetricsEnabled = all.MetricsEnabled
	o.MetricsEnabledDefault = all.MetricsEnabledDefault
	o.NewClientId = all.NewClientId
	o.NewTenantName = all.NewTenantName
	o.ResourceCollectionEnabled = all.ResourceCollectionEnabled
	o.ResourceProviderConfigs = all.ResourceProviderConfigs
	o.TenantName = all.TenantName
	o.UsageMetricsEnabled = all.UsageMetricsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
