// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPUsageCostConfigPostRequestType Type of Google Cloud Usage Cost config post request.
type GCPUsageCostConfigPostRequestType string

// List of GCPUsageCostConfigPostRequestType.
const (
	GCPUSAGECOSTCONFIGPOSTREQUESTTYPE_GCP_USAGE_COST_CONFIG_POST_REQUEST GCPUsageCostConfigPostRequestType = "gcp_uc_config_post_request"
)

var allowedGCPUsageCostConfigPostRequestTypeEnumValues = []GCPUsageCostConfigPostRequestType{
	GCPUSAGECOSTCONFIGPOSTREQUESTTYPE_GCP_USAGE_COST_CONFIG_POST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GCPUsageCostConfigPostRequestType) GetAllowedValues() []GCPUsageCostConfigPostRequestType {
	return allowedGCPUsageCostConfigPostRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GCPUsageCostConfigPostRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GCPUsageCostConfigPostRequestType(value)
	return nil
}

// NewGCPUsageCostConfigPostRequestTypeFromValue returns a pointer to a valid GCPUsageCostConfigPostRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGCPUsageCostConfigPostRequestTypeFromValue(v string) (*GCPUsageCostConfigPostRequestType, error) {
	ev := GCPUsageCostConfigPostRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GCPUsageCostConfigPostRequestType: valid values are %v", v, allowedGCPUsageCostConfigPostRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GCPUsageCostConfigPostRequestType) IsValid() bool {
	for _, existing := range allowedGCPUsageCostConfigPostRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GCPUsageCostConfigPostRequestType value.
func (v GCPUsageCostConfigPostRequestType) Ptr() *GCPUsageCostConfigPostRequestType {
	return &v
}
