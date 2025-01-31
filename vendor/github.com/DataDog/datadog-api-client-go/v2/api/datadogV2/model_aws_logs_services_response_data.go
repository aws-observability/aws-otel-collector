// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogsServicesResponseData AWS Logs Services response body
type AWSLogsServicesResponseData struct {
	// AWS Logs Services response body
	Attributes *AWSLogsServicesResponseAttributes `json:"attributes,omitempty"`
	// The `AWSLogsServicesResponseData` `id`.
	Id string `json:"id"`
	// The `AWSLogsServicesResponseData` `type`.
	Type AWSLogsServicesResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLogsServicesResponseData instantiates a new AWSLogsServicesResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLogsServicesResponseData(id string, typeVar AWSLogsServicesResponseDataType) *AWSLogsServicesResponseData {
	this := AWSLogsServicesResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewAWSLogsServicesResponseDataWithDefaults instantiates a new AWSLogsServicesResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLogsServicesResponseDataWithDefaults() *AWSLogsServicesResponseData {
	this := AWSLogsServicesResponseData{}
	var id string = "logs_services"
	this.Id = id
	var typeVar AWSLogsServicesResponseDataType = AWSLOGSSERVICESRESPONSEDATATYPE_LOGS_SERVICES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AWSLogsServicesResponseData) GetAttributes() AWSLogsServicesResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret AWSLogsServicesResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogsServicesResponseData) GetAttributesOk() (*AWSLogsServicesResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AWSLogsServicesResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AWSLogsServicesResponseAttributes and assigns it to the Attributes field.
func (o *AWSLogsServicesResponseData) SetAttributes(v AWSLogsServicesResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *AWSLogsServicesResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AWSLogsServicesResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AWSLogsServicesResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AWSLogsServicesResponseData) GetType() AWSLogsServicesResponseDataType {
	if o == nil {
		var ret AWSLogsServicesResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSLogsServicesResponseData) GetTypeOk() (*AWSLogsServicesResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSLogsServicesResponseData) SetType(v AWSLogsServicesResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLogsServicesResponseData) MarshalJSON() ([]byte, error) {
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
func (o *AWSLogsServicesResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AWSLogsServicesResponseAttributes `json:"attributes,omitempty"`
		Id         *string                            `json:"id"`
		Type       *AWSLogsServicesResponseDataType   `json:"type"`
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
