// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerProduct Datadog product onto which Sensitive Data Scanner can be activated.
type SensitiveDataScannerProduct string

// List of SensitiveDataScannerProduct.
const (
	SENSITIVEDATASCANNERPRODUCT_LOGS   SensitiveDataScannerProduct = "logs"
	SENSITIVEDATASCANNERPRODUCT_RUM    SensitiveDataScannerProduct = "rum"
	SENSITIVEDATASCANNERPRODUCT_EVENTS SensitiveDataScannerProduct = "events"
	SENSITIVEDATASCANNERPRODUCT_APM    SensitiveDataScannerProduct = "apm"
)

var allowedSensitiveDataScannerProductEnumValues = []SensitiveDataScannerProduct{
	SENSITIVEDATASCANNERPRODUCT_LOGS,
	SENSITIVEDATASCANNERPRODUCT_RUM,
	SENSITIVEDATASCANNERPRODUCT_EVENTS,
	SENSITIVEDATASCANNERPRODUCT_APM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SensitiveDataScannerProduct) GetAllowedValues() []SensitiveDataScannerProduct {
	return allowedSensitiveDataScannerProductEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SensitiveDataScannerProduct) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SensitiveDataScannerProduct(value)
	return nil
}

// NewSensitiveDataScannerProductFromValue returns a pointer to a valid SensitiveDataScannerProduct
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSensitiveDataScannerProductFromValue(v string) (*SensitiveDataScannerProduct, error) {
	ev := SensitiveDataScannerProduct(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SensitiveDataScannerProduct: valid values are %v", v, allowedSensitiveDataScannerProductEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SensitiveDataScannerProduct) IsValid() bool {
	for _, existing := range allowedSensitiveDataScannerProductEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitiveDataScannerProduct value.
func (v SensitiveDataScannerProduct) Ptr() *SensitiveDataScannerProduct {
	return &v
}
