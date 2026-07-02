// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionsResponse Response containing a list of sample log generation subscriptions.
type SampleLogGenerationSubscriptionsResponse struct {
	// The list of sample log generation subscriptions.
	Data []SampleLogGenerationSubscriptionData `json:"data"`
	// Metadata returned alongside a list of sample log generation subscriptions.
	Meta SampleLogGenerationSubscriptionsResponseMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationSubscriptionsResponse instantiates a new SampleLogGenerationSubscriptionsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationSubscriptionsResponse(data []SampleLogGenerationSubscriptionData, meta SampleLogGenerationSubscriptionsResponseMeta) *SampleLogGenerationSubscriptionsResponse {
	this := SampleLogGenerationSubscriptionsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewSampleLogGenerationSubscriptionsResponseWithDefaults instantiates a new SampleLogGenerationSubscriptionsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationSubscriptionsResponseWithDefaults() *SampleLogGenerationSubscriptionsResponse {
	this := SampleLogGenerationSubscriptionsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *SampleLogGenerationSubscriptionsResponse) GetData() []SampleLogGenerationSubscriptionData {
	if o == nil {
		var ret []SampleLogGenerationSubscriptionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionsResponse) GetDataOk() (*[]SampleLogGenerationSubscriptionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *SampleLogGenerationSubscriptionsResponse) SetData(v []SampleLogGenerationSubscriptionData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *SampleLogGenerationSubscriptionsResponse) GetMeta() SampleLogGenerationSubscriptionsResponseMeta {
	if o == nil {
		var ret SampleLogGenerationSubscriptionsResponseMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionsResponse) GetMetaOk() (*SampleLogGenerationSubscriptionsResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *SampleLogGenerationSubscriptionsResponse) SetMeta(v SampleLogGenerationSubscriptionsResponseMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationSubscriptionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]SampleLogGenerationSubscriptionData        `json:"data"`
		Meta *SampleLogGenerationSubscriptionsResponseMeta `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
