// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributesAuthor The entity that made the change. Optional, if provided it must include `type` and `name`.
type ChangeEventCustomAttributesAuthor struct {
	// The name of the user or system that made the change. Limited to 128 characters.
	Name string `json:"name"`
	// Author's type.
	Type ChangeEventCustomAttributesAuthorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeEventCustomAttributesAuthor instantiates a new ChangeEventCustomAttributesAuthor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventCustomAttributesAuthor(name string, typeVar ChangeEventCustomAttributesAuthorType) *ChangeEventCustomAttributesAuthor {
	this := ChangeEventCustomAttributesAuthor{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewChangeEventCustomAttributesAuthorWithDefaults instantiates a new ChangeEventCustomAttributesAuthor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventCustomAttributesAuthorWithDefaults() *ChangeEventCustomAttributesAuthor {
	this := ChangeEventCustomAttributesAuthor{}
	return &this
}

// GetName returns the Name field value.
func (o *ChangeEventCustomAttributesAuthor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributesAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ChangeEventCustomAttributesAuthor) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ChangeEventCustomAttributesAuthor) GetType() ChangeEventCustomAttributesAuthorType {
	if o == nil {
		var ret ChangeEventCustomAttributesAuthorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributesAuthor) GetTypeOk() (*ChangeEventCustomAttributesAuthorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChangeEventCustomAttributesAuthor) SetType(v ChangeEventCustomAttributesAuthorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventCustomAttributesAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEventCustomAttributesAuthor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                                `json:"name"`
		Type *ChangeEventCustomAttributesAuthorType `json:"type"`
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
