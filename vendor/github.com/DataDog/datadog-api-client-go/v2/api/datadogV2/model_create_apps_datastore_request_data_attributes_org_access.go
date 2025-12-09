// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreRequestDataAttributesOrgAccess The organization access level for the datastore. For example, 'contributor'.
type CreateAppsDatastoreRequestDataAttributesOrgAccess string

// List of CreateAppsDatastoreRequestDataAttributesOrgAccess.
const (
	CREATEAPPSDATASTOREREQUESTDATAATTRIBUTESORGACCESS_CONTRIBUTOR CreateAppsDatastoreRequestDataAttributesOrgAccess = "contributor"
	CREATEAPPSDATASTOREREQUESTDATAATTRIBUTESORGACCESS_VIEWER      CreateAppsDatastoreRequestDataAttributesOrgAccess = "viewer"
	CREATEAPPSDATASTOREREQUESTDATAATTRIBUTESORGACCESS_MANAGER     CreateAppsDatastoreRequestDataAttributesOrgAccess = "manager"
)

var allowedCreateAppsDatastoreRequestDataAttributesOrgAccessEnumValues = []CreateAppsDatastoreRequestDataAttributesOrgAccess{
	CREATEAPPSDATASTOREREQUESTDATAATTRIBUTESORGACCESS_CONTRIBUTOR,
	CREATEAPPSDATASTOREREQUESTDATAATTRIBUTESORGACCESS_VIEWER,
	CREATEAPPSDATASTOREREQUESTDATAATTRIBUTESORGACCESS_MANAGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppsDatastoreRequestDataAttributesOrgAccess) GetAllowedValues() []CreateAppsDatastoreRequestDataAttributesOrgAccess {
	return allowedCreateAppsDatastoreRequestDataAttributesOrgAccessEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppsDatastoreRequestDataAttributesOrgAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppsDatastoreRequestDataAttributesOrgAccess(value)
	return nil
}

// NewCreateAppsDatastoreRequestDataAttributesOrgAccessFromValue returns a pointer to a valid CreateAppsDatastoreRequestDataAttributesOrgAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppsDatastoreRequestDataAttributesOrgAccessFromValue(v string) (*CreateAppsDatastoreRequestDataAttributesOrgAccess, error) {
	ev := CreateAppsDatastoreRequestDataAttributesOrgAccess(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppsDatastoreRequestDataAttributesOrgAccess: valid values are %v", v, allowedCreateAppsDatastoreRequestDataAttributesOrgAccessEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppsDatastoreRequestDataAttributesOrgAccess) IsValid() bool {
	for _, existing := range allowedCreateAppsDatastoreRequestDataAttributesOrgAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppsDatastoreRequestDataAttributesOrgAccess value.
func (v CreateAppsDatastoreRequestDataAttributesOrgAccess) Ptr() *CreateAppsDatastoreRequestDataAttributesOrgAccess {
	return &v
}
