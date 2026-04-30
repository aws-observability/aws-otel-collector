// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyOverrideListResponse Response containing a list of org group policy overrides.
type OrgGroupPolicyOverrideListResponse struct {
	// An array of org group policy overrides.
	Data []OrgGroupPolicyOverrideData `json:"data"`
	// Pagination metadata.
	Meta *OrgGroupPaginationMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyOverrideListResponse instantiates a new OrgGroupPolicyOverrideListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyOverrideListResponse(data []OrgGroupPolicyOverrideData) *OrgGroupPolicyOverrideListResponse {
	this := OrgGroupPolicyOverrideListResponse{}
	this.Data = data
	return &this
}

// NewOrgGroupPolicyOverrideListResponseWithDefaults instantiates a new OrgGroupPolicyOverrideListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyOverrideListResponseWithDefaults() *OrgGroupPolicyOverrideListResponse {
	this := OrgGroupPolicyOverrideListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *OrgGroupPolicyOverrideListResponse) GetData() []OrgGroupPolicyOverrideData {
	if o == nil {
		var ret []OrgGroupPolicyOverrideData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideListResponse) GetDataOk() (*[]OrgGroupPolicyOverrideData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *OrgGroupPolicyOverrideListResponse) SetData(v []OrgGroupPolicyOverrideData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OrgGroupPolicyOverrideListResponse) GetMeta() OrgGroupPaginationMeta {
	if o == nil || o.Meta == nil {
		var ret OrgGroupPaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyOverrideListResponse) GetMetaOk() (*OrgGroupPaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OrgGroupPolicyOverrideListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given OrgGroupPaginationMeta and assigns it to the Meta field.
func (o *OrgGroupPolicyOverrideListResponse) SetMeta(v OrgGroupPaginationMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyOverrideListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyOverrideListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]OrgGroupPolicyOverrideData `json:"data"`
		Meta *OrgGroupPaginationMeta       `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
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
