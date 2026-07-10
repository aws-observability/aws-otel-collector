// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmAgentlessHostAttributes Attributes of an agentless host.
type CsmAgentlessHostAttributes struct {
	// The ID of the cloud account that the host belongs to.
	AccountId string `json:"account_id"`
	// The cloud provider of a host resource.
	CloudProvider CsmCloudProvider `json:"cloud_provider"`
	// Whether CSM Misconfigurations is enabled for this host. `true` if enabled; `false` if disabled.
	HasPostureManagement bool `json:"has_posture_management"`
	// Whether CSM Vulnerabilities is enabled for this host. `true` if enabled; `false` if disabled.
	HasVulnerabilityScanning bool `json:"has_vulnerability_scanning"`
	// The type of cloud resource for an agentless host.
	ResourceType CsmAgentlessHostResourceType `json:"resource_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmAgentlessHostAttributes instantiates a new CsmAgentlessHostAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmAgentlessHostAttributes(accountId string, cloudProvider CsmCloudProvider, hasPostureManagement bool, hasVulnerabilityScanning bool, resourceType CsmAgentlessHostResourceType) *CsmAgentlessHostAttributes {
	this := CsmAgentlessHostAttributes{}
	this.AccountId = accountId
	this.CloudProvider = cloudProvider
	this.HasPostureManagement = hasPostureManagement
	this.HasVulnerabilityScanning = hasVulnerabilityScanning
	this.ResourceType = resourceType
	return &this
}

// NewCsmAgentlessHostAttributesWithDefaults instantiates a new CsmAgentlessHostAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmAgentlessHostAttributesWithDefaults() *CsmAgentlessHostAttributes {
	this := CsmAgentlessHostAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *CsmAgentlessHostAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *CsmAgentlessHostAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetCloudProvider returns the CloudProvider field value.
func (o *CsmAgentlessHostAttributes) GetCloudProvider() CsmCloudProvider {
	if o == nil {
		var ret CsmCloudProvider
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostAttributes) GetCloudProviderOk() (*CsmCloudProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *CsmAgentlessHostAttributes) SetCloudProvider(v CsmCloudProvider) {
	o.CloudProvider = v
}

// GetHasPostureManagement returns the HasPostureManagement field value.
func (o *CsmAgentlessHostAttributes) GetHasPostureManagement() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasPostureManagement
}

// GetHasPostureManagementOk returns a tuple with the HasPostureManagement field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostAttributes) GetHasPostureManagementOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPostureManagement, true
}

// SetHasPostureManagement sets field value.
func (o *CsmAgentlessHostAttributes) SetHasPostureManagement(v bool) {
	o.HasPostureManagement = v
}

// GetHasVulnerabilityScanning returns the HasVulnerabilityScanning field value.
func (o *CsmAgentlessHostAttributes) GetHasVulnerabilityScanning() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasVulnerabilityScanning
}

// GetHasVulnerabilityScanningOk returns a tuple with the HasVulnerabilityScanning field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostAttributes) GetHasVulnerabilityScanningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasVulnerabilityScanning, true
}

// SetHasVulnerabilityScanning sets field value.
func (o *CsmAgentlessHostAttributes) SetHasVulnerabilityScanning(v bool) {
	o.HasVulnerabilityScanning = v
}

// GetResourceType returns the ResourceType field value.
func (o *CsmAgentlessHostAttributes) GetResourceType() CsmAgentlessHostResourceType {
	if o == nil {
		var ret CsmAgentlessHostResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostAttributes) GetResourceTypeOk() (*CsmAgentlessHostResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *CsmAgentlessHostAttributes) SetResourceType(v CsmAgentlessHostResourceType) {
	o.ResourceType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmAgentlessHostAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["has_posture_management"] = o.HasPostureManagement
	toSerialize["has_vulnerability_scanning"] = o.HasVulnerabilityScanning
	toSerialize["resource_type"] = o.ResourceType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmAgentlessHostAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId                *string                       `json:"account_id"`
		CloudProvider            *CsmCloudProvider             `json:"cloud_provider"`
		HasPostureManagement     *bool                         `json:"has_posture_management"`
		HasVulnerabilityScanning *bool                         `json:"has_vulnerability_scanning"`
		ResourceType             *CsmAgentlessHostResourceType `json:"resource_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloud_provider missing")
	}
	if all.HasPostureManagement == nil {
		return fmt.Errorf("required field has_posture_management missing")
	}
	if all.HasVulnerabilityScanning == nil {
		return fmt.Errorf("required field has_vulnerability_scanning missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "cloud_provider", "has_posture_management", "has_vulnerability_scanning", "resource_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = *all.AccountId
	if !all.CloudProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.CloudProvider = *all.CloudProvider
	}
	o.HasPostureManagement = *all.HasPostureManagement
	o.HasVulnerabilityScanning = *all.HasVulnerabilityScanning
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
