// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleActionParameters The definition of `ApplicationSecurityWafCustomRuleActionParameters` object.
type ApplicationSecurityWafCustomRuleActionParameters struct {
	// The location to redirect to when the WAF custom rule triggers.
	Location *string `json:"location,omitempty"`
	// The status code to return when the WAF custom rule triggers.
	StatusCode *int64 `json:"status_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleActionParameters instantiates a new ApplicationSecurityWafCustomRuleActionParameters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleActionParameters() *ApplicationSecurityWafCustomRuleActionParameters {
	this := ApplicationSecurityWafCustomRuleActionParameters{}
	var statusCode int64 = 403
	this.StatusCode = &statusCode
	return &this
}

// NewApplicationSecurityWafCustomRuleActionParametersWithDefaults instantiates a new ApplicationSecurityWafCustomRuleActionParameters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleActionParametersWithDefaults() *ApplicationSecurityWafCustomRuleActionParameters {
	this := ApplicationSecurityWafCustomRuleActionParameters{}
	var statusCode int64 = 403
	this.StatusCode = &statusCode
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleActionParameters) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleActionParameters) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleActionParameters) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ApplicationSecurityWafCustomRuleActionParameters) SetLocation(v string) {
	o.Location = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleActionParameters) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleActionParameters) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleActionParameters) HasStatusCode() bool {
	return o != nil && o.StatusCode != nil
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *ApplicationSecurityWafCustomRuleActionParameters) SetStatusCode(v int64) {
	o.StatusCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleActionParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleActionParameters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Location   *string `json:"location,omitempty"`
		StatusCode *int64  `json:"status_code,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"location", "status_code"})
	} else {
		return err
	}
	o.Location = all.Location
	o.StatusCode = all.StatusCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
