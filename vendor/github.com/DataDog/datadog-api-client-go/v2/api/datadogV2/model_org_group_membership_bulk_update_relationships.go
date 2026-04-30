// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipBulkUpdateRelationships Relationships for bulk updating memberships.
type OrgGroupMembershipBulkUpdateRelationships struct {
	// Relationship to a single org group.
	SourceOrgGroup OrgGroupRelationshipToOne `json:"source_org_group"`
	// Relationship to a single org group.
	TargetOrgGroup OrgGroupRelationshipToOne `json:"target_org_group"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupMembershipBulkUpdateRelationships instantiates a new OrgGroupMembershipBulkUpdateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupMembershipBulkUpdateRelationships(sourceOrgGroup OrgGroupRelationshipToOne, targetOrgGroup OrgGroupRelationshipToOne) *OrgGroupMembershipBulkUpdateRelationships {
	this := OrgGroupMembershipBulkUpdateRelationships{}
	this.SourceOrgGroup = sourceOrgGroup
	this.TargetOrgGroup = targetOrgGroup
	return &this
}

// NewOrgGroupMembershipBulkUpdateRelationshipsWithDefaults instantiates a new OrgGroupMembershipBulkUpdateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupMembershipBulkUpdateRelationshipsWithDefaults() *OrgGroupMembershipBulkUpdateRelationships {
	this := OrgGroupMembershipBulkUpdateRelationships{}
	return &this
}

// GetSourceOrgGroup returns the SourceOrgGroup field value.
func (o *OrgGroupMembershipBulkUpdateRelationships) GetSourceOrgGroup() OrgGroupRelationshipToOne {
	if o == nil {
		var ret OrgGroupRelationshipToOne
		return ret
	}
	return o.SourceOrgGroup
}

// GetSourceOrgGroupOk returns a tuple with the SourceOrgGroup field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipBulkUpdateRelationships) GetSourceOrgGroupOk() (*OrgGroupRelationshipToOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceOrgGroup, true
}

// SetSourceOrgGroup sets field value.
func (o *OrgGroupMembershipBulkUpdateRelationships) SetSourceOrgGroup(v OrgGroupRelationshipToOne) {
	o.SourceOrgGroup = v
}

// GetTargetOrgGroup returns the TargetOrgGroup field value.
func (o *OrgGroupMembershipBulkUpdateRelationships) GetTargetOrgGroup() OrgGroupRelationshipToOne {
	if o == nil {
		var ret OrgGroupRelationshipToOne
		return ret
	}
	return o.TargetOrgGroup
}

// GetTargetOrgGroupOk returns a tuple with the TargetOrgGroup field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipBulkUpdateRelationships) GetTargetOrgGroupOk() (*OrgGroupRelationshipToOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetOrgGroup, true
}

// SetTargetOrgGroup sets field value.
func (o *OrgGroupMembershipBulkUpdateRelationships) SetTargetOrgGroup(v OrgGroupRelationshipToOne) {
	o.TargetOrgGroup = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupMembershipBulkUpdateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["source_org_group"] = o.SourceOrgGroup
	toSerialize["target_org_group"] = o.TargetOrgGroup

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupMembershipBulkUpdateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SourceOrgGroup *OrgGroupRelationshipToOne `json:"source_org_group"`
		TargetOrgGroup *OrgGroupRelationshipToOne `json:"target_org_group"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SourceOrgGroup == nil {
		return fmt.Errorf("required field source_org_group missing")
	}
	if all.TargetOrgGroup == nil {
		return fmt.Errorf("required field target_org_group missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"source_org_group", "target_org_group"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SourceOrgGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SourceOrgGroup = *all.SourceOrgGroup
	if all.TargetOrgGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TargetOrgGroup = *all.TargetOrgGroup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
