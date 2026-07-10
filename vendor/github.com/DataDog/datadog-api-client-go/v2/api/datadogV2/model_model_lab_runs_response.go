// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabRunsResponse Response containing a list of Model Lab runs with pagination metadata.
type ModelLabRunsResponse struct {
	// The list of runs.
	Data []ModelLabRunData `json:"data"`
	// Pagination links for navigating list responses.
	Links *ModelLabPaginationLinks `json:"links,omitempty"`
	// Pagination metadata for a list response.
	Meta ModelLabPageMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabRunsResponse instantiates a new ModelLabRunsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabRunsResponse(data []ModelLabRunData, meta ModelLabPageMeta) *ModelLabRunsResponse {
	this := ModelLabRunsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewModelLabRunsResponseWithDefaults instantiates a new ModelLabRunsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabRunsResponseWithDefaults() *ModelLabRunsResponse {
	this := ModelLabRunsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ModelLabRunsResponse) GetData() []ModelLabRunData {
	if o == nil {
		var ret []ModelLabRunData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunsResponse) GetDataOk() (*[]ModelLabRunData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ModelLabRunsResponse) SetData(v []ModelLabRunData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ModelLabRunsResponse) GetLinks() ModelLabPaginationLinks {
	if o == nil || o.Links == nil {
		var ret ModelLabPaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabRunsResponse) GetLinksOk() (*ModelLabPaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ModelLabRunsResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ModelLabPaginationLinks and assigns it to the Links field.
func (o *ModelLabRunsResponse) SetLinks(v ModelLabPaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value.
func (o *ModelLabRunsResponse) GetMeta() ModelLabPageMeta {
	if o == nil {
		var ret ModelLabPageMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunsResponse) GetMetaOk() (*ModelLabPageMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ModelLabRunsResponse) SetMeta(v ModelLabPageMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabRunsResponse) MarshalJSON() ([]byte, error) {
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
func (o *ModelLabRunsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]ModelLabRunData       `json:"data"`
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
