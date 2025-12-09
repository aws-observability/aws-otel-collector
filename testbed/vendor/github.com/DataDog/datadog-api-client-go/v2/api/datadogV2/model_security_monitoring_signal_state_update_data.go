// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalStateUpdateData Data containing the patch for changing the state of a signal.
type SecurityMonitoringSignalStateUpdateData struct {
	// Attributes describing the change of state of a security signal.
	Attributes SecurityMonitoringSignalStateUpdateAttributes `json:"attributes"`
	// The unique ID of the security signal.
	Id interface{} `json:"id,omitempty"`
	// The type of event.
	Type *SecurityMonitoringSignalMetadataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalStateUpdateData instantiates a new SecurityMonitoringSignalStateUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalStateUpdateData(attributes SecurityMonitoringSignalStateUpdateAttributes) *SecurityMonitoringSignalStateUpdateData {
	this := SecurityMonitoringSignalStateUpdateData{}
	this.Attributes = attributes
	var typeVar SecurityMonitoringSignalMetadataType = SECURITYMONITORINGSIGNALMETADATATYPE_SIGNAL_METADATA
	this.Type = &typeVar
	return &this
}

// NewSecurityMonitoringSignalStateUpdateDataWithDefaults instantiates a new SecurityMonitoringSignalStateUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalStateUpdateDataWithDefaults() *SecurityMonitoringSignalStateUpdateData {
	this := SecurityMonitoringSignalStateUpdateData{}
	var typeVar SecurityMonitoringSignalMetadataType = SECURITYMONITORINGSIGNALMETADATATYPE_SIGNAL_METADATA
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringSignalStateUpdateData) GetAttributes() SecurityMonitoringSignalStateUpdateAttributes {
	if o == nil {
		var ret SecurityMonitoringSignalStateUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalStateUpdateData) GetAttributesOk() (*SecurityMonitoringSignalStateUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringSignalStateUpdateData) SetAttributes(v SecurityMonitoringSignalStateUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalStateUpdateData) GetId() interface{} {
	if o == nil || o.Id == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalStateUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalStateUpdateData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *SecurityMonitoringSignalStateUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalStateUpdateData) GetType() SecurityMonitoringSignalMetadataType {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringSignalMetadataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalStateUpdateData) GetTypeOk() (*SecurityMonitoringSignalMetadataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalStateUpdateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SecurityMonitoringSignalMetadataType and assigns it to the Type field.
func (o *SecurityMonitoringSignalStateUpdateData) SetType(v SecurityMonitoringSignalMetadataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalStateUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *SecurityMonitoringSignalStateUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringSignalStateUpdateAttributes `json:"attributes"`
		Id         interface{}                                    `json:"id,omitempty"`
		Type       *SecurityMonitoringSignalMetadataType          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
