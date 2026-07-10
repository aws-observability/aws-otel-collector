// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleDeliveryFormat How a PDF-export report is delivered. `pdf` attaches a PDF file, `png` embeds
// an inline PNG image, and `pdf_and_png` delivers both.
type ReportScheduleDeliveryFormat string

// List of ReportScheduleDeliveryFormat.
const (
	REPORTSCHEDULEDELIVERYFORMAT_PDF         ReportScheduleDeliveryFormat = "pdf"
	REPORTSCHEDULEDELIVERYFORMAT_PNG         ReportScheduleDeliveryFormat = "png"
	REPORTSCHEDULEDELIVERYFORMAT_PDF_AND_PNG ReportScheduleDeliveryFormat = "pdf_and_png"
)

var allowedReportScheduleDeliveryFormatEnumValues = []ReportScheduleDeliveryFormat{
	REPORTSCHEDULEDELIVERYFORMAT_PDF,
	REPORTSCHEDULEDELIVERYFORMAT_PNG,
	REPORTSCHEDULEDELIVERYFORMAT_PDF_AND_PNG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleDeliveryFormat) GetAllowedValues() []ReportScheduleDeliveryFormat {
	return allowedReportScheduleDeliveryFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleDeliveryFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleDeliveryFormat(value)
	return nil
}

// NewReportScheduleDeliveryFormatFromValue returns a pointer to a valid ReportScheduleDeliveryFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleDeliveryFormatFromValue(v string) (*ReportScheduleDeliveryFormat, error) {
	ev := ReportScheduleDeliveryFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleDeliveryFormat: valid values are %v", v, allowedReportScheduleDeliveryFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleDeliveryFormat) IsValid() bool {
	for _, existing := range allowedReportScheduleDeliveryFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleDeliveryFormat value.
func (v ReportScheduleDeliveryFormat) Ptr() *ReportScheduleDeliveryFormat {
	return &v
}
