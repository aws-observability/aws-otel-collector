// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionsResponseMeta Metadata returned alongside a list of sample log generation subscriptions.
type SampleLogGenerationSubscriptionsResponseMeta struct {
	// The total number of subscriptions matching the request, irrespective of pagination.
	TotalSubscriptions int32 `json:"total_subscriptions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationSubscriptionsResponseMeta instantiates a new SampleLogGenerationSubscriptionsResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationSubscriptionsResponseMeta(totalSubscriptions int32) *SampleLogGenerationSubscriptionsResponseMeta {
	this := SampleLogGenerationSubscriptionsResponseMeta{}
	this.TotalSubscriptions = totalSubscriptions
	return &this
}

// NewSampleLogGenerationSubscriptionsResponseMetaWithDefaults instantiates a new SampleLogGenerationSubscriptionsResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationSubscriptionsResponseMetaWithDefaults() *SampleLogGenerationSubscriptionsResponseMeta {
	this := SampleLogGenerationSubscriptionsResponseMeta{}
	return &this
}

// GetTotalSubscriptions returns the TotalSubscriptions field value.
func (o *SampleLogGenerationSubscriptionsResponseMeta) GetTotalSubscriptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.TotalSubscriptions
}

// GetTotalSubscriptionsOk returns a tuple with the TotalSubscriptions field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionsResponseMeta) GetTotalSubscriptionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSubscriptions, true
}

// SetTotalSubscriptions sets field value.
func (o *SampleLogGenerationSubscriptionsResponseMeta) SetTotalSubscriptions(v int32) {
	o.TotalSubscriptions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationSubscriptionsResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["total_subscriptions"] = o.TotalSubscriptions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationSubscriptionsResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalSubscriptions *int32 `json:"total_subscriptions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalSubscriptions == nil {
		return fmt.Errorf("required field total_subscriptions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_subscriptions"})
	} else {
		return err
	}
	o.TotalSubscriptions = *all.TotalSubscriptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
