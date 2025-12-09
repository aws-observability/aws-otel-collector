// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoRequestDataAttributes
type FacetInfoRequestDataAttributes struct {
	//
	FacetId string `json:"facet_id"`
	//
	Limit int64 `json:"limit"`
	//
	Search *FacetInfoRequestDataAttributesSearch `json:"search,omitempty"`
	//
	TermSearch *FacetInfoRequestDataAttributesTermSearch `json:"term_search,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetInfoRequestDataAttributes instantiates a new FacetInfoRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetInfoRequestDataAttributes(facetId string, limit int64) *FacetInfoRequestDataAttributes {
	this := FacetInfoRequestDataAttributes{}
	this.FacetId = facetId
	this.Limit = limit
	return &this
}

// NewFacetInfoRequestDataAttributesWithDefaults instantiates a new FacetInfoRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetInfoRequestDataAttributesWithDefaults() *FacetInfoRequestDataAttributes {
	this := FacetInfoRequestDataAttributes{}
	return &this
}

// GetFacetId returns the FacetId field value.
func (o *FacetInfoRequestDataAttributes) GetFacetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FacetId
}

// GetFacetIdOk returns a tuple with the FacetId field value
// and a boolean to check if the value has been set.
func (o *FacetInfoRequestDataAttributes) GetFacetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetId, true
}

// SetFacetId sets field value.
func (o *FacetInfoRequestDataAttributes) SetFacetId(v string) {
	o.FacetId = v
}

// GetLimit returns the Limit field value.
func (o *FacetInfoRequestDataAttributes) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *FacetInfoRequestDataAttributes) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *FacetInfoRequestDataAttributes) SetLimit(v int64) {
	o.Limit = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FacetInfoRequestDataAttributes) GetSearch() FacetInfoRequestDataAttributesSearch {
	if o == nil || o.Search == nil {
		var ret FacetInfoRequestDataAttributesSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoRequestDataAttributes) GetSearchOk() (*FacetInfoRequestDataAttributesSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FacetInfoRequestDataAttributes) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FacetInfoRequestDataAttributesSearch and assigns it to the Search field.
func (o *FacetInfoRequestDataAttributes) SetSearch(v FacetInfoRequestDataAttributesSearch) {
	o.Search = &v
}

// GetTermSearch returns the TermSearch field value if set, zero value otherwise.
func (o *FacetInfoRequestDataAttributes) GetTermSearch() FacetInfoRequestDataAttributesTermSearch {
	if o == nil || o.TermSearch == nil {
		var ret FacetInfoRequestDataAttributesTermSearch
		return ret
	}
	return *o.TermSearch
}

// GetTermSearchOk returns a tuple with the TermSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoRequestDataAttributes) GetTermSearchOk() (*FacetInfoRequestDataAttributesTermSearch, bool) {
	if o == nil || o.TermSearch == nil {
		return nil, false
	}
	return o.TermSearch, true
}

// HasTermSearch returns a boolean if a field has been set.
func (o *FacetInfoRequestDataAttributes) HasTermSearch() bool {
	return o != nil && o.TermSearch != nil
}

// SetTermSearch gets a reference to the given FacetInfoRequestDataAttributesTermSearch and assigns it to the TermSearch field.
func (o *FacetInfoRequestDataAttributes) SetTermSearch(v FacetInfoRequestDataAttributesTermSearch) {
	o.TermSearch = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetInfoRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet_id"] = o.FacetId
	toSerialize["limit"] = o.Limit
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.TermSearch != nil {
		toSerialize["term_search"] = o.TermSearch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FacetInfoRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FacetId    *string                                   `json:"facet_id"`
		Limit      *int64                                    `json:"limit"`
		Search     *FacetInfoRequestDataAttributesSearch     `json:"search,omitempty"`
		TermSearch *FacetInfoRequestDataAttributesTermSearch `json:"term_search,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FacetId == nil {
		return fmt.Errorf("required field facet_id missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet_id", "limit", "search", "term_search"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FacetId = *all.FacetId
	o.Limit = *all.Limit
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.TermSearch != nil && all.TermSearch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TermSearch = all.TermSearch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
