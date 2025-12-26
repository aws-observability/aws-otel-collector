// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationDetails Detailed information about a single integration.
type FleetIntegrationDetails struct {
	// Type of data collected (metrics, logs).
	DataType *string `json:"data_type,omitempty"`
	// Error messages if the integration has issues.
	ErrorMessages []string `json:"error_messages,omitempty"`
	// Initialization configuration (YAML format).
	InitConfig *string `json:"init_config,omitempty"`
	// Instance-specific configuration (YAML format).
	InstanceConfig *string `json:"instance_config,omitempty"`
	// Whether this is a custom integration.
	IsCustomCheck *bool `json:"is_custom_check,omitempty"`
	// Log collection configuration (YAML format).
	LogConfig *string `json:"log_config,omitempty"`
	// Name of the integration instance.
	Name *string `json:"name,omitempty"`
	// Index in the configuration file.
	SourceIndex *int64 `json:"source_index,omitempty"`
	// Path to the configuration file.
	SourcePath *string `json:"source_path,omitempty"`
	// Integration type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationDetails instantiates a new FleetIntegrationDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationDetails() *FleetIntegrationDetails {
	this := FleetIntegrationDetails{}
	return &this
}

// NewFleetIntegrationDetailsWithDefaults instantiates a new FleetIntegrationDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationDetailsWithDefaults() *FleetIntegrationDetails {
	this := FleetIntegrationDetails{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasDataType() bool {
	return o != nil && o.DataType != nil
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *FleetIntegrationDetails) SetDataType(v string) {
	o.DataType = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages == nil {
		var ret []string
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil || o.ErrorMessages == nil {
		return nil, false
	}
	return &o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages != nil
}

// SetErrorMessages gets a reference to the given []string and assigns it to the ErrorMessages field.
func (o *FleetIntegrationDetails) SetErrorMessages(v []string) {
	o.ErrorMessages = v
}

// GetInitConfig returns the InitConfig field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetInitConfig() string {
	if o == nil || o.InitConfig == nil {
		var ret string
		return ret
	}
	return *o.InitConfig
}

// GetInitConfigOk returns a tuple with the InitConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetInitConfigOk() (*string, bool) {
	if o == nil || o.InitConfig == nil {
		return nil, false
	}
	return o.InitConfig, true
}

// HasInitConfig returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasInitConfig() bool {
	return o != nil && o.InitConfig != nil
}

// SetInitConfig gets a reference to the given string and assigns it to the InitConfig field.
func (o *FleetIntegrationDetails) SetInitConfig(v string) {
	o.InitConfig = &v
}

// GetInstanceConfig returns the InstanceConfig field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetInstanceConfig() string {
	if o == nil || o.InstanceConfig == nil {
		var ret string
		return ret
	}
	return *o.InstanceConfig
}

// GetInstanceConfigOk returns a tuple with the InstanceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetInstanceConfigOk() (*string, bool) {
	if o == nil || o.InstanceConfig == nil {
		return nil, false
	}
	return o.InstanceConfig, true
}

// HasInstanceConfig returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasInstanceConfig() bool {
	return o != nil && o.InstanceConfig != nil
}

// SetInstanceConfig gets a reference to the given string and assigns it to the InstanceConfig field.
func (o *FleetIntegrationDetails) SetInstanceConfig(v string) {
	o.InstanceConfig = &v
}

// GetIsCustomCheck returns the IsCustomCheck field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetIsCustomCheck() bool {
	if o == nil || o.IsCustomCheck == nil {
		var ret bool
		return ret
	}
	return *o.IsCustomCheck
}

// GetIsCustomCheckOk returns a tuple with the IsCustomCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetIsCustomCheckOk() (*bool, bool) {
	if o == nil || o.IsCustomCheck == nil {
		return nil, false
	}
	return o.IsCustomCheck, true
}

// HasIsCustomCheck returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasIsCustomCheck() bool {
	return o != nil && o.IsCustomCheck != nil
}

// SetIsCustomCheck gets a reference to the given bool and assigns it to the IsCustomCheck field.
func (o *FleetIntegrationDetails) SetIsCustomCheck(v bool) {
	o.IsCustomCheck = &v
}

// GetLogConfig returns the LogConfig field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetLogConfig() string {
	if o == nil || o.LogConfig == nil {
		var ret string
		return ret
	}
	return *o.LogConfig
}

// GetLogConfigOk returns a tuple with the LogConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetLogConfigOk() (*string, bool) {
	if o == nil || o.LogConfig == nil {
		return nil, false
	}
	return o.LogConfig, true
}

// HasLogConfig returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasLogConfig() bool {
	return o != nil && o.LogConfig != nil
}

// SetLogConfig gets a reference to the given string and assigns it to the LogConfig field.
func (o *FleetIntegrationDetails) SetLogConfig(v string) {
	o.LogConfig = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetIntegrationDetails) SetName(v string) {
	o.Name = &v
}

// GetSourceIndex returns the SourceIndex field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetSourceIndex() int64 {
	if o == nil || o.SourceIndex == nil {
		var ret int64
		return ret
	}
	return *o.SourceIndex
}

// GetSourceIndexOk returns a tuple with the SourceIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetSourceIndexOk() (*int64, bool) {
	if o == nil || o.SourceIndex == nil {
		return nil, false
	}
	return o.SourceIndex, true
}

// HasSourceIndex returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasSourceIndex() bool {
	return o != nil && o.SourceIndex != nil
}

// SetSourceIndex gets a reference to the given int64 and assigns it to the SourceIndex field.
func (o *FleetIntegrationDetails) SetSourceIndex(v int64) {
	o.SourceIndex = &v
}

// GetSourcePath returns the SourcePath field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetSourcePath() string {
	if o == nil || o.SourcePath == nil {
		var ret string
		return ret
	}
	return *o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetSourcePathOk() (*string, bool) {
	if o == nil || o.SourcePath == nil {
		return nil, false
	}
	return o.SourcePath, true
}

// HasSourcePath returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasSourcePath() bool {
	return o != nil && o.SourcePath != nil
}

// SetSourcePath gets a reference to the given string and assigns it to the SourcePath field.
func (o *FleetIntegrationDetails) SetSourcePath(v string) {
	o.SourcePath = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FleetIntegrationDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FleetIntegrationDetails) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FleetIntegrationDetails) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataType != nil {
		toSerialize["data_type"] = o.DataType
	}
	if o.ErrorMessages != nil {
		toSerialize["error_messages"] = o.ErrorMessages
	}
	if o.InitConfig != nil {
		toSerialize["init_config"] = o.InitConfig
	}
	if o.InstanceConfig != nil {
		toSerialize["instance_config"] = o.InstanceConfig
	}
	if o.IsCustomCheck != nil {
		toSerialize["is_custom_check"] = o.IsCustomCheck
	}
	if o.LogConfig != nil {
		toSerialize["log_config"] = o.LogConfig
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SourceIndex != nil {
		toSerialize["source_index"] = o.SourceIndex
	}
	if o.SourcePath != nil {
		toSerialize["source_path"] = o.SourcePath
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataType       *string  `json:"data_type,omitempty"`
		ErrorMessages  []string `json:"error_messages,omitempty"`
		InitConfig     *string  `json:"init_config,omitempty"`
		InstanceConfig *string  `json:"instance_config,omitempty"`
		IsCustomCheck  *bool    `json:"is_custom_check,omitempty"`
		LogConfig      *string  `json:"log_config,omitempty"`
		Name           *string  `json:"name,omitempty"`
		SourceIndex    *int64   `json:"source_index,omitempty"`
		SourcePath     *string  `json:"source_path,omitempty"`
		Type           *string  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_type", "error_messages", "init_config", "instance_config", "is_custom_check", "log_config", "name", "source_index", "source_path", "type"})
	} else {
		return err
	}
	o.DataType = all.DataType
	o.ErrorMessages = all.ErrorMessages
	o.InitConfig = all.InitConfig
	o.InstanceConfig = all.InstanceConfig
	o.IsCustomCheck = all.IsCustomCheck
	o.LogConfig = all.LogConfig
	o.Name = all.Name
	o.SourceIndex = all.SourceIndex
	o.SourcePath = all.SourcePath
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
