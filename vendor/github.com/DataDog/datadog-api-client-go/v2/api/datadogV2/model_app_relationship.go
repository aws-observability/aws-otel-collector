// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppRelationship The app's publication relationship and custom connections.
type AppRelationship struct {
	// Array of custom connections used by the app.
	Connections []CustomConnection `json:"connections,omitempty"`
	// Information pointing to the app's publication status.
	Deployment *DeploymentRelationship `json:"deployment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAppRelationship instantiates a new AppRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAppRelationship() *AppRelationship {
	this := AppRelationship{}
	return &this
}

// NewAppRelationshipWithDefaults instantiates a new AppRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAppRelationshipWithDefaults() *AppRelationship {
	this := AppRelationship{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *AppRelationship) GetConnections() []CustomConnection {
	if o == nil || o.Connections == nil {
		var ret []CustomConnection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationship) GetConnectionsOk() (*[]CustomConnection, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *AppRelationship) HasConnections() bool {
	return o != nil && o.Connections != nil
}

// SetConnections gets a reference to the given []CustomConnection and assigns it to the Connections field.
func (o *AppRelationship) SetConnections(v []CustomConnection) {
	o.Connections = v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *AppRelationship) GetDeployment() DeploymentRelationship {
	if o == nil || o.Deployment == nil {
		var ret DeploymentRelationship
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationship) GetDeploymentOk() (*DeploymentRelationship, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *AppRelationship) HasDeployment() bool {
	return o != nil && o.Deployment != nil
}

// SetDeployment gets a reference to the given DeploymentRelationship and assigns it to the Deployment field.
func (o *AppRelationship) SetDeployment(v DeploymentRelationship) {
	o.Deployment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AppRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AppRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connections []CustomConnection      `json:"connections,omitempty"`
		Deployment  *DeploymentRelationship `json:"deployment,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connections", "deployment"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Connections = all.Connections
	if all.Deployment != nil && all.Deployment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Deployment = all.Deployment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
