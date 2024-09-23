// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitDimension The property by which the graph splits
type SplitDimension struct {
	// The system interprets this attribute differently depending on the data source of the query being split. For metrics, it's a tag. For the events platform, it's an attribute or tag.
	OneGraphPer string `json:"one_graph_per"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSplitDimension instantiates a new SplitDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitDimension(oneGraphPer string) *SplitDimension {
	this := SplitDimension{}
	this.OneGraphPer = oneGraphPer
	return &this
}

// NewSplitDimensionWithDefaults instantiates a new SplitDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitDimensionWithDefaults() *SplitDimension {
	this := SplitDimension{}
	return &this
}

// GetOneGraphPer returns the OneGraphPer field value.
func (o *SplitDimension) GetOneGraphPer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OneGraphPer
}

// GetOneGraphPerOk returns a tuple with the OneGraphPer field value
// and a boolean to check if the value has been set.
func (o *SplitDimension) GetOneGraphPerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OneGraphPer, true
}

// SetOneGraphPer sets field value.
func (o *SplitDimension) SetOneGraphPer(v string) {
	o.OneGraphPer = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["one_graph_per"] = o.OneGraphPer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OneGraphPer *string `json:"one_graph_per"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OneGraphPer == nil {
		return fmt.Errorf("required field one_graph_per missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"one_graph_per"})
	} else {
		return err
	}
	o.OneGraphPer = *all.OneGraphPer

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
