// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCreatePipelineEventRequestData Data of the pipeline event to create.
type CIAppCreatePipelineEventRequestData struct {
	// Attributes of the pipeline event to create.
	Attributes *CIAppCreatePipelineEventRequestAttributes `json:"attributes,omitempty"`
	// Type of the event.
	Type *CIAppCreatePipelineEventRequestDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppCreatePipelineEventRequestData instantiates a new CIAppCreatePipelineEventRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppCreatePipelineEventRequestData() *CIAppCreatePipelineEventRequestData {
	this := CIAppCreatePipelineEventRequestData{}
	var typeVar CIAppCreatePipelineEventRequestDataType = CIAPPCREATEPIPELINEEVENTREQUESTDATATYPE_CIPIPELINE_RESOURCE_REQUEST
	this.Type = &typeVar
	return &this
}

// NewCIAppCreatePipelineEventRequestDataWithDefaults instantiates a new CIAppCreatePipelineEventRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppCreatePipelineEventRequestDataWithDefaults() *CIAppCreatePipelineEventRequestData {
	this := CIAppCreatePipelineEventRequestData{}
	var typeVar CIAppCreatePipelineEventRequestDataType = CIAPPCREATEPIPELINEEVENTREQUESTDATATYPE_CIPIPELINE_RESOURCE_REQUEST
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CIAppCreatePipelineEventRequestData) GetAttributes() CIAppCreatePipelineEventRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret CIAppCreatePipelineEventRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequestData) GetAttributesOk() (*CIAppCreatePipelineEventRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CIAppCreatePipelineEventRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CIAppCreatePipelineEventRequestAttributes and assigns it to the Attributes field.
func (o *CIAppCreatePipelineEventRequestData) SetAttributes(v CIAppCreatePipelineEventRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CIAppCreatePipelineEventRequestData) GetType() CIAppCreatePipelineEventRequestDataType {
	if o == nil || o.Type == nil {
		var ret CIAppCreatePipelineEventRequestDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequestData) GetTypeOk() (*CIAppCreatePipelineEventRequestDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CIAppCreatePipelineEventRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CIAppCreatePipelineEventRequestDataType and assigns it to the Type field.
func (o *CIAppCreatePipelineEventRequestData) SetType(v CIAppCreatePipelineEventRequestDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppCreatePipelineEventRequestData) MarshalJSON() ([]byte, error) {
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
func (o *CIAppCreatePipelineEventRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CIAppCreatePipelineEventRequestAttributes `json:"attributes,omitempty"`
		Type       *CIAppCreatePipelineEventRequestDataType   `json:"type,omitempty"`
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
