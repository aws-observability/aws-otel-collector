// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot2 Service definition v2.2 for providing service metadata and integrations.
type ServiceDefinitionV2Dot2 struct {
	// Identifier for a group of related services serving a product feature, which the service is a part of.
	Application *string `json:"application,omitempty"`
	// A set of CI fingerprints.
	CiPipelineFingerprints []string `json:"ci-pipeline-fingerprints,omitempty"`
	// A list of contacts related to the services.
	Contacts []ServiceDefinitionV2Dot2Contact `json:"contacts,omitempty"`
	// Unique identifier of the service. Must be unique across all services and is used to match with a service in Datadog.
	DdService string `json:"dd-service"`
	// A short description of the service.
	Description *string `json:"description,omitempty"`
	// Extensions to v2.2 schema.
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	// Third party integrations that Datadog supports.
	Integrations *ServiceDefinitionV2Dot2Integrations `json:"integrations,omitempty"`
	// The service's programming language. Datadog recognizes the following languages: `dotnet`, `go`, `java`, `js`, `php`, `python`, `ruby`, and `c++`.
	Languages []string `json:"languages,omitempty"`
	// The current life cycle phase of the service.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// A list of links related to the services.
	Links []ServiceDefinitionV2Dot2Link `json:"links,omitempty"`
	// Schema version being used.
	SchemaVersion ServiceDefinitionV2Dot2Version `json:"schema-version"`
	// A set of custom tags.
	Tags []string `json:"tags,omitempty"`
	// Team that owns the service. It is used to locate a team defined in Datadog Teams if it exists.
	Team *string `json:"team,omitempty"`
	// Importance of the service.
	Tier *string `json:"tier,omitempty"`
	// The type of service.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionV2Dot2 instantiates a new ServiceDefinitionV2Dot2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Dot2(ddService string, schemaVersion ServiceDefinitionV2Dot2Version) *ServiceDefinitionV2Dot2 {
	this := ServiceDefinitionV2Dot2{}
	this.DdService = ddService
	this.SchemaVersion = schemaVersion
	return &this
}

// NewServiceDefinitionV2Dot2WithDefaults instantiates a new ServiceDefinitionV2Dot2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2Dot2WithDefaults() *ServiceDefinitionV2Dot2 {
	this := ServiceDefinitionV2Dot2{}
	var schemaVersion ServiceDefinitionV2Dot2Version = SERVICEDEFINITIONV2DOT2VERSION_V2_2
	this.SchemaVersion = schemaVersion
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetApplicationOk() (*string, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasApplication() bool {
	return o != nil && o.Application != nil
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ServiceDefinitionV2Dot2) SetApplication(v string) {
	o.Application = &v
}

// GetCiPipelineFingerprints returns the CiPipelineFingerprints field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetCiPipelineFingerprints() []string {
	if o == nil || o.CiPipelineFingerprints == nil {
		var ret []string
		return ret
	}
	return o.CiPipelineFingerprints
}

// GetCiPipelineFingerprintsOk returns a tuple with the CiPipelineFingerprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetCiPipelineFingerprintsOk() (*[]string, bool) {
	if o == nil || o.CiPipelineFingerprints == nil {
		return nil, false
	}
	return &o.CiPipelineFingerprints, true
}

// HasCiPipelineFingerprints returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasCiPipelineFingerprints() bool {
	return o != nil && o.CiPipelineFingerprints != nil
}

// SetCiPipelineFingerprints gets a reference to the given []string and assigns it to the CiPipelineFingerprints field.
func (o *ServiceDefinitionV2Dot2) SetCiPipelineFingerprints(v []string) {
	o.CiPipelineFingerprints = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetContacts() []ServiceDefinitionV2Dot2Contact {
	if o == nil || o.Contacts == nil {
		var ret []ServiceDefinitionV2Dot2Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetContactsOk() (*[]ServiceDefinitionV2Dot2Contact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return &o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasContacts() bool {
	return o != nil && o.Contacts != nil
}

// SetContacts gets a reference to the given []ServiceDefinitionV2Dot2Contact and assigns it to the Contacts field.
func (o *ServiceDefinitionV2Dot2) SetContacts(v []ServiceDefinitionV2Dot2Contact) {
	o.Contacts = v
}

// GetDdService returns the DdService field value.
func (o *ServiceDefinitionV2Dot2) GetDdService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DdService
}

// GetDdServiceOk returns a tuple with the DdService field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetDdServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdService, true
}

// SetDdService sets field value.
func (o *ServiceDefinitionV2Dot2) SetDdService(v string) {
	o.DdService = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceDefinitionV2Dot2) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetExtensionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasExtensions() bool {
	return o != nil && o.Extensions != nil
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *ServiceDefinitionV2Dot2) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetIntegrations() ServiceDefinitionV2Dot2Integrations {
	if o == nil || o.Integrations == nil {
		var ret ServiceDefinitionV2Dot2Integrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetIntegrationsOk() (*ServiceDefinitionV2Dot2Integrations, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given ServiceDefinitionV2Dot2Integrations and assigns it to the Integrations field.
func (o *ServiceDefinitionV2Dot2) SetIntegrations(v ServiceDefinitionV2Dot2Integrations) {
	o.Integrations = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetLanguages() []string {
	if o == nil || o.Languages == nil {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetLanguagesOk() (*[]string, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return &o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasLanguages() bool {
	return o != nil && o.Languages != nil
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *ServiceDefinitionV2Dot2) SetLanguages(v []string) {
	o.Languages = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *ServiceDefinitionV2Dot2) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetLinks() []ServiceDefinitionV2Dot2Link {
	if o == nil || o.Links == nil {
		var ret []ServiceDefinitionV2Dot2Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetLinksOk() (*[]ServiceDefinitionV2Dot2Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []ServiceDefinitionV2Dot2Link and assigns it to the Links field.
func (o *ServiceDefinitionV2Dot2) SetLinks(v []ServiceDefinitionV2Dot2Link) {
	o.Links = v
}

// GetSchemaVersion returns the SchemaVersion field value.
func (o *ServiceDefinitionV2Dot2) GetSchemaVersion() ServiceDefinitionV2Dot2Version {
	if o == nil {
		var ret ServiceDefinitionV2Dot2Version
		return ret
	}
	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetSchemaVersionOk() (*ServiceDefinitionV2Dot2Version, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value.
func (o *ServiceDefinitionV2Dot2) SetSchemaVersion(v ServiceDefinitionV2Dot2Version) {
	o.SchemaVersion = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceDefinitionV2Dot2) SetTags(v []string) {
	o.Tags = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *ServiceDefinitionV2Dot2) SetTeam(v string) {
	o.Team = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *ServiceDefinitionV2Dot2) SetTier(v string) {
	o.Tier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot2) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceDefinitionV2Dot2) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Dot2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.CiPipelineFingerprints != nil {
		toSerialize["ci-pipeline-fingerprints"] = o.CiPipelineFingerprints
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
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Dot2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Application            *string                              `json:"application,omitempty"`
		CiPipelineFingerprints []string                             `json:"ci-pipeline-fingerprints,omitempty"`
		Contacts               []ServiceDefinitionV2Dot2Contact     `json:"contacts,omitempty"`
		DdService              *string                              `json:"dd-service"`
		Description            *string                              `json:"description,omitempty"`
		Extensions             map[string]interface{}               `json:"extensions,omitempty"`
		Integrations           *ServiceDefinitionV2Dot2Integrations `json:"integrations,omitempty"`
		Languages              []string                             `json:"languages,omitempty"`
		Lifecycle              *string                              `json:"lifecycle,omitempty"`
		Links                  []ServiceDefinitionV2Dot2Link        `json:"links,omitempty"`
		SchemaVersion          *ServiceDefinitionV2Dot2Version      `json:"schema-version"`
		Tags                   []string                             `json:"tags,omitempty"`
		Team                   *string                              `json:"team,omitempty"`
		Tier                   *string                              `json:"tier,omitempty"`
		Type                   *string                              `json:"type,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"application", "ci-pipeline-fingerprints", "contacts", "dd-service", "description", "extensions", "integrations", "languages", "lifecycle", "links", "schema-version", "tags", "team", "tier", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Application = all.Application
	o.CiPipelineFingerprints = all.CiPipelineFingerprints
	o.Contacts = all.Contacts
	o.DdService = *all.DdService
	o.Description = all.Description
	o.Extensions = all.Extensions
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integrations = all.Integrations
	o.Languages = all.Languages
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
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
