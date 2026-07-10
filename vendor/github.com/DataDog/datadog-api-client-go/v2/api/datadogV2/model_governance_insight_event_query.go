// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightEventQuery An event query used to compute an insight value.
type GovernanceInsightEventQuery struct {
	// The aggregation applied to an event query.
	Compute *GovernanceInsightEventCompute `json:"compute,omitempty"`
	// The event indexes the query runs against.
	Indexes []string `json:"indexes"`
	// The event search query string.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightEventQuery instantiates a new GovernanceInsightEventQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightEventQuery(indexes []string, query string) *GovernanceInsightEventQuery {
	this := GovernanceInsightEventQuery{}
	this.Indexes = indexes
	this.Query = query
	return &this
}

// NewGovernanceInsightEventQueryWithDefaults instantiates a new GovernanceInsightEventQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightEventQueryWithDefaults() *GovernanceInsightEventQuery {
	this := GovernanceInsightEventQuery{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *GovernanceInsightEventQuery) GetCompute() GovernanceInsightEventCompute {
	if o == nil || o.Compute == nil {
		var ret GovernanceInsightEventCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightEventQuery) GetComputeOk() (*GovernanceInsightEventCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *GovernanceInsightEventQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given GovernanceInsightEventCompute and assigns it to the Compute field.
func (o *GovernanceInsightEventQuery) SetCompute(v GovernanceInsightEventCompute) {
	o.Compute = &v
}

// GetIndexes returns the Indexes field value.
func (o *GovernanceInsightEventQuery) GetIndexes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightEventQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// SetIndexes sets field value.
func (o *GovernanceInsightEventQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQuery returns the Query field value.
func (o *GovernanceInsightEventQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightEventQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceInsightEventQuery) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightEventQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["indexes"] = o.Indexes
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightEventQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *GovernanceInsightEventCompute `json:"compute,omitempty"`
		Indexes *[]string                      `json:"indexes"`
		Query   *string                        `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Indexes == nil {
		return fmt.Errorf("required field indexes missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "indexes", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	o.Indexes = *all.Indexes
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
