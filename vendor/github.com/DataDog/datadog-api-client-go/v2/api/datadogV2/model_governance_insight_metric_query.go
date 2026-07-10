// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightMetricQuery A metric query used to compute an insight value.
type GovernanceInsightMetricQuery struct {
	// The query string.
	Query string `json:"query"`
	// How the query result series is reduced to a single value.
	Reducer string `json:"reducer"`
	// The data source the query runs against.
	Source string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightMetricQuery instantiates a new GovernanceInsightMetricQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightMetricQuery(query string, reducer string, source string) *GovernanceInsightMetricQuery {
	this := GovernanceInsightMetricQuery{}
	this.Query = query
	this.Reducer = reducer
	this.Source = source
	return &this
}

// NewGovernanceInsightMetricQueryWithDefaults instantiates a new GovernanceInsightMetricQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightMetricQueryWithDefaults() *GovernanceInsightMetricQuery {
	this := GovernanceInsightMetricQuery{}
	return &this
}

// GetQuery returns the Query field value.
func (o *GovernanceInsightMetricQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightMetricQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceInsightMetricQuery) SetQuery(v string) {
	o.Query = v
}

// GetReducer returns the Reducer field value.
func (o *GovernanceInsightMetricQuery) GetReducer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reducer
}

// GetReducerOk returns a tuple with the Reducer field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightMetricQuery) GetReducerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reducer, true
}

// SetReducer sets field value.
func (o *GovernanceInsightMetricQuery) SetReducer(v string) {
	o.Reducer = v
}

// GetSource returns the Source field value.
func (o *GovernanceInsightMetricQuery) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightMetricQuery) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *GovernanceInsightMetricQuery) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightMetricQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["reducer"] = o.Reducer
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightMetricQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query   *string `json:"query"`
		Reducer *string `json:"reducer"`
		Source  *string `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Reducer == nil {
		return fmt.Errorf("required field reducer missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "reducer", "source"})
	} else {
		return err
	}
	o.Query = *all.Query
	o.Reducer = *all.Reducer
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
