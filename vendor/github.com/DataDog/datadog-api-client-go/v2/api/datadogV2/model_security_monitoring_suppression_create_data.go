// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionCreateData Object for a single suppression rule.
type SecurityMonitoringSuppressionCreateData struct {
	// Object containing the attributes of the suppression rule to be created.
	Attributes SecurityMonitoringSuppressionCreateAttributes `json:"attributes"`
	// The type of the resource. The value should always be `suppressions`.
	Type SecurityMonitoringSuppressionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSuppressionCreateData instantiates a new SecurityMonitoringSuppressionCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSuppressionCreateData(attributes SecurityMonitoringSuppressionCreateAttributes, typeVar SecurityMonitoringSuppressionType) *SecurityMonitoringSuppressionCreateData {
	this := SecurityMonitoringSuppressionCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringSuppressionCreateDataWithDefaults instantiates a new SecurityMonitoringSuppressionCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSuppressionCreateDataWithDefaults() *SecurityMonitoringSuppressionCreateData {
	this := SecurityMonitoringSuppressionCreateData{}
	var typeVar SecurityMonitoringSuppressionType = SECURITYMONITORINGSUPPRESSIONTYPE_SUPPRESSIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringSuppressionCreateData) GetAttributes() SecurityMonitoringSuppressionCreateAttributes {
	if o == nil {
		var ret SecurityMonitoringSuppressionCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateData) GetAttributesOk() (*SecurityMonitoringSuppressionCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringSuppressionCreateData) SetAttributes(v SecurityMonitoringSuppressionCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringSuppressionCreateData) GetType() SecurityMonitoringSuppressionType {
	if o == nil {
		var ret SecurityMonitoringSuppressionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSuppressionCreateData) GetTypeOk() (*SecurityMonitoringSuppressionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringSuppressionCreateData) SetType(v SecurityMonitoringSuppressionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSuppressionCreateData) MarshalJSON() ([]byte, error) {
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
func (o *SecurityMonitoringSuppressionCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringSuppressionCreateAttributes `json:"attributes"`
		Type       *SecurityMonitoringSuppressionType             `json:"type"`
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
