// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPair Azure config pair.
type AzureUCConfigPair struct {
	// Attributes for Azure config pair.
	Attributes AzureUCConfigPairAttributes `json:"attributes"`
	// The ID of Cloud Cost Management account.
	Id *int64 `json:"id,omitempty"`
	// Type of Azure config pair.
	Type AzureUCConfigPairType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureUCConfigPair instantiates a new AzureUCConfigPair object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureUCConfigPair(attributes AzureUCConfigPairAttributes, typeVar AzureUCConfigPairType) *AzureUCConfigPair {
	this := AzureUCConfigPair{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewAzureUCConfigPairWithDefaults instantiates a new AzureUCConfigPair object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureUCConfigPairWithDefaults() *AzureUCConfigPair {
	this := AzureUCConfigPair{}
	var typeVar AzureUCConfigPairType = AZUREUCCONFIGPAIRTYPE_AZURE_UC_CONFIGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AzureUCConfigPair) GetAttributes() AzureUCConfigPairAttributes {
	if o == nil {
		var ret AzureUCConfigPairAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPair) GetAttributesOk() (*AzureUCConfigPairAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AzureUCConfigPair) SetAttributes(v AzureUCConfigPairAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureUCConfigPair) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPair) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureUCConfigPair) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AzureUCConfigPair) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *AzureUCConfigPair) GetType() AzureUCConfigPairType {
	if o == nil {
		var ret AzureUCConfigPairType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPair) GetTypeOk() (*AzureUCConfigPairType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AzureUCConfigPair) SetType(v AzureUCConfigPairType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureUCConfigPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureUCConfigPair) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AzureUCConfigPairAttributes `json:"attributes"`
		Id         *int64                       `json:"id,omitempty"`
		Type       *AzureUCConfigPairType       `json:"type"`
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
