// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseInsightsAttributes Attributes for adding or removing insights from a case.
type CaseInsightsAttributes struct {
	// Array of insights to add to or remove from a case.
	Insights []CaseInsight `json:"insights"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseInsightsAttributes instantiates a new CaseInsightsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseInsightsAttributes(insights []CaseInsight) *CaseInsightsAttributes {
	this := CaseInsightsAttributes{}
	this.Insights = insights
	return &this
}

// NewCaseInsightsAttributesWithDefaults instantiates a new CaseInsightsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseInsightsAttributesWithDefaults() *CaseInsightsAttributes {
	this := CaseInsightsAttributes{}
	return &this
}

// GetInsights returns the Insights field value.
func (o *CaseInsightsAttributes) GetInsights() []CaseInsight {
	if o == nil {
		var ret []CaseInsight
		return ret
	}
	return o.Insights
}

// GetInsightsOk returns a tuple with the Insights field value
// and a boolean to check if the value has been set.
func (o *CaseInsightsAttributes) GetInsightsOk() (*[]CaseInsight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Insights, true
}

// SetInsights sets field value.
func (o *CaseInsightsAttributes) SetInsights(v []CaseInsight) {
	o.Insights = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseInsightsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["insights"] = o.Insights

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseInsightsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Insights *[]CaseInsight `json:"insights"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Insights == nil {
		return fmt.Errorf("required field insights missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"insights"})
	} else {
		return err
	}
	o.Insights = *all.Insights

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
