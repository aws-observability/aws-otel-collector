// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoleCreateAttributes Attributes of the created role.
type RoleCreateAttributes struct {
	// Creation time of the role.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Time of last role modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Name of the role.
	Name string `json:"name"`
	// The managed role from which this role automatically inherits new permissions.
	// Specify one of the following: "Datadog Admin Role", "Datadog Standard Role", or "Datadog Read Only Role".
	// If empty or not specified, the role does not automatically inherit permissions from any managed role.
	ReceivesPermissionsFrom []string `json:"receives_permissions_from,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoleCreateAttributes instantiates a new RoleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoleCreateAttributes(name string) *RoleCreateAttributes {
	this := RoleCreateAttributes{}
	this.Name = name
	return &this
}

// NewRoleCreateAttributesWithDefaults instantiates a new RoleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoleCreateAttributesWithDefaults() *RoleCreateAttributes {
	this := RoleCreateAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RoleCreateAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreateAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RoleCreateAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RoleCreateAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *RoleCreateAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreateAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *RoleCreateAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *RoleCreateAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value.
func (o *RoleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RoleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetReceivesPermissionsFrom returns the ReceivesPermissionsFrom field value if set, zero value otherwise.
func (o *RoleCreateAttributes) GetReceivesPermissionsFrom() []string {
	if o == nil || o.ReceivesPermissionsFrom == nil {
		var ret []string
		return ret
	}
	return o.ReceivesPermissionsFrom
}

// GetReceivesPermissionsFromOk returns a tuple with the ReceivesPermissionsFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreateAttributes) GetReceivesPermissionsFromOk() (*[]string, bool) {
	if o == nil || o.ReceivesPermissionsFrom == nil {
		return nil, false
	}
	return &o.ReceivesPermissionsFrom, true
}

// HasReceivesPermissionsFrom returns a boolean if a field has been set.
func (o *RoleCreateAttributes) HasReceivesPermissionsFrom() bool {
	return o != nil && o.ReceivesPermissionsFrom != nil
}

// SetReceivesPermissionsFrom gets a reference to the given []string and assigns it to the ReceivesPermissionsFrom field.
func (o *RoleCreateAttributes) SetReceivesPermissionsFrom(v []string) {
	o.ReceivesPermissionsFrom = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	if o.ReceivesPermissionsFrom != nil {
		toSerialize["receives_permissions_from"] = o.ReceivesPermissionsFrom
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt               *time.Time `json:"created_at,omitempty"`
		ModifiedAt              *time.Time `json:"modified_at,omitempty"`
		Name                    *string    `json:"name"`
		ReceivesPermissionsFrom []string   `json:"receives_permissions_from,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "name", "receives_permissions_from"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	o.ReceivesPermissionsFrom = all.ReceivesPermissionsFrom

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
