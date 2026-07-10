// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleResponseAttributesDeliveryFormat The delivery format for dashboard report schedules, or `null` if not set.
type ReportScheduleResponseAttributesDeliveryFormat string

// List of ReportScheduleResponseAttributesDeliveryFormat.
const (
	REPORTSCHEDULERESPONSEATTRIBUTESDELIVERYFORMAT_PDF         ReportScheduleResponseAttributesDeliveryFormat = "pdf"
	REPORTSCHEDULERESPONSEATTRIBUTESDELIVERYFORMAT_PNG         ReportScheduleResponseAttributesDeliveryFormat = "png"
	REPORTSCHEDULERESPONSEATTRIBUTESDELIVERYFORMAT_PDF_AND_PNG ReportScheduleResponseAttributesDeliveryFormat = "pdf_and_png"
)

var allowedReportScheduleResponseAttributesDeliveryFormatEnumValues = []ReportScheduleResponseAttributesDeliveryFormat{
	REPORTSCHEDULERESPONSEATTRIBUTESDELIVERYFORMAT_PDF,
	REPORTSCHEDULERESPONSEATTRIBUTESDELIVERYFORMAT_PNG,
	REPORTSCHEDULERESPONSEATTRIBUTESDELIVERYFORMAT_PDF_AND_PNG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleResponseAttributesDeliveryFormat) GetAllowedValues() []ReportScheduleResponseAttributesDeliveryFormat {
	return allowedReportScheduleResponseAttributesDeliveryFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleResponseAttributesDeliveryFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleResponseAttributesDeliveryFormat(value)
	return nil
}

// NewReportScheduleResponseAttributesDeliveryFormatFromValue returns a pointer to a valid ReportScheduleResponseAttributesDeliveryFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleResponseAttributesDeliveryFormatFromValue(v string) (*ReportScheduleResponseAttributesDeliveryFormat, error) {
	ev := ReportScheduleResponseAttributesDeliveryFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleResponseAttributesDeliveryFormat: valid values are %v", v, allowedReportScheduleResponseAttributesDeliveryFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleResponseAttributesDeliveryFormat) IsValid() bool {
	for _, existing := range allowedReportScheduleResponseAttributesDeliveryFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleResponseAttributesDeliveryFormat value.
func (v ReportScheduleResponseAttributesDeliveryFormat) Ptr() *ReportScheduleResponseAttributesDeliveryFormat {
	return &v
}

// NullableReportScheduleResponseAttributesDeliveryFormat handles when a null is used for ReportScheduleResponseAttributesDeliveryFormat.
type NullableReportScheduleResponseAttributesDeliveryFormat struct {
	value *ReportScheduleResponseAttributesDeliveryFormat
	isSet bool
}

// Get returns the associated value.
func (v NullableReportScheduleResponseAttributesDeliveryFormat) Get() *ReportScheduleResponseAttributesDeliveryFormat {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableReportScheduleResponseAttributesDeliveryFormat) Set(val *ReportScheduleResponseAttributesDeliveryFormat) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableReportScheduleResponseAttributesDeliveryFormat) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableReportScheduleResponseAttributesDeliveryFormat) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableReportScheduleResponseAttributesDeliveryFormat initializes the struct as if Set has been called.
func NewNullableReportScheduleResponseAttributesDeliveryFormat(val *ReportScheduleResponseAttributesDeliveryFormat) *NullableReportScheduleResponseAttributesDeliveryFormat {
	return &NullableReportScheduleResponseAttributesDeliveryFormat{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableReportScheduleResponseAttributesDeliveryFormat) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableReportScheduleResponseAttributesDeliveryFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
