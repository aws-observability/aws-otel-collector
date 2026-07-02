// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeyDetails Additional details for a Cloud Cost Management tag key, including its description and example tag values.
type CostTagKeyDetails struct {
	// Description of the tag key.
	Description string `json:"description"`
	// Example tag values observed for this tag key.
	TagValues []string `json:"tag_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagKeyDetails instantiates a new CostTagKeyDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagKeyDetails(description string, tagValues []string) *CostTagKeyDetails {
	this := CostTagKeyDetails{}
	this.Description = description
	this.TagValues = tagValues
	return &this
}

// NewCostTagKeyDetailsWithDefaults instantiates a new CostTagKeyDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagKeyDetailsWithDefaults() *CostTagKeyDetails {
	this := CostTagKeyDetails{}
	return &this
}

// GetDescription returns the Description field value.
func (o *CostTagKeyDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyDetails) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CostTagKeyDetails) SetDescription(v string) {
	o.Description = v
}

// GetTagValues returns the TagValues field value.
func (o *CostTagKeyDetails) GetTagValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyDetails) GetTagValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValues, true
}

// SetTagValues sets field value.
func (o *CostTagKeyDetails) SetTagValues(v []string) {
	o.TagValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagKeyDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["tag_values"] = o.TagValues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagKeyDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string   `json:"description"`
		TagValues   *[]string `json:"tag_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.TagValues == nil {
		return fmt.Errorf("required field tag_values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "tag_values"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.TagValues = *all.TagValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
