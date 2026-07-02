// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManagedOrgsRelationships Relationships of the managed organizations resource.
type ManagedOrgsRelationships struct {
	// Relationship to the current organization.
	CurrentOrg ManagedOrgsRelationshipToOrg `json:"current_org"`
	// Relationship to the managed organizations.
	ManagedOrgs ManagedOrgsRelationshipToOrgs `json:"managed_orgs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewManagedOrgsRelationships instantiates a new ManagedOrgsRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewManagedOrgsRelationships(currentOrg ManagedOrgsRelationshipToOrg, managedOrgs ManagedOrgsRelationshipToOrgs) *ManagedOrgsRelationships {
	this := ManagedOrgsRelationships{}
	this.CurrentOrg = currentOrg
	this.ManagedOrgs = managedOrgs
	return &this
}

// NewManagedOrgsRelationshipsWithDefaults instantiates a new ManagedOrgsRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewManagedOrgsRelationshipsWithDefaults() *ManagedOrgsRelationships {
	this := ManagedOrgsRelationships{}
	return &this
}

// GetCurrentOrg returns the CurrentOrg field value.
func (o *ManagedOrgsRelationships) GetCurrentOrg() ManagedOrgsRelationshipToOrg {
	if o == nil {
		var ret ManagedOrgsRelationshipToOrg
		return ret
	}
	return o.CurrentOrg
}

// GetCurrentOrgOk returns a tuple with the CurrentOrg field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsRelationships) GetCurrentOrgOk() (*ManagedOrgsRelationshipToOrg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentOrg, true
}

// SetCurrentOrg sets field value.
func (o *ManagedOrgsRelationships) SetCurrentOrg(v ManagedOrgsRelationshipToOrg) {
	o.CurrentOrg = v
}

// GetManagedOrgs returns the ManagedOrgs field value.
func (o *ManagedOrgsRelationships) GetManagedOrgs() ManagedOrgsRelationshipToOrgs {
	if o == nil {
		var ret ManagedOrgsRelationshipToOrgs
		return ret
	}
	return o.ManagedOrgs
}

// GetManagedOrgsOk returns a tuple with the ManagedOrgs field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsRelationships) GetManagedOrgsOk() (*ManagedOrgsRelationshipToOrgs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagedOrgs, true
}

// SetManagedOrgs sets field value.
func (o *ManagedOrgsRelationships) SetManagedOrgs(v ManagedOrgsRelationshipToOrgs) {
	o.ManagedOrgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ManagedOrgsRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["current_org"] = o.CurrentOrg
	toSerialize["managed_orgs"] = o.ManagedOrgs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ManagedOrgsRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentOrg  *ManagedOrgsRelationshipToOrg  `json:"current_org"`
		ManagedOrgs *ManagedOrgsRelationshipToOrgs `json:"managed_orgs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CurrentOrg == nil {
		return fmt.Errorf("required field current_org missing")
	}
	if all.ManagedOrgs == nil {
		return fmt.Errorf("required field managed_orgs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current_org", "managed_orgs"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CurrentOrg.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CurrentOrg = *all.CurrentOrg
	if all.ManagedOrgs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ManagedOrgs = *all.ManagedOrgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
