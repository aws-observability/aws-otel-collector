// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedUncompressedResource Aggregated uncompressed resource detection grouped by URL path.
type AggregatedUncompressedResource struct {
	// Average uncompressed body size in bytes.
	AvgBodySize int64 `json:"avg_body_size"`
	// Average resource loading duration in nanoseconds.
	AvgDuration int64 `json:"avg_duration"`
	// Unique fingerprint identifying this detection group.
	Fingerprint string `json:"fingerprint"`
	// Impact score combining view frequency and resource size.
	ImpactScore float64 `json:"impact_score"`
	// Total number of detection instances across sampled views.
	InstanceCount int32 `json:"instance_count"`
	// CDN or hosting provider type for the resource.
	ProviderType datadog.NullableString `json:"provider_type"`
	// Whether the resource is render-blocking.
	RenderBlocking datadog.NullableString `json:"render_blocking"`
	// Type of the resource (JS, CSS, image, fetch, and so on).
	ResourceType string `json:"resource_type"`
	// Normalized URL path pattern for the uncompressed resource.
	UrlPathGroup string `json:"url_path_group"`
	// Number of sampled views where this detection occurred.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedUncompressedResource instantiates a new AggregatedUncompressedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedUncompressedResource(avgBodySize int64, avgDuration int64, fingerprint string, impactScore float64, instanceCount int32, providerType datadog.NullableString, renderBlocking datadog.NullableString, resourceType string, urlPathGroup string, viewOccurrences int32) *AggregatedUncompressedResource {
	this := AggregatedUncompressedResource{}
	this.AvgBodySize = avgBodySize
	this.AvgDuration = avgDuration
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.InstanceCount = instanceCount
	this.ProviderType = providerType
	this.RenderBlocking = renderBlocking
	this.ResourceType = resourceType
	this.UrlPathGroup = urlPathGroup
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedUncompressedResourceWithDefaults instantiates a new AggregatedUncompressedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedUncompressedResourceWithDefaults() *AggregatedUncompressedResource {
	this := AggregatedUncompressedResource{}
	return &this
}

// GetAvgBodySize returns the AvgBodySize field value.
func (o *AggregatedUncompressedResource) GetAvgBodySize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgBodySize
}

// GetAvgBodySizeOk returns a tuple with the AvgBodySize field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetAvgBodySizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgBodySize, true
}

// SetAvgBodySize sets field value.
func (o *AggregatedUncompressedResource) SetAvgBodySize(v int64) {
	o.AvgBodySize = v
}

// GetAvgDuration returns the AvgDuration field value.
func (o *AggregatedUncompressedResource) GetAvgDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgDuration
}

// GetAvgDurationOk returns a tuple with the AvgDuration field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetAvgDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDuration, true
}

// SetAvgDuration sets field value.
func (o *AggregatedUncompressedResource) SetAvgDuration(v int64) {
	o.AvgDuration = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedUncompressedResource) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedUncompressedResource) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedUncompressedResource) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedUncompressedResource) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AggregatedUncompressedResource) GetInstanceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetInstanceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AggregatedUncompressedResource) SetInstanceCount(v int32) {
	o.InstanceCount = v
}

// GetProviderType returns the ProviderType field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedUncompressedResource) GetProviderType() string {
	if o == nil || o.ProviderType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProviderType.Get()
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedUncompressedResource) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderType.Get(), o.ProviderType.IsSet()
}

// SetProviderType sets field value.
func (o *AggregatedUncompressedResource) SetProviderType(v string) {
	o.ProviderType.Set(&v)
}

// GetRenderBlocking returns the RenderBlocking field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedUncompressedResource) GetRenderBlocking() string {
	if o == nil || o.RenderBlocking.Get() == nil {
		var ret string
		return ret
	}
	return *o.RenderBlocking.Get()
}

// GetRenderBlockingOk returns a tuple with the RenderBlocking field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedUncompressedResource) GetRenderBlockingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenderBlocking.Get(), o.RenderBlocking.IsSet()
}

// SetRenderBlocking sets field value.
func (o *AggregatedUncompressedResource) SetRenderBlocking(v string) {
	o.RenderBlocking.Set(&v)
}

// GetResourceType returns the ResourceType field value.
func (o *AggregatedUncompressedResource) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *AggregatedUncompressedResource) SetResourceType(v string) {
	o.ResourceType = v
}

// GetUrlPathGroup returns the UrlPathGroup field value.
func (o *AggregatedUncompressedResource) GetUrlPathGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UrlPathGroup
}

// GetUrlPathGroupOk returns a tuple with the UrlPathGroup field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetUrlPathGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlPathGroup, true
}

// SetUrlPathGroup sets field value.
func (o *AggregatedUncompressedResource) SetUrlPathGroup(v string) {
	o.UrlPathGroup = v
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedUncompressedResource) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedUncompressedResource) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedUncompressedResource) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedUncompressedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_body_size"] = o.AvgBodySize
	toSerialize["avg_duration"] = o.AvgDuration
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["provider_type"] = o.ProviderType.Get()
	toSerialize["render_blocking"] = o.RenderBlocking.Get()
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["url_path_group"] = o.UrlPathGroup
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedUncompressedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgBodySize     *int64                 `json:"avg_body_size"`
		AvgDuration     *int64                 `json:"avg_duration"`
		Fingerprint     *string                `json:"fingerprint"`
		ImpactScore     *float64               `json:"impact_score"`
		InstanceCount   *int32                 `json:"instance_count"`
		ProviderType    datadog.NullableString `json:"provider_type"`
		RenderBlocking  datadog.NullableString `json:"render_blocking"`
		ResourceType    *string                `json:"resource_type"`
		UrlPathGroup    *string                `json:"url_path_group"`
		ViewOccurrences *int32                 `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgBodySize == nil {
		return fmt.Errorf("required field avg_body_size missing")
	}
	if all.AvgDuration == nil {
		return fmt.Errorf("required field avg_duration missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.ImpactScore == nil {
		return fmt.Errorf("required field impact_score missing")
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if !all.ProviderType.IsSet() {
		return fmt.Errorf("required field provider_type missing")
	}
	if !all.RenderBlocking.IsSet() {
		return fmt.Errorf("required field render_blocking missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.UrlPathGroup == nil {
		return fmt.Errorf("required field url_path_group missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_body_size", "avg_duration", "fingerprint", "impact_score", "instance_count", "provider_type", "render_blocking", "resource_type", "url_path_group", "view_occurrences"})
	} else {
		return err
	}
	o.AvgBodySize = *all.AvgBodySize
	o.AvgDuration = *all.AvgDuration
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.InstanceCount = *all.InstanceCount
	o.ProviderType = all.ProviderType
	o.RenderBlocking = all.RenderBlocking
	o.ResourceType = *all.ResourceType
	o.UrlPathGroup = *all.UrlPathGroup
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
