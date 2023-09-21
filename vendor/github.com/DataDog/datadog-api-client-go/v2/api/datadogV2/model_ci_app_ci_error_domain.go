// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// CIAppCIErrorDomain Error category used to differentiate between issues related to the developer or provider environments.
type CIAppCIErrorDomain string

// List of CIAppCIErrorDomain.
const (
	CIAPPCIERRORDOMAIN_PROVIDER CIAppCIErrorDomain = "provider"
	CIAPPCIERRORDOMAIN_USER     CIAppCIErrorDomain = "user"
	CIAPPCIERRORDOMAIN_UNKNOWN  CIAppCIErrorDomain = "unknown"
)

var allowedCIAppCIErrorDomainEnumValues = []CIAppCIErrorDomain{
	CIAPPCIERRORDOMAIN_PROVIDER,
	CIAPPCIERRORDOMAIN_USER,
	CIAPPCIERRORDOMAIN_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppCIErrorDomain) GetAllowedValues() []CIAppCIErrorDomain {
	return allowedCIAppCIErrorDomainEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppCIErrorDomain) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppCIErrorDomain(value)
	return nil
}

// NewCIAppCIErrorDomainFromValue returns a pointer to a valid CIAppCIErrorDomain
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppCIErrorDomainFromValue(v string) (*CIAppCIErrorDomain, error) {
	ev := CIAppCIErrorDomain(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppCIErrorDomain: valid values are %v", v, allowedCIAppCIErrorDomainEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppCIErrorDomain) IsValid() bool {
	for _, existing := range allowedCIAppCIErrorDomainEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppCIErrorDomain value.
func (v CIAppCIErrorDomain) Ptr() *CIAppCIErrorDomain {
	return &v
}
<<<<<<< HEAD

// NullableCIAppCIErrorDomain handles when a null is used for CIAppCIErrorDomain.
type NullableCIAppCIErrorDomain struct {
	value *CIAppCIErrorDomain
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppCIErrorDomain) Get() *CIAppCIErrorDomain {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppCIErrorDomain) Set(val *CIAppCIErrorDomain) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppCIErrorDomain) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableCIAppCIErrorDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppCIErrorDomain initializes the struct as if Set has been called.
func NewNullableCIAppCIErrorDomain(val *CIAppCIErrorDomain) *NullableCIAppCIErrorDomain {
	return &NullableCIAppCIErrorDomain{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppCIErrorDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppCIErrorDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
