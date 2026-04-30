// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListCampaignsResponse Response containing a list of campaigns.
type ListCampaignsResponse struct {
	// Array of campaigns.
	Data []CampaignResponseData `json:"data"`
	// Metadata for scores response.
	Meta PaginatedResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListCampaignsResponse instantiates a new ListCampaignsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListCampaignsResponse(data []CampaignResponseData, meta PaginatedResponseMeta) *ListCampaignsResponse {
	this := ListCampaignsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewListCampaignsResponseWithDefaults instantiates a new ListCampaignsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListCampaignsResponseWithDefaults() *ListCampaignsResponse {
	this := ListCampaignsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListCampaignsResponse) GetData() []CampaignResponseData {
	if o == nil {
		var ret []CampaignResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCampaignsResponse) GetDataOk() (*[]CampaignResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListCampaignsResponse) SetData(v []CampaignResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *ListCampaignsResponse) GetMeta() PaginatedResponseMeta {
	if o == nil {
		var ret PaginatedResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListCampaignsResponse) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ListCampaignsResponse) SetMeta(v PaginatedResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListCampaignsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListCampaignsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]CampaignResponseData `json:"data"`
		Meta *PaginatedResponseMeta  `json:"meta"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
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
