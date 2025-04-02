// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalJobResponseAttributes Historical job attributes.
type HistoricalJobResponseAttributes struct {
	// Time when the job was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The handle of the user who created the job.
	CreatedByHandle *string `json:"createdByHandle,omitempty"`
	// The name of the user who created the job.
	CreatedByName *string `json:"createdByName,omitempty"`
	// ID of the rule used to create the job (if it is created from a rule).
	CreatedFromRuleId *string `json:"createdFromRuleId,omitempty"`
	// Definition of a historical job.
	JobDefinition *JobDefinition `json:"jobDefinition,omitempty"`
	// Job name.
	JobName *string `json:"jobName,omitempty"`
	// Job status.
	JobStatus *string `json:"jobStatus,omitempty"`
	// Last modification time of the job.
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHistoricalJobResponseAttributes instantiates a new HistoricalJobResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHistoricalJobResponseAttributes() *HistoricalJobResponseAttributes {
	this := HistoricalJobResponseAttributes{}
	return &this
}

// NewHistoricalJobResponseAttributesWithDefaults instantiates a new HistoricalJobResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHistoricalJobResponseAttributesWithDefaults() *HistoricalJobResponseAttributes {
	this := HistoricalJobResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *HistoricalJobResponseAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *HistoricalJobResponseAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetCreatedByName() string {
	if o == nil || o.CreatedByName == nil {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetCreatedByNameOk() (*string, bool) {
	if o == nil || o.CreatedByName == nil {
		return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasCreatedByName() bool {
	return o != nil && o.CreatedByName != nil
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *HistoricalJobResponseAttributes) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetCreatedFromRuleId returns the CreatedFromRuleId field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetCreatedFromRuleId() string {
	if o == nil || o.CreatedFromRuleId == nil {
		var ret string
		return ret
	}
	return *o.CreatedFromRuleId
}

// GetCreatedFromRuleIdOk returns a tuple with the CreatedFromRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetCreatedFromRuleIdOk() (*string, bool) {
	if o == nil || o.CreatedFromRuleId == nil {
		return nil, false
	}
	return o.CreatedFromRuleId, true
}

// HasCreatedFromRuleId returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasCreatedFromRuleId() bool {
	return o != nil && o.CreatedFromRuleId != nil
}

// SetCreatedFromRuleId gets a reference to the given string and assigns it to the CreatedFromRuleId field.
func (o *HistoricalJobResponseAttributes) SetCreatedFromRuleId(v string) {
	o.CreatedFromRuleId = &v
}

// GetJobDefinition returns the JobDefinition field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetJobDefinition() JobDefinition {
	if o == nil || o.JobDefinition == nil {
		var ret JobDefinition
		return ret
	}
	return *o.JobDefinition
}

// GetJobDefinitionOk returns a tuple with the JobDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetJobDefinitionOk() (*JobDefinition, bool) {
	if o == nil || o.JobDefinition == nil {
		return nil, false
	}
	return o.JobDefinition, true
}

// HasJobDefinition returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasJobDefinition() bool {
	return o != nil && o.JobDefinition != nil
}

// SetJobDefinition gets a reference to the given JobDefinition and assigns it to the JobDefinition field.
func (o *HistoricalJobResponseAttributes) SetJobDefinition(v JobDefinition) {
	o.JobDefinition = &v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetJobName() string {
	if o == nil || o.JobName == nil {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetJobNameOk() (*string, bool) {
	if o == nil || o.JobName == nil {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasJobName() bool {
	return o != nil && o.JobName != nil
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *HistoricalJobResponseAttributes) SetJobName(v string) {
	o.JobName = &v
}

// GetJobStatus returns the JobStatus field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetJobStatus() string {
	if o == nil || o.JobStatus == nil {
		var ret string
		return ret
	}
	return *o.JobStatus
}

// GetJobStatusOk returns a tuple with the JobStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetJobStatusOk() (*string, bool) {
	if o == nil || o.JobStatus == nil {
		return nil, false
	}
	return o.JobStatus, true
}

// HasJobStatus returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasJobStatus() bool {
	return o != nil && o.JobStatus != nil
}

// SetJobStatus gets a reference to the given string and assigns it to the JobStatus field.
func (o *HistoricalJobResponseAttributes) SetJobStatus(v string) {
	o.JobStatus = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *HistoricalJobResponseAttributes) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobResponseAttributes) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *HistoricalJobResponseAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *HistoricalJobResponseAttributes) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HistoricalJobResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedByHandle != nil {
		toSerialize["createdByHandle"] = o.CreatedByHandle
	}
	if o.CreatedByName != nil {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if o.CreatedFromRuleId != nil {
		toSerialize["createdFromRuleId"] = o.CreatedFromRuleId
	}
	if o.JobDefinition != nil {
		toSerialize["jobDefinition"] = o.JobDefinition
	}
	if o.JobName != nil {
		toSerialize["jobName"] = o.JobName
	}
	if o.JobStatus != nil {
		toSerialize["jobStatus"] = o.JobStatus
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HistoricalJobResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *string        `json:"createdAt,omitempty"`
		CreatedByHandle   *string        `json:"createdByHandle,omitempty"`
		CreatedByName     *string        `json:"createdByName,omitempty"`
		CreatedFromRuleId *string        `json:"createdFromRuleId,omitempty"`
		JobDefinition     *JobDefinition `json:"jobDefinition,omitempty"`
		JobName           *string        `json:"jobName,omitempty"`
		JobStatus         *string        `json:"jobStatus,omitempty"`
		ModifiedAt        *string        `json:"modifiedAt,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "createdByHandle", "createdByName", "createdFromRuleId", "jobDefinition", "jobName", "jobStatus", "modifiedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedByHandle = all.CreatedByHandle
	o.CreatedByName = all.CreatedByName
	o.CreatedFromRuleId = all.CreatedFromRuleId
	if all.JobDefinition != nil && all.JobDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JobDefinition = all.JobDefinition
	o.JobName = all.JobName
	o.JobStatus = all.JobStatus
	o.ModifiedAt = all.ModifiedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
