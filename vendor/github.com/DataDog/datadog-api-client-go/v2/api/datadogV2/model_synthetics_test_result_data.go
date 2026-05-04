// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultData Wrapper object for a Synthetic test result.
type SyntheticsTestResultData struct {
	// Attributes of a Synthetic test result.
	Attributes *SyntheticsTestResultAttributes `json:"attributes,omitempty"`
	// The result ID.
	Id *string `json:"id,omitempty"`
	// Relationships for a Synthetic test result.
	Relationships *SyntheticsTestResultRelationships `json:"relationships,omitempty"`
	// Type of the Synthetic test result resource, `result`.
	Type *SyntheticsTestResultType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultData instantiates a new SyntheticsTestResultData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultData() *SyntheticsTestResultData {
	this := SyntheticsTestResultData{}
	var typeVar SyntheticsTestResultType = SYNTHETICSTESTRESULTTYPE_RESULT
	this.Type = &typeVar
	return &this
}

// NewSyntheticsTestResultDataWithDefaults instantiates a new SyntheticsTestResultData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDataWithDefaults() *SyntheticsTestResultData {
	this := SyntheticsTestResultData{}
	var typeVar SyntheticsTestResultType = SYNTHETICSTESTRESULTTYPE_RESULT
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SyntheticsTestResultData) GetAttributes() SyntheticsTestResultAttributes {
	if o == nil || o.Attributes == nil {
		var ret SyntheticsTestResultAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultData) GetAttributesOk() (*SyntheticsTestResultAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SyntheticsTestResultData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SyntheticsTestResultAttributes and assigns it to the Attributes field.
func (o *SyntheticsTestResultData) SetAttributes(v SyntheticsTestResultAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SyntheticsTestResultData) GetRelationships() SyntheticsTestResultRelationships {
	if o == nil || o.Relationships == nil {
		var ret SyntheticsTestResultRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultData) GetRelationshipsOk() (*SyntheticsTestResultRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SyntheticsTestResultData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given SyntheticsTestResultRelationships and assigns it to the Relationships field.
func (o *SyntheticsTestResultData) SetRelationships(v SyntheticsTestResultRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultData) GetType() SyntheticsTestResultType {
	if o == nil || o.Type == nil {
		var ret SyntheticsTestResultType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultData) GetTypeOk() (*SyntheticsTestResultType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SyntheticsTestResultType and assigns it to the Type field.
func (o *SyntheticsTestResultData) SetType(v SyntheticsTestResultType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultData) MarshalJSON() ([]byte, error) {
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
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
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
func (o *SyntheticsTestResultData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *SyntheticsTestResultAttributes    `json:"attributes,omitempty"`
		Id            *string                            `json:"id,omitempty"`
		Relationships *SyntheticsTestResultRelationships `json:"relationships,omitempty"`
		Type          *SyntheticsTestResultType          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
