// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFetchUptimesPayload Object containing IDs of Synthetic tests and a timeframe.
type SyntheticsFetchUptimesPayload struct {
	// Timestamp in seconds (Unix epoch) for the start of uptime.
	FromTs int64 `json:"from_ts"`
	// An array of Synthetic test IDs you want uptimes for.
	PublicIds []string `json:"public_ids"`
	// Timestamp in seconds (Unix epoch) for the end of uptime.
	ToTs int64 `json:"to_ts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsFetchUptimesPayload instantiates a new SyntheticsFetchUptimesPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsFetchUptimesPayload(fromTs int64, publicIds []string, toTs int64) *SyntheticsFetchUptimesPayload {
	this := SyntheticsFetchUptimesPayload{}
	this.FromTs = fromTs
	this.PublicIds = publicIds
	this.ToTs = toTs
	return &this
}

// NewSyntheticsFetchUptimesPayloadWithDefaults instantiates a new SyntheticsFetchUptimesPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsFetchUptimesPayloadWithDefaults() *SyntheticsFetchUptimesPayload {
	this := SyntheticsFetchUptimesPayload{}
	return &this
}

// GetFromTs returns the FromTs field value.
func (o *SyntheticsFetchUptimesPayload) GetFromTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value
// and a boolean to check if the value has been set.
func (o *SyntheticsFetchUptimesPayload) GetFromTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTs, true
}

// SetFromTs sets field value.
func (o *SyntheticsFetchUptimesPayload) SetFromTs(v int64) {
	o.FromTs = v
}

// GetPublicIds returns the PublicIds field value.
func (o *SyntheticsFetchUptimesPayload) GetPublicIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PublicIds
}

// GetPublicIdsOk returns a tuple with the PublicIds field value
// and a boolean to check if the value has been set.
func (o *SyntheticsFetchUptimesPayload) GetPublicIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicIds, true
}

// SetPublicIds sets field value.
func (o *SyntheticsFetchUptimesPayload) SetPublicIds(v []string) {
	o.PublicIds = v
}

// GetToTs returns the ToTs field value.
func (o *SyntheticsFetchUptimesPayload) GetToTs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value
// and a boolean to check if the value has been set.
func (o *SyntheticsFetchUptimesPayload) GetToTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTs, true
}

// SetToTs sets field value.
func (o *SyntheticsFetchUptimesPayload) SetToTs(v int64) {
	o.ToTs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsFetchUptimesPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from_ts"] = o.FromTs
	toSerialize["public_ids"] = o.PublicIds
	toSerialize["to_ts"] = o.ToTs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsFetchUptimesPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FromTs    *int64    `json:"from_ts"`
		PublicIds *[]string `json:"public_ids"`
		ToTs      *int64    `json:"to_ts"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FromTs == nil {
		return fmt.Errorf("required field from_ts missing")
	}
	if all.PublicIds == nil {
		return fmt.Errorf("required field public_ids missing")
	}
	if all.ToTs == nil {
		return fmt.Errorf("required field to_ts missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from_ts", "public_ids", "to_ts"})
	} else {
		return err
	}
	o.FromTs = *all.FromTs
	o.PublicIds = *all.PublicIds
	o.ToTs = *all.ToTs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
