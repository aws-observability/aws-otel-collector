// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringCriticalAssetUpdateData The new critical asset properties; partial updates are supported.
type SecurityMonitoringCriticalAssetUpdateData struct {
	// The critical asset properties to be updated.
	Attributes SecurityMonitoringCriticalAssetUpdateAttributes `json:"attributes"`
	// The type of the resource. The value should always be `critical_assets`.
	Type SecurityMonitoringCriticalAssetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringCriticalAssetUpdateData instantiates a new SecurityMonitoringCriticalAssetUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringCriticalAssetUpdateData(attributes SecurityMonitoringCriticalAssetUpdateAttributes, typeVar SecurityMonitoringCriticalAssetType) *SecurityMonitoringCriticalAssetUpdateData {
	this := SecurityMonitoringCriticalAssetUpdateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringCriticalAssetUpdateDataWithDefaults instantiates a new SecurityMonitoringCriticalAssetUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringCriticalAssetUpdateDataWithDefaults() *SecurityMonitoringCriticalAssetUpdateData {
	this := SecurityMonitoringCriticalAssetUpdateData{}
	var typeVar SecurityMonitoringCriticalAssetType = SECURITYMONITORINGCRITICALASSETTYPE_CRITICAL_ASSETS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringCriticalAssetUpdateData) GetAttributes() SecurityMonitoringCriticalAssetUpdateAttributes {
	if o == nil {
		var ret SecurityMonitoringCriticalAssetUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateData) GetAttributesOk() (*SecurityMonitoringCriticalAssetUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringCriticalAssetUpdateData) SetAttributes(v SecurityMonitoringCriticalAssetUpdateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringCriticalAssetUpdateData) GetType() SecurityMonitoringCriticalAssetType {
	if o == nil {
		var ret SecurityMonitoringCriticalAssetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringCriticalAssetUpdateData) GetTypeOk() (*SecurityMonitoringCriticalAssetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringCriticalAssetUpdateData) SetType(v SecurityMonitoringCriticalAssetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringCriticalAssetUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringCriticalAssetUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringCriticalAssetUpdateAttributes `json:"attributes"`
		Type       *SecurityMonitoringCriticalAssetType             `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
