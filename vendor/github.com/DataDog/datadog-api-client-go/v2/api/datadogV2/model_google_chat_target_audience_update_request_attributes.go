// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatTargetAudienceUpdateRequestAttributes Attributes for updating a Google Chat target audience.
type GoogleChatTargetAudienceUpdateRequestAttributes struct {
	// The audience ID.
	AudienceId *string `json:"audience_id,omitempty"`
	// The audience name.
	AudienceName *string `json:"audience_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatTargetAudienceUpdateRequestAttributes instantiates a new GoogleChatTargetAudienceUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatTargetAudienceUpdateRequestAttributes() *GoogleChatTargetAudienceUpdateRequestAttributes {
	this := GoogleChatTargetAudienceUpdateRequestAttributes{}
	return &this
}

// NewGoogleChatTargetAudienceUpdateRequestAttributesWithDefaults instantiates a new GoogleChatTargetAudienceUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatTargetAudienceUpdateRequestAttributesWithDefaults() *GoogleChatTargetAudienceUpdateRequestAttributes {
	this := GoogleChatTargetAudienceUpdateRequestAttributes{}
	return &this
}

// GetAudienceId returns the AudienceId field value if set, zero value otherwise.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) GetAudienceId() string {
	if o == nil || o.AudienceId == nil {
		var ret string
		return ret
	}
	return *o.AudienceId
}

// GetAudienceIdOk returns a tuple with the AudienceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) GetAudienceIdOk() (*string, bool) {
	if o == nil || o.AudienceId == nil {
		return nil, false
	}
	return o.AudienceId, true
}

// HasAudienceId returns a boolean if a field has been set.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) HasAudienceId() bool {
	return o != nil && o.AudienceId != nil
}

// SetAudienceId gets a reference to the given string and assigns it to the AudienceId field.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) SetAudienceId(v string) {
	o.AudienceId = &v
}

// GetAudienceName returns the AudienceName field value if set, zero value otherwise.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) GetAudienceName() string {
	if o == nil || o.AudienceName == nil {
		var ret string
		return ret
	}
	return *o.AudienceName
}

// GetAudienceNameOk returns a tuple with the AudienceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) GetAudienceNameOk() (*string, bool) {
	if o == nil || o.AudienceName == nil {
		return nil, false
	}
	return o.AudienceName, true
}

// HasAudienceName returns a boolean if a field has been set.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) HasAudienceName() bool {
	return o != nil && o.AudienceName != nil
}

// SetAudienceName gets a reference to the given string and assigns it to the AudienceName field.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) SetAudienceName(v string) {
	o.AudienceName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatTargetAudienceUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceId != nil {
		toSerialize["audience_id"] = o.AudienceId
	}
	if o.AudienceName != nil {
		toSerialize["audience_name"] = o.AudienceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatTargetAudienceUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceId   *string `json:"audience_id,omitempty"`
		AudienceName *string `json:"audience_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_id", "audience_name"})
	} else {
		return err
	}
	o.AudienceId = all.AudienceId
	o.AudienceName = all.AudienceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
