// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultCIPipeline Details of the CI pipeline.
type SyntheticsTestResultCIPipeline struct {
	// Pipeline identifier.
	Id *string `json:"id,omitempty"`
	// Pipeline name.
	Name *string `json:"name,omitempty"`
	// Pipeline number.
	Number *int64 `json:"number,omitempty"`
	// Pipeline URL.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultCIPipeline instantiates a new SyntheticsTestResultCIPipeline object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultCIPipeline() *SyntheticsTestResultCIPipeline {
	this := SyntheticsTestResultCIPipeline{}
	return &this
}

// NewSyntheticsTestResultCIPipelineWithDefaults instantiates a new SyntheticsTestResultCIPipeline object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultCIPipelineWithDefaults() *SyntheticsTestResultCIPipeline {
	this := SyntheticsTestResultCIPipeline{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultCIPipeline) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCIPipeline) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultCIPipeline) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultCIPipeline) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultCIPipeline) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCIPipeline) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultCIPipeline) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultCIPipeline) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SyntheticsTestResultCIPipeline) GetNumber() int64 {
	if o == nil || o.Number == nil {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCIPipeline) GetNumberOk() (*int64, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SyntheticsTestResultCIPipeline) HasNumber() bool {
	return o != nil && o.Number != nil
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *SyntheticsTestResultCIPipeline) SetNumber(v int64) {
	o.Number = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultCIPipeline) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCIPipeline) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultCIPipeline) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestResultCIPipeline) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultCIPipeline) MarshalJSON() ([]byte, error) {
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
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultCIPipeline) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string `json:"id,omitempty"`
		Name   *string `json:"name,omitempty"`
		Number *int64  `json:"number,omitempty"`
		Url    *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "number", "url"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.Number = all.Number
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
