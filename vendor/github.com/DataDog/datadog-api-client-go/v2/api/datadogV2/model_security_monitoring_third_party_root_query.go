// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringThirdPartyRootQuery A query to be combined with the third party case query.
type SecurityMonitoringThirdPartyRootQuery struct {
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// Query to run on logs.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringThirdPartyRootQuery instantiates a new SecurityMonitoringThirdPartyRootQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringThirdPartyRootQuery() *SecurityMonitoringThirdPartyRootQuery {
	this := SecurityMonitoringThirdPartyRootQuery{}
	return &this
}

// NewSecurityMonitoringThirdPartyRootQueryWithDefaults instantiates a new SecurityMonitoringThirdPartyRootQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringThirdPartyRootQueryWithDefaults() *SecurityMonitoringThirdPartyRootQuery {
	this := SecurityMonitoringThirdPartyRootQuery{}
	return &this
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringThirdPartyRootQuery) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringThirdPartyRootQuery) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringThirdPartyRootQuery) HasGroupByFields() bool {
	return o != nil && o.GroupByFields != nil
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *SecurityMonitoringThirdPartyRootQuery) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringThirdPartyRootQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringThirdPartyRootQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringThirdPartyRootQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringThirdPartyRootQuery) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringThirdPartyRootQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GroupByFields != nil {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringThirdPartyRootQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupByFields []string `json:"groupByFields,omitempty"`
		Query         *string  `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"groupByFields", "query"})
	} else {
		return err
	}
	o.GroupByFields = all.GroupByFields
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
