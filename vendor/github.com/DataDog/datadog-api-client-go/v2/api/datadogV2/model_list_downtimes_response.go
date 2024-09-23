// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListDowntimesResponse Response for retrieving all downtimes.
type ListDowntimesResponse struct {
	// An array of downtimes.
	Data []DowntimeResponseData `json:"data,omitempty"`
	// Array of objects related to the downtimes.
	Included []DowntimeResponseIncludedItem `json:"included,omitempty"`
	// Pagination metadata returned by the API.
	Meta *DowntimeMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListDowntimesResponse instantiates a new ListDowntimesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListDowntimesResponse() *ListDowntimesResponse {
	this := ListDowntimesResponse{}
	return &this
}

// NewListDowntimesResponseWithDefaults instantiates a new ListDowntimesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListDowntimesResponseWithDefaults() *ListDowntimesResponse {
	this := ListDowntimesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListDowntimesResponse) GetData() []DowntimeResponseData {
	if o == nil || o.Data == nil {
		var ret []DowntimeResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDowntimesResponse) GetDataOk() (*[]DowntimeResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListDowntimesResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []DowntimeResponseData and assigns it to the Data field.
func (o *ListDowntimesResponse) SetData(v []DowntimeResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ListDowntimesResponse) GetIncluded() []DowntimeResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []DowntimeResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDowntimesResponse) GetIncludedOk() (*[]DowntimeResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ListDowntimesResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []DowntimeResponseIncludedItem and assigns it to the Included field.
func (o *ListDowntimesResponse) SetIncluded(v []DowntimeResponseIncludedItem) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListDowntimesResponse) GetMeta() DowntimeMeta {
	if o == nil || o.Meta == nil {
		var ret DowntimeMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDowntimesResponse) GetMetaOk() (*DowntimeMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListDowntimesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DowntimeMeta and assigns it to the Meta field.
func (o *ListDowntimesResponse) SetMeta(v DowntimeMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListDowntimesResponse) MarshalJSON() ([]byte, error) {
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
func (o *ListDowntimesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []DowntimeResponseData         `json:"data,omitempty"`
		Included []DowntimeResponseIncludedItem `json:"included,omitempty"`
		Meta     *DowntimeMeta                  `json:"meta,omitempty"`
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
