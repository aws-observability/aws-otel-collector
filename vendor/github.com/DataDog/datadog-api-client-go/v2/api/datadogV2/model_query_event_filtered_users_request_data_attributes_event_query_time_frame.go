// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame
type QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame struct {
	//
	End *int64 `json:"end,omitempty"`
	//
	Start *int64 `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame instantiates a new QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame() *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame {
	this := QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame{}
	return &this
}

// NewQueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrameWithDefaults instantiates a new QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrameWithDefaults() *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame {
	this := QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) SetEnd(v int64) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) SetStart(v int64) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *int64 `json:"end,omitempty"`
		Start *int64 `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start"})
	} else {
		return err
	}
	o.End = all.End
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
