// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityServicesMetadata Metadata returned alongside the list of services.
type ApplicationSecurityServicesMetadata struct {
	// The number of services with Application Security Management (Threats) enabled.
	NumServicesWithAppsec int64 `json:"num_services_with_appsec"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityServicesMetadata instantiates a new ApplicationSecurityServicesMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityServicesMetadata(numServicesWithAppsec int64) *ApplicationSecurityServicesMetadata {
	this := ApplicationSecurityServicesMetadata{}
	this.NumServicesWithAppsec = numServicesWithAppsec
	return &this
}

// NewApplicationSecurityServicesMetadataWithDefaults instantiates a new ApplicationSecurityServicesMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityServicesMetadataWithDefaults() *ApplicationSecurityServicesMetadata {
	this := ApplicationSecurityServicesMetadata{}
	return &this
}

// GetNumServicesWithAppsec returns the NumServicesWithAppsec field value.
func (o *ApplicationSecurityServicesMetadata) GetNumServicesWithAppsec() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NumServicesWithAppsec
}

// GetNumServicesWithAppsecOk returns a tuple with the NumServicesWithAppsec field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServicesMetadata) GetNumServicesWithAppsecOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumServicesWithAppsec, true
}

// SetNumServicesWithAppsec sets field value.
func (o *ApplicationSecurityServicesMetadata) SetNumServicesWithAppsec(v int64) {
	o.NumServicesWithAppsec = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityServicesMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["num_services_with_appsec"] = o.NumServicesWithAppsec

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityServicesMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumServicesWithAppsec *int64 `json:"num_services_with_appsec"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NumServicesWithAppsec == nil {
		return fmt.Errorf("required field num_services_with_appsec missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"num_services_with_appsec"})
	} else {
		return err
	}
	o.NumServicesWithAppsec = *all.NumServicesWithAppsec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
