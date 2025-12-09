// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributesChangedResource A uniquely identified resource.
type ChangeEventCustomAttributesChangedResource struct {
	// The name of the resource that was changed. Limited to 128 characters.
	Name string `json:"name"`
	// The type of the resource that was changed.
	Type ChangeEventCustomAttributesChangedResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeEventCustomAttributesChangedResource instantiates a new ChangeEventCustomAttributesChangedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventCustomAttributesChangedResource(name string, typeVar ChangeEventCustomAttributesChangedResourceType) *ChangeEventCustomAttributesChangedResource {
	this := ChangeEventCustomAttributesChangedResource{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewChangeEventCustomAttributesChangedResourceWithDefaults instantiates a new ChangeEventCustomAttributesChangedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventCustomAttributesChangedResourceWithDefaults() *ChangeEventCustomAttributesChangedResource {
	this := ChangeEventCustomAttributesChangedResource{}
	return &this
}

// GetName returns the Name field value.
func (o *ChangeEventCustomAttributesChangedResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributesChangedResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ChangeEventCustomAttributesChangedResource) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ChangeEventCustomAttributesChangedResource) GetType() ChangeEventCustomAttributesChangedResourceType {
	if o == nil {
		var ret ChangeEventCustomAttributesChangedResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributesChangedResource) GetTypeOk() (*ChangeEventCustomAttributesChangedResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChangeEventCustomAttributesChangedResource) SetType(v ChangeEventCustomAttributesChangedResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventCustomAttributesChangedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEventCustomAttributesChangedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                                         `json:"name"`
		Type *ChangeEventCustomAttributesChangedResourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
