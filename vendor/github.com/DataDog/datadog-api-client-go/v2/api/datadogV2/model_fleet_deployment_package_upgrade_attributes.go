// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentPackageUpgradeAttributes Attributes for creating a new package upgrade deployment.
type FleetDeploymentPackageUpgradeAttributes struct {
	// Query used to filter and select target hosts for the deployment. Uses the Datadog query syntax.
	FilterQuery *string `json:"filter_query,omitempty"`
	// List of packages and their target versions to deploy to the selected hosts.
	TargetPackages []FleetDeploymentPackage `json:"target_packages"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentPackageUpgradeAttributes instantiates a new FleetDeploymentPackageUpgradeAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentPackageUpgradeAttributes(targetPackages []FleetDeploymentPackage) *FleetDeploymentPackageUpgradeAttributes {
	this := FleetDeploymentPackageUpgradeAttributes{}
	this.TargetPackages = targetPackages
	return &this
}

// NewFleetDeploymentPackageUpgradeAttributesWithDefaults instantiates a new FleetDeploymentPackageUpgradeAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentPackageUpgradeAttributesWithDefaults() *FleetDeploymentPackageUpgradeAttributes {
	this := FleetDeploymentPackageUpgradeAttributes{}
	return &this
}

// GetFilterQuery returns the FilterQuery field value if set, zero value otherwise.
func (o *FleetDeploymentPackageUpgradeAttributes) GetFilterQuery() string {
	if o == nil || o.FilterQuery == nil {
		var ret string
		return ret
	}
	return *o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentPackageUpgradeAttributes) GetFilterQueryOk() (*string, bool) {
	if o == nil || o.FilterQuery == nil {
		return nil, false
	}
	return o.FilterQuery, true
}

// HasFilterQuery returns a boolean if a field has been set.
func (o *FleetDeploymentPackageUpgradeAttributes) HasFilterQuery() bool {
	return o != nil && o.FilterQuery != nil
}

// SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.
func (o *FleetDeploymentPackageUpgradeAttributes) SetFilterQuery(v string) {
	o.FilterQuery = &v
}

// GetTargetPackages returns the TargetPackages field value.
func (o *FleetDeploymentPackageUpgradeAttributes) GetTargetPackages() []FleetDeploymentPackage {
	if o == nil {
		var ret []FleetDeploymentPackage
		return ret
	}
	return o.TargetPackages
}

// GetTargetPackagesOk returns a tuple with the TargetPackages field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentPackageUpgradeAttributes) GetTargetPackagesOk() (*[]FleetDeploymentPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPackages, true
}

// SetTargetPackages sets field value.
func (o *FleetDeploymentPackageUpgradeAttributes) SetTargetPackages(v []FleetDeploymentPackage) {
	o.TargetPackages = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentPackageUpgradeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FilterQuery != nil {
		toSerialize["filter_query"] = o.FilterQuery
	}
	toSerialize["target_packages"] = o.TargetPackages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentPackageUpgradeAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilterQuery    *string                   `json:"filter_query,omitempty"`
		TargetPackages *[]FleetDeploymentPackage `json:"target_packages"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TargetPackages == nil {
		return fmt.Errorf("required field target_packages missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter_query", "target_packages"})
	} else {
		return err
	}
	o.FilterQuery = all.FilterQuery
	o.TargetPackages = *all.TargetPackages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
