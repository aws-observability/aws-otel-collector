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

// OrgGroupPolicyOverrideAttributes Attributes of an org group policy override.
type OrgGroupPolicyOverrideAttributes struct {
	// The override content as key-value pairs.
	Content map[string]interface{} `json:"content,omitempty"`
	// Timestamp when the override was created.
	CreatedAt time.Time `json:"created_at"`
	// Timestamp when the override was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The site of the organization that has the override.
	OrgSite string `json:"org_site"`
	// The UUID of the organization that has the override.
	OrgUuid uuid.UUID `json:"org_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyOverrideAttributes instantiates a new OrgGroupPolicyOverrideAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyOverrideAttributes(createdAt time.Time, modifiedAt time.Time, orgSite string, orgUuid uuid.UUID) *OrgGroupPolicyOverrideAttributes {
	this := OrgGroupPolicyOverrideAttributes{}
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	this.OrgSite = orgSite
	this.OrgUuid = orgUuid
	return &this
}

// NewOrgGroupPolicyOverrideAttributesWithDefaults instantiates a new OrgGroupPolicyOverrideAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyOverrideAttributesWithDefaults() *OrgGroupPolicyOverrideAttributes {
	this := OrgGroupPolicyOverrideAttributes{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OrgGroupPolicyOverrideAttributes) GetContent() map[string]interface{} {
	if o == nil || o.Content == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideAttributes) GetContentOk() (*map[string]interface{}, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OrgGroupPolicyOverrideAttributes) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *OrgGroupPolicyOverrideAttributes) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OrgGroupPolicyOverrideAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OrgGroupPolicyOverrideAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *OrgGroupPolicyOverrideAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *OrgGroupPolicyOverrideAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetOrgSite returns the OrgSite field value.
func (o *OrgGroupPolicyOverrideAttributes) GetOrgSite() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgSite
}

// GetOrgSiteOk returns a tuple with the OrgSite field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideAttributes) GetOrgSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSite, true
}

// SetOrgSite sets field value.
func (o *OrgGroupPolicyOverrideAttributes) SetOrgSite(v string) {
	o.OrgSite = v
}

// GetOrgUuid returns the OrgUuid field value.
func (o *OrgGroupPolicyOverrideAttributes) GetOrgUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.OrgUuid
}

// GetOrgUuidOk returns a tuple with the OrgUuid field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideAttributes) GetOrgUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgUuid, true
}

// SetOrgUuid sets field value.
func (o *OrgGroupPolicyOverrideAttributes) SetOrgUuid(v uuid.UUID) {
	o.OrgUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyOverrideAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
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
	toSerialize["org_site"] = o.OrgSite
	toSerialize["org_uuid"] = o.OrgUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyOverrideAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content    map[string]interface{} `json:"content,omitempty"`
		CreatedAt  *time.Time             `json:"created_at"`
		ModifiedAt *time.Time             `json:"modified_at"`
		OrgSite    *string                `json:"org_site"`
		OrgUuid    *uuid.UUID             `json:"org_uuid"`
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
	if all.OrgSite == nil {
		return fmt.Errorf("required field org_site missing")
	}
	if all.OrgUuid == nil {
		return fmt.Errorf("required field org_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "created_at", "modified_at", "org_site", "org_uuid"})
	} else {
		return err
	}
	o.Content = all.Content
	o.CreatedAt = *all.CreatedAt
	o.ModifiedAt = *all.ModifiedAt
	o.OrgSite = *all.OrgSite
	o.OrgUuid = *all.OrgUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
