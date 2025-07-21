// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPoliciesListResponse Response object that includes a list of Agent policies
type CloudWorkloadSecurityAgentPoliciesListResponse struct {
	// A list of Agent policy objects
	Data []CloudWorkloadSecurityAgentPolicyData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentPoliciesListResponse instantiates a new CloudWorkloadSecurityAgentPoliciesListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentPoliciesListResponse() *CloudWorkloadSecurityAgentPoliciesListResponse {
	this := CloudWorkloadSecurityAgentPoliciesListResponse{}
	return &this
}

// NewCloudWorkloadSecurityAgentPoliciesListResponseWithDefaults instantiates a new CloudWorkloadSecurityAgentPoliciesListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentPoliciesListResponseWithDefaults() *CloudWorkloadSecurityAgentPoliciesListResponse {
	this := CloudWorkloadSecurityAgentPoliciesListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPoliciesListResponse) GetData() []CloudWorkloadSecurityAgentPolicyData {
	if o == nil || o.Data == nil {
		var ret []CloudWorkloadSecurityAgentPolicyData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPoliciesListResponse) GetDataOk() (*[]CloudWorkloadSecurityAgentPolicyData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPoliciesListResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []CloudWorkloadSecurityAgentPolicyData and assigns it to the Data field.
func (o *CloudWorkloadSecurityAgentPoliciesListResponse) SetData(v []CloudWorkloadSecurityAgentPolicyData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentPoliciesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentPoliciesListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []CloudWorkloadSecurityAgentPolicyData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
