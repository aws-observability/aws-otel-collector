// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseUpdatePriorityAttributes Case update priority attributes
type CaseUpdatePriorityAttributes struct {
	// Case priority
	Priority CasePriority `json:"priority"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseUpdatePriorityAttributes instantiates a new CaseUpdatePriorityAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseUpdatePriorityAttributes(priority CasePriority) *CaseUpdatePriorityAttributes {
	this := CaseUpdatePriorityAttributes{}
	this.Priority = priority
	return &this
}

// NewCaseUpdatePriorityAttributesWithDefaults instantiates a new CaseUpdatePriorityAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseUpdatePriorityAttributesWithDefaults() *CaseUpdatePriorityAttributes {
	this := CaseUpdatePriorityAttributes{}
	var priority CasePriority = CASEPRIORITY_NOT_DEFINED
	this.Priority = priority
	return &this
}

// GetPriority returns the Priority field value.
func (o *CaseUpdatePriorityAttributes) GetPriority() CasePriority {
	if o == nil {
		var ret CasePriority
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CaseUpdatePriorityAttributes) GetPriorityOk() (*CasePriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *CaseUpdatePriorityAttributes) SetPriority(v CasePriority) {
	o.Priority = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseUpdatePriorityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["priority"] = o.Priority

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseUpdatePriorityAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Priority *CasePriority `json:"priority"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"priority"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = *all.Priority
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
