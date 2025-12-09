// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageFlavor Container Image breakdown by supported platform.
type ContainerImageFlavor struct {
	// Time the platform-specific Container Image was built.
	BuiltAt *string `json:"built_at,omitempty"`
	// Operating System architecture supported by the Container Image.
	OsArchitecture *string `json:"os_architecture,omitempty"`
	// Operating System name supported by the Container Image.
	OsName *string `json:"os_name,omitempty"`
	// Operating System version supported by the Container Image.
	OsVersion *string `json:"os_version,omitempty"`
	// Size of the platform-specific Container Image.
	Size *int64 `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerImageFlavor instantiates a new ContainerImageFlavor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageFlavor() *ContainerImageFlavor {
	this := ContainerImageFlavor{}
	return &this
}

// NewContainerImageFlavorWithDefaults instantiates a new ContainerImageFlavor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageFlavorWithDefaults() *ContainerImageFlavor {
	this := ContainerImageFlavor{}
	return &this
}

// GetBuiltAt returns the BuiltAt field value if set, zero value otherwise.
func (o *ContainerImageFlavor) GetBuiltAt() string {
	if o == nil || o.BuiltAt == nil {
		var ret string
		return ret
	}
	return *o.BuiltAt
}

// GetBuiltAtOk returns a tuple with the BuiltAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageFlavor) GetBuiltAtOk() (*string, bool) {
	if o == nil || o.BuiltAt == nil {
		return nil, false
	}
	return o.BuiltAt, true
}

// HasBuiltAt returns a boolean if a field has been set.
func (o *ContainerImageFlavor) HasBuiltAt() bool {
	return o != nil && o.BuiltAt != nil
}

// SetBuiltAt gets a reference to the given string and assigns it to the BuiltAt field.
func (o *ContainerImageFlavor) SetBuiltAt(v string) {
	o.BuiltAt = &v
}

// GetOsArchitecture returns the OsArchitecture field value if set, zero value otherwise.
func (o *ContainerImageFlavor) GetOsArchitecture() string {
	if o == nil || o.OsArchitecture == nil {
		var ret string
		return ret
	}
	return *o.OsArchitecture
}

// GetOsArchitectureOk returns a tuple with the OsArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageFlavor) GetOsArchitectureOk() (*string, bool) {
	if o == nil || o.OsArchitecture == nil {
		return nil, false
	}
	return o.OsArchitecture, true
}

// HasOsArchitecture returns a boolean if a field has been set.
func (o *ContainerImageFlavor) HasOsArchitecture() bool {
	return o != nil && o.OsArchitecture != nil
}

// SetOsArchitecture gets a reference to the given string and assigns it to the OsArchitecture field.
func (o *ContainerImageFlavor) SetOsArchitecture(v string) {
	o.OsArchitecture = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *ContainerImageFlavor) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageFlavor) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *ContainerImageFlavor) HasOsName() bool {
	return o != nil && o.OsName != nil
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *ContainerImageFlavor) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *ContainerImageFlavor) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageFlavor) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *ContainerImageFlavor) HasOsVersion() bool {
	return o != nil && o.OsVersion != nil
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *ContainerImageFlavor) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ContainerImageFlavor) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageFlavor) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ContainerImageFlavor) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ContainerImageFlavor) SetSize(v int64) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageFlavor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BuiltAt != nil {
		toSerialize["built_at"] = o.BuiltAt
	}
	if o.OsArchitecture != nil {
		toSerialize["os_architecture"] = o.OsArchitecture
	}
	if o.OsName != nil {
		toSerialize["os_name"] = o.OsName
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerImageFlavor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BuiltAt        *string `json:"built_at,omitempty"`
		OsArchitecture *string `json:"os_architecture,omitempty"`
		OsName         *string `json:"os_name,omitempty"`
		OsVersion      *string `json:"os_version,omitempty"`
		Size           *int64  `json:"size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"built_at", "os_architecture", "os_name", "os_version", "size"})
	} else {
		return err
	}
	o.BuiltAt = all.BuiltAt
	o.OsArchitecture = all.OsArchitecture
	o.OsName = all.OsName
	o.OsVersion = all.OsVersion
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
