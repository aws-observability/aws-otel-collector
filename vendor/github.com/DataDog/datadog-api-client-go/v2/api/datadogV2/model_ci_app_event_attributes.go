// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"time"
=======
	"github.com/goccy/go-json"
>>>>>>> main

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppEventAttributes JSON object containing all event attributes and their associated values.
type CIAppEventAttributes struct {
<<<<<<< HEAD
	// JSON object of attributes from CI Visibility events.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The name of the application or service generating CI Visibility events.
	// It is used to switch from CI Visibility to APM, so make sure you define the same
	// value when you use both products.
	Service *string `json:"service,omitempty"`
	// Array of tags associated with your event.
	Tags []string `json:"tags,omitempty"`
	// Timestamp of your event.
	Timestamp *time.Time `json:"timestamp,omitempty"`
=======
	// JSON object of attributes from CI Visibility test events.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Array of tags associated with your event.
	Tags []string `json:"tags,omitempty"`
	// Test run level.
	TestLevel *CIAppTestLevel `json:"test_level,omitempty"`
>>>>>>> main
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppEventAttributes instantiates a new CIAppEventAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppEventAttributes() *CIAppEventAttributes {
	this := CIAppEventAttributes{}
	return &this
}

// NewCIAppEventAttributesWithDefaults instantiates a new CIAppEventAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppEventAttributesWithDefaults() *CIAppEventAttributes {
	this := CIAppEventAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CIAppEventAttributes) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppEventAttributes) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CIAppEventAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CIAppEventAttributes) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

<<<<<<< HEAD
// GetService returns the Service field value if set, zero value otherwise.
func (o *CIAppEventAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppEventAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CIAppEventAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CIAppEventAttributes) SetService(v string) {
	o.Service = &v
}

=======
>>>>>>> main
// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CIAppEventAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppEventAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CIAppEventAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CIAppEventAttributes) SetTags(v []string) {
	o.Tags = v
}

<<<<<<< HEAD
// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CIAppEventAttributes) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppEventAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CIAppEventAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *CIAppEventAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = &v
=======
// GetTestLevel returns the TestLevel field value if set, zero value otherwise.
func (o *CIAppEventAttributes) GetTestLevel() CIAppTestLevel {
	if o == nil || o.TestLevel == nil {
		var ret CIAppTestLevel
		return ret
	}
	return *o.TestLevel
}

// GetTestLevelOk returns a tuple with the TestLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppEventAttributes) GetTestLevelOk() (*CIAppTestLevel, bool) {
	if o == nil || o.TestLevel == nil {
		return nil, false
	}
	return o.TestLevel, true
}

// HasTestLevel returns a boolean if a field has been set.
func (o *CIAppEventAttributes) HasTestLevel() bool {
	return o != nil && o.TestLevel != nil
}

// SetTestLevel gets a reference to the given CIAppTestLevel and assigns it to the TestLevel field.
func (o *CIAppEventAttributes) SetTestLevel(v CIAppTestLevel) {
	o.TestLevel = &v
>>>>>>> main
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
<<<<<<< HEAD
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		if o.Timestamp.Nanosecond() == 0 {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
=======
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TestLevel != nil {
		toSerialize["test_level"] = o.TestLevel
>>>>>>> main
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppEventAttributes) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
	all := struct {
		Attributes map[string]interface{} `json:"attributes,omitempty"`
		Service    *string                `json:"service,omitempty"`
		Tags       []string               `json:"tags,omitempty"`
		Timestamp  *time.Time             `json:"timestamp,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "service", "tags", "timestamp"})
	} else {
		return err
	}
	o.Attributes = all.Attributes
	o.Service = all.Service
	o.Tags = all.Tags
	o.Timestamp = all.Timestamp
=======
	all := struct {
		Attributes map[string]interface{} `json:"attributes,omitempty"`
		Tags       []string               `json:"tags,omitempty"`
		TestLevel  *CIAppTestLevel        `json:"test_level,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "tags", "test_level"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attributes = all.Attributes
	o.Tags = all.Tags
	if all.TestLevel != nil && !all.TestLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.TestLevel = all.TestLevel
	}

>>>>>>> main
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

<<<<<<< HEAD
=======
	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

>>>>>>> main
	return nil
}
