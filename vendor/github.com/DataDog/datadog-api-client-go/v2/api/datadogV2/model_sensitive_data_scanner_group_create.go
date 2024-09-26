// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerGroupCreate Data related to the creation of a group.
type SensitiveDataScannerGroupCreate struct {
	// Attributes of the Sensitive Data Scanner group.
	Attributes SensitiveDataScannerGroupAttributes `json:"attributes"`
	// Relationships of the group.
	Relationships *SensitiveDataScannerGroupRelationships `json:"relationships,omitempty"`
	// Sensitive Data Scanner group type.
	Type SensitiveDataScannerGroupType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerGroupCreate instantiates a new SensitiveDataScannerGroupCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerGroupCreate(attributes SensitiveDataScannerGroupAttributes, typeVar SensitiveDataScannerGroupType) *SensitiveDataScannerGroupCreate {
	this := SensitiveDataScannerGroupCreate{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSensitiveDataScannerGroupCreateWithDefaults instantiates a new SensitiveDataScannerGroupCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerGroupCreateWithDefaults() *SensitiveDataScannerGroupCreate {
	this := SensitiveDataScannerGroupCreate{}
	var typeVar SensitiveDataScannerGroupType = SENSITIVEDATASCANNERGROUPTYPE_SENSITIVE_DATA_SCANNER_GROUP
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SensitiveDataScannerGroupCreate) GetAttributes() SensitiveDataScannerGroupAttributes {
	if o == nil {
		var ret SensitiveDataScannerGroupAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupCreate) GetAttributesOk() (*SensitiveDataScannerGroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SensitiveDataScannerGroupCreate) SetAttributes(v SensitiveDataScannerGroupAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SensitiveDataScannerGroupCreate) GetRelationships() SensitiveDataScannerGroupRelationships {
	if o == nil || o.Relationships == nil {
		var ret SensitiveDataScannerGroupRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupCreate) GetRelationshipsOk() (*SensitiveDataScannerGroupRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SensitiveDataScannerGroupCreate) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given SensitiveDataScannerGroupRelationships and assigns it to the Relationships field.
func (o *SensitiveDataScannerGroupCreate) SetRelationships(v SensitiveDataScannerGroupRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *SensitiveDataScannerGroupCreate) GetType() SensitiveDataScannerGroupType {
	if o == nil {
		var ret SensitiveDataScannerGroupType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerGroupCreate) GetTypeOk() (*SensitiveDataScannerGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SensitiveDataScannerGroupCreate) SetType(v SensitiveDataScannerGroupType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerGroupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *SensitiveDataScannerGroupAttributes    `json:"attributes"`
		Relationships *SensitiveDataScannerGroupRelationships `json:"relationships,omitempty"`
		Type          *SensitiveDataScannerGroupType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
