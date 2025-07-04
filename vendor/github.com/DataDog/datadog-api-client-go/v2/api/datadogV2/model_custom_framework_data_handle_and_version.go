// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomFrameworkDataHandleAndVersion Framework Handle and Version.
type CustomFrameworkDataHandleAndVersion struct {
	// Framework Handle
	Handle *string `json:"handle,omitempty"`
	// Framework Version
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomFrameworkDataHandleAndVersion instantiates a new CustomFrameworkDataHandleAndVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomFrameworkDataHandleAndVersion() *CustomFrameworkDataHandleAndVersion {
	this := CustomFrameworkDataHandleAndVersion{}
	return &this
}

// NewCustomFrameworkDataHandleAndVersionWithDefaults instantiates a new CustomFrameworkDataHandleAndVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomFrameworkDataHandleAndVersionWithDefaults() *CustomFrameworkDataHandleAndVersion {
	this := CustomFrameworkDataHandleAndVersion{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *CustomFrameworkDataHandleAndVersion) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataHandleAndVersion) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *CustomFrameworkDataHandleAndVersion) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *CustomFrameworkDataHandleAndVersion) SetHandle(v string) {
	o.Handle = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CustomFrameworkDataHandleAndVersion) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFrameworkDataHandleAndVersion) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CustomFrameworkDataHandleAndVersion) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CustomFrameworkDataHandleAndVersion) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomFrameworkDataHandleAndVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomFrameworkDataHandleAndVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle  *string `json:"handle,omitempty"`
		Version *string `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "version"})
	} else {
		return err
	}
	o.Handle = all.Handle
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
