// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachCaseRequestDataRelationships Relationships of the case to attach security findings to.
type AttachCaseRequestDataRelationships struct {
	// A list of security findings.
	Findings Findings `json:"findings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachCaseRequestDataRelationships instantiates a new AttachCaseRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachCaseRequestDataRelationships(findings Findings) *AttachCaseRequestDataRelationships {
	this := AttachCaseRequestDataRelationships{}
	this.Findings = findings
	return &this
}

// NewAttachCaseRequestDataRelationshipsWithDefaults instantiates a new AttachCaseRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachCaseRequestDataRelationshipsWithDefaults() *AttachCaseRequestDataRelationships {
	this := AttachCaseRequestDataRelationships{}
	return &this
}

// GetFindings returns the Findings field value.
func (o *AttachCaseRequestDataRelationships) GetFindings() Findings {
	if o == nil {
		var ret Findings
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *AttachCaseRequestDataRelationships) GetFindingsOk() (*Findings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *AttachCaseRequestDataRelationships) SetFindings(v Findings) {
	o.Findings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachCaseRequestDataRelationships) MarshalJSON() ([]byte, error) {
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
func (o *AttachCaseRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
