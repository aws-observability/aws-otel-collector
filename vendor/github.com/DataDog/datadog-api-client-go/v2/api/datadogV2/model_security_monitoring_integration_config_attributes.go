// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigAttributes The attributes of an entity context sync configuration as returned by the API.
type SecurityMonitoringIntegrationConfigAttributes struct {
	// The time at which the entity context sync configuration was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The domain associated with the external entity source (for example, the customer's identity provider domain).
	Domain string `json:"domain"`
	// Whether the sync is enabled and actively ingesting entities into Cloud SIEM.
	Enabled bool `json:"enabled"`
	// The type of external source that provides entities to Cloud SIEM.
	IntegrationType SecurityMonitoringIntegrationType `json:"integration_type"`
	// The time at which the entity context sync configuration was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The display name of the entity context sync configuration.
	Name *string `json:"name,omitempty"`
	// Free-form, non-sensitive settings for the entity context sync. The accepted keys depend on the source type.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// The state of the credentials configured on the entity context sync.
	State *SecurityMonitoringIntegrationConfigState `json:"state,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringIntegrationConfigAttributes instantiates a new SecurityMonitoringIntegrationConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringIntegrationConfigAttributes(domain string, enabled bool, integrationType SecurityMonitoringIntegrationType) *SecurityMonitoringIntegrationConfigAttributes {
	this := SecurityMonitoringIntegrationConfigAttributes{}
	this.Domain = domain
	this.Enabled = enabled
	this.IntegrationType = integrationType
	return &this
}

// NewSecurityMonitoringIntegrationConfigAttributesWithDefaults instantiates a new SecurityMonitoringIntegrationConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringIntegrationConfigAttributesWithDefaults() *SecurityMonitoringIntegrationConfigAttributes {
	this := SecurityMonitoringIntegrationConfigAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDomain returns the Domain field value.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetEnabled returns the Enabled field value.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetIntegrationType() SecurityMonitoringIntegrationType {
	if o == nil {
		var ret SecurityMonitoringIntegrationType
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetIntegrationTypeOk() (*SecurityMonitoringIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetIntegrationType(v SecurityMonitoringIntegrationType) {
	o.IntegrationType = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetState() SecurityMonitoringIntegrationConfigState {
	if o == nil || o.State == nil {
		var ret SecurityMonitoringIntegrationConfigState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) GetStateOk() (*SecurityMonitoringIntegrationConfigState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SecurityMonitoringIntegrationConfigAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given SecurityMonitoringIntegrationConfigState and assigns it to the State field.
func (o *SecurityMonitoringIntegrationConfigAttributes) SetState(v SecurityMonitoringIntegrationConfigState) {
	o.State = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringIntegrationConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["domain"] = o.Domain
	toSerialize["enabled"] = o.Enabled
	toSerialize["integration_type"] = o.IntegrationType
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringIntegrationConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *time.Time                                `json:"created_at,omitempty"`
		Domain          *string                                   `json:"domain"`
		Enabled         *bool                                     `json:"enabled"`
		IntegrationType *SecurityMonitoringIntegrationType        `json:"integration_type"`
		ModifiedAt      *time.Time                                `json:"modified_at,omitempty"`
		Name            *string                                   `json:"name,omitempty"`
		Settings        map[string]interface{}                    `json:"settings,omitempty"`
		State           *SecurityMonitoringIntegrationConfigState `json:"state,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "domain", "enabled", "integration_type", "modified_at", "name", "settings", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Domain = *all.Domain
	o.Enabled = *all.Enabled
	if !all.IntegrationType.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationType = *all.IntegrationType
	}
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.Settings = all.Settings
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
