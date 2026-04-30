// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetInstrumentedPodGroupAttributes Attributes of a group of instrumented pods targeted for SSI injection.
type FleetInstrumentedPodGroupAttributes struct {
	// The SSI injection target configuration applied to the pod group.
	AppliedTarget map[string]interface{} `json:"applied_target,omitempty"`
	// The name of the applied SSI injection target.
	AppliedTargetName *string `json:"applied_target_name,omitempty"`
	// Tags injected into the pods by the Admission Controller.
	InjectedTags []string `json:"injected_tags,omitempty"`
	// The kind of the Kubernetes owner reference.
	KubeOwnerrefKind *string `json:"kube_ownerref_kind,omitempty"`
	// The name of the Kubernetes owner reference (deployment, statefulset, etc.).
	KubeOwnerrefName *string `json:"kube_ownerref_name,omitempty"`
	// Library injection annotations on the pod group.
	LibInjectionAnnotations []string `json:"lib_injection_annotations,omitempty"`
	// The Kubernetes namespace of the pod group.
	Namespace *string `json:"namespace,omitempty"`
	// Total number of pods in the group.
	PodCount *int64 `json:"pod_count,omitempty"`
	// Names of the individual pods in the group.
	PodNames []string `json:"pod_names,omitempty"`
	// Additional tags associated with the pod group.
	Tags map[string]string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetInstrumentedPodGroupAttributes instantiates a new FleetInstrumentedPodGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetInstrumentedPodGroupAttributes() *FleetInstrumentedPodGroupAttributes {
	this := FleetInstrumentedPodGroupAttributes{}
	return &this
}

// NewFleetInstrumentedPodGroupAttributesWithDefaults instantiates a new FleetInstrumentedPodGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetInstrumentedPodGroupAttributesWithDefaults() *FleetInstrumentedPodGroupAttributes {
	this := FleetInstrumentedPodGroupAttributes{}
	return &this
}

// GetAppliedTarget returns the AppliedTarget field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetAppliedTarget() map[string]interface{} {
	if o == nil || o.AppliedTarget == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AppliedTarget
}

// GetAppliedTargetOk returns a tuple with the AppliedTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetAppliedTargetOk() (*map[string]interface{}, bool) {
	if o == nil || o.AppliedTarget == nil {
		return nil, false
	}
	return &o.AppliedTarget, true
}

// HasAppliedTarget returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasAppliedTarget() bool {
	return o != nil && o.AppliedTarget != nil
}

// SetAppliedTarget gets a reference to the given map[string]interface{} and assigns it to the AppliedTarget field.
func (o *FleetInstrumentedPodGroupAttributes) SetAppliedTarget(v map[string]interface{}) {
	o.AppliedTarget = v
}

// GetAppliedTargetName returns the AppliedTargetName field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetAppliedTargetName() string {
	if o == nil || o.AppliedTargetName == nil {
		var ret string
		return ret
	}
	return *o.AppliedTargetName
}

// GetAppliedTargetNameOk returns a tuple with the AppliedTargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetAppliedTargetNameOk() (*string, bool) {
	if o == nil || o.AppliedTargetName == nil {
		return nil, false
	}
	return o.AppliedTargetName, true
}

// HasAppliedTargetName returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasAppliedTargetName() bool {
	return o != nil && o.AppliedTargetName != nil
}

// SetAppliedTargetName gets a reference to the given string and assigns it to the AppliedTargetName field.
func (o *FleetInstrumentedPodGroupAttributes) SetAppliedTargetName(v string) {
	o.AppliedTargetName = &v
}

// GetInjectedTags returns the InjectedTags field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetInjectedTags() []string {
	if o == nil || o.InjectedTags == nil {
		var ret []string
		return ret
	}
	return o.InjectedTags
}

// GetInjectedTagsOk returns a tuple with the InjectedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetInjectedTagsOk() (*[]string, bool) {
	if o == nil || o.InjectedTags == nil {
		return nil, false
	}
	return &o.InjectedTags, true
}

// HasInjectedTags returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasInjectedTags() bool {
	return o != nil && o.InjectedTags != nil
}

// SetInjectedTags gets a reference to the given []string and assigns it to the InjectedTags field.
func (o *FleetInstrumentedPodGroupAttributes) SetInjectedTags(v []string) {
	o.InjectedTags = v
}

// GetKubeOwnerrefKind returns the KubeOwnerrefKind field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetKubeOwnerrefKind() string {
	if o == nil || o.KubeOwnerrefKind == nil {
		var ret string
		return ret
	}
	return *o.KubeOwnerrefKind
}

// GetKubeOwnerrefKindOk returns a tuple with the KubeOwnerrefKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetKubeOwnerrefKindOk() (*string, bool) {
	if o == nil || o.KubeOwnerrefKind == nil {
		return nil, false
	}
	return o.KubeOwnerrefKind, true
}

// HasKubeOwnerrefKind returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasKubeOwnerrefKind() bool {
	return o != nil && o.KubeOwnerrefKind != nil
}

// SetKubeOwnerrefKind gets a reference to the given string and assigns it to the KubeOwnerrefKind field.
func (o *FleetInstrumentedPodGroupAttributes) SetKubeOwnerrefKind(v string) {
	o.KubeOwnerrefKind = &v
}

// GetKubeOwnerrefName returns the KubeOwnerrefName field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetKubeOwnerrefName() string {
	if o == nil || o.KubeOwnerrefName == nil {
		var ret string
		return ret
	}
	return *o.KubeOwnerrefName
}

// GetKubeOwnerrefNameOk returns a tuple with the KubeOwnerrefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetKubeOwnerrefNameOk() (*string, bool) {
	if o == nil || o.KubeOwnerrefName == nil {
		return nil, false
	}
	return o.KubeOwnerrefName, true
}

// HasKubeOwnerrefName returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasKubeOwnerrefName() bool {
	return o != nil && o.KubeOwnerrefName != nil
}

// SetKubeOwnerrefName gets a reference to the given string and assigns it to the KubeOwnerrefName field.
func (o *FleetInstrumentedPodGroupAttributes) SetKubeOwnerrefName(v string) {
	o.KubeOwnerrefName = &v
}

// GetLibInjectionAnnotations returns the LibInjectionAnnotations field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetLibInjectionAnnotations() []string {
	if o == nil || o.LibInjectionAnnotations == nil {
		var ret []string
		return ret
	}
	return o.LibInjectionAnnotations
}

// GetLibInjectionAnnotationsOk returns a tuple with the LibInjectionAnnotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetLibInjectionAnnotationsOk() (*[]string, bool) {
	if o == nil || o.LibInjectionAnnotations == nil {
		return nil, false
	}
	return &o.LibInjectionAnnotations, true
}

// HasLibInjectionAnnotations returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasLibInjectionAnnotations() bool {
	return o != nil && o.LibInjectionAnnotations != nil
}

// SetLibInjectionAnnotations gets a reference to the given []string and assigns it to the LibInjectionAnnotations field.
func (o *FleetInstrumentedPodGroupAttributes) SetLibInjectionAnnotations(v []string) {
	o.LibInjectionAnnotations = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *FleetInstrumentedPodGroupAttributes) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPodCount returns the PodCount field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetPodCount() int64 {
	if o == nil || o.PodCount == nil {
		var ret int64
		return ret
	}
	return *o.PodCount
}

// GetPodCountOk returns a tuple with the PodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetPodCountOk() (*int64, bool) {
	if o == nil || o.PodCount == nil {
		return nil, false
	}
	return o.PodCount, true
}

// HasPodCount returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasPodCount() bool {
	return o != nil && o.PodCount != nil
}

// SetPodCount gets a reference to the given int64 and assigns it to the PodCount field.
func (o *FleetInstrumentedPodGroupAttributes) SetPodCount(v int64) {
	o.PodCount = &v
}

// GetPodNames returns the PodNames field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetPodNames() []string {
	if o == nil || o.PodNames == nil {
		var ret []string
		return ret
	}
	return o.PodNames
}

// GetPodNamesOk returns a tuple with the PodNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetPodNamesOk() (*[]string, bool) {
	if o == nil || o.PodNames == nil {
		return nil, false
	}
	return &o.PodNames, true
}

// HasPodNames returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasPodNames() bool {
	return o != nil && o.PodNames != nil
}

// SetPodNames gets a reference to the given []string and assigns it to the PodNames field.
func (o *FleetInstrumentedPodGroupAttributes) SetPodNames(v []string) {
	o.PodNames = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FleetInstrumentedPodGroupAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetInstrumentedPodGroupAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FleetInstrumentedPodGroupAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *FleetInstrumentedPodGroupAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetInstrumentedPodGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppliedTarget != nil {
		toSerialize["applied_target"] = o.AppliedTarget
	}
	if o.AppliedTargetName != nil {
		toSerialize["applied_target_name"] = o.AppliedTargetName
	}
	if o.InjectedTags != nil {
		toSerialize["injected_tags"] = o.InjectedTags
	}
	if o.KubeOwnerrefKind != nil {
		toSerialize["kube_ownerref_kind"] = o.KubeOwnerrefKind
	}
	if o.KubeOwnerrefName != nil {
		toSerialize["kube_ownerref_name"] = o.KubeOwnerrefName
	}
	if o.LibInjectionAnnotations != nil {
		toSerialize["lib_injection_annotations"] = o.LibInjectionAnnotations
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.PodCount != nil {
		toSerialize["pod_count"] = o.PodCount
	}
	if o.PodNames != nil {
		toSerialize["pod_names"] = o.PodNames
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
func (o *FleetInstrumentedPodGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppliedTarget           map[string]interface{} `json:"applied_target,omitempty"`
		AppliedTargetName       *string                `json:"applied_target_name,omitempty"`
		InjectedTags            []string               `json:"injected_tags,omitempty"`
		KubeOwnerrefKind        *string                `json:"kube_ownerref_kind,omitempty"`
		KubeOwnerrefName        *string                `json:"kube_ownerref_name,omitempty"`
		LibInjectionAnnotations []string               `json:"lib_injection_annotations,omitempty"`
		Namespace               *string                `json:"namespace,omitempty"`
		PodCount                *int64                 `json:"pod_count,omitempty"`
		PodNames                []string               `json:"pod_names,omitempty"`
		Tags                    map[string]string      `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"applied_target", "applied_target_name", "injected_tags", "kube_ownerref_kind", "kube_ownerref_name", "lib_injection_annotations", "namespace", "pod_count", "pod_names", "tags"})
	} else {
		return err
	}
	o.AppliedTarget = all.AppliedTarget
	o.AppliedTargetName = all.AppliedTargetName
	o.InjectedTags = all.InjectedTags
	o.KubeOwnerrefKind = all.KubeOwnerrefKind
	o.KubeOwnerrefName = all.KubeOwnerrefName
	o.LibInjectionAnnotations = all.LibInjectionAnnotations
	o.Namespace = all.Namespace
	o.PodCount = all.PodCount
	o.PodNames = all.PodNames
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
