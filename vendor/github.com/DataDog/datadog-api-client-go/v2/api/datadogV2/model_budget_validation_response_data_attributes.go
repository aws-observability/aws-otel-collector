// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetValidationResponseDataAttributes The attributes of a budget validation response, including any validation errors and the validity status.
type BudgetValidationResponseDataAttributes struct {
	// A list of validation error messages for the budget.
	Errors []string `json:"errors,omitempty"`
	// Whether the budget configuration is valid.
	Valid *bool `json:"valid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetValidationResponseDataAttributes instantiates a new BudgetValidationResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetValidationResponseDataAttributes() *BudgetValidationResponseDataAttributes {
	this := BudgetValidationResponseDataAttributes{}
	return &this
}

// NewBudgetValidationResponseDataAttributesWithDefaults instantiates a new BudgetValidationResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetValidationResponseDataAttributesWithDefaults() *BudgetValidationResponseDataAttributes {
	this := BudgetValidationResponseDataAttributes{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BudgetValidationResponseDataAttributes) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetValidationResponseDataAttributes) GetErrorsOk() (*[]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BudgetValidationResponseDataAttributes) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *BudgetValidationResponseDataAttributes) SetErrors(v []string) {
	o.Errors = v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BudgetValidationResponseDataAttributes) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetValidationResponseDataAttributes) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BudgetValidationResponseDataAttributes) HasValid() bool {
	return o != nil && o.Valid != nil
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BudgetValidationResponseDataAttributes) SetValid(v bool) {
	o.Valid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetValidationResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BudgetValidationResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Errors []string `json:"errors,omitempty"`
		Valid  *bool    `json:"valid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"errors", "valid"})
	} else {
		return err
	}
	o.Errors = all.Errors
	o.Valid = all.Valid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
