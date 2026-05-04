// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOCountSpec A metric SLI specification.
type SLOCountSpec struct {
	// A count-based (metric) SLI specification, composed of three parts: the good events formula,
	// the bad or total events formula, and the underlying queries.
	// Exactly one of `total_events_formula` or `bad_events_formula` must be provided.
	Count SLOCountDefinition `json:"count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSLOCountSpec instantiates a new SLOCountSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLOCountSpec(count SLOCountDefinition) *SLOCountSpec {
	this := SLOCountSpec{}
	this.Count = count
	return &this
}

// NewSLOCountSpecWithDefaults instantiates a new SLOCountSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLOCountSpecWithDefaults() *SLOCountSpec {
	this := SLOCountSpec{}
	return &this
}

// GetCount returns the Count field value.
func (o *SLOCountSpec) GetCount() SLOCountDefinition {
	if o == nil {
		var ret SLOCountDefinition
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SLOCountSpec) GetCountOk() (*SLOCountDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *SLOCountSpec) SetCount(v SLOCountDefinition) {
	o.Count = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLOCountSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLOCountSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *SLOCountDefinition `json:"count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	o.Count = *all.Count

	return nil
}
