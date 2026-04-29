// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationRequestAttributes Attributes for a deployment gate evaluation request.
type DeploymentGatesEvaluationRequestAttributes struct {
	// The environment of the deployment.
	Env string `json:"env"`
	// The identifier of the deployment gate. Defaults to "default".
	Identifier *string `json:"identifier,omitempty"`
	// A primary tag to scope APM Faulty Deployment Detection rules.
	PrimaryTag *string `json:"primary_tag,omitempty"`
	// The service being deployed.
	Service string `json:"service"`
	// The version of the deployment. Required for APM Faulty Deployment Detection rules.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesEvaluationRequestAttributes instantiates a new DeploymentGatesEvaluationRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesEvaluationRequestAttributes(env string, service string) *DeploymentGatesEvaluationRequestAttributes {
	this := DeploymentGatesEvaluationRequestAttributes{}
	this.Env = env
	var identifier string = "default"
	this.Identifier = &identifier
	this.Service = service
	return &this
}

// NewDeploymentGatesEvaluationRequestAttributesWithDefaults instantiates a new DeploymentGatesEvaluationRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesEvaluationRequestAttributesWithDefaults() *DeploymentGatesEvaluationRequestAttributes {
	this := DeploymentGatesEvaluationRequestAttributes{}
	var identifier string = "default"
	this.Identifier = &identifier
	return &this
}

// GetEnv returns the Env field value.
func (o *DeploymentGatesEvaluationRequestAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *DeploymentGatesEvaluationRequestAttributes) SetEnv(v string) {
	o.Env = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *DeploymentGatesEvaluationRequestAttributes) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) HasIdentifier() bool {
	return o != nil && o.Identifier != nil
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *DeploymentGatesEvaluationRequestAttributes) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetPrimaryTag returns the PrimaryTag field value if set, zero value otherwise.
func (o *DeploymentGatesEvaluationRequestAttributes) GetPrimaryTag() string {
	if o == nil || o.PrimaryTag == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTag
}

// GetPrimaryTagOk returns a tuple with the PrimaryTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) GetPrimaryTagOk() (*string, bool) {
	if o == nil || o.PrimaryTag == nil {
		return nil, false
	}
	return o.PrimaryTag, true
}

// HasPrimaryTag returns a boolean if a field has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) HasPrimaryTag() bool {
	return o != nil && o.PrimaryTag != nil
}

// SetPrimaryTag gets a reference to the given string and assigns it to the PrimaryTag field.
func (o *DeploymentGatesEvaluationRequestAttributes) SetPrimaryTag(v string) {
	o.PrimaryTag = &v
}

// GetService returns the Service field value.
func (o *DeploymentGatesEvaluationRequestAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DeploymentGatesEvaluationRequestAttributes) SetService(v string) {
	o.Service = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeploymentGatesEvaluationRequestAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeploymentGatesEvaluationRequestAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeploymentGatesEvaluationRequestAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesEvaluationRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["env"] = o.Env
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.PrimaryTag != nil {
		toSerialize["primary_tag"] = o.PrimaryTag
	}
	toSerialize["service"] = o.Service
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesEvaluationRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Env        *string `json:"env"`
		Identifier *string `json:"identifier,omitempty"`
		PrimaryTag *string `json:"primary_tag,omitempty"`
		Service    *string `json:"service"`
		Version    *string `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"env", "identifier", "primary_tag", "service", "version"})
	} else {
		return err
	}
	o.Env = *all.Env
	o.Identifier = all.Identifier
	o.PrimaryTag = all.PrimaryTag
	o.Service = *all.Service
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
