// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCIVisibilityHour CI visibility usage in a given hour.
type UsageCIVisibilityHour struct {
	// The number of spans for pipelines in the queried hour.
	CiPipelineIndexedSpans datadog.NullableInt64 `json:"ci_pipeline_indexed_spans,omitempty"`
	// The number of spans for tests in the queried hour.
	CiTestIndexedSpans datadog.NullableInt64 `json:"ci_test_indexed_spans,omitempty"`
	// Shows the total count of all active Git committers for Intelligent Test Runner in the current month. A committer is active if they commit at least 3 times in a given month.
	CiVisibilityItrCommitters datadog.NullableInt64 `json:"ci_visibility_itr_committers,omitempty"`
	// Shows the total count of all active Git committers for Pipelines in the current month. A committer is active if they commit at least 3 times in a given month.
	CiVisibilityPipelineCommitters datadog.NullableInt64 `json:"ci_visibility_pipeline_committers,omitempty"`
	// The total count of all active Git committers for tests in the current month. A committer is active if they commit at least 3 times in a given month.
	CiVisibilityTestCommitters datadog.NullableInt64 `json:"ci_visibility_test_committers,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageCIVisibilityHour instantiates a new UsageCIVisibilityHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCIVisibilityHour() *UsageCIVisibilityHour {
	this := UsageCIVisibilityHour{}
	return &this
}

// NewUsageCIVisibilityHourWithDefaults instantiates a new UsageCIVisibilityHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCIVisibilityHourWithDefaults() *UsageCIVisibilityHour {
	this := UsageCIVisibilityHour{}
	return &this
}

// GetCiPipelineIndexedSpans returns the CiPipelineIndexedSpans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCIVisibilityHour) GetCiPipelineIndexedSpans() int64 {
	if o == nil || o.CiPipelineIndexedSpans.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiPipelineIndexedSpans.Get()
}

// GetCiPipelineIndexedSpansOk returns a tuple with the CiPipelineIndexedSpans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCIVisibilityHour) GetCiPipelineIndexedSpansOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpans.Get(), o.CiPipelineIndexedSpans.IsSet()
}

// HasCiPipelineIndexedSpans returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasCiPipelineIndexedSpans() bool {
	return o != nil && o.CiPipelineIndexedSpans.IsSet()
}

// SetCiPipelineIndexedSpans gets a reference to the given datadog.NullableInt64 and assigns it to the CiPipelineIndexedSpans field.
func (o *UsageCIVisibilityHour) SetCiPipelineIndexedSpans(v int64) {
	o.CiPipelineIndexedSpans.Set(&v)
}

// SetCiPipelineIndexedSpansNil sets the value for CiPipelineIndexedSpans to be an explicit nil.
func (o *UsageCIVisibilityHour) SetCiPipelineIndexedSpansNil() {
	o.CiPipelineIndexedSpans.Set(nil)
}

// UnsetCiPipelineIndexedSpans ensures that no value is present for CiPipelineIndexedSpans, not even an explicit nil.
func (o *UsageCIVisibilityHour) UnsetCiPipelineIndexedSpans() {
	o.CiPipelineIndexedSpans.Unset()
}

// GetCiTestIndexedSpans returns the CiTestIndexedSpans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCIVisibilityHour) GetCiTestIndexedSpans() int64 {
	if o == nil || o.CiTestIndexedSpans.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiTestIndexedSpans.Get()
}

// GetCiTestIndexedSpansOk returns a tuple with the CiTestIndexedSpans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCIVisibilityHour) GetCiTestIndexedSpansOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiTestIndexedSpans.Get(), o.CiTestIndexedSpans.IsSet()
}

// HasCiTestIndexedSpans returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasCiTestIndexedSpans() bool {
	return o != nil && o.CiTestIndexedSpans.IsSet()
}

// SetCiTestIndexedSpans gets a reference to the given datadog.NullableInt64 and assigns it to the CiTestIndexedSpans field.
func (o *UsageCIVisibilityHour) SetCiTestIndexedSpans(v int64) {
	o.CiTestIndexedSpans.Set(&v)
}

// SetCiTestIndexedSpansNil sets the value for CiTestIndexedSpans to be an explicit nil.
func (o *UsageCIVisibilityHour) SetCiTestIndexedSpansNil() {
	o.CiTestIndexedSpans.Set(nil)
}

// UnsetCiTestIndexedSpans ensures that no value is present for CiTestIndexedSpans, not even an explicit nil.
func (o *UsageCIVisibilityHour) UnsetCiTestIndexedSpans() {
	o.CiTestIndexedSpans.Unset()
}

// GetCiVisibilityItrCommitters returns the CiVisibilityItrCommitters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCIVisibilityHour) GetCiVisibilityItrCommitters() int64 {
	if o == nil || o.CiVisibilityItrCommitters.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityItrCommitters.Get()
}

// GetCiVisibilityItrCommittersOk returns a tuple with the CiVisibilityItrCommitters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCIVisibilityHour) GetCiVisibilityItrCommittersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityItrCommitters.Get(), o.CiVisibilityItrCommitters.IsSet()
}

// HasCiVisibilityItrCommitters returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasCiVisibilityItrCommitters() bool {
	return o != nil && o.CiVisibilityItrCommitters.IsSet()
}

// SetCiVisibilityItrCommitters gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityItrCommitters field.
func (o *UsageCIVisibilityHour) SetCiVisibilityItrCommitters(v int64) {
	o.CiVisibilityItrCommitters.Set(&v)
}

// SetCiVisibilityItrCommittersNil sets the value for CiVisibilityItrCommitters to be an explicit nil.
func (o *UsageCIVisibilityHour) SetCiVisibilityItrCommittersNil() {
	o.CiVisibilityItrCommitters.Set(nil)
}

// UnsetCiVisibilityItrCommitters ensures that no value is present for CiVisibilityItrCommitters, not even an explicit nil.
func (o *UsageCIVisibilityHour) UnsetCiVisibilityItrCommitters() {
	o.CiVisibilityItrCommitters.Unset()
}

// GetCiVisibilityPipelineCommitters returns the CiVisibilityPipelineCommitters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCIVisibilityHour) GetCiVisibilityPipelineCommitters() int64 {
	if o == nil || o.CiVisibilityPipelineCommitters.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityPipelineCommitters.Get()
}

// GetCiVisibilityPipelineCommittersOk returns a tuple with the CiVisibilityPipelineCommitters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCIVisibilityHour) GetCiVisibilityPipelineCommittersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityPipelineCommitters.Get(), o.CiVisibilityPipelineCommitters.IsSet()
}

// HasCiVisibilityPipelineCommitters returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasCiVisibilityPipelineCommitters() bool {
	return o != nil && o.CiVisibilityPipelineCommitters.IsSet()
}

// SetCiVisibilityPipelineCommitters gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityPipelineCommitters field.
func (o *UsageCIVisibilityHour) SetCiVisibilityPipelineCommitters(v int64) {
	o.CiVisibilityPipelineCommitters.Set(&v)
}

// SetCiVisibilityPipelineCommittersNil sets the value for CiVisibilityPipelineCommitters to be an explicit nil.
func (o *UsageCIVisibilityHour) SetCiVisibilityPipelineCommittersNil() {
	o.CiVisibilityPipelineCommitters.Set(nil)
}

// UnsetCiVisibilityPipelineCommitters ensures that no value is present for CiVisibilityPipelineCommitters, not even an explicit nil.
func (o *UsageCIVisibilityHour) UnsetCiVisibilityPipelineCommitters() {
	o.CiVisibilityPipelineCommitters.Unset()
}

// GetCiVisibilityTestCommitters returns the CiVisibilityTestCommitters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCIVisibilityHour) GetCiVisibilityTestCommitters() int64 {
	if o == nil || o.CiVisibilityTestCommitters.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityTestCommitters.Get()
}

// GetCiVisibilityTestCommittersOk returns a tuple with the CiVisibilityTestCommitters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCIVisibilityHour) GetCiVisibilityTestCommittersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CiVisibilityTestCommitters.Get(), o.CiVisibilityTestCommitters.IsSet()
}

// HasCiVisibilityTestCommitters returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasCiVisibilityTestCommitters() bool {
	return o != nil && o.CiVisibilityTestCommitters.IsSet()
}

// SetCiVisibilityTestCommitters gets a reference to the given datadog.NullableInt64 and assigns it to the CiVisibilityTestCommitters field.
func (o *UsageCIVisibilityHour) SetCiVisibilityTestCommitters(v int64) {
	o.CiVisibilityTestCommitters.Set(&v)
}

// SetCiVisibilityTestCommittersNil sets the value for CiVisibilityTestCommitters to be an explicit nil.
func (o *UsageCIVisibilityHour) SetCiVisibilityTestCommittersNil() {
	o.CiVisibilityTestCommitters.Set(nil)
}

// UnsetCiVisibilityTestCommitters ensures that no value is present for CiVisibilityTestCommitters, not even an explicit nil.
func (o *UsageCIVisibilityHour) UnsetCiVisibilityTestCommitters() {
	o.CiVisibilityTestCommitters.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageCIVisibilityHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCIVisibilityHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageCIVisibilityHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageCIVisibilityHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCIVisibilityHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageCIVisibilityHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageCIVisibilityHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCIVisibilityHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CiPipelineIndexedSpans.IsSet() {
		toSerialize["ci_pipeline_indexed_spans"] = o.CiPipelineIndexedSpans.Get()
	}
	if o.CiTestIndexedSpans.IsSet() {
		toSerialize["ci_test_indexed_spans"] = o.CiTestIndexedSpans.Get()
	}
	if o.CiVisibilityItrCommitters.IsSet() {
		toSerialize["ci_visibility_itr_committers"] = o.CiVisibilityItrCommitters.Get()
	}
	if o.CiVisibilityPipelineCommitters.IsSet() {
		toSerialize["ci_visibility_pipeline_committers"] = o.CiVisibilityPipelineCommitters.Get()
	}
	if o.CiVisibilityTestCommitters.IsSet() {
		toSerialize["ci_visibility_test_committers"] = o.CiVisibilityTestCommitters.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageCIVisibilityHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CiPipelineIndexedSpans         datadog.NullableInt64 `json:"ci_pipeline_indexed_spans,omitempty"`
		CiTestIndexedSpans             datadog.NullableInt64 `json:"ci_test_indexed_spans,omitempty"`
		CiVisibilityItrCommitters      datadog.NullableInt64 `json:"ci_visibility_itr_committers,omitempty"`
		CiVisibilityPipelineCommitters datadog.NullableInt64 `json:"ci_visibility_pipeline_committers,omitempty"`
		CiVisibilityTestCommitters     datadog.NullableInt64 `json:"ci_visibility_test_committers,omitempty"`
		OrgName                        *string               `json:"org_name,omitempty"`
		PublicId                       *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ci_pipeline_indexed_spans", "ci_test_indexed_spans", "ci_visibility_itr_committers", "ci_visibility_pipeline_committers", "ci_visibility_test_committers", "org_name", "public_id"})
	} else {
		return err
	}
	o.CiPipelineIndexedSpans = all.CiPipelineIndexedSpans
	o.CiTestIndexedSpans = all.CiTestIndexedSpans
	o.CiVisibilityItrCommitters = all.CiVisibilityItrCommitters
	o.CiVisibilityPipelineCommitters = all.CiVisibilityPipelineCommitters
	o.CiVisibilityTestCommitters = all.CiVisibilityTestCommitters
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
