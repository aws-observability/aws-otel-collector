// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityContextRevision A single historical revision of an entity, including the time range during which the revision was observed.
type EntityContextRevision struct {
	// The set of attributes recorded for the entity at this revision. The keys depend on the kind of entity.
	Attributes map[string]interface{} `json:"attributes"`
	// The first time the entity was observed at this revision.
	FirstSeenAt time.Time `json:"first_seen_at"`
	// The last time the entity was observed at this revision.
	LastSeenAt time.Time `json:"last_seen_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityContextRevision instantiates a new EntityContextRevision object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityContextRevision(attributes map[string]interface{}, firstSeenAt time.Time, lastSeenAt time.Time) *EntityContextRevision {
	this := EntityContextRevision{}
	this.Attributes = attributes
	this.FirstSeenAt = firstSeenAt
	this.LastSeenAt = lastSeenAt
	return &this
}

// NewEntityContextRevisionWithDefaults instantiates a new EntityContextRevision object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityContextRevisionWithDefaults() *EntityContextRevision {
	this := EntityContextRevision{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *EntityContextRevision) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EntityContextRevision) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *EntityContextRevision) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetFirstSeenAt returns the FirstSeenAt field value.
func (o *EntityContextRevision) GetFirstSeenAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value
// and a boolean to check if the value has been set.
func (o *EntityContextRevision) GetFirstSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSeenAt, true
}

// SetFirstSeenAt sets field value.
func (o *EntityContextRevision) SetFirstSeenAt(v time.Time) {
	o.FirstSeenAt = v
}

// GetLastSeenAt returns the LastSeenAt field value.
func (o *EntityContextRevision) GetLastSeenAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value
// and a boolean to check if the value has been set.
func (o *EntityContextRevision) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeenAt, true
}

// SetLastSeenAt sets field value.
func (o *EntityContextRevision) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityContextRevision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.FirstSeenAt.Nanosecond() == 0 {
		toSerialize["first_seen_at"] = o.FirstSeenAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["first_seen_at"] = o.FirstSeenAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.LastSeenAt.Nanosecond() == 0 {
		toSerialize["last_seen_at"] = o.LastSeenAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_seen_at"] = o.LastSeenAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityContextRevision) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes  *map[string]interface{} `json:"attributes"`
		FirstSeenAt *time.Time              `json:"first_seen_at"`
		LastSeenAt  *time.Time              `json:"last_seen_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.FirstSeenAt == nil {
		return fmt.Errorf("required field first_seen_at missing")
	}
	if all.LastSeenAt == nil {
		return fmt.Errorf("required field last_seen_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "first_seen_at", "last_seen_at"})
	} else {
		return err
	}
	o.Attributes = *all.Attributes
	o.FirstSeenAt = *all.FirstSeenAt
	o.LastSeenAt = *all.LastSeenAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
