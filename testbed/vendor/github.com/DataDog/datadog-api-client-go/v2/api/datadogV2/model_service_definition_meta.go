// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionMeta Metadata about a service definition.
type ServiceDefinitionMeta struct {
	// GitHub HTML URL.
	GithubHtmlUrl *string `json:"github-html-url,omitempty"`
	// Ingestion schema version.
	IngestedSchemaVersion *string `json:"ingested-schema-version,omitempty"`
	// Ingestion source of the service definition.
	IngestionSource *string `json:"ingestion-source,omitempty"`
	// Last modified time of the service definition.
	LastModifiedTime *string `json:"last-modified-time,omitempty"`
	// User defined origin of the service definition.
	Origin *string `json:"origin,omitempty"`
	// User defined origin's detail of the service definition.
	OriginDetail *string `json:"origin-detail,omitempty"`
	// A list of schema validation warnings.
	Warnings []ServiceDefinitionMetaWarnings `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionMeta instantiates a new ServiceDefinitionMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionMeta() *ServiceDefinitionMeta {
	this := ServiceDefinitionMeta{}
	return &this
}

// NewServiceDefinitionMetaWithDefaults instantiates a new ServiceDefinitionMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionMetaWithDefaults() *ServiceDefinitionMeta {
	this := ServiceDefinitionMeta{}
	return &this
}

// GetGithubHtmlUrl returns the GithubHtmlUrl field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetGithubHtmlUrl() string {
	if o == nil || o.GithubHtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.GithubHtmlUrl
}

// GetGithubHtmlUrlOk returns a tuple with the GithubHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetGithubHtmlUrlOk() (*string, bool) {
	if o == nil || o.GithubHtmlUrl == nil {
		return nil, false
	}
	return o.GithubHtmlUrl, true
}

// HasGithubHtmlUrl returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasGithubHtmlUrl() bool {
	return o != nil && o.GithubHtmlUrl != nil
}

// SetGithubHtmlUrl gets a reference to the given string and assigns it to the GithubHtmlUrl field.
func (o *ServiceDefinitionMeta) SetGithubHtmlUrl(v string) {
	o.GithubHtmlUrl = &v
}

// GetIngestedSchemaVersion returns the IngestedSchemaVersion field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetIngestedSchemaVersion() string {
	if o == nil || o.IngestedSchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.IngestedSchemaVersion
}

// GetIngestedSchemaVersionOk returns a tuple with the IngestedSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetIngestedSchemaVersionOk() (*string, bool) {
	if o == nil || o.IngestedSchemaVersion == nil {
		return nil, false
	}
	return o.IngestedSchemaVersion, true
}

// HasIngestedSchemaVersion returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasIngestedSchemaVersion() bool {
	return o != nil && o.IngestedSchemaVersion != nil
}

// SetIngestedSchemaVersion gets a reference to the given string and assigns it to the IngestedSchemaVersion field.
func (o *ServiceDefinitionMeta) SetIngestedSchemaVersion(v string) {
	o.IngestedSchemaVersion = &v
}

// GetIngestionSource returns the IngestionSource field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetIngestionSource() string {
	if o == nil || o.IngestionSource == nil {
		var ret string
		return ret
	}
	return *o.IngestionSource
}

// GetIngestionSourceOk returns a tuple with the IngestionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetIngestionSourceOk() (*string, bool) {
	if o == nil || o.IngestionSource == nil {
		return nil, false
	}
	return o.IngestionSource, true
}

// HasIngestionSource returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasIngestionSource() bool {
	return o != nil && o.IngestionSource != nil
}

// SetIngestionSource gets a reference to the given string and assigns it to the IngestionSource field.
func (o *ServiceDefinitionMeta) SetIngestionSource(v string) {
	o.IngestionSource = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetLastModifiedTime() string {
	if o == nil || o.LastModifiedTime == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetLastModifiedTimeOk() (*string, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasLastModifiedTime() bool {
	return o != nil && o.LastModifiedTime != nil
}

// SetLastModifiedTime gets a reference to the given string and assigns it to the LastModifiedTime field.
func (o *ServiceDefinitionMeta) SetLastModifiedTime(v string) {
	o.LastModifiedTime = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasOrigin() bool {
	return o != nil && o.Origin != nil
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *ServiceDefinitionMeta) SetOrigin(v string) {
	o.Origin = &v
}

// GetOriginDetail returns the OriginDetail field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetOriginDetail() string {
	if o == nil || o.OriginDetail == nil {
		var ret string
		return ret
	}
	return *o.OriginDetail
}

// GetOriginDetailOk returns a tuple with the OriginDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetOriginDetailOk() (*string, bool) {
	if o == nil || o.OriginDetail == nil {
		return nil, false
	}
	return o.OriginDetail, true
}

// HasOriginDetail returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasOriginDetail() bool {
	return o != nil && o.OriginDetail != nil
}

// SetOriginDetail gets a reference to the given string and assigns it to the OriginDetail field.
func (o *ServiceDefinitionMeta) SetOriginDetail(v string) {
	o.OriginDetail = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ServiceDefinitionMeta) GetWarnings() []ServiceDefinitionMetaWarnings {
	if o == nil || o.Warnings == nil {
		var ret []ServiceDefinitionMetaWarnings
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionMeta) GetWarningsOk() (*[]ServiceDefinitionMetaWarnings, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ServiceDefinitionMeta) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []ServiceDefinitionMetaWarnings and assigns it to the Warnings field.
func (o *ServiceDefinitionMeta) SetWarnings(v []ServiceDefinitionMetaWarnings) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GithubHtmlUrl != nil {
		toSerialize["github-html-url"] = o.GithubHtmlUrl
	}
	if o.IngestedSchemaVersion != nil {
		toSerialize["ingested-schema-version"] = o.IngestedSchemaVersion
	}
	if o.IngestionSource != nil {
		toSerialize["ingestion-source"] = o.IngestionSource
	}
	if o.LastModifiedTime != nil {
		toSerialize["last-modified-time"] = o.LastModifiedTime
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.OriginDetail != nil {
		toSerialize["origin-detail"] = o.OriginDetail
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GithubHtmlUrl         *string                         `json:"github-html-url,omitempty"`
		IngestedSchemaVersion *string                         `json:"ingested-schema-version,omitempty"`
		IngestionSource       *string                         `json:"ingestion-source,omitempty"`
		LastModifiedTime      *string                         `json:"last-modified-time,omitempty"`
		Origin                *string                         `json:"origin,omitempty"`
		OriginDetail          *string                         `json:"origin-detail,omitempty"`
		Warnings              []ServiceDefinitionMetaWarnings `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"github-html-url", "ingested-schema-version", "ingestion-source", "last-modified-time", "origin", "origin-detail", "warnings"})
	} else {
		return err
	}
	o.GithubHtmlUrl = all.GithubHtmlUrl
	o.IngestedSchemaVersion = all.IngestedSchemaVersion
	o.IngestionSource = all.IngestionSource
	o.LastModifiedTime = all.LastModifiedTime
	o.Origin = all.Origin
	o.OriginDetail = all.OriginDetail
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
