// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPhoneNotificationRuleMethod Specifies the method in which a phone is used in a notification rule
type OnCallPhoneNotificationRuleMethod string

// List of OnCallPhoneNotificationRuleMethod.
const (
	ONCALLPHONENOTIFICATIONRULEMETHOD_SMS   OnCallPhoneNotificationRuleMethod = "sms"
	ONCALLPHONENOTIFICATIONRULEMETHOD_VOICE OnCallPhoneNotificationRuleMethod = "voice"
)

var allowedOnCallPhoneNotificationRuleMethodEnumValues = []OnCallPhoneNotificationRuleMethod{
	ONCALLPHONENOTIFICATIONRULEMETHOD_SMS,
	ONCALLPHONENOTIFICATIONRULEMETHOD_VOICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPhoneNotificationRuleMethod) GetAllowedValues() []OnCallPhoneNotificationRuleMethod {
	return allowedOnCallPhoneNotificationRuleMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPhoneNotificationRuleMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPhoneNotificationRuleMethod(value)
	return nil
}

// NewOnCallPhoneNotificationRuleMethodFromValue returns a pointer to a valid OnCallPhoneNotificationRuleMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPhoneNotificationRuleMethodFromValue(v string) (*OnCallPhoneNotificationRuleMethod, error) {
	ev := OnCallPhoneNotificationRuleMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPhoneNotificationRuleMethod: valid values are %v", v, allowedOnCallPhoneNotificationRuleMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPhoneNotificationRuleMethod) IsValid() bool {
	for _, existing := range allowedOnCallPhoneNotificationRuleMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPhoneNotificationRuleMethod value.
func (v OnCallPhoneNotificationRuleMethod) Ptr() *OnCallPhoneNotificationRuleMethod {
	return &v
}
