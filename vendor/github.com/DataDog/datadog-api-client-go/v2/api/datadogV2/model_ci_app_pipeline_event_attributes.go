// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventAttributes JSON object containing all event attributes and their associated values.
type CIAppPipelineEventAttributes struct {
	// JSON object of attributes from CI Visibility pipeline events.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Pipeline execution level.
	CiLevel *CIAppPipelineLevel `json:"ci_level,omitempty"`
	// Array of tags associated with your event.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppPipelineEventAttributes instantiates a new CIAppPipelineEventAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelineEventAttributes() *CIAppPipelineEventAttributes {
	this := CIAppPipelineEventAttributes{}
	return &this
}

// NewCIAppPipelineEventAttributesWithDefaults instantiates a new CIAppPipelineEventAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelineEventAttributesWithDefaults() *CIAppPipelineEventAttributes {
	this := CIAppPipelineEventAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CIAppPipelineEventAttributes) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventAttributes) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CIAppPipelineEventAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CIAppPipelineEventAttributes) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetCiLevel returns the CiLevel field value if set, zero value otherwise.
func (o *CIAppPipelineEventAttributes) GetCiLevel() CIAppPipelineLevel {
	if o == nil || o.CiLevel == nil {
		var ret CIAppPipelineLevel
		return ret
	}
	return *o.CiLevel
}

// GetCiLevelOk returns a tuple with the CiLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventAttributes) GetCiLevelOk() (*CIAppPipelineLevel, bool) {
	if o == nil || o.CiLevel == nil {
		return nil, false
	}
	return o.CiLevel, true
}

// HasCiLevel returns a boolean if a field has been set.
func (o *CIAppPipelineEventAttributes) HasCiLevel() bool {
	return o != nil && o.CiLevel != nil
}

// SetCiLevel gets a reference to the given CIAppPipelineLevel and assigns it to the CiLevel field.
func (o *CIAppPipelineEventAttributes) SetCiLevel(v CIAppPipelineLevel) {
	o.CiLevel = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CIAppPipelineEventAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CIAppPipelineEventAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CIAppPipelineEventAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelineEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.CiLevel != nil {
		toSerialize["ci_level"] = o.CiLevel
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppPipelineEventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes map[string]interface{} `json:"attributes,omitempty"`
		CiLevel    *CIAppPipelineLevel    `json:"ci_level,omitempty"`
		Tags       []string               `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "ci_level", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attributes = all.Attributes
	if all.CiLevel != nil && !all.CiLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.CiLevel = all.CiLevel
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
