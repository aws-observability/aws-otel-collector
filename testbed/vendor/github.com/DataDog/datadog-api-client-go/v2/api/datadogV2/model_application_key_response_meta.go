// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationKeyResponseMeta Additional information related to the application key response.
type ApplicationKeyResponseMeta struct {
	// Max allowed number of application keys per user.
	MaxAllowedPerUser *int64 `json:"max_allowed_per_user,omitempty"`
	// Additional information related to the application key response.
	Page *ApplicationKeyResponseMetaPage `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationKeyResponseMeta instantiates a new ApplicationKeyResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyResponseMeta() *ApplicationKeyResponseMeta {
	this := ApplicationKeyResponseMeta{}
	return &this
}

// NewApplicationKeyResponseMetaWithDefaults instantiates a new ApplicationKeyResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyResponseMetaWithDefaults() *ApplicationKeyResponseMeta {
	this := ApplicationKeyResponseMeta{}
	return &this
}

// GetMaxAllowedPerUser returns the MaxAllowedPerUser field value if set, zero value otherwise.
func (o *ApplicationKeyResponseMeta) GetMaxAllowedPerUser() int64 {
	if o == nil || o.MaxAllowedPerUser == nil {
		var ret int64
		return ret
	}
	return *o.MaxAllowedPerUser
}

// GetMaxAllowedPerUserOk returns a tuple with the MaxAllowedPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyResponseMeta) GetMaxAllowedPerUserOk() (*int64, bool) {
	if o == nil || o.MaxAllowedPerUser == nil {
		return nil, false
	}
	return o.MaxAllowedPerUser, true
}

// HasMaxAllowedPerUser returns a boolean if a field has been set.
func (o *ApplicationKeyResponseMeta) HasMaxAllowedPerUser() bool {
	return o != nil && o.MaxAllowedPerUser != nil
}

// SetMaxAllowedPerUser gets a reference to the given int64 and assigns it to the MaxAllowedPerUser field.
func (o *ApplicationKeyResponseMeta) SetMaxAllowedPerUser(v int64) {
	o.MaxAllowedPerUser = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ApplicationKeyResponseMeta) GetPage() ApplicationKeyResponseMetaPage {
	if o == nil || o.Page == nil {
		var ret ApplicationKeyResponseMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyResponseMeta) GetPageOk() (*ApplicationKeyResponseMetaPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ApplicationKeyResponseMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given ApplicationKeyResponseMetaPage and assigns it to the Page field.
func (o *ApplicationKeyResponseMeta) SetPage(v ApplicationKeyResponseMetaPage) {
	o.Page = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MaxAllowedPerUser != nil {
		toSerialize["max_allowed_per_user"] = o.MaxAllowedPerUser
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxAllowedPerUser *int64                          `json:"max_allowed_per_user,omitempty"`
		Page              *ApplicationKeyResponseMetaPage `json:"page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_allowed_per_user", "page"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MaxAllowedPerUser = all.MaxAllowedPerUser
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
