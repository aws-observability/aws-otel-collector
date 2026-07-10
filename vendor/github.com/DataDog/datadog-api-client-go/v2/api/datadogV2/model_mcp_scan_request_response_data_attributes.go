// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// McpScanRequestResponseDataAttributes The attributes returned when a scan request has been accepted, containing the job identifier used to poll for results.
type McpScanRequestResponseDataAttributes struct {
	// The job identifier assigned to the scan, used to retrieve the scan result.
	JobId string `json:"job_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMcpScanRequestResponseDataAttributes instantiates a new McpScanRequestResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMcpScanRequestResponseDataAttributes(jobId string) *McpScanRequestResponseDataAttributes {
	this := McpScanRequestResponseDataAttributes{}
	this.JobId = jobId
	return &this
}

// NewMcpScanRequestResponseDataAttributesWithDefaults instantiates a new McpScanRequestResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMcpScanRequestResponseDataAttributesWithDefaults() *McpScanRequestResponseDataAttributes {
	this := McpScanRequestResponseDataAttributes{}
	return &this
}

// GetJobId returns the JobId field value.
func (o *McpScanRequestResponseDataAttributes) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestResponseDataAttributes) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value.
func (o *McpScanRequestResponseDataAttributes) SetJobId(v string) {
	o.JobId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o McpScanRequestResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["job_id"] = o.JobId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *McpScanRequestResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JobId *string `json:"job_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JobId == nil {
		return fmt.Errorf("required field job_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"job_id"})
	} else {
		return err
	}
	o.JobId = *all.JobId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
