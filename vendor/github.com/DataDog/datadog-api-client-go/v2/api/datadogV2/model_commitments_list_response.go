// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsListResponse Response containing a list of cloud commitment details.
type CommitmentsListResponse struct {
	// Array of commitment items.
	Commitments []CommitmentsListItem `json:"commitments"`
	// Metadata for a commitments list response.
	Meta *CommitmentsListMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsListResponse instantiates a new CommitmentsListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsListResponse(commitments []CommitmentsListItem) *CommitmentsListResponse {
	this := CommitmentsListResponse{}
	this.Commitments = commitments
	return &this
}

// NewCommitmentsListResponseWithDefaults instantiates a new CommitmentsListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsListResponseWithDefaults() *CommitmentsListResponse {
	this := CommitmentsListResponse{}
	return &this
}

// GetCommitments returns the Commitments field value.
func (o *CommitmentsListResponse) GetCommitments() []CommitmentsListItem {
	if o == nil {
		var ret []CommitmentsListItem
		return ret
	}
	return o.Commitments
}

// GetCommitmentsOk returns a tuple with the Commitments field value
// and a boolean to check if the value has been set.
func (o *CommitmentsListResponse) GetCommitmentsOk() (*[]CommitmentsListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commitments, true
}

// SetCommitments sets field value.
func (o *CommitmentsListResponse) SetCommitments(v []CommitmentsListItem) {
	o.Commitments = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsListResponse) GetMeta() CommitmentsListMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsListMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsListResponse) GetMetaOk() (*CommitmentsListMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsListMeta and assigns it to the Meta field.
func (o *CommitmentsListResponse) SetMeta(v CommitmentsListMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commitments"] = o.Commitments
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Commitments *[]CommitmentsListItem `json:"commitments"`
		Meta        *CommitmentsListMeta   `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Commitments == nil {
		return fmt.Errorf("required field commitments missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commitments", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Commitments = *all.Commitments
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
