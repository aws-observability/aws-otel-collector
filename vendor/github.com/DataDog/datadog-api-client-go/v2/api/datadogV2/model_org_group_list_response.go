// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupListResponse Response containing a list of org groups.
type OrgGroupListResponse struct {
	// An array of org groups.
	Data []OrgGroupData `json:"data"`
	// Pagination links for navigating between pages of an org group list response.
	Links *OrgGroupPaginationLinks `json:"links,omitempty"`
	// Pagination metadata for org group list responses.
	Meta *OrgGroupPaginationMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupListResponse instantiates a new OrgGroupListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupListResponse(data []OrgGroupData) *OrgGroupListResponse {
	this := OrgGroupListResponse{}
	this.Data = data
	return &this
}

// NewOrgGroupListResponseWithDefaults instantiates a new OrgGroupListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupListResponseWithDefaults() *OrgGroupListResponse {
	this := OrgGroupListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *OrgGroupListResponse) GetData() []OrgGroupData {
	if o == nil {
		var ret []OrgGroupData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrgGroupListResponse) GetDataOk() (*[]OrgGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *OrgGroupListResponse) SetData(v []OrgGroupData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgGroupListResponse) GetLinks() OrgGroupPaginationLinks {
	if o == nil || o.Links == nil {
		var ret OrgGroupPaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupListResponse) GetLinksOk() (*OrgGroupPaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgGroupListResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given OrgGroupPaginationLinks and assigns it to the Links field.
func (o *OrgGroupListResponse) SetLinks(v OrgGroupPaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OrgGroupListResponse) GetMeta() OrgGroupPaginationMeta {
	if o == nil || o.Meta == nil {
		var ret OrgGroupPaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupListResponse) GetMetaOk() (*OrgGroupPaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OrgGroupListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given OrgGroupPaginationMeta and assigns it to the Meta field.
func (o *OrgGroupListResponse) SetMeta(v OrgGroupPaginationMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
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
func (o *OrgGroupListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]OrgGroupData          `json:"data"`
		Links *OrgGroupPaginationLinks `json:"links,omitempty"`
		Meta  *OrgGroupPaginationMeta  `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
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
