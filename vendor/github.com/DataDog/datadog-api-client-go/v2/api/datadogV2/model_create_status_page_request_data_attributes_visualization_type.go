// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateStatusPageRequestDataAttributesVisualizationType The visualization type of the status page.
type CreateStatusPageRequestDataAttributesVisualizationType string

// List of CreateStatusPageRequestDataAttributesVisualizationType.
const (
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_BARS_AND_UPTIME_PERCENTAGE CreateStatusPageRequestDataAttributesVisualizationType = "bars_and_uptime_percentage"
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_BARS_ONLY                  CreateStatusPageRequestDataAttributesVisualizationType = "bars_only"
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_COMPONENT_NAME_ONLY        CreateStatusPageRequestDataAttributesVisualizationType = "component_name_only"
)

var allowedCreateStatusPageRequestDataAttributesVisualizationTypeEnumValues = []CreateStatusPageRequestDataAttributesVisualizationType{
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_BARS_AND_UPTIME_PERCENTAGE,
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_BARS_ONLY,
	CREATESTATUSPAGEREQUESTDATAATTRIBUTESVISUALIZATIONTYPE_COMPONENT_NAME_ONLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateStatusPageRequestDataAttributesVisualizationType) GetAllowedValues() []CreateStatusPageRequestDataAttributesVisualizationType {
	return allowedCreateStatusPageRequestDataAttributesVisualizationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateStatusPageRequestDataAttributesVisualizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateStatusPageRequestDataAttributesVisualizationType(value)
	return nil
}

// NewCreateStatusPageRequestDataAttributesVisualizationTypeFromValue returns a pointer to a valid CreateStatusPageRequestDataAttributesVisualizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateStatusPageRequestDataAttributesVisualizationTypeFromValue(v string) (*CreateStatusPageRequestDataAttributesVisualizationType, error) {
	ev := CreateStatusPageRequestDataAttributesVisualizationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateStatusPageRequestDataAttributesVisualizationType: valid values are %v", v, allowedCreateStatusPageRequestDataAttributesVisualizationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateStatusPageRequestDataAttributesVisualizationType) IsValid() bool {
	for _, existing := range allowedCreateStatusPageRequestDataAttributesVisualizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateStatusPageRequestDataAttributesVisualizationType value.
func (v CreateStatusPageRequestDataAttributesVisualizationType) Ptr() *CreateStatusPageRequestDataAttributesVisualizationType {
	return &v
}
