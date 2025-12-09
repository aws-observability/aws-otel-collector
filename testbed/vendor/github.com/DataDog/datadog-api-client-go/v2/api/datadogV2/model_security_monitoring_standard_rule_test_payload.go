// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringStandardRuleTestPayload The payload of a rule to test
type SecurityMonitoringStandardRuleTestPayload struct {
	// Calculated fields. Only allowed for scheduled rules - in other words, when schedulingOptions is also defined.
	CalculatedFields []CalculatedField `json:"calculatedFields,omitempty"`
	// Cases for generating signals.
	Cases []SecurityMonitoringRuleCaseCreate `json:"cases"`
	// Additional queries to filter matched events before they are processed. This field is deprecated for log detection, signal correlation, and workload security rules.
	Filters []SecurityMonitoringFilter `json:"filters,omitempty"`
	// Additional grouping to perform on top of the existing groups in the query section. Must be a subset of the existing groups.
	GroupSignalsBy []string `json:"groupSignalsBy,omitempty"`
	// Whether the notifications include the triggering group-by values in their title.
	HasExtendedTitle *bool `json:"hasExtendedTitle,omitempty"`
	// Whether the rule is enabled.
	IsEnabled bool `json:"isEnabled"`
	// Message for generated signals.
	Message string `json:"message"`
	// The name of the rule.
	Name string `json:"name"`
	// Options.
	Options SecurityMonitoringRuleOptions `json:"options"`
	// Queries for selecting logs which are part of the rule.
	Queries []SecurityMonitoringStandardRuleQuery `json:"queries"`
	// Reference tables for the rule.
	ReferenceTables []SecurityMonitoringReferenceTable `json:"referenceTables,omitempty"`
	// Options for scheduled rules. When this field is present, the rule runs based on the schedule. When absent, it runs real-time on ingested logs.
	SchedulingOptions NullableSecurityMonitoringSchedulingOptions `json:"schedulingOptions,omitempty"`
	// Tags for generated signals.
	Tags []string `json:"tags,omitempty"`
	// Cases for generating signals from third-party rules. Only available for third-party rules.
	ThirdPartyCases []SecurityMonitoringThirdPartyRuleCaseCreate `json:"thirdPartyCases,omitempty"`
	// The rule type.
	Type *SecurityMonitoringRuleTypeTest `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringStandardRuleTestPayload instantiates a new SecurityMonitoringStandardRuleTestPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringStandardRuleTestPayload(cases []SecurityMonitoringRuleCaseCreate, isEnabled bool, message string, name string, options SecurityMonitoringRuleOptions, queries []SecurityMonitoringStandardRuleQuery) *SecurityMonitoringStandardRuleTestPayload {
	this := SecurityMonitoringStandardRuleTestPayload{}
	this.Cases = cases
	this.IsEnabled = isEnabled
	this.Message = message
	this.Name = name
	this.Options = options
	this.Queries = queries
	return &this
}

// NewSecurityMonitoringStandardRuleTestPayloadWithDefaults instantiates a new SecurityMonitoringStandardRuleTestPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringStandardRuleTestPayloadWithDefaults() *SecurityMonitoringStandardRuleTestPayload {
	this := SecurityMonitoringStandardRuleTestPayload{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetCalculatedFields() []CalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []CalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetCalculatedFieldsOk() (*[]CalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []CalculatedField and assigns it to the CalculatedFields field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetCalculatedFields(v []CalculatedField) {
	o.CalculatedFields = v
}

// GetCases returns the Cases field value.
func (o *SecurityMonitoringStandardRuleTestPayload) GetCases() []SecurityMonitoringRuleCaseCreate {
	if o == nil {
		var ret []SecurityMonitoringRuleCaseCreate
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetCasesOk() (*[]SecurityMonitoringRuleCaseCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cases, true
}

// SetCases sets field value.
func (o *SecurityMonitoringStandardRuleTestPayload) SetCases(v []SecurityMonitoringRuleCaseCreate) {
	o.Cases = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetFilters() []SecurityMonitoringFilter {
	if o == nil || o.Filters == nil {
		var ret []SecurityMonitoringFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetFiltersOk() (*[]SecurityMonitoringFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []SecurityMonitoringFilter and assigns it to the Filters field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetFilters(v []SecurityMonitoringFilter) {
	o.Filters = v
}

// GetGroupSignalsBy returns the GroupSignalsBy field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetGroupSignalsBy() []string {
	if o == nil || o.GroupSignalsBy == nil {
		var ret []string
		return ret
	}
	return o.GroupSignalsBy
}

// GetGroupSignalsByOk returns a tuple with the GroupSignalsBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetGroupSignalsByOk() (*[]string, bool) {
	if o == nil || o.GroupSignalsBy == nil {
		return nil, false
	}
	return &o.GroupSignalsBy, true
}

// HasGroupSignalsBy returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasGroupSignalsBy() bool {
	return o != nil && o.GroupSignalsBy != nil
}

// SetGroupSignalsBy gets a reference to the given []string and assigns it to the GroupSignalsBy field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetGroupSignalsBy(v []string) {
	o.GroupSignalsBy = v
}

// GetHasExtendedTitle returns the HasExtendedTitle field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetHasExtendedTitle() bool {
	if o == nil || o.HasExtendedTitle == nil {
		var ret bool
		return ret
	}
	return *o.HasExtendedTitle
}

// GetHasExtendedTitleOk returns a tuple with the HasExtendedTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetHasExtendedTitleOk() (*bool, bool) {
	if o == nil || o.HasExtendedTitle == nil {
		return nil, false
	}
	return o.HasExtendedTitle, true
}

// HasHasExtendedTitle returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasHasExtendedTitle() bool {
	return o != nil && o.HasExtendedTitle != nil
}

// SetHasExtendedTitle gets a reference to the given bool and assigns it to the HasExtendedTitle field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetHasExtendedTitle(v bool) {
	o.HasExtendedTitle = &v
}

// GetIsEnabled returns the IsEnabled field value.
func (o *SecurityMonitoringStandardRuleTestPayload) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *SecurityMonitoringStandardRuleTestPayload) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetMessage returns the Message field value.
func (o *SecurityMonitoringStandardRuleTestPayload) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *SecurityMonitoringStandardRuleTestPayload) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringStandardRuleTestPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringStandardRuleTestPayload) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *SecurityMonitoringStandardRuleTestPayload) GetOptions() SecurityMonitoringRuleOptions {
	if o == nil {
		var ret SecurityMonitoringRuleOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *SecurityMonitoringStandardRuleTestPayload) SetOptions(v SecurityMonitoringRuleOptions) {
	o.Options = v
}

// GetQueries returns the Queries field value.
func (o *SecurityMonitoringStandardRuleTestPayload) GetQueries() []SecurityMonitoringStandardRuleQuery {
	if o == nil {
		var ret []SecurityMonitoringStandardRuleQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetQueriesOk() (*[]SecurityMonitoringStandardRuleQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *SecurityMonitoringStandardRuleTestPayload) SetQueries(v []SecurityMonitoringStandardRuleQuery) {
	o.Queries = v
}

// GetReferenceTables returns the ReferenceTables field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetReferenceTables() []SecurityMonitoringReferenceTable {
	if o == nil || o.ReferenceTables == nil {
		var ret []SecurityMonitoringReferenceTable
		return ret
	}
	return o.ReferenceTables
}

// GetReferenceTablesOk returns a tuple with the ReferenceTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetReferenceTablesOk() (*[]SecurityMonitoringReferenceTable, bool) {
	if o == nil || o.ReferenceTables == nil {
		return nil, false
	}
	return &o.ReferenceTables, true
}

// HasReferenceTables returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasReferenceTables() bool {
	return o != nil && o.ReferenceTables != nil
}

// SetReferenceTables gets a reference to the given []SecurityMonitoringReferenceTable and assigns it to the ReferenceTables field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetReferenceTables(v []SecurityMonitoringReferenceTable) {
	o.ReferenceTables = v
}

// GetSchedulingOptions returns the SchedulingOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringStandardRuleTestPayload) GetSchedulingOptions() SecurityMonitoringSchedulingOptions {
	if o == nil || o.SchedulingOptions.Get() == nil {
		var ret SecurityMonitoringSchedulingOptions
		return ret
	}
	return *o.SchedulingOptions.Get()
}

// GetSchedulingOptionsOk returns a tuple with the SchedulingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringStandardRuleTestPayload) GetSchedulingOptionsOk() (*SecurityMonitoringSchedulingOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.SchedulingOptions.Get(), o.SchedulingOptions.IsSet()
}

// HasSchedulingOptions returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasSchedulingOptions() bool {
	return o != nil && o.SchedulingOptions.IsSet()
}

// SetSchedulingOptions gets a reference to the given NullableSecurityMonitoringSchedulingOptions and assigns it to the SchedulingOptions field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetSchedulingOptions(v SecurityMonitoringSchedulingOptions) {
	o.SchedulingOptions.Set(&v)
}

// SetSchedulingOptionsNil sets the value for SchedulingOptions to be an explicit nil.
func (o *SecurityMonitoringStandardRuleTestPayload) SetSchedulingOptionsNil() {
	o.SchedulingOptions.Set(nil)
}

// UnsetSchedulingOptions ensures that no value is present for SchedulingOptions, not even an explicit nil.
func (o *SecurityMonitoringStandardRuleTestPayload) UnsetSchedulingOptions() {
	o.SchedulingOptions.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetTags(v []string) {
	o.Tags = v
}

// GetThirdPartyCases returns the ThirdPartyCases field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetThirdPartyCases() []SecurityMonitoringThirdPartyRuleCaseCreate {
	if o == nil || o.ThirdPartyCases == nil {
		var ret []SecurityMonitoringThirdPartyRuleCaseCreate
		return ret
	}
	return o.ThirdPartyCases
}

// GetThirdPartyCasesOk returns a tuple with the ThirdPartyCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetThirdPartyCasesOk() (*[]SecurityMonitoringThirdPartyRuleCaseCreate, bool) {
	if o == nil || o.ThirdPartyCases == nil {
		return nil, false
	}
	return &o.ThirdPartyCases, true
}

// HasThirdPartyCases returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasThirdPartyCases() bool {
	return o != nil && o.ThirdPartyCases != nil
}

// SetThirdPartyCases gets a reference to the given []SecurityMonitoringThirdPartyRuleCaseCreate and assigns it to the ThirdPartyCases field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetThirdPartyCases(v []SecurityMonitoringThirdPartyRuleCaseCreate) {
	o.ThirdPartyCases = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleTestPayload) GetType() SecurityMonitoringRuleTypeTest {
	if o == nil || o.Type == nil {
		var ret SecurityMonitoringRuleTypeTest
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) GetTypeOk() (*SecurityMonitoringRuleTypeTest, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleTestPayload) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SecurityMonitoringRuleTypeTest and assigns it to the Type field.
func (o *SecurityMonitoringStandardRuleTestPayload) SetType(v SecurityMonitoringRuleTypeTest) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringStandardRuleTestPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculatedFields"] = o.CalculatedFields
	}
	toSerialize["cases"] = o.Cases
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.GroupSignalsBy != nil {
		toSerialize["groupSignalsBy"] = o.GroupSignalsBy
	}
	if o.HasExtendedTitle != nil {
		toSerialize["hasExtendedTitle"] = o.HasExtendedTitle
	}
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["queries"] = o.Queries
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringStandardRuleTestPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields  []CalculatedField                            `json:"calculatedFields,omitempty"`
		Cases             *[]SecurityMonitoringRuleCaseCreate          `json:"cases"`
		Filters           []SecurityMonitoringFilter                   `json:"filters,omitempty"`
		GroupSignalsBy    []string                                     `json:"groupSignalsBy,omitempty"`
		HasExtendedTitle  *bool                                        `json:"hasExtendedTitle,omitempty"`
		IsEnabled         *bool                                        `json:"isEnabled"`
		Message           *string                                      `json:"message"`
		Name              *string                                      `json:"name"`
		Options           *SecurityMonitoringRuleOptions               `json:"options"`
		Queries           *[]SecurityMonitoringStandardRuleQuery       `json:"queries"`
		ReferenceTables   []SecurityMonitoringReferenceTable           `json:"referenceTables,omitempty"`
		SchedulingOptions NullableSecurityMonitoringSchedulingOptions  `json:"schedulingOptions,omitempty"`
		Tags              []string                                     `json:"tags,omitempty"`
		ThirdPartyCases   []SecurityMonitoringThirdPartyRuleCaseCreate `json:"thirdPartyCases,omitempty"`
		Type              *SecurityMonitoringRuleTypeTest              `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cases == nil {
		return fmt.Errorf("required field cases missing")
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field isEnabled missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"calculatedFields", "cases", "filters", "groupSignalsBy", "hasExtendedTitle", "isEnabled", "message", "name", "options", "queries", "referenceTables", "schedulingOptions", "tags", "thirdPartyCases", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CalculatedFields = all.CalculatedFields
	o.Cases = *all.Cases
	o.Filters = all.Filters
	o.GroupSignalsBy = all.GroupSignalsBy
	o.HasExtendedTitle = all.HasExtendedTitle
	o.IsEnabled = *all.IsEnabled
	o.Message = *all.Message
	o.Name = *all.Name
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options
	o.Queries = *all.Queries
	o.ReferenceTables = all.ReferenceTables
	o.SchedulingOptions = all.SchedulingOptions
	o.Tags = all.Tags
	o.ThirdPartyCases = all.ThirdPartyCases
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
