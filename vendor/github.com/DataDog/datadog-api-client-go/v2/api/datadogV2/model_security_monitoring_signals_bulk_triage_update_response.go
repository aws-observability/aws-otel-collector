// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalsBulkTriageUpdateResponse Response for a bulk triage update of security signals.
type SecurityMonitoringSignalsBulkTriageUpdateResponse struct {
	// The result payload of a bulk signal triage update.
	Result SecurityMonitoringSignalsBulkTriageUpdateResult `json:"result"`
	// The status of the bulk operation.
	Status string `json:"status"`
	// The type of the response.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalsBulkTriageUpdateResponse instantiates a new SecurityMonitoringSignalsBulkTriageUpdateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalsBulkTriageUpdateResponse(result SecurityMonitoringSignalsBulkTriageUpdateResult, status string, typeVar string) *SecurityMonitoringSignalsBulkTriageUpdateResponse {
	this := SecurityMonitoringSignalsBulkTriageUpdateResponse{}
	this.Result = result
	this.Status = status
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringSignalsBulkTriageUpdateResponseWithDefaults instantiates a new SecurityMonitoringSignalsBulkTriageUpdateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalsBulkTriageUpdateResponseWithDefaults() *SecurityMonitoringSignalsBulkTriageUpdateResponse {
	this := SecurityMonitoringSignalsBulkTriageUpdateResponse{}
	return &this
}

// GetResult returns the Result field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) GetResult() SecurityMonitoringSignalsBulkTriageUpdateResult {
	if o == nil {
		var ret SecurityMonitoringSignalsBulkTriageUpdateResult
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) GetResultOk() (*SecurityMonitoringSignalsBulkTriageUpdateResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) SetResult(v SecurityMonitoringSignalsBulkTriageUpdateResult) {
	o.Result = v
}

// GetStatus returns the Status field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalsBulkTriageUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["result"] = o.Result
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalsBulkTriageUpdateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Result *SecurityMonitoringSignalsBulkTriageUpdateResult `json:"result"`
		Status *string                                          `json:"status"`
		Type   *string                                          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Result == nil {
		return fmt.Errorf("required field result missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"result", "status", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = *all.Result
	o.Status = *all.Status
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
