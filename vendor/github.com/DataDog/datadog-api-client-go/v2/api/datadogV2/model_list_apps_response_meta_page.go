// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAppsResponseMetaPage Information on the total number of apps, to be used for pagination.
type ListAppsResponseMetaPage struct {
	// The total number of apps under the Datadog organization, disregarding any filters applied.
	TotalCount *int64 `json:"totalCount,omitempty"`
	// The total number of apps that match the specified filters.
	TotalFilteredCount *int64 `json:"totalFilteredCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAppsResponseMetaPage instantiates a new ListAppsResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAppsResponseMetaPage() *ListAppsResponseMetaPage {
	this := ListAppsResponseMetaPage{}
	return &this
}

// NewListAppsResponseMetaPageWithDefaults instantiates a new ListAppsResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAppsResponseMetaPageWithDefaults() *ListAppsResponseMetaPage {
	this := ListAppsResponseMetaPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListAppsResponseMetaPage) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseMetaPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListAppsResponseMetaPage) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ListAppsResponseMetaPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
func (o *ListAppsResponseMetaPage) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseMetaPage) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *ListAppsResponseMetaPage) HasTotalFilteredCount() bool {
	return o != nil && o.TotalFilteredCount != nil
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
func (o *ListAppsResponseMetaPage) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAppsResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.TotalFilteredCount != nil {
		toSerialize["totalFilteredCount"] = o.TotalFilteredCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAppsResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount         *int64 `json:"totalCount,omitempty"`
		TotalFilteredCount *int64 `json:"totalFilteredCount,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"totalCount", "totalFilteredCount"})
	} else {
		return err
	}
	o.TotalCount = all.TotalCount
	o.TotalFilteredCount = all.TotalFilteredCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
