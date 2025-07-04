// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutboundEdge The definition of `OutboundEdge` object.
type OutboundEdge struct {
	// The `OutboundEdge` `branchName`.
	BranchName string `json:"branchName"`
	// The `OutboundEdge` `nextStepName`.
	NextStepName string `json:"nextStepName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutboundEdge instantiates a new OutboundEdge object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutboundEdge(branchName string, nextStepName string) *OutboundEdge {
	this := OutboundEdge{}
	this.BranchName = branchName
	this.NextStepName = nextStepName
	return &this
}

// NewOutboundEdgeWithDefaults instantiates a new OutboundEdge object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutboundEdgeWithDefaults() *OutboundEdge {
	this := OutboundEdge{}
	return &this
}

// GetBranchName returns the BranchName field value.
func (o *OutboundEdge) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *OutboundEdge) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value.
func (o *OutboundEdge) SetBranchName(v string) {
	o.BranchName = v
}

// GetNextStepName returns the NextStepName field value.
func (o *OutboundEdge) GetNextStepName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NextStepName
}

// GetNextStepNameOk returns a tuple with the NextStepName field value
// and a boolean to check if the value has been set.
func (o *OutboundEdge) GetNextStepNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextStepName, true
}

// SetNextStepName sets field value.
func (o *OutboundEdge) SetNextStepName(v string) {
	o.NextStepName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutboundEdge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["branchName"] = o.BranchName
	toSerialize["nextStepName"] = o.NextStepName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutboundEdge) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BranchName   *string `json:"branchName"`
		NextStepName *string `json:"nextStepName"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BranchName == nil {
		return fmt.Errorf("required field branchName missing")
	}
	if all.NextStepName == nil {
		return fmt.Errorf("required field nextStepName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branchName", "nextStepName"})
	} else {
		return err
	}
	o.BranchName = *all.BranchName
	o.NextStepName = *all.NextStepName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
