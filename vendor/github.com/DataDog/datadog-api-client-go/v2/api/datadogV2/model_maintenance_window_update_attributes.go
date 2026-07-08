// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceWindowUpdateAttributes Attributes that can be updated on a maintenance window. All fields are optional; only provided fields are changed.
type MaintenanceWindowUpdateAttributes struct {
	// The end time of the maintenance window.
	EndAt *time.Time `json:"end_at,omitempty"`
	// The name of the maintenance window.
	Name *string `json:"name,omitempty"`
	// The query to filter event management cases for this maintenance window.
	Query *string `json:"query,omitempty"`
	// The start time of the maintenance window.
	StartAt *time.Time `json:"start_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceWindowUpdateAttributes instantiates a new MaintenanceWindowUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceWindowUpdateAttributes() *MaintenanceWindowUpdateAttributes {
	this := MaintenanceWindowUpdateAttributes{}
	return &this
}

// NewMaintenanceWindowUpdateAttributesWithDefaults instantiates a new MaintenanceWindowUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceWindowUpdateAttributesWithDefaults() *MaintenanceWindowUpdateAttributes {
	this := MaintenanceWindowUpdateAttributes{}
	return &this
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *MaintenanceWindowUpdateAttributes) GetEndAt() time.Time {
	if o == nil || o.EndAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowUpdateAttributes) GetEndAtOk() (*time.Time, bool) {
	if o == nil || o.EndAt == nil {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *MaintenanceWindowUpdateAttributes) HasEndAt() bool {
	return o != nil && o.EndAt != nil
}

// SetEndAt gets a reference to the given time.Time and assigns it to the EndAt field.
func (o *MaintenanceWindowUpdateAttributes) SetEndAt(v time.Time) {
	o.EndAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MaintenanceWindowUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MaintenanceWindowUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MaintenanceWindowUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MaintenanceWindowUpdateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowUpdateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MaintenanceWindowUpdateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MaintenanceWindowUpdateAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *MaintenanceWindowUpdateAttributes) GetStartAt() time.Time {
	if o == nil || o.StartAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowUpdateAttributes) GetStartAtOk() (*time.Time, bool) {
	if o == nil || o.StartAt == nil {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *MaintenanceWindowUpdateAttributes) HasStartAt() bool {
	return o != nil && o.StartAt != nil
}

// SetStartAt gets a reference to the given time.Time and assigns it to the StartAt field.
func (o *MaintenanceWindowUpdateAttributes) SetStartAt(v time.Time) {
	o.StartAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceWindowUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EndAt != nil {
		if o.EndAt.Nanosecond() == 0 {
			toSerialize["end_at"] = o.EndAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_at"] = o.EndAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.StartAt != nil {
		if o.StartAt.Nanosecond() == 0 {
			toSerialize["start_at"] = o.StartAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_at"] = o.StartAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaintenanceWindowUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndAt   *time.Time `json:"end_at,omitempty"`
		Name    *string    `json:"name,omitempty"`
		Query   *string    `json:"query,omitempty"`
		StartAt *time.Time `json:"start_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_at", "name", "query", "start_at"})
	} else {
		return err
	}
	o.EndAt = all.EndAt
	o.Name = all.Name
	o.Query = all.Query
	o.StartAt = all.StartAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
