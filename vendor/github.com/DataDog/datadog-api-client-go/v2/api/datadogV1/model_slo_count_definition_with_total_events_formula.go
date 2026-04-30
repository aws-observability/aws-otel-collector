// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCountDefinitionWithTotalEventsFormula SLO count definition using a total events formula alongside a good events formula.
type SLOCountDefinitionWithTotalEventsFormula struct {
	// A formula that specifies how to combine the results of multiple queries.
	GoodEventsFormula SLOFormula `json:"good_events_formula"`
	//
	Queries []SLODataSourceQueryDefinition `json:"queries"`
	// A formula that specifies how to combine the results of multiple queries.
	TotalEventsFormula SLOFormula `json:"total_events_formula"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSLOCountDefinitionWithTotalEventsFormula instantiates a new SLOCountDefinitionWithTotalEventsFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCountDefinitionWithTotalEventsFormula(goodEventsFormula SLOFormula, queries []SLODataSourceQueryDefinition, totalEventsFormula SLOFormula) *SLOCountDefinitionWithTotalEventsFormula {
	this := SLOCountDefinitionWithTotalEventsFormula{}
	this.GoodEventsFormula = goodEventsFormula
	this.Queries = queries
	this.TotalEventsFormula = totalEventsFormula
	return &this
}

// NewSLOCountDefinitionWithTotalEventsFormulaWithDefaults instantiates a new SLOCountDefinitionWithTotalEventsFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCountDefinitionWithTotalEventsFormulaWithDefaults() *SLOCountDefinitionWithTotalEventsFormula {
	this := SLOCountDefinitionWithTotalEventsFormula{}
	return &this
}

// GetGoodEventsFormula returns the GoodEventsFormula field value.
func (o *SLOCountDefinitionWithTotalEventsFormula) GetGoodEventsFormula() SLOFormula {
	if o == nil {
		var ret SLOFormula
		return ret
	}
	return o.GoodEventsFormula
}

// GetGoodEventsFormulaOk returns a tuple with the GoodEventsFormula field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinitionWithTotalEventsFormula) GetGoodEventsFormulaOk() (*SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoodEventsFormula, true
}

// SetGoodEventsFormula sets field value.
func (o *SLOCountDefinitionWithTotalEventsFormula) SetGoodEventsFormula(v SLOFormula) {
	o.GoodEventsFormula = v
}

// GetQueries returns the Queries field value.
func (o *SLOCountDefinitionWithTotalEventsFormula) GetQueries() []SLODataSourceQueryDefinition {
	if o == nil {
		var ret []SLODataSourceQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinitionWithTotalEventsFormula) GetQueriesOk() (*[]SLODataSourceQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *SLOCountDefinitionWithTotalEventsFormula) SetQueries(v []SLODataSourceQueryDefinition) {
	o.Queries = v
}

// GetTotalEventsFormula returns the TotalEventsFormula field value.
func (o *SLOCountDefinitionWithTotalEventsFormula) GetTotalEventsFormula() SLOFormula {
	if o == nil {
		var ret SLOFormula
		return ret
	}
	return o.TotalEventsFormula
}

// GetTotalEventsFormulaOk returns a tuple with the TotalEventsFormula field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinitionWithTotalEventsFormula) GetTotalEventsFormulaOk() (*SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalEventsFormula, true
}

// SetTotalEventsFormula sets field value.
func (o *SLOCountDefinitionWithTotalEventsFormula) SetTotalEventsFormula(v SLOFormula) {
	o.TotalEventsFormula = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCountDefinitionWithTotalEventsFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["good_events_formula"] = o.GoodEventsFormula
	toSerialize["queries"] = o.Queries
	toSerialize["total_events_formula"] = o.TotalEventsFormula
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOCountDefinitionWithTotalEventsFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GoodEventsFormula  *SLOFormula                     `json:"good_events_formula"`
		Queries            *[]SLODataSourceQueryDefinition `json:"queries"`
		TotalEventsFormula *SLOFormula                     `json:"total_events_formula"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GoodEventsFormula == nil {
		return fmt.Errorf("required field good_events_formula missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.TotalEventsFormula == nil {
		return fmt.Errorf("required field total_events_formula missing")
	}

	hasInvalidField := false
	if all.GoodEventsFormula.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GoodEventsFormula = *all.GoodEventsFormula
	o.Queries = *all.Queries
	if all.TotalEventsFormula.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalEventsFormula = *all.TotalEventsFormula

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
