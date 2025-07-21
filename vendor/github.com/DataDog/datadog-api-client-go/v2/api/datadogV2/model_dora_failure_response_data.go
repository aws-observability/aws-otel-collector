// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAFailureResponseData Response after receiving a DORA failure event.
type DORAFailureResponseData struct {
	// The ID of the received DORA failure event.
	Id string `json:"id"`
	// JSON:API type for DORA failure events.
	Type *DORAFailureType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORAFailureResponseData instantiates a new DORAFailureResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAFailureResponseData(id string) *DORAFailureResponseData {
	this := DORAFailureResponseData{}
	this.Id = id
	var typeVar DORAFailureType = DORAFAILURETYPE_DORA_FAILURE
	this.Type = &typeVar
	return &this
}

// NewDORAFailureResponseDataWithDefaults instantiates a new DORAFailureResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAFailureResponseDataWithDefaults() *DORAFailureResponseData {
	this := DORAFailureResponseData{}
	var typeVar DORAFailureType = DORAFAILURETYPE_DORA_FAILURE
	this.Type = &typeVar
	return &this
}

// GetId returns the Id field value.
func (o *DORAFailureResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DORAFailureResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DORAFailureResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DORAFailureResponseData) GetType() DORAFailureType {
	if o == nil || o.Type == nil {
		var ret DORAFailureType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAFailureResponseData) GetTypeOk() (*DORAFailureType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DORAFailureResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DORAFailureType and assigns it to the Type field.
func (o *DORAFailureResponseData) SetType(v DORAFailureType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAFailureResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORAFailureResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string          `json:"id"`
		Type *DORAFailureType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
