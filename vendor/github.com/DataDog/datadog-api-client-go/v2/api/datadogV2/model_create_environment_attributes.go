// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateEnvironmentAttributes Attributes for creating a new environment.
type CreateEnvironmentAttributes struct {
	// Indicates whether this is a production environment.
	IsProduction *bool `json:"is_production,omitempty"`
	// The name of the environment.
	Name string `json:"name"`
	// List of queries to define the environment scope.
	Queries []string `json:"queries"`
	// Indicates whether feature flag changes require approval in this environment.
	RequireFeatureFlagApproval *bool `json:"require_feature_flag_approval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateEnvironmentAttributes instantiates a new CreateEnvironmentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateEnvironmentAttributes(name string, queries []string) *CreateEnvironmentAttributes {
	this := CreateEnvironmentAttributes{}
	var isProduction bool = false
	this.IsProduction = &isProduction
	this.Name = name
	this.Queries = queries
	var requireFeatureFlagApproval bool = false
	this.RequireFeatureFlagApproval = &requireFeatureFlagApproval
	return &this
}

// NewCreateEnvironmentAttributesWithDefaults instantiates a new CreateEnvironmentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateEnvironmentAttributesWithDefaults() *CreateEnvironmentAttributes {
	this := CreateEnvironmentAttributes{}
	var isProduction bool = false
	this.IsProduction = &isProduction
	var requireFeatureFlagApproval bool = false
	this.RequireFeatureFlagApproval = &requireFeatureFlagApproval
	return &this
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *CreateEnvironmentAttributes) GetIsProduction() bool {
	if o == nil || o.IsProduction == nil {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentAttributes) GetIsProductionOk() (*bool, bool) {
	if o == nil || o.IsProduction == nil {
		return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *CreateEnvironmentAttributes) HasIsProduction() bool {
	return o != nil && o.IsProduction != nil
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *CreateEnvironmentAttributes) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetName returns the Name field value.
func (o *CreateEnvironmentAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateEnvironmentAttributes) SetName(v string) {
	o.Name = v
}

// GetQueries returns the Queries field value.
func (o *CreateEnvironmentAttributes) GetQueries() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentAttributes) GetQueriesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *CreateEnvironmentAttributes) SetQueries(v []string) {
	o.Queries = v
}

// GetRequireFeatureFlagApproval returns the RequireFeatureFlagApproval field value if set, zero value otherwise.
func (o *CreateEnvironmentAttributes) GetRequireFeatureFlagApproval() bool {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireFeatureFlagApproval
}

// GetRequireFeatureFlagApprovalOk returns a tuple with the RequireFeatureFlagApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentAttributes) GetRequireFeatureFlagApprovalOk() (*bool, bool) {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		return nil, false
	}
	return o.RequireFeatureFlagApproval, true
}

// HasRequireFeatureFlagApproval returns a boolean if a field has been set.
func (o *CreateEnvironmentAttributes) HasRequireFeatureFlagApproval() bool {
	return o != nil && o.RequireFeatureFlagApproval != nil
}

// SetRequireFeatureFlagApproval gets a reference to the given bool and assigns it to the RequireFeatureFlagApproval field.
func (o *CreateEnvironmentAttributes) SetRequireFeatureFlagApproval(v bool) {
	o.RequireFeatureFlagApproval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateEnvironmentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsProduction != nil {
		toSerialize["is_production"] = o.IsProduction
	}
	toSerialize["name"] = o.Name
	toSerialize["queries"] = o.Queries
	if o.RequireFeatureFlagApproval != nil {
		toSerialize["require_feature_flag_approval"] = o.RequireFeatureFlagApproval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateEnvironmentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsProduction               *bool     `json:"is_production,omitempty"`
		Name                       *string   `json:"name"`
		Queries                    *[]string `json:"queries"`
		RequireFeatureFlagApproval *bool     `json:"require_feature_flag_approval,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_production", "name", "queries", "require_feature_flag_approval"})
	} else {
		return err
	}
	o.IsProduction = all.IsProduction
	o.Name = *all.Name
	o.Queries = *all.Queries
	o.RequireFeatureFlagApproval = all.RequireFeatureFlagApproval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
