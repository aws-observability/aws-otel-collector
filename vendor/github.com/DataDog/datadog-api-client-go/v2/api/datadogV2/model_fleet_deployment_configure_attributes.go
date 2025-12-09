// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentConfigureAttributes Attributes for creating a new configuration deployment.
type FleetDeploymentConfigureAttributes struct {
	// Ordered list of configuration file operations to perform on the target hosts.
	ConfigOperations []FleetDeploymentOperation `json:"config_operations"`
	// Query used to filter and select target hosts for the deployment. Uses the Datadog query syntax.
	FilterQuery *string `json:"filter_query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentConfigureAttributes instantiates a new FleetDeploymentConfigureAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentConfigureAttributes(configOperations []FleetDeploymentOperation) *FleetDeploymentConfigureAttributes {
	this := FleetDeploymentConfigureAttributes{}
	this.ConfigOperations = configOperations
	return &this
}

// NewFleetDeploymentConfigureAttributesWithDefaults instantiates a new FleetDeploymentConfigureAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentConfigureAttributesWithDefaults() *FleetDeploymentConfigureAttributes {
	this := FleetDeploymentConfigureAttributes{}
	return &this
}

// GetConfigOperations returns the ConfigOperations field value.
func (o *FleetDeploymentConfigureAttributes) GetConfigOperations() []FleetDeploymentOperation {
	if o == nil {
		var ret []FleetDeploymentOperation
		return ret
	}
	return o.ConfigOperations
}

// GetConfigOperationsOk returns a tuple with the ConfigOperations field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureAttributes) GetConfigOperationsOk() (*[]FleetDeploymentOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigOperations, true
}

// SetConfigOperations sets field value.
func (o *FleetDeploymentConfigureAttributes) SetConfigOperations(v []FleetDeploymentOperation) {
	o.ConfigOperations = v
}

// GetFilterQuery returns the FilterQuery field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureAttributes) GetFilterQuery() string {
	if o == nil || o.FilterQuery == nil {
		var ret string
		return ret
	}
	return *o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureAttributes) GetFilterQueryOk() (*string, bool) {
	if o == nil || o.FilterQuery == nil {
		return nil, false
	}
	return o.FilterQuery, true
}

// HasFilterQuery returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureAttributes) HasFilterQuery() bool {
	return o != nil && o.FilterQuery != nil
}

// SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.
func (o *FleetDeploymentConfigureAttributes) SetFilterQuery(v string) {
	o.FilterQuery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentConfigureAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config_operations"] = o.ConfigOperations
	if o.FilterQuery != nil {
		toSerialize["filter_query"] = o.FilterQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentConfigureAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigOperations *[]FleetDeploymentOperation `json:"config_operations"`
		FilterQuery      *string                     `json:"filter_query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigOperations == nil {
		return fmt.Errorf("required field config_operations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config_operations", "filter_query"})
	} else {
		return err
	}
	o.ConfigOperations = *all.ConfigOperations
	o.FilterQuery = all.FilterQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
