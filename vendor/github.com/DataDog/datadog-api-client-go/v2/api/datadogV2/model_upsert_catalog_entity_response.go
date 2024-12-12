// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertCatalogEntityResponse Upsert entity response.
type UpsertCatalogEntityResponse struct {
	// List of entity data.
	Data []EntityData `json:"data,omitempty"`
	// Upsert entity response included.
	Included []UpsertCatalogEntityResponseIncludedItem `json:"included,omitempty"`
	// Entity metadata.
	Meta *EntityResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertCatalogEntityResponse instantiates a new UpsertCatalogEntityResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertCatalogEntityResponse() *UpsertCatalogEntityResponse {
	this := UpsertCatalogEntityResponse{}
	return &this
}

// NewUpsertCatalogEntityResponseWithDefaults instantiates a new UpsertCatalogEntityResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertCatalogEntityResponseWithDefaults() *UpsertCatalogEntityResponse {
	this := UpsertCatalogEntityResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpsertCatalogEntityResponse) GetData() []EntityData {
	if o == nil || o.Data == nil {
		var ret []EntityData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCatalogEntityResponse) GetDataOk() (*[]EntityData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpsertCatalogEntityResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []EntityData and assigns it to the Data field.
func (o *UpsertCatalogEntityResponse) SetData(v []EntityData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *UpsertCatalogEntityResponse) GetIncluded() []UpsertCatalogEntityResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []UpsertCatalogEntityResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCatalogEntityResponse) GetIncludedOk() (*[]UpsertCatalogEntityResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *UpsertCatalogEntityResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []UpsertCatalogEntityResponseIncludedItem and assigns it to the Included field.
func (o *UpsertCatalogEntityResponse) SetIncluded(v []UpsertCatalogEntityResponseIncludedItem) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UpsertCatalogEntityResponse) GetMeta() EntityResponseMeta {
	if o == nil || o.Meta == nil {
		var ret EntityResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCatalogEntityResponse) GetMetaOk() (*EntityResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UpsertCatalogEntityResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given EntityResponseMeta and assigns it to the Meta field.
func (o *UpsertCatalogEntityResponse) SetMeta(v EntityResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertCatalogEntityResponse) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertCatalogEntityResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []EntityData                              `json:"data,omitempty"`
		Included []UpsertCatalogEntityResponseIncludedItem `json:"included,omitempty"`
		Meta     *EntityResponseMeta                       `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	o.Included = all.Included
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
