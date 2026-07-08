// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationSearchDataAttributesRequest Attributes for an experimentation search request.
type LLMObsExperimentationSearchDataAttributesRequest struct {
	// Options to control content preview truncation.
	ContentPreview *LLMObsExperimentationContentPreview `json:"content_preview,omitempty"`
	// Filter criteria for an experimentation search request.
	Filter LLMObsExperimentationFilter `json:"filter"`
	// Additional data to include in the response.
	Include *LLMObsExperimentationInclude `json:"include,omitempty"`
	// Cursor-based pagination parameters.
	Page *LLMObsExperimentationCursorPage `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationSearchDataAttributesRequest instantiates a new LLMObsExperimentationSearchDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationSearchDataAttributesRequest(filter LLMObsExperimentationFilter) *LLMObsExperimentationSearchDataAttributesRequest {
	this := LLMObsExperimentationSearchDataAttributesRequest{}
	this.Filter = filter
	return &this
}

// NewLLMObsExperimentationSearchDataAttributesRequestWithDefaults instantiates a new LLMObsExperimentationSearchDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationSearchDataAttributesRequestWithDefaults() *LLMObsExperimentationSearchDataAttributesRequest {
	this := LLMObsExperimentationSearchDataAttributesRequest{}
	return &this
}

// GetContentPreview returns the ContentPreview field value if set, zero value otherwise.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetContentPreview() LLMObsExperimentationContentPreview {
	if o == nil || o.ContentPreview == nil {
		var ret LLMObsExperimentationContentPreview
		return ret
	}
	return *o.ContentPreview
}

// GetContentPreviewOk returns a tuple with the ContentPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetContentPreviewOk() (*LLMObsExperimentationContentPreview, bool) {
	if o == nil || o.ContentPreview == nil {
		return nil, false
	}
	return o.ContentPreview, true
}

// HasContentPreview returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) HasContentPreview() bool {
	return o != nil && o.ContentPreview != nil
}

// SetContentPreview gets a reference to the given LLMObsExperimentationContentPreview and assigns it to the ContentPreview field.
func (o *LLMObsExperimentationSearchDataAttributesRequest) SetContentPreview(v LLMObsExperimentationContentPreview) {
	o.ContentPreview = &v
}

// GetFilter returns the Filter field value.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetFilter() LLMObsExperimentationFilter {
	if o == nil {
		var ret LLMObsExperimentationFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetFilterOk() (*LLMObsExperimentationFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *LLMObsExperimentationSearchDataAttributesRequest) SetFilter(v LLMObsExperimentationFilter) {
	o.Filter = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetInclude() LLMObsExperimentationInclude {
	if o == nil || o.Include == nil {
		var ret LLMObsExperimentationInclude
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetIncludeOk() (*LLMObsExperimentationInclude, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) HasInclude() bool {
	return o != nil && o.Include != nil
}

// SetInclude gets a reference to the given LLMObsExperimentationInclude and assigns it to the Include field.
func (o *LLMObsExperimentationSearchDataAttributesRequest) SetInclude(v LLMObsExperimentationInclude) {
	o.Include = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetPage() LLMObsExperimentationCursorPage {
	if o == nil || o.Page == nil {
		var ret LLMObsExperimentationCursorPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) GetPageOk() (*LLMObsExperimentationCursorPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchDataAttributesRequest) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given LLMObsExperimentationCursorPage and assigns it to the Page field.
func (o *LLMObsExperimentationSearchDataAttributesRequest) SetPage(v LLMObsExperimentationCursorPage) {
	o.Page = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationSearchDataAttributesRequest) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationSearchDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentPreview *LLMObsExperimentationContentPreview `json:"content_preview,omitempty"`
		Filter         *LLMObsExperimentationFilter         `json:"filter"`
		Include        *LLMObsExperimentationInclude        `json:"include,omitempty"`
		Page           *LLMObsExperimentationCursorPage     `json:"page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content_preview", "filter", "include", "page"})
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
