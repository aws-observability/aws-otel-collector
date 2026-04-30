// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateCampaignRequestAttributes Attributes for updating a campaign.
type UpdateCampaignRequestAttributes struct {
	// The description of the campaign.
	Description *string `json:"description,omitempty"`
	// The due date of the campaign.
	DueDate *time.Time `json:"due_date,omitempty"`
	// Entity scope query to filter entities for this campaign.
	EntityScope *string `json:"entity_scope,omitempty"`
	// Guidance for the campaign.
	Guidance *string `json:"guidance,omitempty"`
	// The unique key for the campaign.
	Key *string `json:"key,omitempty"`
	// The name of the campaign.
	Name string `json:"name"`
	// The UUID of the campaign owner.
	OwnerId string `json:"owner_id"`
	// Array of rule IDs associated with this campaign.
	RuleIds []string `json:"rule_ids"`
	// The start date of the campaign.
	StartDate time.Time `json:"start_date"`
	// The status of the campaign.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateCampaignRequestAttributes instantiates a new UpdateCampaignRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateCampaignRequestAttributes(name string, ownerId string, ruleIds []string, startDate time.Time, status string) *UpdateCampaignRequestAttributes {
	this := UpdateCampaignRequestAttributes{}
	this.Name = name
	this.OwnerId = ownerId
	this.RuleIds = ruleIds
	this.StartDate = startDate
	this.Status = status
	return &this
}

// NewUpdateCampaignRequestAttributesWithDefaults instantiates a new UpdateCampaignRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateCampaignRequestAttributesWithDefaults() *UpdateCampaignRequestAttributes {
	this := UpdateCampaignRequestAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCampaignRequestAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCampaignRequestAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCampaignRequestAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *UpdateCampaignRequestAttributes) GetDueDate() time.Time {
	if o == nil || o.DueDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetDueDateOk() (*time.Time, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *UpdateCampaignRequestAttributes) HasDueDate() bool {
	return o != nil && o.DueDate != nil
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *UpdateCampaignRequestAttributes) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetEntityScope returns the EntityScope field value if set, zero value otherwise.
func (o *UpdateCampaignRequestAttributes) GetEntityScope() string {
	if o == nil || o.EntityScope == nil {
		var ret string
		return ret
	}
	return *o.EntityScope
}

// GetEntityScopeOk returns a tuple with the EntityScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetEntityScopeOk() (*string, bool) {
	if o == nil || o.EntityScope == nil {
		return nil, false
	}
	return o.EntityScope, true
}

// HasEntityScope returns a boolean if a field has been set.
func (o *UpdateCampaignRequestAttributes) HasEntityScope() bool {
	return o != nil && o.EntityScope != nil
}

// SetEntityScope gets a reference to the given string and assigns it to the EntityScope field.
func (o *UpdateCampaignRequestAttributes) SetEntityScope(v string) {
	o.EntityScope = &v
}

// GetGuidance returns the Guidance field value if set, zero value otherwise.
func (o *UpdateCampaignRequestAttributes) GetGuidance() string {
	if o == nil || o.Guidance == nil {
		var ret string
		return ret
	}
	return *o.Guidance
}

// GetGuidanceOk returns a tuple with the Guidance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetGuidanceOk() (*string, bool) {
	if o == nil || o.Guidance == nil {
		return nil, false
	}
	return o.Guidance, true
}

// HasGuidance returns a boolean if a field has been set.
func (o *UpdateCampaignRequestAttributes) HasGuidance() bool {
	return o != nil && o.Guidance != nil
}

// SetGuidance gets a reference to the given string and assigns it to the Guidance field.
func (o *UpdateCampaignRequestAttributes) SetGuidance(v string) {
	o.Guidance = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateCampaignRequestAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateCampaignRequestAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UpdateCampaignRequestAttributes) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value.
func (o *UpdateCampaignRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UpdateCampaignRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value.
func (o *UpdateCampaignRequestAttributes) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value.
func (o *UpdateCampaignRequestAttributes) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetRuleIds returns the RuleIds field value.
func (o *UpdateCampaignRequestAttributes) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetRuleIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleIds, true
}

// SetRuleIds sets field value.
func (o *UpdateCampaignRequestAttributes) SetRuleIds(v []string) {
	o.RuleIds = v
}

// GetStartDate returns the StartDate field value.
func (o *UpdateCampaignRequestAttributes) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value.
func (o *UpdateCampaignRequestAttributes) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetStatus returns the Status field value.
func (o *UpdateCampaignRequestAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignRequestAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *UpdateCampaignRequestAttributes) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateCampaignRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DueDate != nil {
		if o.DueDate.Nanosecond() == 0 {
			toSerialize["due_date"] = o.DueDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["due_date"] = o.DueDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EntityScope != nil {
		toSerialize["entity_scope"] = o.EntityScope
	}
	if o.Guidance != nil {
		toSerialize["guidance"] = o.Guidance
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	toSerialize["rule_ids"] = o.RuleIds
	if o.StartDate.Nanosecond() == 0 {
		toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateCampaignRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string    `json:"description,omitempty"`
		DueDate     *time.Time `json:"due_date,omitempty"`
		EntityScope *string    `json:"entity_scope,omitempty"`
		Guidance    *string    `json:"guidance,omitempty"`
		Key         *string    `json:"key,omitempty"`
		Name        *string    `json:"name"`
		OwnerId     *string    `json:"owner_id"`
		RuleIds     *[]string  `json:"rule_ids"`
		StartDate   *time.Time `json:"start_date"`
		Status      *string    `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OwnerId == nil {
		return fmt.Errorf("required field owner_id missing")
	}
	if all.RuleIds == nil {
		return fmt.Errorf("required field rule_ids missing")
	}
	if all.StartDate == nil {
		return fmt.Errorf("required field start_date missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "due_date", "entity_scope", "guidance", "key", "name", "owner_id", "rule_ids", "start_date", "status"})
	} else {
		return err
	}
	o.Description = all.Description
	o.DueDate = all.DueDate
	o.EntityScope = all.EntityScope
	o.Guidance = all.Guidance
	o.Key = all.Key
	o.Name = *all.Name
	o.OwnerId = *all.OwnerId
	o.RuleIds = *all.RuleIds
	o.StartDate = *all.StartDate
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
