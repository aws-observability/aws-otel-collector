// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListDashboardsUsageResponse Paginated list of dashboard usage records.
type ListDashboardsUsageResponse struct {
	// Dashboard usage records, one per dashboard in the caller's organization.
	Data []DashboardUsage `json:"data"`
	// Pagination links for a list of dashboard usage records.
	Links *ListDashboardsUsageResponseLinks `json:"links,omitempty"`
	// Pagination metadata for a list of dashboard usage records.
	Meta ListDashboardsUsageResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListDashboardsUsageResponse instantiates a new ListDashboardsUsageResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListDashboardsUsageResponse(data []DashboardUsage, meta ListDashboardsUsageResponseMeta) *ListDashboardsUsageResponse {
	this := ListDashboardsUsageResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewListDashboardsUsageResponseWithDefaults instantiates a new ListDashboardsUsageResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListDashboardsUsageResponseWithDefaults() *ListDashboardsUsageResponse {
	this := ListDashboardsUsageResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListDashboardsUsageResponse) GetData() []DashboardUsage {
	if o == nil {
		var ret []DashboardUsage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListDashboardsUsageResponse) GetDataOk() (*[]DashboardUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListDashboardsUsageResponse) SetData(v []DashboardUsage) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListDashboardsUsageResponse) GetLinks() ListDashboardsUsageResponseLinks {
	if o == nil || o.Links == nil {
		var ret ListDashboardsUsageResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDashboardsUsageResponse) GetLinksOk() (*ListDashboardsUsageResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListDashboardsUsageResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ListDashboardsUsageResponseLinks and assigns it to the Links field.
func (o *ListDashboardsUsageResponse) SetLinks(v ListDashboardsUsageResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value.
func (o *ListDashboardsUsageResponse) GetMeta() ListDashboardsUsageResponseMeta {
	if o == nil {
		var ret ListDashboardsUsageResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListDashboardsUsageResponse) GetMetaOk() (*ListDashboardsUsageResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ListDashboardsUsageResponse) SetMeta(v ListDashboardsUsageResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListDashboardsUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListDashboardsUsageResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]DashboardUsage                 `json:"data"`
		Links *ListDashboardsUsageResponseLinks `json:"links,omitempty"`
		Meta  *ListDashboardsUsageResponseMeta  `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
