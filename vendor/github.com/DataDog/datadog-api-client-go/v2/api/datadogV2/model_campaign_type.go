// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CampaignType The JSON:API type for campaigns.
type CampaignType string

// List of CampaignType.
const (
	CAMPAIGNTYPE_CAMPAIGN CampaignType = "campaign"
)

var allowedCampaignTypeEnumValues = []CampaignType{
	CAMPAIGNTYPE_CAMPAIGN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CampaignType) GetAllowedValues() []CampaignType {
	return allowedCampaignTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CampaignType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CampaignType(value)
	return nil
}

// NewCampaignTypeFromValue returns a pointer to a valid CampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCampaignTypeFromValue(v string) (*CampaignType, error) {
	ev := CampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CampaignType: valid values are %v", v, allowedCampaignTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CampaignType) IsValid() bool {
	for _, existing := range allowedCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignType value.
func (v CampaignType) Ptr() *CampaignType {
	return &v
}
