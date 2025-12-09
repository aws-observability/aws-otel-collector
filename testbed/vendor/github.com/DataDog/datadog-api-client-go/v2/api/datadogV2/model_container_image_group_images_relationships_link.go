// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageGroupImagesRelationshipsLink Relationships to Container Images inside a Container Image Group.
type ContainerImageGroupImagesRelationshipsLink struct {
	// Links data.
	Data []string `json:"data,omitempty"`
	// Links attributes.
	Links *ContainerImageGroupRelationshipsLinks `json:"links,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerImageGroupImagesRelationshipsLink instantiates a new ContainerImageGroupImagesRelationshipsLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageGroupImagesRelationshipsLink() *ContainerImageGroupImagesRelationshipsLink {
	this := ContainerImageGroupImagesRelationshipsLink{}
	return &this
}

// NewContainerImageGroupImagesRelationshipsLinkWithDefaults instantiates a new ContainerImageGroupImagesRelationshipsLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageGroupImagesRelationshipsLinkWithDefaults() *ContainerImageGroupImagesRelationshipsLink {
	this := ContainerImageGroupImagesRelationshipsLink{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainerImageGroupImagesRelationshipsLink) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroupImagesRelationshipsLink) GetDataOk() (*[]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainerImageGroupImagesRelationshipsLink) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *ContainerImageGroupImagesRelationshipsLink) SetData(v []string) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ContainerImageGroupImagesRelationshipsLink) GetLinks() ContainerImageGroupRelationshipsLinks {
	if o == nil || o.Links == nil {
		var ret ContainerImageGroupRelationshipsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageGroupImagesRelationshipsLink) GetLinksOk() (*ContainerImageGroupRelationshipsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ContainerImageGroupImagesRelationshipsLink) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ContainerImageGroupRelationshipsLinks and assigns it to the Links field.
func (o *ContainerImageGroupImagesRelationshipsLink) SetLinks(v ContainerImageGroupRelationshipsLinks) {
	o.Links = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageGroupImagesRelationshipsLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerImageGroupImagesRelationshipsLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  []string                               `json:"data,omitempty"`
		Links *ContainerImageGroupRelationshipsLinks `json:"links,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
