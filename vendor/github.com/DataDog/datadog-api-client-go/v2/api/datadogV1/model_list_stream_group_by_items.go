// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamGroupByItems List of facets on which to group.
type ListStreamGroupByItems struct {
	// Facet name.
	Facet string `json:"facet"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListStreamGroupByItems instantiates a new ListStreamGroupByItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListStreamGroupByItems(facet string) *ListStreamGroupByItems {
	this := ListStreamGroupByItems{}
	this.Facet = facet
	return &this
}

// NewListStreamGroupByItemsWithDefaults instantiates a new ListStreamGroupByItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListStreamGroupByItemsWithDefaults() *ListStreamGroupByItems {
	this := ListStreamGroupByItems{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *ListStreamGroupByItems) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *ListStreamGroupByItems) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *ListStreamGroupByItems) SetFacet(v string) {
	o.Facet = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListStreamGroupByItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListStreamGroupByItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet *string `json:"facet"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet"})
	} else {
		return err
	}
	o.Facet = *all.Facet

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
