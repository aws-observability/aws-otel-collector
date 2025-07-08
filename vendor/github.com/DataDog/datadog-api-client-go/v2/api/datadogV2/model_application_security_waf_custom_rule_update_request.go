// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleUpdateRequest Request object that includes the Custom Rule to update.
type ApplicationSecurityWafCustomRuleUpdateRequest struct {
	// Object for a single WAF Custom Rule.
	Data ApplicationSecurityWafCustomRuleUpdateData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleUpdateRequest instantiates a new ApplicationSecurityWafCustomRuleUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleUpdateRequest(data ApplicationSecurityWafCustomRuleUpdateData) *ApplicationSecurityWafCustomRuleUpdateRequest {
	this := ApplicationSecurityWafCustomRuleUpdateRequest{}
	this.Data = data
	return &this
}

// NewApplicationSecurityWafCustomRuleUpdateRequestWithDefaults instantiates a new ApplicationSecurityWafCustomRuleUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleUpdateRequestWithDefaults() *ApplicationSecurityWafCustomRuleUpdateRequest {
	this := ApplicationSecurityWafCustomRuleUpdateRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *ApplicationSecurityWafCustomRuleUpdateRequest) GetData() ApplicationSecurityWafCustomRuleUpdateData {
	if o == nil {
		var ret ApplicationSecurityWafCustomRuleUpdateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleUpdateRequest) GetDataOk() (*ApplicationSecurityWafCustomRuleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ApplicationSecurityWafCustomRuleUpdateRequest) SetData(v ApplicationSecurityWafCustomRuleUpdateData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ApplicationSecurityWafCustomRuleUpdateData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
