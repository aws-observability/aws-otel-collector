// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PartialAPIKeyAttributes Attributes of a partial API key.
type PartialAPIKeyAttributes struct {
	// The category of the API key.
	Category *string `json:"category,omitempty"`
	// Creation date of the API key.
	CreatedAt *string `json:"created_at,omitempty"`
	// Date the API Key was last used.
	DateLastUsed datadog.NullableTime `json:"date_last_used,omitempty"`
	// The last four characters of the API key.
	Last4 *string `json:"last4,omitempty"`
	// Date the API key was last modified.
	ModifiedAt *string `json:"modified_at,omitempty"`
	// Name of the API key.
	Name *string `json:"name,omitempty"`
	// The remote config read enabled status.
	RemoteConfigReadEnabled *bool `json:"remote_config_read_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPartialAPIKeyAttributes instantiates a new PartialAPIKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPartialAPIKeyAttributes() *PartialAPIKeyAttributes {
	this := PartialAPIKeyAttributes{}
	return &this
}

// NewPartialAPIKeyAttributesWithDefaults instantiates a new PartialAPIKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPartialAPIKeyAttributesWithDefaults() *PartialAPIKeyAttributes {
	this := PartialAPIKeyAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *PartialAPIKeyAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialAPIKeyAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *PartialAPIKeyAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PartialAPIKeyAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialAPIKeyAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PartialAPIKeyAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDateLastUsed returns the DateLastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAPIKeyAttributes) GetDateLastUsed() time.Time {
	if o == nil || o.DateLastUsed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DateLastUsed.Get()
}

// GetDateLastUsedOk returns a tuple with the DateLastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PartialAPIKeyAttributes) GetDateLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateLastUsed.Get(), o.DateLastUsed.IsSet()
}

// HasDateLastUsed returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasDateLastUsed() bool {
	return o != nil && o.DateLastUsed.IsSet()
}

// SetDateLastUsed gets a reference to the given datadog.NullableTime and assigns it to the DateLastUsed field.
func (o *PartialAPIKeyAttributes) SetDateLastUsed(v time.Time) {
	o.DateLastUsed.Set(&v)
}

// SetDateLastUsedNil sets the value for DateLastUsed to be an explicit nil.
func (o *PartialAPIKeyAttributes) SetDateLastUsedNil() {
	o.DateLastUsed.Set(nil)
}

// UnsetDateLastUsed ensures that no value is present for DateLastUsed, not even an explicit nil.
func (o *PartialAPIKeyAttributes) UnsetDateLastUsed() {
	o.DateLastUsed.Unset()
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PartialAPIKeyAttributes) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialAPIKeyAttributes) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasLast4() bool {
	return o != nil && o.Last4 != nil
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PartialAPIKeyAttributes) SetLast4(v string) {
	o.Last4 = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *PartialAPIKeyAttributes) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialAPIKeyAttributes) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *PartialAPIKeyAttributes) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialAPIKeyAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialAPIKeyAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialAPIKeyAttributes) SetName(v string) {
	o.Name = &v
}

// GetRemoteConfigReadEnabled returns the RemoteConfigReadEnabled field value if set, zero value otherwise.
func (o *PartialAPIKeyAttributes) GetRemoteConfigReadEnabled() bool {
	if o == nil || o.RemoteConfigReadEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteConfigReadEnabled
}

// GetRemoteConfigReadEnabledOk returns a tuple with the RemoteConfigReadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialAPIKeyAttributes) GetRemoteConfigReadEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteConfigReadEnabled == nil {
		return nil, false
	}
	return o.RemoteConfigReadEnabled, true
}

// HasRemoteConfigReadEnabled returns a boolean if a field has been set.
func (o *PartialAPIKeyAttributes) HasRemoteConfigReadEnabled() bool {
	return o != nil && o.RemoteConfigReadEnabled != nil
}

// SetRemoteConfigReadEnabled gets a reference to the given bool and assigns it to the RemoteConfigReadEnabled field.
func (o *PartialAPIKeyAttributes) SetRemoteConfigReadEnabled(v bool) {
	o.RemoteConfigReadEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PartialAPIKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DateLastUsed.IsSet() {
		toSerialize["date_last_used"] = o.DateLastUsed.Get()
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RemoteConfigReadEnabled != nil {
		toSerialize["remote_config_read_enabled"] = o.RemoteConfigReadEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PartialAPIKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category                *string              `json:"category,omitempty"`
		CreatedAt               *string              `json:"created_at,omitempty"`
		DateLastUsed            datadog.NullableTime `json:"date_last_used,omitempty"`
		Last4                   *string              `json:"last4,omitempty"`
		ModifiedAt              *string              `json:"modified_at,omitempty"`
		Name                    *string              `json:"name,omitempty"`
		RemoteConfigReadEnabled *bool                `json:"remote_config_read_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "created_at", "date_last_used", "last4", "modified_at", "name", "remote_config_read_enabled"})
	} else {
		return err
	}
	o.Category = all.Category
	o.CreatedAt = all.CreatedAt
	o.DateLastUsed = all.DateLastUsed
	o.Last4 = all.Last4
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.RemoteConfigReadEnabled = all.RemoteConfigReadEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
