// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerRelationships Holds references to objects related to the Layer entity, such as its members.
type LayerRelationships struct {
	// Holds an array of references to the members of a Layer, each containing member IDs.
	Members *LayerRelationshipsMembers `json:"members,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLayerRelationships instantiates a new LayerRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLayerRelationships() *LayerRelationships {
	this := LayerRelationships{}
	return &this
}

// NewLayerRelationshipsWithDefaults instantiates a new LayerRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLayerRelationshipsWithDefaults() *LayerRelationships {
	this := LayerRelationships{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *LayerRelationships) GetMembers() LayerRelationshipsMembers {
	if o == nil || o.Members == nil {
		var ret LayerRelationshipsMembers
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerRelationships) GetMembersOk() (*LayerRelationshipsMembers, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *LayerRelationships) HasMembers() bool {
	return o != nil && o.Members != nil
}

// SetMembers gets a reference to the given LayerRelationshipsMembers and assigns it to the Members field.
func (o *LayerRelationships) SetMembers(v LayerRelationshipsMembers) {
	o.Members = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LayerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LayerRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Members *LayerRelationshipsMembers `json:"members,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"members"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Members != nil && all.Members.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Members = all.Members

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
