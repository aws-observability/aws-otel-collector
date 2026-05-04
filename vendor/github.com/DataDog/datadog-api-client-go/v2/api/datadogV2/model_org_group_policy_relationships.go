// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyRelationships Relationships of an org group policy.
type OrgGroupPolicyRelationships struct {
	// Relationship to a single org group.
	OrgGroup *OrgGroupRelationshipToOne `json:"org_group,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyRelationships instantiates a new OrgGroupPolicyRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyRelationships() *OrgGroupPolicyRelationships {
	this := OrgGroupPolicyRelationships{}
	return &this
}

// NewOrgGroupPolicyRelationshipsWithDefaults instantiates a new OrgGroupPolicyRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyRelationshipsWithDefaults() *OrgGroupPolicyRelationships {
	this := OrgGroupPolicyRelationships{}
	return &this
}

// GetOrgGroup returns the OrgGroup field value if set, zero value otherwise.
func (o *OrgGroupPolicyRelationships) GetOrgGroup() OrgGroupRelationshipToOne {
	if o == nil || o.OrgGroup == nil {
		var ret OrgGroupRelationshipToOne
		return ret
	}
	return *o.OrgGroup
}

// GetOrgGroupOk returns a tuple with the OrgGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyRelationships) GetOrgGroupOk() (*OrgGroupRelationshipToOne, bool) {
	if o == nil || o.OrgGroup == nil {
		return nil, false
	}
	return o.OrgGroup, true
}

// HasOrgGroup returns a boolean if a field has been set.
func (o *OrgGroupPolicyRelationships) HasOrgGroup() bool {
	return o != nil && o.OrgGroup != nil
}

// SetOrgGroup gets a reference to the given OrgGroupRelationshipToOne and assigns it to the OrgGroup field.
func (o *OrgGroupPolicyRelationships) SetOrgGroup(v OrgGroupRelationshipToOne) {
	o.OrgGroup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OrgGroup != nil {
		toSerialize["org_group"] = o.OrgGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgGroup *OrgGroupRelationshipToOne `json:"org_group,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_group"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OrgGroup != nil && all.OrgGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgGroup = all.OrgGroup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
