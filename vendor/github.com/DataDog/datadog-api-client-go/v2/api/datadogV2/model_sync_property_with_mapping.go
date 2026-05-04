// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyncPropertyWithMapping Sync property with mapping configuration.
type SyncPropertyWithMapping struct {
	// Map of source values to destination values for synchronization.
	Mapping map[string]string `json:"mapping,omitempty"`
	// Map of source names to display names used during synchronization.
	NameMapping map[string]string `json:"name_mapping,omitempty"`
	// The direction and type of synchronization for this property.
	SyncType *string `json:"sync_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyncPropertyWithMapping instantiates a new SyncPropertyWithMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyncPropertyWithMapping() *SyncPropertyWithMapping {
	this := SyncPropertyWithMapping{}
	return &this
}

// NewSyncPropertyWithMappingWithDefaults instantiates a new SyncPropertyWithMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyncPropertyWithMappingWithDefaults() *SyncPropertyWithMapping {
	this := SyncPropertyWithMapping{}
	return &this
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *SyncPropertyWithMapping) GetMapping() map[string]string {
	if o == nil || o.Mapping == nil {
		var ret map[string]string
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPropertyWithMapping) GetMappingOk() (*map[string]string, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return &o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *SyncPropertyWithMapping) HasMapping() bool {
	return o != nil && o.Mapping != nil
}

// SetMapping gets a reference to the given map[string]string and assigns it to the Mapping field.
func (o *SyncPropertyWithMapping) SetMapping(v map[string]string) {
	o.Mapping = v
}

// GetNameMapping returns the NameMapping field value if set, zero value otherwise.
func (o *SyncPropertyWithMapping) GetNameMapping() map[string]string {
	if o == nil || o.NameMapping == nil {
		var ret map[string]string
		return ret
	}
	return o.NameMapping
}

// GetNameMappingOk returns a tuple with the NameMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPropertyWithMapping) GetNameMappingOk() (*map[string]string, bool) {
	if o == nil || o.NameMapping == nil {
		return nil, false
	}
	return &o.NameMapping, true
}

// HasNameMapping returns a boolean if a field has been set.
func (o *SyncPropertyWithMapping) HasNameMapping() bool {
	return o != nil && o.NameMapping != nil
}

// SetNameMapping gets a reference to the given map[string]string and assigns it to the NameMapping field.
func (o *SyncPropertyWithMapping) SetNameMapping(v map[string]string) {
	o.NameMapping = v
}

// GetSyncType returns the SyncType field value if set, zero value otherwise.
func (o *SyncPropertyWithMapping) GetSyncType() string {
	if o == nil || o.SyncType == nil {
		var ret string
		return ret
	}
	return *o.SyncType
}

// GetSyncTypeOk returns a tuple with the SyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPropertyWithMapping) GetSyncTypeOk() (*string, bool) {
	if o == nil || o.SyncType == nil {
		return nil, false
	}
	return o.SyncType, true
}

// HasSyncType returns a boolean if a field has been set.
func (o *SyncPropertyWithMapping) HasSyncType() bool {
	return o != nil && o.SyncType != nil
}

// SetSyncType gets a reference to the given string and assigns it to the SyncType field.
func (o *SyncPropertyWithMapping) SetSyncType(v string) {
	o.SyncType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyncPropertyWithMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Mapping != nil {
		toSerialize["mapping"] = o.Mapping
	}
	if o.NameMapping != nil {
		toSerialize["name_mapping"] = o.NameMapping
	}
	if o.SyncType != nil {
		toSerialize["sync_type"] = o.SyncType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyncPropertyWithMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mapping     map[string]string `json:"mapping,omitempty"`
		NameMapping map[string]string `json:"name_mapping,omitempty"`
		SyncType    *string           `json:"sync_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mapping", "name_mapping", "sync_type"})
	} else {
		return err
	}
	o.Mapping = all.Mapping
	o.NameMapping = all.NameMapping
	o.SyncType = all.SyncType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
