// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatTargetAudienceCreateRequestAttributes Attributes for creating a Google Chat target audience.
type GoogleChatTargetAudienceCreateRequestAttributes struct {
	// The audience ID.
	AudienceId string `json:"audience_id"`
	// The audience name.
	AudienceName string `json:"audience_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatTargetAudienceCreateRequestAttributes instantiates a new GoogleChatTargetAudienceCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatTargetAudienceCreateRequestAttributes(audienceId string, audienceName string) *GoogleChatTargetAudienceCreateRequestAttributes {
	this := GoogleChatTargetAudienceCreateRequestAttributes{}
	this.AudienceId = audienceId
	this.AudienceName = audienceName
	return &this
}

// NewGoogleChatTargetAudienceCreateRequestAttributesWithDefaults instantiates a new GoogleChatTargetAudienceCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatTargetAudienceCreateRequestAttributesWithDefaults() *GoogleChatTargetAudienceCreateRequestAttributes {
	this := GoogleChatTargetAudienceCreateRequestAttributes{}
	return &this
}

// GetAudienceId returns the AudienceId field value.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) GetAudienceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AudienceId
}

// GetAudienceIdOk returns a tuple with the AudienceId field value
// and a boolean to check if the value has been set.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) GetAudienceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceId, true
}

// SetAudienceId sets field value.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) SetAudienceId(v string) {
	o.AudienceId = v
}

// GetAudienceName returns the AudienceName field value.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) GetAudienceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AudienceName
}

// GetAudienceNameOk returns a tuple with the AudienceName field value
// and a boolean to check if the value has been set.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) GetAudienceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceName, true
}

// SetAudienceName sets field value.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) SetAudienceName(v string) {
	o.AudienceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatTargetAudienceCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["audience_id"] = o.AudienceId
	toSerialize["audience_name"] = o.AudienceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatTargetAudienceCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceId   *string `json:"audience_id"`
		AudienceName *string `json:"audience_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AudienceId == nil {
		return fmt.Errorf("required field audience_id missing")
	}
	if all.AudienceName == nil {
		return fmt.Errorf("required field audience_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_id", "audience_name"})
	} else {
		return err
	}
	o.AudienceId = *all.AudienceId
	o.AudienceName = *all.AudienceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
