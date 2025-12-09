// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceAttributesInterfaceStatuses Count of the device interfaces by status
type DeviceAttributesInterfaceStatuses struct {
	// The number of interfaces that are down
	Down *int64 `json:"down,omitempty"`
	// The number of interfaces that are off
	Off *int64 `json:"off,omitempty"`
	// The number of interfaces that are up
	Up *int64 `json:"up,omitempty"`
	// The number of interfaces that are in a warning state
	Warning *int64 `json:"warning,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeviceAttributesInterfaceStatuses instantiates a new DeviceAttributesInterfaceStatuses object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceAttributesInterfaceStatuses() *DeviceAttributesInterfaceStatuses {
	this := DeviceAttributesInterfaceStatuses{}
	return &this
}

// NewDeviceAttributesInterfaceStatusesWithDefaults instantiates a new DeviceAttributesInterfaceStatuses object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceAttributesInterfaceStatusesWithDefaults() *DeviceAttributesInterfaceStatuses {
	this := DeviceAttributesInterfaceStatuses{}
	return &this
}

// GetDown returns the Down field value if set, zero value otherwise.
func (o *DeviceAttributesInterfaceStatuses) GetDown() int64 {
	if o == nil || o.Down == nil {
		var ret int64
		return ret
	}
	return *o.Down
}

// GetDownOk returns a tuple with the Down field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributesInterfaceStatuses) GetDownOk() (*int64, bool) {
	if o == nil || o.Down == nil {
		return nil, false
	}
	return o.Down, true
}

// HasDown returns a boolean if a field has been set.
func (o *DeviceAttributesInterfaceStatuses) HasDown() bool {
	return o != nil && o.Down != nil
}

// SetDown gets a reference to the given int64 and assigns it to the Down field.
func (o *DeviceAttributesInterfaceStatuses) SetDown(v int64) {
	o.Down = &v
}

// GetOff returns the Off field value if set, zero value otherwise.
func (o *DeviceAttributesInterfaceStatuses) GetOff() int64 {
	if o == nil || o.Off == nil {
		var ret int64
		return ret
	}
	return *o.Off
}

// GetOffOk returns a tuple with the Off field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributesInterfaceStatuses) GetOffOk() (*int64, bool) {
	if o == nil || o.Off == nil {
		return nil, false
	}
	return o.Off, true
}

// HasOff returns a boolean if a field has been set.
func (o *DeviceAttributesInterfaceStatuses) HasOff() bool {
	return o != nil && o.Off != nil
}

// SetOff gets a reference to the given int64 and assigns it to the Off field.
func (o *DeviceAttributesInterfaceStatuses) SetOff(v int64) {
	o.Off = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *DeviceAttributesInterfaceStatuses) GetUp() int64 {
	if o == nil || o.Up == nil {
		var ret int64
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributesInterfaceStatuses) GetUpOk() (*int64, bool) {
	if o == nil || o.Up == nil {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *DeviceAttributesInterfaceStatuses) HasUp() bool {
	return o != nil && o.Up != nil
}

// SetUp gets a reference to the given int64 and assigns it to the Up field.
func (o *DeviceAttributesInterfaceStatuses) SetUp(v int64) {
	o.Up = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *DeviceAttributesInterfaceStatuses) GetWarning() int64 {
	if o == nil || o.Warning == nil {
		var ret int64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributesInterfaceStatuses) GetWarningOk() (*int64, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *DeviceAttributesInterfaceStatuses) HasWarning() bool {
	return o != nil && o.Warning != nil
}

// SetWarning gets a reference to the given int64 and assigns it to the Warning field.
func (o *DeviceAttributesInterfaceStatuses) SetWarning(v int64) {
	o.Warning = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceAttributesInterfaceStatuses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Down != nil {
		toSerialize["down"] = o.Down
	}
	if o.Off != nil {
		toSerialize["off"] = o.Off
	}
	if o.Up != nil {
		toSerialize["up"] = o.Up
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeviceAttributesInterfaceStatuses) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Down    *int64 `json:"down,omitempty"`
		Off     *int64 `json:"off,omitempty"`
		Up      *int64 `json:"up,omitempty"`
		Warning *int64 `json:"warning,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"down", "off", "up", "warning"})
	} else {
		return err
	}
	o.Down = all.Down
	o.Off = all.Off
	o.Up = all.Up
	o.Warning = all.Warning

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
