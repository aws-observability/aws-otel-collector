// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansListResponse Response object with all spans matching the request and pagination information.
type SpansListResponse struct {
	// Array of spans matching the request.
	Data []Span `json:"data,omitempty"`
	// Links attributes.
	Links *SpansListResponseLinks `json:"links,omitempty"`
	// The metadata associated with a request.
	Meta *SpansListResponseMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansListResponse instantiates a new SpansListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansListResponse() *SpansListResponse {
	this := SpansListResponse{}
	return &this
}

// NewSpansListResponseWithDefaults instantiates a new SpansListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansListResponseWithDefaults() *SpansListResponse {
	this := SpansListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SpansListResponse) GetData() []Span {
	if o == nil || o.Data == nil {
		var ret []Span
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansListResponse) GetDataOk() (*[]Span, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SpansListResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []Span and assigns it to the Data field.
func (o *SpansListResponse) SetData(v []Span) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SpansListResponse) GetLinks() SpansListResponseLinks {
	if o == nil || o.Links == nil {
		var ret SpansListResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansListResponse) GetLinksOk() (*SpansListResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SpansListResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given SpansListResponseLinks and assigns it to the Links field.
func (o *SpansListResponse) SetLinks(v SpansListResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SpansListResponse) GetMeta() SpansListResponseMetadata {
	if o == nil || o.Meta == nil {
		var ret SpansListResponseMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansListResponse) GetMetaOk() (*SpansListResponseMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SpansListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SpansListResponseMetadata and assigns it to the Meta field.
func (o *SpansListResponse) SetMeta(v SpansListResponseMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  []Span                     `json:"data,omitempty"`
		Links *SpansListResponseLinks    `json:"links,omitempty"`
		Meta  *SpansListResponseMetadata `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
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
