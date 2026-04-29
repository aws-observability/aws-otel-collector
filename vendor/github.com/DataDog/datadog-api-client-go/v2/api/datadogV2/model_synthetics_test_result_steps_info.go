// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultStepsInfo Step execution summary for a Synthetic test result.
type SyntheticsTestResultStepsInfo struct {
	// Number of completed steps.
	Completed *int64 `json:"completed,omitempty"`
	// Number of steps with errors.
	Errors *int64 `json:"errors,omitempty"`
	// Total number of steps.
	Total *int64 `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultStepsInfo instantiates a new SyntheticsTestResultStepsInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultStepsInfo() *SyntheticsTestResultStepsInfo {
	this := SyntheticsTestResultStepsInfo{}
	return &this
}

// NewSyntheticsTestResultStepsInfoWithDefaults instantiates a new SyntheticsTestResultStepsInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultStepsInfoWithDefaults() *SyntheticsTestResultStepsInfo {
	this := SyntheticsTestResultStepsInfo{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepsInfo) GetCompleted() int64 {
	if o == nil || o.Completed == nil {
		var ret int64
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepsInfo) GetCompletedOk() (*int64, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepsInfo) HasCompleted() bool {
	return o != nil && o.Completed != nil
}

// SetCompleted gets a reference to the given int64 and assigns it to the Completed field.
func (o *SyntheticsTestResultStepsInfo) SetCompleted(v int64) {
	o.Completed = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepsInfo) GetErrors() int64 {
	if o == nil || o.Errors == nil {
		var ret int64
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepsInfo) GetErrorsOk() (*int64, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepsInfo) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given int64 and assigns it to the Errors field.
func (o *SyntheticsTestResultStepsInfo) SetErrors(v int64) {
	o.Errors = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SyntheticsTestResultStepsInfo) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultStepsInfo) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SyntheticsTestResultStepsInfo) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *SyntheticsTestResultStepsInfo) SetTotal(v int64) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultStepsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultStepsInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Completed *int64 `json:"completed,omitempty"`
		Errors    *int64 `json:"errors,omitempty"`
		Total     *int64 `json:"total,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"completed", "errors", "total"})
	} else {
		return err
	}
	o.Completed = all.Completed
	o.Errors = all.Errors
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
