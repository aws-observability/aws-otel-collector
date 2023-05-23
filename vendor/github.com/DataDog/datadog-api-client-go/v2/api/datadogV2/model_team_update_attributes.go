// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// TeamUpdateAttributes Team update attributes
type TeamUpdateAttributes struct {
	// An identifier for the color representing the team
	Color *int32 `json:"color,omitempty"`
	// Free-form markdown description/content for the team's homepage
	Description *string `json:"description,omitempty"`
	// The team's identifier
	Handle string `json:"handle"`
	// The number of links belonging to the team
	LinkCount *int32 `json:"link_count,omitempty"`
	// The name of the team
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewTeamUpdateAttributes instantiates a new TeamUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamUpdateAttributes(handle string, name string) *TeamUpdateAttributes {
	this := TeamUpdateAttributes{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewTeamUpdateAttributesWithDefaults instantiates a new TeamUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamUpdateAttributesWithDefaults() *TeamUpdateAttributes {
	this := TeamUpdateAttributes{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetColor() int32 {
	if o == nil || o.Color == nil {
		var ret int32
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetColorOk() (*int32, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasColor() bool {
	return o != nil && o.Color != nil
}

// SetColor gets a reference to the given int32 and assigns it to the Color field.
func (o *TeamUpdateAttributes) SetColor(v int32) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value.
func (o *TeamUpdateAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *TeamUpdateAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetLinkCount returns the LinkCount field value if set, zero value otherwise.
func (o *TeamUpdateAttributes) GetLinkCount() int32 {
	if o == nil || o.LinkCount == nil {
		var ret int32
		return ret
	}
	return *o.LinkCount
}

// GetLinkCountOk returns a tuple with the LinkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetLinkCountOk() (*int32, bool) {
	if o == nil || o.LinkCount == nil {
		return nil, false
	}
	return o.LinkCount, true
}

// HasLinkCount returns a boolean if a field has been set.
func (o *TeamUpdateAttributes) HasLinkCount() bool {
	return o != nil && o.LinkCount != nil
}

// SetLinkCount gets a reference to the given int32 and assigns it to the LinkCount field.
func (o *TeamUpdateAttributes) SetLinkCount(v int32) {
	o.LinkCount = &v
}

// GetName returns the Name field value.
func (o *TeamUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TeamUpdateAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["handle"] = o.Handle
	if o.LinkCount != nil {
		toSerialize["link_count"] = o.LinkCount
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Handle *string `json:"handle"`
		Name   *string `json:"name"`
	}{}
	all := struct {
		Color       *int32  `json:"color,omitempty"`
		Description *string `json:"description,omitempty"`
		Handle      string  `json:"handle"`
		LinkCount   *int32  `json:"link_count,omitempty"`
		Name        string  `json:"name"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if required.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Color = all.Color
	o.Description = all.Description
	o.Handle = all.Handle
	o.LinkCount = all.LinkCount
	o.Name = all.Name
	return nil
}
