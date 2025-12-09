// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScannedAssetMetadataLastSuccess Metadata for the last successful scan of an asset.
type ScannedAssetMetadataLastSuccess struct {
	// The environment of the last success scan of the asset.
	Env *string `json:"env,omitempty"`
	// The list of origins of the last success scan of the asset.
	Origin []string `json:"origin,omitempty"`
	// The timestamp of the last success scan of the asset.
	Timestamp string `json:"timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScannedAssetMetadataLastSuccess instantiates a new ScannedAssetMetadataLastSuccess object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScannedAssetMetadataLastSuccess(timestamp string) *ScannedAssetMetadataLastSuccess {
	this := ScannedAssetMetadataLastSuccess{}
	this.Timestamp = timestamp
	return &this
}

// NewScannedAssetMetadataLastSuccessWithDefaults instantiates a new ScannedAssetMetadataLastSuccess object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScannedAssetMetadataLastSuccessWithDefaults() *ScannedAssetMetadataLastSuccess {
	this := ScannedAssetMetadataLastSuccess{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ScannedAssetMetadataLastSuccess) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadataLastSuccess) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ScannedAssetMetadataLastSuccess) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *ScannedAssetMetadataLastSuccess) SetEnv(v string) {
	o.Env = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ScannedAssetMetadataLastSuccess) GetOrigin() []string {
	if o == nil || o.Origin == nil {
		var ret []string
		return ret
	}
	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadataLastSuccess) GetOriginOk() (*[]string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return &o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ScannedAssetMetadataLastSuccess) HasOrigin() bool {
	return o != nil && o.Origin != nil
}

// SetOrigin gets a reference to the given []string and assigns it to the Origin field.
func (o *ScannedAssetMetadataLastSuccess) SetOrigin(v []string) {
	o.Origin = v
}

// GetTimestamp returns the Timestamp field value.
func (o *ScannedAssetMetadataLastSuccess) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadataLastSuccess) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *ScannedAssetMetadataLastSuccess) SetTimestamp(v string) {
	o.Timestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScannedAssetMetadataLastSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScannedAssetMetadataLastSuccess) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env       *string  `json:"env,omitempty"`
		Origin    []string `json:"origin,omitempty"`
		Timestamp *string  `json:"timestamp"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "origin", "timestamp"})
	} else {
		return err
	}
	o.Env = all.Env
	o.Origin = all.Origin
	o.Timestamp = *all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
