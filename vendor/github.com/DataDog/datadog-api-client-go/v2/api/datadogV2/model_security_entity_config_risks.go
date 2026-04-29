// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityConfigRisks Configuration risks associated with the entity
type SecurityEntityConfigRisks struct {
	// Whether the entity has identity risks
	HasIdentityRisk bool `json:"hasIdentityRisk"`
	// Whether the entity has misconfigurations
	HasMisconfiguration bool `json:"hasMisconfiguration"`
	// Whether the entity has privileged roles
	HasPrivilegedRole bool `json:"hasPrivilegedRole"`
	// Whether the entity has privileged access
	IsPrivileged bool `json:"isPrivileged"`
	// Whether the entity is in a production environment
	IsProduction bool `json:"isProduction"`
	// Whether the entity is publicly accessible
	IsPubliclyAccessible bool `json:"isPubliclyAccessible"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityEntityConfigRisks instantiates a new SecurityEntityConfigRisks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityEntityConfigRisks(hasIdentityRisk bool, hasMisconfiguration bool, hasPrivilegedRole bool, isPrivileged bool, isProduction bool, isPubliclyAccessible bool) *SecurityEntityConfigRisks {
	this := SecurityEntityConfigRisks{}
	this.HasIdentityRisk = hasIdentityRisk
	this.HasMisconfiguration = hasMisconfiguration
	this.HasPrivilegedRole = hasPrivilegedRole
	this.IsPrivileged = isPrivileged
	this.IsProduction = isProduction
	this.IsPubliclyAccessible = isPubliclyAccessible
	return &this
}

// NewSecurityEntityConfigRisksWithDefaults instantiates a new SecurityEntityConfigRisks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityEntityConfigRisksWithDefaults() *SecurityEntityConfigRisks {
	this := SecurityEntityConfigRisks{}
	return &this
}

// GetHasIdentityRisk returns the HasIdentityRisk field value.
func (o *SecurityEntityConfigRisks) GetHasIdentityRisk() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasIdentityRisk
}

// GetHasIdentityRiskOk returns a tuple with the HasIdentityRisk field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityConfigRisks) GetHasIdentityRiskOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasIdentityRisk, true
}

// SetHasIdentityRisk sets field value.
func (o *SecurityEntityConfigRisks) SetHasIdentityRisk(v bool) {
	o.HasIdentityRisk = v
}

// GetHasMisconfiguration returns the HasMisconfiguration field value.
func (o *SecurityEntityConfigRisks) GetHasMisconfiguration() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasMisconfiguration
}

// GetHasMisconfigurationOk returns a tuple with the HasMisconfiguration field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityConfigRisks) GetHasMisconfigurationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMisconfiguration, true
}

// SetHasMisconfiguration sets field value.
func (o *SecurityEntityConfigRisks) SetHasMisconfiguration(v bool) {
	o.HasMisconfiguration = v
}

// GetHasPrivilegedRole returns the HasPrivilegedRole field value.
func (o *SecurityEntityConfigRisks) GetHasPrivilegedRole() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasPrivilegedRole
}

// GetHasPrivilegedRoleOk returns a tuple with the HasPrivilegedRole field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityConfigRisks) GetHasPrivilegedRoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrivilegedRole, true
}

// SetHasPrivilegedRole sets field value.
func (o *SecurityEntityConfigRisks) SetHasPrivilegedRole(v bool) {
	o.HasPrivilegedRole = v
}

// GetIsPrivileged returns the IsPrivileged field value.
func (o *SecurityEntityConfigRisks) GetIsPrivileged() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPrivileged
}

// GetIsPrivilegedOk returns a tuple with the IsPrivileged field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityConfigRisks) GetIsPrivilegedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrivileged, true
}

// SetIsPrivileged sets field value.
func (o *SecurityEntityConfigRisks) SetIsPrivileged(v bool) {
	o.IsPrivileged = v
}

// GetIsProduction returns the IsProduction field value.
func (o *SecurityEntityConfigRisks) GetIsProduction() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityConfigRisks) GetIsProductionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsProduction, true
}

// SetIsProduction sets field value.
func (o *SecurityEntityConfigRisks) SetIsProduction(v bool) {
	o.IsProduction = v
}

// GetIsPubliclyAccessible returns the IsPubliclyAccessible field value.
func (o *SecurityEntityConfigRisks) GetIsPubliclyAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPubliclyAccessible
}

// GetIsPubliclyAccessibleOk returns a tuple with the IsPubliclyAccessible field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityConfigRisks) GetIsPubliclyAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPubliclyAccessible, true
}

// SetIsPubliclyAccessible sets field value.
func (o *SecurityEntityConfigRisks) SetIsPubliclyAccessible(v bool) {
	o.IsPubliclyAccessible = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityEntityConfigRisks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["hasIdentityRisk"] = o.HasIdentityRisk
	toSerialize["hasMisconfiguration"] = o.HasMisconfiguration
	toSerialize["hasPrivilegedRole"] = o.HasPrivilegedRole
	toSerialize["isPrivileged"] = o.IsPrivileged
	toSerialize["isProduction"] = o.IsProduction
	toSerialize["isPubliclyAccessible"] = o.IsPubliclyAccessible

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityEntityConfigRisks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasIdentityRisk      *bool `json:"hasIdentityRisk"`
		HasMisconfiguration  *bool `json:"hasMisconfiguration"`
		HasPrivilegedRole    *bool `json:"hasPrivilegedRole"`
		IsPrivileged         *bool `json:"isPrivileged"`
		IsProduction         *bool `json:"isProduction"`
		IsPubliclyAccessible *bool `json:"isPubliclyAccessible"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HasIdentityRisk == nil {
		return fmt.Errorf("required field hasIdentityRisk missing")
	}
	if all.HasMisconfiguration == nil {
		return fmt.Errorf("required field hasMisconfiguration missing")
	}
	if all.HasPrivilegedRole == nil {
		return fmt.Errorf("required field hasPrivilegedRole missing")
	}
	if all.IsPrivileged == nil {
		return fmt.Errorf("required field isPrivileged missing")
	}
	if all.IsProduction == nil {
		return fmt.Errorf("required field isProduction missing")
	}
	if all.IsPubliclyAccessible == nil {
		return fmt.Errorf("required field isPubliclyAccessible missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hasIdentityRisk", "hasMisconfiguration", "hasPrivilegedRole", "isPrivileged", "isProduction", "isPubliclyAccessible"})
	} else {
		return err
	}
	o.HasIdentityRisk = *all.HasIdentityRisk
	o.HasMisconfiguration = *all.HasMisconfiguration
	o.HasPrivilegedRole = *all.HasPrivilegedRole
	o.IsPrivileged = *all.IsPrivileged
	o.IsProduction = *all.IsProduction
	o.IsPubliclyAccessible = *all.IsPubliclyAccessible

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
