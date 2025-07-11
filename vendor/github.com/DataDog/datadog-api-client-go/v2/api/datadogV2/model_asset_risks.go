// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssetRisks Asset risks.
type AssetRisks struct {
	// Whether the asset has access to sensitive data or not.
	HasAccessToSensitiveData *bool `json:"has_access_to_sensitive_data,omitempty"`
	// Whether the asset has privileged access or not.
	HasPrivilegedAccess *bool `json:"has_privileged_access,omitempty"`
	// Whether the asset is in production or not.
	InProduction bool `json:"in_production"`
	// Whether the asset is publicly accessible or not.
	IsPubliclyAccessible *bool `json:"is_publicly_accessible,omitempty"`
	// Whether the asset is under attack or not.
	UnderAttack *bool `json:"under_attack,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssetRisks instantiates a new AssetRisks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetRisks(inProduction bool) *AssetRisks {
	this := AssetRisks{}
	this.InProduction = inProduction
	return &this
}

// NewAssetRisksWithDefaults instantiates a new AssetRisks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetRisksWithDefaults() *AssetRisks {
	this := AssetRisks{}
	return &this
}

// GetHasAccessToSensitiveData returns the HasAccessToSensitiveData field value if set, zero value otherwise.
func (o *AssetRisks) GetHasAccessToSensitiveData() bool {
	if o == nil || o.HasAccessToSensitiveData == nil {
		var ret bool
		return ret
	}
	return *o.HasAccessToSensitiveData
}

// GetHasAccessToSensitiveDataOk returns a tuple with the HasAccessToSensitiveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetRisks) GetHasAccessToSensitiveDataOk() (*bool, bool) {
	if o == nil || o.HasAccessToSensitiveData == nil {
		return nil, false
	}
	return o.HasAccessToSensitiveData, true
}

// HasHasAccessToSensitiveData returns a boolean if a field has been set.
func (o *AssetRisks) HasHasAccessToSensitiveData() bool {
	return o != nil && o.HasAccessToSensitiveData != nil
}

// SetHasAccessToSensitiveData gets a reference to the given bool and assigns it to the HasAccessToSensitiveData field.
func (o *AssetRisks) SetHasAccessToSensitiveData(v bool) {
	o.HasAccessToSensitiveData = &v
}

// GetHasPrivilegedAccess returns the HasPrivilegedAccess field value if set, zero value otherwise.
func (o *AssetRisks) GetHasPrivilegedAccess() bool {
	if o == nil || o.HasPrivilegedAccess == nil {
		var ret bool
		return ret
	}
	return *o.HasPrivilegedAccess
}

// GetHasPrivilegedAccessOk returns a tuple with the HasPrivilegedAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetRisks) GetHasPrivilegedAccessOk() (*bool, bool) {
	if o == nil || o.HasPrivilegedAccess == nil {
		return nil, false
	}
	return o.HasPrivilegedAccess, true
}

// HasHasPrivilegedAccess returns a boolean if a field has been set.
func (o *AssetRisks) HasHasPrivilegedAccess() bool {
	return o != nil && o.HasPrivilegedAccess != nil
}

// SetHasPrivilegedAccess gets a reference to the given bool and assigns it to the HasPrivilegedAccess field.
func (o *AssetRisks) SetHasPrivilegedAccess(v bool) {
	o.HasPrivilegedAccess = &v
}

// GetInProduction returns the InProduction field value.
func (o *AssetRisks) GetInProduction() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.InProduction
}

// GetInProductionOk returns a tuple with the InProduction field value
// and a boolean to check if the value has been set.
func (o *AssetRisks) GetInProductionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InProduction, true
}

// SetInProduction sets field value.
func (o *AssetRisks) SetInProduction(v bool) {
	o.InProduction = v
}

// GetIsPubliclyAccessible returns the IsPubliclyAccessible field value if set, zero value otherwise.
func (o *AssetRisks) GetIsPubliclyAccessible() bool {
	if o == nil || o.IsPubliclyAccessible == nil {
		var ret bool
		return ret
	}
	return *o.IsPubliclyAccessible
}

// GetIsPubliclyAccessibleOk returns a tuple with the IsPubliclyAccessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetRisks) GetIsPubliclyAccessibleOk() (*bool, bool) {
	if o == nil || o.IsPubliclyAccessible == nil {
		return nil, false
	}
	return o.IsPubliclyAccessible, true
}

// HasIsPubliclyAccessible returns a boolean if a field has been set.
func (o *AssetRisks) HasIsPubliclyAccessible() bool {
	return o != nil && o.IsPubliclyAccessible != nil
}

// SetIsPubliclyAccessible gets a reference to the given bool and assigns it to the IsPubliclyAccessible field.
func (o *AssetRisks) SetIsPubliclyAccessible(v bool) {
	o.IsPubliclyAccessible = &v
}

// GetUnderAttack returns the UnderAttack field value if set, zero value otherwise.
func (o *AssetRisks) GetUnderAttack() bool {
	if o == nil || o.UnderAttack == nil {
		var ret bool
		return ret
	}
	return *o.UnderAttack
}

// GetUnderAttackOk returns a tuple with the UnderAttack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetRisks) GetUnderAttackOk() (*bool, bool) {
	if o == nil || o.UnderAttack == nil {
		return nil, false
	}
	return o.UnderAttack, true
}

// HasUnderAttack returns a boolean if a field has been set.
func (o *AssetRisks) HasUnderAttack() bool {
	return o != nil && o.UnderAttack != nil
}

// SetUnderAttack gets a reference to the given bool and assigns it to the UnderAttack field.
func (o *AssetRisks) SetUnderAttack(v bool) {
	o.UnderAttack = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetRisks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HasAccessToSensitiveData != nil {
		toSerialize["has_access_to_sensitive_data"] = o.HasAccessToSensitiveData
	}
	if o.HasPrivilegedAccess != nil {
		toSerialize["has_privileged_access"] = o.HasPrivilegedAccess
	}
	toSerialize["in_production"] = o.InProduction
	if o.IsPubliclyAccessible != nil {
		toSerialize["is_publicly_accessible"] = o.IsPubliclyAccessible
	}
	if o.UnderAttack != nil {
		toSerialize["under_attack"] = o.UnderAttack
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssetRisks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasAccessToSensitiveData *bool `json:"has_access_to_sensitive_data,omitempty"`
		HasPrivilegedAccess      *bool `json:"has_privileged_access,omitempty"`
		InProduction             *bool `json:"in_production"`
		IsPubliclyAccessible     *bool `json:"is_publicly_accessible,omitempty"`
		UnderAttack              *bool `json:"under_attack,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InProduction == nil {
		return fmt.Errorf("required field in_production missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_access_to_sensitive_data", "has_privileged_access", "in_production", "is_publicly_accessible", "under_attack"})
	} else {
		return err
	}
	o.HasAccessToSensitiveData = all.HasAccessToSensitiveData
	o.HasPrivilegedAccess = all.HasPrivilegedAccess
	o.InProduction = *all.InProduction
	o.IsPubliclyAccessible = all.IsPubliclyAccessible
	o.UnderAttack = all.UnderAttack

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
