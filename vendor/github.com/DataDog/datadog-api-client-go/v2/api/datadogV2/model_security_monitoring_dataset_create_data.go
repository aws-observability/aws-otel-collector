// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetCreateData The data wrapper of a dataset create request.
type SecurityMonitoringDatasetCreateData struct {
	// The attributes of a dataset create or update request.
	Attributes SecurityMonitoringDatasetAttributesRequest `json:"attributes"`
	// The type of resource for a dataset create request.
	Type SecurityMonitoringDatasetCreateType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetCreateData instantiates a new SecurityMonitoringDatasetCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetCreateData(attributes SecurityMonitoringDatasetAttributesRequest, typeVar SecurityMonitoringDatasetCreateType) *SecurityMonitoringDatasetCreateData {
	this := SecurityMonitoringDatasetCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringDatasetCreateDataWithDefaults instantiates a new SecurityMonitoringDatasetCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetCreateDataWithDefaults() *SecurityMonitoringDatasetCreateData {
	this := SecurityMonitoringDatasetCreateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SecurityMonitoringDatasetCreateData) GetAttributes() SecurityMonitoringDatasetAttributesRequest {
	if o == nil {
		var ret SecurityMonitoringDatasetAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetCreateData) GetAttributesOk() (*SecurityMonitoringDatasetAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SecurityMonitoringDatasetCreateData) SetAttributes(v SecurityMonitoringDatasetAttributesRequest) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringDatasetCreateData) GetType() SecurityMonitoringDatasetCreateType {
	if o == nil {
		var ret SecurityMonitoringDatasetCreateType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetCreateData) GetTypeOk() (*SecurityMonitoringDatasetCreateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringDatasetCreateData) SetType(v SecurityMonitoringDatasetCreateType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetCreateData) MarshalJSON() ([]byte, error) {
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
func (o *SecurityMonitoringDatasetCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SecurityMonitoringDatasetAttributesRequest `json:"attributes"`
		Type       *SecurityMonitoringDatasetCreateType        `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
