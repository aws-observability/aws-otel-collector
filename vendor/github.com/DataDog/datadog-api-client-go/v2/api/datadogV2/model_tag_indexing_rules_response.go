// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRulesResponse Response containing a page of tag indexing rules.
type TagIndexingRulesResponse struct {
	// Array of tag indexing rule objects.
	Data []TagIndexingRuleData `json:"data,omitempty"`
	// Pagination links. Only present if pagination query parameters were provided.
	Links *MetricsListResponseLinks `json:"links,omitempty"`
	// Pagination metadata for a list of tag indexing rules.
	Meta *TagIndexingRulesResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRulesResponse instantiates a new TagIndexingRulesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRulesResponse() *TagIndexingRulesResponse {
	this := TagIndexingRulesResponse{}
	return &this
}

// NewTagIndexingRulesResponseWithDefaults instantiates a new TagIndexingRulesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRulesResponseWithDefaults() *TagIndexingRulesResponse {
	this := TagIndexingRulesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TagIndexingRulesResponse) GetData() []TagIndexingRuleData {
	if o == nil || o.Data == nil {
		var ret []TagIndexingRuleData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRulesResponse) GetDataOk() (*[]TagIndexingRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TagIndexingRulesResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []TagIndexingRuleData and assigns it to the Data field.
func (o *TagIndexingRulesResponse) SetData(v []TagIndexingRuleData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TagIndexingRulesResponse) GetLinks() MetricsListResponseLinks {
	if o == nil || o.Links == nil {
		var ret MetricsListResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRulesResponse) GetLinksOk() (*MetricsListResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TagIndexingRulesResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given MetricsListResponseLinks and assigns it to the Links field.
func (o *TagIndexingRulesResponse) SetLinks(v MetricsListResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TagIndexingRulesResponse) GetMeta() TagIndexingRulesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret TagIndexingRulesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRulesResponse) GetMetaOk() (*TagIndexingRulesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TagIndexingRulesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given TagIndexingRulesResponseMeta and assigns it to the Meta field.
func (o *TagIndexingRulesResponse) SetMeta(v TagIndexingRulesResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRulesResponse) MarshalJSON() ([]byte, error) {
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
func (o *TagIndexingRulesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  []TagIndexingRuleData         `json:"data,omitempty"`
		Links *MetricsListResponseLinks     `json:"links,omitempty"`
		Meta  *TagIndexingRulesResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
