// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightAuditQuery An audit log query used to compute an insight value.
type GovernanceInsightAuditQuery struct {
	// The aggregation applied to an audit log query.
	Compute GovernanceInsightAuditCompute `json:"compute"`
	// The audit log indexes the query runs against.
	Indexes []string `json:"indexes"`
	// The audit log search query string.
	Query string `json:"query"`
	// The data source the query runs against.
	Source string `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightAuditQuery instantiates a new GovernanceInsightAuditQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightAuditQuery(compute GovernanceInsightAuditCompute, indexes []string, query string, source string) *GovernanceInsightAuditQuery {
	this := GovernanceInsightAuditQuery{}
	this.Compute = compute
	this.Indexes = indexes
	this.Query = query
	this.Source = source
	return &this
}

// NewGovernanceInsightAuditQueryWithDefaults instantiates a new GovernanceInsightAuditQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightAuditQueryWithDefaults() *GovernanceInsightAuditQuery {
	this := GovernanceInsightAuditQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *GovernanceInsightAuditQuery) GetCompute() GovernanceInsightAuditCompute {
	if o == nil {
		var ret GovernanceInsightAuditCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditQuery) GetComputeOk() (*GovernanceInsightAuditCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *GovernanceInsightAuditQuery) SetCompute(v GovernanceInsightAuditCompute) {
	o.Compute = v
}

// GetIndexes returns the Indexes field value.
func (o *GovernanceInsightAuditQuery) GetIndexes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// SetIndexes sets field value.
func (o *GovernanceInsightAuditQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQuery returns the Query field value.
func (o *GovernanceInsightAuditQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceInsightAuditQuery) SetQuery(v string) {
	o.Query = v
}

// GetSource returns the Source field value.
func (o *GovernanceInsightAuditQuery) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAuditQuery) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *GovernanceInsightAuditQuery) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightAuditQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["indexes"] = o.Indexes
	toSerialize["query"] = o.Query
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightAuditQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *GovernanceInsightAuditCompute `json:"compute"`
		Indexes *[]string                      `json:"indexes"`
		Query   *string                        `json:"query"`
		Source  *string                        `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.Indexes == nil {
		return fmt.Errorf("required field indexes missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "indexes", "query", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.Indexes = *all.Indexes
	o.Query = *all.Query
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
