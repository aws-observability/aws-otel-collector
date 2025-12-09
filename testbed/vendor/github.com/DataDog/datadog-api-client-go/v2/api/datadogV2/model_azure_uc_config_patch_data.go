// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPatchData Azure config Patch data.
type AzureUCConfigPatchData struct {
	// Attributes for Azure config Patch Request.
	Attributes *AzureUCConfigPatchRequestAttributes `json:"attributes,omitempty"`
	// Type of Azure config Patch Request.
	Type AzureUCConfigPatchRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureUCConfigPatchData instantiates a new AzureUCConfigPatchData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureUCConfigPatchData(typeVar AzureUCConfigPatchRequestType) *AzureUCConfigPatchData {
	this := AzureUCConfigPatchData{}
	this.Type = typeVar
	return &this
}

// NewAzureUCConfigPatchDataWithDefaults instantiates a new AzureUCConfigPatchData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureUCConfigPatchDataWithDefaults() *AzureUCConfigPatchData {
	this := AzureUCConfigPatchData{}
	var typeVar AzureUCConfigPatchRequestType = AZUREUCCONFIGPATCHREQUESTTYPE_AZURE_UC_CONFIG_PATCH_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AzureUCConfigPatchData) GetAttributes() AzureUCConfigPatchRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret AzureUCConfigPatchRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPatchData) GetAttributesOk() (*AzureUCConfigPatchRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AzureUCConfigPatchData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AzureUCConfigPatchRequestAttributes and assigns it to the Attributes field.
func (o *AzureUCConfigPatchData) SetAttributes(v AzureUCConfigPatchRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *AzureUCConfigPatchData) GetType() AzureUCConfigPatchRequestType {
	if o == nil {
		var ret AzureUCConfigPatchRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPatchData) GetTypeOk() (*AzureUCConfigPatchRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AzureUCConfigPatchData) SetType(v AzureUCConfigPatchRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureUCConfigPatchData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureUCConfigPatchData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AzureUCConfigPatchRequestAttributes `json:"attributes,omitempty"`
		Type       *AzureUCConfigPatchRequestType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
