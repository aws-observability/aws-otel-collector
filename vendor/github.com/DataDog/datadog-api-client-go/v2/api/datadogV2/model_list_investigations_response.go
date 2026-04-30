// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListInvestigationsResponse Response for listing investigations.
type ListInvestigationsResponse struct {
	// List of investigations.
	Data []ListInvestigationsResponseData `json:"data"`
	// Pagination links for the list investigations response.
	Links ListInvestigationsResponseLinks `json:"links"`
	// Metadata for the list investigations response.
	Meta ListInvestigationsResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListInvestigationsResponse instantiates a new ListInvestigationsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListInvestigationsResponse(data []ListInvestigationsResponseData, links ListInvestigationsResponseLinks, meta ListInvestigationsResponseMeta) *ListInvestigationsResponse {
	this := ListInvestigationsResponse{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewListInvestigationsResponseWithDefaults instantiates a new ListInvestigationsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListInvestigationsResponseWithDefaults() *ListInvestigationsResponse {
	this := ListInvestigationsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListInvestigationsResponse) GetData() []ListInvestigationsResponseData {
	if o == nil {
		var ret []ListInvestigationsResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponse) GetDataOk() (*[]ListInvestigationsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListInvestigationsResponse) SetData(v []ListInvestigationsResponseData) {
	o.Data = v
}

// GetLinks returns the Links field value.
func (o *ListInvestigationsResponse) GetLinks() ListInvestigationsResponseLinks {
	if o == nil {
		var ret ListInvestigationsResponseLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponse) GetLinksOk() (*ListInvestigationsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value.
func (o *ListInvestigationsResponse) SetLinks(v ListInvestigationsResponseLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value.
func (o *ListInvestigationsResponse) GetMeta() ListInvestigationsResponseMeta {
	if o == nil {
		var ret ListInvestigationsResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponse) GetMetaOk() (*ListInvestigationsResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ListInvestigationsResponse) SetMeta(v ListInvestigationsResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListInvestigationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListInvestigationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]ListInvestigationsResponseData `json:"data"`
		Links *ListInvestigationsResponseLinks  `json:"links"`
		Meta  *ListInvestigationsResponseMeta   `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Links == nil {
		return fmt.Errorf("required field links missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = *all.Links
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
