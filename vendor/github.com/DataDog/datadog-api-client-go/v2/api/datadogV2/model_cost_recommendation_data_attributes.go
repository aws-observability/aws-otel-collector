// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationDataAttributes Attributes describing a single cost recommendation.
type CostRecommendationDataAttributes struct {
	// Datadog resource key identifying the recommended resource.
	DdResourceKey *string `json:"dd_resource_key,omitempty"`
	// Estimated daily savings if the recommendation is applied.
	PotentialDailySavings *CostRecommendationDataAttributesPotentialDailySavings `json:"potential_daily_savings,omitempty"`
	// The kind of recommendation (for example, `terminate` or `rightsize`).
	RecommendationType *string `json:"recommendation_type,omitempty"`
	// Cloud provider identifier of the resource.
	ResourceId *string `json:"resource_id,omitempty"`
	// Resource type (for example, `aws_ec2_instance`).
	ResourceType *string `json:"resource_type,omitempty"`
	// Tags attached to the recommended resource.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationDataAttributes instantiates a new CostRecommendationDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationDataAttributes() *CostRecommendationDataAttributes {
	this := CostRecommendationDataAttributes{}
	return &this
}

// NewCostRecommendationDataAttributesWithDefaults instantiates a new CostRecommendationDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationDataAttributesWithDefaults() *CostRecommendationDataAttributes {
	this := CostRecommendationDataAttributes{}
	return &this
}

// GetDdResourceKey returns the DdResourceKey field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributes) GetDdResourceKey() string {
	if o == nil || o.DdResourceKey == nil {
		var ret string
		return ret
	}
	return *o.DdResourceKey
}

// GetDdResourceKeyOk returns a tuple with the DdResourceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributes) GetDdResourceKeyOk() (*string, bool) {
	if o == nil || o.DdResourceKey == nil {
		return nil, false
	}
	return o.DdResourceKey, true
}

// HasDdResourceKey returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributes) HasDdResourceKey() bool {
	return o != nil && o.DdResourceKey != nil
}

// SetDdResourceKey gets a reference to the given string and assigns it to the DdResourceKey field.
func (o *CostRecommendationDataAttributes) SetDdResourceKey(v string) {
	o.DdResourceKey = &v
}

// GetPotentialDailySavings returns the PotentialDailySavings field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributes) GetPotentialDailySavings() CostRecommendationDataAttributesPotentialDailySavings {
	if o == nil || o.PotentialDailySavings == nil {
		var ret CostRecommendationDataAttributesPotentialDailySavings
		return ret
	}
	return *o.PotentialDailySavings
}

// GetPotentialDailySavingsOk returns a tuple with the PotentialDailySavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributes) GetPotentialDailySavingsOk() (*CostRecommendationDataAttributesPotentialDailySavings, bool) {
	if o == nil || o.PotentialDailySavings == nil {
		return nil, false
	}
	return o.PotentialDailySavings, true
}

// HasPotentialDailySavings returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributes) HasPotentialDailySavings() bool {
	return o != nil && o.PotentialDailySavings != nil
}

// SetPotentialDailySavings gets a reference to the given CostRecommendationDataAttributesPotentialDailySavings and assigns it to the PotentialDailySavings field.
func (o *CostRecommendationDataAttributes) SetPotentialDailySavings(v CostRecommendationDataAttributesPotentialDailySavings) {
	o.PotentialDailySavings = &v
}

// GetRecommendationType returns the RecommendationType field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributes) GetRecommendationType() string {
	if o == nil || o.RecommendationType == nil {
		var ret string
		return ret
	}
	return *o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributes) GetRecommendationTypeOk() (*string, bool) {
	if o == nil || o.RecommendationType == nil {
		return nil, false
	}
	return o.RecommendationType, true
}

// HasRecommendationType returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributes) HasRecommendationType() bool {
	return o != nil && o.RecommendationType != nil
}

// SetRecommendationType gets a reference to the given string and assigns it to the RecommendationType field.
func (o *CostRecommendationDataAttributes) SetRecommendationType(v string) {
	o.RecommendationType = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributes) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributes) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CostRecommendationDataAttributes) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributes) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributes) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *CostRecommendationDataAttributes) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CostRecommendationDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CostRecommendationDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CostRecommendationDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DdResourceKey != nil {
		toSerialize["dd_resource_key"] = o.DdResourceKey
	}
	if o.PotentialDailySavings != nil {
		toSerialize["potential_daily_savings"] = o.PotentialDailySavings
	}
	if o.RecommendationType != nil {
		toSerialize["recommendation_type"] = o.RecommendationType
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DdResourceKey         *string                                                `json:"dd_resource_key,omitempty"`
		PotentialDailySavings *CostRecommendationDataAttributesPotentialDailySavings `json:"potential_daily_savings,omitempty"`
		RecommendationType    *string                                                `json:"recommendation_type,omitempty"`
		ResourceId            *string                                                `json:"resource_id,omitempty"`
		ResourceType          *string                                                `json:"resource_type,omitempty"`
		Tags                  []string                                               `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dd_resource_key", "potential_daily_savings", "recommendation_type", "resource_id", "resource_type", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DdResourceKey = all.DdResourceKey
	if all.PotentialDailySavings != nil && all.PotentialDailySavings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PotentialDailySavings = all.PotentialDailySavings
	o.RecommendationType = all.RecommendationType
	o.ResourceId = all.ResourceId
	o.ResourceType = all.ResourceType
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
