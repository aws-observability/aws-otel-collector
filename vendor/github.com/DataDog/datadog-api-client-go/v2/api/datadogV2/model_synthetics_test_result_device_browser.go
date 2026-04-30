// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultDeviceBrowser Browser information for the device used to run the test.
type SyntheticsTestResultDeviceBrowser struct {
	// Browser type (for example, `chrome`, `firefox`).
	Type *string `json:"type,omitempty"`
	// User agent string reported by the browser.
	UserAgent *string `json:"user_agent,omitempty"`
	// Browser version.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultDeviceBrowser instantiates a new SyntheticsTestResultDeviceBrowser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultDeviceBrowser() *SyntheticsTestResultDeviceBrowser {
	this := SyntheticsTestResultDeviceBrowser{}
	return &this
}

// NewSyntheticsTestResultDeviceBrowserWithDefaults instantiates a new SyntheticsTestResultDeviceBrowser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDeviceBrowserWithDefaults() *SyntheticsTestResultDeviceBrowser {
	this := SyntheticsTestResultDeviceBrowser{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultDeviceBrowser) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDeviceBrowser) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultDeviceBrowser) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultDeviceBrowser) SetType(v string) {
	o.Type = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *SyntheticsTestResultDeviceBrowser) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDeviceBrowser) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *SyntheticsTestResultDeviceBrowser) HasUserAgent() bool {
	return o != nil && o.UserAgent != nil
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *SyntheticsTestResultDeviceBrowser) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SyntheticsTestResultDeviceBrowser) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDeviceBrowser) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SyntheticsTestResultDeviceBrowser) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SyntheticsTestResultDeviceBrowser) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultDeviceBrowser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
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
func (o *SyntheticsTestResultDeviceBrowser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type      *string `json:"type,omitempty"`
		UserAgent *string `json:"user_agent,omitempty"`
		Version   *string `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "user_agent", "version"})
	} else {
		return err
	}
	o.Type = all.Type
	o.UserAgent = all.UserAgent
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
