// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteFindingResponseAttributes The JSON:API attributes of the finding.
type MuteFindingResponseAttributes struct {
	// The evaluation of the finding.
	Evaluation *FindingEvaluation `json:"evaluation,omitempty"`
	// The date on which the evaluation for this finding changed (Unix ms).
	EvaluationChangedAt *int64 `json:"evaluation_changed_at,omitempty"`
	// Information about the mute status of this finding.
	Mute *MuteFindingResponseProperties `json:"mute,omitempty"`
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
	AdditionalProperties map[string]interface{}
}

// NewMuteFindingResponseAttributes instantiates a new MuteFindingResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingResponseAttributes() *MuteFindingResponseAttributes {
	this := MuteFindingResponseAttributes{}
	return &this
}

// NewMuteFindingResponseAttributesWithDefaults instantiates a new MuteFindingResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingResponseAttributesWithDefaults() *MuteFindingResponseAttributes {
	this := MuteFindingResponseAttributes{}
	return &this
}

// GetEvaluation returns the Evaluation field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetEvaluation() FindingEvaluation {
	if o == nil || o.Evaluation == nil {
		var ret FindingEvaluation
		return ret
	}
	return *o.Evaluation
}

// GetEvaluationOk returns a tuple with the Evaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetEvaluationOk() (*FindingEvaluation, bool) {
	if o == nil || o.Evaluation == nil {
		return nil, false
	}
	return o.Evaluation, true
}

// HasEvaluation returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasEvaluation() bool {
	return o != nil && o.Evaluation != nil
}

// SetEvaluation gets a reference to the given FindingEvaluation and assigns it to the Evaluation field.
func (o *MuteFindingResponseAttributes) SetEvaluation(v FindingEvaluation) {
	o.Evaluation = &v
}

// GetEvaluationChangedAt returns the EvaluationChangedAt field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetEvaluationChangedAt() int64 {
	if o == nil || o.EvaluationChangedAt == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationChangedAt
}

// GetEvaluationChangedAtOk returns a tuple with the EvaluationChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetEvaluationChangedAtOk() (*int64, bool) {
	if o == nil || o.EvaluationChangedAt == nil {
		return nil, false
	}
	return o.EvaluationChangedAt, true
}

// HasEvaluationChangedAt returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasEvaluationChangedAt() bool {
	return o != nil && o.EvaluationChangedAt != nil
}

// SetEvaluationChangedAt gets a reference to the given int64 and assigns it to the EvaluationChangedAt field.
func (o *MuteFindingResponseAttributes) SetEvaluationChangedAt(v int64) {
	o.EvaluationChangedAt = &v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetMute() MuteFindingResponseProperties {
	if o == nil || o.Mute == nil {
		var ret MuteFindingResponseProperties
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetMuteOk() (*MuteFindingResponseProperties, bool) {
	if o == nil || o.Mute == nil {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasMute() bool {
	return o != nil && o.Mute != nil
}

// SetMute gets a reference to the given MuteFindingResponseProperties and assigns it to the Mute field.
func (o *MuteFindingResponseAttributes) SetMute(v MuteFindingResponseProperties) {
	o.Mute = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *MuteFindingResponseAttributes) SetResource(v string) {
	o.Resource = &v
}

// GetResourceDiscoveryDate returns the ResourceDiscoveryDate field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetResourceDiscoveryDate() int64 {
	if o == nil || o.ResourceDiscoveryDate == nil {
		var ret int64
		return ret
	}
	return *o.ResourceDiscoveryDate
}

// GetResourceDiscoveryDateOk returns a tuple with the ResourceDiscoveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetResourceDiscoveryDateOk() (*int64, bool) {
	if o == nil || o.ResourceDiscoveryDate == nil {
		return nil, false
	}
	return o.ResourceDiscoveryDate, true
}

// HasResourceDiscoveryDate returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasResourceDiscoveryDate() bool {
	return o != nil && o.ResourceDiscoveryDate != nil
}

// SetResourceDiscoveryDate gets a reference to the given int64 and assigns it to the ResourceDiscoveryDate field.
func (o *MuteFindingResponseAttributes) SetResourceDiscoveryDate(v int64) {
	o.ResourceDiscoveryDate = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *MuteFindingResponseAttributes) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetRule() FindingRule {
	if o == nil || o.Rule == nil {
		var ret FindingRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetRuleOk() (*FindingRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given FindingRule and assigns it to the Rule field.
func (o *MuteFindingResponseAttributes) SetRule(v FindingRule) {
	o.Rule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetStatus() FindingStatus {
	if o == nil || o.Status == nil {
		var ret FindingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetStatusOk() (*FindingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given FindingStatus and assigns it to the Status field.
func (o *MuteFindingResponseAttributes) SetStatus(v FindingStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MuteFindingResponseAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteFindingResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MuteFindingResponseAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MuteFindingResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
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
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Evaluation            *FindingEvaluation             `json:"evaluation,omitempty"`
		EvaluationChangedAt   *int64                         `json:"evaluation_changed_at,omitempty"`
		Mute                  *MuteFindingResponseProperties `json:"mute,omitempty"`
		Resource              *string                        `json:"resource,omitempty"`
		ResourceDiscoveryDate *int64                         `json:"resource_discovery_date,omitempty"`
		ResourceType          *string                        `json:"resource_type,omitempty"`
		Rule                  *FindingRule                   `json:"rule,omitempty"`
		Status                *FindingStatus                 `json:"status,omitempty"`
		Tags                  []string                       `json:"tags,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
