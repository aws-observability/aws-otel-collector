// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimelineCellAuthorUser timeline cell user author
type TimelineCellAuthorUser struct {
	// user author content.
	Content *TimelineCellAuthorUserContent `json:"content,omitempty"`
	// user author type.
	Type *TimelineCellAuthorUserType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimelineCellAuthorUser instantiates a new TimelineCellAuthorUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimelineCellAuthorUser() *TimelineCellAuthorUser {
	this := TimelineCellAuthorUser{}
	return &this
}

// NewTimelineCellAuthorUserWithDefaults instantiates a new TimelineCellAuthorUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimelineCellAuthorUserWithDefaults() *TimelineCellAuthorUser {
	this := TimelineCellAuthorUser{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TimelineCellAuthorUser) GetContent() TimelineCellAuthorUserContent {
	if o == nil || o.Content == nil {
		var ret TimelineCellAuthorUserContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCellAuthorUser) GetContentOk() (*TimelineCellAuthorUserContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TimelineCellAuthorUser) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given TimelineCellAuthorUserContent and assigns it to the Content field.
func (o *TimelineCellAuthorUser) SetContent(v TimelineCellAuthorUserContent) {
	o.Content = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TimelineCellAuthorUser) GetType() TimelineCellAuthorUserType {
	if o == nil || o.Type == nil {
		var ret TimelineCellAuthorUserType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCellAuthorUser) GetTypeOk() (*TimelineCellAuthorUserType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TimelineCellAuthorUser) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given TimelineCellAuthorUserType and assigns it to the Type field.
func (o *TimelineCellAuthorUser) SetType(v TimelineCellAuthorUserType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimelineCellAuthorUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimelineCellAuthorUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content *TimelineCellAuthorUserContent `json:"content,omitempty"`
		Type    *TimelineCellAuthorUserType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Content != nil && all.Content.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Content = all.Content
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
