// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalsBulkUpdateData Data for updating a single security signal in a bulk update operation.
type SecurityMonitoringSignalsBulkUpdateData struct {
	// Attributes for updating the triage state or assignee of a security signal.
	Attributes SecurityMonitoringSignalUpdateAttributes `json:"attributes"`
	// The unique ID of the security signal.
	Id string `json:"id"`
	// The type of event.
	Type *SecurityMonitoringSignalType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalsBulkUpdateData instantiates a new SecurityMonitoringSignalsBulkUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalsBulkUpdateData(attributes SecurityMonitoringSignalUpdateAttributes, id string) *SecurityMonitoringSignalsBulkUpdateData {
	this := SecurityMonitoringSignalsBulkUpdateData{}
	this.Attributes = attributes
	this.Id = id
	var typeVar SecurityMonitoringSignalType = SECURITYMONITORINGSIGNALTYPE_SIGNAL
	this.Type = &typeVar
	return &this
}

// NewSecurityMonitoringSignalsBulkUpdateDataWithDefaults instantiates a new SecurityMonitoringSignalsBulkUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalsBulkUpdateDataWithDefaults() *SecurityMonitoringSignalsBulkUpdateData {
	this := SecurityMonitoringSignalsBulkUpdateData{}
	var typeVar SecurityMonitoringSignalType = SECURITYMONITORINGSIGNALTYPE_SIGNAL
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringSignalsBulkUpdateData) GetAttributes() SecurityMonitoringSignalUpdateAttributes {
	if o == nil {
		var ret SecurityMonitoringSignalUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkUpdateData) GetAttributesOk() (*SecurityMonitoringSignalUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringSignalsBulkUpdateData) SetAttributes(v SecurityMonitoringSignalUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *SecurityMonitoringSignalsBulkUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityMonitoringSignalsBulkUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalsBulkUpdateData) GetType() SecurityMonitoringSignalType {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringSignalType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkUpdateData) GetTypeOk() (*SecurityMonitoringSignalType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalsBulkUpdateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SecurityMonitoringSignalType and assigns it to the Type field.
func (o *SecurityMonitoringSignalsBulkUpdateData) SetType(v SecurityMonitoringSignalType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalsBulkUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *SecurityMonitoringSignalsBulkUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringSignalUpdateAttributes `json:"attributes"`
		Id         *string                                   `json:"id"`
		Type       *SecurityMonitoringSignalType             `json:"type,omitempty"`
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
