// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageAttributes Attributes for a Container Image.
type ContainerImageAttributes struct {
	// Number of containers running the image.
	ContainerCount *int64 `json:"container_count,omitempty"`
	// List of platform-specific images associated with the image record.
	// The list contains more than 1 entry for multi-architecture images.
	ImageFlavors []ContainerImageFlavor `json:"image_flavors,omitempty"`
	// List of image tags associated with the Container Image.
	ImageTags []string `json:"image_tags,omitempty"`
	// List of build times associated with the Container Image.
	// The list contains more than 1 entry for multi-architecture images.
	ImagesBuiltAt []string `json:"images_built_at,omitempty"`
	// Name of the Container Image.
	Name *string `json:"name,omitempty"`
	// List of Operating System architectures supported by the Container Image.
	OsArchitectures []string `json:"os_architectures,omitempty"`
	// List of Operating System names supported by the Container Image.
	OsNames []string `json:"os_names,omitempty"`
	// List of Operating System versions supported by the Container Image.
	OsVersions []string `json:"os_versions,omitempty"`
	// Time the image was pushed to the container registry.
	PublishedAt *string `json:"published_at,omitempty"`
	// Registry the Container Image was pushed to.
	Registry *string `json:"registry,omitempty"`
	// Digest of the compressed image manifest.
	RepoDigest *string `json:"repo_digest,omitempty"`
	// Repository where the Container Image is stored in.
	Repository *string `json:"repository,omitempty"`
	// Short version of the Container Image name.
	ShortImage *string `json:"short_image,omitempty"`
	// List of size for each platform-specific image associated with the image record.
	// The list contains more than 1 entry for multi-architecture images.
	Sizes []int64 `json:"sizes,omitempty"`
	// List of sources where the Container Image was collected from.
	Sources []string `json:"sources,omitempty"`
	// List of tags associated with the Container Image.
	Tags []string `json:"tags,omitempty"`
	// Vulnerability counts associated with the Container Image.
	VulnerabilityCount *ContainerImageVulnerabilities `json:"vulnerability_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerImageAttributes instantiates a new ContainerImageAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageAttributes() *ContainerImageAttributes {
	this := ContainerImageAttributes{}
	return &this
}

// NewContainerImageAttributesWithDefaults instantiates a new ContainerImageAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageAttributesWithDefaults() *ContainerImageAttributes {
	this := ContainerImageAttributes{}
	return &this
}

// GetContainerCount returns the ContainerCount field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetContainerCount() int64 {
	if o == nil || o.ContainerCount == nil {
		var ret int64
		return ret
	}
	return *o.ContainerCount
}

// GetContainerCountOk returns a tuple with the ContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetContainerCountOk() (*int64, bool) {
	if o == nil || o.ContainerCount == nil {
		return nil, false
	}
	return o.ContainerCount, true
}

// HasContainerCount returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasContainerCount() bool {
	return o != nil && o.ContainerCount != nil
}

// SetContainerCount gets a reference to the given int64 and assigns it to the ContainerCount field.
func (o *ContainerImageAttributes) SetContainerCount(v int64) {
	o.ContainerCount = &v
}

// GetImageFlavors returns the ImageFlavors field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetImageFlavors() []ContainerImageFlavor {
	if o == nil || o.ImageFlavors == nil {
		var ret []ContainerImageFlavor
		return ret
	}
	return o.ImageFlavors
}

// GetImageFlavorsOk returns a tuple with the ImageFlavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetImageFlavorsOk() (*[]ContainerImageFlavor, bool) {
	if o == nil || o.ImageFlavors == nil {
		return nil, false
	}
	return &o.ImageFlavors, true
}

// HasImageFlavors returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasImageFlavors() bool {
	return o != nil && o.ImageFlavors != nil
}

// SetImageFlavors gets a reference to the given []ContainerImageFlavor and assigns it to the ImageFlavors field.
func (o *ContainerImageAttributes) SetImageFlavors(v []ContainerImageFlavor) {
	o.ImageFlavors = v
}

// GetImageTags returns the ImageTags field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetImageTags() []string {
	if o == nil || o.ImageTags == nil {
		var ret []string
		return ret
	}
	return o.ImageTags
}

// GetImageTagsOk returns a tuple with the ImageTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetImageTagsOk() (*[]string, bool) {
	if o == nil || o.ImageTags == nil {
		return nil, false
	}
	return &o.ImageTags, true
}

// HasImageTags returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasImageTags() bool {
	return o != nil && o.ImageTags != nil
}

// SetImageTags gets a reference to the given []string and assigns it to the ImageTags field.
func (o *ContainerImageAttributes) SetImageTags(v []string) {
	o.ImageTags = v
}

// GetImagesBuiltAt returns the ImagesBuiltAt field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetImagesBuiltAt() []string {
	if o == nil || o.ImagesBuiltAt == nil {
		var ret []string
		return ret
	}
	return o.ImagesBuiltAt
}

// GetImagesBuiltAtOk returns a tuple with the ImagesBuiltAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetImagesBuiltAtOk() (*[]string, bool) {
	if o == nil || o.ImagesBuiltAt == nil {
		return nil, false
	}
	return &o.ImagesBuiltAt, true
}

// HasImagesBuiltAt returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasImagesBuiltAt() bool {
	return o != nil && o.ImagesBuiltAt != nil
}

// SetImagesBuiltAt gets a reference to the given []string and assigns it to the ImagesBuiltAt field.
func (o *ContainerImageAttributes) SetImagesBuiltAt(v []string) {
	o.ImagesBuiltAt = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerImageAttributes) SetName(v string) {
	o.Name = &v
}

// GetOsArchitectures returns the OsArchitectures field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetOsArchitectures() []string {
	if o == nil || o.OsArchitectures == nil {
		var ret []string
		return ret
	}
	return o.OsArchitectures
}

// GetOsArchitecturesOk returns a tuple with the OsArchitectures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetOsArchitecturesOk() (*[]string, bool) {
	if o == nil || o.OsArchitectures == nil {
		return nil, false
	}
	return &o.OsArchitectures, true
}

// HasOsArchitectures returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasOsArchitectures() bool {
	return o != nil && o.OsArchitectures != nil
}

// SetOsArchitectures gets a reference to the given []string and assigns it to the OsArchitectures field.
func (o *ContainerImageAttributes) SetOsArchitectures(v []string) {
	o.OsArchitectures = v
}

// GetOsNames returns the OsNames field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetOsNames() []string {
	if o == nil || o.OsNames == nil {
		var ret []string
		return ret
	}
	return o.OsNames
}

// GetOsNamesOk returns a tuple with the OsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetOsNamesOk() (*[]string, bool) {
	if o == nil || o.OsNames == nil {
		return nil, false
	}
	return &o.OsNames, true
}

// HasOsNames returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasOsNames() bool {
	return o != nil && o.OsNames != nil
}

// SetOsNames gets a reference to the given []string and assigns it to the OsNames field.
func (o *ContainerImageAttributes) SetOsNames(v []string) {
	o.OsNames = v
}

// GetOsVersions returns the OsVersions field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetOsVersions() []string {
	if o == nil || o.OsVersions == nil {
		var ret []string
		return ret
	}
	return o.OsVersions
}

// GetOsVersionsOk returns a tuple with the OsVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetOsVersionsOk() (*[]string, bool) {
	if o == nil || o.OsVersions == nil {
		return nil, false
	}
	return &o.OsVersions, true
}

// HasOsVersions returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasOsVersions() bool {
	return o != nil && o.OsVersions != nil
}

// SetOsVersions gets a reference to the given []string and assigns it to the OsVersions field.
func (o *ContainerImageAttributes) SetOsVersions(v []string) {
	o.OsVersions = v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetPublishedAt() string {
	if o == nil || o.PublishedAt == nil {
		var ret string
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetPublishedAtOk() (*string, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasPublishedAt() bool {
	return o != nil && o.PublishedAt != nil
}

// SetPublishedAt gets a reference to the given string and assigns it to the PublishedAt field.
func (o *ContainerImageAttributes) SetPublishedAt(v string) {
	o.PublishedAt = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasRegistry() bool {
	return o != nil && o.Registry != nil
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ContainerImageAttributes) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepoDigest returns the RepoDigest field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetRepoDigest() string {
	if o == nil || o.RepoDigest == nil {
		var ret string
		return ret
	}
	return *o.RepoDigest
}

// GetRepoDigestOk returns a tuple with the RepoDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetRepoDigestOk() (*string, bool) {
	if o == nil || o.RepoDigest == nil {
		return nil, false
	}
	return o.RepoDigest, true
}

// HasRepoDigest returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasRepoDigest() bool {
	return o != nil && o.RepoDigest != nil
}

// SetRepoDigest gets a reference to the given string and assigns it to the RepoDigest field.
func (o *ContainerImageAttributes) SetRepoDigest(v string) {
	o.RepoDigest = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasRepository() bool {
	return o != nil && o.Repository != nil
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ContainerImageAttributes) SetRepository(v string) {
	o.Repository = &v
}

// GetShortImage returns the ShortImage field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetShortImage() string {
	if o == nil || o.ShortImage == nil {
		var ret string
		return ret
	}
	return *o.ShortImage
}

// GetShortImageOk returns a tuple with the ShortImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetShortImageOk() (*string, bool) {
	if o == nil || o.ShortImage == nil {
		return nil, false
	}
	return o.ShortImage, true
}

// HasShortImage returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasShortImage() bool {
	return o != nil && o.ShortImage != nil
}

// SetShortImage gets a reference to the given string and assigns it to the ShortImage field.
func (o *ContainerImageAttributes) SetShortImage(v string) {
	o.ShortImage = &v
}

// GetSizes returns the Sizes field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetSizes() []int64 {
	if o == nil || o.Sizes == nil {
		var ret []int64
		return ret
	}
	return o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetSizesOk() (*[]int64, bool) {
	if o == nil || o.Sizes == nil {
		return nil, false
	}
	return &o.Sizes, true
}

// HasSizes returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasSizes() bool {
	return o != nil && o.Sizes != nil
}

// SetSizes gets a reference to the given []int64 and assigns it to the Sizes field.
func (o *ContainerImageAttributes) SetSizes(v []int64) {
	o.Sizes = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetSourcesOk() (*[]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasSources() bool {
	return o != nil && o.Sources != nil
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *ContainerImageAttributes) SetSources(v []string) {
	o.Sources = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ContainerImageAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value if set, zero value otherwise.
func (o *ContainerImageAttributes) GetVulnerabilityCount() ContainerImageVulnerabilities {
	if o == nil || o.VulnerabilityCount == nil {
		var ret ContainerImageVulnerabilities
		return ret
	}
	return *o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageAttributes) GetVulnerabilityCountOk() (*ContainerImageVulnerabilities, bool) {
	if o == nil || o.VulnerabilityCount == nil {
		return nil, false
	}
	return o.VulnerabilityCount, true
}

// HasVulnerabilityCount returns a boolean if a field has been set.
func (o *ContainerImageAttributes) HasVulnerabilityCount() bool {
	return o != nil && o.VulnerabilityCount != nil
}

// SetVulnerabilityCount gets a reference to the given ContainerImageVulnerabilities and assigns it to the VulnerabilityCount field.
func (o *ContainerImageAttributes) SetVulnerabilityCount(v ContainerImageVulnerabilities) {
	o.VulnerabilityCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContainerCount != nil {
		toSerialize["container_count"] = o.ContainerCount
	}
	if o.ImageFlavors != nil {
		toSerialize["image_flavors"] = o.ImageFlavors
	}
	if o.ImageTags != nil {
		toSerialize["image_tags"] = o.ImageTags
	}
	if o.ImagesBuiltAt != nil {
		toSerialize["images_built_at"] = o.ImagesBuiltAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OsArchitectures != nil {
		toSerialize["os_architectures"] = o.OsArchitectures
	}
	if o.OsNames != nil {
		toSerialize["os_names"] = o.OsNames
	}
	if o.OsVersions != nil {
		toSerialize["os_versions"] = o.OsVersions
	}
	if o.PublishedAt != nil {
		toSerialize["published_at"] = o.PublishedAt
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.RepoDigest != nil {
		toSerialize["repo_digest"] = o.RepoDigest
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.ShortImage != nil {
		toSerialize["short_image"] = o.ShortImage
	}
	if o.Sizes != nil {
		toSerialize["sizes"] = o.Sizes
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.VulnerabilityCount != nil {
		toSerialize["vulnerability_count"] = o.VulnerabilityCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerImageAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContainerCount     *int64                         `json:"container_count,omitempty"`
		ImageFlavors       []ContainerImageFlavor         `json:"image_flavors,omitempty"`
		ImageTags          []string                       `json:"image_tags,omitempty"`
		ImagesBuiltAt      []string                       `json:"images_built_at,omitempty"`
		Name               *string                        `json:"name,omitempty"`
		OsArchitectures    []string                       `json:"os_architectures,omitempty"`
		OsNames            []string                       `json:"os_names,omitempty"`
		OsVersions         []string                       `json:"os_versions,omitempty"`
		PublishedAt        *string                        `json:"published_at,omitempty"`
		Registry           *string                        `json:"registry,omitempty"`
		RepoDigest         *string                        `json:"repo_digest,omitempty"`
		Repository         *string                        `json:"repository,omitempty"`
		ShortImage         *string                        `json:"short_image,omitempty"`
		Sizes              []int64                        `json:"sizes,omitempty"`
		Sources            []string                       `json:"sources,omitempty"`
		Tags               []string                       `json:"tags,omitempty"`
		VulnerabilityCount *ContainerImageVulnerabilities `json:"vulnerability_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"container_count", "image_flavors", "image_tags", "images_built_at", "name", "os_architectures", "os_names", "os_versions", "published_at", "registry", "repo_digest", "repository", "short_image", "sizes", "sources", "tags", "vulnerability_count"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ContainerCount = all.ContainerCount
	o.ImageFlavors = all.ImageFlavors
	o.ImageTags = all.ImageTags
	o.ImagesBuiltAt = all.ImagesBuiltAt
	o.Name = all.Name
	o.OsArchitectures = all.OsArchitectures
	o.OsNames = all.OsNames
	o.OsVersions = all.OsVersions
	o.PublishedAt = all.PublishedAt
	o.Registry = all.Registry
	o.RepoDigest = all.RepoDigest
	o.Repository = all.Repository
	o.ShortImage = all.ShortImage
	o.Sizes = all.Sizes
	o.Sources = all.Sources
	o.Tags = all.Tags
	if all.VulnerabilityCount != nil && all.VulnerabilityCount.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.VulnerabilityCount = all.VulnerabilityCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
