// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyOverrideRelationships Relationships of an org group policy override.
type OrgGroupPolicyOverrideRelationships struct {
	// Relationship to a single org group.
	OrgGroup *OrgGroupRelationshipToOne `json:"org_group,omitempty"`
	// Relationship to a single org group policy.
	OrgGroupPolicy *OrgGroupPolicyRelationshipToOne `json:"org_group_policy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyOverrideRelationships instantiates a new OrgGroupPolicyOverrideRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyOverrideRelationships() *OrgGroupPolicyOverrideRelationships {
	this := OrgGroupPolicyOverrideRelationships{}
	return &this
}

// NewOrgGroupPolicyOverrideRelationshipsWithDefaults instantiates a new OrgGroupPolicyOverrideRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyOverrideRelationshipsWithDefaults() *OrgGroupPolicyOverrideRelationships {
	this := OrgGroupPolicyOverrideRelationships{}
	return &this
}

// GetOrgGroup returns the OrgGroup field value if set, zero value otherwise.
func (o *OrgGroupPolicyOverrideRelationships) GetOrgGroup() OrgGroupRelationshipToOne {
	if o == nil || o.OrgGroup == nil {
		var ret OrgGroupRelationshipToOne
		return ret
	}
	return *o.OrgGroup
}

// GetOrgGroupOk returns a tuple with the OrgGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideRelationships) GetOrgGroupOk() (*OrgGroupRelationshipToOne, bool) {
	if o == nil || o.OrgGroup == nil {
		return nil, false
	}
	return o.OrgGroup, true
}

// HasOrgGroup returns a boolean if a field has been set.
func (o *OrgGroupPolicyOverrideRelationships) HasOrgGroup() bool {
	return o != nil && o.OrgGroup != nil
}

// SetOrgGroup gets a reference to the given OrgGroupRelationshipToOne and assigns it to the OrgGroup field.
func (o *OrgGroupPolicyOverrideRelationships) SetOrgGroup(v OrgGroupRelationshipToOne) {
	o.OrgGroup = &v
}

// GetOrgGroupPolicy returns the OrgGroupPolicy field value if set, zero value otherwise.
func (o *OrgGroupPolicyOverrideRelationships) GetOrgGroupPolicy() OrgGroupPolicyRelationshipToOne {
	if o == nil || o.OrgGroupPolicy == nil {
		var ret OrgGroupPolicyRelationshipToOne
		return ret
	}
	return *o.OrgGroupPolicy
}

// GetOrgGroupPolicyOk returns a tuple with the OrgGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideRelationships) GetOrgGroupPolicyOk() (*OrgGroupPolicyRelationshipToOne, bool) {
	if o == nil || o.OrgGroupPolicy == nil {
		return nil, false
	}
	return o.OrgGroupPolicy, true
}

// HasOrgGroupPolicy returns a boolean if a field has been set.
func (o *OrgGroupPolicyOverrideRelationships) HasOrgGroupPolicy() bool {
	return o != nil && o.OrgGroupPolicy != nil
}

// SetOrgGroupPolicy gets a reference to the given OrgGroupPolicyRelationshipToOne and assigns it to the OrgGroupPolicy field.
func (o *OrgGroupPolicyOverrideRelationships) SetOrgGroupPolicy(v OrgGroupPolicyRelationshipToOne) {
	o.OrgGroupPolicy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyOverrideRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OrgGroup != nil {
		toSerialize["org_group"] = o.OrgGroup
	}
	if o.OrgGroupPolicy != nil {
		toSerialize["org_group_policy"] = o.OrgGroupPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyOverrideRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgGroup       *OrgGroupRelationshipToOne       `json:"org_group,omitempty"`
		OrgGroupPolicy *OrgGroupPolicyRelationshipToOne `json:"org_group_policy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_group", "org_group_policy"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OrgGroup != nil && all.OrgGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgGroup = all.OrgGroup
	if all.OrgGroupPolicy != nil && all.OrgGroupPolicy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgGroupPolicy = all.OrgGroupPolicy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
