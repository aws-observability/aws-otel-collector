// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerMeta Response metadata object.
type ContainerMeta struct {
	// Paging attributes.
	Pagination *ContainerMetaPage `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerMeta instantiates a new ContainerMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerMeta() *ContainerMeta {
	this := ContainerMeta{}
	return &this
}

// NewContainerMetaWithDefaults instantiates a new ContainerMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerMetaWithDefaults() *ContainerMeta {
	this := ContainerMeta{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ContainerMeta) GetPagination() ContainerMetaPage {
	if o == nil || o.Pagination == nil {
		var ret ContainerMetaPage
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerMeta) GetPaginationOk() (*ContainerMetaPage, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ContainerMeta) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given ContainerMetaPage and assigns it to the Pagination field.
func (o *ContainerMeta) SetPagination(v ContainerMetaPage) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pagination *ContainerMetaPage `json:"pagination,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
