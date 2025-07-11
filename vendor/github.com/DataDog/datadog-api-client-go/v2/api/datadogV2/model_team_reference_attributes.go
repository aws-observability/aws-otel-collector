// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamReferenceAttributes Encapsulates the basic attributes of a Team reference, such as name, handle, and an optional avatar or description.
type TeamReferenceAttributes struct {
	// URL or reference for the team's avatar (if available).
	Avatar *string `json:"avatar,omitempty"`
	// A short text describing the team.
	Description *string `json:"description,omitempty"`
	// A unique handle/slug for the team.
	Handle *string `json:"handle,omitempty"`
	// The full, human-readable name of the team.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamReferenceAttributes instantiates a new TeamReferenceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamReferenceAttributes() *TeamReferenceAttributes {
	this := TeamReferenceAttributes{}
	return &this
}

// NewTeamReferenceAttributesWithDefaults instantiates a new TeamReferenceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamReferenceAttributesWithDefaults() *TeamReferenceAttributes {
	this := TeamReferenceAttributes{}
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *TeamReferenceAttributes) GetAvatar() string {
	if o == nil || o.Avatar == nil {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamReferenceAttributes) GetAvatarOk() (*string, bool) {
	if o == nil || o.Avatar == nil {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamReferenceAttributes) HasAvatar() bool {
	return o != nil && o.Avatar != nil
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *TeamReferenceAttributes) SetAvatar(v string) {
	o.Avatar = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamReferenceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamReferenceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamReferenceAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamReferenceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *TeamReferenceAttributes) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamReferenceAttributes) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *TeamReferenceAttributes) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *TeamReferenceAttributes) SetHandle(v string) {
	o.Handle = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamReferenceAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamReferenceAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamReferenceAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamReferenceAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamReferenceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamReferenceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Avatar      *string `json:"avatar,omitempty"`
		Description *string `json:"description,omitempty"`
		Handle      *string `json:"handle,omitempty"`
		Name        *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avatar", "description", "handle", "name"})
	} else {
		return err
	}
	o.Avatar = all.Avatar
	o.Description = all.Description
	o.Handle = all.Handle
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
