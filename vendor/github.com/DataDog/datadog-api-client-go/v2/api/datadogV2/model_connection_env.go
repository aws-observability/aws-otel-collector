// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionEnv A list of connections or connection groups used in the workflow.
type ConnectionEnv struct {
	// The `ConnectionEnv` `connectionGroups`.
	ConnectionGroups []ConnectionGroup `json:"connectionGroups,omitempty"`
	// The `ConnectionEnv` `connections`.
	Connections []Connection `json:"connections,omitempty"`
	// The definition of `ConnectionEnvEnv` object.
	Env ConnectionEnvEnv `json:"env"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConnectionEnv instantiates a new ConnectionEnv object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConnectionEnv(env ConnectionEnvEnv) *ConnectionEnv {
	this := ConnectionEnv{}
	this.Env = env
	return &this
}

// NewConnectionEnvWithDefaults instantiates a new ConnectionEnv object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConnectionEnvWithDefaults() *ConnectionEnv {
	this := ConnectionEnv{}
	return &this
}

// GetConnectionGroups returns the ConnectionGroups field value if set, zero value otherwise.
func (o *ConnectionEnv) GetConnectionGroups() []ConnectionGroup {
	if o == nil || o.ConnectionGroups == nil {
		var ret []ConnectionGroup
		return ret
	}
	return o.ConnectionGroups
}

// GetConnectionGroupsOk returns a tuple with the ConnectionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEnv) GetConnectionGroupsOk() (*[]ConnectionGroup, bool) {
	if o == nil || o.ConnectionGroups == nil {
		return nil, false
	}
	return &o.ConnectionGroups, true
}

// HasConnectionGroups returns a boolean if a field has been set.
func (o *ConnectionEnv) HasConnectionGroups() bool {
	return o != nil && o.ConnectionGroups != nil
}

// SetConnectionGroups gets a reference to the given []ConnectionGroup and assigns it to the ConnectionGroups field.
func (o *ConnectionEnv) SetConnectionGroups(v []ConnectionGroup) {
	o.ConnectionGroups = v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ConnectionEnv) GetConnections() []Connection {
	if o == nil || o.Connections == nil {
		var ret []Connection
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEnv) GetConnectionsOk() (*[]Connection, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ConnectionEnv) HasConnections() bool {
	return o != nil && o.Connections != nil
}

// SetConnections gets a reference to the given []Connection and assigns it to the Connections field.
func (o *ConnectionEnv) SetConnections(v []Connection) {
	o.Connections = v
}

// GetEnv returns the Env field value.
func (o *ConnectionEnv) GetEnv() ConnectionEnvEnv {
	if o == nil {
		var ret ConnectionEnvEnv
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ConnectionEnv) GetEnvOk() (*ConnectionEnvEnv, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ConnectionEnv) SetEnv(v ConnectionEnvEnv) {
	o.Env = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConnectionEnv) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConnectionGroups != nil {
		toSerialize["connectionGroups"] = o.ConnectionGroups
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	toSerialize["env"] = o.Env

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConnectionEnv) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionGroups []ConnectionGroup `json:"connectionGroups,omitempty"`
		Connections      []Connection      `json:"connections,omitempty"`
		Env              *ConnectionEnvEnv `json:"env"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connectionGroups", "connections", "env"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConnectionGroups = all.ConnectionGroups
	o.Connections = all.Connections
	if !all.Env.IsValid() {
		hasInvalidField = true
	} else {
		o.Env = *all.Env
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
