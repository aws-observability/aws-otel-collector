// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateResponseDataAttributes Basic information about a deployment gate.
type DeploymentGateResponseDataAttributes struct {
	// The timestamp when the deployment gate was created.
	CreatedAt time.Time `json:"created_at"`
	// Information about the user who created the deployment gate.
	CreatedBy DeploymentGateResponseDataAttributesCreatedBy `json:"created_by"`
	// Whether this gate is run in dry-run mode.
	DryRun bool `json:"dry_run"`
	// The environment of the deployment gate.
	Env string `json:"env"`
	// The identifier of the deployment gate.
	Identifier string `json:"identifier"`
	// The service of the deployment gate.
	Service string `json:"service"`
	// The timestamp when the deployment gate was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Information about the user who updated the deployment gate.
	UpdatedBy *DeploymentGateResponseDataAttributesUpdatedBy `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateResponseDataAttributes instantiates a new DeploymentGateResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateResponseDataAttributes(createdAt time.Time, createdBy DeploymentGateResponseDataAttributesCreatedBy, dryRun bool, env string, identifier string, service string) *DeploymentGateResponseDataAttributes {
	this := DeploymentGateResponseDataAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DryRun = dryRun
	this.Env = env
	this.Identifier = identifier
	this.Service = service
	return &this
}

// NewDeploymentGateResponseDataAttributesWithDefaults instantiates a new DeploymentGateResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateResponseDataAttributesWithDefaults() *DeploymentGateResponseDataAttributes {
	this := DeploymentGateResponseDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DeploymentGateResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DeploymentGateResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *DeploymentGateResponseDataAttributes) GetCreatedBy() DeploymentGateResponseDataAttributesCreatedBy {
	if o == nil {
		var ret DeploymentGateResponseDataAttributesCreatedBy
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetCreatedByOk() (*DeploymentGateResponseDataAttributesCreatedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *DeploymentGateResponseDataAttributes) SetCreatedBy(v DeploymentGateResponseDataAttributesCreatedBy) {
	o.CreatedBy = v
}

// GetDryRun returns the DryRun field value.
func (o *DeploymentGateResponseDataAttributes) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *DeploymentGateResponseDataAttributes) SetDryRun(v bool) {
	o.DryRun = v
}

// GetEnv returns the Env field value.
func (o *DeploymentGateResponseDataAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *DeploymentGateResponseDataAttributes) SetEnv(v string) {
	o.Env = v
}

// GetIdentifier returns the Identifier field value.
func (o *DeploymentGateResponseDataAttributes) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value.
func (o *DeploymentGateResponseDataAttributes) SetIdentifier(v string) {
	o.Identifier = v
}

// GetService returns the Service field value.
func (o *DeploymentGateResponseDataAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *DeploymentGateResponseDataAttributes) SetService(v string) {
	o.Service = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentGateResponseDataAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentGateResponseDataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentGateResponseDataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DeploymentGateResponseDataAttributes) GetUpdatedBy() DeploymentGateResponseDataAttributesUpdatedBy {
	if o == nil || o.UpdatedBy == nil {
		var ret DeploymentGateResponseDataAttributesUpdatedBy
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateResponseDataAttributes) GetUpdatedByOk() (*DeploymentGateResponseDataAttributesUpdatedBy, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DeploymentGateResponseDataAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given DeploymentGateResponseDataAttributesUpdatedBy and assigns it to the UpdatedBy field.
func (o *DeploymentGateResponseDataAttributes) SetUpdatedBy(v DeploymentGateResponseDataAttributesUpdatedBy) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["env"] = o.Env
	toSerialize["identifier"] = o.Identifier
	toSerialize["service"] = o.Service
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
func (o *DeploymentGateResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time                                     `json:"created_at"`
		CreatedBy  *DeploymentGateResponseDataAttributesCreatedBy `json:"created_by"`
		DryRun     *bool                                          `json:"dry_run"`
		Env        *string                                        `json:"env"`
		Identifier *string                                        `json:"identifier"`
		Service    *string                                        `json:"service"`
		UpdatedAt  *time.Time                                     `json:"updated_at,omitempty"`
		UpdatedBy  *DeploymentGateResponseDataAttributesUpdatedBy `json:"updated_by,omitempty"`
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
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Identifier == nil {
		return fmt.Errorf("required field identifier missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "dry_run", "env", "identifier", "service", "updated_at", "updated_by"})
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
	o.Env = *all.Env
	o.Identifier = *all.Identifier
	o.Service = *all.Service
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
