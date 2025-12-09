// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleQueryPayloadData Payload used to test the rule query.
type SecurityMonitoringRuleQueryPayloadData struct {
	// Source of the payload.
	Ddsource *string `json:"ddsource,omitempty"`
	// Tags associated with your data.
	Ddtags *string `json:"ddtags,omitempty"`
	// The name of the originating host of the log.
	Hostname *string `json:"hostname,omitempty"`
	// The message of the payload.
	Message *string `json:"message,omitempty"`
	// The name of the application or service generating the data.
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleQueryPayloadData instantiates a new SecurityMonitoringRuleQueryPayloadData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleQueryPayloadData() *SecurityMonitoringRuleQueryPayloadData {
	this := SecurityMonitoringRuleQueryPayloadData{}
	return &this
}

// NewSecurityMonitoringRuleQueryPayloadDataWithDefaults instantiates a new SecurityMonitoringRuleQueryPayloadData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleQueryPayloadDataWithDefaults() *SecurityMonitoringRuleQueryPayloadData {
	this := SecurityMonitoringRuleQueryPayloadData{}
	return &this
}

// GetDdsource returns the Ddsource field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayloadData) GetDdsource() string {
	if o == nil || o.Ddsource == nil {
		var ret string
		return ret
	}
	return *o.Ddsource
}

// GetDdsourceOk returns a tuple with the Ddsource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) GetDdsourceOk() (*string, bool) {
	if o == nil || o.Ddsource == nil {
		return nil, false
	}
	return o.Ddsource, true
}

// HasDdsource returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) HasDdsource() bool {
	return o != nil && o.Ddsource != nil
}

// SetDdsource gets a reference to the given string and assigns it to the Ddsource field.
func (o *SecurityMonitoringRuleQueryPayloadData) SetDdsource(v string) {
	o.Ddsource = &v
}

// GetDdtags returns the Ddtags field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayloadData) GetDdtags() string {
	if o == nil || o.Ddtags == nil {
		var ret string
		return ret
	}
	return *o.Ddtags
}

// GetDdtagsOk returns a tuple with the Ddtags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) GetDdtagsOk() (*string, bool) {
	if o == nil || o.Ddtags == nil {
		return nil, false
	}
	return o.Ddtags, true
}

// HasDdtags returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) HasDdtags() bool {
	return o != nil && o.Ddtags != nil
}

// SetDdtags gets a reference to the given string and assigns it to the Ddtags field.
func (o *SecurityMonitoringRuleQueryPayloadData) SetDdtags(v string) {
	o.Ddtags = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayloadData) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SecurityMonitoringRuleQueryPayloadData) SetHostname(v string) {
	o.Hostname = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayloadData) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityMonitoringRuleQueryPayloadData) SetMessage(v string) {
	o.Message = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleQueryPayloadData) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleQueryPayloadData) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SecurityMonitoringRuleQueryPayloadData) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleQueryPayloadData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ddsource != nil {
		toSerialize["ddsource"] = o.Ddsource
	}
	if o.Ddtags != nil {
		toSerialize["ddtags"] = o.Ddtags
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleQueryPayloadData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ddsource *string `json:"ddsource,omitempty"`
		Ddtags   *string `json:"ddtags,omitempty"`
		Hostname *string `json:"hostname,omitempty"`
		Message  *string `json:"message,omitempty"`
		Service  *string `json:"service,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ddsource", "ddtags", "hostname", "message", "service"})
	} else {
		return err
	}
	o.Ddsource = all.Ddsource
	o.Ddtags = all.Ddtags
	o.Hostname = all.Hostname
	o.Message = all.Message
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
