// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomFrameworkControl Framework Control.
type CustomFrameworkControl struct {
	// Control Name.
	Name string `json:"name"`
	// Rule IDs.
	RulesId []string `json:"rules_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomFrameworkControl instantiates a new CustomFrameworkControl object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomFrameworkControl(name string, rulesId []string) *CustomFrameworkControl {
	this := CustomFrameworkControl{}
	this.Name = name
	this.RulesId = rulesId
	return &this
}

// NewCustomFrameworkControlWithDefaults instantiates a new CustomFrameworkControl object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomFrameworkControlWithDefaults() *CustomFrameworkControl {
	this := CustomFrameworkControl{}
	return &this
}

// GetName returns the Name field value.
func (o *CustomFrameworkControl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkControl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomFrameworkControl) SetName(v string) {
	o.Name = v
}

// GetRulesId returns the RulesId field value.
func (o *CustomFrameworkControl) GetRulesId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RulesId
}

// GetRulesIdOk returns a tuple with the RulesId field value
// and a boolean to check if the value has been set.
func (o *CustomFrameworkControl) GetRulesIdOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RulesId, true
}

// SetRulesId sets field value.
func (o *CustomFrameworkControl) SetRulesId(v []string) {
	o.RulesId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomFrameworkControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["rules_id"] = o.RulesId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomFrameworkControl) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string   `json:"name"`
		RulesId *[]string `json:"rules_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RulesId == nil {
		return fmt.Errorf("required field rules_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "rules_id"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.RulesId = *all.RulesId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
