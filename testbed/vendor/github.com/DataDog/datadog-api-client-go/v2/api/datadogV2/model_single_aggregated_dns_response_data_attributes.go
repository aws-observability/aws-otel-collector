// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SingleAggregatedDnsResponseDataAttributes Attributes for an aggregated DNS flow.
type SingleAggregatedDnsResponseDataAttributes struct {
	// The key, value pairs for each group by.
	GroupBys []SingleAggregatedDnsResponseDataAttributesGroupByItems `json:"group_bys,omitempty"`
	// Metrics associated with an aggregated DNS flow.
	Metrics []SingleAggregatedDnsResponseDataAttributesMetricsItems `json:"metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSingleAggregatedDnsResponseDataAttributes instantiates a new SingleAggregatedDnsResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSingleAggregatedDnsResponseDataAttributes() *SingleAggregatedDnsResponseDataAttributes {
	this := SingleAggregatedDnsResponseDataAttributes{}
	return &this
}

// NewSingleAggregatedDnsResponseDataAttributesWithDefaults instantiates a new SingleAggregatedDnsResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSingleAggregatedDnsResponseDataAttributesWithDefaults() *SingleAggregatedDnsResponseDataAttributes {
	this := SingleAggregatedDnsResponseDataAttributes{}
	return &this
}

// GetGroupBys returns the GroupBys field value if set, zero value otherwise.
func (o *SingleAggregatedDnsResponseDataAttributes) GetGroupBys() []SingleAggregatedDnsResponseDataAttributesGroupByItems {
	if o == nil || o.GroupBys == nil {
		var ret []SingleAggregatedDnsResponseDataAttributesGroupByItems
		return ret
	}
	return o.GroupBys
}

// GetGroupBysOk returns a tuple with the GroupBys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedDnsResponseDataAttributes) GetGroupBysOk() (*[]SingleAggregatedDnsResponseDataAttributesGroupByItems, bool) {
	if o == nil || o.GroupBys == nil {
		return nil, false
	}
	return &o.GroupBys, true
}

// HasGroupBys returns a boolean if a field has been set.
func (o *SingleAggregatedDnsResponseDataAttributes) HasGroupBys() bool {
	return o != nil && o.GroupBys != nil
}

// SetGroupBys gets a reference to the given []SingleAggregatedDnsResponseDataAttributesGroupByItems and assigns it to the GroupBys field.
func (o *SingleAggregatedDnsResponseDataAttributes) SetGroupBys(v []SingleAggregatedDnsResponseDataAttributesGroupByItems) {
	o.GroupBys = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SingleAggregatedDnsResponseDataAttributes) GetMetrics() []SingleAggregatedDnsResponseDataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []SingleAggregatedDnsResponseDataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedDnsResponseDataAttributes) GetMetricsOk() (*[]SingleAggregatedDnsResponseDataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SingleAggregatedDnsResponseDataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []SingleAggregatedDnsResponseDataAttributesMetricsItems and assigns it to the Metrics field.
func (o *SingleAggregatedDnsResponseDataAttributes) SetMetrics(v []SingleAggregatedDnsResponseDataAttributesMetricsItems) {
	o.Metrics = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SingleAggregatedDnsResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GroupBys != nil {
		toSerialize["group_bys"] = o.GroupBys
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SingleAggregatedDnsResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBys []SingleAggregatedDnsResponseDataAttributesGroupByItems `json:"group_bys,omitempty"`
		Metrics  []SingleAggregatedDnsResponseDataAttributesMetricsItems `json:"metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_bys", "metrics"})
	} else {
		return err
	}
	o.GroupBys = all.GroupBys
	o.Metrics = all.Metrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
