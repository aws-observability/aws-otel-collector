// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleBulkDeleteResponse Response for bulk deleting security monitoring rules.
type SecurityMonitoringRuleBulkDeleteResponse struct {
	// List of successfully deleted rule IDs.
	DeletedRules []string `json:"deletedRules,omitempty"`
	// List of rule IDs that could not be deleted.
	FailedRules []string `json:"failedRules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleBulkDeleteResponse instantiates a new SecurityMonitoringRuleBulkDeleteResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleBulkDeleteResponse() *SecurityMonitoringRuleBulkDeleteResponse {
	this := SecurityMonitoringRuleBulkDeleteResponse{}
	return &this
}

// NewSecurityMonitoringRuleBulkDeleteResponseWithDefaults instantiates a new SecurityMonitoringRuleBulkDeleteResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleBulkDeleteResponseWithDefaults() *SecurityMonitoringRuleBulkDeleteResponse {
	this := SecurityMonitoringRuleBulkDeleteResponse{}
	return &this
}

// GetDeletedRules returns the DeletedRules field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleBulkDeleteResponse) GetDeletedRules() []string {
	if o == nil || o.DeletedRules == nil {
		var ret []string
		return ret
	}
	return o.DeletedRules
}

// GetDeletedRulesOk returns a tuple with the DeletedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponse) GetDeletedRulesOk() (*[]string, bool) {
	if o == nil || o.DeletedRules == nil {
		return nil, false
	}
	return &o.DeletedRules, true
}

// HasDeletedRules returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponse) HasDeletedRules() bool {
	return o != nil && o.DeletedRules != nil
}

// SetDeletedRules gets a reference to the given []string and assigns it to the DeletedRules field.
func (o *SecurityMonitoringRuleBulkDeleteResponse) SetDeletedRules(v []string) {
	o.DeletedRules = v
}

// GetFailedRules returns the FailedRules field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleBulkDeleteResponse) GetFailedRules() []string {
	if o == nil || o.FailedRules == nil {
		var ret []string
		return ret
	}
	return o.FailedRules
}

// GetFailedRulesOk returns a tuple with the FailedRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponse) GetFailedRulesOk() (*[]string, bool) {
	if o == nil || o.FailedRules == nil {
		return nil, false
	}
	return &o.FailedRules, true
}

// HasFailedRules returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleBulkDeleteResponse) HasFailedRules() bool {
	return o != nil && o.FailedRules != nil
}

// SetFailedRules gets a reference to the given []string and assigns it to the FailedRules field.
func (o *SecurityMonitoringRuleBulkDeleteResponse) SetFailedRules(v []string) {
	o.FailedRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleBulkDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedRules != nil {
		toSerialize["deletedRules"] = o.DeletedRules
	}
	if o.FailedRules != nil {
		toSerialize["failedRules"] = o.FailedRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleBulkDeleteResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedRules []string `json:"deletedRules,omitempty"`
		FailedRules  []string `json:"failedRules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deletedRules", "failedRules"})
	} else {
		return err
	}
	o.DeletedRules = all.DeletedRules
	o.FailedRules = all.FailedRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
