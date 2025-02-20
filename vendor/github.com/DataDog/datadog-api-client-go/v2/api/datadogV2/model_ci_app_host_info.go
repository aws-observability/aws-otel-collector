// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppHostInfo Contains information of the host running the pipeline, stage, job, or step.
type CIAppHostInfo struct {
	// FQDN of the host.
	Hostname *string `json:"hostname,omitempty"`
	// A list of labels used to select or identify the node.
	Labels []string `json:"labels,omitempty"`
	// Name for the host.
	Name *string `json:"name,omitempty"`
	// The path where the code is checked out.
	Workspace *string `json:"workspace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppHostInfo instantiates a new CIAppHostInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppHostInfo() *CIAppHostInfo {
	this := CIAppHostInfo{}
	return &this
}

// NewCIAppHostInfoWithDefaults instantiates a new CIAppHostInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppHostInfoWithDefaults() *CIAppHostInfo {
	this := CIAppHostInfo{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CIAppHostInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppHostInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CIAppHostInfo) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CIAppHostInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CIAppHostInfo) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppHostInfo) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CIAppHostInfo) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *CIAppHostInfo) SetLabels(v []string) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CIAppHostInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppHostInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CIAppHostInfo) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CIAppHostInfo) SetName(v string) {
	o.Name = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *CIAppHostInfo) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppHostInfo) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *CIAppHostInfo) HasWorkspace() bool {
	return o != nil && o.Workspace != nil
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *CIAppHostInfo) SetWorkspace(v string) {
	o.Workspace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppHostInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppHostInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hostname  *string  `json:"hostname,omitempty"`
		Labels    []string `json:"labels,omitempty"`
		Name      *string  `json:"name,omitempty"`
		Workspace *string  `json:"workspace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hostname", "labels", "name", "workspace"})
	} else {
		return err
	}
	o.Hostname = all.Hostname
	o.Labels = all.Labels
	o.Name = all.Name
	o.Workspace = all.Workspace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableCIAppHostInfo handles when a null is used for CIAppHostInfo.
type NullableCIAppHostInfo struct {
	value *CIAppHostInfo
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppHostInfo) Get() *CIAppHostInfo {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppHostInfo) Set(val *CIAppHostInfo) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppHostInfo) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCIAppHostInfo) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppHostInfo initializes the struct as if Set has been called.
func NewNullableCIAppHostInfo(val *CIAppHostInfo) *NullableCIAppHostInfo {
	return &NullableCIAppHostInfo{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppHostInfo) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppHostInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
