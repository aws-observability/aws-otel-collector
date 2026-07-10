// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsOnDemandHotspotsScalarMeta Metadata for the on-demand hot-spots scalar response.
type CommitmentsOnDemandHotspotsScalarMeta struct {
	// Active on-demand filters applied to the response.
	OnDemandFilters string `json:"on_demand_filters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsOnDemandHotspotsScalarMeta instantiates a new CommitmentsOnDemandHotspotsScalarMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsOnDemandHotspotsScalarMeta(onDemandFilters string) *CommitmentsOnDemandHotspotsScalarMeta {
	this := CommitmentsOnDemandHotspotsScalarMeta{}
	this.OnDemandFilters = onDemandFilters
	return &this
}

// NewCommitmentsOnDemandHotspotsScalarMetaWithDefaults instantiates a new CommitmentsOnDemandHotspotsScalarMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsOnDemandHotspotsScalarMetaWithDefaults() *CommitmentsOnDemandHotspotsScalarMeta {
	this := CommitmentsOnDemandHotspotsScalarMeta{}
	return &this
}

// GetOnDemandFilters returns the OnDemandFilters field value.
func (o *CommitmentsOnDemandHotspotsScalarMeta) GetOnDemandFilters() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OnDemandFilters
}

// GetOnDemandFiltersOk returns a tuple with the OnDemandFilters field value
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarMeta) GetOnDemandFiltersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnDemandFilters, true
}

// SetOnDemandFilters sets field value.
func (o *CommitmentsOnDemandHotspotsScalarMeta) SetOnDemandFilters(v string) {
	o.OnDemandFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsOnDemandHotspotsScalarMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["on_demand_filters"] = o.OnDemandFilters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsOnDemandHotspotsScalarMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OnDemandFilters *string `json:"on_demand_filters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OnDemandFilters == nil {
		return fmt.Errorf("required field on_demand_filters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"on_demand_filters"})
	} else {
		return err
	}
	o.OnDemandFilters = *all.OnDemandFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
