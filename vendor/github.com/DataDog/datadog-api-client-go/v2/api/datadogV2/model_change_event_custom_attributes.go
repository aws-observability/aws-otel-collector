// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCustomAttributes Change event attributes.
type ChangeEventCustomAttributes struct {
	// The entity that made the change. Optional, if provided it must include `type` and `name`.
	Author *ChangeEventCustomAttributesAuthor `json:"author,omitempty"`
	// Free form JSON object with information related to the `change` event. Supports up to 100 properties per object and a maximum nesting depth of 10 levels.
	ChangeMetadata map[string]interface{} `json:"change_metadata,omitempty"`
	// A uniquely identified resource.
	ChangedResource ChangeEventCustomAttributesChangedResource `json:"changed_resource"`
	// A list of resources impacted by this change. It is recommended to provide an impacted resource to display
	// the change event at the correct location. Only resources of type `service` are supported. Maximum of 100 impacted resources allowed.
	ImpactedResources []ChangeEventCustomAttributesImpactedResourcesItems `json:"impacted_resources,omitempty"`
	// Free form JSON object representing the new state of the changed resource.
	NewValue map[string]interface{} `json:"new_value,omitempty"`
	// Free form JSON object representing the previous state of the changed resource.
	PrevValue map[string]interface{} `json:"prev_value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeEventCustomAttributes instantiates a new ChangeEventCustomAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventCustomAttributes(changedResource ChangeEventCustomAttributesChangedResource) *ChangeEventCustomAttributes {
	this := ChangeEventCustomAttributes{}
	this.ChangedResource = changedResource
	return &this
}

// NewChangeEventCustomAttributesWithDefaults instantiates a new ChangeEventCustomAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventCustomAttributesWithDefaults() *ChangeEventCustomAttributes {
	this := ChangeEventCustomAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ChangeEventCustomAttributes) GetAuthor() ChangeEventCustomAttributesAuthor {
	if o == nil || o.Author == nil {
		var ret ChangeEventCustomAttributesAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributes) GetAuthorOk() (*ChangeEventCustomAttributesAuthor, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ChangeEventCustomAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given ChangeEventCustomAttributesAuthor and assigns it to the Author field.
func (o *ChangeEventCustomAttributes) SetAuthor(v ChangeEventCustomAttributesAuthor) {
	o.Author = &v
}

// GetChangeMetadata returns the ChangeMetadata field value if set, zero value otherwise.
func (o *ChangeEventCustomAttributes) GetChangeMetadata() map[string]interface{} {
	if o == nil || o.ChangeMetadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ChangeMetadata
}

// GetChangeMetadataOk returns a tuple with the ChangeMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributes) GetChangeMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.ChangeMetadata == nil {
		return nil, false
	}
	return &o.ChangeMetadata, true
}

// HasChangeMetadata returns a boolean if a field has been set.
func (o *ChangeEventCustomAttributes) HasChangeMetadata() bool {
	return o != nil && o.ChangeMetadata != nil
}

// SetChangeMetadata gets a reference to the given map[string]interface{} and assigns it to the ChangeMetadata field.
func (o *ChangeEventCustomAttributes) SetChangeMetadata(v map[string]interface{}) {
	o.ChangeMetadata = v
}

// GetChangedResource returns the ChangedResource field value.
func (o *ChangeEventCustomAttributes) GetChangedResource() ChangeEventCustomAttributesChangedResource {
	if o == nil {
		var ret ChangeEventCustomAttributesChangedResource
		return ret
	}
	return o.ChangedResource
}

// GetChangedResourceOk returns a tuple with the ChangedResource field value
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributes) GetChangedResourceOk() (*ChangeEventCustomAttributesChangedResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangedResource, true
}

// SetChangedResource sets field value.
func (o *ChangeEventCustomAttributes) SetChangedResource(v ChangeEventCustomAttributesChangedResource) {
	o.ChangedResource = v
}

// GetImpactedResources returns the ImpactedResources field value if set, zero value otherwise.
func (o *ChangeEventCustomAttributes) GetImpactedResources() []ChangeEventCustomAttributesImpactedResourcesItems {
	if o == nil || o.ImpactedResources == nil {
		var ret []ChangeEventCustomAttributesImpactedResourcesItems
		return ret
	}
	return o.ImpactedResources
}

// GetImpactedResourcesOk returns a tuple with the ImpactedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributes) GetImpactedResourcesOk() (*[]ChangeEventCustomAttributesImpactedResourcesItems, bool) {
	if o == nil || o.ImpactedResources == nil {
		return nil, false
	}
	return &o.ImpactedResources, true
}

// HasImpactedResources returns a boolean if a field has been set.
func (o *ChangeEventCustomAttributes) HasImpactedResources() bool {
	return o != nil && o.ImpactedResources != nil
}

// SetImpactedResources gets a reference to the given []ChangeEventCustomAttributesImpactedResourcesItems and assigns it to the ImpactedResources field.
func (o *ChangeEventCustomAttributes) SetImpactedResources(v []ChangeEventCustomAttributesImpactedResourcesItems) {
	o.ImpactedResources = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *ChangeEventCustomAttributes) GetNewValue() map[string]interface{} {
	if o == nil || o.NewValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributes) GetNewValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *ChangeEventCustomAttributes) HasNewValue() bool {
	return o != nil && o.NewValue != nil
}

// SetNewValue gets a reference to the given map[string]interface{} and assigns it to the NewValue field.
func (o *ChangeEventCustomAttributes) SetNewValue(v map[string]interface{}) {
	o.NewValue = v
}

// GetPrevValue returns the PrevValue field value if set, zero value otherwise.
func (o *ChangeEventCustomAttributes) GetPrevValue() map[string]interface{} {
	if o == nil || o.PrevValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PrevValue
}

// GetPrevValueOk returns a tuple with the PrevValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventCustomAttributes) GetPrevValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.PrevValue == nil {
		return nil, false
	}
	return &o.PrevValue, true
}

// HasPrevValue returns a boolean if a field has been set.
func (o *ChangeEventCustomAttributes) HasPrevValue() bool {
	return o != nil && o.PrevValue != nil
}

// SetPrevValue gets a reference to the given map[string]interface{} and assigns it to the PrevValue field.
func (o *ChangeEventCustomAttributes) SetPrevValue(v map[string]interface{}) {
	o.PrevValue = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventCustomAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.ChangeMetadata != nil {
		toSerialize["change_metadata"] = o.ChangeMetadata
	}
	toSerialize["changed_resource"] = o.ChangedResource
	if o.ImpactedResources != nil {
		toSerialize["impacted_resources"] = o.ImpactedResources
	}
	if o.NewValue != nil {
		toSerialize["new_value"] = o.NewValue
	}
	if o.PrevValue != nil {
		toSerialize["prev_value"] = o.PrevValue
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEventCustomAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author            *ChangeEventCustomAttributesAuthor                  `json:"author,omitempty"`
		ChangeMetadata    map[string]interface{}                              `json:"change_metadata,omitempty"`
		ChangedResource   *ChangeEventCustomAttributesChangedResource         `json:"changed_resource"`
		ImpactedResources []ChangeEventCustomAttributesImpactedResourcesItems `json:"impacted_resources,omitempty"`
		NewValue          map[string]interface{}                              `json:"new_value,omitempty"`
		PrevValue         map[string]interface{}                              `json:"prev_value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChangedResource == nil {
		return fmt.Errorf("required field changed_resource missing")
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	o.ChangeMetadata = all.ChangeMetadata
	if all.ChangedResource.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChangedResource = *all.ChangedResource
	o.ImpactedResources = all.ImpactedResources
	o.NewValue = all.NewValue
	o.PrevValue = all.PrevValue

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
