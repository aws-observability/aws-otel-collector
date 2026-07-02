// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipEvidenceAttributes The attributes of an ownership evidence response.
type OwnershipEvidenceAttributes struct {
	// The list of evidence versions associated with an inference.
	EvidenceVersions datadog.NullableList[map[string]interface{}] `json:"evidence_versions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipEvidenceAttributes instantiates a new OwnershipEvidenceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipEvidenceAttributes(evidenceVersions datadog.NullableList[map[string]interface{}]) *OwnershipEvidenceAttributes {
	this := OwnershipEvidenceAttributes{}
	this.EvidenceVersions = evidenceVersions
	return &this
}

// NewOwnershipEvidenceAttributesWithDefaults instantiates a new OwnershipEvidenceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipEvidenceAttributesWithDefaults() *OwnershipEvidenceAttributes {
	this := OwnershipEvidenceAttributes{}
	return &this
}

// GetEvidenceVersions returns the EvidenceVersions field value.
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned.
func (o *OwnershipEvidenceAttributes) GetEvidenceVersions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.EvidenceVersions.Get()
}

// GetEvidenceVersionsOk returns a tuple with the EvidenceVersions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipEvidenceAttributes) GetEvidenceVersionsOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvidenceVersions.Get(), o.EvidenceVersions.IsSet()
}

// SetEvidenceVersions sets field value.
func (o *OwnershipEvidenceAttributes) SetEvidenceVersions(v []map[string]interface{}) {
	o.EvidenceVersions.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipEvidenceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["evidence_versions"] = o.EvidenceVersions.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipEvidenceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EvidenceVersions datadog.NullableList[map[string]interface{}] `json:"evidence_versions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.EvidenceVersions.IsSet() {
		return fmt.Errorf("required field evidence_versions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"evidence_versions"})
	} else {
		return err
	}
	o.EvidenceVersions = all.EvidenceVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
