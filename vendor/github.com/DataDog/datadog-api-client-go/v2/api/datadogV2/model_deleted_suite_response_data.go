// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedSuiteResponseData Data object for a deleted Synthetic test suite.
type DeletedSuiteResponseData struct {
	// Attributes of a deleted Synthetic test suite, including deletion timestamp and public ID.
	Attributes *DeletedSuiteResponseDataAttributes `json:"attributes,omitempty"`
	// The public ID of the deleted Synthetic test suite.
	Id *string `json:"id,omitempty"`
	// Type for the Synthetics suites responses, `suites`.
	Type *SyntheticsSuiteTypes `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeletedSuiteResponseData instantiates a new DeletedSuiteResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeletedSuiteResponseData() *DeletedSuiteResponseData {
	this := DeletedSuiteResponseData{}
	var typeVar SyntheticsSuiteTypes = SYNTHETICSSUITETYPES_SUITES
	this.Type = &typeVar
	return &this
}

// NewDeletedSuiteResponseDataWithDefaults instantiates a new DeletedSuiteResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeletedSuiteResponseDataWithDefaults() *DeletedSuiteResponseData {
	this := DeletedSuiteResponseData{}
	var typeVar SyntheticsSuiteTypes = SYNTHETICSSUITETYPES_SUITES
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DeletedSuiteResponseData) GetAttributes() DeletedSuiteResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret DeletedSuiteResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedSuiteResponseData) GetAttributesOk() (*DeletedSuiteResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DeletedSuiteResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DeletedSuiteResponseDataAttributes and assigns it to the Attributes field.
func (o *DeletedSuiteResponseData) SetAttributes(v DeletedSuiteResponseDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeletedSuiteResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedSuiteResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeletedSuiteResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeletedSuiteResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeletedSuiteResponseData) GetType() SyntheticsSuiteTypes {
	if o == nil || o.Type == nil {
		var ret SyntheticsSuiteTypes
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedSuiteResponseData) GetTypeOk() (*SyntheticsSuiteTypes, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeletedSuiteResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsSuiteTypes and assigns it to the Type field.
func (o *DeletedSuiteResponseData) SetType(v SyntheticsSuiteTypes) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeletedSuiteResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeletedSuiteResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DeletedSuiteResponseDataAttributes `json:"attributes,omitempty"`
		Id         *string                             `json:"id,omitempty"`
		Type       *SyntheticsSuiteTypes               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
