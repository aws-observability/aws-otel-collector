// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareIntegrationType The definition of the `CloudflareIntegrationType` object.
type CloudflareIntegrationType string

// List of CloudflareIntegrationType.
const (
	CLOUDFLAREINTEGRATIONTYPE_CLOUDFLARE CloudflareIntegrationType = "Cloudflare"
)

var allowedCloudflareIntegrationTypeEnumValues = []CloudflareIntegrationType{
	CLOUDFLAREINTEGRATIONTYPE_CLOUDFLARE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudflareIntegrationType) GetAllowedValues() []CloudflareIntegrationType {
	return allowedCloudflareIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudflareIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudflareIntegrationType(value)
	return nil
}

// NewCloudflareIntegrationTypeFromValue returns a pointer to a valid CloudflareIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudflareIntegrationTypeFromValue(v string) (*CloudflareIntegrationType, error) {
	ev := CloudflareIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudflareIntegrationType: valid values are %v", v, allowedCloudflareIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudflareIntegrationType) IsValid() bool {
	for _, existing := range allowedCloudflareIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudflareIntegrationType value.
func (v CloudflareIntegrationType) Ptr() *CloudflareIntegrationType {
	return &v
}
