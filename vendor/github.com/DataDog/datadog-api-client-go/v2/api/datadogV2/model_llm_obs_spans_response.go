// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpansResponse Response containing a list of LLM Observability spans.
type LLMObsSpansResponse struct {
	// List of spans matching the query.
	Data []LLMObsSpanData `json:"data"`
	// Pagination links accompanying the spans response.
	Links *LLMObsSpansResponseLinks `json:"links,omitempty"`
	// Metadata accompanying the spans response.
	Meta LLMObsSpansResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpansResponse instantiates a new LLMObsSpansResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpansResponse(data []LLMObsSpanData, meta LLMObsSpansResponseMeta) *LLMObsSpansResponse {
	this := LLMObsSpansResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewLLMObsSpansResponseWithDefaults instantiates a new LLMObsSpansResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpansResponseWithDefaults() *LLMObsSpansResponse {
	this := LLMObsSpansResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *LLMObsSpansResponse) GetData() []LLMObsSpanData {
	if o == nil {
		var ret []LLMObsSpanData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponse) GetDataOk() (*[]LLMObsSpanData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *LLMObsSpansResponse) SetData(v []LLMObsSpanData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LLMObsSpansResponse) GetLinks() LLMObsSpansResponseLinks {
	if o == nil || o.Links == nil {
		var ret LLMObsSpansResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponse) GetLinksOk() (*LLMObsSpansResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LLMObsSpansResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given LLMObsSpansResponseLinks and assigns it to the Links field.
func (o *LLMObsSpansResponse) SetLinks(v LLMObsSpansResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value.
func (o *LLMObsSpansResponse) GetMeta() LLMObsSpansResponseMeta {
	if o == nil {
		var ret LLMObsSpansResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponse) GetMetaOk() (*LLMObsSpansResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *LLMObsSpansResponse) SetMeta(v LLMObsSpansResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpansResponse) MarshalJSON() ([]byte, error) {
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
func (o *LLMObsSpansResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]LLMObsSpanData         `json:"data"`
		Links *LLMObsSpansResponseLinks `json:"links,omitempty"`
		Meta  *LLMObsSpansResponseMeta  `json:"meta"`
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
