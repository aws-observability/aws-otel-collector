// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultOCSPUpdates OCSP response update timestamps.
type SyntheticsTestResultOCSPUpdates struct {
	// Unix timestamp (ms) of the next expected OCSP update.
	NextUpdate *int64 `json:"next_update,omitempty"`
	// Unix timestamp (ms) of when the OCSP response was produced.
	ProducedAt *int64 `json:"produced_at,omitempty"`
	// Unix timestamp (ms) of this OCSP update.
	ThisUpdate *int64 `json:"this_update,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultOCSPUpdates instantiates a new SyntheticsTestResultOCSPUpdates object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultOCSPUpdates() *SyntheticsTestResultOCSPUpdates {
	this := SyntheticsTestResultOCSPUpdates{}
	return &this
}

// NewSyntheticsTestResultOCSPUpdatesWithDefaults instantiates a new SyntheticsTestResultOCSPUpdates object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultOCSPUpdatesWithDefaults() *SyntheticsTestResultOCSPUpdates {
	this := SyntheticsTestResultOCSPUpdates{}
	return &this
}

// GetNextUpdate returns the NextUpdate field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPUpdates) GetNextUpdate() int64 {
	if o == nil || o.NextUpdate == nil {
		var ret int64
		return ret
	}
	return *o.NextUpdate
}

// GetNextUpdateOk returns a tuple with the NextUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPUpdates) GetNextUpdateOk() (*int64, bool) {
	if o == nil || o.NextUpdate == nil {
		return nil, false
	}
	return o.NextUpdate, true
}

// HasNextUpdate returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPUpdates) HasNextUpdate() bool {
	return o != nil && o.NextUpdate != nil
}

// SetNextUpdate gets a reference to the given int64 and assigns it to the NextUpdate field.
func (o *SyntheticsTestResultOCSPUpdates) SetNextUpdate(v int64) {
	o.NextUpdate = &v
}

// GetProducedAt returns the ProducedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPUpdates) GetProducedAt() int64 {
	if o == nil || o.ProducedAt == nil {
		var ret int64
		return ret
	}
	return *o.ProducedAt
}

// GetProducedAtOk returns a tuple with the ProducedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPUpdates) GetProducedAtOk() (*int64, bool) {
	if o == nil || o.ProducedAt == nil {
		return nil, false
	}
	return o.ProducedAt, true
}

// HasProducedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPUpdates) HasProducedAt() bool {
	return o != nil && o.ProducedAt != nil
}

// SetProducedAt gets a reference to the given int64 and assigns it to the ProducedAt field.
func (o *SyntheticsTestResultOCSPUpdates) SetProducedAt(v int64) {
	o.ProducedAt = &v
}

// GetThisUpdate returns the ThisUpdate field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPUpdates) GetThisUpdate() int64 {
	if o == nil || o.ThisUpdate == nil {
		var ret int64
		return ret
	}
	return *o.ThisUpdate
}

// GetThisUpdateOk returns a tuple with the ThisUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPUpdates) GetThisUpdateOk() (*int64, bool) {
	if o == nil || o.ThisUpdate == nil {
		return nil, false
	}
	return o.ThisUpdate, true
}

// HasThisUpdate returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPUpdates) HasThisUpdate() bool {
	return o != nil && o.ThisUpdate != nil
}

// SetThisUpdate gets a reference to the given int64 and assigns it to the ThisUpdate field.
func (o *SyntheticsTestResultOCSPUpdates) SetThisUpdate(v int64) {
	o.ThisUpdate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultOCSPUpdates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextUpdate != nil {
		toSerialize["next_update"] = o.NextUpdate
	}
	if o.ProducedAt != nil {
		toSerialize["produced_at"] = o.ProducedAt
	}
	if o.ThisUpdate != nil {
		toSerialize["this_update"] = o.ThisUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultOCSPUpdates) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextUpdate *int64 `json:"next_update,omitempty"`
		ProducedAt *int64 `json:"produced_at,omitempty"`
		ThisUpdate *int64 `json:"this_update,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_update", "produced_at", "this_update"})
	} else {
		return err
	}
	o.NextUpdate = all.NextUpdate
	o.ProducedAt = all.ProducedAt
	o.ThisUpdate = all.ThisUpdate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
