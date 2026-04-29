// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyOverrideCreateRelationships Relationships for creating a policy override.
type OrgGroupPolicyOverrideCreateRelationships struct {
	// Relationship to a single org group.
	OrgGroup OrgGroupRelationshipToOne `json:"org_group"`
	// Relationship to a single org group policy.
	OrgGroupPolicy OrgGroupPolicyRelationshipToOne `json:"org_group_policy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyOverrideCreateRelationships instantiates a new OrgGroupPolicyOverrideCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyOverrideCreateRelationships(orgGroup OrgGroupRelationshipToOne, orgGroupPolicy OrgGroupPolicyRelationshipToOne) *OrgGroupPolicyOverrideCreateRelationships {
	this := OrgGroupPolicyOverrideCreateRelationships{}
	this.OrgGroup = orgGroup
	this.OrgGroupPolicy = orgGroupPolicy
	return &this
}

// NewOrgGroupPolicyOverrideCreateRelationshipsWithDefaults instantiates a new OrgGroupPolicyOverrideCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyOverrideCreateRelationshipsWithDefaults() *OrgGroupPolicyOverrideCreateRelationships {
	this := OrgGroupPolicyOverrideCreateRelationships{}
	return &this
}

// GetOrgGroup returns the OrgGroup field value.
func (o *OrgGroupPolicyOverrideCreateRelationships) GetOrgGroup() OrgGroupRelationshipToOne {
	if o == nil {
		var ret OrgGroupRelationshipToOne
		return ret
	}
	return o.OrgGroup
}

// GetOrgGroupOk returns a tuple with the OrgGroup field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideCreateRelationships) GetOrgGroupOk() (*OrgGroupRelationshipToOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgGroup, true
}

// SetOrgGroup sets field value.
func (o *OrgGroupPolicyOverrideCreateRelationships) SetOrgGroup(v OrgGroupRelationshipToOne) {
	o.OrgGroup = v
}

// GetOrgGroupPolicy returns the OrgGroupPolicy field value.
func (o *OrgGroupPolicyOverrideCreateRelationships) GetOrgGroupPolicy() OrgGroupPolicyRelationshipToOne {
	if o == nil {
		var ret OrgGroupPolicyRelationshipToOne
		return ret
	}
	return o.OrgGroupPolicy
}

// GetOrgGroupPolicyOk returns a tuple with the OrgGroupPolicy field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideCreateRelationships) GetOrgGroupPolicyOk() (*OrgGroupPolicyRelationshipToOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgGroupPolicy, true
}

// SetOrgGroupPolicy sets field value.
func (o *OrgGroupPolicyOverrideCreateRelationships) SetOrgGroupPolicy(v OrgGroupPolicyRelationshipToOne) {
	o.OrgGroupPolicy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyOverrideCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["org_group"] = o.OrgGroup
	toSerialize["org_group_policy"] = o.OrgGroupPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyOverrideCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgGroup       *OrgGroupRelationshipToOne       `json:"org_group"`
		OrgGroupPolicy *OrgGroupPolicyRelationshipToOne `json:"org_group_policy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OrgGroup == nil {
		return fmt.Errorf("required field org_group missing")
	}
	if all.OrgGroupPolicy == nil {
		return fmt.Errorf("required field org_group_policy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_group", "org_group_policy"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OrgGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgGroup = *all.OrgGroup
	if all.OrgGroupPolicy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgGroupPolicy = *all.OrgGroupPolicy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
