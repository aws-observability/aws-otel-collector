// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3Queue Schema for queue entities.
type EntityV3Queue struct {
	// The version of the schema data that was used to populate this entity's data. This could be via the API, Terraform, or YAML file in a repository. The field is known as schema-version in the previous version.
	ApiVersion EntityV3APIVersion `json:"apiVersion"`
	// Datadog product integrations for the datastore entity.
	Datadog *EntityV3QueueDatadog `json:"datadog,omitempty"`
	// Custom extensions. This is the free-formed field to send client-side metadata. No Datadog features are affected by this field.
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	// A base schema for defining third-party integrations.
	Integrations *EntityV3Integrations `json:"integrations,omitempty"`
	// The definition of Entity V3 Queue Kind object.
	Kind EntityV3QueueKind `json:"kind"`
	// The definition of Entity V3 Metadata object.
	Metadata EntityV3Metadata `json:"metadata"`
	// The definition of Entity V3 Queue Spec object.
	Spec *EntityV3QueueSpec `json:"spec,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3Queue instantiates a new EntityV3Queue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3Queue(apiVersion EntityV3APIVersion, kind EntityV3QueueKind, metadata EntityV3Metadata) *EntityV3Queue {
	this := EntityV3Queue{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewEntityV3QueueWithDefaults instantiates a new EntityV3Queue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3QueueWithDefaults() *EntityV3Queue {
	this := EntityV3Queue{}
	return &this
}

// GetApiVersion returns the ApiVersion field value.
func (o *EntityV3Queue) GetApiVersion() EntityV3APIVersion {
	if o == nil {
		var ret EntityV3APIVersion
		return ret
	}
	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetApiVersionOk() (*EntityV3APIVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value.
func (o *EntityV3Queue) SetApiVersion(v EntityV3APIVersion) {
	o.ApiVersion = v
}

// GetDatadog returns the Datadog field value if set, zero value otherwise.
func (o *EntityV3Queue) GetDatadog() EntityV3QueueDatadog {
	if o == nil || o.Datadog == nil {
		var ret EntityV3QueueDatadog
		return ret
	}
	return *o.Datadog
}

// GetDatadogOk returns a tuple with the Datadog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetDatadogOk() (*EntityV3QueueDatadog, bool) {
	if o == nil || o.Datadog == nil {
		return nil, false
	}
	return o.Datadog, true
}

// HasDatadog returns a boolean if a field has been set.
func (o *EntityV3Queue) HasDatadog() bool {
	return o != nil && o.Datadog != nil
}

// SetDatadog gets a reference to the given EntityV3QueueDatadog and assigns it to the Datadog field.
func (o *EntityV3Queue) SetDatadog(v EntityV3QueueDatadog) {
	o.Datadog = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *EntityV3Queue) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetExtensionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *EntityV3Queue) HasExtensions() bool {
	return o != nil && o.Extensions != nil
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *EntityV3Queue) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *EntityV3Queue) GetIntegrations() EntityV3Integrations {
	if o == nil || o.Integrations == nil {
		var ret EntityV3Integrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetIntegrationsOk() (*EntityV3Integrations, bool) {
	if o == nil || o.Integrations == nil {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *EntityV3Queue) HasIntegrations() bool {
	return o != nil && o.Integrations != nil
}

// SetIntegrations gets a reference to the given EntityV3Integrations and assigns it to the Integrations field.
func (o *EntityV3Queue) SetIntegrations(v EntityV3Integrations) {
	o.Integrations = &v
}

// GetKind returns the Kind field value.
func (o *EntityV3Queue) GetKind() EntityV3QueueKind {
	if o == nil {
		var ret EntityV3QueueKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetKindOk() (*EntityV3QueueKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *EntityV3Queue) SetKind(v EntityV3QueueKind) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value.
func (o *EntityV3Queue) GetMetadata() EntityV3Metadata {
	if o == nil {
		var ret EntityV3Metadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetMetadataOk() (*EntityV3Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *EntityV3Queue) SetMetadata(v EntityV3Metadata) {
	o.Metadata = v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *EntityV3Queue) GetSpec() EntityV3QueueSpec {
	if o == nil || o.Spec == nil {
		var ret EntityV3QueueSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3Queue) GetSpecOk() (*EntityV3QueueSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *EntityV3Queue) HasSpec() bool {
	return o != nil && o.Spec != nil
}

// SetSpec gets a reference to the given EntityV3QueueSpec and assigns it to the Spec field.
func (o *EntityV3Queue) SetSpec(v EntityV3QueueSpec) {
	o.Spec = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3Queue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["apiVersion"] = o.ApiVersion
	if o.Datadog != nil {
		toSerialize["datadog"] = o.Datadog
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3Queue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiVersion   *EntityV3APIVersion    `json:"apiVersion"`
		Datadog      *EntityV3QueueDatadog  `json:"datadog,omitempty"`
		Extensions   map[string]interface{} `json:"extensions,omitempty"`
		Integrations *EntityV3Integrations  `json:"integrations,omitempty"`
		Kind         *EntityV3QueueKind     `json:"kind"`
		Metadata     *EntityV3Metadata      `json:"metadata"`
		Spec         *EntityV3QueueSpec     `json:"spec,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiVersion == nil {
		return fmt.Errorf("required field apiVersion missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}

	hasInvalidField := false
	if !all.ApiVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.ApiVersion = *all.ApiVersion
	}
	if all.Datadog != nil && all.Datadog.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Datadog = all.Datadog
	o.Extensions = all.Extensions
	if all.Integrations != nil && all.Integrations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integrations = all.Integrations
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	if all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = *all.Metadata
	if all.Spec != nil && all.Spec.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Spec = all.Spec

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
