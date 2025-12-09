// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactAttributes The incident impact's attributes.
type IncidentImpactAttributes struct {
	// Timestamp when the impact was created.
	Created *time.Time `json:"created,omitempty"`
	// Description of the impact.
	Description string `json:"description"`
	// Timestamp when the impact ended.
	EndAt datadog.NullableTime `json:"end_at,omitempty"`
	// An object mapping impact field names to field values.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// The type of impact.
	ImpactType *string `json:"impact_type,omitempty"`
	// Timestamp when the impact was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// Timestamp representing when the impact started.
	StartAt time.Time `json:"start_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImpactAttributes instantiates a new IncidentImpactAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImpactAttributes(description string, startAt time.Time) *IncidentImpactAttributes {
	this := IncidentImpactAttributes{}
	this.Description = description
	this.StartAt = startAt
	return &this
}

// NewIncidentImpactAttributesWithDefaults instantiates a new IncidentImpactAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImpactAttributesWithDefaults() *IncidentImpactAttributes {
	this := IncidentImpactAttributes{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IncidentImpactAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImpactAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IncidentImpactAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IncidentImpactAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value.
func (o *IncidentImpactAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IncidentImpactAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *IncidentImpactAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImpactAttributes) GetEndAt() time.Time {
	if o == nil || o.EndAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImpactAttributes) GetEndAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *IncidentImpactAttributes) HasEndAt() bool {
	return o != nil && o.EndAt.IsSet()
}

// SetEndAt gets a reference to the given datadog.NullableTime and assigns it to the EndAt field.
func (o *IncidentImpactAttributes) SetEndAt(v time.Time) {
	o.EndAt.Set(&v)
}

// SetEndAtNil sets the value for EndAt to be an explicit nil.
func (o *IncidentImpactAttributes) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil.
func (o *IncidentImpactAttributes) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImpactAttributes) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImpactAttributes) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentImpactAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *IncidentImpactAttributes) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetImpactType returns the ImpactType field value if set, zero value otherwise.
func (o *IncidentImpactAttributes) GetImpactType() string {
	if o == nil || o.ImpactType == nil {
		var ret string
		return ret
	}
	return *o.ImpactType
}

// GetImpactTypeOk returns a tuple with the ImpactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImpactAttributes) GetImpactTypeOk() (*string, bool) {
	if o == nil || o.ImpactType == nil {
		return nil, false
	}
	return o.ImpactType, true
}

// HasImpactType returns a boolean if a field has been set.
func (o *IncidentImpactAttributes) HasImpactType() bool {
	return o != nil && o.ImpactType != nil
}

// SetImpactType gets a reference to the given string and assigns it to the ImpactType field.
func (o *IncidentImpactAttributes) SetImpactType(v string) {
	o.ImpactType = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentImpactAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImpactAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentImpactAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentImpactAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetStartAt returns the StartAt field value.
func (o *IncidentImpactAttributes) GetStartAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *IncidentImpactAttributes) GetStartAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value.
func (o *IncidentImpactAttributes) SetStartAt(v time.Time) {
	o.StartAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImpactAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["description"] = o.Description
	if o.EndAt.IsSet() {
		toSerialize["end_at"] = o.EndAt.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.ImpactType != nil {
		toSerialize["impact_type"] = o.ImpactType
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.StartAt.Nanosecond() == 0 {
		toSerialize["start_at"] = o.StartAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start_at"] = o.StartAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentImpactAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created     *time.Time             `json:"created,omitempty"`
		Description *string                `json:"description"`
		EndAt       datadog.NullableTime   `json:"end_at,omitempty"`
		Fields      map[string]interface{} `json:"fields,omitempty"`
		ImpactType  *string                `json:"impact_type,omitempty"`
		Modified    *time.Time             `json:"modified,omitempty"`
		StartAt     *time.Time             `json:"start_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.StartAt == nil {
		return fmt.Errorf("required field start_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "description", "end_at", "fields", "impact_type", "modified", "start_at"})
	} else {
		return err
	}
	o.Created = all.Created
	o.Description = *all.Description
	o.EndAt = all.EndAt
	o.Fields = all.Fields
	o.ImpactType = all.ImpactType
	o.Modified = all.Modified
	o.StartAt = *all.StartAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
