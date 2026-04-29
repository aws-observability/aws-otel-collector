// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyOverrideUpdateAttributes Attributes for updating a policy override. The `org_uuid` and `org_site` fields must match the existing override and cannot be changed.
type OrgGroupPolicyOverrideUpdateAttributes struct {
	// The site of the organization.
	OrgSite string `json:"org_site"`
	// The UUID of the organization.
	OrgUuid uuid.UUID `json:"org_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyOverrideUpdateAttributes instantiates a new OrgGroupPolicyOverrideUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyOverrideUpdateAttributes(orgSite string, orgUuid uuid.UUID) *OrgGroupPolicyOverrideUpdateAttributes {
	this := OrgGroupPolicyOverrideUpdateAttributes{}
	this.OrgSite = orgSite
	this.OrgUuid = orgUuid
	return &this
}

// NewOrgGroupPolicyOverrideUpdateAttributesWithDefaults instantiates a new OrgGroupPolicyOverrideUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyOverrideUpdateAttributesWithDefaults() *OrgGroupPolicyOverrideUpdateAttributes {
	this := OrgGroupPolicyOverrideUpdateAttributes{}
	return &this
}

// GetOrgSite returns the OrgSite field value.
func (o *OrgGroupPolicyOverrideUpdateAttributes) GetOrgSite() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgSite
}

// GetOrgSiteOk returns a tuple with the OrgSite field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideUpdateAttributes) GetOrgSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSite, true
}

// SetOrgSite sets field value.
func (o *OrgGroupPolicyOverrideUpdateAttributes) SetOrgSite(v string) {
	o.OrgSite = v
}

// GetOrgUuid returns the OrgUuid field value.
func (o *OrgGroupPolicyOverrideUpdateAttributes) GetOrgUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.OrgUuid
}

// GetOrgUuidOk returns a tuple with the OrgUuid field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideUpdateAttributes) GetOrgUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgUuid, true
}

// SetOrgUuid sets field value.
func (o *OrgGroupPolicyOverrideUpdateAttributes) SetOrgUuid(v uuid.UUID) {
	o.OrgUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyOverrideUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["org_site"] = o.OrgSite
	toSerialize["org_uuid"] = o.OrgUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyOverrideUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgSite *string    `json:"org_site"`
		OrgUuid *uuid.UUID `json:"org_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OrgSite == nil {
		return fmt.Errorf("required field org_site missing")
	}
	if all.OrgUuid == nil {
		return fmt.Errorf("required field org_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_site", "org_uuid"})
	} else {
		return err
	}
	o.OrgSite = *all.OrgSite
	o.OrgUuid = *all.OrgUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
