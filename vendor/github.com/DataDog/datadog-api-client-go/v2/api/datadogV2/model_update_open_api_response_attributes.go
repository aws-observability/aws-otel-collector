// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateOpenAPIResponseAttributes Attributes for `UpdateOpenAPI`.
type UpdateOpenAPIResponseAttributes struct {
	// List of endpoints which couldn't be parsed.
	FailedEndpoints []OpenAPIEndpoint `json:"failed_endpoints,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateOpenAPIResponseAttributes instantiates a new UpdateOpenAPIResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateOpenAPIResponseAttributes() *UpdateOpenAPIResponseAttributes {
	this := UpdateOpenAPIResponseAttributes{}
	return &this
}

// NewUpdateOpenAPIResponseAttributesWithDefaults instantiates a new UpdateOpenAPIResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateOpenAPIResponseAttributesWithDefaults() *UpdateOpenAPIResponseAttributes {
	this := UpdateOpenAPIResponseAttributes{}
	return &this
}

// GetFailedEndpoints returns the FailedEndpoints field value if set, zero value otherwise.
func (o *UpdateOpenAPIResponseAttributes) GetFailedEndpoints() []OpenAPIEndpoint {
	if o == nil || o.FailedEndpoints == nil {
		var ret []OpenAPIEndpoint
		return ret
	}
	return o.FailedEndpoints
}

// GetFailedEndpointsOk returns a tuple with the FailedEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOpenAPIResponseAttributes) GetFailedEndpointsOk() (*[]OpenAPIEndpoint, bool) {
	if o == nil || o.FailedEndpoints == nil {
		return nil, false
	}
	return &o.FailedEndpoints, true
}

// HasFailedEndpoints returns a boolean if a field has been set.
func (o *UpdateOpenAPIResponseAttributes) HasFailedEndpoints() bool {
	return o != nil && o.FailedEndpoints != nil
}

// SetFailedEndpoints gets a reference to the given []OpenAPIEndpoint and assigns it to the FailedEndpoints field.
func (o *UpdateOpenAPIResponseAttributes) SetFailedEndpoints(v []OpenAPIEndpoint) {
	o.FailedEndpoints = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateOpenAPIResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FailedEndpoints != nil {
		toSerialize["failed_endpoints"] = o.FailedEndpoints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateOpenAPIResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FailedEndpoints []OpenAPIEndpoint `json:"failed_endpoints,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failed_endpoints"})
	} else {
		return err
	}
	o.FailedEndpoints = all.FailedEndpoints

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
