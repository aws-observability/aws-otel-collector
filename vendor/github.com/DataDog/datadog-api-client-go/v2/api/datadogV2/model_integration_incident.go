// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationIncident Incident integration settings.
type IntegrationIncident struct {
	// Query for auto-escalation.
	AutoEscalationQuery *string `json:"auto_escalation_query,omitempty"`
	// Default incident commander.
	DefaultIncidentCommander *string `json:"default_incident_commander,omitempty"`
	// Whether incident integration is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of mappings between incident fields and case fields.
	FieldMappings []IntegrationIncidentFieldMappingsItems `json:"field_mappings,omitempty"`
	// Incident type.
	IncidentType *string `json:"incident_type,omitempty"`
	// Severity configuration for mapping incident priorities to case priorities.
	SeverityConfig *IntegrationIncidentSeverityConfig `json:"severity_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationIncident instantiates a new IntegrationIncident object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationIncident() *IntegrationIncident {
	this := IntegrationIncident{}
	return &this
}

// NewIntegrationIncidentWithDefaults instantiates a new IntegrationIncident object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationIncidentWithDefaults() *IntegrationIncident {
	this := IntegrationIncident{}
	return &this
}

// GetAutoEscalationQuery returns the AutoEscalationQuery field value if set, zero value otherwise.
func (o *IntegrationIncident) GetAutoEscalationQuery() string {
	if o == nil || o.AutoEscalationQuery == nil {
		var ret string
		return ret
	}
	return *o.AutoEscalationQuery
}

// GetAutoEscalationQueryOk returns a tuple with the AutoEscalationQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncident) GetAutoEscalationQueryOk() (*string, bool) {
	if o == nil || o.AutoEscalationQuery == nil {
		return nil, false
	}
	return o.AutoEscalationQuery, true
}

// HasAutoEscalationQuery returns a boolean if a field has been set.
func (o *IntegrationIncident) HasAutoEscalationQuery() bool {
	return o != nil && o.AutoEscalationQuery != nil
}

// SetAutoEscalationQuery gets a reference to the given string and assigns it to the AutoEscalationQuery field.
func (o *IntegrationIncident) SetAutoEscalationQuery(v string) {
	o.AutoEscalationQuery = &v
}

// GetDefaultIncidentCommander returns the DefaultIncidentCommander field value if set, zero value otherwise.
func (o *IntegrationIncident) GetDefaultIncidentCommander() string {
	if o == nil || o.DefaultIncidentCommander == nil {
		var ret string
		return ret
	}
	return *o.DefaultIncidentCommander
}

// GetDefaultIncidentCommanderOk returns a tuple with the DefaultIncidentCommander field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncident) GetDefaultIncidentCommanderOk() (*string, bool) {
	if o == nil || o.DefaultIncidentCommander == nil {
		return nil, false
	}
	return o.DefaultIncidentCommander, true
}

// HasDefaultIncidentCommander returns a boolean if a field has been set.
func (o *IntegrationIncident) HasDefaultIncidentCommander() bool {
	return o != nil && o.DefaultIncidentCommander != nil
}

// SetDefaultIncidentCommander gets a reference to the given string and assigns it to the DefaultIncidentCommander field.
func (o *IntegrationIncident) SetDefaultIncidentCommander(v string) {
	o.DefaultIncidentCommander = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationIncident) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncident) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationIncident) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationIncident) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise.
func (o *IntegrationIncident) GetFieldMappings() []IntegrationIncidentFieldMappingsItems {
	if o == nil || o.FieldMappings == nil {
		var ret []IntegrationIncidentFieldMappingsItems
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncident) GetFieldMappingsOk() (*[]IntegrationIncidentFieldMappingsItems, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *IntegrationIncident) HasFieldMappings() bool {
	return o != nil && o.FieldMappings != nil
}

// SetFieldMappings gets a reference to the given []IntegrationIncidentFieldMappingsItems and assigns it to the FieldMappings field.
func (o *IntegrationIncident) SetFieldMappings(v []IntegrationIncidentFieldMappingsItems) {
	o.FieldMappings = v
}

// GetIncidentType returns the IncidentType field value if set, zero value otherwise.
func (o *IntegrationIncident) GetIncidentType() string {
	if o == nil || o.IncidentType == nil {
		var ret string
		return ret
	}
	return *o.IncidentType
}

// GetIncidentTypeOk returns a tuple with the IncidentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncident) GetIncidentTypeOk() (*string, bool) {
	if o == nil || o.IncidentType == nil {
		return nil, false
	}
	return o.IncidentType, true
}

// HasIncidentType returns a boolean if a field has been set.
func (o *IntegrationIncident) HasIncidentType() bool {
	return o != nil && o.IncidentType != nil
}

// SetIncidentType gets a reference to the given string and assigns it to the IncidentType field.
func (o *IntegrationIncident) SetIncidentType(v string) {
	o.IncidentType = &v
}

// GetSeverityConfig returns the SeverityConfig field value if set, zero value otherwise.
func (o *IntegrationIncident) GetSeverityConfig() IntegrationIncidentSeverityConfig {
	if o == nil || o.SeverityConfig == nil {
		var ret IntegrationIncidentSeverityConfig
		return ret
	}
	return *o.SeverityConfig
}

// GetSeverityConfigOk returns a tuple with the SeverityConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncident) GetSeverityConfigOk() (*IntegrationIncidentSeverityConfig, bool) {
	if o == nil || o.SeverityConfig == nil {
		return nil, false
	}
	return o.SeverityConfig, true
}

// HasSeverityConfig returns a boolean if a field has been set.
func (o *IntegrationIncident) HasSeverityConfig() bool {
	return o != nil && o.SeverityConfig != nil
}

// SetSeverityConfig gets a reference to the given IntegrationIncidentSeverityConfig and assigns it to the SeverityConfig field.
func (o *IntegrationIncident) SetSeverityConfig(v IntegrationIncidentSeverityConfig) {
	o.SeverityConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationIncident) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoEscalationQuery != nil {
		toSerialize["auto_escalation_query"] = o.AutoEscalationQuery
	}
	if o.DefaultIncidentCommander != nil {
		toSerialize["default_incident_commander"] = o.DefaultIncidentCommander
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	if o.IncidentType != nil {
		toSerialize["incident_type"] = o.IncidentType
	}
	if o.SeverityConfig != nil {
		toSerialize["severity_config"] = o.SeverityConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationIncident) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoEscalationQuery      *string                                 `json:"auto_escalation_query,omitempty"`
		DefaultIncidentCommander *string                                 `json:"default_incident_commander,omitempty"`
		Enabled                  *bool                                   `json:"enabled,omitempty"`
		FieldMappings            []IntegrationIncidentFieldMappingsItems `json:"field_mappings,omitempty"`
		IncidentType             *string                                 `json:"incident_type,omitempty"`
		SeverityConfig           *IntegrationIncidentSeverityConfig      `json:"severity_config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_escalation_query", "default_incident_commander", "enabled", "field_mappings", "incident_type", "severity_config"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoEscalationQuery = all.AutoEscalationQuery
	o.DefaultIncidentCommander = all.DefaultIncidentCommander
	o.Enabled = all.Enabled
	o.FieldMappings = all.FieldMappings
	o.IncidentType = all.IncidentType
	if all.SeverityConfig != nil && all.SeverityConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SeverityConfig = all.SeverityConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
