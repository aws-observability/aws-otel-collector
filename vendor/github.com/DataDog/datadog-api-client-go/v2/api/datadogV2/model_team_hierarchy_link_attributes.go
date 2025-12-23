// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkAttributes Team hierarchy link attributes
type TeamHierarchyLinkAttributes struct {
	// Timestamp when the team hierarchy link was created
	CreatedAt time.Time `json:"created_at"`
	// The provisioner of the team hierarchy link
	ProvisionedBy string `json:"provisioned_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinkAttributes instantiates a new TeamHierarchyLinkAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinkAttributes(createdAt time.Time, provisionedBy string) *TeamHierarchyLinkAttributes {
	this := TeamHierarchyLinkAttributes{}
	this.CreatedAt = createdAt
	this.ProvisionedBy = provisionedBy
	return &this
}

// NewTeamHierarchyLinkAttributesWithDefaults instantiates a new TeamHierarchyLinkAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkAttributesWithDefaults() *TeamHierarchyLinkAttributes {
	this := TeamHierarchyLinkAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TeamHierarchyLinkAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TeamHierarchyLinkAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetProvisionedBy returns the ProvisionedBy field value.
func (o *TeamHierarchyLinkAttributes) GetProvisionedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProvisionedBy
}

// GetProvisionedByOk returns a tuple with the ProvisionedBy field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkAttributes) GetProvisionedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisionedBy, true
}

// SetProvisionedBy sets field value.
func (o *TeamHierarchyLinkAttributes) SetProvisionedBy(v string) {
	o.ProvisionedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinkAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["provisioned_by"] = o.ProvisionedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyLinkAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time `json:"created_at"`
		ProvisionedBy *string    `json:"provisioned_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ProvisionedBy == nil {
		return fmt.Errorf("required field provisioned_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "provisioned_by"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ProvisionedBy = *all.ProvisionedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
