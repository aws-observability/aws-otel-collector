// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupRelationships Relationships of an org group.
type OrgGroupRelationships struct {
	// Relationship to org group memberships.
	Memberships *OrgGroupMembershipsRelationship `json:"memberships,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupRelationships instantiates a new OrgGroupRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupRelationships() *OrgGroupRelationships {
	this := OrgGroupRelationships{}
	return &this
}

// NewOrgGroupRelationshipsWithDefaults instantiates a new OrgGroupRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupRelationshipsWithDefaults() *OrgGroupRelationships {
	this := OrgGroupRelationships{}
	return &this
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *OrgGroupRelationships) GetMemberships() OrgGroupMembershipsRelationship {
	if o == nil || o.Memberships == nil {
		var ret OrgGroupMembershipsRelationship
		return ret
	}
	return *o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupRelationships) GetMembershipsOk() (*OrgGroupMembershipsRelationship, bool) {
	if o == nil || o.Memberships == nil {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *OrgGroupRelationships) HasMemberships() bool {
	return o != nil && o.Memberships != nil
}

// SetMemberships gets a reference to the given OrgGroupMembershipsRelationship and assigns it to the Memberships field.
func (o *OrgGroupRelationships) SetMemberships(v OrgGroupMembershipsRelationship) {
	o.Memberships = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Memberships != nil {
		toSerialize["memberships"] = o.Memberships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Memberships *OrgGroupMembershipsRelationship `json:"memberships,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"memberships"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Memberships != nil && all.Memberships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Memberships = all.Memberships

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
