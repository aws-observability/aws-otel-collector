// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleResponseDataAttributes Basic information about a deployment rule.
type DeploymentRuleResponseDataAttributes struct {
	// The timestamp when the deployment rule was created.
	CreatedAt time.Time `json:"created_at"`
	// Information about the user who created the deployment rule.
	CreatedBy DeploymentRuleResponseDataAttributesCreatedBy `json:"created_by"`
	// Whether this rule is run in dry-run mode.
	DryRun bool `json:"dry_run"`
	// The ID of the deployment gate.
	GateId string `json:"gate_id"`
	// The name of the deployment rule.
	Name string `json:"name"`
	// Options for deployment rule response representing either faulty deployment detection or monitor options.
	Options DeploymentRulesOptions `json:"options"`
	// The type of the deployment rule.
	Type DeploymentRuleResponseDataAttributesType `json:"type"`
	// The timestamp when the deployment rule was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Information about the user who updated the deployment rule.
	UpdatedBy *DeploymentRuleResponseDataAttributesUpdatedBy `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentRuleResponseDataAttributes instantiates a new DeploymentRuleResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRuleResponseDataAttributes(createdAt time.Time, createdBy DeploymentRuleResponseDataAttributesCreatedBy, dryRun bool, gateId string, name string, options DeploymentRulesOptions, typeVar DeploymentRuleResponseDataAttributesType) *DeploymentRuleResponseDataAttributes {
	this := DeploymentRuleResponseDataAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DryRun = dryRun
	this.GateId = gateId
	this.Name = name
	this.Options = options
	this.Type = typeVar
	return &this
}

// NewDeploymentRuleResponseDataAttributesWithDefaults instantiates a new DeploymentRuleResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRuleResponseDataAttributesWithDefaults() *DeploymentRuleResponseDataAttributes {
	this := DeploymentRuleResponseDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DeploymentRuleResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *DeploymentRuleResponseDataAttributes) GetCreatedBy() DeploymentRuleResponseDataAttributesCreatedBy {
	if o == nil {
		var ret DeploymentRuleResponseDataAttributesCreatedBy
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetCreatedByOk() (*DeploymentRuleResponseDataAttributesCreatedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetCreatedBy(v DeploymentRuleResponseDataAttributesCreatedBy) {
	o.CreatedBy = v
}

// GetDryRun returns the DryRun field value.
func (o *DeploymentRuleResponseDataAttributes) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetDryRun(v bool) {
	o.DryRun = v
}

// GetGateId returns the GateId field value.
func (o *DeploymentRuleResponseDataAttributes) GetGateId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GateId
}

// GetGateIdOk returns a tuple with the GateId field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetGateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GateId, true
}

// SetGateId sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetGateId(v string) {
	o.GateId = v
}

// GetName returns the Name field value.
func (o *DeploymentRuleResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *DeploymentRuleResponseDataAttributes) GetOptions() DeploymentRulesOptions {
	if o == nil {
		var ret DeploymentRulesOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetOptionsOk() (*DeploymentRulesOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetOptions(v DeploymentRulesOptions) {
	o.Options = v
}

// GetType returns the Type field value.
func (o *DeploymentRuleResponseDataAttributes) GetType() DeploymentRuleResponseDataAttributesType {
	if o == nil {
		var ret DeploymentRuleResponseDataAttributesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetTypeOk() (*DeploymentRuleResponseDataAttributesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DeploymentRuleResponseDataAttributes) SetType(v DeploymentRuleResponseDataAttributesType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentRuleResponseDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentRuleResponseDataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentRuleResponseDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DeploymentRuleResponseDataAttributes) GetUpdatedBy() DeploymentRuleResponseDataAttributesUpdatedBy {
	if o == nil || o.UpdatedBy == nil {
		var ret DeploymentRuleResponseDataAttributesUpdatedBy
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRuleResponseDataAttributes) GetUpdatedByOk() (*DeploymentRuleResponseDataAttributesUpdatedBy, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DeploymentRuleResponseDataAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given DeploymentRuleResponseDataAttributesUpdatedBy and assigns it to the UpdatedBy field.
func (o *DeploymentRuleResponseDataAttributes) SetUpdatedBy(v DeploymentRuleResponseDataAttributesUpdatedBy) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRuleResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["dry_run"] = o.DryRun
	toSerialize["gate_id"] = o.GateId
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["type"] = o.Type
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRuleResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                                     `json:"created_at"`
		CreatedBy *DeploymentRuleResponseDataAttributesCreatedBy `json:"created_by"`
		DryRun    *bool                                          `json:"dry_run"`
		GateId    *string                                        `json:"gate_id"`
		Name      *string                                        `json:"name"`
		Options   *DeploymentRulesOptions                        `json:"options"`
		Type      *DeploymentRuleResponseDataAttributesType      `json:"type"`
		UpdatedAt *time.Time                                     `json:"updated_at,omitempty"`
		UpdatedBy *DeploymentRuleResponseDataAttributesUpdatedBy `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.DryRun == nil {
		return fmt.Errorf("required field dry_run missing")
	}
	if all.GateId == nil {
		return fmt.Errorf("required field gate_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "dry_run", "gate_id", "name", "options", "type", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	o.DryRun = *all.DryRun
	o.GateId = *all.GateId
	o.Name = *all.Name
	o.Options = *all.Options
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UpdatedAt = all.UpdatedAt
	if all.UpdatedBy != nil && all.UpdatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
