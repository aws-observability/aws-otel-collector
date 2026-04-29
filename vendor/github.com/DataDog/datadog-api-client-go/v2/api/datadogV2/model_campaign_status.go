// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CampaignStatus The status of the campaign.
type CampaignStatus string

// List of CampaignStatus.
const (
	CAMPAIGNSTATUS_IN_PROGRESS CampaignStatus = "in_progress"
	CAMPAIGNSTATUS_NOT_STARTED CampaignStatus = "not_started"
	CAMPAIGNSTATUS_COMPLETED   CampaignStatus = "completed"
)

var allowedCampaignStatusEnumValues = []CampaignStatus{
	CAMPAIGNSTATUS_IN_PROGRESS,
	CAMPAIGNSTATUS_NOT_STARTED,
	CAMPAIGNSTATUS_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CampaignStatus) GetAllowedValues() []CampaignStatus {
	return allowedCampaignStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CampaignStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CampaignStatus(value)
	return nil
}

// NewCampaignStatusFromValue returns a pointer to a valid CampaignStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCampaignStatusFromValue(v string) (*CampaignStatus, error) {
	ev := CampaignStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CampaignStatus: valid values are %v", v, allowedCampaignStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CampaignStatus) IsValid() bool {
	for _, existing := range allowedCampaignStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CampaignStatus value.
func (v CampaignStatus) Ptr() *CampaignStatus {
	return &v
}
