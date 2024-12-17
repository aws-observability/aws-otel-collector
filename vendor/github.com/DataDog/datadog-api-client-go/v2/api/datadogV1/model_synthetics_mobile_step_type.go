// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepType Step type used in your mobile Synthetic test.
type SyntheticsMobileStepType string

// List of SyntheticsMobileStepType.
const (
	SYNTHETICSMOBILESTEPTYPE_ASSERTELEMENTCONTENT SyntheticsMobileStepType = "assertElementContent"
	SYNTHETICSMOBILESTEPTYPE_ASSERTSCREENCONTAINS SyntheticsMobileStepType = "assertScreenContains"
	SYNTHETICSMOBILESTEPTYPE_ASSERTSCREENLACKS    SyntheticsMobileStepType = "assertScreenLacks"
	SYNTHETICSMOBILESTEPTYPE_DOUBLETAP            SyntheticsMobileStepType = "doubleTap"
	SYNTHETICSMOBILESTEPTYPE_EXTRACTVARIABLE      SyntheticsMobileStepType = "extractVariable"
	SYNTHETICSMOBILESTEPTYPE_FLICK                SyntheticsMobileStepType = "flick"
	SYNTHETICSMOBILESTEPTYPE_OPENDEEPLINK         SyntheticsMobileStepType = "openDeeplink"
	SYNTHETICSMOBILESTEPTYPE_PLAYSUBTEST          SyntheticsMobileStepType = "playSubTest"
	SYNTHETICSMOBILESTEPTYPE_PRESSBACK            SyntheticsMobileStepType = "pressBack"
	SYNTHETICSMOBILESTEPTYPE_RESTARTAPPLICATION   SyntheticsMobileStepType = "restartApplication"
	SYNTHETICSMOBILESTEPTYPE_ROTATE               SyntheticsMobileStepType = "rotate"
	SYNTHETICSMOBILESTEPTYPE_SCROLL               SyntheticsMobileStepType = "scroll"
	SYNTHETICSMOBILESTEPTYPE_SCROLLTOELEMENT      SyntheticsMobileStepType = "scrollToElement"
	SYNTHETICSMOBILESTEPTYPE_TAP                  SyntheticsMobileStepType = "tap"
	SYNTHETICSMOBILESTEPTYPE_TOGGLEWIFI           SyntheticsMobileStepType = "toggleWiFi"
	SYNTHETICSMOBILESTEPTYPE_TYPETEXT             SyntheticsMobileStepType = "typeText"
	SYNTHETICSMOBILESTEPTYPE_WAIT                 SyntheticsMobileStepType = "wait"
)

var allowedSyntheticsMobileStepTypeEnumValues = []SyntheticsMobileStepType{
	SYNTHETICSMOBILESTEPTYPE_ASSERTELEMENTCONTENT,
	SYNTHETICSMOBILESTEPTYPE_ASSERTSCREENCONTAINS,
	SYNTHETICSMOBILESTEPTYPE_ASSERTSCREENLACKS,
	SYNTHETICSMOBILESTEPTYPE_DOUBLETAP,
	SYNTHETICSMOBILESTEPTYPE_EXTRACTVARIABLE,
	SYNTHETICSMOBILESTEPTYPE_FLICK,
	SYNTHETICSMOBILESTEPTYPE_OPENDEEPLINK,
	SYNTHETICSMOBILESTEPTYPE_PLAYSUBTEST,
	SYNTHETICSMOBILESTEPTYPE_PRESSBACK,
	SYNTHETICSMOBILESTEPTYPE_RESTARTAPPLICATION,
	SYNTHETICSMOBILESTEPTYPE_ROTATE,
	SYNTHETICSMOBILESTEPTYPE_SCROLL,
	SYNTHETICSMOBILESTEPTYPE_SCROLLTOELEMENT,
	SYNTHETICSMOBILESTEPTYPE_TAP,
	SYNTHETICSMOBILESTEPTYPE_TOGGLEWIFI,
	SYNTHETICSMOBILESTEPTYPE_TYPETEXT,
	SYNTHETICSMOBILESTEPTYPE_WAIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileStepType) GetAllowedValues() []SyntheticsMobileStepType {
	return allowedSyntheticsMobileStepTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileStepType(value)
	return nil
}

// NewSyntheticsMobileStepTypeFromValue returns a pointer to a valid SyntheticsMobileStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileStepTypeFromValue(v string) (*SyntheticsMobileStepType, error) {
	ev := SyntheticsMobileStepType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileStepType: valid values are %v", v, allowedSyntheticsMobileStepTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileStepType) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileStepType value.
func (v SyntheticsMobileStepType) Ptr() *SyntheticsMobileStepType {
	return &v
}
