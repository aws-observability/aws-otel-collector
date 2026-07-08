// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabProjectsResponse Response containing a list of Model Lab projects with pagination metadata.
type ModelLabProjectsResponse struct {
	// The list of projects.
	Data []ModelLabProjectData `json:"data"`
	// Pagination links for navigating list responses.
	Links *ModelLabPaginationLinks `json:"links,omitempty"`
	// Pagination metadata for a list response.
	Meta ModelLabPageMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabProjectsResponse instantiates a new ModelLabProjectsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabProjectsResponse(data []ModelLabProjectData, meta ModelLabPageMeta) *ModelLabProjectsResponse {
	this := ModelLabProjectsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewModelLabProjectsResponseWithDefaults instantiates a new ModelLabProjectsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabProjectsResponseWithDefaults() *ModelLabProjectsResponse {
	this := ModelLabProjectsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ModelLabProjectsResponse) GetData() []ModelLabProjectData {
	if o == nil {
		var ret []ModelLabProjectData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectsResponse) GetDataOk() (*[]ModelLabProjectData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ModelLabProjectsResponse) SetData(v []ModelLabProjectData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ModelLabProjectsResponse) GetLinks() ModelLabPaginationLinks {
	if o == nil || o.Links == nil {
		var ret ModelLabPaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabProjectsResponse) GetLinksOk() (*ModelLabPaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ModelLabProjectsResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ModelLabPaginationLinks and assigns it to the Links field.
func (o *ModelLabProjectsResponse) SetLinks(v ModelLabPaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value.
func (o *ModelLabProjectsResponse) GetMeta() ModelLabPageMeta {
	if o == nil {
		var ret ModelLabPageMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectsResponse) GetMetaOk() (*ModelLabPageMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ModelLabProjectsResponse) SetMeta(v ModelLabPageMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabProjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabProjectsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]ModelLabProjectData   `json:"data"`
		Links *ModelLabPaginationLinks `json:"links,omitempty"`
		Meta  *ModelLabPageMeta        `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
