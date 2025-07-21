// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafExclusionFilterMetadata Extra information about the exclusion filter.
type ApplicationSecurityWafExclusionFilterMetadata struct {
	// The creation date of the exclusion filter.
	AddedAt *time.Time `json:"added_at,omitempty"`
	// The handle of the user who created the exclusion filter.
	AddedBy *string `json:"added_by,omitempty"`
	// The name of the user who created the exclusion filter.
	AddedByName *string `json:"added_by_name,omitempty"`
	// The last modification date of the exclusion filter.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The handle of the user who last modified the exclusion filter.
	ModifiedBy *string `json:"modified_by,omitempty"`
	// The name of the user who last modified the exclusion filter.
	ModifiedByName *string `json:"modified_by_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafExclusionFilterMetadata instantiates a new ApplicationSecurityWafExclusionFilterMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafExclusionFilterMetadata() *ApplicationSecurityWafExclusionFilterMetadata {
	this := ApplicationSecurityWafExclusionFilterMetadata{}
	return &this
}

// NewApplicationSecurityWafExclusionFilterMetadataWithDefaults instantiates a new ApplicationSecurityWafExclusionFilterMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafExclusionFilterMetadataWithDefaults() *ApplicationSecurityWafExclusionFilterMetadata {
	this := ApplicationSecurityWafExclusionFilterMetadata{}
	return &this
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetAddedAt() time.Time {
	if o == nil || o.AddedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || o.AddedAt == nil {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) HasAddedAt() bool {
	return o != nil && o.AddedAt != nil
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *ApplicationSecurityWafExclusionFilterMetadata) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetAddedBy() string {
	if o == nil || o.AddedBy == nil {
		var ret string
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetAddedByOk() (*string, bool) {
	if o == nil || o.AddedBy == nil {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) HasAddedBy() bool {
	return o != nil && o.AddedBy != nil
}

// SetAddedBy gets a reference to the given string and assigns it to the AddedBy field.
func (o *ApplicationSecurityWafExclusionFilterMetadata) SetAddedBy(v string) {
	o.AddedBy = &v
}

// GetAddedByName returns the AddedByName field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetAddedByName() string {
	if o == nil || o.AddedByName == nil {
		var ret string
		return ret
	}
	return *o.AddedByName
}

// GetAddedByNameOk returns a tuple with the AddedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetAddedByNameOk() (*string, bool) {
	if o == nil || o.AddedByName == nil {
		return nil, false
	}
	return o.AddedByName, true
}

// HasAddedByName returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) HasAddedByName() bool {
	return o != nil && o.AddedByName != nil
}

// SetAddedByName gets a reference to the given string and assigns it to the AddedByName field.
func (o *ApplicationSecurityWafExclusionFilterMetadata) SetAddedByName(v string) {
	o.AddedByName = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ApplicationSecurityWafExclusionFilterMetadata) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *ApplicationSecurityWafExclusionFilterMetadata) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedByName returns the ModifiedByName field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetModifiedByName() string {
	if o == nil || o.ModifiedByName == nil {
		var ret string
		return ret
	}
	return *o.ModifiedByName
}

// GetModifiedByNameOk returns a tuple with the ModifiedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) GetModifiedByNameOk() (*string, bool) {
	if o == nil || o.ModifiedByName == nil {
		return nil, false
	}
	return o.ModifiedByName, true
}

// HasModifiedByName returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterMetadata) HasModifiedByName() bool {
	return o != nil && o.ModifiedByName != nil
}

// SetModifiedByName gets a reference to the given string and assigns it to the ModifiedByName field.
func (o *ApplicationSecurityWafExclusionFilterMetadata) SetModifiedByName(v string) {
	o.ModifiedByName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafExclusionFilterMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddedAt != nil {
		if o.AddedAt.Nanosecond() == 0 {
			toSerialize["added_at"] = o.AddedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["added_at"] = o.AddedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.AddedBy != nil {
		toSerialize["added_by"] = o.AddedBy
	}
	if o.AddedByName != nil {
		toSerialize["added_by_name"] = o.AddedByName
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.ModifiedByName != nil {
		toSerialize["modified_by_name"] = o.ModifiedByName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafExclusionFilterMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddedAt        *time.Time `json:"added_at,omitempty"`
		AddedBy        *string    `json:"added_by,omitempty"`
		AddedByName    *string    `json:"added_by_name,omitempty"`
		ModifiedAt     *time.Time `json:"modified_at,omitempty"`
		ModifiedBy     *string    `json:"modified_by,omitempty"`
		ModifiedByName *string    `json:"modified_by_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"added_at", "added_by", "added_by_name", "modified_at", "modified_by", "modified_by_name"})
	} else {
		return err
	}
	o.AddedAt = all.AddedAt
	o.AddedBy = all.AddedBy
	o.AddedByName = all.AddedByName
	o.ModifiedAt = all.ModifiedAt
	o.ModifiedBy = all.ModifiedBy
	o.ModifiedByName = all.ModifiedByName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
