// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventParentPipeline If the pipeline is triggered as child of another pipeline, this should contain the details of the parent pipeline.
type CIAppPipelineEventParentPipeline struct {
	// UUID of a pipeline.
	Id string `json:"id"`
	// The URL to look at the pipeline in the CI provider UI.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCIAppPipelineEventParentPipeline instantiates a new CIAppPipelineEventParentPipeline object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppPipelineEventParentPipeline(id string) *CIAppPipelineEventParentPipeline {
	this := CIAppPipelineEventParentPipeline{}
	this.Id = id
	return &this
}

// NewCIAppPipelineEventParentPipelineWithDefaults instantiates a new CIAppPipelineEventParentPipeline object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppPipelineEventParentPipelineWithDefaults() *CIAppPipelineEventParentPipeline {
	this := CIAppPipelineEventParentPipeline{}
	return &this
}

// GetId returns the Id field value.
func (o *CIAppPipelineEventParentPipeline) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventParentPipeline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CIAppPipelineEventParentPipeline) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CIAppPipelineEventParentPipeline) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppPipelineEventParentPipeline) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CIAppPipelineEventParentPipeline) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CIAppPipelineEventParentPipeline) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppPipelineEventParentPipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CIAppPipelineEventParentPipeline) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id  *string `json:"id"`
		Url *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "url"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableCIAppPipelineEventParentPipeline handles when a null is used for CIAppPipelineEventParentPipeline.
type NullableCIAppPipelineEventParentPipeline struct {
	value *CIAppPipelineEventParentPipeline
	isSet bool
}

// Get returns the associated value.
func (v NullableCIAppPipelineEventParentPipeline) Get() *CIAppPipelineEventParentPipeline {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCIAppPipelineEventParentPipeline) Set(val *CIAppPipelineEventParentPipeline) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCIAppPipelineEventParentPipeline) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCIAppPipelineEventParentPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCIAppPipelineEventParentPipeline initializes the struct as if Set has been called.
func NewNullableCIAppPipelineEventParentPipeline(val *CIAppPipelineEventParentPipeline) *NullableCIAppPipelineEventParentPipeline {
	return &NullableCIAppPipelineEventParentPipeline{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCIAppPipelineEventParentPipeline) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCIAppPipelineEventParentPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
