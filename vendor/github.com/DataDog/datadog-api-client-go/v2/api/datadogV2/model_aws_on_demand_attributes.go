// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsOnDemandAttributes Attributes for the AWS on demand task.
type AwsOnDemandAttributes struct {
	// The arn of the resource to scan.
	Arn *string `json:"arn,omitempty"`
	// Specifies the assignment timestamp if the task has been already assigned to a scanner.
	AssignedAt *string `json:"assigned_at,omitempty"`
	// The task submission timestamp.
	CreatedAt *string `json:"created_at,omitempty"`
	// Indicates the status of the task.
	// QUEUED: the task has been submitted successfully and the resource has not been assigned to a scanner yet.
	// ASSIGNED: the task has been assigned.
	// ABORTED: the scan has been aborted after a period of time due to technical reasons, such as resource not found, insufficient permissions, or the absence of a configured scanner.
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsOnDemandAttributes instantiates a new AwsOnDemandAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsOnDemandAttributes() *AwsOnDemandAttributes {
	this := AwsOnDemandAttributes{}
	return &this
}

// NewAwsOnDemandAttributesWithDefaults instantiates a new AwsOnDemandAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsOnDemandAttributesWithDefaults() *AwsOnDemandAttributes {
	this := AwsOnDemandAttributes{}
	return &this
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *AwsOnDemandAttributes) GetArn() string {
	if o == nil || o.Arn == nil {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsOnDemandAttributes) GetArnOk() (*string, bool) {
	if o == nil || o.Arn == nil {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *AwsOnDemandAttributes) HasArn() bool {
	return o != nil && o.Arn != nil
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *AwsOnDemandAttributes) SetArn(v string) {
	o.Arn = &v
}

// GetAssignedAt returns the AssignedAt field value if set, zero value otherwise.
func (o *AwsOnDemandAttributes) GetAssignedAt() string {
	if o == nil || o.AssignedAt == nil {
		var ret string
		return ret
	}
	return *o.AssignedAt
}

// GetAssignedAtOk returns a tuple with the AssignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsOnDemandAttributes) GetAssignedAtOk() (*string, bool) {
	if o == nil || o.AssignedAt == nil {
		return nil, false
	}
	return o.AssignedAt, true
}

// HasAssignedAt returns a boolean if a field has been set.
func (o *AwsOnDemandAttributes) HasAssignedAt() bool {
	return o != nil && o.AssignedAt != nil
}

// SetAssignedAt gets a reference to the given string and assigns it to the AssignedAt field.
func (o *AwsOnDemandAttributes) SetAssignedAt(v string) {
	o.AssignedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AwsOnDemandAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsOnDemandAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AwsOnDemandAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AwsOnDemandAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AwsOnDemandAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsOnDemandAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsOnDemandAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AwsOnDemandAttributes) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsOnDemandAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arn != nil {
		toSerialize["arn"] = o.Arn
	}
	if o.AssignedAt != nil {
		toSerialize["assigned_at"] = o.AssignedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsOnDemandAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arn        *string `json:"arn,omitempty"`
		AssignedAt *string `json:"assigned_at,omitempty"`
		CreatedAt  *string `json:"created_at,omitempty"`
		Status     *string `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arn", "assigned_at", "created_at", "status"})
	} else {
		return err
	}
	o.Arn = all.Arn
	o.AssignedAt = all.AssignedAt
	o.CreatedAt = all.CreatedAt
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
