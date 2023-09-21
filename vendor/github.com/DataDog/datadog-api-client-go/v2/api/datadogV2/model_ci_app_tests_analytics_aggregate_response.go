// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
=======
	"github.com/goccy/go-json"
>>>>>>> main

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppTestsAnalyticsAggregateResponse The response object for the test events aggregate API endpoint.
type CIAppTestsAnalyticsAggregateResponse struct {
	// The query results.
	Data *CIAppTestsAggregationBucketsResponse `json:"data,omitempty"`
	// Links attributes.
	Links *CIAppResponseLinks `json:"links,omitempty"`
	// The metadata associated with a request.
<<<<<<< HEAD
	Meta *CIAppResponseMetadata `json:"meta,omitempty"`
=======
	Meta *CIAppResponseMetadataWithPagination `json:"meta,omitempty"`
>>>>>>> main
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppTestsAnalyticsAggregateResponse instantiates a new CIAppTestsAnalyticsAggregateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppTestsAnalyticsAggregateResponse() *CIAppTestsAnalyticsAggregateResponse {
	this := CIAppTestsAnalyticsAggregateResponse{}
	return &this
}

// NewCIAppTestsAnalyticsAggregateResponseWithDefaults instantiates a new CIAppTestsAnalyticsAggregateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppTestsAnalyticsAggregateResponseWithDefaults() *CIAppTestsAnalyticsAggregateResponse {
	this := CIAppTestsAnalyticsAggregateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CIAppTestsAnalyticsAggregateResponse) GetData() CIAppTestsAggregationBucketsResponse {
	if o == nil || o.Data == nil {
		var ret CIAppTestsAggregationBucketsResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestsAnalyticsAggregateResponse) GetDataOk() (*CIAppTestsAggregationBucketsResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CIAppTestsAnalyticsAggregateResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given CIAppTestsAggregationBucketsResponse and assigns it to the Data field.
func (o *CIAppTestsAnalyticsAggregateResponse) SetData(v CIAppTestsAggregationBucketsResponse) {
	o.Data = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CIAppTestsAnalyticsAggregateResponse) GetLinks() CIAppResponseLinks {
	if o == nil || o.Links == nil {
		var ret CIAppResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppTestsAnalyticsAggregateResponse) GetLinksOk() (*CIAppResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CIAppTestsAnalyticsAggregateResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given CIAppResponseLinks and assigns it to the Links field.
func (o *CIAppTestsAnalyticsAggregateResponse) SetLinks(v CIAppResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *CIAppTestsAnalyticsAggregateResponse) GetMeta() CIAppResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret CIAppResponseMetadata
=======
func (o *CIAppTestsAnalyticsAggregateResponse) GetMeta() CIAppResponseMetadataWithPagination {
	if o == nil || o.Meta == nil {
		var ret CIAppResponseMetadataWithPagination
>>>>>>> main
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *CIAppTestsAnalyticsAggregateResponse) GetMetaOk() (*CIAppResponseMetadata, bool) {
=======
func (o *CIAppTestsAnalyticsAggregateResponse) GetMetaOk() (*CIAppResponseMetadataWithPagination, bool) {
>>>>>>> main
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CIAppTestsAnalyticsAggregateResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

<<<<<<< HEAD
// SetMeta gets a reference to the given CIAppResponseMetadata and assigns it to the Meta field.
func (o *CIAppTestsAnalyticsAggregateResponse) SetMeta(v CIAppResponseMetadata) {
=======
// SetMeta gets a reference to the given CIAppResponseMetadataWithPagination and assigns it to the Meta field.
func (o *CIAppTestsAnalyticsAggregateResponse) SetMeta(v CIAppResponseMetadataWithPagination) {
>>>>>>> main
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppTestsAnalyticsAggregateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
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
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppTestsAnalyticsAggregateResponse) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
	all := struct {
		Data  *CIAppTestsAggregationBucketsResponse `json:"data,omitempty"`
		Links *CIAppResponseLinks                   `json:"links,omitempty"`
		Meta  *CIAppResponseMetadata                `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
=======
	all := struct {
		Data  *CIAppTestsAggregationBucketsResponse `json:"data,omitempty"`
		Links *CIAppResponseLinks                   `json:"links,omitempty"`
		Meta  *CIAppResponseMetadataWithPagination  `json:"meta,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
>>>>>>> main
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}
<<<<<<< HEAD
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Meta = all.Meta
=======

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

>>>>>>> main
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

<<<<<<< HEAD
=======
	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

>>>>>>> main
	return nil
}
