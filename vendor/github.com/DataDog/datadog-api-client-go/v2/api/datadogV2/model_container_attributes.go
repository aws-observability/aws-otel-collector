// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerAttributes Attributes for a container.
type ContainerAttributes struct {
	// The ID of the container.
	ContainerId *string `json:"container_id,omitempty"`
	// Time the container was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Hostname of the host running the container.
	Host *string `json:"host,omitempty"`
	// Digest of the compressed image manifest.
	ImageDigest datadog.NullableString `json:"image_digest,omitempty"`
	// Name of the associated container image.
	ImageName *string `json:"image_name,omitempty"`
	// List of image tags associated with the container image.
	ImageTags datadog.NullableList[string] `json:"image_tags,omitempty"`
	// Name of the container.
	Name *string `json:"name,omitempty"`
	// Time the container was started.
	StartedAt *string `json:"started_at,omitempty"`
	// State of the container. This depends on the container runtime.
	State *string `json:"state,omitempty"`
	// List of tags associated with the container.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerAttributes instantiates a new ContainerAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerAttributes() *ContainerAttributes {
	this := ContainerAttributes{}
	return &this
}

// NewContainerAttributesWithDefaults instantiates a new ContainerAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerAttributesWithDefaults() *ContainerAttributes {
	this := ContainerAttributes{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *ContainerAttributes) GetContainerId() string {
	if o == nil || o.ContainerId == nil {
		var ret string
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetContainerIdOk() (*string, bool) {
	if o == nil || o.ContainerId == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *ContainerAttributes) HasContainerId() bool {
	return o != nil && o.ContainerId != nil
}

// SetContainerId gets a reference to the given string and assigns it to the ContainerId field.
func (o *ContainerAttributes) SetContainerId(v string) {
	o.ContainerId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContainerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContainerAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ContainerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ContainerAttributes) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ContainerAttributes) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ContainerAttributes) SetHost(v string) {
	o.Host = &v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerAttributes) GetImageDigest() string {
	if o == nil || o.ImageDigest.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest.Get()
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ContainerAttributes) GetImageDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageDigest.Get(), o.ImageDigest.IsSet()
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ContainerAttributes) HasImageDigest() bool {
	return o != nil && o.ImageDigest.IsSet()
}

// SetImageDigest gets a reference to the given datadog.NullableString and assigns it to the ImageDigest field.
func (o *ContainerAttributes) SetImageDigest(v string) {
	o.ImageDigest.Set(&v)
}

// SetImageDigestNil sets the value for ImageDigest to be an explicit nil.
func (o *ContainerAttributes) SetImageDigestNil() {
	o.ImageDigest.Set(nil)
}

// UnsetImageDigest ensures that no value is present for ImageDigest, not even an explicit nil.
func (o *ContainerAttributes) UnsetImageDigest() {
	o.ImageDigest.Unset()
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ContainerAttributes) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ContainerAttributes) HasImageName() bool {
	return o != nil && o.ImageName != nil
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ContainerAttributes) SetImageName(v string) {
	o.ImageName = &v
}

// GetImageTags returns the ImageTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerAttributes) GetImageTags() []string {
	if o == nil || o.ImageTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ImageTags.Get()
}

// GetImageTagsOk returns a tuple with the ImageTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ContainerAttributes) GetImageTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageTags.Get(), o.ImageTags.IsSet()
}

// HasImageTags returns a boolean if a field has been set.
func (o *ContainerAttributes) HasImageTags() bool {
	return o != nil && o.ImageTags.IsSet()
}

// SetImageTags gets a reference to the given datadog.NullableList[string] and assigns it to the ImageTags field.
func (o *ContainerAttributes) SetImageTags(v []string) {
	o.ImageTags.Set(&v)
}

// SetImageTagsNil sets the value for ImageTags to be an explicit nil.
func (o *ContainerAttributes) SetImageTagsNil() {
	o.ImageTags.Set(nil)
}

// UnsetImageTags ensures that no value is present for ImageTags, not even an explicit nil.
func (o *ContainerAttributes) UnsetImageTags() {
	o.ImageTags.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerAttributes) SetName(v string) {
	o.Name = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *ContainerAttributes) GetStartedAt() string {
	if o == nil || o.StartedAt == nil {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetStartedAtOk() (*string, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ContainerAttributes) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *ContainerAttributes) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ContainerAttributes) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ContainerAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ContainerAttributes) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContainerAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContainerAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ContainerAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContainerId != nil {
		toSerialize["container_id"] = o.ContainerId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.ImageDigest.IsSet() {
		toSerialize["image_digest"] = o.ImageDigest.Get()
	}
	if o.ImageName != nil {
		toSerialize["image_name"] = o.ImageName
	}
	if o.ImageTags.IsSet() {
		toSerialize["image_tags"] = o.ImageTags.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.State != nil {
		toSerialize["state"] = o.State
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
func (o *ContainerAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContainerId *string                      `json:"container_id,omitempty"`
		CreatedAt   *string                      `json:"created_at,omitempty"`
		Host        *string                      `json:"host,omitempty"`
		ImageDigest datadog.NullableString       `json:"image_digest,omitempty"`
		ImageName   *string                      `json:"image_name,omitempty"`
		ImageTags   datadog.NullableList[string] `json:"image_tags,omitempty"`
		Name        *string                      `json:"name,omitempty"`
		StartedAt   *string                      `json:"started_at,omitempty"`
		State       *string                      `json:"state,omitempty"`
		Tags        []string                     `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"container_id", "created_at", "host", "image_digest", "image_name", "image_tags", "name", "started_at", "state", "tags"})
	} else {
		return err
	}
	o.ContainerId = all.ContainerId
	o.CreatedAt = all.CreatedAt
	o.Host = all.Host
	o.ImageDigest = all.ImageDigest
	o.ImageName = all.ImageName
	o.ImageTags = all.ImageTags
	o.Name = all.Name
	o.StartedAt = all.StartedAt
	o.State = all.State
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
