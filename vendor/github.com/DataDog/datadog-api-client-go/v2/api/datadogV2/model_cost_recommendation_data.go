// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationData A single cost recommendation entry in JSON:API form.
type CostRecommendationData struct {
	// Attributes describing a single cost recommendation.
	Attributes *CostRecommendationDataAttributes `json:"attributes,omitempty"`
	// Unique identifier for the recommendation.
	Id *string `json:"id,omitempty"`
	// Recommendation resource type.
	Type CostRecommendationDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationData instantiates a new CostRecommendationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationData(typeVar CostRecommendationDataType) *CostRecommendationData {
	this := CostRecommendationData{}
	this.Type = typeVar
	return &this
}

// NewCostRecommendationDataWithDefaults instantiates a new CostRecommendationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationDataWithDefaults() *CostRecommendationData {
	this := CostRecommendationData{}
	var typeVar CostRecommendationDataType = COSTRECOMMENDATIONDATATYPE_RECOMMENDATION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CostRecommendationData) GetAttributes() CostRecommendationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CostRecommendationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationData) GetAttributesOk() (*CostRecommendationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CostRecommendationData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CostRecommendationDataAttributes and assigns it to the Attributes field.
func (o *CostRecommendationData) SetAttributes(v CostRecommendationDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CostRecommendationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CostRecommendationData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CostRecommendationData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *CostRecommendationData) GetType() CostRecommendationDataType {
	if o == nil {
		var ret CostRecommendationDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationData) GetTypeOk() (*CostRecommendationDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CostRecommendationData) SetType(v CostRecommendationDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *CostRecommendationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CostRecommendationDataAttributes `json:"attributes,omitempty"`
		Id         *string                           `json:"id,omitempty"`
		Type       *CostRecommendationDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
