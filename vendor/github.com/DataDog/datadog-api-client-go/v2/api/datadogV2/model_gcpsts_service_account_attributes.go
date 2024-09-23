// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPSTSServiceAccountAttributes Attributes associated with your service account.
type GCPSTSServiceAccountAttributes struct {
	// Tags to be associated with GCP metrics and service checks from your account.
	AccountTags []string `json:"account_tags,omitempty"`
	// Silence monitors for expected GCE instance shutdowns.
	Automute *bool `json:"automute,omitempty"`
	// Your service account email address.
	ClientEmail *string `json:"client_email,omitempty"`
	// List of filters to limit the Cloud Run revisions that are pulled into Datadog by using tags.
	// Only Cloud Run revision resources that apply to specified filters are imported into Datadog.
	CloudRunRevisionFilters []string `json:"cloud_run_revision_filters,omitempty"`
	// Your Host Filters.
	HostFilters []string `json:"host_filters,omitempty"`
	// When enabled, Datadog will activate the Cloud Security Monitoring product for this service account. Note: This requires resource_collection_enabled to be set to true.
	IsCspmEnabled *bool `json:"is_cspm_enabled,omitempty"`
	// When enabled, Datadog will attempt to collect Security Command Center Findings. Note: This requires additional permissions on the service account.
	IsSecurityCommandCenterEnabled *bool `json:"is_security_command_center_enabled,omitempty"`
	// When enabled, Datadog scans for all resources in your GCP environment.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPSTSServiceAccountAttributes instantiates a new GCPSTSServiceAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPSTSServiceAccountAttributes() *GCPSTSServiceAccountAttributes {
	this := GCPSTSServiceAccountAttributes{}
	var isSecurityCommandCenterEnabled bool = false
	this.IsSecurityCommandCenterEnabled = &isSecurityCommandCenterEnabled
	return &this
}

// NewGCPSTSServiceAccountAttributesWithDefaults instantiates a new GCPSTSServiceAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPSTSServiceAccountAttributesWithDefaults() *GCPSTSServiceAccountAttributes {
	this := GCPSTSServiceAccountAttributes{}
	var isSecurityCommandCenterEnabled bool = false
	this.IsSecurityCommandCenterEnabled = &isSecurityCommandCenterEnabled
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetAccountTags() []string {
	if o == nil || o.AccountTags == nil {
		var ret []string
		return ret
	}
	return o.AccountTags
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetAccountTagsOk() (*[]string, bool) {
	if o == nil || o.AccountTags == nil {
		return nil, false
	}
	return &o.AccountTags, true
}

// HasAccountTags returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasAccountTags() bool {
	return o != nil && o.AccountTags != nil
}

// SetAccountTags gets a reference to the given []string and assigns it to the AccountTags field.
func (o *GCPSTSServiceAccountAttributes) SetAccountTags(v []string) {
	o.AccountTags = v
}

// GetAutomute returns the Automute field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetAutomute() bool {
	if o == nil || o.Automute == nil {
		var ret bool
		return ret
	}
	return *o.Automute
}

// GetAutomuteOk returns a tuple with the Automute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetAutomuteOk() (*bool, bool) {
	if o == nil || o.Automute == nil {
		return nil, false
	}
	return o.Automute, true
}

// HasAutomute returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasAutomute() bool {
	return o != nil && o.Automute != nil
}

// SetAutomute gets a reference to the given bool and assigns it to the Automute field.
func (o *GCPSTSServiceAccountAttributes) SetAutomute(v bool) {
	o.Automute = &v
}

// GetClientEmail returns the ClientEmail field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetClientEmail() string {
	if o == nil || o.ClientEmail == nil {
		var ret string
		return ret
	}
	return *o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetClientEmailOk() (*string, bool) {
	if o == nil || o.ClientEmail == nil {
		return nil, false
	}
	return o.ClientEmail, true
}

// HasClientEmail returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasClientEmail() bool {
	return o != nil && o.ClientEmail != nil
}

// SetClientEmail gets a reference to the given string and assigns it to the ClientEmail field.
func (o *GCPSTSServiceAccountAttributes) SetClientEmail(v string) {
	o.ClientEmail = &v
}

// GetCloudRunRevisionFilters returns the CloudRunRevisionFilters field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetCloudRunRevisionFilters() []string {
	if o == nil || o.CloudRunRevisionFilters == nil {
		var ret []string
		return ret
	}
	return o.CloudRunRevisionFilters
}

// GetCloudRunRevisionFiltersOk returns a tuple with the CloudRunRevisionFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetCloudRunRevisionFiltersOk() (*[]string, bool) {
	if o == nil || o.CloudRunRevisionFilters == nil {
		return nil, false
	}
	return &o.CloudRunRevisionFilters, true
}

// HasCloudRunRevisionFilters returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasCloudRunRevisionFilters() bool {
	return o != nil && o.CloudRunRevisionFilters != nil
}

// SetCloudRunRevisionFilters gets a reference to the given []string and assigns it to the CloudRunRevisionFilters field.
func (o *GCPSTSServiceAccountAttributes) SetCloudRunRevisionFilters(v []string) {
	o.CloudRunRevisionFilters = v
}

// GetHostFilters returns the HostFilters field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetHostFilters() []string {
	if o == nil || o.HostFilters == nil {
		var ret []string
		return ret
	}
	return o.HostFilters
}

// GetHostFiltersOk returns a tuple with the HostFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetHostFiltersOk() (*[]string, bool) {
	if o == nil || o.HostFilters == nil {
		return nil, false
	}
	return &o.HostFilters, true
}

// HasHostFilters returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasHostFilters() bool {
	return o != nil && o.HostFilters != nil
}

// SetHostFilters gets a reference to the given []string and assigns it to the HostFilters field.
func (o *GCPSTSServiceAccountAttributes) SetHostFilters(v []string) {
	o.HostFilters = v
}

// GetIsCspmEnabled returns the IsCspmEnabled field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetIsCspmEnabled() bool {
	if o == nil || o.IsCspmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCspmEnabled
}

// GetIsCspmEnabledOk returns a tuple with the IsCspmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetIsCspmEnabledOk() (*bool, bool) {
	if o == nil || o.IsCspmEnabled == nil {
		return nil, false
	}
	return o.IsCspmEnabled, true
}

// HasIsCspmEnabled returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasIsCspmEnabled() bool {
	return o != nil && o.IsCspmEnabled != nil
}

// SetIsCspmEnabled gets a reference to the given bool and assigns it to the IsCspmEnabled field.
func (o *GCPSTSServiceAccountAttributes) SetIsCspmEnabled(v bool) {
	o.IsCspmEnabled = &v
}

// GetIsSecurityCommandCenterEnabled returns the IsSecurityCommandCenterEnabled field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetIsSecurityCommandCenterEnabled() bool {
	if o == nil || o.IsSecurityCommandCenterEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSecurityCommandCenterEnabled
}

// GetIsSecurityCommandCenterEnabledOk returns a tuple with the IsSecurityCommandCenterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetIsSecurityCommandCenterEnabledOk() (*bool, bool) {
	if o == nil || o.IsSecurityCommandCenterEnabled == nil {
		return nil, false
	}
	return o.IsSecurityCommandCenterEnabled, true
}

// HasIsSecurityCommandCenterEnabled returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasIsSecurityCommandCenterEnabled() bool {
	return o != nil && o.IsSecurityCommandCenterEnabled != nil
}

// SetIsSecurityCommandCenterEnabled gets a reference to the given bool and assigns it to the IsSecurityCommandCenterEnabled field.
func (o *GCPSTSServiceAccountAttributes) SetIsSecurityCommandCenterEnabled(v bool) {
	o.IsSecurityCommandCenterEnabled = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountAttributes) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountAttributes) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountAttributes) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *GCPSTSServiceAccountAttributes) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPSTSServiceAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountTags != nil {
		toSerialize["account_tags"] = o.AccountTags
	}
	if o.Automute != nil {
		toSerialize["automute"] = o.Automute
	}
	if o.ClientEmail != nil {
		toSerialize["client_email"] = o.ClientEmail
	}
	if o.CloudRunRevisionFilters != nil {
		toSerialize["cloud_run_revision_filters"] = o.CloudRunRevisionFilters
	}
	if o.HostFilters != nil {
		toSerialize["host_filters"] = o.HostFilters
	}
	if o.IsCspmEnabled != nil {
		toSerialize["is_cspm_enabled"] = o.IsCspmEnabled
	}
	if o.IsSecurityCommandCenterEnabled != nil {
		toSerialize["is_security_command_center_enabled"] = o.IsSecurityCommandCenterEnabled
	}
	if o.ResourceCollectionEnabled != nil {
		toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPSTSServiceAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags                    []string `json:"account_tags,omitempty"`
		Automute                       *bool    `json:"automute,omitempty"`
		ClientEmail                    *string  `json:"client_email,omitempty"`
		CloudRunRevisionFilters        []string `json:"cloud_run_revision_filters,omitempty"`
		HostFilters                    []string `json:"host_filters,omitempty"`
		IsCspmEnabled                  *bool    `json:"is_cspm_enabled,omitempty"`
		IsSecurityCommandCenterEnabled *bool    `json:"is_security_command_center_enabled,omitempty"`
		ResourceCollectionEnabled      *bool    `json:"resource_collection_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_tags", "automute", "client_email", "cloud_run_revision_filters", "host_filters", "is_cspm_enabled", "is_security_command_center_enabled", "resource_collection_enabled"})
	} else {
		return err
	}
	o.AccountTags = all.AccountTags
	o.Automute = all.Automute
	o.ClientEmail = all.ClientEmail
	o.CloudRunRevisionFilters = all.CloudRunRevisionFilters
	o.HostFilters = all.HostFilters
	o.IsCspmEnabled = all.IsCspmEnabled
	o.IsSecurityCommandCenterEnabled = all.IsSecurityCommandCenterEnabled
	o.ResourceCollectionEnabled = all.ResourceCollectionEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
