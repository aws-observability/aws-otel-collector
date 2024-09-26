// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionMetadata The object containing document metadata.
type MonthlyUsageAttributionMetadata struct {
	// An array of available aggregates.
	Aggregates []UsageAttributionAggregatesBody `json:"aggregates,omitempty"`
	// The metadata for the current pagination.
	Pagination *MonthlyUsageAttributionPagination `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyUsageAttributionMetadata instantiates a new MonthlyUsageAttributionMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyUsageAttributionMetadata() *MonthlyUsageAttributionMetadata {
	this := MonthlyUsageAttributionMetadata{}
	return &this
}

// NewMonthlyUsageAttributionMetadataWithDefaults instantiates a new MonthlyUsageAttributionMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyUsageAttributionMetadataWithDefaults() *MonthlyUsageAttributionMetadata {
	this := MonthlyUsageAttributionMetadata{}
	return &this
}

// GetAggregates returns the Aggregates field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionMetadata) GetAggregates() []UsageAttributionAggregatesBody {
	if o == nil || o.Aggregates == nil {
		var ret []UsageAttributionAggregatesBody
		return ret
	}
	return o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionMetadata) GetAggregatesOk() (*[]UsageAttributionAggregatesBody, bool) {
	if o == nil || o.Aggregates == nil {
		return nil, false
	}
	return &o.Aggregates, true
}

// HasAggregates returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionMetadata) HasAggregates() bool {
	return o != nil && o.Aggregates != nil
}

// SetAggregates gets a reference to the given []UsageAttributionAggregatesBody and assigns it to the Aggregates field.
func (o *MonthlyUsageAttributionMetadata) SetAggregates(v []UsageAttributionAggregatesBody) {
	o.Aggregates = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionMetadata) GetPagination() MonthlyUsageAttributionPagination {
	if o == nil || o.Pagination == nil {
		var ret MonthlyUsageAttributionPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionMetadata) GetPaginationOk() (*MonthlyUsageAttributionPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionMetadata) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given MonthlyUsageAttributionPagination and assigns it to the Pagination field.
func (o *MonthlyUsageAttributionMetadata) SetPagination(v MonthlyUsageAttributionPagination) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyUsageAttributionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregates != nil {
		toSerialize["aggregates"] = o.Aggregates
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyUsageAttributionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregates []UsageAttributionAggregatesBody   `json:"aggregates,omitempty"`
		Pagination *MonthlyUsageAttributionPagination `json:"pagination,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregates", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Aggregates = all.Aggregates
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
