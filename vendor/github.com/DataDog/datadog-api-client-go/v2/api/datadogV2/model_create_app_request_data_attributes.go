// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppRequestDataAttributes The definition of `CreateAppRequestDataAttributes` object.
type CreateAppRequestDataAttributes struct {
	// The `attributes` `components`.
	Components []ComponentGrid `json:"components,omitempty"`
	// The `attributes` `description`.
	Description *string `json:"description,omitempty"`
	// The `attributes` `embeddedQueries`.
	EmbeddedQueries []Query `json:"embeddedQueries,omitempty"`
	// The definition of `InputSchema` object.
	InputSchema *InputSchema `json:"inputSchema,omitempty"`
	// The `attributes` `name`.
	Name *string `json:"name,omitempty"`
	// The `attributes` `rootInstanceName`.
	RootInstanceName *string `json:"rootInstanceName,omitempty"`
	// The `attributes` `scripts`.
	Scripts []Script `json:"scripts,omitempty"`
	// The `attributes` `tags`.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppRequestDataAttributes instantiates a new CreateAppRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppRequestDataAttributes() *CreateAppRequestDataAttributes {
	this := CreateAppRequestDataAttributes{}
	return &this
}

// NewCreateAppRequestDataAttributesWithDefaults instantiates a new CreateAppRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppRequestDataAttributesWithDefaults() *CreateAppRequestDataAttributes {
	this := CreateAppRequestDataAttributes{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetComponents() []ComponentGrid {
	if o == nil || o.Components == nil {
		var ret []ComponentGrid
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetComponentsOk() (*[]ComponentGrid, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []ComponentGrid and assigns it to the Components field.
func (o *CreateAppRequestDataAttributes) SetComponents(v []ComponentGrid) {
	o.Components = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAppRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEmbeddedQueries returns the EmbeddedQueries field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetEmbeddedQueries() []Query {
	if o == nil || o.EmbeddedQueries == nil {
		var ret []Query
		return ret
	}
	return o.EmbeddedQueries
}

// GetEmbeddedQueriesOk returns a tuple with the EmbeddedQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetEmbeddedQueriesOk() (*[]Query, bool) {
	if o == nil || o.EmbeddedQueries == nil {
		return nil, false
	}
	return &o.EmbeddedQueries, true
}

// HasEmbeddedQueries returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasEmbeddedQueries() bool {
	return o != nil && o.EmbeddedQueries != nil
}

// SetEmbeddedQueries gets a reference to the given []Query and assigns it to the EmbeddedQueries field.
func (o *CreateAppRequestDataAttributes) SetEmbeddedQueries(v []Query) {
	o.EmbeddedQueries = v
}

// GetInputSchema returns the InputSchema field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetInputSchema() InputSchema {
	if o == nil || o.InputSchema == nil {
		var ret InputSchema
		return ret
	}
	return *o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetInputSchemaOk() (*InputSchema, bool) {
	if o == nil || o.InputSchema == nil {
		return nil, false
	}
	return o.InputSchema, true
}

// HasInputSchema returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasInputSchema() bool {
	return o != nil && o.InputSchema != nil
}

// SetInputSchema gets a reference to the given InputSchema and assigns it to the InputSchema field.
func (o *CreateAppRequestDataAttributes) SetInputSchema(v InputSchema) {
	o.InputSchema = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateAppRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRootInstanceName returns the RootInstanceName field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetRootInstanceName() string {
	if o == nil || o.RootInstanceName == nil {
		var ret string
		return ret
	}
	return *o.RootInstanceName
}

// GetRootInstanceNameOk returns a tuple with the RootInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetRootInstanceNameOk() (*string, bool) {
	if o == nil || o.RootInstanceName == nil {
		return nil, false
	}
	return o.RootInstanceName, true
}

// HasRootInstanceName returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasRootInstanceName() bool {
	return o != nil && o.RootInstanceName != nil
}

// SetRootInstanceName gets a reference to the given string and assigns it to the RootInstanceName field.
func (o *CreateAppRequestDataAttributes) SetRootInstanceName(v string) {
	o.RootInstanceName = &v
}

// GetScripts returns the Scripts field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetScripts() []Script {
	if o == nil || o.Scripts == nil {
		var ret []Script
		return ret
	}
	return o.Scripts
}

// GetScriptsOk returns a tuple with the Scripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetScriptsOk() (*[]Script, bool) {
	if o == nil || o.Scripts == nil {
		return nil, false
	}
	return &o.Scripts, true
}

// HasScripts returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasScripts() bool {
	return o != nil && o.Scripts != nil
}

// SetScripts gets a reference to the given []Script and assigns it to the Scripts field.
func (o *CreateAppRequestDataAttributes) SetScripts(v []Script) {
	o.Scripts = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateAppRequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateAppRequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateAppRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EmbeddedQueries != nil {
		toSerialize["embeddedQueries"] = o.EmbeddedQueries
	}
	if o.InputSchema != nil {
		toSerialize["inputSchema"] = o.InputSchema
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RootInstanceName != nil {
		toSerialize["rootInstanceName"] = o.RootInstanceName
	}
	if o.Scripts != nil {
		toSerialize["scripts"] = o.Scripts
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAppRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components       []ComponentGrid `json:"components,omitempty"`
		Description      *string         `json:"description,omitempty"`
		EmbeddedQueries  []Query         `json:"embeddedQueries,omitempty"`
		InputSchema      *InputSchema    `json:"inputSchema,omitempty"`
		Name             *string         `json:"name,omitempty"`
		RootInstanceName *string         `json:"rootInstanceName,omitempty"`
		Scripts          []Script        `json:"scripts,omitempty"`
		Tags             []string        `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components", "description", "embeddedQueries", "inputSchema", "name", "rootInstanceName", "scripts", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Components = all.Components
	o.Description = all.Description
	o.EmbeddedQueries = all.EmbeddedQueries
	if all.InputSchema != nil && all.InputSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InputSchema = all.InputSchema
	o.Name = all.Name
	o.RootInstanceName = all.RootInstanceName
	o.Scripts = all.Scripts
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
