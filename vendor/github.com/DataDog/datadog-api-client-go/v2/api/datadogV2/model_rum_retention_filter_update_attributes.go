// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterUpdateAttributes The object describing attributes of a RUM retention filter to update.
type RumRetentionFilterUpdateAttributes struct {
	// Whether the retention filter is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The type of RUM events to filter on.
	EventType *RumRetentionFilterEventType `json:"event_type,omitempty"`
	// The name of a RUM retention filter.
	Name *string `json:"name,omitempty"`
	// The query string for a RUM retention filter.
	Query *string `json:"query,omitempty"`
	// The sample rate for a RUM retention filter, between 0 and 100.
	SampleRate *int64 `json:"sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRetentionFilterUpdateAttributes instantiates a new RumRetentionFilterUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRetentionFilterUpdateAttributes() *RumRetentionFilterUpdateAttributes {
	this := RumRetentionFilterUpdateAttributes{}
	return &this
}

// NewRumRetentionFilterUpdateAttributesWithDefaults instantiates a new RumRetentionFilterUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRetentionFilterUpdateAttributesWithDefaults() *RumRetentionFilterUpdateAttributes {
	this := RumRetentionFilterUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RumRetentionFilterUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RumRetentionFilterUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RumRetentionFilterUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RumRetentionFilterUpdateAttributes) GetEventType() RumRetentionFilterEventType {
	if o == nil || o.EventType == nil {
		var ret RumRetentionFilterEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterUpdateAttributes) GetEventTypeOk() (*RumRetentionFilterEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RumRetentionFilterUpdateAttributes) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given RumRetentionFilterEventType and assigns it to the EventType field.
func (o *RumRetentionFilterUpdateAttributes) SetEventType(v RumRetentionFilterEventType) {
	o.EventType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RumRetentionFilterUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RumRetentionFilterUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RumRetentionFilterUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RumRetentionFilterUpdateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterUpdateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RumRetentionFilterUpdateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RumRetentionFilterUpdateAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *RumRetentionFilterUpdateAttributes) GetSampleRate() int64 {
	if o == nil || o.SampleRate == nil {
		var ret int64
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterUpdateAttributes) GetSampleRateOk() (*int64, bool) {
	if o == nil || o.SampleRate == nil {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *RumRetentionFilterUpdateAttributes) HasSampleRate() bool {
	return o != nil && o.SampleRate != nil
}

// SetSampleRate gets a reference to the given int64 and assigns it to the SampleRate field.
func (o *RumRetentionFilterUpdateAttributes) SetSampleRate(v int64) {
	o.SampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRetentionFilterUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.SampleRate != nil {
		toSerialize["sample_rate"] = o.SampleRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRetentionFilterUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled    *bool                        `json:"enabled,omitempty"`
		EventType  *RumRetentionFilterEventType `json:"event_type,omitempty"`
		Name       *string                      `json:"name,omitempty"`
		Query      *string                      `json:"query,omitempty"`
		SampleRate *int64                       `json:"sample_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "event_type", "name", "query", "sample_rate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.EventType != nil && !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
	}
	o.Name = all.Name
	o.Query = all.Query
	o.SampleRate = all.SampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
