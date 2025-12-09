// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerAttributes Describes key properties of a Layer, including rotation details, name, start/end times, and any restrictions.
type LayerAttributes struct {
	// When the layer becomes active (ISO 8601).
	EffectiveDate *time.Time `json:"effective_date,omitempty"`
	// When the layer ceases to be active (ISO 8601).
	EndDate *time.Time `json:"end_date,omitempty"`
	// Defines how often the rotation repeats, using a combination of days and optional seconds. Should be at least 1 hour.
	Interval *LayerAttributesInterval `json:"interval,omitempty"`
	// The name of this layer.
	Name *string `json:"name,omitempty"`
	// An optional list of time restrictions for when this layer is in effect.
	Restrictions []TimeRestriction `json:"restrictions,omitempty"`
	// The date/time when the rotation starts (ISO 8601).
	RotationStart *time.Time `json:"rotation_start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLayerAttributes instantiates a new LayerAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLayerAttributes() *LayerAttributes {
	this := LayerAttributes{}
	return &this
}

// NewLayerAttributesWithDefaults instantiates a new LayerAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLayerAttributesWithDefaults() *LayerAttributes {
	this := LayerAttributes{}
	return &this
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *LayerAttributes) GetEffectiveDate() time.Time {
	if o == nil || o.EffectiveDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributes) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil || o.EffectiveDate == nil {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *LayerAttributes) HasEffectiveDate() bool {
	return o != nil && o.EffectiveDate != nil
}

// SetEffectiveDate gets a reference to the given time.Time and assigns it to the EffectiveDate field.
func (o *LayerAttributes) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *LayerAttributes) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributes) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *LayerAttributes) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *LayerAttributes) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *LayerAttributes) GetInterval() LayerAttributesInterval {
	if o == nil || o.Interval == nil {
		var ret LayerAttributesInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributes) GetIntervalOk() (*LayerAttributesInterval, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *LayerAttributes) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given LayerAttributesInterval and assigns it to the Interval field.
func (o *LayerAttributes) SetInterval(v LayerAttributesInterval) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LayerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LayerAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LayerAttributes) SetName(v string) {
	o.Name = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *LayerAttributes) GetRestrictions() []TimeRestriction {
	if o == nil || o.Restrictions == nil {
		var ret []TimeRestriction
		return ret
	}
	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributes) GetRestrictionsOk() (*[]TimeRestriction, bool) {
	if o == nil || o.Restrictions == nil {
		return nil, false
	}
	return &o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *LayerAttributes) HasRestrictions() bool {
	return o != nil && o.Restrictions != nil
}

// SetRestrictions gets a reference to the given []TimeRestriction and assigns it to the Restrictions field.
func (o *LayerAttributes) SetRestrictions(v []TimeRestriction) {
	o.Restrictions = v
}

// GetRotationStart returns the RotationStart field value if set, zero value otherwise.
func (o *LayerAttributes) GetRotationStart() time.Time {
	if o == nil || o.RotationStart == nil {
		var ret time.Time
		return ret
	}
	return *o.RotationStart
}

// GetRotationStartOk returns a tuple with the RotationStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayerAttributes) GetRotationStartOk() (*time.Time, bool) {
	if o == nil || o.RotationStart == nil {
		return nil, false
	}
	return o.RotationStart, true
}

// HasRotationStart returns a boolean if a field has been set.
func (o *LayerAttributes) HasRotationStart() bool {
	return o != nil && o.RotationStart != nil
}

// SetRotationStart gets a reference to the given time.Time and assigns it to the RotationStart field.
func (o *LayerAttributes) SetRotationStart(v time.Time) {
	o.RotationStart = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LayerAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
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
func (o *LayerAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EffectiveDate *time.Time               `json:"effective_date,omitempty"`
		EndDate       *time.Time               `json:"end_date,omitempty"`
		Interval      *LayerAttributesInterval `json:"interval,omitempty"`
		Name          *string                  `json:"name,omitempty"`
		Restrictions  []TimeRestriction        `json:"restrictions,omitempty"`
		RotationStart *time.Time               `json:"rotation_start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"effective_date", "end_date", "interval", "name", "restrictions", "rotation_start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EffectiveDate = all.EffectiveDate
	o.EndDate = all.EndDate
	if all.Interval != nil && all.Interval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Interval = all.Interval
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
