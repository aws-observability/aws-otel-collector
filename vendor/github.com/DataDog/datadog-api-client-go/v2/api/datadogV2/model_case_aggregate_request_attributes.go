// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAggregateRequestAttributes Attributes for the aggregation request, including the search query and grouping configuration.
type CaseAggregateRequestAttributes struct {
	// Configuration for grouping aggregated results by one or more case fields.
	GroupBy CaseAggregateGroupBy `json:"group_by"`
	// A search query to filter which cases are included in the aggregation. Uses the same syntax as the Case Management search bar.
	QueryFilter string `json:"query_filter"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseAggregateRequestAttributes instantiates a new CaseAggregateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseAggregateRequestAttributes(groupBy CaseAggregateGroupBy, queryFilter string) *CaseAggregateRequestAttributes {
	this := CaseAggregateRequestAttributes{}
	this.GroupBy = groupBy
	this.QueryFilter = queryFilter
	return &this
}

// NewCaseAggregateRequestAttributesWithDefaults instantiates a new CaseAggregateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseAggregateRequestAttributesWithDefaults() *CaseAggregateRequestAttributes {
	this := CaseAggregateRequestAttributes{}
	return &this
}

// GetGroupBy returns the GroupBy field value.
func (o *CaseAggregateRequestAttributes) GetGroupBy() CaseAggregateGroupBy {
	if o == nil {
		var ret CaseAggregateGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateRequestAttributes) GetGroupByOk() (*CaseAggregateGroupBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *CaseAggregateRequestAttributes) SetGroupBy(v CaseAggregateGroupBy) {
	o.GroupBy = v
}

// GetQueryFilter returns the QueryFilter field value.
func (o *CaseAggregateRequestAttributes) GetQueryFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value
// and a boolean to check if the value has been set.
func (o *CaseAggregateRequestAttributes) GetQueryFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryFilter, true
}

// SetQueryFilter sets field value.
func (o *CaseAggregateRequestAttributes) SetQueryFilter(v string) {
	o.QueryFilter = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseAggregateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_by"] = o.GroupBy
	toSerialize["query_filter"] = o.QueryFilter

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseAggregateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBy     *CaseAggregateGroupBy `json:"group_by"`
		QueryFilter *string               `json:"query_filter"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.QueryFilter == nil {
		return fmt.Errorf("required field query_filter missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_by", "query_filter"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GroupBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GroupBy = *all.GroupBy
	o.QueryFilter = *all.QueryFilter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
