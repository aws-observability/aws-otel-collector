// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentHostsPage Pagination details for the list of hosts in a deployment.
type FleetDeploymentHostsPage struct {
	// Current page index (zero-based).
	CurrentPage *int64 `json:"current_page,omitempty"`
	// Number of hosts returned per page.
	PageSize *int64 `json:"page_size,omitempty"`
	// Total number of hosts in this deployment.
	TotalHosts *int64 `json:"total_hosts,omitempty"`
	// Total number of pages available.
	TotalPages *int64 `json:"total_pages,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentHostsPage instantiates a new FleetDeploymentHostsPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentHostsPage() *FleetDeploymentHostsPage {
	this := FleetDeploymentHostsPage{}
	return &this
}

// NewFleetDeploymentHostsPageWithDefaults instantiates a new FleetDeploymentHostsPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentHostsPageWithDefaults() *FleetDeploymentHostsPage {
	this := FleetDeploymentHostsPage{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *FleetDeploymentHostsPage) GetCurrentPage() int64 {
	if o == nil || o.CurrentPage == nil {
		var ret int64
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostsPage) GetCurrentPageOk() (*int64, bool) {
	if o == nil || o.CurrentPage == nil {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *FleetDeploymentHostsPage) HasCurrentPage() bool {
	return o != nil && o.CurrentPage != nil
}

// SetCurrentPage gets a reference to the given int64 and assigns it to the CurrentPage field.
func (o *FleetDeploymentHostsPage) SetCurrentPage(v int64) {
	o.CurrentPage = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *FleetDeploymentHostsPage) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostsPage) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *FleetDeploymentHostsPage) HasPageSize() bool {
	return o != nil && o.PageSize != nil
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *FleetDeploymentHostsPage) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetTotalHosts returns the TotalHosts field value if set, zero value otherwise.
func (o *FleetDeploymentHostsPage) GetTotalHosts() int64 {
	if o == nil || o.TotalHosts == nil {
		var ret int64
		return ret
	}
	return *o.TotalHosts
}

// GetTotalHostsOk returns a tuple with the TotalHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostsPage) GetTotalHostsOk() (*int64, bool) {
	if o == nil || o.TotalHosts == nil {
		return nil, false
	}
	return o.TotalHosts, true
}

// HasTotalHosts returns a boolean if a field has been set.
func (o *FleetDeploymentHostsPage) HasTotalHosts() bool {
	return o != nil && o.TotalHosts != nil
}

// SetTotalHosts gets a reference to the given int64 and assigns it to the TotalHosts field.
func (o *FleetDeploymentHostsPage) SetTotalHosts(v int64) {
	o.TotalHosts = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *FleetDeploymentHostsPage) GetTotalPages() int64 {
	if o == nil || o.TotalPages == nil {
		var ret int64
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostsPage) GetTotalPagesOk() (*int64, bool) {
	if o == nil || o.TotalPages == nil {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *FleetDeploymentHostsPage) HasTotalPages() bool {
	return o != nil && o.TotalPages != nil
}

// SetTotalPages gets a reference to the given int64 and assigns it to the TotalPages field.
func (o *FleetDeploymentHostsPage) SetTotalPages(v int64) {
	o.TotalPages = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentHostsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CurrentPage != nil {
		toSerialize["current_page"] = o.CurrentPage
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.TotalHosts != nil {
		toSerialize["total_hosts"] = o.TotalHosts
	}
	if o.TotalPages != nil {
		toSerialize["total_pages"] = o.TotalPages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentHostsPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentPage *int64 `json:"current_page,omitempty"`
		PageSize    *int64 `json:"page_size,omitempty"`
		TotalHosts  *int64 `json:"total_hosts,omitempty"`
		TotalPages  *int64 `json:"total_pages,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current_page", "page_size", "total_hosts", "total_pages"})
	} else {
		return err
	}
	o.CurrentPage = all.CurrentPage
	o.PageSize = all.PageSize
	o.TotalHosts = all.TotalHosts
	o.TotalPages = all.TotalPages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
