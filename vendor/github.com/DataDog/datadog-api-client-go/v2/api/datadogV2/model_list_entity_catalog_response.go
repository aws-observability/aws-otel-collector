// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListEntityCatalogResponse List entity response.
type ListEntityCatalogResponse struct {
	// List of entity data.
	Data []EntityData `json:"data,omitempty"`
	// List entity response included.
	Included []ListEntityCatalogResponseIncludedItem `json:"included,omitempty"`
	// List entity response links.
	Links *ListEntityCatalogResponseLinks `json:"links,omitempty"`
	// Entity metadata.
	Meta *EntityResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListEntityCatalogResponse instantiates a new ListEntityCatalogResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListEntityCatalogResponse() *ListEntityCatalogResponse {
	this := ListEntityCatalogResponse{}
	return &this
}

// NewListEntityCatalogResponseWithDefaults instantiates a new ListEntityCatalogResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListEntityCatalogResponseWithDefaults() *ListEntityCatalogResponse {
	this := ListEntityCatalogResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListEntityCatalogResponse) GetData() []EntityData {
	if o == nil || o.Data == nil {
		var ret []EntityData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponse) GetDataOk() (*[]EntityData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListEntityCatalogResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []EntityData and assigns it to the Data field.
func (o *ListEntityCatalogResponse) SetData(v []EntityData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ListEntityCatalogResponse) GetIncluded() []ListEntityCatalogResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []ListEntityCatalogResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponse) GetIncludedOk() (*[]ListEntityCatalogResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ListEntityCatalogResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []ListEntityCatalogResponseIncludedItem and assigns it to the Included field.
func (o *ListEntityCatalogResponse) SetIncluded(v []ListEntityCatalogResponseIncludedItem) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListEntityCatalogResponse) GetLinks() ListEntityCatalogResponseLinks {
	if o == nil || o.Links == nil {
		var ret ListEntityCatalogResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponse) GetLinksOk() (*ListEntityCatalogResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListEntityCatalogResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ListEntityCatalogResponseLinks and assigns it to the Links field.
func (o *ListEntityCatalogResponse) SetLinks(v ListEntityCatalogResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListEntityCatalogResponse) GetMeta() EntityResponseMeta {
	if o == nil || o.Meta == nil {
		var ret EntityResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponse) GetMetaOk() (*EntityResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListEntityCatalogResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given EntityResponseMeta and assigns it to the Meta field.
func (o *ListEntityCatalogResponse) SetMeta(v EntityResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListEntityCatalogResponse) MarshalJSON() ([]byte, error) {
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
func (o *ListEntityCatalogResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []EntityData                            `json:"data,omitempty"`
		Included []ListEntityCatalogResponseIncludedItem `json:"included,omitempty"`
		Links    *ListEntityCatalogResponseLinks         `json:"links,omitempty"`
		Meta     *EntityResponseMeta                     `json:"meta,omitempty"`
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
