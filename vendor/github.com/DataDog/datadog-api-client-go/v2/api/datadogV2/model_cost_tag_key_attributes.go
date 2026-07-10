// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeyAttributes Attributes of a Cloud Cost Management tag key.
type CostTagKeyAttributes struct {
	// Additional details for a Cloud Cost Management tag key, including its description and example tag values.
	Details *CostTagKeyDetails `json:"details,omitempty"`
	// List of sources that define this tag key.
	Sources []string `json:"sources"`
	// The tag key name.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostTagKeyAttributes instantiates a new CostTagKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostTagKeyAttributes(sources []string, value string) *CostTagKeyAttributes {
	this := CostTagKeyAttributes{}
	this.Sources = sources
	this.Value = value
	return &this
}

// NewCostTagKeyAttributesWithDefaults instantiates a new CostTagKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostTagKeyAttributesWithDefaults() *CostTagKeyAttributes {
	this := CostTagKeyAttributes{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CostTagKeyAttributes) GetDetails() CostTagKeyDetails {
	if o == nil || o.Details == nil {
		var ret CostTagKeyDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTagKeyAttributes) GetDetailsOk() (*CostTagKeyDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CostTagKeyAttributes) HasDetails() bool {
	return o != nil && o.Details != nil
}

// SetDetails gets a reference to the given CostTagKeyDetails and assigns it to the Details field.
func (o *CostTagKeyAttributes) SetDetails(v CostTagKeyDetails) {
	o.Details = &v
}

// GetSources returns the Sources field value.
func (o *CostTagKeyAttributes) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyAttributes) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *CostTagKeyAttributes) SetSources(v []string) {
	o.Sources = v
}

// GetValue returns the Value field value.
func (o *CostTagKeyAttributes) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CostTagKeyAttributes) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *CostTagKeyAttributes) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostTagKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	toSerialize["sources"] = o.Sources
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostTagKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Details *CostTagKeyDetails `json:"details,omitempty"`
		Sources *[]string          `json:"sources"`
		Value   *string            `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"details", "sources", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Details != nil && all.Details.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Details = all.Details
	o.Sources = *all.Sources
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
