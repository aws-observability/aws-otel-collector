// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAttributes Attributes of an organization.
type OrgAttributes struct {
	// The creation timestamp of the organization.
	CreatedAt time.Time `json:"created_at"`
	// A description of the organization.
	Description string `json:"description"`
	// Whether the organization is disabled.
	Disabled bool `json:"disabled"`
	// The last modification timestamp of the organization.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the organization.
	Name string `json:"name"`
	// The public identifier of the organization.
	PublicId string `json:"public_id"`
	// The sharing setting of the organization.
	Sharing string `json:"sharing"`
	// The URL of the organization.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgAttributes instantiates a new OrgAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgAttributes(createdAt time.Time, description string, disabled bool, modifiedAt time.Time, name string, publicId string, sharing string, url string) *OrgAttributes {
	this := OrgAttributes{}
	this.CreatedAt = createdAt
	this.Description = description
	this.Disabled = disabled
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.PublicId = publicId
	this.Sharing = sharing
	this.Url = url
	return &this
}

// NewOrgAttributesWithDefaults instantiates a new OrgAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgAttributesWithDefaults() *OrgAttributes {
	this := OrgAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OrgAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OrgAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *OrgAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *OrgAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDisabled returns the Disabled field value.
func (o *OrgAttributes) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value.
func (o *OrgAttributes) SetDisabled(v bool) {
	o.Disabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *OrgAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *OrgAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *OrgAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrgAttributes) SetName(v string) {
	o.Name = v
}

// GetPublicId returns the PublicId field value.
func (o *OrgAttributes) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value.
func (o *OrgAttributes) SetPublicId(v string) {
	o.PublicId = v
}

// GetSharing returns the Sharing field value.
func (o *OrgAttributes) GetSharing() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetSharingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sharing, true
}

// SetSharing sets field value.
func (o *OrgAttributes) SetSharing(v string) {
	o.Sharing = v
}

// GetUrl returns the Url field value.
func (o *OrgAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrgAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *OrgAttributes) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	toSerialize["disabled"] = o.Disabled
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["public_id"] = o.PublicId
	toSerialize["sharing"] = o.Sharing
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"created_at"`
		Description *string    `json:"description"`
		Disabled    *bool      `json:"disabled"`
		ModifiedAt  *time.Time `json:"modified_at"`
		Name        *string    `json:"name"`
		PublicId    *string    `json:"public_id"`
		Sharing     *string    `json:"sharing"`
		Url         *string    `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Disabled == nil {
		return fmt.Errorf("required field disabled missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.PublicId == nil {
		return fmt.Errorf("required field public_id missing")
	}
	if all.Sharing == nil {
		return fmt.Errorf("required field sharing missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "disabled", "modified_at", "name", "public_id", "sharing", "url"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.Disabled = *all.Disabled
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.PublicId = *all.PublicId
	o.Sharing = *all.Sharing
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
