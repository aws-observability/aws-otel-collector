// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionEnvEnv The definition of `ConnectionEnvEnv` object.
type ConnectionEnvEnv string

// List of ConnectionEnvEnv.
const (
	CONNECTIONENVENV_DEFAULT ConnectionEnvEnv = "default"
)

var allowedConnectionEnvEnvEnumValues = []ConnectionEnvEnv{
	CONNECTIONENVENV_DEFAULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConnectionEnvEnv) GetAllowedValues() []ConnectionEnvEnv {
	return allowedConnectionEnvEnvEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConnectionEnvEnv) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConnectionEnvEnv(value)
	return nil
}

// NewConnectionEnvEnvFromValue returns a pointer to a valid ConnectionEnvEnv
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConnectionEnvEnvFromValue(v string) (*ConnectionEnvEnv, error) {
	ev := ConnectionEnvEnv(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConnectionEnvEnv: valid values are %v", v, allowedConnectionEnvEnvEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConnectionEnvEnv) IsValid() bool {
	for _, existing := range allowedConnectionEnvEnvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionEnvEnv value.
func (v ConnectionEnvEnv) Ptr() *ConnectionEnvEnv {
	return &v
}
