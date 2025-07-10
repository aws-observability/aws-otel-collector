// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListRelationCatalogResponse List entity relation response.
type ListRelationCatalogResponse struct {
	// Array of relation responses
	Data []RelationResponse `json:"data,omitempty"`
	// List relation response included entities.
	Included []EntityData `json:"included,omitempty"`
	// List relation response links.
	Links *ListRelationCatalogResponseLinks `json:"links,omitempty"`
	// Relation response metadata.
	Meta *RelationResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListRelationCatalogResponse instantiates a new ListRelationCatalogResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListRelationCatalogResponse() *ListRelationCatalogResponse {
	this := ListRelationCatalogResponse{}
	return &this
}

// NewListRelationCatalogResponseWithDefaults instantiates a new ListRelationCatalogResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListRelationCatalogResponseWithDefaults() *ListRelationCatalogResponse {
	this := ListRelationCatalogResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListRelationCatalogResponse) GetData() []RelationResponse {
	if o == nil || o.Data == nil {
		var ret []RelationResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRelationCatalogResponse) GetDataOk() (*[]RelationResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListRelationCatalogResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []RelationResponse and assigns it to the Data field.
func (o *ListRelationCatalogResponse) SetData(v []RelationResponse) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ListRelationCatalogResponse) GetIncluded() []EntityData {
	if o == nil || o.Included == nil {
		var ret []EntityData
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRelationCatalogResponse) GetIncludedOk() (*[]EntityData, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ListRelationCatalogResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []EntityData and assigns it to the Included field.
func (o *ListRelationCatalogResponse) SetIncluded(v []EntityData) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListRelationCatalogResponse) GetLinks() ListRelationCatalogResponseLinks {
	if o == nil || o.Links == nil {
		var ret ListRelationCatalogResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRelationCatalogResponse) GetLinksOk() (*ListRelationCatalogResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListRelationCatalogResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ListRelationCatalogResponseLinks and assigns it to the Links field.
func (o *ListRelationCatalogResponse) SetLinks(v ListRelationCatalogResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListRelationCatalogResponse) GetMeta() RelationResponseMeta {
	if o == nil || o.Meta == nil {
		var ret RelationResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRelationCatalogResponse) GetMetaOk() (*RelationResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListRelationCatalogResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given RelationResponseMeta and assigns it to the Meta field.
func (o *ListRelationCatalogResponse) SetMeta(v RelationResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListRelationCatalogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
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
func (o *ListRelationCatalogResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []RelationResponse                `json:"data,omitempty"`
		Included []EntityData                      `json:"included,omitempty"`
		Links    *ListRelationCatalogResponseLinks `json:"links,omitempty"`
		Meta     *RelationResponseMeta             `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	o.Included = all.Included
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
