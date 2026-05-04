// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleAttributesRequest Incident handle attributes for requests
type IncidentHandleAttributesRequest struct {
	// Dynamic fields associated with the handle
	Fields *IncidentHandleAttributesFields `json:"fields,omitempty"`
	// The handle name
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleAttributesRequest instantiates a new IncidentHandleAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleAttributesRequest(name string) *IncidentHandleAttributesRequest {
	this := IncidentHandleAttributesRequest{}
	this.Name = name
	return &this
}

// NewIncidentHandleAttributesRequestWithDefaults instantiates a new IncidentHandleAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleAttributesRequestWithDefaults() *IncidentHandleAttributesRequest {
	this := IncidentHandleAttributesRequest{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentHandleAttributesRequest) GetFields() IncidentHandleAttributesFields {
	if o == nil || o.Fields == nil {
		var ret IncidentHandleAttributesFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesRequest) GetFieldsOk() (*IncidentHandleAttributesFields, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentHandleAttributesRequest) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given IncidentHandleAttributesFields and assigns it to the Fields field.
func (o *IncidentHandleAttributesRequest) SetFields(v IncidentHandleAttributesFields) {
	o.Fields = &v
}

// GetName returns the Name field value.
func (o *IncidentHandleAttributesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentHandleAttributesRequest) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields *IncidentHandleAttributesFields `json:"fields,omitempty"`
		Name   *string                         `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Fields != nil && all.Fields.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Fields = all.Fields
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
