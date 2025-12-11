// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPAccount Your Google Cloud Platform Account.
type GCPAccount struct {
	// Should be `https://www.googleapis.com/oauth2/v1/certs`.
	AuthProviderX509CertUrl *string `json:"auth_provider_x509_cert_url,omitempty"`
	// Should be `https://accounts.google.com/o/oauth2/auth`.
	AuthUri *string `json:"auth_uri,omitempty"`
	// Silence monitors for expected GCE instance shutdowns.
	Automute *bool `json:"automute,omitempty"`
	// Your email found in your JSON service account key.
	ClientEmail *string `json:"client_email,omitempty"`
	// Your ID found in your JSON service account key.
	ClientId *string `json:"client_id,omitempty"`
	// Should be `https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL`
	// where `$CLIENT_EMAIL` is the email found in your JSON service account key.
	ClientX509CertUrl *string `json:"client_x509_cert_url,omitempty"`
	// List of filters to limit the Cloud Run revisions that are pulled into Datadog by using tags.
	// Only Cloud Run revision resources that apply to specified filters are imported into Datadog.
	// **Note:** This field is deprecated. Instead, use `monitored_resource_configs` with `type=cloud_run_revision`
	// Deprecated
	CloudRunRevisionFilters []string `json:"cloud_run_revision_filters,omitempty"`
	// An array of errors.
	Errors []string `json:"errors,omitempty"`
	// A comma-separated list of filters to limit the VM instances that are pulled into Datadog by using tags.
	// Only VM instance resources that apply to specified filters are imported into Datadog.
	// **Note:** This field is deprecated. Instead, use `monitored_resource_configs` with `type=gce_instance`
	// Deprecated
	HostFilters *string `json:"host_filters,omitempty"`
	// When enabled, Datadog will activate the Cloud Security Monitoring product for this service account. Note: This requires resource_collection_enabled to be set to true.
	IsCspmEnabled *bool `json:"is_cspm_enabled,omitempty"`
	// When enabled, Datadog scans for all resource change data in your Google Cloud environment.
	IsResourceChangeCollectionEnabled *bool `json:"is_resource_change_collection_enabled,omitempty"`
	// When enabled, Datadog will attempt to collect Security Command Center Findings. Note: This requires additional permissions on the service account.
	IsSecurityCommandCenterEnabled *bool `json:"is_security_command_center_enabled,omitempty"`
	// Configurations for GCP monitored resources.
	MonitoredResourceConfigs []GCPMonitoredResourceConfig `json:"monitored_resource_configs,omitempty"`
	// Your private key name found in your JSON service account key.
	PrivateKey *string `json:"private_key,omitempty"`
	// Your private key ID found in your JSON service account key.
	PrivateKeyId *string `json:"private_key_id,omitempty"`
	// Your Google Cloud project ID found in your JSON service account key.
	ProjectId *string `json:"project_id,omitempty"`
	// When enabled, Datadog scans for all resources in your GCP environment.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// Should be `https://accounts.google.com/o/oauth2/token`.
	TokenUri *string `json:"token_uri,omitempty"`
	// The value for service_account found in your JSON service account key.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPAccount instantiates a new GCPAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPAccount() *GCPAccount {
	this := GCPAccount{}
	var isResourceChangeCollectionEnabled bool = false
	this.IsResourceChangeCollectionEnabled = &isResourceChangeCollectionEnabled
	var isSecurityCommandCenterEnabled bool = false
	this.IsSecurityCommandCenterEnabled = &isSecurityCommandCenterEnabled
	return &this
}

// NewGCPAccountWithDefaults instantiates a new GCPAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPAccountWithDefaults() *GCPAccount {
	this := GCPAccount{}
	var isResourceChangeCollectionEnabled bool = false
	this.IsResourceChangeCollectionEnabled = &isResourceChangeCollectionEnabled
	var isSecurityCommandCenterEnabled bool = false
	this.IsSecurityCommandCenterEnabled = &isSecurityCommandCenterEnabled
	return &this
}

// GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field value if set, zero value otherwise.
func (o *GCPAccount) GetAuthProviderX509CertUrl() string {
	if o == nil || o.AuthProviderX509CertUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthProviderX509CertUrl
}

// GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetAuthProviderX509CertUrlOk() (*string, bool) {
	if o == nil || o.AuthProviderX509CertUrl == nil {
		return nil, false
	}
	return o.AuthProviderX509CertUrl, true
}

// HasAuthProviderX509CertUrl returns a boolean if a field has been set.
func (o *GCPAccount) HasAuthProviderX509CertUrl() bool {
	return o != nil && o.AuthProviderX509CertUrl != nil
}

// SetAuthProviderX509CertUrl gets a reference to the given string and assigns it to the AuthProviderX509CertUrl field.
func (o *GCPAccount) SetAuthProviderX509CertUrl(v string) {
	o.AuthProviderX509CertUrl = &v
}

// GetAuthUri returns the AuthUri field value if set, zero value otherwise.
func (o *GCPAccount) GetAuthUri() string {
	if o == nil || o.AuthUri == nil {
		var ret string
		return ret
	}
	return *o.AuthUri
}

// GetAuthUriOk returns a tuple with the AuthUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetAuthUriOk() (*string, bool) {
	if o == nil || o.AuthUri == nil {
		return nil, false
	}
	return o.AuthUri, true
}

// HasAuthUri returns a boolean if a field has been set.
func (o *GCPAccount) HasAuthUri() bool {
	return o != nil && o.AuthUri != nil
}

// SetAuthUri gets a reference to the given string and assigns it to the AuthUri field.
func (o *GCPAccount) SetAuthUri(v string) {
	o.AuthUri = &v
}

// GetAutomute returns the Automute field value if set, zero value otherwise.
func (o *GCPAccount) GetAutomute() bool {
	if o == nil || o.Automute == nil {
		var ret bool
		return ret
	}
	return *o.Automute
}

// GetAutomuteOk returns a tuple with the Automute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetAutomuteOk() (*bool, bool) {
	if o == nil || o.Automute == nil {
		return nil, false
	}
	return o.Automute, true
}

// HasAutomute returns a boolean if a field has been set.
func (o *GCPAccount) HasAutomute() bool {
	return o != nil && o.Automute != nil
}

// SetAutomute gets a reference to the given bool and assigns it to the Automute field.
func (o *GCPAccount) SetAutomute(v bool) {
	o.Automute = &v
}

// GetClientEmail returns the ClientEmail field value if set, zero value otherwise.
func (o *GCPAccount) GetClientEmail() string {
	if o == nil || o.ClientEmail == nil {
		var ret string
		return ret
	}
	return *o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetClientEmailOk() (*string, bool) {
	if o == nil || o.ClientEmail == nil {
		return nil, false
	}
	return o.ClientEmail, true
}

// HasClientEmail returns a boolean if a field has been set.
func (o *GCPAccount) HasClientEmail() bool {
	return o != nil && o.ClientEmail != nil
}

// SetClientEmail gets a reference to the given string and assigns it to the ClientEmail field.
func (o *GCPAccount) SetClientEmail(v string) {
	o.ClientEmail = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GCPAccount) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GCPAccount) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GCPAccount) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientX509CertUrl returns the ClientX509CertUrl field value if set, zero value otherwise.
func (o *GCPAccount) GetClientX509CertUrl() string {
	if o == nil || o.ClientX509CertUrl == nil {
		var ret string
		return ret
	}
	return *o.ClientX509CertUrl
}

// GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetClientX509CertUrlOk() (*string, bool) {
	if o == nil || o.ClientX509CertUrl == nil {
		return nil, false
	}
	return o.ClientX509CertUrl, true
}

// HasClientX509CertUrl returns a boolean if a field has been set.
func (o *GCPAccount) HasClientX509CertUrl() bool {
	return o != nil && o.ClientX509CertUrl != nil
}

// SetClientX509CertUrl gets a reference to the given string and assigns it to the ClientX509CertUrl field.
func (o *GCPAccount) SetClientX509CertUrl(v string) {
	o.ClientX509CertUrl = &v
}

// GetCloudRunRevisionFilters returns the CloudRunRevisionFilters field value if set, zero value otherwise.
// Deprecated
func (o *GCPAccount) GetCloudRunRevisionFilters() []string {
	if o == nil || o.CloudRunRevisionFilters == nil {
		var ret []string
		return ret
	}
	return o.CloudRunRevisionFilters
}

// GetCloudRunRevisionFiltersOk returns a tuple with the CloudRunRevisionFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GCPAccount) GetCloudRunRevisionFiltersOk() (*[]string, bool) {
	if o == nil || o.CloudRunRevisionFilters == nil {
		return nil, false
	}
	return &o.CloudRunRevisionFilters, true
}

// HasCloudRunRevisionFilters returns a boolean if a field has been set.
func (o *GCPAccount) HasCloudRunRevisionFilters() bool {
	return o != nil && o.CloudRunRevisionFilters != nil
}

// SetCloudRunRevisionFilters gets a reference to the given []string and assigns it to the CloudRunRevisionFilters field.
// Deprecated
func (o *GCPAccount) SetCloudRunRevisionFilters(v []string) {
	o.CloudRunRevisionFilters = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GCPAccount) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetErrorsOk() (*[]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GCPAccount) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GCPAccount) SetErrors(v []string) {
	o.Errors = v
}

// GetHostFilters returns the HostFilters field value if set, zero value otherwise.
// Deprecated
func (o *GCPAccount) GetHostFilters() string {
	if o == nil || o.HostFilters == nil {
		var ret string
		return ret
	}
	return *o.HostFilters
}

// GetHostFiltersOk returns a tuple with the HostFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GCPAccount) GetHostFiltersOk() (*string, bool) {
	if o == nil || o.HostFilters == nil {
		return nil, false
	}
	return o.HostFilters, true
}

// HasHostFilters returns a boolean if a field has been set.
func (o *GCPAccount) HasHostFilters() bool {
	return o != nil && o.HostFilters != nil
}

// SetHostFilters gets a reference to the given string and assigns it to the HostFilters field.
// Deprecated
func (o *GCPAccount) SetHostFilters(v string) {
	o.HostFilters = &v
}

// GetIsCspmEnabled returns the IsCspmEnabled field value if set, zero value otherwise.
func (o *GCPAccount) GetIsCspmEnabled() bool {
	if o == nil || o.IsCspmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCspmEnabled
}

// GetIsCspmEnabledOk returns a tuple with the IsCspmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetIsCspmEnabledOk() (*bool, bool) {
	if o == nil || o.IsCspmEnabled == nil {
		return nil, false
	}
	return o.IsCspmEnabled, true
}

// HasIsCspmEnabled returns a boolean if a field has been set.
func (o *GCPAccount) HasIsCspmEnabled() bool {
	return o != nil && o.IsCspmEnabled != nil
}

// SetIsCspmEnabled gets a reference to the given bool and assigns it to the IsCspmEnabled field.
func (o *GCPAccount) SetIsCspmEnabled(v bool) {
	o.IsCspmEnabled = &v
}

// GetIsResourceChangeCollectionEnabled returns the IsResourceChangeCollectionEnabled field value if set, zero value otherwise.
func (o *GCPAccount) GetIsResourceChangeCollectionEnabled() bool {
	if o == nil || o.IsResourceChangeCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsResourceChangeCollectionEnabled
}

// GetIsResourceChangeCollectionEnabledOk returns a tuple with the IsResourceChangeCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetIsResourceChangeCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.IsResourceChangeCollectionEnabled == nil {
		return nil, false
	}
	return o.IsResourceChangeCollectionEnabled, true
}

// HasIsResourceChangeCollectionEnabled returns a boolean if a field has been set.
func (o *GCPAccount) HasIsResourceChangeCollectionEnabled() bool {
	return o != nil && o.IsResourceChangeCollectionEnabled != nil
}

// SetIsResourceChangeCollectionEnabled gets a reference to the given bool and assigns it to the IsResourceChangeCollectionEnabled field.
func (o *GCPAccount) SetIsResourceChangeCollectionEnabled(v bool) {
	o.IsResourceChangeCollectionEnabled = &v
}

// GetIsSecurityCommandCenterEnabled returns the IsSecurityCommandCenterEnabled field value if set, zero value otherwise.
func (o *GCPAccount) GetIsSecurityCommandCenterEnabled() bool {
	if o == nil || o.IsSecurityCommandCenterEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSecurityCommandCenterEnabled
}

// GetIsSecurityCommandCenterEnabledOk returns a tuple with the IsSecurityCommandCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetIsSecurityCommandCenterEnabledOk() (*bool, bool) {
	if o == nil || o.IsSecurityCommandCenterEnabled == nil {
		return nil, false
	}
	return o.IsSecurityCommandCenterEnabled, true
}

// HasIsSecurityCommandCenterEnabled returns a boolean if a field has been set.
func (o *GCPAccount) HasIsSecurityCommandCenterEnabled() bool {
	return o != nil && o.IsSecurityCommandCenterEnabled != nil
}

// SetIsSecurityCommandCenterEnabled gets a reference to the given bool and assigns it to the IsSecurityCommandCenterEnabled field.
func (o *GCPAccount) SetIsSecurityCommandCenterEnabled(v bool) {
	o.IsSecurityCommandCenterEnabled = &v
}

// GetMonitoredResourceConfigs returns the MonitoredResourceConfigs field value if set, zero value otherwise.
func (o *GCPAccount) GetMonitoredResourceConfigs() []GCPMonitoredResourceConfig {
	if o == nil || o.MonitoredResourceConfigs == nil {
		var ret []GCPMonitoredResourceConfig
		return ret
	}
	return o.MonitoredResourceConfigs
}

// GetMonitoredResourceConfigsOk returns a tuple with the MonitoredResourceConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetMonitoredResourceConfigsOk() (*[]GCPMonitoredResourceConfig, bool) {
	if o == nil || o.MonitoredResourceConfigs == nil {
		return nil, false
	}
	return &o.MonitoredResourceConfigs, true
}

// HasMonitoredResourceConfigs returns a boolean if a field has been set.
func (o *GCPAccount) HasMonitoredResourceConfigs() bool {
	return o != nil && o.MonitoredResourceConfigs != nil
}

// SetMonitoredResourceConfigs gets a reference to the given []GCPMonitoredResourceConfig and assigns it to the MonitoredResourceConfigs field.
func (o *GCPAccount) SetMonitoredResourceConfigs(v []GCPMonitoredResourceConfig) {
	o.MonitoredResourceConfigs = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *GCPAccount) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GCPAccount) HasPrivateKey() bool {
	return o != nil && o.PrivateKey != nil
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *GCPAccount) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPrivateKeyId returns the PrivateKeyId field value if set, zero value otherwise.
func (o *GCPAccount) GetPrivateKeyId() string {
	if o == nil || o.PrivateKeyId == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyId
}

// GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetPrivateKeyIdOk() (*string, bool) {
	if o == nil || o.PrivateKeyId == nil {
		return nil, false
	}
	return o.PrivateKeyId, true
}

// HasPrivateKeyId returns a boolean if a field has been set.
func (o *GCPAccount) HasPrivateKeyId() bool {
	return o != nil && o.PrivateKeyId != nil
}

// SetPrivateKeyId gets a reference to the given string and assigns it to the PrivateKeyId field.
func (o *GCPAccount) SetPrivateKeyId(v string) {
	o.PrivateKeyId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GCPAccount) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GCPAccount) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GCPAccount) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *GCPAccount) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *GCPAccount) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *GCPAccount) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// GetTokenUri returns the TokenUri field value if set, zero value otherwise.
func (o *GCPAccount) GetTokenUri() string {
	if o == nil || o.TokenUri == nil {
		var ret string
		return ret
	}
	return *o.TokenUri
}

// GetTokenUriOk returns a tuple with the TokenUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetTokenUriOk() (*string, bool) {
	if o == nil || o.TokenUri == nil {
		return nil, false
	}
	return o.TokenUri, true
}

// HasTokenUri returns a boolean if a field has been set.
func (o *GCPAccount) HasTokenUri() bool {
	return o != nil && o.TokenUri != nil
}

// SetTokenUri gets a reference to the given string and assigns it to the TokenUri field.
func (o *GCPAccount) SetTokenUri(v string) {
	o.TokenUri = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GCPAccount) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPAccount) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GCPAccount) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GCPAccount) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthProviderX509CertUrl != nil {
		toSerialize["auth_provider_x509_cert_url"] = o.AuthProviderX509CertUrl
	}
	if o.AuthUri != nil {
		toSerialize["auth_uri"] = o.AuthUri
	}
	if o.Automute != nil {
		toSerialize["automute"] = o.Automute
	}
	if o.ClientEmail != nil {
		toSerialize["client_email"] = o.ClientEmail
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientX509CertUrl != nil {
		toSerialize["client_x509_cert_url"] = o.ClientX509CertUrl
	}
	if o.CloudRunRevisionFilters != nil {
		toSerialize["cloud_run_revision_filters"] = o.CloudRunRevisionFilters
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.HostFilters != nil {
		toSerialize["host_filters"] = o.HostFilters
	}
	if o.IsCspmEnabled != nil {
		toSerialize["is_cspm_enabled"] = o.IsCspmEnabled
	}
	if o.IsResourceChangeCollectionEnabled != nil {
		toSerialize["is_resource_change_collection_enabled"] = o.IsResourceChangeCollectionEnabled
	}
	if o.IsSecurityCommandCenterEnabled != nil {
		toSerialize["is_security_command_center_enabled"] = o.IsSecurityCommandCenterEnabled
	}
	if o.MonitoredResourceConfigs != nil {
		toSerialize["monitored_resource_configs"] = o.MonitoredResourceConfigs
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.PrivateKeyId != nil {
		toSerialize["private_key_id"] = o.PrivateKeyId
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.ResourceCollectionEnabled != nil {
		toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	}
	if o.TokenUri != nil {
		toSerialize["token_uri"] = o.TokenUri
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthProviderX509CertUrl           *string                      `json:"auth_provider_x509_cert_url,omitempty"`
		AuthUri                           *string                      `json:"auth_uri,omitempty"`
		Automute                          *bool                        `json:"automute,omitempty"`
		ClientEmail                       *string                      `json:"client_email,omitempty"`
		ClientId                          *string                      `json:"client_id,omitempty"`
		ClientX509CertUrl                 *string                      `json:"client_x509_cert_url,omitempty"`
		CloudRunRevisionFilters           []string                     `json:"cloud_run_revision_filters,omitempty"`
		Errors                            []string                     `json:"errors,omitempty"`
		HostFilters                       *string                      `json:"host_filters,omitempty"`
		IsCspmEnabled                     *bool                        `json:"is_cspm_enabled,omitempty"`
		IsResourceChangeCollectionEnabled *bool                        `json:"is_resource_change_collection_enabled,omitempty"`
		IsSecurityCommandCenterEnabled    *bool                        `json:"is_security_command_center_enabled,omitempty"`
		MonitoredResourceConfigs          []GCPMonitoredResourceConfig `json:"monitored_resource_configs,omitempty"`
		PrivateKey                        *string                      `json:"private_key,omitempty"`
		PrivateKeyId                      *string                      `json:"private_key_id,omitempty"`
		ProjectId                         *string                      `json:"project_id,omitempty"`
		ResourceCollectionEnabled         *bool                        `json:"resource_collection_enabled,omitempty"`
		TokenUri                          *string                      `json:"token_uri,omitempty"`
		Type                              *string                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_provider_x509_cert_url", "auth_uri", "automute", "client_email", "client_id", "client_x509_cert_url", "cloud_run_revision_filters", "errors", "host_filters", "is_cspm_enabled", "is_resource_change_collection_enabled", "is_security_command_center_enabled", "monitored_resource_configs", "private_key", "private_key_id", "project_id", "resource_collection_enabled", "token_uri", "type"})
	} else {
		return err
	}
	o.AuthProviderX509CertUrl = all.AuthProviderX509CertUrl
	o.AuthUri = all.AuthUri
	o.Automute = all.Automute
	o.ClientEmail = all.ClientEmail
	o.ClientId = all.ClientId
	o.ClientX509CertUrl = all.ClientX509CertUrl
	o.CloudRunRevisionFilters = all.CloudRunRevisionFilters
	o.Errors = all.Errors
	o.HostFilters = all.HostFilters
	o.IsCspmEnabled = all.IsCspmEnabled
	o.IsResourceChangeCollectionEnabled = all.IsResourceChangeCollectionEnabled
	o.IsSecurityCommandCenterEnabled = all.IsSecurityCommandCenterEnabled
	o.MonitoredResourceConfigs = all.MonitoredResourceConfigs
	o.PrivateKey = all.PrivateKey
	o.PrivateKeyId = all.PrivateKeyId
	o.ProjectId = all.ProjectId
	o.ResourceCollectionEnabled = all.ResourceCollectionEnabled
	o.TokenUri = all.TokenUri
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
