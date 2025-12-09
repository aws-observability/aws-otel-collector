// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPolicyVersion The versions of the policy
type CloudWorkloadSecurityAgentPolicyVersion struct {
	// The date and time the version was created
	Date datadog.NullableString `json:"Date,omitempty"`
	// The version of the policy
	Name *string `json:"Name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentPolicyVersion instantiates a new CloudWorkloadSecurityAgentPolicyVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentPolicyVersion() *CloudWorkloadSecurityAgentPolicyVersion {
	this := CloudWorkloadSecurityAgentPolicyVersion{}
	return &this
}

// NewCloudWorkloadSecurityAgentPolicyVersionWithDefaults instantiates a new CloudWorkloadSecurityAgentPolicyVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentPolicyVersionWithDefaults() *CloudWorkloadSecurityAgentPolicyVersion {
	this := CloudWorkloadSecurityAgentPolicyVersion{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudWorkloadSecurityAgentPolicyVersion) GetDate() string {
	if o == nil || o.Date.Get() == nil {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CloudWorkloadSecurityAgentPolicyVersion) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyVersion) HasDate() bool {
	return o != nil && o.Date.IsSet()
}

// SetDate gets a reference to the given datadog.NullableString and assigns it to the Date field.
func (o *CloudWorkloadSecurityAgentPolicyVersion) SetDate(v string) {
	o.Date.Set(&v)
}

// SetDateNil sets the value for Date to be an explicit nil.
func (o *CloudWorkloadSecurityAgentPolicyVersion) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil.
func (o *CloudWorkloadSecurityAgentPolicyVersion) UnsetDate() {
	o.Date.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyVersion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyVersion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyVersion) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudWorkloadSecurityAgentPolicyVersion) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentPolicyVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Date.IsSet() {
		toSerialize["Date"] = o.Date.Get()
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentPolicyVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Date datadog.NullableString `json:"Date,omitempty"`
		Name *string                `json:"Name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"Date", "Name"})
	} else {
		return err
	}
	o.Date = all.Date
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
