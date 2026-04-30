// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssignSeatsUserRequestDataAttributes Attributes specifying the product and users to whom seats will be assigned.
type AssignSeatsUserRequestDataAttributes struct {
	// The product code for which to assign seats.
	ProductCode string `json:"product_code"`
	// The list of user IDs to assign seats to.
	UserUuids []string `json:"user_uuids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssignSeatsUserRequestDataAttributes instantiates a new AssignSeatsUserRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssignSeatsUserRequestDataAttributes(productCode string, userUuids []string) *AssignSeatsUserRequestDataAttributes {
	this := AssignSeatsUserRequestDataAttributes{}
	this.ProductCode = productCode
	this.UserUuids = userUuids
	return &this
}

// NewAssignSeatsUserRequestDataAttributesWithDefaults instantiates a new AssignSeatsUserRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssignSeatsUserRequestDataAttributesWithDefaults() *AssignSeatsUserRequestDataAttributes {
	this := AssignSeatsUserRequestDataAttributes{}
	return &this
}

// GetProductCode returns the ProductCode field value.
func (o *AssignSeatsUserRequestDataAttributes) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserRequestDataAttributes) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value.
func (o *AssignSeatsUserRequestDataAttributes) SetProductCode(v string) {
	o.ProductCode = v
}

// GetUserUuids returns the UserUuids field value.
func (o *AssignSeatsUserRequestDataAttributes) GetUserUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserUuids
}

// GetUserUuidsOk returns a tuple with the UserUuids field value
// and a boolean to check if the value has been set.
func (o *AssignSeatsUserRequestDataAttributes) GetUserUuidsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuids, true
}

// SetUserUuids sets field value.
func (o *AssignSeatsUserRequestDataAttributes) SetUserUuids(v []string) {
	o.UserUuids = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssignSeatsUserRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["product_code"] = o.ProductCode
	toSerialize["user_uuids"] = o.UserUuids

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssignSeatsUserRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProductCode *string   `json:"product_code"`
		UserUuids   *[]string `json:"user_uuids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProductCode == nil {
		return fmt.Errorf("required field product_code missing")
	}
	if all.UserUuids == nil {
		return fmt.Errorf("required field user_uuids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"product_code", "user_uuids"})
	} else {
		return err
	}
	o.ProductCode = *all.ProductCode
	o.UserUuids = *all.UserUuids

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
