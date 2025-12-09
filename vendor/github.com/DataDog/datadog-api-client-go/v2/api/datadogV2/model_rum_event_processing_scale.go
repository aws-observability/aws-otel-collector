// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMEventProcessingScale RUM event processing scale configuration.
type RUMEventProcessingScale struct {
	// Timestamp in milliseconds when this scale was last modified.
	LastModifiedAt *int64 `json:"last_modified_at,omitempty"`
	// Configures which RUM events are processed and stored for the application.
	State *RUMEventProcessingState `json:"state,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMEventProcessingScale instantiates a new RUMEventProcessingScale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMEventProcessingScale() *RUMEventProcessingScale {
	this := RUMEventProcessingScale{}
	return &this
}

// NewRUMEventProcessingScaleWithDefaults instantiates a new RUMEventProcessingScale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMEventProcessingScaleWithDefaults() *RUMEventProcessingScale {
	this := RUMEventProcessingScale{}
	return &this
}

// GetLastModifiedAt returns the LastModifiedAt field value if set, zero value otherwise.
func (o *RUMEventProcessingScale) GetLastModifiedAt() int64 {
	if o == nil || o.LastModifiedAt == nil {
		var ret int64
		return ret
	}
	return *o.LastModifiedAt
}

// GetLastModifiedAtOk returns a tuple with the LastModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMEventProcessingScale) GetLastModifiedAtOk() (*int64, bool) {
	if o == nil || o.LastModifiedAt == nil {
		return nil, false
	}
	return o.LastModifiedAt, true
}

// HasLastModifiedAt returns a boolean if a field has been set.
func (o *RUMEventProcessingScale) HasLastModifiedAt() bool {
	return o != nil && o.LastModifiedAt != nil
}

// SetLastModifiedAt gets a reference to the given int64 and assigns it to the LastModifiedAt field.
func (o *RUMEventProcessingScale) SetLastModifiedAt(v int64) {
	o.LastModifiedAt = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RUMEventProcessingScale) GetState() RUMEventProcessingState {
	if o == nil || o.State == nil {
		var ret RUMEventProcessingState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMEventProcessingScale) GetStateOk() (*RUMEventProcessingState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RUMEventProcessingScale) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given RUMEventProcessingState and assigns it to the State field.
func (o *RUMEventProcessingScale) SetState(v RUMEventProcessingState) {
	o.State = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMEventProcessingScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LastModifiedAt != nil {
		toSerialize["last_modified_at"] = o.LastModifiedAt
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMEventProcessingScale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LastModifiedAt *int64                   `json:"last_modified_at,omitempty"`
		State          *RUMEventProcessingState `json:"state,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"last_modified_at", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.LastModifiedAt = all.LastModifiedAt
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
