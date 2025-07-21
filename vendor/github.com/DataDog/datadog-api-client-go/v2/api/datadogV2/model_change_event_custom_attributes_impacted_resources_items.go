// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributesImpactedResourcesItems Object representing a uniquely identified resource.
type ChangeEventCustomAttributesImpactedResourcesItems struct {
	// The name of the impacted resource. Limited to 128 characters.
	Name string `json:"name"`
	// The type of the impacted resource.
	Type ChangeEventCustomAttributesImpactedResourcesItemsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeEventCustomAttributesImpactedResourcesItems instantiates a new ChangeEventCustomAttributesImpactedResourcesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventCustomAttributesImpactedResourcesItems(name string, typeVar ChangeEventCustomAttributesImpactedResourcesItemsType) *ChangeEventCustomAttributesImpactedResourcesItems {
	this := ChangeEventCustomAttributesImpactedResourcesItems{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewChangeEventCustomAttributesImpactedResourcesItemsWithDefaults instantiates a new ChangeEventCustomAttributesImpactedResourcesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventCustomAttributesImpactedResourcesItemsWithDefaults() *ChangeEventCustomAttributesImpactedResourcesItems {
	this := ChangeEventCustomAttributesImpactedResourcesItems{}
	return &this
}

// GetName returns the Name field value.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) GetType() ChangeEventCustomAttributesImpactedResourcesItemsType {
	if o == nil {
		var ret ChangeEventCustomAttributesImpactedResourcesItemsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) GetTypeOk() (*ChangeEventCustomAttributesImpactedResourcesItemsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) SetType(v ChangeEventCustomAttributesImpactedResourcesItemsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventCustomAttributesImpactedResourcesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEventCustomAttributesImpactedResourcesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                                                `json:"name"`
		Type *ChangeEventCustomAttributesImpactedResourcesItemsType `json:"type"`
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
