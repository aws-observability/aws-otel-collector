// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CSMAgentsMetadata Metadata related to the paginated response.
type CSMAgentsMetadata struct {
	// The index of the current page in the paginated results.
	PageIndex *int64 `json:"page_index,omitempty"`
	// The number of items per page in the paginated results.
	PageSize *int64 `json:"page_size,omitempty"`
	// Total number of items that match the filter criteria.
	TotalFiltered *int64 `json:"total_filtered,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCSMAgentsMetadata instantiates a new CSMAgentsMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCSMAgentsMetadata() *CSMAgentsMetadata {
	this := CSMAgentsMetadata{}
	return &this
}

// NewCSMAgentsMetadataWithDefaults instantiates a new CSMAgentsMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCSMAgentsMetadataWithDefaults() *CSMAgentsMetadata {
	this := CSMAgentsMetadata{}
	return &this
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise.
func (o *CSMAgentsMetadata) GetPageIndex() int64 {
	if o == nil || o.PageIndex == nil {
		var ret int64
		return ret
	}
	return *o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSMAgentsMetadata) GetPageIndexOk() (*int64, bool) {
	if o == nil || o.PageIndex == nil {
		return nil, false
	}
	return o.PageIndex, true
}

// HasPageIndex returns a boolean if a field has been set.
func (o *CSMAgentsMetadata) HasPageIndex() bool {
	return o != nil && o.PageIndex != nil
}

// SetPageIndex gets a reference to the given int64 and assigns it to the PageIndex field.
func (o *CSMAgentsMetadata) SetPageIndex(v int64) {
	o.PageIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CSMAgentsMetadata) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSMAgentsMetadata) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CSMAgentsMetadata) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *CSMAgentsMetadata) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetTotalFiltered returns the TotalFiltered field value if set, zero value otherwise.
func (o *CSMAgentsMetadata) GetTotalFiltered() int64 {
	if o == nil || o.TotalFiltered == nil {
		var ret int64
		return ret
	}
	return *o.TotalFiltered
}

// GetTotalFilteredOk returns a tuple with the TotalFiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSMAgentsMetadata) GetTotalFilteredOk() (*int64, bool) {
	if o == nil || o.TotalFiltered == nil {
		return nil, false
	}
	return o.TotalFiltered, true
}

// HasTotalFiltered returns a boolean if a field has been set.
func (o *CSMAgentsMetadata) HasTotalFiltered() bool {
	return o != nil && o.TotalFiltered != nil
}

// SetTotalFiltered gets a reference to the given int64 and assigns it to the TotalFiltered field.
func (o *CSMAgentsMetadata) SetTotalFiltered(v int64) {
	o.TotalFiltered = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CSMAgentsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PageIndex != nil {
		toSerialize["page_index"] = o.PageIndex
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
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
func (o *CSMAgentsMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PageIndex     *int64 `json:"page_index,omitempty"`
		PageSize      *int64 `json:"page_size,omitempty"`
		TotalFiltered *int64 `json:"total_filtered,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page_index", "page_size", "total_filtered"})
	} else {
		return err
	}
	o.PageIndex = all.PageIndex
	o.PageSize = all.PageSize
	o.TotalFiltered = all.TotalFiltered

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
