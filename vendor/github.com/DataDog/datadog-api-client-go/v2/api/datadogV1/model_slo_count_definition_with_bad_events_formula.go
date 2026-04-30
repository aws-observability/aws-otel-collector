// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCountDefinitionWithBadEventsFormula SLO count definition using a bad events formula alongside a good events formula.
type SLOCountDefinitionWithBadEventsFormula struct {
	// A formula that specifies how to combine the results of multiple queries.
	BadEventsFormula SLOFormula `json:"bad_events_formula"`
	// A formula that specifies how to combine the results of multiple queries.
	GoodEventsFormula SLOFormula `json:"good_events_formula"`
	//
	Queries []SLODataSourceQueryDefinition `json:"queries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSLOCountDefinitionWithBadEventsFormula instantiates a new SLOCountDefinitionWithBadEventsFormula object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCountDefinitionWithBadEventsFormula(badEventsFormula SLOFormula, goodEventsFormula SLOFormula, queries []SLODataSourceQueryDefinition) *SLOCountDefinitionWithBadEventsFormula {
	this := SLOCountDefinitionWithBadEventsFormula{}
	this.BadEventsFormula = badEventsFormula
	this.GoodEventsFormula = goodEventsFormula
	this.Queries = queries
	return &this
}

// NewSLOCountDefinitionWithBadEventsFormulaWithDefaults instantiates a new SLOCountDefinitionWithBadEventsFormula object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCountDefinitionWithBadEventsFormulaWithDefaults() *SLOCountDefinitionWithBadEventsFormula {
	this := SLOCountDefinitionWithBadEventsFormula{}
	return &this
}

// GetBadEventsFormula returns the BadEventsFormula field value.
func (o *SLOCountDefinitionWithBadEventsFormula) GetBadEventsFormula() SLOFormula {
	if o == nil {
		var ret SLOFormula
		return ret
	}
	return o.BadEventsFormula
}

// GetBadEventsFormulaOk returns a tuple with the BadEventsFormula field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinitionWithBadEventsFormula) GetBadEventsFormulaOk() (*SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BadEventsFormula, true
}

// SetBadEventsFormula sets field value.
func (o *SLOCountDefinitionWithBadEventsFormula) SetBadEventsFormula(v SLOFormula) {
	o.BadEventsFormula = v
}

// GetGoodEventsFormula returns the GoodEventsFormula field value.
func (o *SLOCountDefinitionWithBadEventsFormula) GetGoodEventsFormula() SLOFormula {
	if o == nil {
		var ret SLOFormula
		return ret
	}
	return o.GoodEventsFormula
}

// GetGoodEventsFormulaOk returns a tuple with the GoodEventsFormula field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinitionWithBadEventsFormula) GetGoodEventsFormulaOk() (*SLOFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoodEventsFormula, true
}

// SetGoodEventsFormula sets field value.
func (o *SLOCountDefinitionWithBadEventsFormula) SetGoodEventsFormula(v SLOFormula) {
	o.GoodEventsFormula = v
}

// GetQueries returns the Queries field value.
func (o *SLOCountDefinitionWithBadEventsFormula) GetQueries() []SLODataSourceQueryDefinition {
	if o == nil {
		var ret []SLODataSourceQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SLOCountDefinitionWithBadEventsFormula) GetQueriesOk() (*[]SLODataSourceQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *SLOCountDefinitionWithBadEventsFormula) SetQueries(v []SLODataSourceQueryDefinition) {
	o.Queries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCountDefinitionWithBadEventsFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bad_events_formula"] = o.BadEventsFormula
	toSerialize["good_events_formula"] = o.GoodEventsFormula
	toSerialize["queries"] = o.Queries
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOCountDefinitionWithBadEventsFormula) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BadEventsFormula  *SLOFormula                     `json:"bad_events_formula"`
		GoodEventsFormula *SLOFormula                     `json:"good_events_formula"`
		Queries           *[]SLODataSourceQueryDefinition `json:"queries"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BadEventsFormula == nil {
		return fmt.Errorf("required field bad_events_formula missing")
	}
	if all.GoodEventsFormula == nil {
		return fmt.Errorf("required field good_events_formula missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}

	hasInvalidField := false
	if all.BadEventsFormula.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BadEventsFormula = *all.BadEventsFormula
	if all.GoodEventsFormula.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GoodEventsFormula = *all.GoodEventsFormula
	o.Queries = *all.Queries

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
