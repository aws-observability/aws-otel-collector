// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAppKeyRegistrationsResponseMeta The definition of `ListAppKeyRegistrationsResponseMeta` object.
type ListAppKeyRegistrationsResponseMeta struct {
	// The total number of app key registrations.
	Total *int64 `json:"total,omitempty"`
	// The total number of app key registrations that match the specified filters.
	TotalFiltered *int64 `json:"total_filtered,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAppKeyRegistrationsResponseMeta instantiates a new ListAppKeyRegistrationsResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAppKeyRegistrationsResponseMeta() *ListAppKeyRegistrationsResponseMeta {
	this := ListAppKeyRegistrationsResponseMeta{}
	return &this
}

// NewListAppKeyRegistrationsResponseMetaWithDefaults instantiates a new ListAppKeyRegistrationsResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAppKeyRegistrationsResponseMetaWithDefaults() *ListAppKeyRegistrationsResponseMeta {
	this := ListAppKeyRegistrationsResponseMeta{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ListAppKeyRegistrationsResponseMeta) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppKeyRegistrationsResponseMeta) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ListAppKeyRegistrationsResponseMeta) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ListAppKeyRegistrationsResponseMeta) SetTotal(v int64) {
	o.Total = &v
}

// GetTotalFiltered returns the TotalFiltered field value if set, zero value otherwise.
func (o *ListAppKeyRegistrationsResponseMeta) GetTotalFiltered() int64 {
	if o == nil || o.TotalFiltered == nil {
		var ret int64
		return ret
	}
	return *o.TotalFiltered
}

// GetTotalFilteredOk returns a tuple with the TotalFiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppKeyRegistrationsResponseMeta) GetTotalFilteredOk() (*int64, bool) {
	if o == nil || o.TotalFiltered == nil {
		return nil, false
	}
	return o.TotalFiltered, true
}

// HasTotalFiltered returns a boolean if a field has been set.
func (o *ListAppKeyRegistrationsResponseMeta) HasTotalFiltered() bool {
	return o != nil && o.TotalFiltered != nil
}

// SetTotalFiltered gets a reference to the given int64 and assigns it to the TotalFiltered field.
func (o *ListAppKeyRegistrationsResponseMeta) SetTotalFiltered(v int64) {
	o.TotalFiltered = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAppKeyRegistrationsResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.TotalFiltered != nil {
		toSerialize["total_filtered"] = o.TotalFiltered
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAppKeyRegistrationsResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total         *int64 `json:"total,omitempty"`
		TotalFiltered *int64 `json:"total_filtered,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total", "total_filtered"})
	} else {
		return err
	}
	o.Total = all.Total
	o.TotalFiltered = all.TotalFiltered

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
