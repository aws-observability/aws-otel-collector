// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationsByStatus Integrations organized by their status.
type FleetIntegrationsByStatus struct {
	// Configuration files for integrations.
	ConfigurationFiles []FleetConfigurationFile `json:"configuration_files,omitempty"`
	// The unique agent key identifier.
	DatadogAgentKey *string `json:"datadog_agent_key,omitempty"`
	// Integrations with errors.
	ErrorIntegrations []FleetIntegrationDetails `json:"error_integrations,omitempty"`
	// Detected but not configured integrations.
	MissingIntegrations []FleetDetectedIntegration `json:"missing_integrations,omitempty"`
	// Integrations with warnings.
	WarningIntegrations []FleetIntegrationDetails `json:"warning_integrations,omitempty"`
	// Integrations that are working correctly.
	WorkingIntegrations []FleetIntegrationDetails `json:"working_integrations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationsByStatus instantiates a new FleetIntegrationsByStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationsByStatus() *FleetIntegrationsByStatus {
	this := FleetIntegrationsByStatus{}
	return &this
}

// NewFleetIntegrationsByStatusWithDefaults instantiates a new FleetIntegrationsByStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationsByStatusWithDefaults() *FleetIntegrationsByStatus {
	this := FleetIntegrationsByStatus{}
	return &this
}

// GetConfigurationFiles returns the ConfigurationFiles field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatus) GetConfigurationFiles() []FleetConfigurationFile {
	if o == nil || o.ConfigurationFiles == nil {
		var ret []FleetConfigurationFile
		return ret
	}
	return o.ConfigurationFiles
}

// GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatus) GetConfigurationFilesOk() (*[]FleetConfigurationFile, bool) {
	if o == nil || o.ConfigurationFiles == nil {
		return nil, false
	}
	return &o.ConfigurationFiles, true
}

// HasConfigurationFiles returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatus) HasConfigurationFiles() bool {
	return o != nil && o.ConfigurationFiles != nil
}

// SetConfigurationFiles gets a reference to the given []FleetConfigurationFile and assigns it to the ConfigurationFiles field.
func (o *FleetIntegrationsByStatus) SetConfigurationFiles(v []FleetConfigurationFile) {
	o.ConfigurationFiles = v
}

// GetDatadogAgentKey returns the DatadogAgentKey field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatus) GetDatadogAgentKey() string {
	if o == nil || o.DatadogAgentKey == nil {
		var ret string
		return ret
	}
	return *o.DatadogAgentKey
}

// GetDatadogAgentKeyOk returns a tuple with the DatadogAgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatus) GetDatadogAgentKeyOk() (*string, bool) {
	if o == nil || o.DatadogAgentKey == nil {
		return nil, false
	}
	return o.DatadogAgentKey, true
}

// HasDatadogAgentKey returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatus) HasDatadogAgentKey() bool {
	return o != nil && o.DatadogAgentKey != nil
}

// SetDatadogAgentKey gets a reference to the given string and assigns it to the DatadogAgentKey field.
func (o *FleetIntegrationsByStatus) SetDatadogAgentKey(v string) {
	o.DatadogAgentKey = &v
}

// GetErrorIntegrations returns the ErrorIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatus) GetErrorIntegrations() []FleetIntegrationDetails {
	if o == nil || o.ErrorIntegrations == nil {
		var ret []FleetIntegrationDetails
		return ret
	}
	return o.ErrorIntegrations
}

// GetErrorIntegrationsOk returns a tuple with the ErrorIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatus) GetErrorIntegrationsOk() (*[]FleetIntegrationDetails, bool) {
	if o == nil || o.ErrorIntegrations == nil {
		return nil, false
	}
	return &o.ErrorIntegrations, true
}

// HasErrorIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatus) HasErrorIntegrations() bool {
	return o != nil && o.ErrorIntegrations != nil
}

// SetErrorIntegrations gets a reference to the given []FleetIntegrationDetails and assigns it to the ErrorIntegrations field.
func (o *FleetIntegrationsByStatus) SetErrorIntegrations(v []FleetIntegrationDetails) {
	o.ErrorIntegrations = v
}

// GetMissingIntegrations returns the MissingIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatus) GetMissingIntegrations() []FleetDetectedIntegration {
	if o == nil || o.MissingIntegrations == nil {
		var ret []FleetDetectedIntegration
		return ret
	}
	return o.MissingIntegrations
}

// GetMissingIntegrationsOk returns a tuple with the MissingIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatus) GetMissingIntegrationsOk() (*[]FleetDetectedIntegration, bool) {
	if o == nil || o.MissingIntegrations == nil {
		return nil, false
	}
	return &o.MissingIntegrations, true
}

// HasMissingIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatus) HasMissingIntegrations() bool {
	return o != nil && o.MissingIntegrations != nil
}

// SetMissingIntegrations gets a reference to the given []FleetDetectedIntegration and assigns it to the MissingIntegrations field.
func (o *FleetIntegrationsByStatus) SetMissingIntegrations(v []FleetDetectedIntegration) {
	o.MissingIntegrations = v
}

// GetWarningIntegrations returns the WarningIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatus) GetWarningIntegrations() []FleetIntegrationDetails {
	if o == nil || o.WarningIntegrations == nil {
		var ret []FleetIntegrationDetails
		return ret
	}
	return o.WarningIntegrations
}

// GetWarningIntegrationsOk returns a tuple with the WarningIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatus) GetWarningIntegrationsOk() (*[]FleetIntegrationDetails, bool) {
	if o == nil || o.WarningIntegrations == nil {
		return nil, false
	}
	return &o.WarningIntegrations, true
}

// HasWarningIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatus) HasWarningIntegrations() bool {
	return o != nil && o.WarningIntegrations != nil
}

// SetWarningIntegrations gets a reference to the given []FleetIntegrationDetails and assigns it to the WarningIntegrations field.
func (o *FleetIntegrationsByStatus) SetWarningIntegrations(v []FleetIntegrationDetails) {
	o.WarningIntegrations = v
}

// GetWorkingIntegrations returns the WorkingIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatus) GetWorkingIntegrations() []FleetIntegrationDetails {
	if o == nil || o.WorkingIntegrations == nil {
		var ret []FleetIntegrationDetails
		return ret
	}
	return o.WorkingIntegrations
}

// GetWorkingIntegrationsOk returns a tuple with the WorkingIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatus) GetWorkingIntegrationsOk() (*[]FleetIntegrationDetails, bool) {
	if o == nil || o.WorkingIntegrations == nil {
		return nil, false
	}
	return &o.WorkingIntegrations, true
}

// HasWorkingIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatus) HasWorkingIntegrations() bool {
	return o != nil && o.WorkingIntegrations != nil
}

// SetWorkingIntegrations gets a reference to the given []FleetIntegrationDetails and assigns it to the WorkingIntegrations field.
func (o *FleetIntegrationsByStatus) SetWorkingIntegrations(v []FleetIntegrationDetails) {
	o.WorkingIntegrations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfigurationFiles != nil {
		toSerialize["configuration_files"] = o.ConfigurationFiles
	}
	if o.DatadogAgentKey != nil {
		toSerialize["datadog_agent_key"] = o.DatadogAgentKey
	}
	if o.ErrorIntegrations != nil {
		toSerialize["error_integrations"] = o.ErrorIntegrations
	}
	if o.MissingIntegrations != nil {
		toSerialize["missing_integrations"] = o.MissingIntegrations
	}
	if o.WarningIntegrations != nil {
		toSerialize["warning_integrations"] = o.WarningIntegrations
	}
	if o.WorkingIntegrations != nil {
		toSerialize["working_integrations"] = o.WorkingIntegrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationsByStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigurationFiles  []FleetConfigurationFile   `json:"configuration_files,omitempty"`
		DatadogAgentKey     *string                    `json:"datadog_agent_key,omitempty"`
		ErrorIntegrations   []FleetIntegrationDetails  `json:"error_integrations,omitempty"`
		MissingIntegrations []FleetDetectedIntegration `json:"missing_integrations,omitempty"`
		WarningIntegrations []FleetIntegrationDetails  `json:"warning_integrations,omitempty"`
		WorkingIntegrations []FleetIntegrationDetails  `json:"working_integrations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration_files", "datadog_agent_key", "error_integrations", "missing_integrations", "warning_integrations", "working_integrations"})
	} else {
		return err
	}
	o.ConfigurationFiles = all.ConfigurationFiles
	o.DatadogAgentKey = all.DatadogAgentKey
	o.ErrorIntegrations = all.ErrorIntegrations
	o.MissingIntegrations = all.MissingIntegrations
	o.WarningIntegrations = all.WarningIntegrations
	o.WorkingIntegrations = all.WorkingIntegrations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
