// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DetailedFindingAttributes The JSON:API attributes of the detailed finding.
type DetailedFindingAttributes struct {
	// The evaluation of the finding.
	Evaluation *FindingEvaluation `json:"evaluation,omitempty"`
	// The date on which the evaluation for this finding changed (Unix ms).
	EvaluationChangedAt *int64 `json:"evaluation_changed_at,omitempty"`
	// The remediation message for this finding.
	Message *string `json:"message,omitempty"`
	// Information about the mute status of this finding.
	Mute *FindingMute `json:"mute,omitempty"`
	// The resource name of this finding.
	Resource *string `json:"resource,omitempty"`
	// The resource configuration for this finding.
	ResourceConfiguration interface{} `json:"resource_configuration,omitempty"`
	// The date on which the resource was discovered (Unix ms).
	ResourceDiscoveryDate *int64 `json:"resource_discovery_date,omitempty"`
	// The resource type of this finding.
	ResourceType *string `json:"resource_type,omitempty"`
	// The rule that triggered this finding.
	Rule *FindingRule `json:"rule,omitempty"`
	// The status of the finding.
	Status *FindingStatus `json:"status,omitempty"`
	// The tags associated with this finding.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDetailedFindingAttributes instantiates a new DetailedFindingAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDetailedFindingAttributes() *DetailedFindingAttributes {
	this := DetailedFindingAttributes{}
	return &this
}

// NewDetailedFindingAttributesWithDefaults instantiates a new DetailedFindingAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDetailedFindingAttributesWithDefaults() *DetailedFindingAttributes {
	this := DetailedFindingAttributes{}
	return &this
}

// GetEvaluation returns the Evaluation field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetEvaluation() FindingEvaluation {
	if o == nil || o.Evaluation == nil {
		var ret FindingEvaluation
		return ret
	}
	return *o.Evaluation
}

// GetEvaluationOk returns a tuple with the Evaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetEvaluationOk() (*FindingEvaluation, bool) {
	if o == nil || o.Evaluation == nil {
		return nil, false
	}
	return o.Evaluation, true
}

// HasEvaluation returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasEvaluation() bool {
	return o != nil && o.Evaluation != nil
}

// SetEvaluation gets a reference to the given FindingEvaluation and assigns it to the Evaluation field.
func (o *DetailedFindingAttributes) SetEvaluation(v FindingEvaluation) {
	o.Evaluation = &v
}

// GetEvaluationChangedAt returns the EvaluationChangedAt field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetEvaluationChangedAt() int64 {
	if o == nil || o.EvaluationChangedAt == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationChangedAt
}

// GetEvaluationChangedAtOk returns a tuple with the EvaluationChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetEvaluationChangedAtOk() (*int64, bool) {
	if o == nil || o.EvaluationChangedAt == nil {
		return nil, false
	}
	return o.EvaluationChangedAt, true
}

// HasEvaluationChangedAt returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasEvaluationChangedAt() bool {
	return o != nil && o.EvaluationChangedAt != nil
}

// SetEvaluationChangedAt gets a reference to the given int64 and assigns it to the EvaluationChangedAt field.
func (o *DetailedFindingAttributes) SetEvaluationChangedAt(v int64) {
	o.EvaluationChangedAt = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DetailedFindingAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetMute() FindingMute {
	if o == nil || o.Mute == nil {
		var ret FindingMute
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetMuteOk() (*FindingMute, bool) {
	if o == nil || o.Mute == nil {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasMute() bool {
	return o != nil && o.Mute != nil
}

// SetMute gets a reference to the given FindingMute and assigns it to the Mute field.
func (o *DetailedFindingAttributes) SetMute(v FindingMute) {
	o.Mute = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *DetailedFindingAttributes) SetResource(v string) {
	o.Resource = &v
}

// GetResourceConfiguration returns the ResourceConfiguration field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetResourceConfiguration() interface{} {
	if o == nil || o.ResourceConfiguration == nil {
		var ret interface{}
		return ret
	}
	return o.ResourceConfiguration
}

// GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetResourceConfigurationOk() (*interface{}, bool) {
	if o == nil || o.ResourceConfiguration == nil {
		return nil, false
	}
	return &o.ResourceConfiguration, true
}

// HasResourceConfiguration returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasResourceConfiguration() bool {
	return o != nil && o.ResourceConfiguration != nil
}

// SetResourceConfiguration gets a reference to the given interface{} and assigns it to the ResourceConfiguration field.
func (o *DetailedFindingAttributes) SetResourceConfiguration(v interface{}) {
	o.ResourceConfiguration = v
}

// GetResourceDiscoveryDate returns the ResourceDiscoveryDate field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetResourceDiscoveryDate() int64 {
	if o == nil || o.ResourceDiscoveryDate == nil {
		var ret int64
		return ret
	}
	return *o.ResourceDiscoveryDate
}

// GetResourceDiscoveryDateOk returns a tuple with the ResourceDiscoveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetResourceDiscoveryDateOk() (*int64, bool) {
	if o == nil || o.ResourceDiscoveryDate == nil {
		return nil, false
	}
	return o.ResourceDiscoveryDate, true
}

// HasResourceDiscoveryDate returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasResourceDiscoveryDate() bool {
	return o != nil && o.ResourceDiscoveryDate != nil
}

// SetResourceDiscoveryDate gets a reference to the given int64 and assigns it to the ResourceDiscoveryDate field.
func (o *DetailedFindingAttributes) SetResourceDiscoveryDate(v int64) {
	o.ResourceDiscoveryDate = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *DetailedFindingAttributes) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetRule() FindingRule {
	if o == nil || o.Rule == nil {
		var ret FindingRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetRuleOk() (*FindingRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given FindingRule and assigns it to the Rule field.
func (o *DetailedFindingAttributes) SetRule(v FindingRule) {
	o.Rule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetStatus() FindingStatus {
	if o == nil || o.Status == nil {
		var ret FindingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetStatusOk() (*FindingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FindingStatus and assigns it to the Status field.
func (o *DetailedFindingAttributes) SetStatus(v FindingStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DetailedFindingAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedFindingAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DetailedFindingAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DetailedFindingAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DetailedFindingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Evaluation != nil {
		toSerialize["evaluation"] = o.Evaluation
	}
	if o.EvaluationChangedAt != nil {
		toSerialize["evaluation_changed_at"] = o.EvaluationChangedAt
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Mute != nil {
		toSerialize["mute"] = o.Mute
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.ResourceConfiguration != nil {
		toSerialize["resource_configuration"] = o.ResourceConfiguration
	}
	if o.ResourceDiscoveryDate != nil {
		toSerialize["resource_discovery_date"] = o.ResourceDiscoveryDate
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DetailedFindingAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Evaluation            *FindingEvaluation `json:"evaluation,omitempty"`
		EvaluationChangedAt   *int64             `json:"evaluation_changed_at,omitempty"`
		Message               *string            `json:"message,omitempty"`
		Mute                  *FindingMute       `json:"mute,omitempty"`
		Resource              *string            `json:"resource,omitempty"`
		ResourceConfiguration interface{}        `json:"resource_configuration,omitempty"`
		ResourceDiscoveryDate *int64             `json:"resource_discovery_date,omitempty"`
		ResourceType          *string            `json:"resource_type,omitempty"`
		Rule                  *FindingRule       `json:"rule,omitempty"`
		Status                *FindingStatus     `json:"status,omitempty"`
		Tags                  []string           `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"evaluation", "evaluation_changed_at", "message", "mute", "resource", "resource_configuration", "resource_discovery_date", "resource_type", "rule", "status", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Evaluation != nil && !all.Evaluation.IsValid() {
		hasInvalidField = true
	} else {
		o.Evaluation = all.Evaluation
	}
	o.EvaluationChangedAt = all.EvaluationChangedAt
	o.Message = all.Message
	if all.Mute != nil && all.Mute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Mute = all.Mute
	o.Resource = all.Resource
	o.ResourceConfiguration = all.ResourceConfiguration
	o.ResourceDiscoveryDate = all.ResourceDiscoveryDate
	o.ResourceType = all.ResourceType
	if all.Rule != nil && all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = all.Rule
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
