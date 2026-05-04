// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EnvironmentAttributes Attributes of an environment.
type EnvironmentAttributes struct {
	// The timestamp when the environment was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The description of the environment.
	Description datadog.NullableString `json:"description,omitempty"`
	// Indicates whether this is a production environment.
	IsProduction *bool `json:"is_production,omitempty"`
	// The unique key of the environment.
	Key *string `json:"key,omitempty"`
	// The name of the environment.
	Name string `json:"name"`
	// List of queries to define the environment scope.
	Queries []string `json:"queries,omitempty"`
	// Indicates whether feature flag changes require approval in this environment.
	RequireFeatureFlagApproval *bool `json:"require_feature_flag_approval,omitempty"`
	// The timestamp when the environment was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentAttributes instantiates a new EnvironmentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentAttributes(name string) *EnvironmentAttributes {
	this := EnvironmentAttributes{}
	this.Name = name
	return &this
}

// NewEnvironmentAttributesWithDefaults instantiates a new EnvironmentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentAttributesWithDefaults() *EnvironmentAttributes {
	this := EnvironmentAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvironmentAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *EnvironmentAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *EnvironmentAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *EnvironmentAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *EnvironmentAttributes) GetIsProduction() bool {
	if o == nil || o.IsProduction == nil {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetIsProductionOk() (*bool, bool) {
	if o == nil || o.IsProduction == nil {
		return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasIsProduction() bool {
	return o != nil && o.IsProduction != nil
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *EnvironmentAttributes) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EnvironmentAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EnvironmentAttributes) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value.
func (o *EnvironmentAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentAttributes) SetName(v string) {
	o.Name = v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *EnvironmentAttributes) GetQueries() []string {
	if o == nil || o.Queries == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetQueriesOk() (*[]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *EnvironmentAttributes) SetQueries(v []string) {
	o.Queries = v
}

// GetRequireFeatureFlagApproval returns the RequireFeatureFlagApproval field value if set, zero value otherwise.
func (o *EnvironmentAttributes) GetRequireFeatureFlagApproval() bool {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireFeatureFlagApproval
}

// GetRequireFeatureFlagApprovalOk returns a tuple with the RequireFeatureFlagApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetRequireFeatureFlagApprovalOk() (*bool, bool) {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		return nil, false
	}
	return o.RequireFeatureFlagApproval, true
}

// HasRequireFeatureFlagApproval returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasRequireFeatureFlagApproval() bool {
	return o != nil && o.RequireFeatureFlagApproval != nil
}

// SetRequireFeatureFlagApproval gets a reference to the given bool and assigns it to the RequireFeatureFlagApproval field.
func (o *EnvironmentAttributes) SetRequireFeatureFlagApproval(v bool) {
	o.RequireFeatureFlagApproval = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IsProduction != nil {
		toSerialize["is_production"] = o.IsProduction
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	toSerialize["name"] = o.Name
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.RequireFeatureFlagApproval != nil {
		toSerialize["require_feature_flag_approval"] = o.RequireFeatureFlagApproval
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                  *time.Time             `json:"created_at,omitempty"`
		Description                datadog.NullableString `json:"description,omitempty"`
		IsProduction               *bool                  `json:"is_production,omitempty"`
		Key                        *string                `json:"key,omitempty"`
		Name                       *string                `json:"name"`
		Queries                    []string               `json:"queries,omitempty"`
		RequireFeatureFlagApproval *bool                  `json:"require_feature_flag_approval,omitempty"`
		UpdatedAt                  *time.Time             `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "is_production", "key", "name", "queries", "require_feature_flag_approval", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.IsProduction = all.IsProduction
	o.Key = all.Key
	o.Name = *all.Name
	o.Queries = all.Queries
	o.RequireFeatureFlagApproval = all.RequireFeatureFlagApproval
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
