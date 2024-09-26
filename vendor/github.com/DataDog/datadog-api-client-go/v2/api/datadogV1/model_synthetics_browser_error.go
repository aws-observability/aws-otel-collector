// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBrowserError Error response object for a browser test.
type SyntheticsBrowserError struct {
	// Description of the error.
	Description string `json:"description"`
	// Name of the error.
	Name string `json:"name"`
	// Status Code of the error.
	Status *int64 `json:"status,omitempty"`
	// Error type returned by a browser test.
	Type SyntheticsBrowserErrorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsBrowserError instantiates a new SyntheticsBrowserError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsBrowserError(description string, name string, typeVar SyntheticsBrowserErrorType) *SyntheticsBrowserError {
	this := SyntheticsBrowserError{}
	this.Description = description
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewSyntheticsBrowserErrorWithDefaults instantiates a new SyntheticsBrowserError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsBrowserErrorWithDefaults() *SyntheticsBrowserError {
	this := SyntheticsBrowserError{}
	return &this
}

// GetDescription returns the Description field value.
func (o *SyntheticsBrowserError) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SyntheticsBrowserError) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *SyntheticsBrowserError) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SyntheticsBrowserError) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsBrowserError) GetStatus() int64 {
	if o == nil || o.Status == nil {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetStatusOk() (*int64, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsBrowserError) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *SyntheticsBrowserError) SetStatus(v int64) {
	o.Status = &v
}

// GetType returns the Type field value.
func (o *SyntheticsBrowserError) GetType() SyntheticsBrowserErrorType {
	if o == nil {
		var ret SyntheticsBrowserErrorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserError) GetTypeOk() (*SyntheticsBrowserErrorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SyntheticsBrowserError) SetType(v SyntheticsBrowserErrorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsBrowserError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsBrowserError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                     `json:"description"`
		Name        *string                     `json:"name"`
		Status      *int64                      `json:"status,omitempty"`
		Type        *SyntheticsBrowserErrorType `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "status", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Name = *all.Name
	o.Status = all.Status
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
