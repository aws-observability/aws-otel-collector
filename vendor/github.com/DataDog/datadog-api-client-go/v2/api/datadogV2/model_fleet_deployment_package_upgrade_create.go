// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentPackageUpgradeCreate Data for creating a new package upgrade deployment.
type FleetDeploymentPackageUpgradeCreate struct {
	// Attributes for creating a new package upgrade deployment.
	Attributes FleetDeploymentPackageUpgradeAttributes `json:"attributes"`
	// The type of deployment resource.
	Type FleetDeploymentResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentPackageUpgradeCreate instantiates a new FleetDeploymentPackageUpgradeCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentPackageUpgradeCreate(attributes FleetDeploymentPackageUpgradeAttributes, typeVar FleetDeploymentResourceType) *FleetDeploymentPackageUpgradeCreate {
	this := FleetDeploymentPackageUpgradeCreate{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewFleetDeploymentPackageUpgradeCreateWithDefaults instantiates a new FleetDeploymentPackageUpgradeCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentPackageUpgradeCreateWithDefaults() *FleetDeploymentPackageUpgradeCreate {
	this := FleetDeploymentPackageUpgradeCreate{}
	var typeVar FleetDeploymentResourceType = FLEETDEPLOYMENTRESOURCETYPE_DEPLOYMENT
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *FleetDeploymentPackageUpgradeCreate) GetAttributes() FleetDeploymentPackageUpgradeAttributes {
	if o == nil {
		var ret FleetDeploymentPackageUpgradeAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentPackageUpgradeCreate) GetAttributesOk() (*FleetDeploymentPackageUpgradeAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *FleetDeploymentPackageUpgradeCreate) SetAttributes(v FleetDeploymentPackageUpgradeAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *FleetDeploymentPackageUpgradeCreate) GetType() FleetDeploymentResourceType {
	if o == nil {
		var ret FleetDeploymentResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentPackageUpgradeCreate) GetTypeOk() (*FleetDeploymentResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FleetDeploymentPackageUpgradeCreate) SetType(v FleetDeploymentResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentPackageUpgradeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentPackageUpgradeCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *FleetDeploymentPackageUpgradeAttributes `json:"attributes"`
		Type       *FleetDeploymentResourceType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
