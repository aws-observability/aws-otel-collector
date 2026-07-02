// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceWindowCreateAttributes Attributes required to create a maintenance window.
type MaintenanceWindowCreateAttributes struct {
	// The end time of the maintenance window.
	EndAt time.Time `json:"end_at"`
	// The name of the maintenance window.
	Name string `json:"name"`
	// The query to filter event management cases for this maintenance window.
	Query string `json:"query"`
	// The start time of the maintenance window.
	StartAt time.Time `json:"start_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceWindowCreateAttributes instantiates a new MaintenanceWindowCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceWindowCreateAttributes(endAt time.Time, name string, query string, startAt time.Time) *MaintenanceWindowCreateAttributes {
	this := MaintenanceWindowCreateAttributes{}
	this.EndAt = endAt
	this.Name = name
	this.Query = query
	this.StartAt = startAt
	return &this
}

// NewMaintenanceWindowCreateAttributesWithDefaults instantiates a new MaintenanceWindowCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceWindowCreateAttributesWithDefaults() *MaintenanceWindowCreateAttributes {
	this := MaintenanceWindowCreateAttributes{}
	return &this
}

// GetEndAt returns the EndAt field value.
func (o *MaintenanceWindowCreateAttributes) GetEndAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowCreateAttributes) GetEndAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value.
func (o *MaintenanceWindowCreateAttributes) SetEndAt(v time.Time) {
	o.EndAt = v
}

// GetName returns the Name field value.
func (o *MaintenanceWindowCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MaintenanceWindowCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *MaintenanceWindowCreateAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *MaintenanceWindowCreateAttributes) SetQuery(v string) {
	o.Query = v
}

// GetStartAt returns the StartAt field value.
func (o *MaintenanceWindowCreateAttributes) GetStartAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowCreateAttributes) GetStartAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value.
func (o *MaintenanceWindowCreateAttributes) SetStartAt(v time.Time) {
	o.StartAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceWindowCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EndAt.Nanosecond() == 0 {
		toSerialize["end_at"] = o.EndAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["end_at"] = o.EndAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
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
func (o *MaintenanceWindowCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndAt   *time.Time `json:"end_at"`
		Name    *string    `json:"name"`
		Query   *string    `json:"query"`
		StartAt *time.Time `json:"start_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EndAt == nil {
		return fmt.Errorf("required field end_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.StartAt == nil {
		return fmt.Errorf("required field start_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end_at", "name", "query", "start_at"})
	} else {
		return err
	}
	o.EndAt = *all.EndAt
	o.Name = *all.Name
	o.Query = *all.Query
	o.StartAt = *all.StartAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
