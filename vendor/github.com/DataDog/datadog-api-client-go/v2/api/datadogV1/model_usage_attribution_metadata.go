// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageAttributionMetadata The object containing document metadata.
type UsageAttributionMetadata struct {
	// An array of available aggregates.
	Aggregates []UsageAttributionAggregatesBody `json:"aggregates,omitempty"`
	// The metadata for the current pagination.
	Pagination *UsageAttributionPagination `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageAttributionMetadata instantiates a new UsageAttributionMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageAttributionMetadata() *UsageAttributionMetadata {
	this := UsageAttributionMetadata{}
	return &this
}

// NewUsageAttributionMetadataWithDefaults instantiates a new UsageAttributionMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageAttributionMetadataWithDefaults() *UsageAttributionMetadata {
	this := UsageAttributionMetadata{}
	return &this
}

// GetAggregates returns the Aggregates field value if set, zero value otherwise.
func (o *UsageAttributionMetadata) GetAggregates() []UsageAttributionAggregatesBody {
	if o == nil || o.Aggregates == nil {
		var ret []UsageAttributionAggregatesBody
		return ret
	}
	return o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributionMetadata) GetAggregatesOk() (*[]UsageAttributionAggregatesBody, bool) {
	if o == nil || o.Aggregates == nil {
		return nil, false
	}
	return &o.Aggregates, true
}

// HasAggregates returns a boolean if a field has been set.
func (o *UsageAttributionMetadata) HasAggregates() bool {
	return o != nil && o.Aggregates != nil
}

// SetAggregates gets a reference to the given []UsageAttributionAggregatesBody and assigns it to the Aggregates field.
func (o *UsageAttributionMetadata) SetAggregates(v []UsageAttributionAggregatesBody) {
	o.Aggregates = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *UsageAttributionMetadata) GetPagination() UsageAttributionPagination {
	if o == nil || o.Pagination == nil {
		var ret UsageAttributionPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributionMetadata) GetPaginationOk() (*UsageAttributionPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *UsageAttributionMetadata) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given UsageAttributionPagination and assigns it to the Pagination field.
func (o *UsageAttributionMetadata) SetPagination(v UsageAttributionPagination) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageAttributionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
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
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageAttributionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregates []UsageAttributionAggregatesBody `json:"aggregates,omitempty"`
		Pagination *UsageAttributionPagination      `json:"pagination,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}