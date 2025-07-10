// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPTokenUpdate The definition of `HTTPTokenUpdate` object.
type HTTPTokenUpdate struct {
	// Should the header be deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// The `HTTPToken` `name`.
	Name string `json:"name"`
	// The definition of `TokenType` object.
	Type TokenType `json:"type"`
	// The `HTTPToken` `value`.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPTokenUpdate instantiates a new HTTPTokenUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPTokenUpdate(name string, typeVar TokenType, value string) *HTTPTokenUpdate {
	this := HTTPTokenUpdate{}
	this.Name = name
	this.Type = typeVar
	this.Value = value
	return &this
}

// NewHTTPTokenUpdateWithDefaults instantiates a new HTTPTokenUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPTokenUpdateWithDefaults() *HTTPTokenUpdate {
	this := HTTPTokenUpdate{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *HTTPTokenUpdate) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenUpdate) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *HTTPTokenUpdate) HasDeleted() bool {
	return o != nil && o.Deleted != nil
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *HTTPTokenUpdate) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetName returns the Name field value.
func (o *HTTPTokenUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HTTPTokenUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *HTTPTokenUpdate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *HTTPTokenUpdate) GetType() TokenType {
	if o == nil {
		var ret TokenType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPTokenUpdate) GetTypeOk() (*TokenType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPTokenUpdate) SetType(v TokenType) {
	o.Type = v
}

// GetValue returns the Value field value.
func (o *HTTPTokenUpdate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HTTPTokenUpdate) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *HTTPTokenUpdate) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPTokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPTokenUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deleted *bool      `json:"deleted,omitempty"`
		Name    *string    `json:"name"`
		Type    *TokenType `json:"type"`
		Value   *string    `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted", "name", "type", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Deleted = all.Deleted
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
