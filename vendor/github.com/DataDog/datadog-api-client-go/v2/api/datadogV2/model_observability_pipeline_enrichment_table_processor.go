// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableProcessor The `enrichment_table` processor enriches logs using a static CSV file, GeoIP database, or reference table. Exactly one of `file`, `geoip`, or `reference_table` must be configured.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineEnrichmentTableProcessor struct {
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// Defines a static enrichment table loaded from a CSV file.
	File *ObservabilityPipelineEnrichmentTableFile `json:"file,omitempty"`
	// Uses a GeoIP database to enrich logs based on an IP field.
	Geoip *ObservabilityPipelineEnrichmentTableGeoIp `json:"geoip,omitempty"`
	// The unique identifier for this processor.
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// Uses a Datadog reference table to enrich logs.
	ReferenceTable *ObservabilityPipelineEnrichmentTableReferenceTable `json:"reference_table,omitempty"`
	// Path where enrichment results should be stored in the log.
	Target string `json:"target"`
	// The processor type. The value should always be `enrichment_table`.
	Type ObservabilityPipelineEnrichmentTableProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableProcessor instantiates a new ObservabilityPipelineEnrichmentTableProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableProcessor(enabled bool, id string, include string, target string, typeVar ObservabilityPipelineEnrichmentTableProcessorType) *ObservabilityPipelineEnrichmentTableProcessor {
	this := ObservabilityPipelineEnrichmentTableProcessor{}
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineEnrichmentTableProcessorWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableProcessorWithDefaults() *ObservabilityPipelineEnrichmentTableProcessor {
	this := ObservabilityPipelineEnrichmentTableProcessor{}
	var typeVar ObservabilityPipelineEnrichmentTableProcessorType = OBSERVABILITYPIPELINEENRICHMENTTABLEPROCESSORTYPE_ENRICHMENT_TABLE
	this.Type = typeVar
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetFile() ObservabilityPipelineEnrichmentTableFile {
	if o == nil || o.File == nil {
		var ret ObservabilityPipelineEnrichmentTableFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetFileOk() (*ObservabilityPipelineEnrichmentTableFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) HasFile() bool {
	return o != nil && o.File != nil
}

// SetFile gets a reference to the given ObservabilityPipelineEnrichmentTableFile and assigns it to the File field.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetFile(v ObservabilityPipelineEnrichmentTableFile) {
	o.File = &v
}

// GetGeoip returns the Geoip field value if set, zero value otherwise.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetGeoip() ObservabilityPipelineEnrichmentTableGeoIp {
	if o == nil || o.Geoip == nil {
		var ret ObservabilityPipelineEnrichmentTableGeoIp
		return ret
	}
	return *o.Geoip
}

// GetGeoipOk returns a tuple with the Geoip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetGeoipOk() (*ObservabilityPipelineEnrichmentTableGeoIp, bool) {
	if o == nil || o.Geoip == nil {
		return nil, false
	}
	return o.Geoip, true
}

// HasGeoip returns a boolean if a field has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) HasGeoip() bool {
	return o != nil && o.Geoip != nil
}

// SetGeoip gets a reference to the given ObservabilityPipelineEnrichmentTableGeoIp and assigns it to the Geoip field.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetGeoip(v ObservabilityPipelineEnrichmentTableGeoIp) {
	o.Geoip = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetInclude(v string) {
	o.Include = v
}

// GetReferenceTable returns the ReferenceTable field value if set, zero value otherwise.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetReferenceTable() ObservabilityPipelineEnrichmentTableReferenceTable {
	if o == nil || o.ReferenceTable == nil {
		var ret ObservabilityPipelineEnrichmentTableReferenceTable
		return ret
	}
	return *o.ReferenceTable
}

// GetReferenceTableOk returns a tuple with the ReferenceTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetReferenceTableOk() (*ObservabilityPipelineEnrichmentTableReferenceTable, bool) {
	if o == nil || o.ReferenceTable == nil {
		return nil, false
	}
	return o.ReferenceTable, true
}

// HasReferenceTable returns a boolean if a field has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) HasReferenceTable() bool {
	return o != nil && o.ReferenceTable != nil
}

// SetReferenceTable gets a reference to the given ObservabilityPipelineEnrichmentTableReferenceTable and assigns it to the ReferenceTable field.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetReferenceTable(v ObservabilityPipelineEnrichmentTableReferenceTable) {
	o.ReferenceTable = &v
}

// GetTarget returns the Target field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetType() ObservabilityPipelineEnrichmentTableProcessorType {
	if o == nil {
		var ret ObservabilityPipelineEnrichmentTableProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableProcessor) GetTypeOk() (*ObservabilityPipelineEnrichmentTableProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineEnrichmentTableProcessor) SetType(v ObservabilityPipelineEnrichmentTableProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Geoip != nil {
		toSerialize["geoip"] = o.Geoip
	}
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	if o.ReferenceTable != nil {
		toSerialize["reference_table"] = o.ReferenceTable
	}
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName    *string                                             `json:"display_name,omitempty"`
		Enabled        *bool                                               `json:"enabled"`
		File           *ObservabilityPipelineEnrichmentTableFile           `json:"file,omitempty"`
		Geoip          *ObservabilityPipelineEnrichmentTableGeoIp          `json:"geoip,omitempty"`
		Id             *string                                             `json:"id"`
		Include        *string                                             `json:"include"`
		ReferenceTable *ObservabilityPipelineEnrichmentTableReferenceTable `json:"reference_table,omitempty"`
		Target         *string                                             `json:"target"`
		Type           *ObservabilityPipelineEnrichmentTableProcessorType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "enabled", "file", "geoip", "id", "include", "reference_table", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	if all.File != nil && all.File.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.File = all.File
	if all.Geoip != nil && all.Geoip.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Geoip = all.Geoip
	o.Id = *all.Id
	o.Include = *all.Include
	if all.ReferenceTable != nil && all.ReferenceTable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReferenceTable = all.ReferenceTable
	o.Target = *all.Target
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
