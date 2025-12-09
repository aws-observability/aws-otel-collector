// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConfigReadAttributes Readable attributes of an Org Config.
type OrgConfigReadAttributes struct {
	// The description of an Org Config.
	Description string `json:"description"`
	// The timestamp of the last Org Config update (if any).
	ModifiedAt datadog.NullableTime `json:"modified_at,omitempty"`
	// The machine-friendly name of an Org Config.
	Name string `json:"name"`
	// The value of an Org Config.
	Value interface{} `json:"value"`
	// The type of an Org Config value.
	ValueType string `json:"value_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConfigReadAttributes instantiates a new OrgConfigReadAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConfigReadAttributes(description string, name string, value interface{}, valueType string) *OrgConfigReadAttributes {
	this := OrgConfigReadAttributes{}
	this.Description = description
	this.Name = name
	this.Value = value
	this.ValueType = valueType
	return &this
}

// NewOrgConfigReadAttributesWithDefaults instantiates a new OrgConfigReadAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConfigReadAttributesWithDefaults() *OrgConfigReadAttributes {
	this := OrgConfigReadAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *OrgConfigReadAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrgConfigReadAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *OrgConfigReadAttributes) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgConfigReadAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OrgConfigReadAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *OrgConfigReadAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt.IsSet()
}

// SetModifiedAt gets a reference to the given datadog.NullableTime and assigns it to the ModifiedAt field.
func (o *OrgConfigReadAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil.
func (o *OrgConfigReadAttributes) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil.
func (o *OrgConfigReadAttributes) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetName returns the Name field value.
func (o *OrgConfigReadAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgConfigReadAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrgConfigReadAttributes) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *OrgConfigReadAttributes) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OrgConfigReadAttributes) GetValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *OrgConfigReadAttributes) SetValue(v interface{}) {
	o.Value = v
}

// GetValueType returns the ValueType field value.
func (o *OrgConfigReadAttributes) GetValueType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *OrgConfigReadAttributes) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value.
func (o *OrgConfigReadAttributes) SetValueType(v string) {
	o.ValueType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConfigReadAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["value_type"] = o.ValueType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConfigReadAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string              `json:"description"`
		ModifiedAt  datadog.NullableTime `json:"modified_at,omitempty"`
		Name        *string              `json:"name"`
		Value       *interface{}         `json:"value"`
		ValueType   *string              `json:"value_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.ValueType == nil {
		return fmt.Errorf("required field value_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "modified_at", "name", "value", "value_type"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	o.Value = *all.Value
	o.ValueType = *all.ValueType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
