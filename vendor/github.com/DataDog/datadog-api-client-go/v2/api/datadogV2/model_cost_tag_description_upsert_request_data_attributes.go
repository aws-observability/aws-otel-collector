// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagDescriptionUpsertRequestDataAttributes Mutable attributes set when creating or updating a Cloud Cost Management tag key description.
type CostTagDescriptionUpsertRequestDataAttributes struct {
	// Cloud provider this description applies to (for example, `aws`). Omit to set the cross-cloud default for the tag key.
	Cloud *string `json:"cloud,omitempty"`
	// The human-readable description for the tag key.
	Description string `json:"description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagDescriptionUpsertRequestDataAttributes instantiates a new CostTagDescriptionUpsertRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagDescriptionUpsertRequestDataAttributes(description string) *CostTagDescriptionUpsertRequestDataAttributes {
	this := CostTagDescriptionUpsertRequestDataAttributes{}
	this.Description = description
	return &this
}

// NewCostTagDescriptionUpsertRequestDataAttributesWithDefaults instantiates a new CostTagDescriptionUpsertRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagDescriptionUpsertRequestDataAttributesWithDefaults() *CostTagDescriptionUpsertRequestDataAttributes {
	this := CostTagDescriptionUpsertRequestDataAttributes{}
	return &this
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CostTagDescriptionUpsertRequestDataAttributes) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionUpsertRequestDataAttributes) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CostTagDescriptionUpsertRequestDataAttributes) HasCloud() bool {
	return o != nil && o.Cloud != nil
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *CostTagDescriptionUpsertRequestDataAttributes) SetCloud(v string) {
	o.Cloud = &v
}

// GetDescription returns the Description field value.
func (o *CostTagDescriptionUpsertRequestDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CostTagDescriptionUpsertRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CostTagDescriptionUpsertRequestDataAttributes) SetDescription(v string) {
	o.Description = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagDescriptionUpsertRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagDescriptionUpsertRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cloud       *string `json:"cloud,omitempty"`
		Description *string `json:"description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud", "description"})
	} else {
		return err
	}
	o.Cloud = all.Cloud
	o.Description = *all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
