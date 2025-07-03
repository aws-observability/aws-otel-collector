// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataAttributesLayersItems Represents a layer within a schedule update, including rotation details, members,
// and optional restrictions.
type ScheduleUpdateRequestDataAttributesLayersItems struct {
	// When this updated layer takes effect (ISO 8601 format).
	EffectiveDate *time.Time `json:"effective_date,omitempty"`
	// When this updated layer should stop being active (ISO 8601 format).
	EndDate *time.Time `json:"end_date,omitempty"`
	// A unique identifier for the layer being updated.
	Id *string `json:"id,omitempty"`
	// Specifies how the rotation repeats: number of days, plus optional seconds, up to the given maximums.
	Interval *ScheduleUpdateRequestDataAttributesLayersItemsInterval `json:"interval,omitempty"`
	// The members assigned to this layer.
	Members []ScheduleUpdateRequestDataAttributesLayersItemsMembersItems `json:"members,omitempty"`
	// The name for this layer (for example, “Secondary Coverage”).
	Name *string `json:"name,omitempty"`
	// Any time restrictions that define when this layer is active.
	Restrictions []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems `json:"restrictions,omitempty"`
	// The date/time at which the rotation begins (ISO 8601 format).
	RotationStart *time.Time `json:"rotation_start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestDataAttributesLayersItems instantiates a new ScheduleUpdateRequestDataAttributesLayersItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestDataAttributesLayersItems() *ScheduleUpdateRequestDataAttributesLayersItems {
	this := ScheduleUpdateRequestDataAttributesLayersItems{}
	return &this
}

// NewScheduleUpdateRequestDataAttributesLayersItemsWithDefaults instantiates a new ScheduleUpdateRequestDataAttributesLayersItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataAttributesLayersItemsWithDefaults() *ScheduleUpdateRequestDataAttributesLayersItems {
	this := ScheduleUpdateRequestDataAttributesLayersItems{}
	return &this
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetEffectiveDate() time.Time {
	if o == nil || o.EffectiveDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil || o.EffectiveDate == nil {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasEffectiveDate() bool {
	return o != nil && o.EffectiveDate != nil
}

// SetEffectiveDate gets a reference to the given time.Time and assigns it to the EffectiveDate field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetInterval() ScheduleUpdateRequestDataAttributesLayersItemsInterval {
	if o == nil || o.Interval == nil {
		var ret ScheduleUpdateRequestDataAttributesLayersItemsInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetIntervalOk() (*ScheduleUpdateRequestDataAttributesLayersItemsInterval, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given ScheduleUpdateRequestDataAttributesLayersItemsInterval and assigns it to the Interval field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetInterval(v ScheduleUpdateRequestDataAttributesLayersItemsInterval) {
	o.Interval = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetMembers() []ScheduleUpdateRequestDataAttributesLayersItemsMembersItems {
	if o == nil || o.Members == nil {
		var ret []ScheduleUpdateRequestDataAttributesLayersItemsMembersItems
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetMembersOk() (*[]ScheduleUpdateRequestDataAttributesLayersItemsMembersItems, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return &o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasMembers() bool {
	return o != nil && o.Members != nil
}

// SetMembers gets a reference to the given []ScheduleUpdateRequestDataAttributesLayersItemsMembersItems and assigns it to the Members field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetMembers(v []ScheduleUpdateRequestDataAttributesLayersItemsMembersItems) {
	o.Members = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetName(v string) {
	o.Name = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetRestrictions() []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems {
	if o == nil || o.Restrictions == nil {
		var ret []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetRestrictionsOk() (*[]ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems, bool) {
	if o == nil || o.Restrictions == nil {
		return nil, false
	}
	return &o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasRestrictions() bool {
	return o != nil && o.Restrictions != nil
}

// SetRestrictions gets a reference to the given []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems and assigns it to the Restrictions field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetRestrictions(v []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems) {
	o.Restrictions = v
}

// GetRotationStart returns the RotationStart field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetRotationStart() time.Time {
	if o == nil || o.RotationStart == nil {
		var ret time.Time
		return ret
	}
	return *o.RotationStart
}

// GetRotationStartOk returns a tuple with the RotationStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) GetRotationStartOk() (*time.Time, bool) {
	if o == nil || o.RotationStart == nil {
		return nil, false
	}
	return o.RotationStart, true
}

// HasRotationStart returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) HasRotationStart() bool {
	return o != nil && o.RotationStart != nil
}

// SetRotationStart gets a reference to the given time.Time and assigns it to the RotationStart field.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) SetRotationStart(v time.Time) {
	o.RotationStart = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestDataAttributesLayersItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EffectiveDate != nil {
		if o.EffectiveDate.Nanosecond() == 0 {
			toSerialize["effective_date"] = o.EffectiveDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["effective_date"] = o.EffectiveDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Restrictions != nil {
		toSerialize["restrictions"] = o.Restrictions
	}
	if o.RotationStart != nil {
		if o.RotationStart.Nanosecond() == 0 {
			toSerialize["rotation_start"] = o.RotationStart.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["rotation_start"] = o.RotationStart.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleUpdateRequestDataAttributesLayersItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EffectiveDate *time.Time                                                        `json:"effective_date,omitempty"`
		EndDate       *time.Time                                                        `json:"end_date,omitempty"`
		Id            *string                                                           `json:"id,omitempty"`
		Interval      *ScheduleUpdateRequestDataAttributesLayersItemsInterval           `json:"interval,omitempty"`
		Members       []ScheduleUpdateRequestDataAttributesLayersItemsMembersItems      `json:"members,omitempty"`
		Name          *string                                                           `json:"name,omitempty"`
		Restrictions  []ScheduleUpdateRequestDataAttributesLayersItemsRestrictionsItems `json:"restrictions,omitempty"`
		RotationStart *time.Time                                                        `json:"rotation_start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effective_date", "end_date", "id", "interval", "members", "name", "restrictions", "rotation_start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EffectiveDate = all.EffectiveDate
	o.EndDate = all.EndDate
	o.Id = all.Id
	if all.Interval != nil && all.Interval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Interval = all.Interval
	o.Members = all.Members
	o.Name = all.Name
	o.Restrictions = all.Restrictions
	o.RotationStart = all.RotationStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
