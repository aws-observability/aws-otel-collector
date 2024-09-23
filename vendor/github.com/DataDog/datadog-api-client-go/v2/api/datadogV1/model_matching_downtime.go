// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MatchingDowntime Object describing a downtime that matches this monitor.
type MatchingDowntime struct {
	// POSIX timestamp to end the downtime.
	End datadog.NullableInt64 `json:"end,omitempty"`
	// The downtime ID.
	Id int64 `json:"id"`
	// The scope(s) to which the downtime applies. Must be in `key:value` format. For example, `host:app2`.
	// Provide multiple scopes as a comma-separated list like `env:dev,env:prod`.
	// The resulting downtime applies to sources that matches ALL provided scopes (`env:dev` **AND** `env:prod`).
	Scope []string `json:"scope,omitempty"`
	// POSIX timestamp to start the downtime.
	Start *int64 `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMatchingDowntime instantiates a new MatchingDowntime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMatchingDowntime(id int64) *MatchingDowntime {
	this := MatchingDowntime{}
	this.Id = id
	return &this
}

// NewMatchingDowntimeWithDefaults instantiates a new MatchingDowntime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMatchingDowntimeWithDefaults() *MatchingDowntime {
	this := MatchingDowntime{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MatchingDowntime) GetEnd() int64 {
	if o == nil || o.End.Get() == nil {
		var ret int64
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MatchingDowntime) GetEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *MatchingDowntime) HasEnd() bool {
	return o != nil && o.End.IsSet()
}

// SetEnd gets a reference to the given datadog.NullableInt64 and assigns it to the End field.
func (o *MatchingDowntime) SetEnd(v int64) {
	o.End.Set(&v)
}

// SetEndNil sets the value for End to be an explicit nil.
func (o *MatchingDowntime) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil.
func (o *MatchingDowntime) UnsetEnd() {
	o.End.Unset()
}

// GetId returns the Id field value.
func (o *MatchingDowntime) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MatchingDowntime) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MatchingDowntime) SetId(v int64) {
	o.Id = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *MatchingDowntime) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchingDowntime) GetScopeOk() (*[]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MatchingDowntime) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *MatchingDowntime) SetScope(v []string) {
	o.Scope = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MatchingDowntime) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchingDowntime) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MatchingDowntime) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *MatchingDowntime) SetStart(v int64) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MatchingDowntime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	toSerialize["id"] = o.Id
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
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
func (o *MatchingDowntime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   datadog.NullableInt64 `json:"end,omitempty"`
		Id    *int64                `json:"id"`
		Scope []string              `json:"scope,omitempty"`
		Start *int64                `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "id", "scope", "start"})
	} else {
		return err
	}
	o.End = all.End
	o.Id = *all.Id
	o.Scope = all.Scope
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
