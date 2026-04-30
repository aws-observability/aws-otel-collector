// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedGlobalTimeLiveSpan Dashboard global time live_span selection.
type SecureEmbedGlobalTimeLiveSpan string

// List of SecureEmbedGlobalTimeLiveSpan.
const (
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_FIFTEEN_MINUTES SecureEmbedGlobalTimeLiveSpan = "15m"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_HOUR        SecureEmbedGlobalTimeLiveSpan = "1h"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_FOUR_HOURS      SecureEmbedGlobalTimeLiveSpan = "4h"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_DAY         SecureEmbedGlobalTimeLiveSpan = "1d"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_TWO_DAYS        SecureEmbedGlobalTimeLiveSpan = "2d"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_WEEK        SecureEmbedGlobalTimeLiveSpan = "1w"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_MONTH       SecureEmbedGlobalTimeLiveSpan = "1mo"
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_THREE_MONTHS    SecureEmbedGlobalTimeLiveSpan = "3mo"
)

var allowedSecureEmbedGlobalTimeLiveSpanEnumValues = []SecureEmbedGlobalTimeLiveSpan{
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_FIFTEEN_MINUTES,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_HOUR,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_FOUR_HOURS,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_DAY,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_TWO_DAYS,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_WEEK,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_ONE_MONTH,
	SECUREEMBEDGLOBALTIMELIVESPAN_PAST_THREE_MONTHS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedGlobalTimeLiveSpan) GetAllowedValues() []SecureEmbedGlobalTimeLiveSpan {
	return allowedSecureEmbedGlobalTimeLiveSpanEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedGlobalTimeLiveSpan) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedGlobalTimeLiveSpan(value)
	return nil
}

// NewSecureEmbedGlobalTimeLiveSpanFromValue returns a pointer to a valid SecureEmbedGlobalTimeLiveSpan
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedGlobalTimeLiveSpanFromValue(v string) (*SecureEmbedGlobalTimeLiveSpan, error) {
	ev := SecureEmbedGlobalTimeLiveSpan(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedGlobalTimeLiveSpan: valid values are %v", v, allowedSecureEmbedGlobalTimeLiveSpanEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedGlobalTimeLiveSpan) IsValid() bool {
	for _, existing := range allowedSecureEmbedGlobalTimeLiveSpanEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedGlobalTimeLiveSpan value.
func (v SecureEmbedGlobalTimeLiveSpan) Ptr() *SecureEmbedGlobalTimeLiveSpan {
	return &v
}
