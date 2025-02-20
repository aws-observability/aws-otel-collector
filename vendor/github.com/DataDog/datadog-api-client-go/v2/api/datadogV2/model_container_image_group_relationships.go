// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageGroupRelationships Relationships inside a Container Image Group.
type ContainerImageGroupRelationships struct {
	// Relationships to Container Images inside a Container Image Group.
	ContainerImages *ContainerImageGroupImagesRelationshipsLink `json:"container_images,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerImageGroupRelationships instantiates a new ContainerImageGroupRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageGroupRelationships() *ContainerImageGroupRelationships {
	this := ContainerImageGroupRelationships{}
	return &this
}

// NewContainerImageGroupRelationshipsWithDefaults instantiates a new ContainerImageGroupRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageGroupRelationshipsWithDefaults() *ContainerImageGroupRelationships {
	this := ContainerImageGroupRelationships{}
	return &this
}

// GetContainerImages returns the ContainerImages field value if set, zero value otherwise.
func (o *ContainerImageGroupRelationships) GetContainerImages() ContainerImageGroupImagesRelationshipsLink {
	if o == nil || o.ContainerImages == nil {
		var ret ContainerImageGroupImagesRelationshipsLink
		return ret
	}
	return *o.ContainerImages
}

// GetContainerImagesOk returns a tuple with the ContainerImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroupRelationships) GetContainerImagesOk() (*ContainerImageGroupImagesRelationshipsLink, bool) {
	if o == nil || o.ContainerImages == nil {
		return nil, false
	}
	return o.ContainerImages, true
}

// HasContainerImages returns a boolean if a field has been set.
func (o *ContainerImageGroupRelationships) HasContainerImages() bool {
	return o != nil && o.ContainerImages != nil
}

// SetContainerImages gets a reference to the given ContainerImageGroupImagesRelationshipsLink and assigns it to the ContainerImages field.
func (o *ContainerImageGroupRelationships) SetContainerImages(v ContainerImageGroupImagesRelationshipsLink) {
	o.ContainerImages = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageGroupRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContainerImages != nil {
		toSerialize["container_images"] = o.ContainerImages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerImageGroupRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContainerImages *ContainerImageGroupImagesRelationshipsLink `json:"container_images,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"container_images"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ContainerImages != nil && all.ContainerImages.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ContainerImages = all.ContainerImages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
