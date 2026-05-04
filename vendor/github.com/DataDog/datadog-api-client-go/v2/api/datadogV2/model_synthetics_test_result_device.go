// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultDevice Device information for the test result (browser and mobile tests).
type SyntheticsTestResultDevice struct {
	// Browser information for the device used to run the test.
	Browser *SyntheticsTestResultDeviceBrowser `json:"browser,omitempty"`
	// Device identifier.
	Id *string `json:"id,omitempty"`
	// Device name.
	Name *string `json:"name,omitempty"`
	// Platform information for the device used to run the test.
	Platform *SyntheticsTestResultDevicePlatform `json:"platform,omitempty"`
	// Screen resolution of the device used to run the test.
	Resolution *SyntheticsTestResultDeviceResolution `json:"resolution,omitempty"`
	// Device type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultDevice instantiates a new SyntheticsTestResultDevice object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultDevice() *SyntheticsTestResultDevice {
	this := SyntheticsTestResultDevice{}
	return &this
}

// NewSyntheticsTestResultDeviceWithDefaults instantiates a new SyntheticsTestResultDevice object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDeviceWithDefaults() *SyntheticsTestResultDevice {
	this := SyntheticsTestResultDevice{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *SyntheticsTestResultDevice) GetBrowser() SyntheticsTestResultDeviceBrowser {
	if o == nil || o.Browser == nil {
		var ret SyntheticsTestResultDeviceBrowser
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDevice) GetBrowserOk() (*SyntheticsTestResultDeviceBrowser, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *SyntheticsTestResultDevice) HasBrowser() bool {
	return o != nil && o.Browser != nil
}

// SetBrowser gets a reference to the given SyntheticsTestResultDeviceBrowser and assigns it to the Browser field.
func (o *SyntheticsTestResultDevice) SetBrowser(v SyntheticsTestResultDeviceBrowser) {
	o.Browser = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultDevice) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDevice) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultDevice) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultDevice) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultDevice) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDevice) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultDevice) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultDevice) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *SyntheticsTestResultDevice) GetPlatform() SyntheticsTestResultDevicePlatform {
	if o == nil || o.Platform == nil {
		var ret SyntheticsTestResultDevicePlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDevice) GetPlatformOk() (*SyntheticsTestResultDevicePlatform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *SyntheticsTestResultDevice) HasPlatform() bool {
	return o != nil && o.Platform != nil
}

// SetPlatform gets a reference to the given SyntheticsTestResultDevicePlatform and assigns it to the Platform field.
func (o *SyntheticsTestResultDevice) SetPlatform(v SyntheticsTestResultDevicePlatform) {
	o.Platform = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *SyntheticsTestResultDevice) GetResolution() SyntheticsTestResultDeviceResolution {
	if o == nil || o.Resolution == nil {
		var ret SyntheticsTestResultDeviceResolution
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDevice) GetResolutionOk() (*SyntheticsTestResultDeviceResolution, bool) {
	if o == nil || o.Resolution == nil {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *SyntheticsTestResultDevice) HasResolution() bool {
	return o != nil && o.Resolution != nil
}

// SetResolution gets a reference to the given SyntheticsTestResultDeviceResolution and assigns it to the Resolution field.
func (o *SyntheticsTestResultDevice) SetResolution(v SyntheticsTestResultDeviceResolution) {
	o.Resolution = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultDevice) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDevice) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultDevice) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultDevice) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Browser != nil {
		toSerialize["browser"] = o.Browser
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Resolution != nil {
		toSerialize["resolution"] = o.Resolution
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultDevice) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Browser    *SyntheticsTestResultDeviceBrowser    `json:"browser,omitempty"`
		Id         *string                               `json:"id,omitempty"`
		Name       *string                               `json:"name,omitempty"`
		Platform   *SyntheticsTestResultDevicePlatform   `json:"platform,omitempty"`
		Resolution *SyntheticsTestResultDeviceResolution `json:"resolution,omitempty"`
		Type       *string                               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"browser", "id", "name", "platform", "resolution", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Browser != nil && all.Browser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Browser = all.Browser
	o.Id = all.Id
	o.Name = all.Name
	if all.Platform != nil && all.Platform.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Platform = all.Platform
	if all.Resolution != nil && all.Resolution.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resolution = all.Resolution
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
