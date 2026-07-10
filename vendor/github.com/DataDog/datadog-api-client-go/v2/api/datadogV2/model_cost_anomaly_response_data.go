// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAnomalyResponseData Resource wrapper for a single cost anomaly.
type CostAnomalyResponseData struct {
	// A single detected Cloud Cost Management anomaly.
	Attributes CostAnomaly `json:"attributes"`
	// The unique identifier of the anomaly.
	Id string `json:"id"`
	// Type of the cost anomalies collection resource. Must be `anomalies`.
	Type CostAnomaliesResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostAnomalyResponseData instantiates a new CostAnomalyResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostAnomalyResponseData(attributes CostAnomaly, id string, typeVar CostAnomaliesResponseDataType) *CostAnomalyResponseData {
	this := CostAnomalyResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCostAnomalyResponseDataWithDefaults instantiates a new CostAnomalyResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostAnomalyResponseDataWithDefaults() *CostAnomalyResponseData {
	this := CostAnomalyResponseData{}
	var typeVar CostAnomaliesResponseDataType = COSTANOMALIESRESPONSEDATATYPE_ANOMALIES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CostAnomalyResponseData) GetAttributes() CostAnomaly {
	if o == nil {
		var ret CostAnomaly
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyResponseData) GetAttributesOk() (*CostAnomaly, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CostAnomalyResponseData) SetAttributes(v CostAnomaly) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CostAnomalyResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CostAnomalyResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CostAnomalyResponseData) GetType() CostAnomaliesResponseDataType {
	if o == nil {
		var ret CostAnomaliesResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyResponseData) GetTypeOk() (*CostAnomaliesResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CostAnomalyResponseData) SetType(v CostAnomaliesResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostAnomalyResponseData) MarshalJSON() ([]byte, error) {
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
func (o *CostAnomalyResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CostAnomaly                   `json:"attributes"`
		Id         *string                        `json:"id"`
		Type       *CostAnomaliesResponseDataType `json:"type"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
