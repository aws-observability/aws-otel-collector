// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorDowntimeMatchResponseAttributes Downtime match details.
type MonitorDowntimeMatchResponseAttributes struct {
	// The end of the downtime.
	End datadog.NullableTime `json:"end,omitempty"`
	// An array of groups associated with the downtime.
	Groups []string `json:"groups,omitempty"`
	// The scope to which the downtime applies. Must follow the [common search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/).
	Scope *string `json:"scope,omitempty"`
	// The start of the downtime.
	Start *time.Time `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorDowntimeMatchResponseAttributes instantiates a new MonitorDowntimeMatchResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDowntimeMatchResponseAttributes() *MonitorDowntimeMatchResponseAttributes {
	this := MonitorDowntimeMatchResponseAttributes{}
	return &this
}

// NewMonitorDowntimeMatchResponseAttributesWithDefaults instantiates a new MonitorDowntimeMatchResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDowntimeMatchResponseAttributesWithDefaults() *MonitorDowntimeMatchResponseAttributes {
	this := MonitorDowntimeMatchResponseAttributes{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorDowntimeMatchResponseAttributes) GetEnd() time.Time {
	if o == nil || o.End.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorDowntimeMatchResponseAttributes) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseAttributes) HasEnd() bool {
	return o != nil && o.End.IsSet()
}

// SetEnd gets a reference to the given datadog.NullableTime and assigns it to the End field.
func (o *MonitorDowntimeMatchResponseAttributes) SetEnd(v time.Time) {
	o.End.Set(&v)
}

// SetEndNil sets the value for End to be an explicit nil.
func (o *MonitorDowntimeMatchResponseAttributes) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil.
func (o *MonitorDowntimeMatchResponseAttributes) UnsetEnd() {
	o.End.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseAttributes) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseAttributes) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseAttributes) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *MonitorDowntimeMatchResponseAttributes) SetGroups(v []string) {
	o.Groups = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseAttributes) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseAttributes) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *MonitorDowntimeMatchResponseAttributes) SetScope(v string) {
	o.Scope = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseAttributes) GetStart() time.Time {
	if o == nil || o.Start == nil {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseAttributes) GetStartOk() (*time.Time, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseAttributes) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *MonitorDowntimeMatchResponseAttributes) SetStart(v time.Time) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorDowntimeMatchResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Start != nil {
		if o.Start.Nanosecond() == 0 {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorDowntimeMatchResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End    datadog.NullableTime `json:"end,omitempty"`
		Groups []string             `json:"groups,omitempty"`
		Scope  *string              `json:"scope,omitempty"`
		Start  *time.Time           `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "groups", "scope", "start"})
	} else {
		return err
	}
	o.End = all.End
	o.Groups = all.Groups
	o.Scope = all.Scope
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
