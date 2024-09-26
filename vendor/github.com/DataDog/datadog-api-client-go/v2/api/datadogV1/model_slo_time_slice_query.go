// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOTimeSliceQuery The queries and formula used to calculate the SLI value.
type SLOTimeSliceQuery struct {
	// A list that contains exactly one formula, as only a single formula may be used in a time-slice SLO.
	Formulas []SLOFormula `json:"formulas"`
	// A list of queries that are used to calculate the SLI value.
	Queries []SLODataSourceQueryDefinition `json:"queries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLOTimeSliceQuery instantiates a new SLOTimeSliceQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOTimeSliceQuery(formulas []SLOFormula, queries []SLODataSourceQueryDefinition) *SLOTimeSliceQuery {
	this := SLOTimeSliceQuery{}
	this.Formulas = formulas
	this.Queries = queries
	return &this
}

// NewSLOTimeSliceQueryWithDefaults instantiates a new SLOTimeSliceQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOTimeSliceQueryWithDefaults() *SLOTimeSliceQuery {
	this := SLOTimeSliceQuery{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *SLOTimeSliceQuery) GetFormulas() []SLOFormula {
	if o == nil {
		var ret []SLOFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceQuery) GetFormulasOk() (*[]SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *SLOTimeSliceQuery) SetFormulas(v []SLOFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *SLOTimeSliceQuery) GetQueries() []SLODataSourceQueryDefinition {
	if o == nil {
		var ret []SLODataSourceQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SLOTimeSliceQuery) GetQueriesOk() (*[]SLODataSourceQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *SLOTimeSliceQuery) SetQueries(v []SLODataSourceQueryDefinition) {
	o.Queries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOTimeSliceQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formulas"] = o.Formulas
	toSerialize["queries"] = o.Queries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOTimeSliceQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas *[]SLOFormula                   `json:"formulas"`
		Queries  *[]SLODataSourceQueryDefinition `json:"queries"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formulas == nil {
		return fmt.Errorf("required field formulas missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formulas", "queries"})
	} else {
		return err
	}
	o.Formulas = *all.Formulas
	o.Queries = *all.Queries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
