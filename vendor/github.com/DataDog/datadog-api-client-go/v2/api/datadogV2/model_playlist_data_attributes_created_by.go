// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PlaylistDataAttributesCreatedBy Information about the user who created the playlist.
type PlaylistDataAttributesCreatedBy struct {
	// Email handle of the user who created the playlist.
	Handle string `json:"handle"`
	// URL or identifier of the user's avatar icon.
	Icon *string `json:"icon,omitempty"`
	// Unique identifier of the user who created the playlist.
	Id string `json:"id"`
	// Display name of the user who created the playlist.
	Name *string `json:"name,omitempty"`
	// UUID of the user who created the playlist.
	Uuid string `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlaylistDataAttributesCreatedBy instantiates a new PlaylistDataAttributesCreatedBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlaylistDataAttributesCreatedBy(handle string, id string, uuid string) *PlaylistDataAttributesCreatedBy {
	this := PlaylistDataAttributesCreatedBy{}
	this.Handle = handle
	this.Id = id
	this.Uuid = uuid
	return &this
}

// NewPlaylistDataAttributesCreatedByWithDefaults instantiates a new PlaylistDataAttributesCreatedBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlaylistDataAttributesCreatedByWithDefaults() *PlaylistDataAttributesCreatedBy {
	this := PlaylistDataAttributesCreatedBy{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *PlaylistDataAttributesCreatedBy) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *PlaylistDataAttributesCreatedBy) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *PlaylistDataAttributesCreatedBy) SetHandle(v string) {
	o.Handle = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *PlaylistDataAttributesCreatedBy) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistDataAttributesCreatedBy) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *PlaylistDataAttributesCreatedBy) HasIcon() bool {
	return o != nil && o.Icon != nil
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *PlaylistDataAttributesCreatedBy) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value.
func (o *PlaylistDataAttributesCreatedBy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlaylistDataAttributesCreatedBy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PlaylistDataAttributesCreatedBy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlaylistDataAttributesCreatedBy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistDataAttributesCreatedBy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlaylistDataAttributesCreatedBy) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlaylistDataAttributesCreatedBy) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value.
func (o *PlaylistDataAttributesCreatedBy) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PlaylistDataAttributesCreatedBy) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *PlaylistDataAttributesCreatedBy) SetUuid(v string) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlaylistDataAttributesCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlaylistDataAttributesCreatedBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string `json:"handle"`
		Icon   *string `json:"icon,omitempty"`
		Id     *string `json:"id"`
		Name   *string `json:"name,omitempty"`
		Uuid   *string `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "icon", "id", "name", "uuid"})
	} else {
		return err
	}
	o.Handle = *all.Handle
	o.Icon = all.Icon
	o.Id = *all.Id
	o.Name = all.Name
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
