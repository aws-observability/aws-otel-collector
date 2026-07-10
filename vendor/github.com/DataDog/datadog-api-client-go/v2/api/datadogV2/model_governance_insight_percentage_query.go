// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightPercentageQuery A percentage query that computes an insight value as a ratio of two metric queries.
type GovernanceInsightPercentageQuery struct {
	// A metric query used to compute an insight value.
	DenominatorQuery GovernanceInsightMetricQuery `json:"denominator_query"`
	// A metric query used to compute an insight value.
	NumeratorQuery GovernanceInsightMetricQuery `json:"numerator_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightPercentageQuery instantiates a new GovernanceInsightPercentageQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightPercentageQuery(denominatorQuery GovernanceInsightMetricQuery, numeratorQuery GovernanceInsightMetricQuery) *GovernanceInsightPercentageQuery {
	this := GovernanceInsightPercentageQuery{}
	this.DenominatorQuery = denominatorQuery
	this.NumeratorQuery = numeratorQuery
	return &this
}

// NewGovernanceInsightPercentageQueryWithDefaults instantiates a new GovernanceInsightPercentageQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightPercentageQueryWithDefaults() *GovernanceInsightPercentageQuery {
	this := GovernanceInsightPercentageQuery{}
	return &this
}

// GetDenominatorQuery returns the DenominatorQuery field value.
func (o *GovernanceInsightPercentageQuery) GetDenominatorQuery() GovernanceInsightMetricQuery {
	if o == nil {
		var ret GovernanceInsightMetricQuery
		return ret
	}
	return o.DenominatorQuery
}

// GetDenominatorQueryOk returns a tuple with the DenominatorQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightPercentageQuery) GetDenominatorQueryOk() (*GovernanceInsightMetricQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenominatorQuery, true
}

// SetDenominatorQuery sets field value.
func (o *GovernanceInsightPercentageQuery) SetDenominatorQuery(v GovernanceInsightMetricQuery) {
	o.DenominatorQuery = v
}

// GetNumeratorQuery returns the NumeratorQuery field value.
func (o *GovernanceInsightPercentageQuery) GetNumeratorQuery() GovernanceInsightMetricQuery {
	if o == nil {
		var ret GovernanceInsightMetricQuery
		return ret
	}
	return o.NumeratorQuery
}

// GetNumeratorQueryOk returns a tuple with the NumeratorQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightPercentageQuery) GetNumeratorQueryOk() (*GovernanceInsightMetricQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeratorQuery, true
}

// SetNumeratorQuery sets field value.
func (o *GovernanceInsightPercentageQuery) SetNumeratorQuery(v GovernanceInsightMetricQuery) {
	o.NumeratorQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightPercentageQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["denominator_query"] = o.DenominatorQuery
	toSerialize["numerator_query"] = o.NumeratorQuery

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightPercentageQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DenominatorQuery *GovernanceInsightMetricQuery `json:"denominator_query"`
		NumeratorQuery   *GovernanceInsightMetricQuery `json:"numerator_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DenominatorQuery == nil {
		return fmt.Errorf("required field denominator_query missing")
	}
	if all.NumeratorQuery == nil {
		return fmt.Errorf("required field numerator_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"denominator_query", "numerator_query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DenominatorQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DenominatorQuery = *all.DenominatorQuery
	if all.NumeratorQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumeratorQuery = *all.NumeratorQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
