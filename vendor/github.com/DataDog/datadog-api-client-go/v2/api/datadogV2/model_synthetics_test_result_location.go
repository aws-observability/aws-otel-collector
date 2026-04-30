// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultLocation Location information for a Synthetic test result.
type SyntheticsTestResultLocation struct {
	// Identifier of the location.
	Id *string `json:"id,omitempty"`
	// Human-readable name of the location.
	Name *string `json:"name,omitempty"`
	// Version of the worker that ran the test.
	Version *string `json:"version,omitempty"`
	// Identifier of the specific worker that ran the test.
	WorkerId *string `json:"worker_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultLocation instantiates a new SyntheticsTestResultLocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultLocation() *SyntheticsTestResultLocation {
	this := SyntheticsTestResultLocation{}
	return &this
}

// NewSyntheticsTestResultLocationWithDefaults instantiates a new SyntheticsTestResultLocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultLocationWithDefaults() *SyntheticsTestResultLocation {
	this := SyntheticsTestResultLocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultLocation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultLocation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultLocation) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultLocation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultLocation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultLocation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultLocation) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultLocation) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SyntheticsTestResultLocation) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultLocation) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SyntheticsTestResultLocation) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SyntheticsTestResultLocation) SetVersion(v string) {
	o.Version = &v
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise.
func (o *SyntheticsTestResultLocation) GetWorkerId() string {
	if o == nil || o.WorkerId == nil {
		var ret string
		return ret
	}
	return *o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultLocation) GetWorkerIdOk() (*string, bool) {
	if o == nil || o.WorkerId == nil {
		return nil, false
	}
	return o.WorkerId, true
}

// HasWorkerId returns a boolean if a field has been set.
func (o *SyntheticsTestResultLocation) HasWorkerId() bool {
	return o != nil && o.WorkerId != nil
}

// SetWorkerId gets a reference to the given string and assigns it to the WorkerId field.
func (o *SyntheticsTestResultLocation) SetWorkerId(v string) {
	o.WorkerId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.WorkerId != nil {
		toSerialize["worker_id"] = o.WorkerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultLocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *string `json:"id,omitempty"`
		Name     *string `json:"name,omitempty"`
		Version  *string `json:"version,omitempty"`
		WorkerId *string `json:"worker_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "version", "worker_id"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.Version = all.Version
	o.WorkerId = all.WorkerId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
