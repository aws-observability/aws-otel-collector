// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListSecurityFindingsResponse The expected response schema when listing security findings.
type ListSecurityFindingsResponse struct {
	// Array of security findings matching the search query.
	Data []SecurityFindingsData `json:"data,omitempty"`
	// Links for pagination.
	Links *SecurityFindingsLinks `json:"links,omitempty"`
	// Metadata about the response.
	Meta *SecurityFindingsMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListSecurityFindingsResponse instantiates a new ListSecurityFindingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListSecurityFindingsResponse() *ListSecurityFindingsResponse {
	this := ListSecurityFindingsResponse{}
	return &this
}

// NewListSecurityFindingsResponseWithDefaults instantiates a new ListSecurityFindingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListSecurityFindingsResponseWithDefaults() *ListSecurityFindingsResponse {
	this := ListSecurityFindingsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSecurityFindingsResponse) GetData() []SecurityFindingsData {
	if o == nil || o.Data == nil {
		var ret []SecurityFindingsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityFindingsResponse) GetDataOk() (*[]SecurityFindingsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSecurityFindingsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []SecurityFindingsData and assigns it to the Data field.
func (o *ListSecurityFindingsResponse) SetData(v []SecurityFindingsData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListSecurityFindingsResponse) GetLinks() SecurityFindingsLinks {
	if o == nil || o.Links == nil {
		var ret SecurityFindingsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityFindingsResponse) GetLinksOk() (*SecurityFindingsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListSecurityFindingsResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given SecurityFindingsLinks and assigns it to the Links field.
func (o *ListSecurityFindingsResponse) SetLinks(v SecurityFindingsLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListSecurityFindingsResponse) GetMeta() SecurityFindingsMeta {
	if o == nil || o.Meta == nil {
		var ret SecurityFindingsMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityFindingsResponse) GetMetaOk() (*SecurityFindingsMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListSecurityFindingsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SecurityFindingsMeta and assigns it to the Meta field.
func (o *ListSecurityFindingsResponse) SetMeta(v SecurityFindingsMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListSecurityFindingsResponse) MarshalJSON() ([]byte, error) {
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
func (o *ListSecurityFindingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  []SecurityFindingsData `json:"data,omitempty"`
		Links *SecurityFindingsLinks `json:"links,omitempty"`
		Meta  *SecurityFindingsMeta  `json:"meta,omitempty"`
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
