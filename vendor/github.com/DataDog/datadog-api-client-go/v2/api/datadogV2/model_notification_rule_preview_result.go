// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRulePreviewResult The preview result for a single rule type.
type NotificationRulePreviewResult struct {
	// The notification status for the given rule type. `SUCCESS` means a matching event was found and the notification was sent successfully. `DEFAULT` means no matching event was found and a default placeholder notification was sent instead. `ERROR` means an error occurred while sending the notification.
	NotificationStatus NotificationRulePreviewNotificationStatus `json:"notification_status"`
	// Security rule type which can be used in security rules.
	// Signal-based notification rules can filter signals based on rule types application_security, log_detection,
	// workload_security, signal_correlation, cloud_configuration and infrastructure_configuration.
	// Vulnerability-based notification rules can filter vulnerabilities based on rule types application_code_vulnerability,
	// application_library_vulnerability, attack_path, container_image_vulnerability, identity_risk, misconfiguration,
	// api_security, host_vulnerability, iac_misconfiguration, sast_vulnerability and secret_vulnerability.
	RuleType RuleTypesItems `json:"rule_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationRulePreviewResult instantiates a new NotificationRulePreviewResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationRulePreviewResult(notificationStatus NotificationRulePreviewNotificationStatus, ruleType RuleTypesItems) *NotificationRulePreviewResult {
	this := NotificationRulePreviewResult{}
	this.NotificationStatus = notificationStatus
	this.RuleType = ruleType
	return &this
}

// NewNotificationRulePreviewResultWithDefaults instantiates a new NotificationRulePreviewResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationRulePreviewResultWithDefaults() *NotificationRulePreviewResult {
	this := NotificationRulePreviewResult{}
	return &this
}

// GetNotificationStatus returns the NotificationStatus field value.
func (o *NotificationRulePreviewResult) GetNotificationStatus() NotificationRulePreviewNotificationStatus {
	if o == nil {
		var ret NotificationRulePreviewNotificationStatus
		return ret
	}
	return o.NotificationStatus
}

// GetNotificationStatusOk returns a tuple with the NotificationStatus field value
// and a boolean to check if the value has been set.
func (o *NotificationRulePreviewResult) GetNotificationStatusOk() (*NotificationRulePreviewNotificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationStatus, true
}

// SetNotificationStatus sets field value.
func (o *NotificationRulePreviewResult) SetNotificationStatus(v NotificationRulePreviewNotificationStatus) {
	o.NotificationStatus = v
}

// GetRuleType returns the RuleType field value.
func (o *NotificationRulePreviewResult) GetRuleType() RuleTypesItems {
	if o == nil {
		var ret RuleTypesItems
		return ret
	}
	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *NotificationRulePreviewResult) GetRuleTypeOk() (*RuleTypesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value.
func (o *NotificationRulePreviewResult) SetRuleType(v RuleTypesItems) {
	o.RuleType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationRulePreviewResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["notification_status"] = o.NotificationStatus
	toSerialize["rule_type"] = o.RuleType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationRulePreviewResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NotificationStatus *NotificationRulePreviewNotificationStatus `json:"notification_status"`
		RuleType           *RuleTypesItems                            `json:"rule_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NotificationStatus == nil {
		return fmt.Errorf("required field notification_status missing")
	}
	if all.RuleType == nil {
		return fmt.Errorf("required field rule_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"notification_status", "rule_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.NotificationStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.NotificationStatus = *all.NotificationStatus
	}
	if !all.RuleType.IsValid() {
		hasInvalidField = true
	} else {
		o.RuleType = *all.RuleType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
