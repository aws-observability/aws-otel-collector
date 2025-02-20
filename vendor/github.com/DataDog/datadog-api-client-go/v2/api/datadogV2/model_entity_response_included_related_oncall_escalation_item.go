// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedOncallEscalationItem Oncall escalation.
type EntityResponseIncludedRelatedOncallEscalationItem struct {
	// Oncall email.
	Email *string `json:"email,omitempty"`
	// Oncall level.
	EscalationLevel *int64 `json:"escalationLevel,omitempty"`
	// Oncall name.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRelatedOncallEscalationItem instantiates a new EntityResponseIncludedRelatedOncallEscalationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRelatedOncallEscalationItem() *EntityResponseIncludedRelatedOncallEscalationItem {
	this := EntityResponseIncludedRelatedOncallEscalationItem{}
	return &this
}

// NewEntityResponseIncludedRelatedOncallEscalationItemWithDefaults instantiates a new EntityResponseIncludedRelatedOncallEscalationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRelatedOncallEscalationItemWithDefaults() *EntityResponseIncludedRelatedOncallEscalationItem {
	this := EntityResponseIncludedRelatedOncallEscalationItem{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) SetEmail(v string) {
	o.Email = &v
}

// GetEscalationLevel returns the EscalationLevel field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) GetEscalationLevel() int64 {
	if o == nil || o.EscalationLevel == nil {
		var ret int64
		return ret
	}
	return *o.EscalationLevel
}

// GetEscalationLevelOk returns a tuple with the EscalationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) GetEscalationLevelOk() (*int64, bool) {
	if o == nil || o.EscalationLevel == nil {
		return nil, false
	}
	return o.EscalationLevel, true
}

// HasEscalationLevel returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) HasEscalationLevel() bool {
	return o != nil && o.EscalationLevel != nil
}

// SetEscalationLevel gets a reference to the given int64 and assigns it to the EscalationLevel field.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) SetEscalationLevel(v int64) {
	o.EscalationLevel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRelatedOncallEscalationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EscalationLevel != nil {
		toSerialize["escalationLevel"] = o.EscalationLevel
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedRelatedOncallEscalationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email           *string `json:"email,omitempty"`
		EscalationLevel *int64  `json:"escalationLevel,omitempty"`
		Name            *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "escalationLevel", "name"})
	} else {
		return err
	}
	o.Email = all.Email
	o.EscalationLevel = all.EscalationLevel
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
