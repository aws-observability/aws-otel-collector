// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MuteFindingsRequestDataRelationships Relationships of the mute request.
type MuteFindingsRequestDataRelationships struct {
	// A list of security findings.
	Findings Findings `json:"findings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMuteFindingsRequestDataRelationships instantiates a new MuteFindingsRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMuteFindingsRequestDataRelationships(findings Findings) *MuteFindingsRequestDataRelationships {
	this := MuteFindingsRequestDataRelationships{}
	this.Findings = findings
	return &this
}

// NewMuteFindingsRequestDataRelationshipsWithDefaults instantiates a new MuteFindingsRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMuteFindingsRequestDataRelationshipsWithDefaults() *MuteFindingsRequestDataRelationships {
	this := MuteFindingsRequestDataRelationships{}
	return &this
}

// GetFindings returns the Findings field value.
func (o *MuteFindingsRequestDataRelationships) GetFindings() Findings {
	if o == nil {
		var ret Findings
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *MuteFindingsRequestDataRelationships) GetFindingsOk() (*Findings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *MuteFindingsRequestDataRelationships) SetFindings(v Findings) {
	o.Findings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MuteFindingsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["findings"] = o.Findings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MuteFindingsRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Findings *Findings `json:"findings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Findings == nil {
		return fmt.Errorf("required field findings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"findings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Findings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Findings = *all.Findings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
