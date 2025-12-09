// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationData JSON:API resource object for SPA Recommendation. Includes type, optional ID, and resource attributes with structured recommendations.
type RecommendationData struct {
	// Attributes of the SPA Recommendation resource. Contains recommendations for both driver and executor components.
	Attributes RecommendationAttributes `json:"attributes"`
	// Resource identifier for the recommendation. Optional in responses.
	Id *string `json:"id,omitempty"`
	// JSON:API resource type for Spark Pod Autosizing recommendations. Identifies the Recommendation resource returned by SPA.
	Type RecommendationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationData instantiates a new RecommendationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationData(attributes RecommendationAttributes, typeVar RecommendationType) *RecommendationData {
	this := RecommendationData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewRecommendationDataWithDefaults instantiates a new RecommendationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationDataWithDefaults() *RecommendationData {
	this := RecommendationData{}
	var typeVar RecommendationType = RECOMMENDATIONTYPE_RECOMMENDATION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RecommendationData) GetAttributes() RecommendationAttributes {
	if o == nil {
		var ret RecommendationAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RecommendationData) GetAttributesOk() (*RecommendationAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RecommendationData) SetAttributes(v RecommendationAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecommendationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecommendationData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RecommendationData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *RecommendationData) GetType() RecommendationType {
	if o == nil {
		var ret RecommendationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecommendationData) GetTypeOk() (*RecommendationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RecommendationData) SetType(v RecommendationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationData) MarshalJSON() ([]byte, error) {
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
func (o *RecommendationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RecommendationAttributes `json:"attributes"`
		Id         *string                   `json:"id,omitempty"`
		Type       *RecommendationType       `json:"type"`
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
