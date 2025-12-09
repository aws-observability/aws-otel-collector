// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionCreateAttributes Attributes for creating an org connection.
type OrgConnectionCreateAttributes struct {
	// List of connection types to establish.
	ConnectionTypes []OrgConnectionTypeEnum `json:"connection_types"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConnectionCreateAttributes instantiates a new OrgConnectionCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConnectionCreateAttributes(connectionTypes []OrgConnectionTypeEnum) *OrgConnectionCreateAttributes {
	this := OrgConnectionCreateAttributes{}
	this.ConnectionTypes = connectionTypes
	return &this
}

// NewOrgConnectionCreateAttributesWithDefaults instantiates a new OrgConnectionCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConnectionCreateAttributesWithDefaults() *OrgConnectionCreateAttributes {
	this := OrgConnectionCreateAttributes{}
	return &this
}

// GetConnectionTypes returns the ConnectionTypes field value.
func (o *OrgConnectionCreateAttributes) GetConnectionTypes() []OrgConnectionTypeEnum {
	if o == nil {
		var ret []OrgConnectionTypeEnum
		return ret
	}
	return o.ConnectionTypes
}

// GetConnectionTypesOk returns a tuple with the ConnectionTypes field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionCreateAttributes) GetConnectionTypesOk() (*[]OrgConnectionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTypes, true
}

// SetConnectionTypes sets field value.
func (o *OrgConnectionCreateAttributes) SetConnectionTypes(v []OrgConnectionTypeEnum) {
	o.ConnectionTypes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConnectionCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connection_types"] = o.ConnectionTypes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConnectionCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionTypes *[]OrgConnectionTypeEnum `json:"connection_types"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionTypes == nil {
		return fmt.Errorf("required field connection_types missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connection_types"})
	} else {
		return err
	}
	o.ConnectionTypes = *all.ConnectionTypes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
