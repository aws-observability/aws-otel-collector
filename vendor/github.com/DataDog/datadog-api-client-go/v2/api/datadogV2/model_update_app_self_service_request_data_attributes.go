// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppSelfServiceRequestDataAttributes Attributes for updating an app's self-service status.
type UpdateAppSelfServiceRequestDataAttributes struct {
	// Whether the app is enabled for self-service.
	SelfService bool `json:"selfService"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppSelfServiceRequestDataAttributes instantiates a new UpdateAppSelfServiceRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppSelfServiceRequestDataAttributes(selfService bool) *UpdateAppSelfServiceRequestDataAttributes {
	this := UpdateAppSelfServiceRequestDataAttributes{}
	this.SelfService = selfService
	return &this
}

// NewUpdateAppSelfServiceRequestDataAttributesWithDefaults instantiates a new UpdateAppSelfServiceRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppSelfServiceRequestDataAttributesWithDefaults() *UpdateAppSelfServiceRequestDataAttributes {
	this := UpdateAppSelfServiceRequestDataAttributes{}
	return &this
}

// GetSelfService returns the SelfService field value.
func (o *UpdateAppSelfServiceRequestDataAttributes) GetSelfService() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SelfService
}

// GetSelfServiceOk returns a tuple with the SelfService field value
// and a boolean to check if the value has been set.
func (o *UpdateAppSelfServiceRequestDataAttributes) GetSelfServiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfService, true
}

// SetSelfService sets field value.
func (o *UpdateAppSelfServiceRequestDataAttributes) SetSelfService(v bool) {
	o.SelfService = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppSelfServiceRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["selfService"] = o.SelfService

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppSelfServiceRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SelfService *bool `json:"selfService"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SelfService == nil {
		return fmt.Errorf("required field selfService missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"selfService"})
	} else {
		return err
	}
	o.SelfService = *all.SelfService

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
