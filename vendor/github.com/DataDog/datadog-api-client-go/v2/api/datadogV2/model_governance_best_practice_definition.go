// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceBestPracticeDefinition The best practice associated with an insight. Populated with the first active best practice
// matched to the insight; `null` when no best practice is attached.
type GovernanceBestPracticeDefinition struct {
	// The value driver the best practice is grouped under, such as `access_governance`,
	// `security`, `compliance`, or `operational_hygiene`.
	Category string `json:"category"`
	// A relative link to the configuration page where the best practice can be acted upon.
	DeepLink string `json:"deep_link"`
	// The full rationale and guidance for the best practice.
	Description string `json:"description"`
	// An optional association to a control's detection type. `null` when not associated with a control.
	DetectionType datadog.NullableString `json:"detection_type,omitempty"`
	// The unique identifier of the best practice.
	Id string `json:"id"`
	// The expected impact of following the best practice.
	Impact string `json:"impact"`
	// A priority hint for ordering best practices by expected impact. Lower values indicate
	// higher priority.
	ImpactHint int64 `json:"impact_hint"`
	// The permissions required for the user to act on the best practice.
	Permissions []string `json:"permissions"`
	// Whether the best practice is currently `active` or `deprecated`.
	Status string `json:"status"`
	// A one-line explanation of why this best practice matters.
	Summary string `json:"summary"`
	// A short, human-readable name for the best practice.
	Title string `json:"title"`
	// The condition that surfaces the best practice. For an `insight` trigger, the insight
	// slug; for a `static` trigger, a descriptive condition key.
	TriggerCondition string `json:"trigger_condition"`
	// How the best practice is surfaced. `insight` ties it to an insight; `static` surfaces it
	// unless its condition is met.
	TriggerType string `json:"trigger_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceBestPracticeDefinition instantiates a new GovernanceBestPracticeDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceBestPracticeDefinition(category string, deepLink string, description string, id string, impact string, impactHint int64, permissions []string, status string, summary string, title string, triggerCondition string, triggerType string) *GovernanceBestPracticeDefinition {
	this := GovernanceBestPracticeDefinition{}
	this.Category = category
	this.DeepLink = deepLink
	this.Description = description
	this.Id = id
	this.Impact = impact
	this.ImpactHint = impactHint
	this.Permissions = permissions
	this.Status = status
	this.Summary = summary
	this.Title = title
	this.TriggerCondition = triggerCondition
	this.TriggerType = triggerType
	return &this
}

// NewGovernanceBestPracticeDefinitionWithDefaults instantiates a new GovernanceBestPracticeDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceBestPracticeDefinitionWithDefaults() *GovernanceBestPracticeDefinition {
	this := GovernanceBestPracticeDefinition{}
	return &this
}

// GetCategory returns the Category field value.
func (o *GovernanceBestPracticeDefinition) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *GovernanceBestPracticeDefinition) SetCategory(v string) {
	o.Category = v
}

// GetDeepLink returns the DeepLink field value.
func (o *GovernanceBestPracticeDefinition) GetDeepLink() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeepLink
}

// GetDeepLinkOk returns a tuple with the DeepLink field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetDeepLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeepLink, true
}

// SetDeepLink sets field value.
func (o *GovernanceBestPracticeDefinition) SetDeepLink(v string) {
	o.DeepLink = v
}

// GetDescription returns the Description field value.
func (o *GovernanceBestPracticeDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceBestPracticeDefinition) SetDescription(v string) {
	o.Description = v
}

// GetDetectionType returns the DetectionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GovernanceBestPracticeDefinition) GetDetectionType() string {
	if o == nil || o.DetectionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DetectionType.Get()
}

// GetDetectionTypeOk returns a tuple with the DetectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceBestPracticeDefinition) GetDetectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetectionType.Get(), o.DetectionType.IsSet()
}

// HasDetectionType returns a boolean if a field has been set.
func (o *GovernanceBestPracticeDefinition) HasDetectionType() bool {
	return o != nil && o.DetectionType.IsSet()
}

// SetDetectionType gets a reference to the given datadog.NullableString and assigns it to the DetectionType field.
func (o *GovernanceBestPracticeDefinition) SetDetectionType(v string) {
	o.DetectionType.Set(&v)
}

// SetDetectionTypeNil sets the value for DetectionType to be an explicit nil.
func (o *GovernanceBestPracticeDefinition) SetDetectionTypeNil() {
	o.DetectionType.Set(nil)
}

// UnsetDetectionType ensures that no value is present for DetectionType, not even an explicit nil.
func (o *GovernanceBestPracticeDefinition) UnsetDetectionType() {
	o.DetectionType.Unset()
}

// GetId returns the Id field value.
func (o *GovernanceBestPracticeDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GovernanceBestPracticeDefinition) SetId(v string) {
	o.Id = v
}

// GetImpact returns the Impact field value.
func (o *GovernanceBestPracticeDefinition) GetImpact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value.
func (o *GovernanceBestPracticeDefinition) SetImpact(v string) {
	o.Impact = v
}

// GetImpactHint returns the ImpactHint field value.
func (o *GovernanceBestPracticeDefinition) GetImpactHint() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ImpactHint
}

// GetImpactHintOk returns a tuple with the ImpactHint field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetImpactHintOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactHint, true
}

// SetImpactHint sets field value.
func (o *GovernanceBestPracticeDefinition) SetImpactHint(v int64) {
	o.ImpactHint = v
}

// GetPermissions returns the Permissions field value.
func (o *GovernanceBestPracticeDefinition) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetPermissionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *GovernanceBestPracticeDefinition) SetPermissions(v []string) {
	o.Permissions = v
}

// GetStatus returns the Status field value.
func (o *GovernanceBestPracticeDefinition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *GovernanceBestPracticeDefinition) SetStatus(v string) {
	o.Status = v
}

// GetSummary returns the Summary field value.
func (o *GovernanceBestPracticeDefinition) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *GovernanceBestPracticeDefinition) SetSummary(v string) {
	o.Summary = v
}

// GetTitle returns the Title field value.
func (o *GovernanceBestPracticeDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *GovernanceBestPracticeDefinition) SetTitle(v string) {
	o.Title = v
}

// GetTriggerCondition returns the TriggerCondition field value.
func (o *GovernanceBestPracticeDefinition) GetTriggerCondition() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TriggerCondition
}

// GetTriggerConditionOk returns a tuple with the TriggerCondition field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetTriggerConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerCondition, true
}

// SetTriggerCondition sets field value.
func (o *GovernanceBestPracticeDefinition) SetTriggerCondition(v string) {
	o.TriggerCondition = v
}

// GetTriggerType returns the TriggerType field value.
func (o *GovernanceBestPracticeDefinition) GetTriggerType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *GovernanceBestPracticeDefinition) GetTriggerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value.
func (o *GovernanceBestPracticeDefinition) SetTriggerType(v string) {
	o.TriggerType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceBestPracticeDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["deep_link"] = o.DeepLink
	toSerialize["description"] = o.Description
	if o.DetectionType.IsSet() {
		toSerialize["detection_type"] = o.DetectionType.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["impact"] = o.Impact
	toSerialize["impact_hint"] = o.ImpactHint
	toSerialize["permissions"] = o.Permissions
	toSerialize["status"] = o.Status
	toSerialize["summary"] = o.Summary
	toSerialize["title"] = o.Title
	toSerialize["trigger_condition"] = o.TriggerCondition
	toSerialize["trigger_type"] = o.TriggerType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceBestPracticeDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category         *string                `json:"category"`
		DeepLink         *string                `json:"deep_link"`
		Description      *string                `json:"description"`
		DetectionType    datadog.NullableString `json:"detection_type,omitempty"`
		Id               *string                `json:"id"`
		Impact           *string                `json:"impact"`
		ImpactHint       *int64                 `json:"impact_hint"`
		Permissions      *[]string              `json:"permissions"`
		Status           *string                `json:"status"`
		Summary          *string                `json:"summary"`
		Title            *string                `json:"title"`
		TriggerCondition *string                `json:"trigger_condition"`
		TriggerType      *string                `json:"trigger_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.DeepLink == nil {
		return fmt.Errorf("required field deep_link missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Impact == nil {
		return fmt.Errorf("required field impact missing")
	}
	if all.ImpactHint == nil {
		return fmt.Errorf("required field impact_hint missing")
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.TriggerCondition == nil {
		return fmt.Errorf("required field trigger_condition missing")
	}
	if all.TriggerType == nil {
		return fmt.Errorf("required field trigger_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "deep_link", "description", "detection_type", "id", "impact", "impact_hint", "permissions", "status", "summary", "title", "trigger_condition", "trigger_type"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.DeepLink = *all.DeepLink
	o.Description = *all.Description
	o.DetectionType = all.DetectionType
	o.Id = *all.Id
	o.Impact = *all.Impact
	o.ImpactHint = *all.ImpactHint
	o.Permissions = *all.Permissions
	o.Status = *all.Status
	o.Summary = *all.Summary
	o.Title = *all.Title
	o.TriggerCondition = *all.TriggerCondition
	o.TriggerType = *all.TriggerType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
