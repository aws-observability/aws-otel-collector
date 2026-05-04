// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityMetadata Metadata about the entity from cloud providers
type SecurityEntityMetadata struct {
	// Cloud account ID (AWS)
	AccountId *string `json:"accountID,omitempty"`
	// Environment tags associated with the entity
	Environments []string `json:"environments"`
	// MITRE ATT&CK tactics detected
	MitreTactics []string `json:"mitreTactics"`
	// MITRE ATT&CK techniques detected
	MitreTechniques []string `json:"mitreTechniques"`
	// Cloud project ID (GCP)
	ProjectId *string `json:"projectID,omitempty"`
	// Services associated with the entity
	Services []string `json:"services"`
	// Data sources that detected this entity
	Sources []string `json:"sources"`
	// Cloud subscription ID (Azure)
	SubscriptionId *string `json:"subscriptionID,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityEntityMetadata instantiates a new SecurityEntityMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityEntityMetadata(environments []string, mitreTactics []string, mitreTechniques []string, services []string, sources []string) *SecurityEntityMetadata {
	this := SecurityEntityMetadata{}
	this.Environments = environments
	this.MitreTactics = mitreTactics
	this.MitreTechniques = mitreTechniques
	this.Services = services
	this.Sources = sources
	return &this
}

// NewSecurityEntityMetadataWithDefaults instantiates a new SecurityEntityMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityEntityMetadataWithDefaults() *SecurityEntityMetadata {
	this := SecurityEntityMetadata{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SecurityEntityMetadata) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SecurityEntityMetadata) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SecurityEntityMetadata) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEnvironments returns the Environments field value.
func (o *SecurityEntityMetadata) GetEnvironments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetEnvironmentsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environments, true
}

// SetEnvironments sets field value.
func (o *SecurityEntityMetadata) SetEnvironments(v []string) {
	o.Environments = v
}

// GetMitreTactics returns the MitreTactics field value.
func (o *SecurityEntityMetadata) GetMitreTactics() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MitreTactics
}

// GetMitreTacticsOk returns a tuple with the MitreTactics field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetMitreTacticsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MitreTactics, true
}

// SetMitreTactics sets field value.
func (o *SecurityEntityMetadata) SetMitreTactics(v []string) {
	o.MitreTactics = v
}

// GetMitreTechniques returns the MitreTechniques field value.
func (o *SecurityEntityMetadata) GetMitreTechniques() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MitreTechniques
}

// GetMitreTechniquesOk returns a tuple with the MitreTechniques field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetMitreTechniquesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MitreTechniques, true
}

// SetMitreTechniques sets field value.
func (o *SecurityEntityMetadata) SetMitreTechniques(v []string) {
	o.MitreTechniques = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *SecurityEntityMetadata) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *SecurityEntityMetadata) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *SecurityEntityMetadata) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetServices returns the Services field value.
func (o *SecurityEntityMetadata) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetServicesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value.
func (o *SecurityEntityMetadata) SetServices(v []string) {
	o.Services = v
}

// GetSources returns the Sources field value.
func (o *SecurityEntityMetadata) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *SecurityEntityMetadata) SetSources(v []string) {
	o.Sources = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SecurityEntityMetadata) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityMetadata) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SecurityEntityMetadata) HasSubscriptionId() bool {
	return o != nil && o.SubscriptionId != nil
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SecurityEntityMetadata) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityEntityMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["accountID"] = o.AccountId
	}
	toSerialize["environments"] = o.Environments
	toSerialize["mitreTactics"] = o.MitreTactics
	toSerialize["mitreTechniques"] = o.MitreTechniques
	if o.ProjectId != nil {
		toSerialize["projectID"] = o.ProjectId
	}
	toSerialize["services"] = o.Services
	toSerialize["sources"] = o.Sources
	if o.SubscriptionId != nil {
		toSerialize["subscriptionID"] = o.SubscriptionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityEntityMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId       *string   `json:"accountID,omitempty"`
		Environments    *[]string `json:"environments"`
		MitreTactics    *[]string `json:"mitreTactics"`
		MitreTechniques *[]string `json:"mitreTechniques"`
		ProjectId       *string   `json:"projectID,omitempty"`
		Services        *[]string `json:"services"`
		Sources         *[]string `json:"sources"`
		SubscriptionId  *string   `json:"subscriptionID,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Environments == nil {
		return fmt.Errorf("required field environments missing")
	}
	if all.MitreTactics == nil {
		return fmt.Errorf("required field mitreTactics missing")
	}
	if all.MitreTechniques == nil {
		return fmt.Errorf("required field mitreTechniques missing")
	}
	if all.Services == nil {
		return fmt.Errorf("required field services missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accountID", "environments", "mitreTactics", "mitreTechniques", "projectID", "services", "sources", "subscriptionID"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.Environments = *all.Environments
	o.MitreTactics = *all.MitreTactics
	o.MitreTechniques = *all.MitreTechniques
	o.ProjectId = all.ProjectId
	o.Services = *all.Services
	o.Sources = *all.Sources
	o.SubscriptionId = all.SubscriptionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
