// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldValidValue A valid value for an incident user-defined field.
type IncidentUserDefinedFieldValidValue struct {
	// A detailed description of the valid value.
	Description *string `json:"description,omitempty"`
	// The human-readable display name for this value.
	DisplayName string `json:"display_name"`
	// A short description of the valid value.
	ShortDescription *string `json:"short_description,omitempty"`
	// The identifier that is stored when this option is selected.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldValidValue instantiates a new IncidentUserDefinedFieldValidValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldValidValue(displayName string, value string) *IncidentUserDefinedFieldValidValue {
	this := IncidentUserDefinedFieldValidValue{}
	this.DisplayName = displayName
	this.Value = value
	return &this
}

// NewIncidentUserDefinedFieldValidValueWithDefaults instantiates a new IncidentUserDefinedFieldValidValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldValidValueWithDefaults() *IncidentUserDefinedFieldValidValue {
	this := IncidentUserDefinedFieldValidValue{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldValidValue) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldValidValue) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldValidValue) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IncidentUserDefinedFieldValidValue) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value.
func (o *IncidentUserDefinedFieldValidValue) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldValidValue) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *IncidentUserDefinedFieldValidValue) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldValidValue) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldValidValue) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldValidValue) HasShortDescription() bool {
	return o != nil && o.ShortDescription != nil
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *IncidentUserDefinedFieldValidValue) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetValue returns the Value field value.
func (o *IncidentUserDefinedFieldValidValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldValidValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *IncidentUserDefinedFieldValidValue) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldValidValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["display_name"] = o.DisplayName
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldValidValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string `json:"description,omitempty"`
		DisplayName      *string `json:"display_name"`
		ShortDescription *string `json:"short_description,omitempty"`
		Value            *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "short_description", "value"})
	} else {
		return err
	}
	o.Description = all.Description
	o.DisplayName = *all.DisplayName
	o.ShortDescription = all.ShortDescription
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
