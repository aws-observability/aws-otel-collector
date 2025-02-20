// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainersResponse List of containers.
type ContainersResponse struct {
	// Array of Container objects.
	Data []ContainerItem `json:"data,omitempty"`
	// Pagination links.
	Links *ContainersResponseLinks `json:"links,omitempty"`
	// Response metadata object.
	Meta *ContainerMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainersResponse instantiates a new ContainersResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainersResponse() *ContainersResponse {
	this := ContainersResponse{}
	return &this
}

// NewContainersResponseWithDefaults instantiates a new ContainersResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainersResponseWithDefaults() *ContainersResponse {
	this := ContainersResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ContainersResponse) GetData() []ContainerItem {
	if o == nil || o.Data == nil {
		var ret []ContainerItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersResponse) GetDataOk() (*[]ContainerItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ContainersResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []ContainerItem and assigns it to the Data field.
func (o *ContainersResponse) SetData(v []ContainerItem) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ContainersResponse) GetLinks() ContainersResponseLinks {
	if o == nil || o.Links == nil {
		var ret ContainersResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersResponse) GetLinksOk() (*ContainersResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ContainersResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ContainersResponseLinks and assigns it to the Links field.
func (o *ContainersResponse) SetLinks(v ContainersResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ContainersResponse) GetMeta() ContainerMeta {
	if o == nil || o.Meta == nil {
		var ret ContainerMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersResponse) GetMetaOk() (*ContainerMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ContainersResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ContainerMeta and assigns it to the Meta field.
func (o *ContainersResponse) SetMeta(v ContainerMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainersResponse) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainersResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  []ContainerItem          `json:"data,omitempty"`
		Links *ContainersResponseLinks `json:"links,omitempty"`
		Meta  *ContainerMeta           `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
