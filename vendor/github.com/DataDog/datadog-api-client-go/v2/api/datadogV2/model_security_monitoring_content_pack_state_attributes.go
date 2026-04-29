// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackStateAttributes Attributes of a content pack state
type SecurityMonitoringContentPackStateAttributes struct {
	// Whether the cloud SIEM index configuration is incorrect (only applies to certain pricing models)
	CloudSiemIndexIncorrect bool `json:"cloud_siem_index_incorrect"`
	// The activation status of a content pack.
	CpActivation SecurityMonitoringContentPackActivation `json:"cp_activation"`
	// Whether filters (Security Filters or Index Query depending on the pricing model) are
	// present and correctly configured to route logs into Cloud SIEM.
	FiltersConfiguredForLogs bool `json:"filters_configured_for_logs"`
	// The installation status of the related integration.
	IntegrationInstalledStatus *SecurityMonitoringContentPackIntegrationStatus `json:"integration_installed_status,omitempty"`
	// Timestamp bucket indicating when logs were last collected.
	LogsLastCollected SecurityMonitoringContentPackTimestampBucket `json:"logs_last_collected"`
	// Whether logs for this content pack have been seen in any Datadog index within the last 72 hours.
	LogsSeenFromAnyIndex bool `json:"logs_seen_from_any_index"`
	// The current operational status of a content pack.
	State SecurityMonitoringContentPackStatus `json:"state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringContentPackStateAttributes instantiates a new SecurityMonitoringContentPackStateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringContentPackStateAttributes(cloudSiemIndexIncorrect bool, cpActivation SecurityMonitoringContentPackActivation, filtersConfiguredForLogs bool, logsLastCollected SecurityMonitoringContentPackTimestampBucket, logsSeenFromAnyIndex bool, state SecurityMonitoringContentPackStatus) *SecurityMonitoringContentPackStateAttributes {
	this := SecurityMonitoringContentPackStateAttributes{}
	this.CloudSiemIndexIncorrect = cloudSiemIndexIncorrect
	this.CpActivation = cpActivation
	this.FiltersConfiguredForLogs = filtersConfiguredForLogs
	this.LogsLastCollected = logsLastCollected
	this.LogsSeenFromAnyIndex = logsSeenFromAnyIndex
	this.State = state
	return &this
}

// NewSecurityMonitoringContentPackStateAttributesWithDefaults instantiates a new SecurityMonitoringContentPackStateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringContentPackStateAttributesWithDefaults() *SecurityMonitoringContentPackStateAttributes {
	this := SecurityMonitoringContentPackStateAttributes{}
	return &this
}

// GetCloudSiemIndexIncorrect returns the CloudSiemIndexIncorrect field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetCloudSiemIndexIncorrect() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CloudSiemIndexIncorrect
}

// GetCloudSiemIndexIncorrectOk returns a tuple with the CloudSiemIndexIncorrect field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetCloudSiemIndexIncorrectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudSiemIndexIncorrect, true
}

// SetCloudSiemIndexIncorrect sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetCloudSiemIndexIncorrect(v bool) {
	o.CloudSiemIndexIncorrect = v
}

// GetCpActivation returns the CpActivation field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetCpActivation() SecurityMonitoringContentPackActivation {
	if o == nil {
		var ret SecurityMonitoringContentPackActivation
		return ret
	}
	return o.CpActivation
}

// GetCpActivationOk returns a tuple with the CpActivation field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetCpActivationOk() (*SecurityMonitoringContentPackActivation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpActivation, true
}

// SetCpActivation sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetCpActivation(v SecurityMonitoringContentPackActivation) {
	o.CpActivation = v
}

// GetFiltersConfiguredForLogs returns the FiltersConfiguredForLogs field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetFiltersConfiguredForLogs() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.FiltersConfiguredForLogs
}

// GetFiltersConfiguredForLogsOk returns a tuple with the FiltersConfiguredForLogs field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetFiltersConfiguredForLogsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiltersConfiguredForLogs, true
}

// SetFiltersConfiguredForLogs sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetFiltersConfiguredForLogs(v bool) {
	o.FiltersConfiguredForLogs = v
}

// GetIntegrationInstalledStatus returns the IntegrationInstalledStatus field value if set, zero value otherwise.
func (o *SecurityMonitoringContentPackStateAttributes) GetIntegrationInstalledStatus() SecurityMonitoringContentPackIntegrationStatus {
	if o == nil || o.IntegrationInstalledStatus == nil {
		var ret SecurityMonitoringContentPackIntegrationStatus
		return ret
	}
	return *o.IntegrationInstalledStatus
}

// GetIntegrationInstalledStatusOk returns a tuple with the IntegrationInstalledStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetIntegrationInstalledStatusOk() (*SecurityMonitoringContentPackIntegrationStatus, bool) {
	if o == nil || o.IntegrationInstalledStatus == nil {
		return nil, false
	}
	return o.IntegrationInstalledStatus, true
}

// HasIntegrationInstalledStatus returns a boolean if a field has been set.
func (o *SecurityMonitoringContentPackStateAttributes) HasIntegrationInstalledStatus() bool {
	return o != nil && o.IntegrationInstalledStatus != nil
}

// SetIntegrationInstalledStatus gets a reference to the given SecurityMonitoringContentPackIntegrationStatus and assigns it to the IntegrationInstalledStatus field.
func (o *SecurityMonitoringContentPackStateAttributes) SetIntegrationInstalledStatus(v SecurityMonitoringContentPackIntegrationStatus) {
	o.IntegrationInstalledStatus = &v
}

// GetLogsLastCollected returns the LogsLastCollected field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetLogsLastCollected() SecurityMonitoringContentPackTimestampBucket {
	if o == nil {
		var ret SecurityMonitoringContentPackTimestampBucket
		return ret
	}
	return o.LogsLastCollected
}

// GetLogsLastCollectedOk returns a tuple with the LogsLastCollected field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetLogsLastCollectedOk() (*SecurityMonitoringContentPackTimestampBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsLastCollected, true
}

// SetLogsLastCollected sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetLogsLastCollected(v SecurityMonitoringContentPackTimestampBucket) {
	o.LogsLastCollected = v
}

// GetLogsSeenFromAnyIndex returns the LogsSeenFromAnyIndex field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetLogsSeenFromAnyIndex() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.LogsSeenFromAnyIndex
}

// GetLogsSeenFromAnyIndexOk returns a tuple with the LogsSeenFromAnyIndex field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetLogsSeenFromAnyIndexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsSeenFromAnyIndex, true
}

// SetLogsSeenFromAnyIndex sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetLogsSeenFromAnyIndex(v bool) {
	o.LogsSeenFromAnyIndex = v
}

// GetState returns the State field value.
func (o *SecurityMonitoringContentPackStateAttributes) GetState() SecurityMonitoringContentPackStatus {
	if o == nil {
		var ret SecurityMonitoringContentPackStatus
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateAttributes) GetStateOk() (*SecurityMonitoringContentPackStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *SecurityMonitoringContentPackStateAttributes) SetState(v SecurityMonitoringContentPackStatus) {
	o.State = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringContentPackStateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cloud_siem_index_incorrect"] = o.CloudSiemIndexIncorrect
	toSerialize["cp_activation"] = o.CpActivation
	toSerialize["filters_configured_for_logs"] = o.FiltersConfiguredForLogs
	if o.IntegrationInstalledStatus != nil {
		toSerialize["integration_installed_status"] = o.IntegrationInstalledStatus
	}
	toSerialize["logs_last_collected"] = o.LogsLastCollected
	toSerialize["logs_seen_from_any_index"] = o.LogsSeenFromAnyIndex
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringContentPackStateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudSiemIndexIncorrect    *bool                                           `json:"cloud_siem_index_incorrect"`
		CpActivation               *SecurityMonitoringContentPackActivation        `json:"cp_activation"`
		FiltersConfiguredForLogs   *bool                                           `json:"filters_configured_for_logs"`
		IntegrationInstalledStatus *SecurityMonitoringContentPackIntegrationStatus `json:"integration_installed_status,omitempty"`
		LogsLastCollected          *SecurityMonitoringContentPackTimestampBucket   `json:"logs_last_collected"`
		LogsSeenFromAnyIndex       *bool                                           `json:"logs_seen_from_any_index"`
		State                      *SecurityMonitoringContentPackStatus            `json:"state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CloudSiemIndexIncorrect == nil {
		return fmt.Errorf("required field cloud_siem_index_incorrect missing")
	}
	if all.CpActivation == nil {
		return fmt.Errorf("required field cp_activation missing")
	}
	if all.FiltersConfiguredForLogs == nil {
		return fmt.Errorf("required field filters_configured_for_logs missing")
	}
	if all.LogsLastCollected == nil {
		return fmt.Errorf("required field logs_last_collected missing")
	}
	if all.LogsSeenFromAnyIndex == nil {
		return fmt.Errorf("required field logs_seen_from_any_index missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud_siem_index_incorrect", "cp_activation", "filters_configured_for_logs", "integration_installed_status", "logs_last_collected", "logs_seen_from_any_index", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CloudSiemIndexIncorrect = *all.CloudSiemIndexIncorrect
	if !all.CpActivation.IsValid() {
		hasInvalidField = true
	} else {
		o.CpActivation = *all.CpActivation
	}
	o.FiltersConfiguredForLogs = *all.FiltersConfiguredForLogs
	if all.IntegrationInstalledStatus != nil && !all.IntegrationInstalledStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.IntegrationInstalledStatus = all.IntegrationInstalledStatus
	}
	if !all.LogsLastCollected.IsValid() {
		hasInvalidField = true
	} else {
		o.LogsLastCollected = *all.LogsLastCollected
	}
	o.LogsSeenFromAnyIndex = *all.LogsSeenFromAnyIndex
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
