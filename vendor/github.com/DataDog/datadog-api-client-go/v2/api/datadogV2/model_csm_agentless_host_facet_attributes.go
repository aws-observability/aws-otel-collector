// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmAgentlessHostFacetAttributes Attributes of an agentless host facet.
type CsmAgentlessHostFacetAttributes struct {
	// Whether the facet has a bounded set of allowed values. `true` indicates a fixed value set and `false` indicates free-form values.
	Bounded bool `json:"bounded"`
	// Whether the facet is bundled as part of the default facet set. `true` indicates bundled and `false` indicates custom.
	Bundled bool `json:"bundled"`
	// Whether the facet is both bundled and actively used. `true` indicates in use; `false` indicates unused.
	BundledAndUsed bool `json:"bundledAndUsed"`
	// The list of default filter values for the facet.
	DefaultValues []string `json:"defaultValues"`
	// A human-readable description of what the facet represents.
	Description string `json:"description"`
	// Whether the facet can be edited by users. `true` indicates editable; `false` indicates read-only.
	Editable bool `json:"editable"`
	// The UI display type for the facet, such as `list`.
	FacetType string `json:"facetType"`
	// The list of UI groups that this facet belongs to.
	Groups []string `json:"groups"`
	// The display name of the facet.
	Name string `json:"name"`
	// The field path used when filtering by this facet.
	Path string `json:"path"`
	// The data source that provides the facet values.
	Source string `json:"source"`
	// The data type of the facet values.
	Type string `json:"type"`
	// The list of allowed filter values for bounded facets. Empty for unbounded facets.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmAgentlessHostFacetAttributes instantiates a new CsmAgentlessHostFacetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmAgentlessHostFacetAttributes(bounded bool, bundled bool, bundledAndUsed bool, defaultValues []string, description string, editable bool, facetType string, groups []string, name string, path string, source string, typeVar string, values []string) *CsmAgentlessHostFacetAttributes {
	this := CsmAgentlessHostFacetAttributes{}
	this.Bounded = bounded
	this.Bundled = bundled
	this.BundledAndUsed = bundledAndUsed
	this.DefaultValues = defaultValues
	this.Description = description
	this.Editable = editable
	this.FacetType = facetType
	this.Groups = groups
	this.Name = name
	this.Path = path
	this.Source = source
	this.Type = typeVar
	this.Values = values
	return &this
}

// NewCsmAgentlessHostFacetAttributesWithDefaults instantiates a new CsmAgentlessHostFacetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmAgentlessHostFacetAttributesWithDefaults() *CsmAgentlessHostFacetAttributes {
	this := CsmAgentlessHostFacetAttributes{}
	return &this
}

// GetBounded returns the Bounded field value.
func (o *CsmAgentlessHostFacetAttributes) GetBounded() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Bounded
}

// GetBoundedOk returns a tuple with the Bounded field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetBoundedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bounded, true
}

// SetBounded sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetBounded(v bool) {
	o.Bounded = v
}

// GetBundled returns the Bundled field value.
func (o *CsmAgentlessHostFacetAttributes) GetBundled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Bundled
}

// GetBundledOk returns a tuple with the Bundled field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetBundledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bundled, true
}

// SetBundled sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetBundled(v bool) {
	o.Bundled = v
}

// GetBundledAndUsed returns the BundledAndUsed field value.
func (o *CsmAgentlessHostFacetAttributes) GetBundledAndUsed() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.BundledAndUsed
}

// GetBundledAndUsedOk returns a tuple with the BundledAndUsed field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetBundledAndUsedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundledAndUsed, true
}

// SetBundledAndUsed sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetBundledAndUsed(v bool) {
	o.BundledAndUsed = v
}

// GetDefaultValues returns the DefaultValues field value.
func (o *CsmAgentlessHostFacetAttributes) GetDefaultValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetDefaultValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValues, true
}

// SetDefaultValues sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetDefaultValues(v []string) {
	o.DefaultValues = v
}

// GetDescription returns the Description field value.
func (o *CsmAgentlessHostFacetAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEditable returns the Editable field value.
func (o *CsmAgentlessHostFacetAttributes) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetEditable(v bool) {
	o.Editable = v
}

// GetFacetType returns the FacetType field value.
func (o *CsmAgentlessHostFacetAttributes) GetFacetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FacetType
}

// GetFacetTypeOk returns a tuple with the FacetType field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetFacetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetType, true
}

// SetFacetType sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetFacetType(v string) {
	o.FacetType = v
}

// GetGroups returns the Groups field value.
func (o *CsmAgentlessHostFacetAttributes) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetGroupsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetGroups(v []string) {
	o.Groups = v
}

// GetName returns the Name field value.
func (o *CsmAgentlessHostFacetAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value.
func (o *CsmAgentlessHostFacetAttributes) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetPath(v string) {
	o.Path = v
}

// GetSource returns the Source field value.
func (o *CsmAgentlessHostFacetAttributes) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value.
func (o *CsmAgentlessHostFacetAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value.
func (o *CsmAgentlessHostFacetAttributes) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CsmAgentlessHostFacetAttributes) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CsmAgentlessHostFacetAttributes) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmAgentlessHostFacetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bounded"] = o.Bounded
	toSerialize["bundled"] = o.Bundled
	toSerialize["bundledAndUsed"] = o.BundledAndUsed
	toSerialize["defaultValues"] = o.DefaultValues
	toSerialize["description"] = o.Description
	toSerialize["editable"] = o.Editable
	toSerialize["facetType"] = o.FacetType
	toSerialize["groups"] = o.Groups
	toSerialize["name"] = o.Name
	toSerialize["path"] = o.Path
	toSerialize["source"] = o.Source
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmAgentlessHostFacetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bounded        *bool     `json:"bounded"`
		Bundled        *bool     `json:"bundled"`
		BundledAndUsed *bool     `json:"bundledAndUsed"`
		DefaultValues  *[]string `json:"defaultValues"`
		Description    *string   `json:"description"`
		Editable       *bool     `json:"editable"`
		FacetType      *string   `json:"facetType"`
		Groups         *[]string `json:"groups"`
		Name           *string   `json:"name"`
		Path           *string   `json:"path"`
		Source         *string   `json:"source"`
		Type           *string   `json:"type"`
		Values         *[]string `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bounded == nil {
		return fmt.Errorf("required field bounded missing")
	}
	if all.Bundled == nil {
		return fmt.Errorf("required field bundled missing")
	}
	if all.BundledAndUsed == nil {
		return fmt.Errorf("required field bundledAndUsed missing")
	}
	if all.DefaultValues == nil {
		return fmt.Errorf("required field defaultValues missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Editable == nil {
		return fmt.Errorf("required field editable missing")
	}
	if all.FacetType == nil {
		return fmt.Errorf("required field facetType missing")
	}
	if all.Groups == nil {
		return fmt.Errorf("required field groups missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bounded", "bundled", "bundledAndUsed", "defaultValues", "description", "editable", "facetType", "groups", "name", "path", "source", "type", "values"})
	} else {
		return err
	}
	o.Bounded = *all.Bounded
	o.Bundled = *all.Bundled
	o.BundledAndUsed = *all.BundledAndUsed
	o.DefaultValues = *all.DefaultValues
	o.Description = *all.Description
	o.Editable = *all.Editable
	o.FacetType = *all.FacetType
	o.Groups = *all.Groups
	o.Name = *all.Name
	o.Path = *all.Path
	o.Source = *all.Source
	o.Type = *all.Type
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
