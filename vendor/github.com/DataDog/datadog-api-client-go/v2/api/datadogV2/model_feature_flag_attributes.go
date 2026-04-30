// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagAttributes Attributes of a feature flag.
type FeatureFlagAttributes struct {
	// The timestamp when the feature flag was archived.
	ArchivedAt datadog.NullableTime `json:"archived_at,omitempty"`
	// The timestamp when the feature flag was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The ID of the user who created the feature flag.
	CreatedBy *uuid.UUID `json:"created_by,omitempty"`
	// The description of the feature flag.
	Description string `json:"description"`
	// Distribution channel for the feature flag.
	DistributionChannel *string `json:"distribution_channel,omitempty"`
	// Environment-specific settings for the feature flag.
	FeatureFlagEnvironments []FeatureFlagEnvironment `json:"feature_flag_environments,omitempty"`
	// JSON schema for validation when value_type is JSON.
	JsonSchema datadog.NullableString `json:"json_schema,omitempty"`
	// The unique key of the feature flag.
	Key string `json:"key"`
	// The ID of the user who last updated the feature flag.
	LastUpdatedBy *uuid.UUID `json:"last_updated_by,omitempty"`
	// The name of the feature flag.
	Name string `json:"name"`
	// Indicates whether this feature flag requires approval for changes.
	RequireApproval *bool `json:"require_approval,omitempty"`
	// Tags associated with the feature flag.
	Tags []string `json:"tags,omitempty"`
	// The timestamp when the feature flag was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The type of values for the feature flag variants.
	ValueType ValueType `json:"value_type"`
	// The variants of the feature flag.
	Variants []Variant `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagAttributes instantiates a new FeatureFlagAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagAttributes(description string, key string, name string, valueType ValueType, variants []Variant) *FeatureFlagAttributes {
	this := FeatureFlagAttributes{}
	this.Description = description
	this.Key = key
	this.Name = name
	this.ValueType = valueType
	this.Variants = variants
	return &this
}

// NewFeatureFlagAttributesWithDefaults instantiates a new FeatureFlagAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagAttributesWithDefaults() *FeatureFlagAttributes {
	this := FeatureFlagAttributes{}
	return &this
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributes) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt.Get()
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributes) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivedAt.Get(), o.ArchivedAt.IsSet()
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasArchivedAt() bool {
	return o != nil && o.ArchivedAt.IsSet()
}

// SetArchivedAt gets a reference to the given datadog.NullableTime and assigns it to the ArchivedAt field.
func (o *FeatureFlagAttributes) SetArchivedAt(v time.Time) {
	o.ArchivedAt.Set(&v)
}

// SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil.
func (o *FeatureFlagAttributes) SetArchivedAtNil() {
	o.ArchivedAt.Set(nil)
}

// UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil.
func (o *FeatureFlagAttributes) UnsetArchivedAt() {
	o.ArchivedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FeatureFlagAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetCreatedBy() uuid.UUID {
	if o == nil || o.CreatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given uuid.UUID and assigns it to the CreatedBy field.
func (o *FeatureFlagAttributes) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value.
func (o *FeatureFlagAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FeatureFlagAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDistributionChannel returns the DistributionChannel field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetDistributionChannel() string {
	if o == nil || o.DistributionChannel == nil {
		var ret string
		return ret
	}
	return *o.DistributionChannel
}

// GetDistributionChannelOk returns a tuple with the DistributionChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetDistributionChannelOk() (*string, bool) {
	if o == nil || o.DistributionChannel == nil {
		return nil, false
	}
	return o.DistributionChannel, true
}

// HasDistributionChannel returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasDistributionChannel() bool {
	return o != nil && o.DistributionChannel != nil
}

// SetDistributionChannel gets a reference to the given string and assigns it to the DistributionChannel field.
func (o *FeatureFlagAttributes) SetDistributionChannel(v string) {
	o.DistributionChannel = &v
}

// GetFeatureFlagEnvironments returns the FeatureFlagEnvironments field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetFeatureFlagEnvironments() []FeatureFlagEnvironment {
	if o == nil || o.FeatureFlagEnvironments == nil {
		var ret []FeatureFlagEnvironment
		return ret
	}
	return o.FeatureFlagEnvironments
}

// GetFeatureFlagEnvironmentsOk returns a tuple with the FeatureFlagEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetFeatureFlagEnvironmentsOk() (*[]FeatureFlagEnvironment, bool) {
	if o == nil || o.FeatureFlagEnvironments == nil {
		return nil, false
	}
	return &o.FeatureFlagEnvironments, true
}

// HasFeatureFlagEnvironments returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasFeatureFlagEnvironments() bool {
	return o != nil && o.FeatureFlagEnvironments != nil
}

// SetFeatureFlagEnvironments gets a reference to the given []FeatureFlagEnvironment and assigns it to the FeatureFlagEnvironments field.
func (o *FeatureFlagAttributes) SetFeatureFlagEnvironments(v []FeatureFlagEnvironment) {
	o.FeatureFlagEnvironments = v
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagAttributes) GetJsonSchema() string {
	if o == nil || o.JsonSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema.Get()
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagAttributes) GetJsonSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonSchema.Get(), o.JsonSchema.IsSet()
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasJsonSchema() bool {
	return o != nil && o.JsonSchema.IsSet()
}

// SetJsonSchema gets a reference to the given datadog.NullableString and assigns it to the JsonSchema field.
func (o *FeatureFlagAttributes) SetJsonSchema(v string) {
	o.JsonSchema.Set(&v)
}

// SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil.
func (o *FeatureFlagAttributes) SetJsonSchemaNil() {
	o.JsonSchema.Set(nil)
}

// UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil.
func (o *FeatureFlagAttributes) UnsetJsonSchema() {
	o.JsonSchema.Unset()
}

// GetKey returns the Key field value.
func (o *FeatureFlagAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *FeatureFlagAttributes) SetKey(v string) {
	o.Key = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetLastUpdatedBy() uuid.UUID {
	if o == nil || o.LastUpdatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetLastUpdatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given uuid.UUID and assigns it to the LastUpdatedBy field.
func (o *FeatureFlagAttributes) SetLastUpdatedBy(v uuid.UUID) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value.
func (o *FeatureFlagAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FeatureFlagAttributes) SetName(v string) {
	o.Name = v
}

// GetRequireApproval returns the RequireApproval field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetRequireApproval() bool {
	if o == nil || o.RequireApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireApproval
}

// GetRequireApprovalOk returns a tuple with the RequireApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetRequireApprovalOk() (*bool, bool) {
	if o == nil || o.RequireApproval == nil {
		return nil, false
	}
	return o.RequireApproval, true
}

// HasRequireApproval returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasRequireApproval() bool {
	return o != nil && o.RequireApproval != nil
}

// SetRequireApproval gets a reference to the given bool and assigns it to the RequireApproval field.
func (o *FeatureFlagAttributes) SetRequireApproval(v bool) {
	o.RequireApproval = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FeatureFlagAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FeatureFlagAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FeatureFlagAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FeatureFlagAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValueType returns the ValueType field value.
func (o *FeatureFlagAttributes) GetValueType() ValueType {
	if o == nil {
		var ret ValueType
		return ret
	}
	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetValueTypeOk() (*ValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value.
func (o *FeatureFlagAttributes) SetValueType(v ValueType) {
	o.ValueType = v
}

// GetVariants returns the Variants field value.
func (o *FeatureFlagAttributes) GetVariants() []Variant {
	if o == nil {
		var ret []Variant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagAttributes) GetVariantsOk() (*[]Variant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *FeatureFlagAttributes) SetVariants(v []Variant) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArchivedAt.IsSet() {
		toSerialize["archived_at"] = o.ArchivedAt.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	toSerialize["description"] = o.Description
	if o.DistributionChannel != nil {
		toSerialize["distribution_channel"] = o.DistributionChannel
	}
	if o.FeatureFlagEnvironments != nil {
		toSerialize["feature_flag_environments"] = o.FeatureFlagEnvironments
	}
	if o.JsonSchema.IsSet() {
		toSerialize["json_schema"] = o.JsonSchema.Get()
	}
	toSerialize["key"] = o.Key
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	toSerialize["name"] = o.Name
	if o.RequireApproval != nil {
		toSerialize["require_approval"] = o.RequireApproval
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["value_type"] = o.ValueType
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchivedAt              datadog.NullableTime     `json:"archived_at,omitempty"`
		CreatedAt               *time.Time               `json:"created_at,omitempty"`
		CreatedBy               *uuid.UUID               `json:"created_by,omitempty"`
		Description             *string                  `json:"description"`
		DistributionChannel     *string                  `json:"distribution_channel,omitempty"`
		FeatureFlagEnvironments []FeatureFlagEnvironment `json:"feature_flag_environments,omitempty"`
		JsonSchema              datadog.NullableString   `json:"json_schema,omitempty"`
		Key                     *string                  `json:"key"`
		LastUpdatedBy           *uuid.UUID               `json:"last_updated_by,omitempty"`
		Name                    *string                  `json:"name"`
		RequireApproval         *bool                    `json:"require_approval,omitempty"`
		Tags                    []string                 `json:"tags,omitempty"`
		UpdatedAt               *time.Time               `json:"updated_at,omitempty"`
		ValueType               *ValueType               `json:"value_type"`
		Variants                *[]Variant               `json:"variants"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ValueType == nil {
		return fmt.Errorf("required field value_type missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"archived_at", "created_at", "created_by", "description", "distribution_channel", "feature_flag_environments", "json_schema", "key", "last_updated_by", "name", "require_approval", "tags", "updated_at", "value_type", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ArchivedAt = all.ArchivedAt
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Description = *all.Description
	o.DistributionChannel = all.DistributionChannel
	o.FeatureFlagEnvironments = all.FeatureFlagEnvironments
	o.JsonSchema = all.JsonSchema
	o.Key = *all.Key
	o.LastUpdatedBy = all.LastUpdatedBy
	o.Name = *all.Name
	o.RequireApproval = all.RequireApproval
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt
	if !all.ValueType.IsValid() {
		hasInvalidField = true
	} else {
		o.ValueType = *all.ValueType
	}
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
