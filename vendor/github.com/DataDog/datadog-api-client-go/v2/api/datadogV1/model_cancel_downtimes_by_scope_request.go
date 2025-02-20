// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CancelDowntimesByScopeRequest Cancel downtimes according to scope.
type CancelDowntimesByScopeRequest struct {
	// The scope(s) to which the downtime applies and must be in `key:value` format. For example, `host:app2`.
	// Provide multiple scopes as a comma-separated list like `env:dev,env:prod`.
	// The resulting downtime applies to sources that matches ALL provided scopes (`env:dev` **AND** `env:prod`).
	Scope string `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCancelDowntimesByScopeRequest instantiates a new CancelDowntimesByScopeRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCancelDowntimesByScopeRequest(scope string) *CancelDowntimesByScopeRequest {
	this := CancelDowntimesByScopeRequest{}
	this.Scope = scope
	return &this
}

// NewCancelDowntimesByScopeRequestWithDefaults instantiates a new CancelDowntimesByScopeRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCancelDowntimesByScopeRequestWithDefaults() *CancelDowntimesByScopeRequest {
	this := CancelDowntimesByScopeRequest{}
	return &this
}

// GetScope returns the Scope field value.
func (o *CancelDowntimesByScopeRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CancelDowntimesByScopeRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *CancelDowntimesByScopeRequest) SetScope(v string) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CancelDowntimesByScopeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CancelDowntimesByScopeRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Scope *string `json:"scope"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"scope"})
	} else {
		return err
	}
	o.Scope = *all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
