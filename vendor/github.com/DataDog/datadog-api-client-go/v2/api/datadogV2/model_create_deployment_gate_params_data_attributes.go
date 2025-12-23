// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDeploymentGateParamsDataAttributes Parameters for creating a deployment gate.
type CreateDeploymentGateParamsDataAttributes struct {
	// Whether this gate is run in dry-run mode.
	DryRun *bool `json:"dry_run,omitempty"`
	// The environment of the deployment gate.
	Env string `json:"env"`
	// The identifier of the deployment gate.
	Identifier *string `json:"identifier,omitempty"`
	// The service of the deployment gate.
	Service string `json:"service"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateDeploymentGateParamsDataAttributes instantiates a new CreateDeploymentGateParamsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateDeploymentGateParamsDataAttributes(env string, service string) *CreateDeploymentGateParamsDataAttributes {
	this := CreateDeploymentGateParamsDataAttributes{}
	var dryRun bool = false
	this.DryRun = &dryRun
	this.Env = env
	var identifier string = "default"
	this.Identifier = &identifier
	this.Service = service
	return &this
}

// NewCreateDeploymentGateParamsDataAttributesWithDefaults instantiates a new CreateDeploymentGateParamsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateDeploymentGateParamsDataAttributesWithDefaults() *CreateDeploymentGateParamsDataAttributes {
	this := CreateDeploymentGateParamsDataAttributes{}
	var dryRun bool = false
	this.DryRun = &dryRun
	var identifier string = "default"
	this.Identifier = &identifier
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateDeploymentGateParamsDataAttributes) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeploymentGateParamsDataAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateDeploymentGateParamsDataAttributes) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateDeploymentGateParamsDataAttributes) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEnv returns the Env field value.
func (o *CreateDeploymentGateParamsDataAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *CreateDeploymentGateParamsDataAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *CreateDeploymentGateParamsDataAttributes) SetEnv(v string) {
	o.Env = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CreateDeploymentGateParamsDataAttributes) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeploymentGateParamsDataAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CreateDeploymentGateParamsDataAttributes) HasIdentifier() bool {
	return o != nil && o.Identifier != nil
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CreateDeploymentGateParamsDataAttributes) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetService returns the Service field value.
func (o *CreateDeploymentGateParamsDataAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *CreateDeploymentGateParamsDataAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *CreateDeploymentGateParamsDataAttributes) SetService(v string) {
	o.Service = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateDeploymentGateParamsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	toSerialize["env"] = o.Env
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	toSerialize["service"] = o.Service

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateDeploymentGateParamsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun     *bool   `json:"dry_run,omitempty"`
		Env        *string `json:"env"`
		Identifier *string `json:"identifier,omitempty"`
		Service    *string `json:"service"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"dry_run", "env", "identifier", "service"})
	} else {
		return err
	}
	o.DryRun = all.DryRun
	o.Env = *all.Env
	o.Identifier = all.Identifier
	o.Service = *all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
