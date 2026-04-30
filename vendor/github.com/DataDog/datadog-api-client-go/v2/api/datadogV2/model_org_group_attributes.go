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

// OrgGroupAttributes Attributes of an org group.
type OrgGroupAttributes struct {
	// Timestamp when the org group was created.
	CreatedAt time.Time `json:"created_at"`
	// Timestamp when the org group was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the org group.
	Name string `json:"name"`
	// The site of the organization that owns this org group.
	OwnerOrgSite string `json:"owner_org_site"`
	// The UUID of the organization that owns this org group.
	OwnerOrgUuid uuid.UUID `json:"owner_org_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupAttributes instantiates a new OrgGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupAttributes(createdAt time.Time, modifiedAt time.Time, name string, ownerOrgSite string, ownerOrgUuid uuid.UUID) *OrgGroupAttributes {
	this := OrgGroupAttributes{}
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.OwnerOrgSite = ownerOrgSite
	this.OwnerOrgUuid = ownerOrgUuid
	return &this
}

// NewOrgGroupAttributesWithDefaults instantiates a new OrgGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupAttributesWithDefaults() *OrgGroupAttributes {
	this := OrgGroupAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OrgGroupAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OrgGroupAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *OrgGroupAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *OrgGroupAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *OrgGroupAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrgGroupAttributes) SetName(v string) {
	o.Name = v
}

// GetOwnerOrgSite returns the OwnerOrgSite field value.
func (o *OrgGroupAttributes) GetOwnerOrgSite() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerOrgSite
}

// GetOwnerOrgSiteOk returns a tuple with the OwnerOrgSite field value
// and a boolean to check if the value has been set.
func (o *OrgGroupAttributes) GetOwnerOrgSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrgSite, true
}

// SetOwnerOrgSite sets field value.
func (o *OrgGroupAttributes) SetOwnerOrgSite(v string) {
	o.OwnerOrgSite = v
}

// GetOwnerOrgUuid returns the OwnerOrgUuid field value.
func (o *OrgGroupAttributes) GetOwnerOrgUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.OwnerOrgUuid
}

// GetOwnerOrgUuidOk returns a tuple with the OwnerOrgUuid field value
// and a boolean to check if the value has been set.
func (o *OrgGroupAttributes) GetOwnerOrgUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerOrgUuid, true
}

// SetOwnerOrgUuid sets field value.
func (o *OrgGroupAttributes) SetOwnerOrgUuid(v uuid.UUID) {
	o.OwnerOrgUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name
	toSerialize["owner_org_site"] = o.OwnerOrgSite
	toSerialize["owner_org_uuid"] = o.OwnerOrgUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *time.Time `json:"created_at"`
		ModifiedAt   *time.Time `json:"modified_at"`
		Name         *string    `json:"name"`
		OwnerOrgSite *string    `json:"owner_org_site"`
		OwnerOrgUuid *uuid.UUID `json:"owner_org_uuid"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OwnerOrgSite == nil {
		return fmt.Errorf("required field owner_org_site missing")
	}
	if all.OwnerOrgUuid == nil {
		return fmt.Errorf("required field owner_org_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "name", "owner_org_site", "owner_org_uuid"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.OwnerOrgSite = *all.OwnerOrgSite
	o.OwnerOrgUuid = *all.OwnerOrgUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
