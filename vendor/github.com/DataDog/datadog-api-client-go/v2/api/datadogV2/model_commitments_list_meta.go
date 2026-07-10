// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsListMeta Metadata for a commitments list response.
type CommitmentsListMeta struct {
	// Unit metadata for a numeric metric.
	CommittedSpendUnit *CommitmentsUnit `json:"committed_spend_unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsListMeta instantiates a new CommitmentsListMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsListMeta() *CommitmentsListMeta {
	this := CommitmentsListMeta{}
	return &this
}

// NewCommitmentsListMetaWithDefaults instantiates a new CommitmentsListMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsListMetaWithDefaults() *CommitmentsListMeta {
	this := CommitmentsListMeta{}
	return &this
}

// GetCommittedSpendUnit returns the CommittedSpendUnit field value if set, zero value otherwise.
func (o *CommitmentsListMeta) GetCommittedSpendUnit() CommitmentsUnit {
	if o == nil || o.CommittedSpendUnit == nil {
		var ret CommitmentsUnit
		return ret
	}
	return *o.CommittedSpendUnit
}

// GetCommittedSpendUnitOk returns a tuple with the CommittedSpendUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsListMeta) GetCommittedSpendUnitOk() (*CommitmentsUnit, bool) {
	if o == nil || o.CommittedSpendUnit == nil {
		return nil, false
	}
	return o.CommittedSpendUnit, true
}

// HasCommittedSpendUnit returns a boolean if a field has been set.
func (o *CommitmentsListMeta) HasCommittedSpendUnit() bool {
	return o != nil && o.CommittedSpendUnit != nil
}

// SetCommittedSpendUnit gets a reference to the given CommitmentsUnit and assigns it to the CommittedSpendUnit field.
func (o *CommitmentsListMeta) SetCommittedSpendUnit(v CommitmentsUnit) {
	o.CommittedSpendUnit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CommittedSpendUnit != nil {
		toSerialize["committed_spend_unit"] = o.CommittedSpendUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsListMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommittedSpendUnit *CommitmentsUnit `json:"committed_spend_unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"committed_spend_unit"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CommittedSpendUnit != nil && all.CommittedSpendUnit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CommittedSpendUnit = all.CommittedSpendUnit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
