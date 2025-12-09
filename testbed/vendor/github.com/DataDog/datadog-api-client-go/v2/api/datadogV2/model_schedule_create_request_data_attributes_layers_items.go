// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataAttributesLayersItems Describes a schedule layer, including rotation intervals, members, restrictions, and timeline settings.
type ScheduleCreateRequestDataAttributesLayersItems struct {
	// The date/time when this layer becomes active (in ISO 8601).
	EffectiveDate time.Time `json:"effective_date"`
	// The date/time after which this layer no longer applies (in ISO 8601).
	EndDate *time.Time `json:"end_date,omitempty"`
	// Defines how often the rotation repeats, using a combination of days and optional seconds. Should be at least 1 hour.
	Interval LayerAttributesInterval `json:"interval"`
	// A list of members who participate in this layer's rotation.
	Members []ScheduleRequestDataAttributesLayersItemsMembersItems `json:"members"`
	// The name of this layer.
	Name string `json:"name"`
	// Zero or more time-based restrictions (for example, only weekdays, during business hours).
	Restrictions []TimeRestriction `json:"restrictions,omitempty"`
	// The date/time when the rotation for this layer starts (in ISO 8601).
	RotationStart time.Time `json:"rotation_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleCreateRequestDataAttributesLayersItems instantiates a new ScheduleCreateRequestDataAttributesLayersItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleCreateRequestDataAttributesLayersItems(effectiveDate time.Time, interval LayerAttributesInterval, members []ScheduleRequestDataAttributesLayersItemsMembersItems, name string, rotationStart time.Time) *ScheduleCreateRequestDataAttributesLayersItems {
	this := ScheduleCreateRequestDataAttributesLayersItems{}
	this.EffectiveDate = effectiveDate
	this.Interval = interval
	this.Members = members
	this.Name = name
	this.RotationStart = rotationStart
	return &this
}

// NewScheduleCreateRequestDataAttributesLayersItemsWithDefaults instantiates a new ScheduleCreateRequestDataAttributesLayersItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleCreateRequestDataAttributesLayersItemsWithDefaults() *ScheduleCreateRequestDataAttributesLayersItems {
	this := ScheduleCreateRequestDataAttributesLayersItems{}
	return &this
}

// GetEffectiveDate returns the EffectiveDate field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetInterval returns the Interval field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetInterval() LayerAttributesInterval {
	if o == nil {
		var ret LayerAttributesInterval
		return ret
	}
	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetIntervalOk() (*LayerAttributesInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetInterval(v LayerAttributesInterval) {
	o.Interval = v
}

// GetMembers returns the Members field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetMembers() []ScheduleRequestDataAttributesLayersItemsMembersItems {
	if o == nil {
		var ret []ScheduleRequestDataAttributesLayersItemsMembersItems
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetMembersOk() (*[]ScheduleRequestDataAttributesLayersItemsMembersItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetMembers(v []ScheduleRequestDataAttributesLayersItemsMembersItems) {
	o.Members = v
}

// GetName returns the Name field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetName(v string) {
	o.Name = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetRestrictions() []TimeRestriction {
	if o == nil || o.Restrictions == nil {
		var ret []TimeRestriction
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetRestrictionsOk() (*[]TimeRestriction, bool) {
	if o == nil || o.Restrictions == nil {
		return nil, false
	}
	return &o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) HasRestrictions() bool {
	return o != nil && o.Restrictions != nil
}

// SetRestrictions gets a reference to the given []TimeRestriction and assigns it to the Restrictions field.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetRestrictions(v []TimeRestriction) {
	o.Restrictions = v
}

// GetRotationStart returns the RotationStart field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetRotationStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.RotationStart
}

// GetRotationStartOk returns a tuple with the RotationStart field value
// and a boolean to check if the value has been set.
func (o *ScheduleCreateRequestDataAttributesLayersItems) GetRotationStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationStart, true
}

// SetRotationStart sets field value.
func (o *ScheduleCreateRequestDataAttributesLayersItems) SetRotationStart(v time.Time) {
	o.RotationStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleCreateRequestDataAttributesLayersItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EffectiveDate.Nanosecond() == 0 {
		toSerialize["effective_date"] = o.EffectiveDate.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["effective_date"] = o.EffectiveDate.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["interval"] = o.Interval
	toSerialize["members"] = o.Members
	toSerialize["name"] = o.Name
	if o.Restrictions != nil {
		toSerialize["restrictions"] = o.Restrictions
	}
	if o.RotationStart.Nanosecond() == 0 {
		toSerialize["rotation_start"] = o.RotationStart.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["rotation_start"] = o.RotationStart.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleCreateRequestDataAttributesLayersItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EffectiveDate *time.Time                                              `json:"effective_date"`
		EndDate       *time.Time                                              `json:"end_date,omitempty"`
		Interval      *LayerAttributesInterval                                `json:"interval"`
		Members       *[]ScheduleRequestDataAttributesLayersItemsMembersItems `json:"members"`
		Name          *string                                                 `json:"name"`
		Restrictions  []TimeRestriction                                       `json:"restrictions,omitempty"`
		RotationStart *time.Time                                              `json:"rotation_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EffectiveDate == nil {
		return fmt.Errorf("required field effective_date missing")
	}
	if all.Interval == nil {
		return fmt.Errorf("required field interval missing")
	}
	if all.Members == nil {
		return fmt.Errorf("required field members missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RotationStart == nil {
		return fmt.Errorf("required field rotation_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effective_date", "end_date", "interval", "members", "name", "restrictions", "rotation_start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EffectiveDate = *all.EffectiveDate
	o.EndDate = all.EndDate
	if all.Interval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Interval = *all.Interval
	o.Members = *all.Members
	o.Name = *all.Name
	o.Restrictions = all.Restrictions
	o.RotationStart = *all.RotationStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
