// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyRuleOverride Override WAF rule parameters for services in a policy.
type ApplicationSecurityPolicyRuleOverride struct {
	// When blocking is enabled, the rule will block the traffic matched by this rule.
	Blocking bool `json:"blocking"`
	// When false, this rule will not match any traffic.
	Enabled bool `json:"enabled"`
	// When true, collects additional data from the WAF for this rule.
	ExtendedDataCollection *bool `json:"extended_data_collection,omitempty"`
	// Override the parameters for this WAF rule identifier.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityPolicyRuleOverride instantiates a new ApplicationSecurityPolicyRuleOverride object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityPolicyRuleOverride(blocking bool, enabled bool, id string) *ApplicationSecurityPolicyRuleOverride {
	this := ApplicationSecurityPolicyRuleOverride{}
	this.Blocking = blocking
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewApplicationSecurityPolicyRuleOverrideWithDefaults instantiates a new ApplicationSecurityPolicyRuleOverride object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityPolicyRuleOverrideWithDefaults() *ApplicationSecurityPolicyRuleOverride {
	this := ApplicationSecurityPolicyRuleOverride{}
	return &this
}

// GetBlocking returns the Blocking field value.
func (o *ApplicationSecurityPolicyRuleOverride) GetBlocking() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRuleOverride) GetBlockingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocking, true
}

// SetBlocking sets field value.
func (o *ApplicationSecurityPolicyRuleOverride) SetBlocking(v bool) {
	o.Blocking = v
}

// GetEnabled returns the Enabled field value.
func (o *ApplicationSecurityPolicyRuleOverride) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRuleOverride) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ApplicationSecurityPolicyRuleOverride) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExtendedDataCollection returns the ExtendedDataCollection field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyRuleOverride) GetExtendedDataCollection() bool {
	if o == nil || o.ExtendedDataCollection == nil {
		var ret bool
		return ret
	}
	return *o.ExtendedDataCollection
}

// GetExtendedDataCollectionOk returns a tuple with the ExtendedDataCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRuleOverride) GetExtendedDataCollectionOk() (*bool, bool) {
	if o == nil || o.ExtendedDataCollection == nil {
		return nil, false
	}
	return o.ExtendedDataCollection, true
}

// HasExtendedDataCollection returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyRuleOverride) HasExtendedDataCollection() bool {
	return o != nil && o.ExtendedDataCollection != nil
}

// SetExtendedDataCollection gets a reference to the given bool and assigns it to the ExtendedDataCollection field.
func (o *ApplicationSecurityPolicyRuleOverride) SetExtendedDataCollection(v bool) {
	o.ExtendedDataCollection = &v
}

// GetId returns the Id field value.
func (o *ApplicationSecurityPolicyRuleOverride) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyRuleOverride) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ApplicationSecurityPolicyRuleOverride) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityPolicyRuleOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["blocking"] = o.Blocking
	toSerialize["enabled"] = o.Enabled
	if o.ExtendedDataCollection != nil {
		toSerialize["extended_data_collection"] = o.ExtendedDataCollection
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityPolicyRuleOverride) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Blocking               *bool   `json:"blocking"`
		Enabled                *bool   `json:"enabled"`
		ExtendedDataCollection *bool   `json:"extended_data_collection,omitempty"`
		Id                     *string `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Blocking == nil {
		return fmt.Errorf("required field blocking missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"blocking", "enabled", "extended_data_collection", "id"})
	} else {
		return err
	}
	o.Blocking = *all.Blocking
	o.Enabled = *all.Enabled
	o.ExtendedDataCollection = all.ExtendedDataCollection
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
