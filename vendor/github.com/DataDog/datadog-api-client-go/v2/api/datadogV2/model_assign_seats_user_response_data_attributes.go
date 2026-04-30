// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssignSeatsUserResponseDataAttributes Attributes of the assign seats response, including the list of users assigned and the product code.
type AssignSeatsUserResponseDataAttributes struct {
	// The list of user IDs to which the seats were assigned.
	AssignedIds []string `json:"assigned_ids,omitempty"`
	// The product code for which the seats were assigned.
	ProductCode *string `json:"product_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssignSeatsUserResponseDataAttributes instantiates a new AssignSeatsUserResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssignSeatsUserResponseDataAttributes() *AssignSeatsUserResponseDataAttributes {
	this := AssignSeatsUserResponseDataAttributes{}
	return &this
}

// NewAssignSeatsUserResponseDataAttributesWithDefaults instantiates a new AssignSeatsUserResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssignSeatsUserResponseDataAttributesWithDefaults() *AssignSeatsUserResponseDataAttributes {
	this := AssignSeatsUserResponseDataAttributes{}
	return &this
}

// GetAssignedIds returns the AssignedIds field value if set, zero value otherwise.
func (o *AssignSeatsUserResponseDataAttributes) GetAssignedIds() []string {
	if o == nil || o.AssignedIds == nil {
		var ret []string
		return ret
	}
	return o.AssignedIds
}

// GetAssignedIdsOk returns a tuple with the AssignedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserResponseDataAttributes) GetAssignedIdsOk() (*[]string, bool) {
	if o == nil || o.AssignedIds == nil {
		return nil, false
	}
	return &o.AssignedIds, true
}

// HasAssignedIds returns a boolean if a field has been set.
func (o *AssignSeatsUserResponseDataAttributes) HasAssignedIds() bool {
	return o != nil && o.AssignedIds != nil
}

// SetAssignedIds gets a reference to the given []string and assigns it to the AssignedIds field.
func (o *AssignSeatsUserResponseDataAttributes) SetAssignedIds(v []string) {
	o.AssignedIds = v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AssignSeatsUserResponseDataAttributes) GetProductCode() string {
	if o == nil || o.ProductCode == nil {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserResponseDataAttributes) GetProductCodeOk() (*string, bool) {
	if o == nil || o.ProductCode == nil {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AssignSeatsUserResponseDataAttributes) HasProductCode() bool {
	return o != nil && o.ProductCode != nil
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AssignSeatsUserResponseDataAttributes) SetProductCode(v string) {
	o.ProductCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssignSeatsUserResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignedIds != nil {
		toSerialize["assigned_ids"] = o.AssignedIds
	}
	if o.ProductCode != nil {
		toSerialize["product_code"] = o.ProductCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssignSeatsUserResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignedIds []string `json:"assigned_ids,omitempty"`
		ProductCode *string  `json:"product_code,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigned_ids", "product_code"})
	} else {
		return err
	}
	o.AssignedIds = all.AssignedIds
	o.ProductCode = all.ProductCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
