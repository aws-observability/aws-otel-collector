// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CampaignResponseAttributes Campaign attributes.
type CampaignResponseAttributes struct {
	// Creation time of the campaign.
	CreatedAt time.Time `json:"created_at"`
	// The description of the campaign.
	Description *string `json:"description,omitempty"`
	// The due date of the campaign.
	DueDate *time.Time `json:"due_date,omitempty"`
	// Entity scope query to filter entities for this campaign.
	EntityScope *string `json:"entity_scope,omitempty"`
	// Guidance for the campaign.
	Guidance *string `json:"guidance,omitempty"`
	// The unique key for the campaign.
	Key string `json:"key"`
	// Time of last campaign modification.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the campaign.
	Name string `json:"name"`
	// The UUID of the campaign owner.
	Owner string `json:"owner"`
	// The start date of the campaign.
	StartDate time.Time `json:"start_date"`
	// The status of the campaign.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCampaignResponseAttributes instantiates a new CampaignResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCampaignResponseAttributes(createdAt time.Time, key string, modifiedAt time.Time, name string, owner string, startDate time.Time, status string) *CampaignResponseAttributes {
	this := CampaignResponseAttributes{}
	this.CreatedAt = createdAt
	this.Key = key
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.Owner = owner
	this.StartDate = startDate
	this.Status = status
	return &this
}

// NewCampaignResponseAttributesWithDefaults instantiates a new CampaignResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCampaignResponseAttributesWithDefaults() *CampaignResponseAttributes {
	this := CampaignResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CampaignResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CampaignResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignResponseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignResponseAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignResponseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *CampaignResponseAttributes) GetDueDate() time.Time {
	if o == nil || o.DueDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetDueDateOk() (*time.Time, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *CampaignResponseAttributes) HasDueDate() bool {
	return o != nil && o.DueDate != nil
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *CampaignResponseAttributes) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetEntityScope returns the EntityScope field value if set, zero value otherwise.
func (o *CampaignResponseAttributes) GetEntityScope() string {
	if o == nil || o.EntityScope == nil {
		var ret string
		return ret
	}
	return *o.EntityScope
}

// GetEntityScopeOk returns a tuple with the EntityScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetEntityScopeOk() (*string, bool) {
	if o == nil || o.EntityScope == nil {
		return nil, false
	}
	return o.EntityScope, true
}

// HasEntityScope returns a boolean if a field has been set.
func (o *CampaignResponseAttributes) HasEntityScope() bool {
	return o != nil && o.EntityScope != nil
}

// SetEntityScope gets a reference to the given string and assigns it to the EntityScope field.
func (o *CampaignResponseAttributes) SetEntityScope(v string) {
	o.EntityScope = &v
}

// GetGuidance returns the Guidance field value if set, zero value otherwise.
func (o *CampaignResponseAttributes) GetGuidance() string {
	if o == nil || o.Guidance == nil {
		var ret string
		return ret
	}
	return *o.Guidance
}

// GetGuidanceOk returns a tuple with the Guidance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetGuidanceOk() (*string, bool) {
	if o == nil || o.Guidance == nil {
		return nil, false
	}
	return o.Guidance, true
}

// HasGuidance returns a boolean if a field has been set.
func (o *CampaignResponseAttributes) HasGuidance() bool {
	return o != nil && o.Guidance != nil
}

// SetGuidance gets a reference to the given string and assigns it to the Guidance field.
func (o *CampaignResponseAttributes) SetGuidance(v string) {
	o.Guidance = &v
}

// GetKey returns the Key field value.
func (o *CampaignResponseAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *CampaignResponseAttributes) SetKey(v string) {
	o.Key = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *CampaignResponseAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *CampaignResponseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *CampaignResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CampaignResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value.
func (o *CampaignResponseAttributes) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value.
func (o *CampaignResponseAttributes) SetOwner(v string) {
	o.Owner = v
}

// GetStartDate returns the StartDate field value.
func (o *CampaignResponseAttributes) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value.
func (o *CampaignResponseAttributes) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetStatus returns the Status field value.
func (o *CampaignResponseAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CampaignResponseAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CampaignResponseAttributes) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CampaignResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
	toSerialize["key"] = o.Key
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["owner"] = o.Owner
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
func (o *CampaignResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"created_at"`
		Description *string    `json:"description,omitempty"`
		DueDate     *time.Time `json:"due_date,omitempty"`
		EntityScope *string    `json:"entity_scope,omitempty"`
		Guidance    *string    `json:"guidance,omitempty"`
		Key         *string    `json:"key"`
		ModifiedAt  *time.Time `json:"modified_at"`
		Name        *string    `json:"name"`
		Owner       *string    `json:"owner"`
		StartDate   *time.Time `json:"start_date"`
		Status      *string    `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Owner == nil {
		return fmt.Errorf("required field owner missing")
	}
	if all.StartDate == nil {
		return fmt.Errorf("required field start_date missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "due_date", "entity_scope", "guidance", "key", "modified_at", "name", "owner", "start_date", "status"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.DueDate = all.DueDate
	o.EntityScope = all.EntityScope
	o.Guidance = all.Guidance
	o.Key = *all.Key
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.Owner = *all.Owner
	o.StartDate = *all.StartDate
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
