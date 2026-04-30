// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationGetServiceSettingsRequestAttributes Attributes for requesting Test Optimization service settings.
type TestOptimizationGetServiceSettingsRequestAttributes struct {
	// The environment name. If omitted, defaults to `none`.
	Env *string `json:"env,omitempty"`
	// The repository identifier.
	RepositoryId string `json:"repository_id"`
	// The service name.
	ServiceName string `json:"service_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestOptimizationGetServiceSettingsRequestAttributes instantiates a new TestOptimizationGetServiceSettingsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestOptimizationGetServiceSettingsRequestAttributes(repositoryId string, serviceName string) *TestOptimizationGetServiceSettingsRequestAttributes {
	this := TestOptimizationGetServiceSettingsRequestAttributes{}
	this.RepositoryId = repositoryId
	this.ServiceName = serviceName
	return &this
}

// NewTestOptimizationGetServiceSettingsRequestAttributesWithDefaults instantiates a new TestOptimizationGetServiceSettingsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestOptimizationGetServiceSettingsRequestAttributesWithDefaults() *TestOptimizationGetServiceSettingsRequestAttributes {
	this := TestOptimizationGetServiceSettingsRequestAttributes{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetRepositoryId returns the RepositoryId field value.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// GetServiceName returns the ServiceName field value.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) SetServiceName(v string) {
	o.ServiceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestOptimizationGetServiceSettingsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	toSerialize["repository_id"] = o.RepositoryId
	toSerialize["service_name"] = o.ServiceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestOptimizationGetServiceSettingsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env          *string `json:"env,omitempty"`
		RepositoryId *string `json:"repository_id"`
		ServiceName  *string `json:"service_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field service_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "repository_id", "service_name"})
	} else {
		return err
	}
	o.Env = all.Env
	o.RepositoryId = *all.RepositoryId
	o.ServiceName = *all.ServiceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
