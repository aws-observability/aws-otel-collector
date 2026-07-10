// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightUsageQuery A usage query used to compute an insight value.
type GovernanceInsightUsageQuery struct {
	// The usage query string.
	Query string `json:"query"`
	// How the query result series is reduced to a single value.
	Reducer string `json:"reducer"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightUsageQuery instantiates a new GovernanceInsightUsageQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightUsageQuery(query string, reducer string) *GovernanceInsightUsageQuery {
	this := GovernanceInsightUsageQuery{}
	this.Query = query
	this.Reducer = reducer
	return &this
}

// NewGovernanceInsightUsageQueryWithDefaults instantiates a new GovernanceInsightUsageQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightUsageQueryWithDefaults() *GovernanceInsightUsageQuery {
	this := GovernanceInsightUsageQuery{}
	return &this
}

// GetQuery returns the Query field value.
func (o *GovernanceInsightUsageQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightUsageQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceInsightUsageQuery) SetQuery(v string) {
	o.Query = v
}

// GetReducer returns the Reducer field value.
func (o *GovernanceInsightUsageQuery) GetReducer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reducer
}

// GetReducerOk returns a tuple with the Reducer field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightUsageQuery) GetReducerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reducer, true
}

// SetReducer sets field value.
func (o *GovernanceInsightUsageQuery) SetReducer(v string) {
	o.Reducer = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightUsageQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["reducer"] = o.Reducer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightUsageQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query   *string `json:"query"`
		Reducer *string `json:"reducer"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "reducer"})
	} else {
		return err
	}
	o.Query = *all.Query
	o.Reducer = *all.Reducer

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
