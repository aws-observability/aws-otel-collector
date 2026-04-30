// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingsSearchRequestDataAttributes Request attributes for searching security findings.
type SecurityFindingsSearchRequestDataAttributes struct {
	// The search query following log search syntax.
	Filter *string `json:"filter,omitempty"`
	// Pagination attributes for the search request.
	Page *SecurityFindingsSearchRequestPage `json:"page,omitempty"`
	// The sort parameters when querying security findings.
	Sort *SecurityFindingsSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityFindingsSearchRequestDataAttributes instantiates a new SecurityFindingsSearchRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityFindingsSearchRequestDataAttributes() *SecurityFindingsSearchRequestDataAttributes {
	this := SecurityFindingsSearchRequestDataAttributes{}
	var filter string = "*"
	this.Filter = &filter
	var sort SecurityFindingsSort = SECURITYFINDINGSSORT_DETECTION_CHANGED_AT_DESC
	this.Sort = &sort
	return &this
}

// NewSecurityFindingsSearchRequestDataAttributesWithDefaults instantiates a new SecurityFindingsSearchRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityFindingsSearchRequestDataAttributesWithDefaults() *SecurityFindingsSearchRequestDataAttributes {
	this := SecurityFindingsSearchRequestDataAttributes{}
	var filter string = "*"
	this.Filter = &filter
	var sort SecurityFindingsSort = SECURITYFINDINGSSORT_DETECTION_CHANGED_AT_DESC
	this.Sort = &sort
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SecurityFindingsSearchRequestDataAttributes) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsSearchRequestDataAttributes) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SecurityFindingsSearchRequestDataAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *SecurityFindingsSearchRequestDataAttributes) SetFilter(v string) {
	o.Filter = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SecurityFindingsSearchRequestDataAttributes) GetPage() SecurityFindingsSearchRequestPage {
	if o == nil || o.Page == nil {
		var ret SecurityFindingsSearchRequestPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsSearchRequestDataAttributes) GetPageOk() (*SecurityFindingsSearchRequestPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SecurityFindingsSearchRequestDataAttributes) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given SecurityFindingsSearchRequestPage and assigns it to the Page field.
func (o *SecurityFindingsSearchRequestDataAttributes) SetPage(v SecurityFindingsSearchRequestPage) {
	o.Page = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SecurityFindingsSearchRequestDataAttributes) GetSort() SecurityFindingsSort {
	if o == nil || o.Sort == nil {
		var ret SecurityFindingsSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsSearchRequestDataAttributes) GetSortOk() (*SecurityFindingsSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SecurityFindingsSearchRequestDataAttributes) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given SecurityFindingsSort and assigns it to the Sort field.
func (o *SecurityFindingsSearchRequestDataAttributes) SetSort(v SecurityFindingsSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityFindingsSearchRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
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
func (o *SecurityFindingsSearchRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter *string                            `json:"filter,omitempty"`
		Page   *SecurityFindingsSearchRequestPage `json:"page,omitempty"`
		Sort   *SecurityFindingsSort              `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "page", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Filter = all.Filter
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	if all.Sort != nil && !all.Sort.IsValid() {
		hasInvalidField = true
	} else {
		o.Sort = all.Sort
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
