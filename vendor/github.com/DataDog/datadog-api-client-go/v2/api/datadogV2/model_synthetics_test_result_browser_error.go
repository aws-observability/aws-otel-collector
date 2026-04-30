// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultBrowserError A browser error captured during a browser test step.
type SyntheticsTestResultBrowserError struct {
	// Error description.
	Description *string `json:"description,omitempty"`
	// HTTP method associated with the error (for network errors).
	Method *string `json:"method,omitempty"`
	// Error name.
	Name *string `json:"name,omitempty"`
	// HTTP status code associated with the error (for network errors).
	Status *int64 `json:"status,omitempty"`
	// Type of the browser error.
	Type *string `json:"type,omitempty"`
	// URL associated with the error.
	Url map[string]interface{} `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultBrowserError instantiates a new SyntheticsTestResultBrowserError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultBrowserError() *SyntheticsTestResultBrowserError {
	this := SyntheticsTestResultBrowserError{}
	return &this
}

// NewSyntheticsTestResultBrowserErrorWithDefaults instantiates a new SyntheticsTestResultBrowserError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultBrowserErrorWithDefaults() *SyntheticsTestResultBrowserError {
	this := SyntheticsTestResultBrowserError{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyntheticsTestResultBrowserError) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBrowserError) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyntheticsTestResultBrowserError) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyntheticsTestResultBrowserError) SetDescription(v string) {
	o.Description = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SyntheticsTestResultBrowserError) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBrowserError) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SyntheticsTestResultBrowserError) HasMethod() bool {
	return o != nil && o.Method != nil
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SyntheticsTestResultBrowserError) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultBrowserError) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBrowserError) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultBrowserError) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultBrowserError) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultBrowserError) GetStatus() int64 {
	if o == nil || o.Status == nil {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBrowserError) GetStatusOk() (*int64, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultBrowserError) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SyntheticsTestResultBrowserError) SetStatus(v int64) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultBrowserError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBrowserError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultBrowserError) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultBrowserError) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultBrowserError) GetUrl() map[string]interface{} {
	if o == nil || o.Url == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBrowserError) GetUrlOk() (*map[string]interface{}, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return &o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultBrowserError) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given map[string]interface{} and assigns it to the Url field.
func (o *SyntheticsTestResultBrowserError) SetUrl(v map[string]interface{}) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultBrowserError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultBrowserError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                `json:"description,omitempty"`
		Method      *string                `json:"method,omitempty"`
		Name        *string                `json:"name,omitempty"`
		Status      *int64                 `json:"status,omitempty"`
		Type        *string                `json:"type,omitempty"`
		Url         map[string]interface{} `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "method", "name", "status", "type", "url"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Method = all.Method
	o.Name = all.Name
	o.Status = all.Status
	o.Type = all.Type
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
