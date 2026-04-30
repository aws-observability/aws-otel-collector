// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipAttributes Attributes of an org group membership.
type OrgGroupMembershipAttributes struct {
	// Timestamp when the membership was created.
	CreatedAt time.Time `json:"created_at"`
	// Timestamp when the membership was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the member organization.
	OrgName string `json:"org_name"`
	// The site of the member organization.
	OrgSite string `json:"org_site"`
	// The UUID of the member organization.
	OrgUuid uuid.UUID `json:"org_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupMembershipAttributes instantiates a new OrgGroupMembershipAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupMembershipAttributes(createdAt time.Time, modifiedAt time.Time, orgName string, orgSite string, orgUuid uuid.UUID) *OrgGroupMembershipAttributes {
	this := OrgGroupMembershipAttributes{}
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.OrgName = orgName
	this.OrgSite = orgSite
	this.OrgUuid = orgUuid
	return &this
}

// NewOrgGroupMembershipAttributesWithDefaults instantiates a new OrgGroupMembershipAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupMembershipAttributesWithDefaults() *OrgGroupMembershipAttributes {
	this := OrgGroupMembershipAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OrgGroupMembershipAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OrgGroupMembershipAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *OrgGroupMembershipAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *OrgGroupMembershipAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetOrgName returns the OrgName field value.
func (o *OrgGroupMembershipAttributes) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipAttributes) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *OrgGroupMembershipAttributes) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgSite returns the OrgSite field value.
func (o *OrgGroupMembershipAttributes) GetOrgSite() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgSite
}

// GetOrgSiteOk returns a tuple with the OrgSite field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipAttributes) GetOrgSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSite, true
}

// SetOrgSite sets field value.
func (o *OrgGroupMembershipAttributes) SetOrgSite(v string) {
	o.OrgSite = v
}

// GetOrgUuid returns the OrgUuid field value.
func (o *OrgGroupMembershipAttributes) GetOrgUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.OrgUuid
}

// GetOrgUuidOk returns a tuple with the OrgUuid field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipAttributes) GetOrgUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgUuid, true
}

// SetOrgUuid sets field value.
func (o *OrgGroupMembershipAttributes) SetOrgUuid(v uuid.UUID) {
	o.OrgUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupMembershipAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["org_name"] = o.OrgName
	toSerialize["org_site"] = o.OrgSite
	toSerialize["org_uuid"] = o.OrgUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupMembershipAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time `json:"created_at"`
		ModifiedAt *time.Time `json:"modified_at"`
		OrgName    *string    `json:"org_name"`
		OrgSite    *string    `json:"org_site"`
		OrgUuid    *uuid.UUID `json:"org_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field org_name missing")
	}
	if all.OrgSite == nil {
		return fmt.Errorf("required field org_site missing")
	}
	if all.OrgUuid == nil {
		return fmt.Errorf("required field org_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "org_name", "org_site", "org_uuid"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ModifiedAt = *all.ModifiedAt
	o.OrgName = *all.OrgName
	o.OrgSite = *all.OrgSite
	o.OrgUuid = *all.OrgUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
