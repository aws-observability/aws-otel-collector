// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryCreateData Data related to the creation of a restriction query.
type RestrictionQueryCreateData struct {
	// Attributes of the created restriction query.
	Attributes *RestrictionQueryCreateAttributes `json:"attributes,omitempty"`
	// Restriction query resource type.
	Type *LogsRestrictionQueriesType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionQueryCreateData instantiates a new RestrictionQueryCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionQueryCreateData() *RestrictionQueryCreateData {
	this := RestrictionQueryCreateData{}
	var typeVar LogsRestrictionQueriesType = LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES
	this.Type = &typeVar
	return &this
}

// NewRestrictionQueryCreateDataWithDefaults instantiates a new RestrictionQueryCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionQueryCreateDataWithDefaults() *RestrictionQueryCreateData {
	this := RestrictionQueryCreateData{}
	var typeVar LogsRestrictionQueriesType = LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RestrictionQueryCreateData) GetAttributes() RestrictionQueryCreateAttributes {
	if o == nil || o.Attributes == nil {
		var ret RestrictionQueryCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryCreateData) GetAttributesOk() (*RestrictionQueryCreateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RestrictionQueryCreateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RestrictionQueryCreateAttributes and assigns it to the Attributes field.
func (o *RestrictionQueryCreateData) SetAttributes(v RestrictionQueryCreateAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestrictionQueryCreateData) GetType() LogsRestrictionQueriesType {
	if o == nil || o.Type == nil {
		var ret LogsRestrictionQueriesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryCreateData) GetTypeOk() (*LogsRestrictionQueriesType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestrictionQueryCreateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LogsRestrictionQueriesType and assigns it to the Type field.
func (o *RestrictionQueryCreateData) SetType(v LogsRestrictionQueriesType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionQueryCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *RestrictionQueryCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RestrictionQueryCreateAttributes `json:"attributes,omitempty"`
		Type       *LogsRestrictionQueriesType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
