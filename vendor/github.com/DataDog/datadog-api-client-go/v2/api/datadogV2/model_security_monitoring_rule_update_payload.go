// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleUpdatePayload Update an existing rule.
type SecurityMonitoringRuleUpdatePayload struct {
	// Calculated fields. Only allowed for scheduled rules - in other words, when schedulingOptions is also defined.
	CalculatedFields []CalculatedField `json:"calculatedFields,omitempty"`
	// Cases for generating signals.
	Cases []SecurityMonitoringRuleCase `json:"cases,omitempty"`
	// How to generate compliance signals. Useful for cloud_configuration rules only.
	ComplianceSignalOptions *CloudConfigurationRuleComplianceSignalOptions `json:"complianceSignalOptions,omitempty"`
	// Custom/Overridden Message for generated signals (used in case of Default rule update).
	CustomMessage *string `json:"customMessage,omitempty"`
	// Custom/Overridden name (used in case of Default rule update).
	CustomName *string `json:"customName,omitempty"`
	// Additional queries to filter matched events before they are processed. This field is deprecated for log detection, signal correlation, and workload security rules.
	Filters []SecurityMonitoringFilter `json:"filters,omitempty"`
	// Additional grouping to perform on top of the existing groups in the query section. Must be a subset of the existing groups.
	GroupSignalsBy []string `json:"groupSignalsBy,omitempty"`
	// Whether the notifications include the triggering group-by values in their title.
	HasExtendedTitle *bool `json:"hasExtendedTitle,omitempty"`
	// Whether the rule is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Message for generated signals.
	Message *string `json:"message,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Options.
	Options *SecurityMonitoringRuleOptions `json:"options,omitempty"`
	// Queries for selecting logs which are part of the rule.
	Queries []SecurityMonitoringRuleQuery `json:"queries,omitempty"`
	// Reference tables for the rule.
	ReferenceTables []SecurityMonitoringReferenceTable `json:"referenceTables,omitempty"`
	// Options for scheduled rules. When this field is present, the rule runs based on the schedule. When absent, it runs real-time on ingested logs.
	SchedulingOptions NullableSecurityMonitoringSchedulingOptions `json:"schedulingOptions,omitempty"`
	// Tags for generated signals.
	Tags []string `json:"tags,omitempty"`
	// Cases for generating signals from third-party rules. Only available for third-party rules.
	ThirdPartyCases []SecurityMonitoringThirdPartyRuleCase `json:"thirdPartyCases,omitempty"`
	// The version of the rule being updated.
	Version *int32 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleUpdatePayload instantiates a new SecurityMonitoringRuleUpdatePayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleUpdatePayload() *SecurityMonitoringRuleUpdatePayload {
	this := SecurityMonitoringRuleUpdatePayload{}
	return &this
}

// NewSecurityMonitoringRuleUpdatePayloadWithDefaults instantiates a new SecurityMonitoringRuleUpdatePayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleUpdatePayloadWithDefaults() *SecurityMonitoringRuleUpdatePayload {
	this := SecurityMonitoringRuleUpdatePayload{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetCalculatedFields() []CalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []CalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetCalculatedFieldsOk() (*[]CalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []CalculatedField and assigns it to the CalculatedFields field.
func (o *SecurityMonitoringRuleUpdatePayload) SetCalculatedFields(v []CalculatedField) {
	o.CalculatedFields = v
}

// GetCases returns the Cases field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetCases() []SecurityMonitoringRuleCase {
	if o == nil || o.Cases == nil {
		var ret []SecurityMonitoringRuleCase
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetCasesOk() (*[]SecurityMonitoringRuleCase, bool) {
	if o == nil || o.Cases == nil {
		return nil, false
	}
	return &o.Cases, true
}

// HasCases returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasCases() bool {
	return o != nil && o.Cases != nil
}

// SetCases gets a reference to the given []SecurityMonitoringRuleCase and assigns it to the Cases field.
func (o *SecurityMonitoringRuleUpdatePayload) SetCases(v []SecurityMonitoringRuleCase) {
	o.Cases = v
}

// GetComplianceSignalOptions returns the ComplianceSignalOptions field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetComplianceSignalOptions() CloudConfigurationRuleComplianceSignalOptions {
	if o == nil || o.ComplianceSignalOptions == nil {
		var ret CloudConfigurationRuleComplianceSignalOptions
		return ret
	}
	return *o.ComplianceSignalOptions
}

// GetComplianceSignalOptionsOk returns a tuple with the ComplianceSignalOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetComplianceSignalOptionsOk() (*CloudConfigurationRuleComplianceSignalOptions, bool) {
	if o == nil || o.ComplianceSignalOptions == nil {
		return nil, false
	}
	return o.ComplianceSignalOptions, true
}

// HasComplianceSignalOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasComplianceSignalOptions() bool {
	return o != nil && o.ComplianceSignalOptions != nil
}

// SetComplianceSignalOptions gets a reference to the given CloudConfigurationRuleComplianceSignalOptions and assigns it to the ComplianceSignalOptions field.
func (o *SecurityMonitoringRuleUpdatePayload) SetComplianceSignalOptions(v CloudConfigurationRuleComplianceSignalOptions) {
	o.ComplianceSignalOptions = &v
}

// GetCustomMessage returns the CustomMessage field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetCustomMessage() string {
	if o == nil || o.CustomMessage == nil {
		var ret string
		return ret
	}
	return *o.CustomMessage
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetCustomMessageOk() (*string, bool) {
	if o == nil || o.CustomMessage == nil {
		return nil, false
	}
	return o.CustomMessage, true
}

// HasCustomMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasCustomMessage() bool {
	return o != nil && o.CustomMessage != nil
}

// SetCustomMessage gets a reference to the given string and assigns it to the CustomMessage field.
func (o *SecurityMonitoringRuleUpdatePayload) SetCustomMessage(v string) {
	o.CustomMessage = &v
}

// GetCustomName returns the CustomName field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetCustomName() string {
	if o == nil || o.CustomName == nil {
		var ret string
		return ret
	}
	return *o.CustomName
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetCustomNameOk() (*string, bool) {
	if o == nil || o.CustomName == nil {
		return nil, false
	}
	return o.CustomName, true
}

// HasCustomName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasCustomName() bool {
	return o != nil && o.CustomName != nil
}

// SetCustomName gets a reference to the given string and assigns it to the CustomName field.
func (o *SecurityMonitoringRuleUpdatePayload) SetCustomName(v string) {
	o.CustomName = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetFilters() []SecurityMonitoringFilter {
	if o == nil || o.Filters == nil {
		var ret []SecurityMonitoringFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetFiltersOk() (*[]SecurityMonitoringFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []SecurityMonitoringFilter and assigns it to the Filters field.
func (o *SecurityMonitoringRuleUpdatePayload) SetFilters(v []SecurityMonitoringFilter) {
	o.Filters = v
}

// GetGroupSignalsBy returns the GroupSignalsBy field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetGroupSignalsBy() []string {
	if o == nil || o.GroupSignalsBy == nil {
		var ret []string
		return ret
	}
	return o.GroupSignalsBy
}

// GetGroupSignalsByOk returns a tuple with the GroupSignalsBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetGroupSignalsByOk() (*[]string, bool) {
	if o == nil || o.GroupSignalsBy == nil {
		return nil, false
	}
	return &o.GroupSignalsBy, true
}

// HasGroupSignalsBy returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasGroupSignalsBy() bool {
	return o != nil && o.GroupSignalsBy != nil
}

// SetGroupSignalsBy gets a reference to the given []string and assigns it to the GroupSignalsBy field.
func (o *SecurityMonitoringRuleUpdatePayload) SetGroupSignalsBy(v []string) {
	o.GroupSignalsBy = v
}

// GetHasExtendedTitle returns the HasExtendedTitle field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetHasExtendedTitle() bool {
	if o == nil || o.HasExtendedTitle == nil {
		var ret bool
		return ret
	}
	return *o.HasExtendedTitle
}

// GetHasExtendedTitleOk returns a tuple with the HasExtendedTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetHasExtendedTitleOk() (*bool, bool) {
	if o == nil || o.HasExtendedTitle == nil {
		return nil, false
	}
	return o.HasExtendedTitle, true
}

// HasHasExtendedTitle returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasHasExtendedTitle() bool {
	return o != nil && o.HasExtendedTitle != nil
}

// SetHasExtendedTitle gets a reference to the given bool and assigns it to the HasExtendedTitle field.
func (o *SecurityMonitoringRuleUpdatePayload) SetHasExtendedTitle(v bool) {
	o.HasExtendedTitle = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SecurityMonitoringRuleUpdatePayload) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SecurityMonitoringRuleUpdatePayload) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringRuleUpdatePayload) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetOptions() SecurityMonitoringRuleOptions {
	if o == nil || o.Options == nil {
		var ret SecurityMonitoringRuleOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given SecurityMonitoringRuleOptions and assigns it to the Options field.
func (o *SecurityMonitoringRuleUpdatePayload) SetOptions(v SecurityMonitoringRuleOptions) {
	o.Options = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetQueries() []SecurityMonitoringRuleQuery {
	if o == nil || o.Queries == nil {
		var ret []SecurityMonitoringRuleQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetQueriesOk() (*[]SecurityMonitoringRuleQuery, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []SecurityMonitoringRuleQuery and assigns it to the Queries field.
func (o *SecurityMonitoringRuleUpdatePayload) SetQueries(v []SecurityMonitoringRuleQuery) {
	o.Queries = v
}

// GetReferenceTables returns the ReferenceTables field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetReferenceTables() []SecurityMonitoringReferenceTable {
	if o == nil || o.ReferenceTables == nil {
		var ret []SecurityMonitoringReferenceTable
		return ret
	}
	return o.ReferenceTables
}

// GetReferenceTablesOk returns a tuple with the ReferenceTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetReferenceTablesOk() (*[]SecurityMonitoringReferenceTable, bool) {
	if o == nil || o.ReferenceTables == nil {
		return nil, false
	}
	return &o.ReferenceTables, true
}

// HasReferenceTables returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasReferenceTables() bool {
	return o != nil && o.ReferenceTables != nil
}

// SetReferenceTables gets a reference to the given []SecurityMonitoringReferenceTable and assigns it to the ReferenceTables field.
func (o *SecurityMonitoringRuleUpdatePayload) SetReferenceTables(v []SecurityMonitoringReferenceTable) {
	o.ReferenceTables = v
}

// GetSchedulingOptions returns the SchedulingOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringRuleUpdatePayload) GetSchedulingOptions() SecurityMonitoringSchedulingOptions {
	if o == nil || o.SchedulingOptions.Get() == nil {
		var ret SecurityMonitoringSchedulingOptions
		return ret
	}
	return *o.SchedulingOptions.Get()
}

// GetSchedulingOptionsOk returns a tuple with the SchedulingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringRuleUpdatePayload) GetSchedulingOptionsOk() (*SecurityMonitoringSchedulingOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.SchedulingOptions.Get(), o.SchedulingOptions.IsSet()
}

// HasSchedulingOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasSchedulingOptions() bool {
	return o != nil && o.SchedulingOptions.IsSet()
}

// SetSchedulingOptions gets a reference to the given NullableSecurityMonitoringSchedulingOptions and assigns it to the SchedulingOptions field.
func (o *SecurityMonitoringRuleUpdatePayload) SetSchedulingOptions(v SecurityMonitoringSchedulingOptions) {
	o.SchedulingOptions.Set(&v)
}

// SetSchedulingOptionsNil sets the value for SchedulingOptions to be an explicit nil.
func (o *SecurityMonitoringRuleUpdatePayload) SetSchedulingOptionsNil() {
	o.SchedulingOptions.Set(nil)
}

// UnsetSchedulingOptions ensures that no value is present for SchedulingOptions, not even an explicit nil.
func (o *SecurityMonitoringRuleUpdatePayload) UnsetSchedulingOptions() {
	o.SchedulingOptions.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringRuleUpdatePayload) SetTags(v []string) {
	o.Tags = v
}

// GetThirdPartyCases returns the ThirdPartyCases field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetThirdPartyCases() []SecurityMonitoringThirdPartyRuleCase {
	if o == nil || o.ThirdPartyCases == nil {
		var ret []SecurityMonitoringThirdPartyRuleCase
		return ret
	}
	return o.ThirdPartyCases
}

// GetThirdPartyCasesOk returns a tuple with the ThirdPartyCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetThirdPartyCasesOk() (*[]SecurityMonitoringThirdPartyRuleCase, bool) {
	if o == nil || o.ThirdPartyCases == nil {
		return nil, false
	}
	return &o.ThirdPartyCases, true
}

// HasThirdPartyCases returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasThirdPartyCases() bool {
	return o != nil && o.ThirdPartyCases != nil
}

// SetThirdPartyCases gets a reference to the given []SecurityMonitoringThirdPartyRuleCase and assigns it to the ThirdPartyCases field.
func (o *SecurityMonitoringRuleUpdatePayload) SetThirdPartyCases(v []SecurityMonitoringThirdPartyRuleCase) {
	o.ThirdPartyCases = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleUpdatePayload) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleUpdatePayload) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleUpdatePayload) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *SecurityMonitoringRuleUpdatePayload) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculatedFields"] = o.CalculatedFields
	}
	if o.Cases != nil {
		toSerialize["cases"] = o.Cases
	}
	if o.ComplianceSignalOptions != nil {
		toSerialize["complianceSignalOptions"] = o.ComplianceSignalOptions
	}
	if o.CustomMessage != nil {
		toSerialize["customMessage"] = o.CustomMessage
	}
	if o.CustomName != nil {
		toSerialize["customName"] = o.CustomName
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.GroupSignalsBy != nil {
		toSerialize["groupSignalsBy"] = o.GroupSignalsBy
	}
	if o.HasExtendedTitle != nil {
		toSerialize["hasExtendedTitle"] = o.HasExtendedTitle
	}
	if o.IsEnabled != nil {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.ReferenceTables != nil {
		toSerialize["referenceTables"] = o.ReferenceTables
	}
	if o.SchedulingOptions.IsSet() {
		toSerialize["schedulingOptions"] = o.SchedulingOptions.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ThirdPartyCases != nil {
		toSerialize["thirdPartyCases"] = o.ThirdPartyCases
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleUpdatePayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields        []CalculatedField                              `json:"calculatedFields,omitempty"`
		Cases                   []SecurityMonitoringRuleCase                   `json:"cases,omitempty"`
		ComplianceSignalOptions *CloudConfigurationRuleComplianceSignalOptions `json:"complianceSignalOptions,omitempty"`
		CustomMessage           *string                                        `json:"customMessage,omitempty"`
		CustomName              *string                                        `json:"customName,omitempty"`
		Filters                 []SecurityMonitoringFilter                     `json:"filters,omitempty"`
		GroupSignalsBy          []string                                       `json:"groupSignalsBy,omitempty"`
		HasExtendedTitle        *bool                                          `json:"hasExtendedTitle,omitempty"`
		IsEnabled               *bool                                          `json:"isEnabled,omitempty"`
		Message                 *string                                        `json:"message,omitempty"`
		Name                    *string                                        `json:"name,omitempty"`
		Options                 *SecurityMonitoringRuleOptions                 `json:"options,omitempty"`
		Queries                 []SecurityMonitoringRuleQuery                  `json:"queries,omitempty"`
		ReferenceTables         []SecurityMonitoringReferenceTable             `json:"referenceTables,omitempty"`
		SchedulingOptions       NullableSecurityMonitoringSchedulingOptions    `json:"schedulingOptions,omitempty"`
		Tags                    []string                                       `json:"tags,omitempty"`
		ThirdPartyCases         []SecurityMonitoringThirdPartyRuleCase         `json:"thirdPartyCases,omitempty"`
		Version                 *int32                                         `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"calculatedFields", "cases", "complianceSignalOptions", "customMessage", "customName", "filters", "groupSignalsBy", "hasExtendedTitle", "isEnabled", "message", "name", "options", "queries", "referenceTables", "schedulingOptions", "tags", "thirdPartyCases", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CalculatedFields = all.CalculatedFields
	o.Cases = all.Cases
	if all.ComplianceSignalOptions != nil && all.ComplianceSignalOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ComplianceSignalOptions = all.ComplianceSignalOptions
	o.CustomMessage = all.CustomMessage
	o.CustomName = all.CustomName
	o.Filters = all.Filters
	o.GroupSignalsBy = all.GroupSignalsBy
	o.HasExtendedTitle = all.HasExtendedTitle
	o.IsEnabled = all.IsEnabled
	o.Message = all.Message
	o.Name = all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	o.Queries = all.Queries
	o.ReferenceTables = all.ReferenceTables
	o.SchedulingOptions = all.SchedulingOptions
	o.Tags = all.Tags
	o.ThirdPartyCases = all.ThirdPartyCases
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
