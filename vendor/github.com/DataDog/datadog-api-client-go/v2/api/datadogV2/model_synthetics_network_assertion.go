// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertion - Object describing an assertion for a Network Path test.
type SyntheticsNetworkAssertion struct {
	SyntheticsNetworkAssertionLatency              *SyntheticsNetworkAssertionLatency
	SyntheticsNetworkAssertionMultiNetworkHop      *SyntheticsNetworkAssertionMultiNetworkHop
	SyntheticsNetworkAssertionPacketLossPercentage *SyntheticsNetworkAssertionPacketLossPercentage
	SyntheticsNetworkAssertionJitter               *SyntheticsNetworkAssertionJitter

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsNetworkAssertionLatencyAsSyntheticsNetworkAssertion is a convenience function that returns SyntheticsNetworkAssertionLatency wrapped in SyntheticsNetworkAssertion.
func SyntheticsNetworkAssertionLatencyAsSyntheticsNetworkAssertion(v *SyntheticsNetworkAssertionLatency) SyntheticsNetworkAssertion {
	return SyntheticsNetworkAssertion{SyntheticsNetworkAssertionLatency: v}
}

// SyntheticsNetworkAssertionMultiNetworkHopAsSyntheticsNetworkAssertion is a convenience function that returns SyntheticsNetworkAssertionMultiNetworkHop wrapped in SyntheticsNetworkAssertion.
func SyntheticsNetworkAssertionMultiNetworkHopAsSyntheticsNetworkAssertion(v *SyntheticsNetworkAssertionMultiNetworkHop) SyntheticsNetworkAssertion {
	return SyntheticsNetworkAssertion{SyntheticsNetworkAssertionMultiNetworkHop: v}
}

// SyntheticsNetworkAssertionPacketLossPercentageAsSyntheticsNetworkAssertion is a convenience function that returns SyntheticsNetworkAssertionPacketLossPercentage wrapped in SyntheticsNetworkAssertion.
func SyntheticsNetworkAssertionPacketLossPercentageAsSyntheticsNetworkAssertion(v *SyntheticsNetworkAssertionPacketLossPercentage) SyntheticsNetworkAssertion {
	return SyntheticsNetworkAssertion{SyntheticsNetworkAssertionPacketLossPercentage: v}
}

// SyntheticsNetworkAssertionJitterAsSyntheticsNetworkAssertion is a convenience function that returns SyntheticsNetworkAssertionJitter wrapped in SyntheticsNetworkAssertion.
func SyntheticsNetworkAssertionJitterAsSyntheticsNetworkAssertion(v *SyntheticsNetworkAssertionJitter) SyntheticsNetworkAssertion {
	return SyntheticsNetworkAssertion{SyntheticsNetworkAssertionJitter: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsNetworkAssertion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsNetworkAssertionLatency
	err = datadog.Unmarshal(data, &obj.SyntheticsNetworkAssertionLatency)
	if err == nil {
		if obj.SyntheticsNetworkAssertionLatency != nil && obj.SyntheticsNetworkAssertionLatency.UnparsedObject == nil {
			jsonSyntheticsNetworkAssertionLatency, _ := datadog.Marshal(obj.SyntheticsNetworkAssertionLatency)
			if string(jsonSyntheticsNetworkAssertionLatency) == "{}" { // empty struct
				obj.SyntheticsNetworkAssertionLatency = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsNetworkAssertionLatency = nil
		}
	} else {
		obj.SyntheticsNetworkAssertionLatency = nil
	}

	// try to unmarshal data into SyntheticsNetworkAssertionMultiNetworkHop
	err = datadog.Unmarshal(data, &obj.SyntheticsNetworkAssertionMultiNetworkHop)
	if err == nil {
		if obj.SyntheticsNetworkAssertionMultiNetworkHop != nil && obj.SyntheticsNetworkAssertionMultiNetworkHop.UnparsedObject == nil {
			jsonSyntheticsNetworkAssertionMultiNetworkHop, _ := datadog.Marshal(obj.SyntheticsNetworkAssertionMultiNetworkHop)
			if string(jsonSyntheticsNetworkAssertionMultiNetworkHop) == "{}" { // empty struct
				obj.SyntheticsNetworkAssertionMultiNetworkHop = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsNetworkAssertionMultiNetworkHop = nil
		}
	} else {
		obj.SyntheticsNetworkAssertionMultiNetworkHop = nil
	}

	// try to unmarshal data into SyntheticsNetworkAssertionPacketLossPercentage
	err = datadog.Unmarshal(data, &obj.SyntheticsNetworkAssertionPacketLossPercentage)
	if err == nil {
		if obj.SyntheticsNetworkAssertionPacketLossPercentage != nil && obj.SyntheticsNetworkAssertionPacketLossPercentage.UnparsedObject == nil {
			jsonSyntheticsNetworkAssertionPacketLossPercentage, _ := datadog.Marshal(obj.SyntheticsNetworkAssertionPacketLossPercentage)
			if string(jsonSyntheticsNetworkAssertionPacketLossPercentage) == "{}" { // empty struct
				obj.SyntheticsNetworkAssertionPacketLossPercentage = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsNetworkAssertionPacketLossPercentage = nil
		}
	} else {
		obj.SyntheticsNetworkAssertionPacketLossPercentage = nil
	}

	// try to unmarshal data into SyntheticsNetworkAssertionJitter
	err = datadog.Unmarshal(data, &obj.SyntheticsNetworkAssertionJitter)
	if err == nil {
		if obj.SyntheticsNetworkAssertionJitter != nil && obj.SyntheticsNetworkAssertionJitter.UnparsedObject == nil {
			jsonSyntheticsNetworkAssertionJitter, _ := datadog.Marshal(obj.SyntheticsNetworkAssertionJitter)
			if string(jsonSyntheticsNetworkAssertionJitter) == "{}" { // empty struct
				obj.SyntheticsNetworkAssertionJitter = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsNetworkAssertionJitter = nil
		}
	} else {
		obj.SyntheticsNetworkAssertionJitter = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsNetworkAssertionLatency = nil
		obj.SyntheticsNetworkAssertionMultiNetworkHop = nil
		obj.SyntheticsNetworkAssertionPacketLossPercentage = nil
		obj.SyntheticsNetworkAssertionJitter = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsNetworkAssertion) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsNetworkAssertionLatency != nil {
		return datadog.Marshal(&obj.SyntheticsNetworkAssertionLatency)
	}

	if obj.SyntheticsNetworkAssertionMultiNetworkHop != nil {
		return datadog.Marshal(&obj.SyntheticsNetworkAssertionMultiNetworkHop)
	}

	if obj.SyntheticsNetworkAssertionPacketLossPercentage != nil {
		return datadog.Marshal(&obj.SyntheticsNetworkAssertionPacketLossPercentage)
	}

	if obj.SyntheticsNetworkAssertionJitter != nil {
		return datadog.Marshal(&obj.SyntheticsNetworkAssertionJitter)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsNetworkAssertion) GetActualInstance() interface{} {
	if obj.SyntheticsNetworkAssertionLatency != nil {
		return obj.SyntheticsNetworkAssertionLatency
	}

	if obj.SyntheticsNetworkAssertionMultiNetworkHop != nil {
		return obj.SyntheticsNetworkAssertionMultiNetworkHop
	}

	if obj.SyntheticsNetworkAssertionPacketLossPercentage != nil {
		return obj.SyntheticsNetworkAssertionPacketLossPercentage
	}

	if obj.SyntheticsNetworkAssertionJitter != nil {
		return obj.SyntheticsNetworkAssertionJitter
	}

	// all schemas are nil
	return nil
}
