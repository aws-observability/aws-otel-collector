// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppFavoriteRequestDataAttributes Attributes for updating an app's favorite status.
type UpdateAppFavoriteRequestDataAttributes struct {
	// Whether the app should be marked as a favorite for the current user.
	Favorite bool `json:"favorite"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppFavoriteRequestDataAttributes instantiates a new UpdateAppFavoriteRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppFavoriteRequestDataAttributes(favorite bool) *UpdateAppFavoriteRequestDataAttributes {
	this := UpdateAppFavoriteRequestDataAttributes{}
	this.Favorite = favorite
	return &this
}

// NewUpdateAppFavoriteRequestDataAttributesWithDefaults instantiates a new UpdateAppFavoriteRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppFavoriteRequestDataAttributesWithDefaults() *UpdateAppFavoriteRequestDataAttributes {
	this := UpdateAppFavoriteRequestDataAttributes{}
	return &this
}

// GetFavorite returns the Favorite field value.
func (o *UpdateAppFavoriteRequestDataAttributes) GetFavorite() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value
// and a boolean to check if the value has been set.
func (o *UpdateAppFavoriteRequestDataAttributes) GetFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favorite, true
}

// SetFavorite sets field value.
func (o *UpdateAppFavoriteRequestDataAttributes) SetFavorite(v bool) {
	o.Favorite = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppFavoriteRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["favorite"] = o.Favorite

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppFavoriteRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Favorite *bool `json:"favorite"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Favorite == nil {
		return fmt.Errorf("required field favorite missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"favorite"})
	} else {
		return err
	}
	o.Favorite = *all.Favorite

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
