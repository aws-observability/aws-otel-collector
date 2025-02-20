// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppEventAttributes JSON object containing all event attributes and their associated values.
type CIAppEventAttributes struct {
	// JSON object of attributes from CI Visibility test events.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Array of tags associated with your event.
	Tags []string `json:"tags,omitempty"`
	// Test run level.
	TestLevel *CIAppTestLevel `json:"test_level,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
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
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TestLevel != nil {
		toSerialize["test_level"] = o.TestLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppEventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes map[string]interface{} `json:"attributes,omitempty"`
		Tags       []string               `json:"tags,omitempty"`
		TestLevel  *CIAppTestLevel        `json:"test_level,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
