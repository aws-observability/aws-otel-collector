// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchResponsePropertyFieldFacetData Facet data for the incident property fields.
type IncidentSearchResponsePropertyFieldFacetData struct {
	// Aggregate information for numeric incident data.
	Aggregates *IncidentSearchResponseNumericFacetDataAggregates `json:"aggregates,omitempty"`
	// Facet data for the property field of an incident.
	Facets []IncidentSearchResponseFieldFacetData `json:"facets"`
	// Name of the incident property field.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentSearchResponsePropertyFieldFacetData instantiates a new IncidentSearchResponsePropertyFieldFacetData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSearchResponsePropertyFieldFacetData(facets []IncidentSearchResponseFieldFacetData, name string) *IncidentSearchResponsePropertyFieldFacetData {
	this := IncidentSearchResponsePropertyFieldFacetData{}
	this.Facets = facets
	this.Name = name
	return &this
}

// NewIncidentSearchResponsePropertyFieldFacetDataWithDefaults instantiates a new IncidentSearchResponsePropertyFieldFacetData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSearchResponsePropertyFieldFacetDataWithDefaults() *IncidentSearchResponsePropertyFieldFacetData {
	this := IncidentSearchResponsePropertyFieldFacetData{}
	return &this
}

// GetAggregates returns the Aggregates field value if set, zero value otherwise.
func (o *IncidentSearchResponsePropertyFieldFacetData) GetAggregates() IncidentSearchResponseNumericFacetDataAggregates {
	if o == nil || o.Aggregates == nil {
		var ret IncidentSearchResponseNumericFacetDataAggregates
		return ret
	}
	return *o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponsePropertyFieldFacetData) GetAggregatesOk() (*IncidentSearchResponseNumericFacetDataAggregates, bool) {
	if o == nil || o.Aggregates == nil {
		return nil, false
	}
	return o.Aggregates, true
}

// HasAggregates returns a boolean if a field has been set.
func (o *IncidentSearchResponsePropertyFieldFacetData) HasAggregates() bool {
	return o != nil && o.Aggregates != nil
}

// SetAggregates gets a reference to the given IncidentSearchResponseNumericFacetDataAggregates and assigns it to the Aggregates field.
func (o *IncidentSearchResponsePropertyFieldFacetData) SetAggregates(v IncidentSearchResponseNumericFacetDataAggregates) {
	o.Aggregates = &v
}

// GetFacets returns the Facets field value.
func (o *IncidentSearchResponsePropertyFieldFacetData) GetFacets() []IncidentSearchResponseFieldFacetData {
	if o == nil {
		var ret []IncidentSearchResponseFieldFacetData
		return ret
	}
	return o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponsePropertyFieldFacetData) GetFacetsOk() (*[]IncidentSearchResponseFieldFacetData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facets, true
}

// SetFacets sets field value.
func (o *IncidentSearchResponsePropertyFieldFacetData) SetFacets(v []IncidentSearchResponseFieldFacetData) {
	o.Facets = v
}

// GetName returns the Name field value.
func (o *IncidentSearchResponsePropertyFieldFacetData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponsePropertyFieldFacetData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentSearchResponsePropertyFieldFacetData) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSearchResponsePropertyFieldFacetData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregates != nil {
		toSerialize["aggregates"] = o.Aggregates
	}
	toSerialize["facets"] = o.Facets
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSearchResponsePropertyFieldFacetData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregates *IncidentSearchResponseNumericFacetDataAggregates `json:"aggregates,omitempty"`
		Facets     *[]IncidentSearchResponseFieldFacetData           `json:"facets"`
		Name       *string                                           `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facets == nil {
		return fmt.Errorf("required field facets missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregates", "facets", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregates != nil && all.Aggregates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Aggregates = all.Aggregates
	o.Facets = *all.Facets
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
