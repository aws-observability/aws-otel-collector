// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3APIVersion The version of the schema data that was used to populate this entity's data. This could be via the API, Terraform, or YAML file in a repository. The field is known as schema-version in the previous version.
type EntityV3APIVersion string

// List of EntityV3APIVersion.
const (
	ENTITYV3APIVERSION_V3   EntityV3APIVersion = "v3"
	ENTITYV3APIVERSION_V2_2 EntityV3APIVersion = "v2.2"
	ENTITYV3APIVERSION_V2_1 EntityV3APIVersion = "v2.1"
	ENTITYV3APIVERSION_V2   EntityV3APIVersion = "v2"
)

var allowedEntityV3APIVersionEnumValues = []EntityV3APIVersion{
	ENTITYV3APIVERSION_V3,
	ENTITYV3APIVERSION_V2_2,
	ENTITYV3APIVERSION_V2_1,
	ENTITYV3APIVERSION_V2,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EntityV3APIVersion) GetAllowedValues() []EntityV3APIVersion {
	return allowedEntityV3APIVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3APIVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3APIVersion(value)
	return nil
}

// NewEntityV3APIVersionFromValue returns a pointer to a valid EntityV3APIVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3APIVersionFromValue(v string) (*EntityV3APIVersion, error) {
	ev := EntityV3APIVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3APIVersion: valid values are %v", v, allowedEntityV3APIVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3APIVersion) IsValid() bool {
	for _, existing := range allowedEntityV3APIVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3APIVersion value.
func (v EntityV3APIVersion) Ptr() *EntityV3APIVersion {
	return &v
}
