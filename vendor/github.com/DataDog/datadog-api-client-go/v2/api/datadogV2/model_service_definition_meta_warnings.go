// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionMetaWarnings Schema validation warnings.
type ServiceDefinitionMetaWarnings struct {
	// The warning instance location.
	InstanceLocation *string `json:"instance-location,omitempty"`
	// The warning keyword location.
	KeywordLocation *string `json:"keyword-location,omitempty"`
	// The warning message.
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionMetaWarnings instantiates a new ServiceDefinitionMetaWarnings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionMetaWarnings() *ServiceDefinitionMetaWarnings {
	this := ServiceDefinitionMetaWarnings{}
	return &this
}

// NewServiceDefinitionMetaWarningsWithDefaults instantiates a new ServiceDefinitionMetaWarnings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionMetaWarningsWithDefaults() *ServiceDefinitionMetaWarnings {
	this := ServiceDefinitionMetaWarnings{}
	return &this
}

// GetInstanceLocation returns the InstanceLocation field value if set, zero value otherwise.
func (o *ServiceDefinitionMetaWarnings) GetInstanceLocation() string {
	if o == nil || o.InstanceLocation == nil {
		var ret string
		return ret
	}
	return *o.InstanceLocation
}

// GetInstanceLocationOk returns a tuple with the InstanceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMetaWarnings) GetInstanceLocationOk() (*string, bool) {
	if o == nil || o.InstanceLocation == nil {
		return nil, false
	}
	return o.InstanceLocation, true
}

// HasInstanceLocation returns a boolean if a field has been set.
func (o *ServiceDefinitionMetaWarnings) HasInstanceLocation() bool {
	return o != nil && o.InstanceLocation != nil
}

// SetInstanceLocation gets a reference to the given string and assigns it to the InstanceLocation field.
func (o *ServiceDefinitionMetaWarnings) SetInstanceLocation(v string) {
	o.InstanceLocation = &v
}

// GetKeywordLocation returns the KeywordLocation field value if set, zero value otherwise.
func (o *ServiceDefinitionMetaWarnings) GetKeywordLocation() string {
	if o == nil || o.KeywordLocation == nil {
		var ret string
		return ret
	}
	return *o.KeywordLocation
}

// GetKeywordLocationOk returns a tuple with the KeywordLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMetaWarnings) GetKeywordLocationOk() (*string, bool) {
	if o == nil || o.KeywordLocation == nil {
		return nil, false
	}
	return o.KeywordLocation, true
}

// HasKeywordLocation returns a boolean if a field has been set.
func (o *ServiceDefinitionMetaWarnings) HasKeywordLocation() bool {
	return o != nil && o.KeywordLocation != nil
}

// SetKeywordLocation gets a reference to the given string and assigns it to the KeywordLocation field.
func (o *ServiceDefinitionMetaWarnings) SetKeywordLocation(v string) {
	o.KeywordLocation = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ServiceDefinitionMetaWarnings) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMetaWarnings) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServiceDefinitionMetaWarnings) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServiceDefinitionMetaWarnings) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionMetaWarnings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.InstanceLocation != nil {
		toSerialize["instance-location"] = o.InstanceLocation
	}
	if o.KeywordLocation != nil {
		toSerialize["keyword-location"] = o.KeywordLocation
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionMetaWarnings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceLocation *string `json:"instance-location,omitempty"`
		KeywordLocation  *string `json:"keyword-location,omitempty"`
		Message          *string `json:"message,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance-location", "keyword-location", "message"})
	} else {
		return err
	}
	o.InstanceLocation = all.InstanceLocation
	o.KeywordLocation = all.KeywordLocation
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
