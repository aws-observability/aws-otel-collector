// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CodeLocation Code vulnerability location.
type CodeLocation struct {
	// Vulnerability location file path.
	FilePath *string `json:"file_path,omitempty"`
	// Vulnerability extracted location.
	Location string `json:"location"`
	// Vulnerability location method.
	Method *string `json:"method,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCodeLocation instantiates a new CodeLocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCodeLocation(location string) *CodeLocation {
	this := CodeLocation{}
	this.Location = location
	return &this
}

// NewCodeLocationWithDefaults instantiates a new CodeLocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCodeLocationWithDefaults() *CodeLocation {
	this := CodeLocation{}
	return &this
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *CodeLocation) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *CodeLocation) HasFilePath() bool {
	return o != nil && o.FilePath != nil
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *CodeLocation) SetFilePath(v string) {
	o.FilePath = &v
}

// GetLocation returns the Location field value.
func (o *CodeLocation) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value.
func (o *CodeLocation) SetLocation(v string) {
	o.Location = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CodeLocation) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeLocation) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CodeLocation) HasMethod() bool {
	return o != nil && o.Method != nil
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CodeLocation) SetMethod(v string) {
	o.Method = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CodeLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FilePath != nil {
		toSerialize["file_path"] = o.FilePath
	}
	toSerialize["location"] = o.Location
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CodeLocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilePath *string `json:"file_path,omitempty"`
		Location *string `json:"location"`
		Method   *string `json:"method,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Location == nil {
		return fmt.Errorf("required field location missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file_path", "location", "method"})
	} else {
		return err
	}
	o.FilePath = all.FilePath
	o.Location = *all.Location
	o.Method = all.Method

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
