// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerGroupRelationships Relationships to containers inside a container group.
type ContainerGroupRelationships struct {
	// Relationships to Containers inside a Container Group.
	Containers *ContainerGroupRelationshipsLink `json:"containers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerGroupRelationships instantiates a new ContainerGroupRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerGroupRelationships() *ContainerGroupRelationships {
	this := ContainerGroupRelationships{}
	return &this
}

// NewContainerGroupRelationshipsWithDefaults instantiates a new ContainerGroupRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerGroupRelationshipsWithDefaults() *ContainerGroupRelationships {
	this := ContainerGroupRelationships{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ContainerGroupRelationships) GetContainers() ContainerGroupRelationshipsLink {
	if o == nil || o.Containers == nil {
		var ret ContainerGroupRelationshipsLink
		return ret
	}
	return *o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupRelationships) GetContainersOk() (*ContainerGroupRelationshipsLink, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ContainerGroupRelationships) HasContainers() bool {
	return o != nil && o.Containers != nil
}

// SetContainers gets a reference to the given ContainerGroupRelationshipsLink and assigns it to the Containers field.
func (o *ContainerGroupRelationships) SetContainers(v ContainerGroupRelationshipsLink) {
	o.Containers = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerGroupRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerGroupRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Containers *ContainerGroupRelationshipsLink `json:"containers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"containers"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Containers != nil && all.Containers.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Containers = all.Containers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
