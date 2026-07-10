// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceWindowAttributes Attributes of a maintenance window, including its schedule and the query that determines which cases are affected.
type MaintenanceWindowAttributes struct {
	// The UUID of the user who created this maintenance window. Read-only.
	CreatedBy *string `json:"created_by,omitempty"`
	// The ISO 8601 timestamp when the maintenance window ends and normal notification behavior resumes.
	EndAt time.Time `json:"end_at"`
	// A human-readable name for the maintenance window (for example, `Database migration - Dec 15`).
	Name string `json:"name"`
	// A case search query that determines which cases are affected during the maintenance window. Uses the same syntax as the Case Management search bar.
	Query string `json:"query"`
	// The ISO 8601 timestamp when the maintenance window begins and notifications start being suppressed.
	StartAt time.Time `json:"start_at"`
	// The UUID of the user who last modified this maintenance window. Read-only.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMaintenanceWindowAttributes instantiates a new MaintenanceWindowAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMaintenanceWindowAttributes(endAt time.Time, name string, query string, startAt time.Time) *MaintenanceWindowAttributes {
	this := MaintenanceWindowAttributes{}
	this.EndAt = endAt
	this.Name = name
	this.Query = query
	this.StartAt = startAt
	return &this
}

// NewMaintenanceWindowAttributesWithDefaults instantiates a new MaintenanceWindowAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMaintenanceWindowAttributesWithDefaults() *MaintenanceWindowAttributes {
	this := MaintenanceWindowAttributes{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *MaintenanceWindowAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MaintenanceWindowAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *MaintenanceWindowAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEndAt returns the EndAt field value.
func (o *MaintenanceWindowAttributes) GetEndAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowAttributes) GetEndAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value.
func (o *MaintenanceWindowAttributes) SetEndAt(v time.Time) {
	o.EndAt = v
}

// GetName returns the Name field value.
func (o *MaintenanceWindowAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MaintenanceWindowAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *MaintenanceWindowAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *MaintenanceWindowAttributes) SetQuery(v string) {
	o.Query = v
}

// GetStartAt returns the StartAt field value.
func (o *MaintenanceWindowAttributes) GetStartAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowAttributes) GetStartAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value.
func (o *MaintenanceWindowAttributes) SetStartAt(v time.Time) {
	o.StartAt = v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *MaintenanceWindowAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *MaintenanceWindowAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *MaintenanceWindowAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MaintenanceWindowAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
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
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MaintenanceWindowAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy *string    `json:"created_by,omitempty"`
		EndAt     *time.Time `json:"end_at"`
		Name      *string    `json:"name"`
		Query     *string    `json:"query"`
		StartAt   *time.Time `json:"start_at"`
		UpdatedBy *string    `json:"updated_by,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "end_at", "name", "query", "start_at", "updated_by"})
	} else {
		return err
	}
	o.CreatedBy = all.CreatedBy
	o.EndAt = *all.EndAt
	o.Name = *all.Name
	o.Query = *all.Query
	o.StartAt = *all.StartAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
