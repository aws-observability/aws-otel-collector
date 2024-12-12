// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestUptime Object containing the uptime for a Synthetic test ID.
type SyntheticsTestUptime struct {
	// Timestamp in seconds for the start of uptime.
	FromTs *int64 `json:"from_ts,omitempty"`
	// Object containing the uptime information.
	Overall *SyntheticsUptime `json:"overall,omitempty"`
	// A Synthetic test ID.
	PublicId *string `json:"public_id,omitempty"`
	// Timestamp in seconds for the end of uptime.
	ToTs *int64 `json:"to_ts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestUptime instantiates a new SyntheticsTestUptime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestUptime() *SyntheticsTestUptime {
	this := SyntheticsTestUptime{}
	return &this
}

// NewSyntheticsTestUptimeWithDefaults instantiates a new SyntheticsTestUptime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestUptimeWithDefaults() *SyntheticsTestUptime {
	this := SyntheticsTestUptime{}
	return &this
}

// GetFromTs returns the FromTs field value if set, zero value otherwise.
func (o *SyntheticsTestUptime) GetFromTs() int64 {
	if o == nil || o.FromTs == nil {
		var ret int64
		return ret
	}
	return *o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestUptime) GetFromTsOk() (*int64, bool) {
	if o == nil || o.FromTs == nil {
		return nil, false
	}
	return o.FromTs, true
}

// HasFromTs returns a boolean if a field has been set.
func (o *SyntheticsTestUptime) HasFromTs() bool {
	return o != nil && o.FromTs != nil
}

// SetFromTs gets a reference to the given int64 and assigns it to the FromTs field.
func (o *SyntheticsTestUptime) SetFromTs(v int64) {
	o.FromTs = &v
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *SyntheticsTestUptime) GetOverall() SyntheticsUptime {
	if o == nil || o.Overall == nil {
		var ret SyntheticsUptime
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestUptime) GetOverallOk() (*SyntheticsUptime, bool) {
	if o == nil || o.Overall == nil {
		return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *SyntheticsTestUptime) HasOverall() bool {
	return o != nil && o.Overall != nil
}

// SetOverall gets a reference to the given SyntheticsUptime and assigns it to the Overall field.
func (o *SyntheticsTestUptime) SetOverall(v SyntheticsUptime) {
	o.Overall = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *SyntheticsTestUptime) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestUptime) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *SyntheticsTestUptime) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *SyntheticsTestUptime) SetPublicId(v string) {
	o.PublicId = &v
}

// GetToTs returns the ToTs field value if set, zero value otherwise.
func (o *SyntheticsTestUptime) GetToTs() int64 {
	if o == nil || o.ToTs == nil {
		var ret int64
		return ret
	}
	return *o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestUptime) GetToTsOk() (*int64, bool) {
	if o == nil || o.ToTs == nil {
		return nil, false
	}
	return o.ToTs, true
}

// HasToTs returns a boolean if a field has been set.
func (o *SyntheticsTestUptime) HasToTs() bool {
	return o != nil && o.ToTs != nil
}

// SetToTs gets a reference to the given int64 and assigns it to the ToTs field.
func (o *SyntheticsTestUptime) SetToTs(v int64) {
	o.ToTs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestUptime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FromTs != nil {
		toSerialize["from_ts"] = o.FromTs
	}
	if o.Overall != nil {
		toSerialize["overall"] = o.Overall
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.ToTs != nil {
		toSerialize["to_ts"] = o.ToTs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestUptime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromTs   *int64            `json:"from_ts,omitempty"`
		Overall  *SyntheticsUptime `json:"overall,omitempty"`
		PublicId *string           `json:"public_id,omitempty"`
		ToTs     *int64            `json:"to_ts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from_ts", "overall", "public_id", "to_ts"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FromTs = all.FromTs
	if all.Overall != nil && all.Overall.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Overall = all.Overall
	o.PublicId = all.PublicId
	o.ToTs = all.ToTs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
