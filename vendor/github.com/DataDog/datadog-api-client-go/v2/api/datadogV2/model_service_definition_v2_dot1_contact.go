// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1Contact - Service owner's contacts information.
type ServiceDefinitionV2Dot1Contact struct {
	ServiceDefinitionV2Dot1Email   *ServiceDefinitionV2Dot1Email
	ServiceDefinitionV2Dot1Slack   *ServiceDefinitionV2Dot1Slack
	ServiceDefinitionV2Dot1MSTeams *ServiceDefinitionV2Dot1MSTeams

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ServiceDefinitionV2Dot1EmailAsServiceDefinitionV2Dot1Contact is a convenience function that returns ServiceDefinitionV2Dot1Email wrapped in ServiceDefinitionV2Dot1Contact.
func ServiceDefinitionV2Dot1EmailAsServiceDefinitionV2Dot1Contact(v *ServiceDefinitionV2Dot1Email) ServiceDefinitionV2Dot1Contact {
	return ServiceDefinitionV2Dot1Contact{ServiceDefinitionV2Dot1Email: v}
}

// ServiceDefinitionV2Dot1SlackAsServiceDefinitionV2Dot1Contact is a convenience function that returns ServiceDefinitionV2Dot1Slack wrapped in ServiceDefinitionV2Dot1Contact.
func ServiceDefinitionV2Dot1SlackAsServiceDefinitionV2Dot1Contact(v *ServiceDefinitionV2Dot1Slack) ServiceDefinitionV2Dot1Contact {
	return ServiceDefinitionV2Dot1Contact{ServiceDefinitionV2Dot1Slack: v}
}

// ServiceDefinitionV2Dot1MSTeamsAsServiceDefinitionV2Dot1Contact is a convenience function that returns ServiceDefinitionV2Dot1MSTeams wrapped in ServiceDefinitionV2Dot1Contact.
func ServiceDefinitionV2Dot1MSTeamsAsServiceDefinitionV2Dot1Contact(v *ServiceDefinitionV2Dot1MSTeams) ServiceDefinitionV2Dot1Contact {
	return ServiceDefinitionV2Dot1Contact{ServiceDefinitionV2Dot1MSTeams: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ServiceDefinitionV2Dot1Contact) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceDefinitionV2Dot1Email
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Dot1Email)
	if err == nil {
		if obj.ServiceDefinitionV2Dot1Email != nil && obj.ServiceDefinitionV2Dot1Email.UnparsedObject == nil {
			jsonServiceDefinitionV2Dot1Email, _ := datadog.Marshal(obj.ServiceDefinitionV2Dot1Email)
			if string(jsonServiceDefinitionV2Dot1Email) == "{}" { // empty struct
				obj.ServiceDefinitionV2Dot1Email = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Dot1Email = nil
		}
	} else {
		obj.ServiceDefinitionV2Dot1Email = nil
	}

	// try to unmarshal data into ServiceDefinitionV2Dot1Slack
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Dot1Slack)
	if err == nil {
		if obj.ServiceDefinitionV2Dot1Slack != nil && obj.ServiceDefinitionV2Dot1Slack.UnparsedObject == nil {
			jsonServiceDefinitionV2Dot1Slack, _ := datadog.Marshal(obj.ServiceDefinitionV2Dot1Slack)
			if string(jsonServiceDefinitionV2Dot1Slack) == "{}" { // empty struct
				obj.ServiceDefinitionV2Dot1Slack = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Dot1Slack = nil
		}
	} else {
		obj.ServiceDefinitionV2Dot1Slack = nil
	}

	// try to unmarshal data into ServiceDefinitionV2Dot1MSTeams
	err = datadog.Unmarshal(data, &obj.ServiceDefinitionV2Dot1MSTeams)
	if err == nil {
		if obj.ServiceDefinitionV2Dot1MSTeams != nil && obj.ServiceDefinitionV2Dot1MSTeams.UnparsedObject == nil {
			jsonServiceDefinitionV2Dot1MSTeams, _ := datadog.Marshal(obj.ServiceDefinitionV2Dot1MSTeams)
			if string(jsonServiceDefinitionV2Dot1MSTeams) == "{}" { // empty struct
				obj.ServiceDefinitionV2Dot1MSTeams = nil
			} else {
				match++
			}
		} else {
			obj.ServiceDefinitionV2Dot1MSTeams = nil
		}
	} else {
		obj.ServiceDefinitionV2Dot1MSTeams = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ServiceDefinitionV2Dot1Email = nil
		obj.ServiceDefinitionV2Dot1Slack = nil
		obj.ServiceDefinitionV2Dot1MSTeams = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ServiceDefinitionV2Dot1Contact) MarshalJSON() ([]byte, error) {
	if obj.ServiceDefinitionV2Dot1Email != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Dot1Email)
	}

	if obj.ServiceDefinitionV2Dot1Slack != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Dot1Slack)
	}

	if obj.ServiceDefinitionV2Dot1MSTeams != nil {
		return datadog.Marshal(&obj.ServiceDefinitionV2Dot1MSTeams)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ServiceDefinitionV2Dot1Contact) GetActualInstance() interface{} {
	if obj.ServiceDefinitionV2Dot1Email != nil {
		return obj.ServiceDefinitionV2Dot1Email
	}

	if obj.ServiceDefinitionV2Dot1Slack != nil {
		return obj.ServiceDefinitionV2Dot1Slack
	}

	if obj.ServiceDefinitionV2Dot1MSTeams != nil {
		return obj.ServiceDefinitionV2Dot1MSTeams
	}

	// all schemas are nil
	return nil
}
