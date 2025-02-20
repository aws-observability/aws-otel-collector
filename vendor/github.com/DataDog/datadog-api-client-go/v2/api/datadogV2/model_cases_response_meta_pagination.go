// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CasesResponseMetaPagination Pagination metadata
type CasesResponseMetaPagination struct {
	// Current page number
	Current *int64 `json:"current,omitempty"`
	// Number of cases in current page
	Size *int64 `json:"size,omitempty"`
	// Total number of pages
	Total *int64 `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCasesResponseMetaPagination instantiates a new CasesResponseMetaPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCasesResponseMetaPagination() *CasesResponseMetaPagination {
	this := CasesResponseMetaPagination{}
	return &this
}

// NewCasesResponseMetaPaginationWithDefaults instantiates a new CasesResponseMetaPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCasesResponseMetaPaginationWithDefaults() *CasesResponseMetaPagination {
	this := CasesResponseMetaPagination{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *CasesResponseMetaPagination) GetCurrent() int64 {
	if o == nil || o.Current == nil {
		var ret int64
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesResponseMetaPagination) GetCurrentOk() (*int64, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *CasesResponseMetaPagination) HasCurrent() bool {
	return o != nil && o.Current != nil
}

// SetCurrent gets a reference to the given int64 and assigns it to the Current field.
func (o *CasesResponseMetaPagination) SetCurrent(v int64) {
	o.Current = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CasesResponseMetaPagination) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesResponseMetaPagination) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CasesResponseMetaPagination) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *CasesResponseMetaPagination) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CasesResponseMetaPagination) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CasesResponseMetaPagination) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CasesResponseMetaPagination) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *CasesResponseMetaPagination) SetTotal(v int64) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CasesResponseMetaPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CasesResponseMetaPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Current *int64 `json:"current,omitempty"`
		Size    *int64 `json:"size,omitempty"`
		Total   *int64 `json:"total,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current", "size", "total"})
	} else {
		return err
	}
	o.Current = all.Current
	o.Size = all.Size
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
