// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCTriageEvent A single entry in an indicator's triage history timeline.
type IoCTriageEvent struct {
	// Current triage state of the indicator.
	TriageState *IoCTriageState `json:"triage_state,omitempty"`
	// Timestamp when this triage action occurred.
	TriagedAt *time.Time `json:"triaged_at,omitempty"`
	// UUID of the user who performed this triage action.
	TriagedBy *string `json:"triaged_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCTriageEvent instantiates a new IoCTriageEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCTriageEvent() *IoCTriageEvent {
	this := IoCTriageEvent{}
	return &this
}

// NewIoCTriageEventWithDefaults instantiates a new IoCTriageEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCTriageEventWithDefaults() *IoCTriageEvent {
	this := IoCTriageEvent{}
	return &this
}

// GetTriageState returns the TriageState field value if set, zero value otherwise.
func (o *IoCTriageEvent) GetTriageState() IoCTriageState {
	if o == nil || o.TriageState == nil {
		var ret IoCTriageState
		return ret
	}
	return *o.TriageState
}

// GetTriageStateOk returns a tuple with the TriageState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCTriageEvent) GetTriageStateOk() (*IoCTriageState, bool) {
	if o == nil || o.TriageState == nil {
		return nil, false
	}
	return o.TriageState, true
}

// HasTriageState returns a boolean if a field has been set.
func (o *IoCTriageEvent) HasTriageState() bool {
	return o != nil && o.TriageState != nil
}

// SetTriageState gets a reference to the given IoCTriageState and assigns it to the TriageState field.
func (o *IoCTriageEvent) SetTriageState(v IoCTriageState) {
	o.TriageState = &v
}

// GetTriagedAt returns the TriagedAt field value if set, zero value otherwise.
func (o *IoCTriageEvent) GetTriagedAt() time.Time {
	if o == nil || o.TriagedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.TriagedAt
}

// GetTriagedAtOk returns a tuple with the TriagedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCTriageEvent) GetTriagedAtOk() (*time.Time, bool) {
	if o == nil || o.TriagedAt == nil {
		return nil, false
	}
	return o.TriagedAt, true
}

// HasTriagedAt returns a boolean if a field has been set.
func (o *IoCTriageEvent) HasTriagedAt() bool {
	return o != nil && o.TriagedAt != nil
}

// SetTriagedAt gets a reference to the given time.Time and assigns it to the TriagedAt field.
func (o *IoCTriageEvent) SetTriagedAt(v time.Time) {
	o.TriagedAt = &v
}

// GetTriagedBy returns the TriagedBy field value if set, zero value otherwise.
func (o *IoCTriageEvent) GetTriagedBy() string {
	if o == nil || o.TriagedBy == nil {
		var ret string
		return ret
	}
	return *o.TriagedBy
}

// GetTriagedByOk returns a tuple with the TriagedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCTriageEvent) GetTriagedByOk() (*string, bool) {
	if o == nil || o.TriagedBy == nil {
		return nil, false
	}
	return o.TriagedBy, true
}

// HasTriagedBy returns a boolean if a field has been set.
func (o *IoCTriageEvent) HasTriagedBy() bool {
	return o != nil && o.TriagedBy != nil
}

// SetTriagedBy gets a reference to the given string and assigns it to the TriagedBy field.
func (o *IoCTriageEvent) SetTriagedBy(v string) {
	o.TriagedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCTriageEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TriageState != nil {
		toSerialize["triage_state"] = o.TriageState
	}
	if o.TriagedAt != nil {
		if o.TriagedAt.Nanosecond() == 0 {
			toSerialize["triaged_at"] = o.TriagedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["triaged_at"] = o.TriagedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.TriagedBy != nil {
		toSerialize["triaged_by"] = o.TriagedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCTriageEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TriageState *IoCTriageState `json:"triage_state,omitempty"`
		TriagedAt   *time.Time      `json:"triaged_at,omitempty"`
		TriagedBy   *string         `json:"triaged_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"triage_state", "triaged_at", "triaged_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TriageState != nil && !all.TriageState.IsValid() {
		hasInvalidField = true
	} else {
		o.TriageState = all.TriageState
	}
	o.TriagedAt = all.TriagedAt
	o.TriagedBy = all.TriagedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
