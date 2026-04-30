// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedTestsRequestDeleteAttributes Attributes for a bulk delete Synthetic tests request.
type DeletedTestsRequestDeleteAttributes struct {
	// Whether to force deletion of tests that have dependent resources.
	ForceDeleteDependencies *bool `json:"force_delete_dependencies,omitempty"`
	// List of public IDs of the Synthetic tests to delete.
	PublicIds []string `json:"public_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeletedTestsRequestDeleteAttributes instantiates a new DeletedTestsRequestDeleteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeletedTestsRequestDeleteAttributes(publicIds []string) *DeletedTestsRequestDeleteAttributes {
	this := DeletedTestsRequestDeleteAttributes{}
	this.PublicIds = publicIds
	return &this
}

// NewDeletedTestsRequestDeleteAttributesWithDefaults instantiates a new DeletedTestsRequestDeleteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeletedTestsRequestDeleteAttributesWithDefaults() *DeletedTestsRequestDeleteAttributes {
	this := DeletedTestsRequestDeleteAttributes{}
	return &this
}

// GetForceDeleteDependencies returns the ForceDeleteDependencies field value if set, zero value otherwise.
func (o *DeletedTestsRequestDeleteAttributes) GetForceDeleteDependencies() bool {
	if o == nil || o.ForceDeleteDependencies == nil {
		var ret bool
		return ret
	}
	return *o.ForceDeleteDependencies
}

// GetForceDeleteDependenciesOk returns a tuple with the ForceDeleteDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedTestsRequestDeleteAttributes) GetForceDeleteDependenciesOk() (*bool, bool) {
	if o == nil || o.ForceDeleteDependencies == nil {
		return nil, false
	}
	return o.ForceDeleteDependencies, true
}

// HasForceDeleteDependencies returns a boolean if a field has been set.
func (o *DeletedTestsRequestDeleteAttributes) HasForceDeleteDependencies() bool {
	return o != nil && o.ForceDeleteDependencies != nil
}

// SetForceDeleteDependencies gets a reference to the given bool and assigns it to the ForceDeleteDependencies field.
func (o *DeletedTestsRequestDeleteAttributes) SetForceDeleteDependencies(v bool) {
	o.ForceDeleteDependencies = &v
}

// GetPublicIds returns the PublicIds field value.
func (o *DeletedTestsRequestDeleteAttributes) GetPublicIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PublicIds
}

// GetPublicIdsOk returns a tuple with the PublicIds field value
// and a boolean to check if the value has been set.
func (o *DeletedTestsRequestDeleteAttributes) GetPublicIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicIds, true
}

// SetPublicIds sets field value.
func (o *DeletedTestsRequestDeleteAttributes) SetPublicIds(v []string) {
	o.PublicIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeletedTestsRequestDeleteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ForceDeleteDependencies != nil {
		toSerialize["force_delete_dependencies"] = o.ForceDeleteDependencies
	}
	toSerialize["public_ids"] = o.PublicIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeletedTestsRequestDeleteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ForceDeleteDependencies *bool     `json:"force_delete_dependencies,omitempty"`
		PublicIds               *[]string `json:"public_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PublicIds == nil {
		return fmt.Errorf("required field public_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"force_delete_dependencies", "public_ids"})
	} else {
		return err
	}
	o.ForceDeleteDependencies = all.ForceDeleteDependencies
	o.PublicIds = *all.PublicIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
