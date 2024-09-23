// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationKeyResponse An application key response.
type ApplicationKeyResponse struct {
	// An application key with its associated metadata.
	ApplicationKey *ApplicationKey `json:"application_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationKeyResponse instantiates a new ApplicationKeyResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyResponse() *ApplicationKeyResponse {
	this := ApplicationKeyResponse{}
	return &this
}

// NewApplicationKeyResponseWithDefaults instantiates a new ApplicationKeyResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyResponseWithDefaults() *ApplicationKeyResponse {
	this := ApplicationKeyResponse{}
	return &this
}

// GetApplicationKey returns the ApplicationKey field value if set, zero value otherwise.
func (o *ApplicationKeyResponse) GetApplicationKey() ApplicationKey {
	if o == nil || o.ApplicationKey == nil {
		var ret ApplicationKey
		return ret
	}
	return *o.ApplicationKey
}

// GetApplicationKeyOk returns a tuple with the ApplicationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyResponse) GetApplicationKeyOk() (*ApplicationKey, bool) {
	if o == nil || o.ApplicationKey == nil {
		return nil, false
	}
	return o.ApplicationKey, true
}

// HasApplicationKey returns a boolean if a field has been set.
func (o *ApplicationKeyResponse) HasApplicationKey() bool {
	return o != nil && o.ApplicationKey != nil
}

// SetApplicationKey gets a reference to the given ApplicationKey and assigns it to the ApplicationKey field.
func (o *ApplicationKeyResponse) SetApplicationKey(v ApplicationKey) {
	o.ApplicationKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationKey != nil {
		toSerialize["application_key"] = o.ApplicationKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationKey *ApplicationKey `json:"application_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApplicationKey != nil && all.ApplicationKey.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApplicationKey = all.ApplicationKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
