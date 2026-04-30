// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalSuggestedAction A suggested action for a security signal.
type SecurityMonitoringSignalSuggestedAction struct {
	// Attributes of a suggested action for a security signal. The available fields depend on the action type.
	Attributes SecurityMonitoringSignalSuggestedActionAttributes `json:"attributes"`
	// The unique ID of the suggested action.
	Id string `json:"id"`
	// The type of the suggested action resource.
	Type SecurityMonitoringSignalSuggestedActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalSuggestedAction instantiates a new SecurityMonitoringSignalSuggestedAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalSuggestedAction(attributes SecurityMonitoringSignalSuggestedActionAttributes, id string, typeVar SecurityMonitoringSignalSuggestedActionType) *SecurityMonitoringSignalSuggestedAction {
	this := SecurityMonitoringSignalSuggestedAction{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringSignalSuggestedActionWithDefaults instantiates a new SecurityMonitoringSignalSuggestedAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalSuggestedActionWithDefaults() *SecurityMonitoringSignalSuggestedAction {
	this := SecurityMonitoringSignalSuggestedAction{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringSignalSuggestedAction) GetAttributes() SecurityMonitoringSignalSuggestedActionAttributes {
	if o == nil {
		var ret SecurityMonitoringSignalSuggestedActionAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedAction) GetAttributesOk() (*SecurityMonitoringSignalSuggestedActionAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringSignalSuggestedAction) SetAttributes(v SecurityMonitoringSignalSuggestedActionAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *SecurityMonitoringSignalSuggestedAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityMonitoringSignalSuggestedAction) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringSignalSuggestedAction) GetType() SecurityMonitoringSignalSuggestedActionType {
	if o == nil {
		var ret SecurityMonitoringSignalSuggestedActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedAction) GetTypeOk() (*SecurityMonitoringSignalSuggestedActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringSignalSuggestedAction) SetType(v SecurityMonitoringSignalSuggestedActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalSuggestedAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalSuggestedAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringSignalSuggestedActionAttributes `json:"attributes"`
		Id         *string                                            `json:"id"`
		Type       *SecurityMonitoringSignalSuggestedActionType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
