// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesBatchResponseMeta Metadata pertaining to the bulk operation.
type OutcomesBatchResponseMeta struct {
	// Total number of scorecard results received during the bulk operation.
	TotalReceived *int64 `json:"total_received,omitempty"`
	// Total number of scorecard results modified during the bulk operation.
	TotalUpdated *int64 `json:"total_updated,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutcomesBatchResponseMeta instantiates a new OutcomesBatchResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutcomesBatchResponseMeta() *OutcomesBatchResponseMeta {
	this := OutcomesBatchResponseMeta{}
	return &this
}

// NewOutcomesBatchResponseMetaWithDefaults instantiates a new OutcomesBatchResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutcomesBatchResponseMetaWithDefaults() *OutcomesBatchResponseMeta {
	this := OutcomesBatchResponseMeta{}
	return &this
}

// GetTotalReceived returns the TotalReceived field value if set, zero value otherwise.
func (o *OutcomesBatchResponseMeta) GetTotalReceived() int64 {
	if o == nil || o.TotalReceived == nil {
		var ret int64
		return ret
	}
	return *o.TotalReceived
}

// GetTotalReceivedOk returns a tuple with the TotalReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseMeta) GetTotalReceivedOk() (*int64, bool) {
	if o == nil || o.TotalReceived == nil {
		return nil, false
	}
	return o.TotalReceived, true
}

// HasTotalReceived returns a boolean if a field has been set.
func (o *OutcomesBatchResponseMeta) HasTotalReceived() bool {
	return o != nil && o.TotalReceived != nil
}

// SetTotalReceived gets a reference to the given int64 and assigns it to the TotalReceived field.
func (o *OutcomesBatchResponseMeta) SetTotalReceived(v int64) {
	o.TotalReceived = &v
}

// GetTotalUpdated returns the TotalUpdated field value if set, zero value otherwise.
func (o *OutcomesBatchResponseMeta) GetTotalUpdated() int64 {
	if o == nil || o.TotalUpdated == nil {
		var ret int64
		return ret
	}
	return *o.TotalUpdated
}

// GetTotalUpdatedOk returns a tuple with the TotalUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseMeta) GetTotalUpdatedOk() (*int64, bool) {
	if o == nil || o.TotalUpdated == nil {
		return nil, false
	}
	return o.TotalUpdated, true
}

// HasTotalUpdated returns a boolean if a field has been set.
func (o *OutcomesBatchResponseMeta) HasTotalUpdated() bool {
	return o != nil && o.TotalUpdated != nil
}

// SetTotalUpdated gets a reference to the given int64 and assigns it to the TotalUpdated field.
func (o *OutcomesBatchResponseMeta) SetTotalUpdated(v int64) {
	o.TotalUpdated = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutcomesBatchResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TotalReceived != nil {
		toSerialize["total_received"] = o.TotalReceived
	}
	if o.TotalUpdated != nil {
		toSerialize["total_updated"] = o.TotalUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutcomesBatchResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalReceived *int64 `json:"total_received,omitempty"`
		TotalUpdated  *int64 `json:"total_updated,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_received", "total_updated"})
	} else {
		return err
	}
	o.TotalReceived = all.TotalReceived
	o.TotalUpdated = all.TotalUpdated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
