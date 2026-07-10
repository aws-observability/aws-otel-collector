// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseUpdateResolvedReasonAttributes Attributes for setting the resolution reason on a security case.
type CaseUpdateResolvedReasonAttributes struct {
	// The reason the security case was resolved (for example, `FALSE_POSITIVE`, `TRUE_POSITIVE`, `BENIGN_POSITIVE`).
	SecurityResolvedReason string `json:"security_resolved_reason"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseUpdateResolvedReasonAttributes instantiates a new CaseUpdateResolvedReasonAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseUpdateResolvedReasonAttributes(securityResolvedReason string) *CaseUpdateResolvedReasonAttributes {
	this := CaseUpdateResolvedReasonAttributes{}
	this.SecurityResolvedReason = securityResolvedReason
	return &this
}

// NewCaseUpdateResolvedReasonAttributesWithDefaults instantiates a new CaseUpdateResolvedReasonAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseUpdateResolvedReasonAttributesWithDefaults() *CaseUpdateResolvedReasonAttributes {
	this := CaseUpdateResolvedReasonAttributes{}
	return &this
}

// GetSecurityResolvedReason returns the SecurityResolvedReason field value.
func (o *CaseUpdateResolvedReasonAttributes) GetSecurityResolvedReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SecurityResolvedReason
}

// GetSecurityResolvedReasonOk returns a tuple with the SecurityResolvedReason field value
// and a boolean to check if the value has been set.
func (o *CaseUpdateResolvedReasonAttributes) GetSecurityResolvedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityResolvedReason, true
}

// SetSecurityResolvedReason sets field value.
func (o *CaseUpdateResolvedReasonAttributes) SetSecurityResolvedReason(v string) {
	o.SecurityResolvedReason = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseUpdateResolvedReasonAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["security_resolved_reason"] = o.SecurityResolvedReason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseUpdateResolvedReasonAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SecurityResolvedReason *string `json:"security_resolved_reason"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SecurityResolvedReason == nil {
		return fmt.Errorf("required field security_resolved_reason missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"security_resolved_reason"})
	} else {
		return err
	}
	o.SecurityResolvedReason = *all.SecurityResolvedReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
