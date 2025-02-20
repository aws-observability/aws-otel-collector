// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowListInstancesResponseMetaPage Page information for the list instances response.
type WorkflowListInstancesResponseMetaPage struct {
	// The total count of items.
	TotalCount *int64 `json:"totalCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowListInstancesResponseMetaPage instantiates a new WorkflowListInstancesResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowListInstancesResponseMetaPage() *WorkflowListInstancesResponseMetaPage {
	this := WorkflowListInstancesResponseMetaPage{}
	return &this
}

// NewWorkflowListInstancesResponseMetaPageWithDefaults instantiates a new WorkflowListInstancesResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowListInstancesResponseMetaPageWithDefaults() *WorkflowListInstancesResponseMetaPage {
	this := WorkflowListInstancesResponseMetaPage{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *WorkflowListInstancesResponseMetaPage) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListInstancesResponseMetaPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *WorkflowListInstancesResponseMetaPage) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *WorkflowListInstancesResponseMetaPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowListInstancesResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowListInstancesResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount *int64 `json:"totalCount,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"totalCount"})
	} else {
		return err
	}
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
