// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSerie A series in a timeseries response.
type ProductAnalyticsSerie struct {
	// The group-by tag values that identify this series.
	GroupTags []string `json:"group_tags,omitempty"`
	// The index of the query that produced this series.
	QueryIndex *int64 `json:"query_index,omitempty"`
	// Unit definitions for the series values.
	Unit []ProductAnalyticsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSerie instantiates a new ProductAnalyticsSerie object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSerie() *ProductAnalyticsSerie {
	this := ProductAnalyticsSerie{}
	return &this
}

// NewProductAnalyticsSerieWithDefaults instantiates a new ProductAnalyticsSerie object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSerieWithDefaults() *ProductAnalyticsSerie {
	this := ProductAnalyticsSerie{}
	return &this
}

// GetGroupTags returns the GroupTags field value if set, zero value otherwise.
func (o *ProductAnalyticsSerie) GetGroupTags() []string {
	if o == nil || o.GroupTags == nil {
		var ret []string
		return ret
	}
	return o.GroupTags
}

// GetGroupTagsOk returns a tuple with the GroupTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSerie) GetGroupTagsOk() (*[]string, bool) {
	if o == nil || o.GroupTags == nil {
		return nil, false
	}
	return &o.GroupTags, true
}

// HasGroupTags returns a boolean if a field has been set.
func (o *ProductAnalyticsSerie) HasGroupTags() bool {
	return o != nil && o.GroupTags != nil
}

// SetGroupTags gets a reference to the given []string and assigns it to the GroupTags field.
func (o *ProductAnalyticsSerie) SetGroupTags(v []string) {
	o.GroupTags = v
}

// GetQueryIndex returns the QueryIndex field value if set, zero value otherwise.
func (o *ProductAnalyticsSerie) GetQueryIndex() int64 {
	if o == nil || o.QueryIndex == nil {
		var ret int64
		return ret
	}
	return *o.QueryIndex
}

// GetQueryIndexOk returns a tuple with the QueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSerie) GetQueryIndexOk() (*int64, bool) {
	if o == nil || o.QueryIndex == nil {
		return nil, false
	}
	return o.QueryIndex, true
}

// HasQueryIndex returns a boolean if a field has been set.
func (o *ProductAnalyticsSerie) HasQueryIndex() bool {
	return o != nil && o.QueryIndex != nil
}

// SetQueryIndex gets a reference to the given int64 and assigns it to the QueryIndex field.
func (o *ProductAnalyticsSerie) SetQueryIndex(v int64) {
	o.QueryIndex = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProductAnalyticsSerie) GetUnit() []ProductAnalyticsUnit {
	if o == nil || o.Unit == nil {
		var ret []ProductAnalyticsUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSerie) GetUnitOk() (*[]ProductAnalyticsUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return &o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProductAnalyticsSerie) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given []ProductAnalyticsUnit and assigns it to the Unit field.
func (o *ProductAnalyticsSerie) SetUnit(v []ProductAnalyticsUnit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSerie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GroupTags != nil {
		toSerialize["group_tags"] = o.GroupTags
	}
	if o.QueryIndex != nil {
		toSerialize["query_index"] = o.QueryIndex
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsSerie) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupTags  []string               `json:"group_tags,omitempty"`
		QueryIndex *int64                 `json:"query_index,omitempty"`
		Unit       []ProductAnalyticsUnit `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_tags", "query_index", "unit"})
	} else {
		return err
	}
	o.GroupTags = all.GroupTags
	o.QueryIndex = all.QueryIndex
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
