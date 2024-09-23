// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageIncidentManagementHour Incident management usage for a given organization for a given hour.
type UsageIncidentManagementHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the total number monthly active users from the start of the given hour's month until the given hour.
	MonthlyActiveUsers datadog.NullableInt64 `json:"monthly_active_users,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageIncidentManagementHour instantiates a new UsageIncidentManagementHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageIncidentManagementHour() *UsageIncidentManagementHour {
	this := UsageIncidentManagementHour{}
	return &this
}

// NewUsageIncidentManagementHourWithDefaults instantiates a new UsageIncidentManagementHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageIncidentManagementHourWithDefaults() *UsageIncidentManagementHour {
	this := UsageIncidentManagementHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageIncidentManagementHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIncidentManagementHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageIncidentManagementHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageIncidentManagementHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetMonthlyActiveUsers returns the MonthlyActiveUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageIncidentManagementHour) GetMonthlyActiveUsers() int64 {
	if o == nil || o.MonthlyActiveUsers.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MonthlyActiveUsers.Get()
}

// GetMonthlyActiveUsersOk returns a tuple with the MonthlyActiveUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageIncidentManagementHour) GetMonthlyActiveUsersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyActiveUsers.Get(), o.MonthlyActiveUsers.IsSet()
}

// HasMonthlyActiveUsers returns a boolean if a field has been set.
func (o *UsageIncidentManagementHour) HasMonthlyActiveUsers() bool {
	return o != nil && o.MonthlyActiveUsers.IsSet()
}

// SetMonthlyActiveUsers gets a reference to the given datadog.NullableInt64 and assigns it to the MonthlyActiveUsers field.
func (o *UsageIncidentManagementHour) SetMonthlyActiveUsers(v int64) {
	o.MonthlyActiveUsers.Set(&v)
}

// SetMonthlyActiveUsersNil sets the value for MonthlyActiveUsers to be an explicit nil.
func (o *UsageIncidentManagementHour) SetMonthlyActiveUsersNil() {
	o.MonthlyActiveUsers.Set(nil)
}

// UnsetMonthlyActiveUsers ensures that no value is present for MonthlyActiveUsers, not even an explicit nil.
func (o *UsageIncidentManagementHour) UnsetMonthlyActiveUsers() {
	o.MonthlyActiveUsers.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageIncidentManagementHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIncidentManagementHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageIncidentManagementHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageIncidentManagementHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageIncidentManagementHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIncidentManagementHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageIncidentManagementHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageIncidentManagementHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageIncidentManagementHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MonthlyActiveUsers.IsSet() {
		toSerialize["monthly_active_users"] = o.MonthlyActiveUsers.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageIncidentManagementHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hour               *time.Time            `json:"hour,omitempty"`
		MonthlyActiveUsers datadog.NullableInt64 `json:"monthly_active_users,omitempty"`
		OrgName            *string               `json:"org_name,omitempty"`
		PublicId           *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hour", "monthly_active_users", "org_name", "public_id"})
	} else {
		return err
	}
	o.Hour = all.Hour
	o.MonthlyActiveUsers = all.MonthlyActiveUsers
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
