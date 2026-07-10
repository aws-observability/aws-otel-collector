// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrg Organization information for a global organization association.
type GlobalOrg struct {
	// The name of the organization.
	Name string `json:"name"`
	// The public identifier of the organization.
	PublicId datadog.NullableString `json:"public_id,omitempty"`
	// The subdomain used to access the organization, if configured.
	Subdomain datadog.NullableString `json:"subdomain,omitempty"`
	// The UUID of the organization.
	Uuid uuid.UUID `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalOrg instantiates a new GlobalOrg object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalOrg(name string, uuid uuid.UUID) *GlobalOrg {
	this := GlobalOrg{}
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewGlobalOrgWithDefaults instantiates a new GlobalOrg object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalOrgWithDefaults() *GlobalOrg {
	this := GlobalOrg{}
	return &this
}

// GetName returns the Name field value.
func (o *GlobalOrg) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GlobalOrg) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GlobalOrg) SetName(v string) {
	o.Name = v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrg) GetPublicId() string {
	if o == nil || o.PublicId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicId.Get()
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrg) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicId.Get(), o.PublicId.IsSet()
}

// HasPublicId returns a boolean if a field has been set.
func (o *GlobalOrg) HasPublicId() bool {
	return o != nil && o.PublicId.IsSet()
}

// SetPublicId gets a reference to the given datadog.NullableString and assigns it to the PublicId field.
func (o *GlobalOrg) SetPublicId(v string) {
	o.PublicId.Set(&v)
}

// SetPublicIdNil sets the value for PublicId to be an explicit nil.
func (o *GlobalOrg) SetPublicIdNil() {
	o.PublicId.Set(nil)
}

// UnsetPublicId ensures that no value is present for PublicId, not even an explicit nil.
func (o *GlobalOrg) UnsetPublicId() {
	o.PublicId.Unset()
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalOrg) GetSubdomain() string {
	if o == nil || o.Subdomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subdomain.Get()
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GlobalOrg) GetSubdomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subdomain.Get(), o.Subdomain.IsSet()
}

// HasSubdomain returns a boolean if a field has been set.
func (o *GlobalOrg) HasSubdomain() bool {
	return o != nil && o.Subdomain.IsSet()
}

// SetSubdomain gets a reference to the given datadog.NullableString and assigns it to the Subdomain field.
func (o *GlobalOrg) SetSubdomain(v string) {
	o.Subdomain.Set(&v)
}

// SetSubdomainNil sets the value for Subdomain to be an explicit nil.
func (o *GlobalOrg) SetSubdomainNil() {
	o.Subdomain.Set(nil)
}

// UnsetSubdomain ensures that no value is present for Subdomain, not even an explicit nil.
func (o *GlobalOrg) UnsetSubdomain() {
	o.Subdomain.Unset()
}

// GetUuid returns the Uuid field value.
func (o *GlobalOrg) GetUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *GlobalOrg) GetUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *GlobalOrg) SetUuid(v uuid.UUID) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalOrg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.PublicId.IsSet() {
		toSerialize["public_id"] = o.PublicId.Get()
	}
	if o.Subdomain.IsSet() {
		toSerialize["subdomain"] = o.Subdomain.Get()
	}
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalOrg) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string                `json:"name"`
		PublicId  datadog.NullableString `json:"public_id,omitempty"`
		Subdomain datadog.NullableString `json:"subdomain,omitempty"`
		Uuid      *uuid.UUID             `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "public_id", "subdomain", "uuid"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.PublicId = all.PublicId
	o.Subdomain = all.Subdomain
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
