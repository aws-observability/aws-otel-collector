// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardWidgetSpecification Vega or Vega-Lite specification for custom visualization rendering. See https://vega.github.io/vega-lite/ for the full grammar reference.
type WildcardWidgetSpecification struct {
	// The Vega or Vega-Lite JSON specification object.
	Contents interface{} `json:"contents"`
	// Type of specification used by the wildcard widget.
	Type WildcardWidgetSpecificationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWildcardWidgetSpecification instantiates a new WildcardWidgetSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWildcardWidgetSpecification(contents interface{}, typeVar WildcardWidgetSpecificationType) *WildcardWidgetSpecification {
	this := WildcardWidgetSpecification{}
	this.Contents = contents
	this.Type = typeVar
	return &this
}

// NewWildcardWidgetSpecificationWithDefaults instantiates a new WildcardWidgetSpecification object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWildcardWidgetSpecificationWithDefaults() *WildcardWidgetSpecification {
	this := WildcardWidgetSpecification{}
	return &this
}

// GetContents returns the Contents field value.
func (o *WildcardWidgetSpecification) GetContents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *WildcardWidgetSpecification) GetContentsOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value.
func (o *WildcardWidgetSpecification) SetContents(v interface{}) {
	o.Contents = v
}

// GetType returns the Type field value.
func (o *WildcardWidgetSpecification) GetType() WildcardWidgetSpecificationType {
	if o == nil {
		var ret WildcardWidgetSpecificationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WildcardWidgetSpecification) GetTypeOk() (*WildcardWidgetSpecificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WildcardWidgetSpecification) SetType(v WildcardWidgetSpecificationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WildcardWidgetSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["contents"] = o.Contents
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WildcardWidgetSpecification) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Contents *interface{}                     `json:"contents"`
		Type     *WildcardWidgetSpecificationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Contents == nil {
		return fmt.Errorf("required field contents missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"contents", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Contents = *all.Contents
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
