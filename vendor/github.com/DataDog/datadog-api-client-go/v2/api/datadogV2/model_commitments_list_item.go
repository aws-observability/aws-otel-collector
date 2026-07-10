// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsListItem - A commitment item, which varies based on the provider, product, and commitment type.
type CommitmentsListItem struct {
	CommitmentsAwsEC2RICommitment         *CommitmentsAwsEC2RICommitment
	CommitmentsAwsRDSRICommitment         *CommitmentsAwsRDSRICommitment
	CommitmentsAwsElasticacheRICommitment *CommitmentsAwsElasticacheRICommitment
	CommitmentsAwsSPCommitment            *CommitmentsAwsSPCommitment
	CommitmentsAzureVMRICommitment        *CommitmentsAzureVMRICommitment
	CommitmentsAzureComputeSPCommitment   *CommitmentsAzureComputeSPCommitment

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CommitmentsAwsEC2RICommitmentAsCommitmentsListItem is a convenience function that returns CommitmentsAwsEC2RICommitment wrapped in CommitmentsListItem.
func CommitmentsAwsEC2RICommitmentAsCommitmentsListItem(v *CommitmentsAwsEC2RICommitment) CommitmentsListItem {
	return CommitmentsListItem{CommitmentsAwsEC2RICommitment: v}
}

// CommitmentsAwsRDSRICommitmentAsCommitmentsListItem is a convenience function that returns CommitmentsAwsRDSRICommitment wrapped in CommitmentsListItem.
func CommitmentsAwsRDSRICommitmentAsCommitmentsListItem(v *CommitmentsAwsRDSRICommitment) CommitmentsListItem {
	return CommitmentsListItem{CommitmentsAwsRDSRICommitment: v}
}

// CommitmentsAwsElasticacheRICommitmentAsCommitmentsListItem is a convenience function that returns CommitmentsAwsElasticacheRICommitment wrapped in CommitmentsListItem.
func CommitmentsAwsElasticacheRICommitmentAsCommitmentsListItem(v *CommitmentsAwsElasticacheRICommitment) CommitmentsListItem {
	return CommitmentsListItem{CommitmentsAwsElasticacheRICommitment: v}
}

// CommitmentsAwsSPCommitmentAsCommitmentsListItem is a convenience function that returns CommitmentsAwsSPCommitment wrapped in CommitmentsListItem.
func CommitmentsAwsSPCommitmentAsCommitmentsListItem(v *CommitmentsAwsSPCommitment) CommitmentsListItem {
	return CommitmentsListItem{CommitmentsAwsSPCommitment: v}
}

// CommitmentsAzureVMRICommitmentAsCommitmentsListItem is a convenience function that returns CommitmentsAzureVMRICommitment wrapped in CommitmentsListItem.
func CommitmentsAzureVMRICommitmentAsCommitmentsListItem(v *CommitmentsAzureVMRICommitment) CommitmentsListItem {
	return CommitmentsListItem{CommitmentsAzureVMRICommitment: v}
}

// CommitmentsAzureComputeSPCommitmentAsCommitmentsListItem is a convenience function that returns CommitmentsAzureComputeSPCommitment wrapped in CommitmentsListItem.
func CommitmentsAzureComputeSPCommitmentAsCommitmentsListItem(v *CommitmentsAzureComputeSPCommitment) CommitmentsListItem {
	return CommitmentsListItem{CommitmentsAzureComputeSPCommitment: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CommitmentsListItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CommitmentsAwsEC2RICommitment
	err = datadog.Unmarshal(data, &obj.CommitmentsAwsEC2RICommitment)
	if err == nil {
		if obj.CommitmentsAwsEC2RICommitment != nil && obj.CommitmentsAwsEC2RICommitment.UnparsedObject == nil {
			jsonCommitmentsAwsEC2RICommitment, _ := datadog.Marshal(obj.CommitmentsAwsEC2RICommitment)
			if string(jsonCommitmentsAwsEC2RICommitment) == "{}" { // empty struct
				obj.CommitmentsAwsEC2RICommitment = nil
			} else {
				match++
			}
		} else {
			obj.CommitmentsAwsEC2RICommitment = nil
		}
	} else {
		obj.CommitmentsAwsEC2RICommitment = nil
	}

	// try to unmarshal data into CommitmentsAwsRDSRICommitment
	err = datadog.Unmarshal(data, &obj.CommitmentsAwsRDSRICommitment)
	if err == nil {
		if obj.CommitmentsAwsRDSRICommitment != nil && obj.CommitmentsAwsRDSRICommitment.UnparsedObject == nil {
			jsonCommitmentsAwsRDSRICommitment, _ := datadog.Marshal(obj.CommitmentsAwsRDSRICommitment)
			if string(jsonCommitmentsAwsRDSRICommitment) == "{}" { // empty struct
				obj.CommitmentsAwsRDSRICommitment = nil
			} else {
				match++
			}
		} else {
			obj.CommitmentsAwsRDSRICommitment = nil
		}
	} else {
		obj.CommitmentsAwsRDSRICommitment = nil
	}

	// try to unmarshal data into CommitmentsAwsElasticacheRICommitment
	err = datadog.Unmarshal(data, &obj.CommitmentsAwsElasticacheRICommitment)
	if err == nil {
		if obj.CommitmentsAwsElasticacheRICommitment != nil && obj.CommitmentsAwsElasticacheRICommitment.UnparsedObject == nil {
			jsonCommitmentsAwsElasticacheRICommitment, _ := datadog.Marshal(obj.CommitmentsAwsElasticacheRICommitment)
			if string(jsonCommitmentsAwsElasticacheRICommitment) == "{}" { // empty struct
				obj.CommitmentsAwsElasticacheRICommitment = nil
			} else {
				match++
			}
		} else {
			obj.CommitmentsAwsElasticacheRICommitment = nil
		}
	} else {
		obj.CommitmentsAwsElasticacheRICommitment = nil
	}

	// try to unmarshal data into CommitmentsAwsSPCommitment
	err = datadog.Unmarshal(data, &obj.CommitmentsAwsSPCommitment)
	if err == nil {
		if obj.CommitmentsAwsSPCommitment != nil && obj.CommitmentsAwsSPCommitment.UnparsedObject == nil {
			jsonCommitmentsAwsSPCommitment, _ := datadog.Marshal(obj.CommitmentsAwsSPCommitment)
			if string(jsonCommitmentsAwsSPCommitment) == "{}" { // empty struct
				obj.CommitmentsAwsSPCommitment = nil
			} else {
				match++
			}
		} else {
			obj.CommitmentsAwsSPCommitment = nil
		}
	} else {
		obj.CommitmentsAwsSPCommitment = nil
	}

	// try to unmarshal data into CommitmentsAzureVMRICommitment
	err = datadog.Unmarshal(data, &obj.CommitmentsAzureVMRICommitment)
	if err == nil {
		if obj.CommitmentsAzureVMRICommitment != nil && obj.CommitmentsAzureVMRICommitment.UnparsedObject == nil {
			jsonCommitmentsAzureVMRICommitment, _ := datadog.Marshal(obj.CommitmentsAzureVMRICommitment)
			if string(jsonCommitmentsAzureVMRICommitment) == "{}" { // empty struct
				obj.CommitmentsAzureVMRICommitment = nil
			} else {
				match++
			}
		} else {
			obj.CommitmentsAzureVMRICommitment = nil
		}
	} else {
		obj.CommitmentsAzureVMRICommitment = nil
	}

	// try to unmarshal data into CommitmentsAzureComputeSPCommitment
	err = datadog.Unmarshal(data, &obj.CommitmentsAzureComputeSPCommitment)
	if err == nil {
		if obj.CommitmentsAzureComputeSPCommitment != nil && obj.CommitmentsAzureComputeSPCommitment.UnparsedObject == nil {
			jsonCommitmentsAzureComputeSPCommitment, _ := datadog.Marshal(obj.CommitmentsAzureComputeSPCommitment)
			if string(jsonCommitmentsAzureComputeSPCommitment) == "{}" { // empty struct
				obj.CommitmentsAzureComputeSPCommitment = nil
			} else {
				match++
			}
		} else {
			obj.CommitmentsAzureComputeSPCommitment = nil
		}
	} else {
		obj.CommitmentsAzureComputeSPCommitment = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CommitmentsAwsEC2RICommitment = nil
		obj.CommitmentsAwsRDSRICommitment = nil
		obj.CommitmentsAwsElasticacheRICommitment = nil
		obj.CommitmentsAwsSPCommitment = nil
		obj.CommitmentsAzureVMRICommitment = nil
		obj.CommitmentsAzureComputeSPCommitment = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CommitmentsListItem) MarshalJSON() ([]byte, error) {
	if obj.CommitmentsAwsEC2RICommitment != nil {
		return datadog.Marshal(&obj.CommitmentsAwsEC2RICommitment)
	}

	if obj.CommitmentsAwsRDSRICommitment != nil {
		return datadog.Marshal(&obj.CommitmentsAwsRDSRICommitment)
	}

	if obj.CommitmentsAwsElasticacheRICommitment != nil {
		return datadog.Marshal(&obj.CommitmentsAwsElasticacheRICommitment)
	}

	if obj.CommitmentsAwsSPCommitment != nil {
		return datadog.Marshal(&obj.CommitmentsAwsSPCommitment)
	}

	if obj.CommitmentsAzureVMRICommitment != nil {
		return datadog.Marshal(&obj.CommitmentsAzureVMRICommitment)
	}

	if obj.CommitmentsAzureComputeSPCommitment != nil {
		return datadog.Marshal(&obj.CommitmentsAzureComputeSPCommitment)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CommitmentsListItem) GetActualInstance() interface{} {
	if obj.CommitmentsAwsEC2RICommitment != nil {
		return obj.CommitmentsAwsEC2RICommitment
	}

	if obj.CommitmentsAwsRDSRICommitment != nil {
		return obj.CommitmentsAwsRDSRICommitment
	}

	if obj.CommitmentsAwsElasticacheRICommitment != nil {
		return obj.CommitmentsAwsElasticacheRICommitment
	}

	if obj.CommitmentsAwsSPCommitment != nil {
		return obj.CommitmentsAwsSPCommitment
	}

	if obj.CommitmentsAzureVMRICommitment != nil {
		return obj.CommitmentsAzureVMRICommitment
	}

	if obj.CommitmentsAzureComputeSPCommitment != nil {
		return obj.CommitmentsAzureComputeSPCommitment
	}

	// all schemas are nil
	return nil
}
