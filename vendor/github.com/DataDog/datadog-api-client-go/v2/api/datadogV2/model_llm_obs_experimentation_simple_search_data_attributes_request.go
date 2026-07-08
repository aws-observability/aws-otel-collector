// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationSimpleSearchDataAttributesRequest Attributes for an experimentation simple search request.
type LLMObsExperimentationSimpleSearchDataAttributesRequest struct {
	// Options to control content preview truncation.
	ContentPreview *LLMObsExperimentationContentPreview `json:"content_preview,omitempty"`
	// Filter criteria for an experimentation search request.
	Filter LLMObsExperimentationFilter `json:"filter"`
	// Additional data to include in the response.
	Include *LLMObsExperimentationInclude `json:"include,omitempty"`
	// Offset-based pagination parameters for simple search.
	Page *LLMObsExperimentationNumberPage `json:"page,omitempty"`
	// Sort order for results.
	Sort []LLMObsExperimentationSortField `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationSimpleSearchDataAttributesRequest instantiates a new LLMObsExperimentationSimpleSearchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationSimpleSearchDataAttributesRequest(filter LLMObsExperimentationFilter) *LLMObsExperimentationSimpleSearchDataAttributesRequest {
	this := LLMObsExperimentationSimpleSearchDataAttributesRequest{}
	this.Filter = filter
	return &this
}

// NewLLMObsExperimentationSimpleSearchDataAttributesRequestWithDefaults instantiates a new LLMObsExperimentationSimpleSearchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationSimpleSearchDataAttributesRequestWithDefaults() *LLMObsExperimentationSimpleSearchDataAttributesRequest {
	this := LLMObsExperimentationSimpleSearchDataAttributesRequest{}
	return &this
}

// GetContentPreview returns the ContentPreview field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetContentPreview() LLMObsExperimentationContentPreview {
	if o == nil || o.ContentPreview == nil {
		var ret LLMObsExperimentationContentPreview
		return ret
	}
	return *o.ContentPreview
}

// GetContentPreviewOk returns a tuple with the ContentPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetContentPreviewOk() (*LLMObsExperimentationContentPreview, bool) {
	if o == nil || o.ContentPreview == nil {
		return nil, false
	}
	return o.ContentPreview, true
}

// HasContentPreview returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) HasContentPreview() bool {
	return o != nil && o.ContentPreview != nil
}

// SetContentPreview gets a reference to the given LLMObsExperimentationContentPreview and assigns it to the ContentPreview field.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) SetContentPreview(v LLMObsExperimentationContentPreview) {
	o.ContentPreview = &v
}

// GetFilter returns the Filter field value.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetFilter() LLMObsExperimentationFilter {
	if o == nil {
		var ret LLMObsExperimentationFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetFilterOk() (*LLMObsExperimentationFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) SetFilter(v LLMObsExperimentationFilter) {
	o.Filter = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetInclude() LLMObsExperimentationInclude {
	if o == nil || o.Include == nil {
		var ret LLMObsExperimentationInclude
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetIncludeOk() (*LLMObsExperimentationInclude, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) HasInclude() bool {
	return o != nil && o.Include != nil
}

// SetInclude gets a reference to the given LLMObsExperimentationInclude and assigns it to the Include field.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) SetInclude(v LLMObsExperimentationInclude) {
	o.Include = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetPage() LLMObsExperimentationNumberPage {
	if o == nil || o.Page == nil {
		var ret LLMObsExperimentationNumberPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetPageOk() (*LLMObsExperimentationNumberPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given LLMObsExperimentationNumberPage and assigns it to the Page field.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) SetPage(v LLMObsExperimentationNumberPage) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetSort() []LLMObsExperimentationSortField {
	if o == nil || o.Sort == nil {
		var ret []LLMObsExperimentationSortField
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) GetSortOk() (*[]LLMObsExperimentationSortField, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given []LLMObsExperimentationSortField and assigns it to the Sort field.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) SetSort(v []LLMObsExperimentationSortField) {
	o.Sort = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationSimpleSearchDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContentPreview != nil {
		toSerialize["content_preview"] = o.ContentPreview
	}
	toSerialize["filter"] = o.Filter
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationSimpleSearchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentPreview *LLMObsExperimentationContentPreview `json:"content_preview,omitempty"`
		Filter         *LLMObsExperimentationFilter         `json:"filter"`
		Include        *LLMObsExperimentationInclude        `json:"include,omitempty"`
		Page           *LLMObsExperimentationNumberPage     `json:"page,omitempty"`
		Sort           []LLMObsExperimentationSortField     `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content_preview", "filter", "include", "page", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ContentPreview != nil && all.ContentPreview.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ContentPreview = all.ContentPreview
	if all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = *all.Filter
	if all.Include != nil && all.Include.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Include = all.Include
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	o.Sort = all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
