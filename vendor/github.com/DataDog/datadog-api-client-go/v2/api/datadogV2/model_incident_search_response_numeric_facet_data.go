// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchResponseNumericFacetData Facet data numeric attributes of an incident.
type IncidentSearchResponseNumericFacetData struct {
	// Aggregate information for numeric incident data.
	Aggregates IncidentSearchResponseNumericFacetDataAggregates `json:"aggregates"`
	// Name of the incident property field.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentSearchResponseNumericFacetData instantiates a new IncidentSearchResponseNumericFacetData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSearchResponseNumericFacetData(aggregates IncidentSearchResponseNumericFacetDataAggregates, name string) *IncidentSearchResponseNumericFacetData {
	this := IncidentSearchResponseNumericFacetData{}
	this.Aggregates = aggregates
	this.Name = name
	return &this
}

// NewIncidentSearchResponseNumericFacetDataWithDefaults instantiates a new IncidentSearchResponseNumericFacetData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSearchResponseNumericFacetDataWithDefaults() *IncidentSearchResponseNumericFacetData {
	this := IncidentSearchResponseNumericFacetData{}
	return &this
}

// GetAggregates returns the Aggregates field value.
func (o *IncidentSearchResponseNumericFacetData) GetAggregates() IncidentSearchResponseNumericFacetDataAggregates {
	if o == nil {
		var ret IncidentSearchResponseNumericFacetDataAggregates
		return ret
	}
	return o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseNumericFacetData) GetAggregatesOk() (*IncidentSearchResponseNumericFacetDataAggregates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregates, true
}

// SetAggregates sets field value.
func (o *IncidentSearchResponseNumericFacetData) SetAggregates(v IncidentSearchResponseNumericFacetDataAggregates) {
	o.Aggregates = v
}

// GetName returns the Name field value.
func (o *IncidentSearchResponseNumericFacetData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponseNumericFacetData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentSearchResponseNumericFacetData) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSearchResponseNumericFacetData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregates"] = o.Aggregates
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSearchResponseNumericFacetData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregates *IncidentSearchResponseNumericFacetDataAggregates `json:"aggregates"`
		Name       *string                                           `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregates == nil {
		return fmt.Errorf("required field aggregates missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregates", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Aggregates = *all.Aggregates
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
