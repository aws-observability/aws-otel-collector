// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAPIEndpoint Endpoint info extracted from an `OpenAPI` specification.
type OpenAPIEndpoint struct {
	// The endpoint method.
	Method *string `json:"method,omitempty"`
	// The endpoint path.
	Path *string `json:"path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAPIEndpoint instantiates a new OpenAPIEndpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAPIEndpoint() *OpenAPIEndpoint {
	this := OpenAPIEndpoint{}
	return &this
}

// NewOpenAPIEndpointWithDefaults instantiates a new OpenAPIEndpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAPIEndpointWithDefaults() *OpenAPIEndpoint {
	this := OpenAPIEndpoint{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *OpenAPIEndpoint) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPIEndpoint) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *OpenAPIEndpoint) HasMethod() bool {
	return o != nil && o.Method != nil
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *OpenAPIEndpoint) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OpenAPIEndpoint) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPIEndpoint) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OpenAPIEndpoint) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OpenAPIEndpoint) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAPIEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAPIEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Method *string `json:"method,omitempty"`
		Path   *string `json:"path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"method", "path"})
	} else {
		return err
	}
	o.Method = all.Method
	o.Path = all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
