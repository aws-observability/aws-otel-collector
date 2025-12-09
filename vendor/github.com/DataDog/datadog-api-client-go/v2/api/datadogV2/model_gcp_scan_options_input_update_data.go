// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GcpScanOptionsInputUpdateData Data object for updating the scan options of a single GCP project.
type GcpScanOptionsInputUpdateData struct {
	// Attributes for updating GCP scan options configuration.
	Attributes *GcpScanOptionsInputUpdateDataAttributes `json:"attributes,omitempty"`
	// The GCP project ID.
	Id string `json:"id"`
	// GCP scan options resource type.
	Type GcpScanOptionsInputUpdateDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGcpScanOptionsInputUpdateData instantiates a new GcpScanOptionsInputUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGcpScanOptionsInputUpdateData(id string, typeVar GcpScanOptionsInputUpdateDataType) *GcpScanOptionsInputUpdateData {
	this := GcpScanOptionsInputUpdateData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewGcpScanOptionsInputUpdateDataWithDefaults instantiates a new GcpScanOptionsInputUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGcpScanOptionsInputUpdateDataWithDefaults() *GcpScanOptionsInputUpdateData {
	this := GcpScanOptionsInputUpdateData{}
	var typeVar GcpScanOptionsInputUpdateDataType = GCPSCANOPTIONSINPUTUPDATEDATATYPE_GCP_SCAN_OPTIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GcpScanOptionsInputUpdateData) GetAttributes() GcpScanOptionsInputUpdateDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret GcpScanOptionsInputUpdateDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpScanOptionsInputUpdateData) GetAttributesOk() (*GcpScanOptionsInputUpdateDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GcpScanOptionsInputUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given GcpScanOptionsInputUpdateDataAttributes and assigns it to the Attributes field.
func (o *GcpScanOptionsInputUpdateData) SetAttributes(v GcpScanOptionsInputUpdateDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *GcpScanOptionsInputUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GcpScanOptionsInputUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GcpScanOptionsInputUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *GcpScanOptionsInputUpdateData) GetType() GcpScanOptionsInputUpdateDataType {
	if o == nil {
		var ret GcpScanOptionsInputUpdateDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GcpScanOptionsInputUpdateData) GetTypeOk() (*GcpScanOptionsInputUpdateDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GcpScanOptionsInputUpdateData) SetType(v GcpScanOptionsInputUpdateDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GcpScanOptionsInputUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GcpScanOptionsInputUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *GcpScanOptionsInputUpdateDataAttributes `json:"attributes,omitempty"`
		Id         *string                                  `json:"id"`
		Type       *GcpScanOptionsInputUpdateDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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
	o.Id = *all.Id
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
