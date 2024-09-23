// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingAttributes The JSON:API attributes of the finding.
type FindingAttributes struct {
	// The evaluation of the finding.
	Evaluation *FindingEvaluation `json:"evaluation,omitempty"`
	// The date on which the evaluation for this finding changed (Unix ms).
	EvaluationChangedAt *int64 `json:"evaluation_changed_at,omitempty"`
	// Information about the mute status of this finding.
	Mute *FindingMute `json:"mute,omitempty"`
	// The resource name of this finding.
	Resource *string `json:"resource,omitempty"`
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

// NewFindingAttributes instantiates a new FindingAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingAttributes() *FindingAttributes {
	this := FindingAttributes{}
	return &this
}

// NewFindingAttributesWithDefaults instantiates a new FindingAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingAttributesWithDefaults() *FindingAttributes {
	this := FindingAttributes{}
	return &this
}

// GetEvaluation returns the Evaluation field value if set, zero value otherwise.
func (o *FindingAttributes) GetEvaluation() FindingEvaluation {
	if o == nil || o.Evaluation == nil {
		var ret FindingEvaluation
		return ret
	}
	return *o.Evaluation
}

// GetEvaluationOk returns a tuple with the Evaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetEvaluationOk() (*FindingEvaluation, bool) {
	if o == nil || o.Evaluation == nil {
		return nil, false
	}
	return o.Evaluation, true
}

// HasEvaluation returns a boolean if a field has been set.
func (o *FindingAttributes) HasEvaluation() bool {
	return o != nil && o.Evaluation != nil
}

// SetEvaluation gets a reference to the given FindingEvaluation and assigns it to the Evaluation field.
func (o *FindingAttributes) SetEvaluation(v FindingEvaluation) {
	o.Evaluation = &v
}

// GetEvaluationChangedAt returns the EvaluationChangedAt field value if set, zero value otherwise.
func (o *FindingAttributes) GetEvaluationChangedAt() int64 {
	if o == nil || o.EvaluationChangedAt == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationChangedAt
}

// GetEvaluationChangedAtOk returns a tuple with the EvaluationChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetEvaluationChangedAtOk() (*int64, bool) {
	if o == nil || o.EvaluationChangedAt == nil {
		return nil, false
	}
	return o.EvaluationChangedAt, true
}

// HasEvaluationChangedAt returns a boolean if a field has been set.
func (o *FindingAttributes) HasEvaluationChangedAt() bool {
	return o != nil && o.EvaluationChangedAt != nil
}

// SetEvaluationChangedAt gets a reference to the given int64 and assigns it to the EvaluationChangedAt field.
func (o *FindingAttributes) SetEvaluationChangedAt(v int64) {
	o.EvaluationChangedAt = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *FindingAttributes) GetMute() FindingMute {
	if o == nil || o.Mute == nil {
		var ret FindingMute
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetMuteOk() (*FindingMute, bool) {
	if o == nil || o.Mute == nil {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *FindingAttributes) HasMute() bool {
	return o != nil && o.Mute != nil
}

// SetMute gets a reference to the given FindingMute and assigns it to the Mute field.
func (o *FindingAttributes) SetMute(v FindingMute) {
	o.Mute = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *FindingAttributes) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *FindingAttributes) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *FindingAttributes) SetResource(v string) {
	o.Resource = &v
}

// GetResourceDiscoveryDate returns the ResourceDiscoveryDate field value if set, zero value otherwise.
func (o *FindingAttributes) GetResourceDiscoveryDate() int64 {
	if o == nil || o.ResourceDiscoveryDate == nil {
		var ret int64
		return ret
	}
	return *o.ResourceDiscoveryDate
}

// GetResourceDiscoveryDateOk returns a tuple with the ResourceDiscoveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetResourceDiscoveryDateOk() (*int64, bool) {
	if o == nil || o.ResourceDiscoveryDate == nil {
		return nil, false
	}
	return o.ResourceDiscoveryDate, true
}

// HasResourceDiscoveryDate returns a boolean if a field has been set.
func (o *FindingAttributes) HasResourceDiscoveryDate() bool {
	return o != nil && o.ResourceDiscoveryDate != nil
}

// SetResourceDiscoveryDate gets a reference to the given int64 and assigns it to the ResourceDiscoveryDate field.
func (o *FindingAttributes) SetResourceDiscoveryDate(v int64) {
	o.ResourceDiscoveryDate = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *FindingAttributes) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *FindingAttributes) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *FindingAttributes) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FindingAttributes) GetRule() FindingRule {
	if o == nil || o.Rule == nil {
		var ret FindingRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetRuleOk() (*FindingRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FindingAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given FindingRule and assigns it to the Rule field.
func (o *FindingAttributes) SetRule(v FindingRule) {
	o.Rule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FindingAttributes) GetStatus() FindingStatus {
	if o == nil || o.Status == nil {
		var ret FindingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetStatusOk() (*FindingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FindingAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FindingStatus and assigns it to the Status field.
func (o *FindingAttributes) SetStatus(v FindingStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FindingAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FindingAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FindingAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Mute != nil {
		toSerialize["mute"] = o.Mute
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
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
func (o *FindingAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Evaluation            *FindingEvaluation `json:"evaluation,omitempty"`
		EvaluationChangedAt   *int64             `json:"evaluation_changed_at,omitempty"`
		Mute                  *FindingMute       `json:"mute,omitempty"`
		Resource              *string            `json:"resource,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"evaluation", "evaluation_changed_at", "mute", "resource", "resource_discovery_date", "resource_type", "rule", "status", "tags"})
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
	if all.Mute != nil && all.Mute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Mute = all.Mute
	o.Resource = all.Resource
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
