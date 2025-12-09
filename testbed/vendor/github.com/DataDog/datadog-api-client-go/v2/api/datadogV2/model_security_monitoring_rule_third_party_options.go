// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleThirdPartyOptions Options on third party detection method.
type SecurityMonitoringRuleThirdPartyOptions struct {
	// Notification targets for the logs that do not correspond to any of the cases.
	DefaultNotifications []string `json:"defaultNotifications,omitempty"`
	// Severity of the Security Signal.
	DefaultStatus *SecurityMonitoringRuleSeverity `json:"defaultStatus,omitempty"`
	// Queries to be combined with third party case queries. Each of them can have different group by fields, to aggregate differently based on the type of alert.
	RootQueries []SecurityMonitoringThirdPartyRootQuery `json:"rootQueries,omitempty"`
	// A template for the signal title; if omitted, the title is generated based on the case name.
	SignalTitleTemplate *string `json:"signalTitleTemplate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleThirdPartyOptions instantiates a new SecurityMonitoringRuleThirdPartyOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleThirdPartyOptions() *SecurityMonitoringRuleThirdPartyOptions {
	this := SecurityMonitoringRuleThirdPartyOptions{}
	return &this
}

// NewSecurityMonitoringRuleThirdPartyOptionsWithDefaults instantiates a new SecurityMonitoringRuleThirdPartyOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleThirdPartyOptionsWithDefaults() *SecurityMonitoringRuleThirdPartyOptions {
	this := SecurityMonitoringRuleThirdPartyOptions{}
	return &this
}

// GetDefaultNotifications returns the DefaultNotifications field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultNotifications() []string {
	if o == nil || o.DefaultNotifications == nil {
		var ret []string
		return ret
	}
	return o.DefaultNotifications
}

// GetDefaultNotificationsOk returns a tuple with the DefaultNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultNotificationsOk() (*[]string, bool) {
	if o == nil || o.DefaultNotifications == nil {
		return nil, false
	}
	return &o.DefaultNotifications, true
}

// HasDefaultNotifications returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) HasDefaultNotifications() bool {
	return o != nil && o.DefaultNotifications != nil
}

// SetDefaultNotifications gets a reference to the given []string and assigns it to the DefaultNotifications field.
func (o *SecurityMonitoringRuleThirdPartyOptions) SetDefaultNotifications(v []string) {
	o.DefaultNotifications = v
}

// GetDefaultStatus returns the DefaultStatus field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultStatus() SecurityMonitoringRuleSeverity {
	if o == nil || o.DefaultStatus == nil {
		var ret SecurityMonitoringRuleSeverity
		return ret
	}
	return *o.DefaultStatus
}

// GetDefaultStatusOk returns a tuple with the DefaultStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultStatusOk() (*SecurityMonitoringRuleSeverity, bool) {
	if o == nil || o.DefaultStatus == nil {
		return nil, false
	}
	return o.DefaultStatus, true
}

// HasDefaultStatus returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) HasDefaultStatus() bool {
	return o != nil && o.DefaultStatus != nil
}

// SetDefaultStatus gets a reference to the given SecurityMonitoringRuleSeverity and assigns it to the DefaultStatus field.
func (o *SecurityMonitoringRuleThirdPartyOptions) SetDefaultStatus(v SecurityMonitoringRuleSeverity) {
	o.DefaultStatus = &v
}

// GetRootQueries returns the RootQueries field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetRootQueries() []SecurityMonitoringThirdPartyRootQuery {
	if o == nil || o.RootQueries == nil {
		var ret []SecurityMonitoringThirdPartyRootQuery
		return ret
	}
	return o.RootQueries
}

// GetRootQueriesOk returns a tuple with the RootQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetRootQueriesOk() (*[]SecurityMonitoringThirdPartyRootQuery, bool) {
	if o == nil || o.RootQueries == nil {
		return nil, false
	}
	return &o.RootQueries, true
}

// HasRootQueries returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) HasRootQueries() bool {
	return o != nil && o.RootQueries != nil
}

// SetRootQueries gets a reference to the given []SecurityMonitoringThirdPartyRootQuery and assigns it to the RootQueries field.
func (o *SecurityMonitoringRuleThirdPartyOptions) SetRootQueries(v []SecurityMonitoringThirdPartyRootQuery) {
	o.RootQueries = v
}

// GetSignalTitleTemplate returns the SignalTitleTemplate field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetSignalTitleTemplate() string {
	if o == nil || o.SignalTitleTemplate == nil {
		var ret string
		return ret
	}
	return *o.SignalTitleTemplate
}

// GetSignalTitleTemplateOk returns a tuple with the SignalTitleTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) GetSignalTitleTemplateOk() (*string, bool) {
	if o == nil || o.SignalTitleTemplate == nil {
		return nil, false
	}
	return o.SignalTitleTemplate, true
}

// HasSignalTitleTemplate returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleThirdPartyOptions) HasSignalTitleTemplate() bool {
	return o != nil && o.SignalTitleTemplate != nil
}

// SetSignalTitleTemplate gets a reference to the given string and assigns it to the SignalTitleTemplate field.
func (o *SecurityMonitoringRuleThirdPartyOptions) SetSignalTitleTemplate(v string) {
	o.SignalTitleTemplate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleThirdPartyOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultNotifications != nil {
		toSerialize["defaultNotifications"] = o.DefaultNotifications
	}
	if o.DefaultStatus != nil {
		toSerialize["defaultStatus"] = o.DefaultStatus
	}
	if o.RootQueries != nil {
		toSerialize["rootQueries"] = o.RootQueries
	}
	if o.SignalTitleTemplate != nil {
		toSerialize["signalTitleTemplate"] = o.SignalTitleTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleThirdPartyOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultNotifications []string                                `json:"defaultNotifications,omitempty"`
		DefaultStatus        *SecurityMonitoringRuleSeverity         `json:"defaultStatus,omitempty"`
		RootQueries          []SecurityMonitoringThirdPartyRootQuery `json:"rootQueries,omitempty"`
		SignalTitleTemplate  *string                                 `json:"signalTitleTemplate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"defaultNotifications", "defaultStatus", "rootQueries", "signalTitleTemplate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DefaultNotifications = all.DefaultNotifications
	if all.DefaultStatus != nil && !all.DefaultStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DefaultStatus = all.DefaultStatus
	}
	o.RootQueries = all.RootQueries
	o.SignalTitleTemplate = all.SignalTitleTemplate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
