// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppProtectionLevelRequestDataAttributes Attributes for updating an app's publication protection level.
type UpdateAppProtectionLevelRequestDataAttributes struct {
	// The publication protection level of the app. `approval_required` means changes must go through an approval workflow before being published.
	ProtectionLevel AppProtectionLevel `json:"protectionLevel"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppProtectionLevelRequestDataAttributes instantiates a new UpdateAppProtectionLevelRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppProtectionLevelRequestDataAttributes(protectionLevel AppProtectionLevel) *UpdateAppProtectionLevelRequestDataAttributes {
	this := UpdateAppProtectionLevelRequestDataAttributes{}
	this.ProtectionLevel = protectionLevel
	return &this
}

// NewUpdateAppProtectionLevelRequestDataAttributesWithDefaults instantiates a new UpdateAppProtectionLevelRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppProtectionLevelRequestDataAttributesWithDefaults() *UpdateAppProtectionLevelRequestDataAttributes {
	this := UpdateAppProtectionLevelRequestDataAttributes{}
	return &this
}

// GetProtectionLevel returns the ProtectionLevel field value.
func (o *UpdateAppProtectionLevelRequestDataAttributes) GetProtectionLevel() AppProtectionLevel {
	if o == nil {
		var ret AppProtectionLevel
		return ret
	}
	return o.ProtectionLevel
}

// GetProtectionLevelOk returns a tuple with the ProtectionLevel field value
// and a boolean to check if the value has been set.
func (o *UpdateAppProtectionLevelRequestDataAttributes) GetProtectionLevelOk() (*AppProtectionLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectionLevel, true
}

// SetProtectionLevel sets field value.
func (o *UpdateAppProtectionLevelRequestDataAttributes) SetProtectionLevel(v AppProtectionLevel) {
	o.ProtectionLevel = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppProtectionLevelRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["protectionLevel"] = o.ProtectionLevel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppProtectionLevelRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProtectionLevel *AppProtectionLevel `json:"protectionLevel"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProtectionLevel == nil {
		return fmt.Errorf("required field protectionLevel missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"protectionLevel"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ProtectionLevel.IsValid() {
		hasInvalidField = true
	} else {
		o.ProtectionLevel = *all.ProtectionLevel
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
