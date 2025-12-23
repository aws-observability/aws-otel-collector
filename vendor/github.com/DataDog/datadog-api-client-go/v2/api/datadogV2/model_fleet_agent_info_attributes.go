// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentInfoAttributes Attributes for agent information.
type FleetAgentInfoAttributes struct {
	// Detailed information about a Datadog Agent.
	AgentInfos *FleetAgentInfoDetails `json:"agent_infos,omitempty"`
	// Configuration information organized by layers.
	ConfigurationFiles *FleetConfigurationLayer `json:"configuration_files,omitempty"`
	// List of detected integrations.
	DetectedIntegrations []FleetDetectedIntegration `json:"detected_integrations,omitempty"`
	// Integrations organized by their status.
	Integrations *FleetIntegrationsByStatus `json:"integrations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAgentInfoAttributes instantiates a new FleetAgentInfoAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAgentInfoAttributes() *FleetAgentInfoAttributes {
	this := FleetAgentInfoAttributes{}
	return &this
}

// NewFleetAgentInfoAttributesWithDefaults instantiates a new FleetAgentInfoAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAgentInfoAttributesWithDefaults() *FleetAgentInfoAttributes {
	this := FleetAgentInfoAttributes{}
	return &this
}

// GetAgentInfos returns the AgentInfos field value if set, zero value otherwise.
func (o *FleetAgentInfoAttributes) GetAgentInfos() FleetAgentInfoDetails {
	if o == nil || o.AgentInfos == nil {
		var ret FleetAgentInfoDetails
		return ret
	}
	return *o.AgentInfos
}

// GetAgentInfosOk returns a tuple with the AgentInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoAttributes) GetAgentInfosOk() (*FleetAgentInfoDetails, bool) {
	if o == nil || o.AgentInfos == nil {
		return nil, false
	}
	return o.AgentInfos, true
}

// HasAgentInfos returns a boolean if a field has been set.
func (o *FleetAgentInfoAttributes) HasAgentInfos() bool {
	return o != nil && o.AgentInfos != nil
}

// SetAgentInfos gets a reference to the given FleetAgentInfoDetails and assigns it to the AgentInfos field.
func (o *FleetAgentInfoAttributes) SetAgentInfos(v FleetAgentInfoDetails) {
	o.AgentInfos = &v
}

// GetConfigurationFiles returns the ConfigurationFiles field value if set, zero value otherwise.
func (o *FleetAgentInfoAttributes) GetConfigurationFiles() FleetConfigurationLayer {
	if o == nil || o.ConfigurationFiles == nil {
		var ret FleetConfigurationLayer
		return ret
	}
	return *o.ConfigurationFiles
}

// GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoAttributes) GetConfigurationFilesOk() (*FleetConfigurationLayer, bool) {
	if o == nil || o.ConfigurationFiles == nil {
		return nil, false
	}
	return o.ConfigurationFiles, true
}

// HasConfigurationFiles returns a boolean if a field has been set.
func (o *FleetAgentInfoAttributes) HasConfigurationFiles() bool {
	return o != nil && o.ConfigurationFiles != nil
}

// SetConfigurationFiles gets a reference to the given FleetConfigurationLayer and assigns it to the ConfigurationFiles field.
func (o *FleetAgentInfoAttributes) SetConfigurationFiles(v FleetConfigurationLayer) {
	o.ConfigurationFiles = &v
}

// GetDetectedIntegrations returns the DetectedIntegrations field value if set, zero value otherwise.
func (o *FleetAgentInfoAttributes) GetDetectedIntegrations() []FleetDetectedIntegration {
	if o == nil || o.DetectedIntegrations == nil {
		var ret []FleetDetectedIntegration
		return ret
	}
	return o.DetectedIntegrations
}

// GetDetectedIntegrationsOk returns a tuple with the DetectedIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoAttributes) GetDetectedIntegrationsOk() (*[]FleetDetectedIntegration, bool) {
	if o == nil || o.DetectedIntegrations == nil {
		return nil, false
	}
	return &o.DetectedIntegrations, true
}

// HasDetectedIntegrations returns a boolean if a field has been set.
func (o *FleetAgentInfoAttributes) HasDetectedIntegrations() bool {
	return o != nil && o.DetectedIntegrations != nil
}

// SetDetectedIntegrations gets a reference to the given []FleetDetectedIntegration and assigns it to the DetectedIntegrations field.
func (o *FleetAgentInfoAttributes) SetDetectedIntegrations(v []FleetDetectedIntegration) {
	o.DetectedIntegrations = v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *FleetAgentInfoAttributes) GetIntegrations() FleetIntegrationsByStatus {
	if o == nil || o.Integrations == nil {
		var ret FleetIntegrationsByStatus
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAgentInfoAttributes) GetIntegrationsOk() (*FleetIntegrationsByStatus, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *FleetAgentInfoAttributes) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given FleetIntegrationsByStatus and assigns it to the Integrations field.
func (o *FleetAgentInfoAttributes) SetIntegrations(v FleetIntegrationsByStatus) {
	o.Integrations = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAgentInfoAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentInfos != nil {
		toSerialize["agent_infos"] = o.AgentInfos
	}
	if o.ConfigurationFiles != nil {
		toSerialize["configuration_files"] = o.ConfigurationFiles
	}
	if o.DetectedIntegrations != nil {
		toSerialize["detected_integrations"] = o.DetectedIntegrations
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAgentInfoAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentInfos           *FleetAgentInfoDetails     `json:"agent_infos,omitempty"`
		ConfigurationFiles   *FleetConfigurationLayer   `json:"configuration_files,omitempty"`
		DetectedIntegrations []FleetDetectedIntegration `json:"detected_integrations,omitempty"`
		Integrations         *FleetIntegrationsByStatus `json:"integrations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_infos", "configuration_files", "detected_integrations", "integrations"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AgentInfos != nil && all.AgentInfos.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AgentInfos = all.AgentInfos
	if all.ConfigurationFiles != nil && all.ConfigurationFiles.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfigurationFiles = all.ConfigurationFiles
	o.DetectedIntegrations = all.DetectedIntegrations
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integrations = all.Integrations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
