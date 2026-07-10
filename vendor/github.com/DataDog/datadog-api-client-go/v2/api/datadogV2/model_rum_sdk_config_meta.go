// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigMeta Metadata associated with a RUM SDK configuration.
type RumSdkConfigMeta struct {
	// The timestamp of the last update to this configuration.
	UpdatedAt time.Time `json:"updated_at"`
	// The handle of the user who last updated this configuration.
	UpdatedBy string `json:"updated_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSdkConfigMeta instantiates a new RumSdkConfigMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSdkConfigMeta(updatedAt time.Time, updatedBy string) *RumSdkConfigMeta {
	this := RumSdkConfigMeta{}
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewRumSdkConfigMetaWithDefaults instantiates a new RumSdkConfigMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSdkConfigMetaWithDefaults() *RumSdkConfigMeta {
	this := RumSdkConfigMeta{}
	return &this
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *RumSdkConfigMeta) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigMeta) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *RumSdkConfigMeta) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *RumSdkConfigMeta) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigMeta) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *RumSdkConfigMeta) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSdkConfigMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["updated_by"] = o.UpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSdkConfigMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UpdatedAt *time.Time `json:"updated_at"`
		UpdatedBy *string    `json:"updated_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"updated_at", "updated_by"})
	} else {
		return err
	}
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
