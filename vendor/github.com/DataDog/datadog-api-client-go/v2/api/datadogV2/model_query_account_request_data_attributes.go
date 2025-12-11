// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryAccountRequestDataAttributes
type QueryAccountRequestDataAttributes struct {
	//
	Limit *int64 `json:"limit,omitempty"`
	//
	Query *string `json:"query,omitempty"`
	//
	SelectColumns []string `json:"select_columns,omitempty"`
	//
	Sort *QueryAccountRequestDataAttributesSort `json:"sort,omitempty"`
	//
	WildcardSearchTerm *string `json:"wildcard_search_term,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueryAccountRequestDataAttributes instantiates a new QueryAccountRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryAccountRequestDataAttributes() *QueryAccountRequestDataAttributes {
	this := QueryAccountRequestDataAttributes{}
	return &this
}

// NewQueryAccountRequestDataAttributesWithDefaults instantiates a new QueryAccountRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryAccountRequestDataAttributesWithDefaults() *QueryAccountRequestDataAttributes {
	this := QueryAccountRequestDataAttributes{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryAccountRequestDataAttributes) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAccountRequestDataAttributes) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryAccountRequestDataAttributes) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *QueryAccountRequestDataAttributes) SetLimit(v int64) {
	o.Limit = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *QueryAccountRequestDataAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAccountRequestDataAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *QueryAccountRequestDataAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *QueryAccountRequestDataAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetSelectColumns returns the SelectColumns field value if set, zero value otherwise.
func (o *QueryAccountRequestDataAttributes) GetSelectColumns() []string {
	if o == nil || o.SelectColumns == nil {
		var ret []string
		return ret
	}
	return o.SelectColumns
}

// GetSelectColumnsOk returns a tuple with the SelectColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAccountRequestDataAttributes) GetSelectColumnsOk() (*[]string, bool) {
	if o == nil || o.SelectColumns == nil {
		return nil, false
	}
	return &o.SelectColumns, true
}

// HasSelectColumns returns a boolean if a field has been set.
func (o *QueryAccountRequestDataAttributes) HasSelectColumns() bool {
	return o != nil && o.SelectColumns != nil
}

// SetSelectColumns gets a reference to the given []string and assigns it to the SelectColumns field.
func (o *QueryAccountRequestDataAttributes) SetSelectColumns(v []string) {
	o.SelectColumns = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *QueryAccountRequestDataAttributes) GetSort() QueryAccountRequestDataAttributesSort {
	if o == nil || o.Sort == nil {
		var ret QueryAccountRequestDataAttributesSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAccountRequestDataAttributes) GetSortOk() (*QueryAccountRequestDataAttributesSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *QueryAccountRequestDataAttributes) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given QueryAccountRequestDataAttributesSort and assigns it to the Sort field.
func (o *QueryAccountRequestDataAttributes) SetSort(v QueryAccountRequestDataAttributesSort) {
	o.Sort = &v
}

// GetWildcardSearchTerm returns the WildcardSearchTerm field value if set, zero value otherwise.
func (o *QueryAccountRequestDataAttributes) GetWildcardSearchTerm() string {
	if o == nil || o.WildcardSearchTerm == nil {
		var ret string
		return ret
	}
	return *o.WildcardSearchTerm
}

// GetWildcardSearchTermOk returns a tuple with the WildcardSearchTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAccountRequestDataAttributes) GetWildcardSearchTermOk() (*string, bool) {
	if o == nil || o.WildcardSearchTerm == nil {
		return nil, false
	}
	return o.WildcardSearchTerm, true
}

// HasWildcardSearchTerm returns a boolean if a field has been set.
func (o *QueryAccountRequestDataAttributes) HasWildcardSearchTerm() bool {
	return o != nil && o.WildcardSearchTerm != nil
}

// SetWildcardSearchTerm gets a reference to the given string and assigns it to the WildcardSearchTerm field.
func (o *QueryAccountRequestDataAttributes) SetWildcardSearchTerm(v string) {
	o.WildcardSearchTerm = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryAccountRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SelectColumns != nil {
		toSerialize["select_columns"] = o.SelectColumns
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.WildcardSearchTerm != nil {
		toSerialize["wildcard_search_term"] = o.WildcardSearchTerm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryAccountRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit              *int64                                 `json:"limit,omitempty"`
		Query              *string                                `json:"query,omitempty"`
		SelectColumns      []string                               `json:"select_columns,omitempty"`
		Sort               *QueryAccountRequestDataAttributesSort `json:"sort,omitempty"`
		WildcardSearchTerm *string                                `json:"wildcard_search_term,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit", "query", "select_columns", "sort", "wildcard_search_term"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Limit = all.Limit
	o.Query = all.Query
	o.SelectColumns = all.SelectColumns
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.WildcardSearchTerm = all.WildcardSearchTerm

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
