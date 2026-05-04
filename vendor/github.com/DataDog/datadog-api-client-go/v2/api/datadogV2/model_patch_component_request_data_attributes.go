// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchComponentRequestDataAttributes The supported attributes for updating a component.
type PatchComponentRequestDataAttributes struct {
	// The name of the component.
	Name *string `json:"name,omitempty"`
	// The position of the component. If the component belongs to a group, the position is relative to the other components in the group.
	Position *int64 `json:"position,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchComponentRequestDataAttributes instantiates a new PatchComponentRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchComponentRequestDataAttributes() *PatchComponentRequestDataAttributes {
	this := PatchComponentRequestDataAttributes{}
	return &this
}

// NewPatchComponentRequestDataAttributesWithDefaults instantiates a new PatchComponentRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchComponentRequestDataAttributesWithDefaults() *PatchComponentRequestDataAttributes {
	this := PatchComponentRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchComponentRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchComponentRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchComponentRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchComponentRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PatchComponentRequestDataAttributes) GetPosition() int64 {
	if o == nil || o.Position == nil {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchComponentRequestDataAttributes) GetPositionOk() (*int64, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PatchComponentRequestDataAttributes) HasPosition() bool {
	return o != nil && o.Position != nil
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *PatchComponentRequestDataAttributes) SetPosition(v int64) {
	o.Position = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchComponentRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchComponentRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name,omitempty"`
		Position *int64  `json:"position,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "position"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Position = all.Position

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
