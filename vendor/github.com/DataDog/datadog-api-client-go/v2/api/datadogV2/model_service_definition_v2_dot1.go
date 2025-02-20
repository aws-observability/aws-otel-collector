// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1 Service definition v2.1 for providing service metadata and integrations.
type ServiceDefinitionV2Dot1 struct {
	// Identifier for a group of related services serving a product feature, which the service is a part of.
	Application *string `json:"application,omitempty"`
	// A list of contacts related to the services.
	Contacts []ServiceDefinitionV2Dot1Contact `json:"contacts,omitempty"`
	// Unique identifier of the service. Must be unique across all services and is used to match with a service in Datadog.
	DdService string `json:"dd-service"`
	// A short description of the service.
	Description *string `json:"description,omitempty"`
	// Extensions to v2.1 schema.
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	// Third party integrations that Datadog supports.
	Integrations *ServiceDefinitionV2Dot1Integrations `json:"integrations,omitempty"`
	// The current life cycle phase of the service.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// A list of links related to the services.
	Links []ServiceDefinitionV2Dot1Link `json:"links,omitempty"`
	// Schema version being used.
	SchemaVersion ServiceDefinitionV2Dot1Version `json:"schema-version"`
	// A set of custom tags.
	Tags []string `json:"tags,omitempty"`
	// Team that owns the service. It is used to locate a team defined in Datadog Teams if it exists.
	Team *string `json:"team,omitempty"`
	// Importance of the service.
	Tier *string `json:"tier,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionV2Dot1 instantiates a new ServiceDefinitionV2Dot1 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Dot1(ddService string, schemaVersion ServiceDefinitionV2Dot1Version) *ServiceDefinitionV2Dot1 {
	this := ServiceDefinitionV2Dot1{}
	this.DdService = ddService
	this.SchemaVersion = schemaVersion
	return &this
}

// NewServiceDefinitionV2Dot1WithDefaults instantiates a new ServiceDefinitionV2Dot1 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2Dot1WithDefaults() *ServiceDefinitionV2Dot1 {
	this := ServiceDefinitionV2Dot1{}
	var schemaVersion ServiceDefinitionV2Dot1Version = SERVICEDEFINITIONV2DOT1VERSION_V2_1
	this.SchemaVersion = schemaVersion
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ServiceDefinitionV2Dot1) SetApplication(v string) {
	o.Application = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetContacts() []ServiceDefinitionV2Dot1Contact {
	if o == nil || o.Contacts == nil {
		var ret []ServiceDefinitionV2Dot1Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetContactsOk() (*[]ServiceDefinitionV2Dot1Contact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return &o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasContacts() bool {
	return o != nil && o.Contacts != nil
}

// SetContacts gets a reference to the given []ServiceDefinitionV2Dot1Contact and assigns it to the Contacts field.
func (o *ServiceDefinitionV2Dot1) SetContacts(v []ServiceDefinitionV2Dot1Contact) {
	o.Contacts = v
}

// GetDdService returns the DdService field value.
func (o *ServiceDefinitionV2Dot1) GetDdService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdService
}

// GetDdServiceOk returns a tuple with the DdService field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetDdServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdService, true
}

// SetDdService sets field value.
func (o *ServiceDefinitionV2Dot1) SetDdService(v string) {
	o.DdService = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceDefinitionV2Dot1) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetExtensionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasExtensions() bool {
	return o != nil && o.Extensions != nil
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *ServiceDefinitionV2Dot1) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetIntegrations() ServiceDefinitionV2Dot1Integrations {
	if o == nil || o.Integrations == nil {
		var ret ServiceDefinitionV2Dot1Integrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetIntegrationsOk() (*ServiceDefinitionV2Dot1Integrations, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given ServiceDefinitionV2Dot1Integrations and assigns it to the Integrations field.
func (o *ServiceDefinitionV2Dot1) SetIntegrations(v ServiceDefinitionV2Dot1Integrations) {
	o.Integrations = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *ServiceDefinitionV2Dot1) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetLinks() []ServiceDefinitionV2Dot1Link {
	if o == nil || o.Links == nil {
		var ret []ServiceDefinitionV2Dot1Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetLinksOk() (*[]ServiceDefinitionV2Dot1Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []ServiceDefinitionV2Dot1Link and assigns it to the Links field.
func (o *ServiceDefinitionV2Dot1) SetLinks(v []ServiceDefinitionV2Dot1Link) {
	o.Links = v
}

// GetSchemaVersion returns the SchemaVersion field value.
func (o *ServiceDefinitionV2Dot1) GetSchemaVersion() ServiceDefinitionV2Dot1Version {
	if o == nil {
		var ret ServiceDefinitionV2Dot1Version
		return ret
	}
	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetSchemaVersionOk() (*ServiceDefinitionV2Dot1Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value.
func (o *ServiceDefinitionV2Dot1) SetSchemaVersion(v ServiceDefinitionV2Dot1Version) {
	o.SchemaVersion = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceDefinitionV2Dot1) SetTags(v []string) {
	o.Tags = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *ServiceDefinitionV2Dot1) SetTeam(v string) {
	o.Team = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *ServiceDefinitionV2Dot1) SetTier(v string) {
	o.Tier = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Dot1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	toSerialize["dd-service"] = o.DdService
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	toSerialize["schema-version"] = o.SchemaVersion
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Dot1) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Application   *string                              `json:"application,omitempty"`
		Contacts      []ServiceDefinitionV2Dot1Contact     `json:"contacts,omitempty"`
		DdService     *string                              `json:"dd-service"`
		Description   *string                              `json:"description,omitempty"`
		Extensions    map[string]interface{}               `json:"extensions,omitempty"`
		Integrations  *ServiceDefinitionV2Dot1Integrations `json:"integrations,omitempty"`
		Lifecycle     *string                              `json:"lifecycle,omitempty"`
		Links         []ServiceDefinitionV2Dot1Link        `json:"links,omitempty"`
		SchemaVersion *ServiceDefinitionV2Dot1Version      `json:"schema-version"`
		Tags          []string                             `json:"tags,omitempty"`
		Team          *string                              `json:"team,omitempty"`
		Tier          *string                              `json:"tier,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DdService == nil {
		return fmt.Errorf("required field dd-service missing")
	}
	if all.SchemaVersion == nil {
		return fmt.Errorf("required field schema-version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application", "contacts", "dd-service", "description", "extensions", "integrations", "lifecycle", "links", "schema-version", "tags", "team", "tier"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Application = all.Application
	o.Contacts = all.Contacts
	o.DdService = *all.DdService
	o.Description = all.Description
	o.Extensions = all.Extensions
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integrations = all.Integrations
	o.Lifecycle = all.Lifecycle
	o.Links = all.Links
	if !all.SchemaVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.SchemaVersion = *all.SchemaVersion
	}
	o.Tags = all.Tags
	o.Team = all.Team
	o.Tier = all.Tier

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
