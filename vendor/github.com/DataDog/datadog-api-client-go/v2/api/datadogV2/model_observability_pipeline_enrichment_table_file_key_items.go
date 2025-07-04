// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFileKeyItems Defines how to map log fields to enrichment table columns during lookups.
type ObservabilityPipelineEnrichmentTableFileKeyItems struct {
	// The `items` `column`.
	Column string `json:"column"`
	// Defines how to compare key fields for enrichment table lookups.
	Comparison ObservabilityPipelineEnrichmentTableFileKeyItemsComparison `json:"comparison"`
	// The `items` `field`.
	Field string `json:"field"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableFileKeyItems instantiates a new ObservabilityPipelineEnrichmentTableFileKeyItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableFileKeyItems(column string, comparison ObservabilityPipelineEnrichmentTableFileKeyItemsComparison, field string) *ObservabilityPipelineEnrichmentTableFileKeyItems {
	this := ObservabilityPipelineEnrichmentTableFileKeyItems{}
	this.Column = column
	this.Comparison = comparison
	this.Field = field
	return &this
}

// NewObservabilityPipelineEnrichmentTableFileKeyItemsWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableFileKeyItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableFileKeyItemsWithDefaults() *ObservabilityPipelineEnrichmentTableFileKeyItems {
	this := ObservabilityPipelineEnrichmentTableFileKeyItems{}
	return &this
}

// GetColumn returns the Column field value.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) SetColumn(v string) {
	o.Column = v
}

// GetComparison returns the Comparison field value.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) GetComparison() ObservabilityPipelineEnrichmentTableFileKeyItemsComparison {
	if o == nil {
		var ret ObservabilityPipelineEnrichmentTableFileKeyItemsComparison
		return ret
	}
	return o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) GetComparisonOk() (*ObservabilityPipelineEnrichmentTableFileKeyItemsComparison, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparison, true
}

// SetComparison sets field value.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) SetComparison(v ObservabilityPipelineEnrichmentTableFileKeyItemsComparison) {
	o.Comparison = v
}

// GetField returns the Field field value.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) SetField(v string) {
	o.Field = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableFileKeyItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["comparison"] = o.Comparison
	toSerialize["field"] = o.Field

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableFileKeyItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column     *string                                                     `json:"column"`
		Comparison *ObservabilityPipelineEnrichmentTableFileKeyItemsComparison `json:"comparison"`
		Field      *string                                                     `json:"field"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Comparison == nil {
		return fmt.Errorf("required field comparison missing")
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column", "comparison", "field"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Column = *all.Column
	if !all.Comparison.IsValid() {
		hasInvalidField = true
	} else {
		o.Comparison = *all.Comparison
	}
	o.Field = *all.Field

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
