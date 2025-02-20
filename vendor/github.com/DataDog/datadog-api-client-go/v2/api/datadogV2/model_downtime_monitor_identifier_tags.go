// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeMonitorIdentifierTags Object of the monitor tags.
type DowntimeMonitorIdentifierTags struct {
	// A list of monitor tags. For example, tags that are applied directly to monitors,
	// not tags that are used in monitor queries (which are filtered by the scope parameter), to which the downtime applies.
	// The resulting downtime applies to monitors that match **all** provided monitor tags. Setting `monitor_tags`
	// to `[*]` configures the downtime to mute all monitors for the given scope.
	MonitorTags []string `json:"monitor_tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeMonitorIdentifierTags instantiates a new DowntimeMonitorIdentifierTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeMonitorIdentifierTags(monitorTags []string) *DowntimeMonitorIdentifierTags {
	this := DowntimeMonitorIdentifierTags{}
	this.MonitorTags = monitorTags
	return &this
}

// NewDowntimeMonitorIdentifierTagsWithDefaults instantiates a new DowntimeMonitorIdentifierTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeMonitorIdentifierTagsWithDefaults() *DowntimeMonitorIdentifierTags {
	this := DowntimeMonitorIdentifierTags{}
	return &this
}

// GetMonitorTags returns the MonitorTags field value.
func (o *DowntimeMonitorIdentifierTags) GetMonitorTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MonitorTags
}

// GetMonitorTagsOk returns a tuple with the MonitorTags field value
// and a boolean to check if the value has been set.
func (o *DowntimeMonitorIdentifierTags) GetMonitorTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorTags, true
}

// SetMonitorTags sets field value.
func (o *DowntimeMonitorIdentifierTags) SetMonitorTags(v []string) {
	o.MonitorTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeMonitorIdentifierTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["monitor_tags"] = o.MonitorTags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeMonitorIdentifierTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MonitorTags *[]string `json:"monitor_tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorTags == nil {
		return fmt.Errorf("required field monitor_tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"monitor_tags"})
	} else {
		return err
	}
	o.MonitorTags = *all.MonitorTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
