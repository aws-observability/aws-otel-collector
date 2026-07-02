// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitConfigAttributes The RUM rate limit configuration properties.
type RumRateLimitConfigAttributes struct {
	// The configuration used when `mode` is `adaptive`.
	Adaptive *RumRateLimitAdaptiveConfig `json:"adaptive,omitempty"`
	// The configuration used when `mode` is `custom`.
	Custom *RumRateLimitCustomConfig `json:"custom,omitempty"`
	// The rate limit mode. `custom` enforces a fixed session limit, while
	// `adaptive` dynamically adjusts retention.
	Mode RumRateLimitMode `json:"mode"`
	// The ID of the organization the rate limit configuration belongs to.
	OrgId int64 `json:"org_id"`
	// The date the rate limit configuration was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// The handle of the user who last updated the rate limit configuration.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRateLimitConfigAttributes instantiates a new RumRateLimitConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRateLimitConfigAttributes(mode RumRateLimitMode, orgId int64) *RumRateLimitConfigAttributes {
	this := RumRateLimitConfigAttributes{}
	this.Mode = mode
	this.OrgId = orgId
	return &this
}

// NewRumRateLimitConfigAttributesWithDefaults instantiates a new RumRateLimitConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRateLimitConfigAttributesWithDefaults() *RumRateLimitConfigAttributes {
	this := RumRateLimitConfigAttributes{}
	return &this
}

// GetAdaptive returns the Adaptive field value if set, zero value otherwise.
func (o *RumRateLimitConfigAttributes) GetAdaptive() RumRateLimitAdaptiveConfig {
	if o == nil || o.Adaptive == nil {
		var ret RumRateLimitAdaptiveConfig
		return ret
	}
	return *o.Adaptive
}

// GetAdaptiveOk returns a tuple with the Adaptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigAttributes) GetAdaptiveOk() (*RumRateLimitAdaptiveConfig, bool) {
	if o == nil || o.Adaptive == nil {
		return nil, false
	}
	return o.Adaptive, true
}

// HasAdaptive returns a boolean if a field has been set.
func (o *RumRateLimitConfigAttributes) HasAdaptive() bool {
	return o != nil && o.Adaptive != nil
}

// SetAdaptive gets a reference to the given RumRateLimitAdaptiveConfig and assigns it to the Adaptive field.
func (o *RumRateLimitConfigAttributes) SetAdaptive(v RumRateLimitAdaptiveConfig) {
	o.Adaptive = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *RumRateLimitConfigAttributes) GetCustom() RumRateLimitCustomConfig {
	if o == nil || o.Custom == nil {
		var ret RumRateLimitCustomConfig
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigAttributes) GetCustomOk() (*RumRateLimitCustomConfig, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *RumRateLimitConfigAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given RumRateLimitCustomConfig and assigns it to the Custom field.
func (o *RumRateLimitConfigAttributes) SetCustom(v RumRateLimitCustomConfig) {
	o.Custom = &v
}

// GetMode returns the Mode field value.
func (o *RumRateLimitConfigAttributes) GetMode() RumRateLimitMode {
	if o == nil {
		var ret RumRateLimitMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigAttributes) GetModeOk() (*RumRateLimitMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *RumRateLimitConfigAttributes) SetMode(v RumRateLimitMode) {
	o.Mode = v
}

// GetOrgId returns the OrgId field value.
func (o *RumRateLimitConfigAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *RumRateLimitConfigAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RumRateLimitConfigAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RumRateLimitConfigAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RumRateLimitConfigAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RumRateLimitConfigAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRateLimitConfigAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RumRateLimitConfigAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *RumRateLimitConfigAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRateLimitConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Adaptive != nil {
		toSerialize["adaptive"] = o.Adaptive
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	toSerialize["mode"] = o.Mode
	toSerialize["org_id"] = o.OrgId
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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
func (o *RumRateLimitConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Adaptive  *RumRateLimitAdaptiveConfig `json:"adaptive,omitempty"`
		Custom    *RumRateLimitCustomConfig   `json:"custom,omitempty"`
		Mode      *RumRateLimitMode           `json:"mode"`
		OrgId     *int64                      `json:"org_id"`
		UpdatedAt *string                     `json:"updated_at,omitempty"`
		UpdatedBy *string                     `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"adaptive", "custom", "mode", "org_id", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Adaptive != nil && all.Adaptive.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Adaptive = all.Adaptive
	if all.Custom != nil && all.Custom.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Custom = all.Custom
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.OrgId = *all.OrgId
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
