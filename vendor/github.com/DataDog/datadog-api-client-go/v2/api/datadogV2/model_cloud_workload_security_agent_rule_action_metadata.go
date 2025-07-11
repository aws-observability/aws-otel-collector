// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleActionMetadata The metadata action applied on the scope matching the rule
type CloudWorkloadSecurityAgentRuleActionMetadata struct {
	// The image tag of the metadata action
	ImageTag *string `json:"image_tag,omitempty"`
	// The service of the metadata action
	Service *string `json:"service,omitempty"`
	// The short image of the metadata action
	ShortImage *string `json:"short_image,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleActionMetadata instantiates a new CloudWorkloadSecurityAgentRuleActionMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleActionMetadata() *CloudWorkloadSecurityAgentRuleActionMetadata {
	this := CloudWorkloadSecurityAgentRuleActionMetadata{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionMetadataWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleActionMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionMetadataWithDefaults() *CloudWorkloadSecurityAgentRuleActionMetadata {
	this := CloudWorkloadSecurityAgentRuleActionMetadata{}
	return &this
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) HasImageTag() bool {
	return o != nil && o.ImageTag != nil
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) SetService(v string) {
	o.Service = &v
}

// GetShortImage returns the ShortImage field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) GetShortImage() string {
	if o == nil || o.ShortImage == nil {
		var ret string
		return ret
	}
	return *o.ShortImage
}

// GetShortImageOk returns a tuple with the ShortImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) GetShortImageOk() (*string, bool) {
	if o == nil || o.ShortImage == nil {
		return nil, false
	}
	return o.ShortImage, true
}

// HasShortImage returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) HasShortImage() bool {
	return o != nil && o.ShortImage != nil
}

// SetShortImage gets a reference to the given string and assigns it to the ShortImage field.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) SetShortImage(v string) {
	o.ShortImage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleActionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ShortImage != nil {
		toSerialize["short_image"] = o.ShortImage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleActionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ImageTag   *string `json:"image_tag,omitempty"`
		Service    *string `json:"service,omitempty"`
		ShortImage *string `json:"short_image,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"image_tag", "service", "short_image"})
	} else {
		return err
	}
	o.ImageTag = all.ImageTag
	o.Service = all.Service
	o.ShortImage = all.ShortImage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
