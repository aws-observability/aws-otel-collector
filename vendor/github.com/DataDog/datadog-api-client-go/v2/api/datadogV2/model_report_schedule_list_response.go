// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleListResponse Response containing a list of report schedules.
type ReportScheduleListResponse struct {
	// The list of report schedules.
	Data []ReportScheduleListResponseData `json:"data"`
	// Related resources included with the report schedules, such as authors and rendered resources.
	Included []ReportScheduleIncludedResource `json:"included,omitempty"`
	// Pagination links for navigating a report schedule list response.
	Links *ReportScheduleListResponseLinks `json:"links,omitempty"`
	// Metadata for a paginated report schedule list response.
	Meta *ReportScheduleListResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleListResponse instantiates a new ReportScheduleListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleListResponse(data []ReportScheduleListResponseData) *ReportScheduleListResponse {
	this := ReportScheduleListResponse{}
	this.Data = data
	return &this
}

// NewReportScheduleListResponseWithDefaults instantiates a new ReportScheduleListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleListResponseWithDefaults() *ReportScheduleListResponse {
	this := ReportScheduleListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ReportScheduleListResponse) GetData() []ReportScheduleListResponseData {
	if o == nil {
		var ret []ReportScheduleListResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponse) GetDataOk() (*[]ReportScheduleListResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ReportScheduleListResponse) SetData(v []ReportScheduleListResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ReportScheduleListResponse) GetIncluded() []ReportScheduleIncludedResource {
	if o == nil || o.Included == nil {
		var ret []ReportScheduleIncludedResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponse) GetIncludedOk() (*[]ReportScheduleIncludedResource, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ReportScheduleListResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []ReportScheduleIncludedResource and assigns it to the Included field.
func (o *ReportScheduleListResponse) SetIncluded(v []ReportScheduleIncludedResource) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReportScheduleListResponse) GetLinks() ReportScheduleListResponseLinks {
	if o == nil || o.Links == nil {
		var ret ReportScheduleListResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponse) GetLinksOk() (*ReportScheduleListResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReportScheduleListResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given ReportScheduleListResponseLinks and assigns it to the Links field.
func (o *ReportScheduleListResponse) SetLinks(v ReportScheduleListResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReportScheduleListResponse) GetMeta() ReportScheduleListResponseMeta {
	if o == nil || o.Meta == nil {
		var ret ReportScheduleListResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponse) GetMetaOk() (*ReportScheduleListResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReportScheduleListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ReportScheduleListResponseMeta and assigns it to the Meta field.
func (o *ReportScheduleListResponse) SetMeta(v ReportScheduleListResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
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
func (o *ReportScheduleListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]ReportScheduleListResponseData `json:"data"`
		Included []ReportScheduleIncludedResource  `json:"included,omitempty"`
		Links    *ReportScheduleListResponseLinks  `json:"links,omitempty"`
		Meta     *ReportScheduleListResponseMeta   `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
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
