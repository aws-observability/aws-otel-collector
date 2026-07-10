// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListSharedDashboardsResponse Response containing shared dashboards for a dashboard.
type ListSharedDashboardsResponse struct {
	// Shared dashboards for the dashboard.
	Data []SharedDashboardResponse `json:"data"`
	// Users and dashboards related to the shared dashboards.
	Included []SharedDashboardIncluded `json:"included"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListSharedDashboardsResponse instantiates a new ListSharedDashboardsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListSharedDashboardsResponse(data []SharedDashboardResponse, included []SharedDashboardIncluded) *ListSharedDashboardsResponse {
	this := ListSharedDashboardsResponse{}
	this.Data = data
	this.Included = included
	return &this
}

// NewListSharedDashboardsResponseWithDefaults instantiates a new ListSharedDashboardsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListSharedDashboardsResponseWithDefaults() *ListSharedDashboardsResponse {
	this := ListSharedDashboardsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListSharedDashboardsResponse) GetData() []SharedDashboardResponse {
	if o == nil {
		var ret []SharedDashboardResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListSharedDashboardsResponse) GetDataOk() (*[]SharedDashboardResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListSharedDashboardsResponse) SetData(v []SharedDashboardResponse) {
	o.Data = v
}

// GetIncluded returns the Included field value.
func (o *ListSharedDashboardsResponse) GetIncluded() []SharedDashboardIncluded {
	if o == nil {
		var ret []SharedDashboardIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
func (o *ListSharedDashboardsResponse) GetIncludedOk() (*[]SharedDashboardIncluded, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Included, true
}

// SetIncluded sets field value.
func (o *ListSharedDashboardsResponse) SetIncluded(v []SharedDashboardIncluded) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListSharedDashboardsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["included"] = o.Included

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListSharedDashboardsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]SharedDashboardResponse `json:"data"`
		Included *[]SharedDashboardIncluded `json:"included"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Included == nil {
		return fmt.Errorf("required field included missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}
	o.Data = *all.Data
	o.Included = *all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
