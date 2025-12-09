// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SingleAggregatedDnsResponseDataAttributesMetricsItems Metrics associated with an aggregated DNS flow.
type SingleAggregatedDnsResponseDataAttributesMetricsItems struct {
	// The metric key for DNS metrics.
	Key *DnsMetricKey `json:"key,omitempty"`
	// The metric value.
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSingleAggregatedDnsResponseDataAttributesMetricsItems instantiates a new SingleAggregatedDnsResponseDataAttributesMetricsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSingleAggregatedDnsResponseDataAttributesMetricsItems() *SingleAggregatedDnsResponseDataAttributesMetricsItems {
	this := SingleAggregatedDnsResponseDataAttributesMetricsItems{}
	return &this
}

// NewSingleAggregatedDnsResponseDataAttributesMetricsItemsWithDefaults instantiates a new SingleAggregatedDnsResponseDataAttributesMetricsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSingleAggregatedDnsResponseDataAttributesMetricsItemsWithDefaults() *SingleAggregatedDnsResponseDataAttributesMetricsItems {
	this := SingleAggregatedDnsResponseDataAttributesMetricsItems{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) GetKey() DnsMetricKey {
	if o == nil || o.Key == nil {
		var ret DnsMetricKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) GetKeyOk() (*DnsMetricKey, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given DnsMetricKey and assigns it to the Key field.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) SetKey(v DnsMetricKey) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SingleAggregatedDnsResponseDataAttributesMetricsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SingleAggregatedDnsResponseDataAttributesMetricsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *DnsMetricKey `json:"key,omitempty"`
		Value *int64        `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Key != nil && !all.Key.IsValid() {
		hasInvalidField = true
	} else {
		o.Key = all.Key
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
