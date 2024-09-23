// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BulkMuteFindingsRequestMetaFindings Finding object containing the finding information.
type BulkMuteFindingsRequestMetaFindings struct {
	// The unique ID for this finding.
	FindingId *string `json:"finding_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBulkMuteFindingsRequestMetaFindings instantiates a new BulkMuteFindingsRequestMetaFindings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBulkMuteFindingsRequestMetaFindings() *BulkMuteFindingsRequestMetaFindings {
	this := BulkMuteFindingsRequestMetaFindings{}
	return &this
}

// NewBulkMuteFindingsRequestMetaFindingsWithDefaults instantiates a new BulkMuteFindingsRequestMetaFindings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBulkMuteFindingsRequestMetaFindingsWithDefaults() *BulkMuteFindingsRequestMetaFindings {
	this := BulkMuteFindingsRequestMetaFindings{}
	return &this
}

// GetFindingId returns the FindingId field value if set, zero value otherwise.
func (o *BulkMuteFindingsRequestMetaFindings) GetFindingId() string {
	if o == nil || o.FindingId == nil {
		var ret string
		return ret
	}
	return *o.FindingId
}

// GetFindingIdOk returns a tuple with the FindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMuteFindingsRequestMetaFindings) GetFindingIdOk() (*string, bool) {
	if o == nil || o.FindingId == nil {
		return nil, false
	}
	return o.FindingId, true
}

// HasFindingId returns a boolean if a field has been set.
func (o *BulkMuteFindingsRequestMetaFindings) HasFindingId() bool {
	return o != nil && o.FindingId != nil
}

// SetFindingId gets a reference to the given string and assigns it to the FindingId field.
func (o *BulkMuteFindingsRequestMetaFindings) SetFindingId(v string) {
	o.FindingId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BulkMuteFindingsRequestMetaFindings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FindingId != nil {
		toSerialize["finding_id"] = o.FindingId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BulkMuteFindingsRequestMetaFindings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FindingId *string `json:"finding_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"finding_id"})
	} else {
		return err
	}
	o.FindingId = all.FindingId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
