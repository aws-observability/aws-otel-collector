// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryEventFilteredUsersRequestDataAttributes
type QueryEventFilteredUsersRequestDataAttributes struct {
	//
	EventQuery *QueryEventFilteredUsersRequestDataAttributesEventQuery `json:"event_query,omitempty"`
	//
	IncludeRowCount *bool `json:"include_row_count,omitempty"`
	//
	Limit *int64 `json:"limit,omitempty"`
	//
	Query *string `json:"query,omitempty"`
	//
	SelectColumns []string `json:"select_columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueryEventFilteredUsersRequestDataAttributes instantiates a new QueryEventFilteredUsersRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryEventFilteredUsersRequestDataAttributes() *QueryEventFilteredUsersRequestDataAttributes {
	this := QueryEventFilteredUsersRequestDataAttributes{}
	return &this
}

// NewQueryEventFilteredUsersRequestDataAttributesWithDefaults instantiates a new QueryEventFilteredUsersRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryEventFilteredUsersRequestDataAttributesWithDefaults() *QueryEventFilteredUsersRequestDataAttributes {
	this := QueryEventFilteredUsersRequestDataAttributes{}
	return &this
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetEventQuery() QueryEventFilteredUsersRequestDataAttributesEventQuery {
	if o == nil || o.EventQuery == nil {
		var ret QueryEventFilteredUsersRequestDataAttributesEventQuery
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetEventQueryOk() (*QueryEventFilteredUsersRequestDataAttributesEventQuery, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) HasEventQuery() bool {
	return o != nil && o.EventQuery != nil
}

// SetEventQuery gets a reference to the given QueryEventFilteredUsersRequestDataAttributesEventQuery and assigns it to the EventQuery field.
func (o *QueryEventFilteredUsersRequestDataAttributes) SetEventQuery(v QueryEventFilteredUsersRequestDataAttributesEventQuery) {
	o.EventQuery = &v
}

// GetIncludeRowCount returns the IncludeRowCount field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetIncludeRowCount() bool {
	if o == nil || o.IncludeRowCount == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRowCount
}

// GetIncludeRowCountOk returns a tuple with the IncludeRowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetIncludeRowCountOk() (*bool, bool) {
	if o == nil || o.IncludeRowCount == nil {
		return nil, false
	}
	return o.IncludeRowCount, true
}

// HasIncludeRowCount returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) HasIncludeRowCount() bool {
	return o != nil && o.IncludeRowCount != nil
}

// SetIncludeRowCount gets a reference to the given bool and assigns it to the IncludeRowCount field.
func (o *QueryEventFilteredUsersRequestDataAttributes) SetIncludeRowCount(v bool) {
	o.IncludeRowCount = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *QueryEventFilteredUsersRequestDataAttributes) SetLimit(v int64) {
	o.Limit = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *QueryEventFilteredUsersRequestDataAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetSelectColumns returns the SelectColumns field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetSelectColumns() []string {
	if o == nil || o.SelectColumns == nil {
		var ret []string
		return ret
	}
	return o.SelectColumns
}

// GetSelectColumnsOk returns a tuple with the SelectColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) GetSelectColumnsOk() (*[]string, bool) {
	if o == nil || o.SelectColumns == nil {
		return nil, false
	}
	return &o.SelectColumns, true
}

// HasSelectColumns returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributes) HasSelectColumns() bool {
	return o != nil && o.SelectColumns != nil
}

// SetSelectColumns gets a reference to the given []string and assigns it to the SelectColumns field.
func (o *QueryEventFilteredUsersRequestDataAttributes) SetSelectColumns(v []string) {
	o.SelectColumns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryEventFilteredUsersRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.IncludeRowCount != nil {
		toSerialize["include_row_count"] = o.IncludeRowCount
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryEventFilteredUsersRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventQuery      *QueryEventFilteredUsersRequestDataAttributesEventQuery `json:"event_query,omitempty"`
		IncludeRowCount *bool                                                   `json:"include_row_count,omitempty"`
		Limit           *int64                                                  `json:"limit,omitempty"`
		Query           *string                                                 `json:"query,omitempty"`
		SelectColumns   []string                                                `json:"select_columns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_query", "include_row_count", "limit", "query", "select_columns"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EventQuery != nil && all.EventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EventQuery = all.EventQuery
	o.IncludeRowCount = all.IncludeRowCount
	o.Limit = all.Limit
	o.Query = all.Query
	o.SelectColumns = all.SelectColumns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
