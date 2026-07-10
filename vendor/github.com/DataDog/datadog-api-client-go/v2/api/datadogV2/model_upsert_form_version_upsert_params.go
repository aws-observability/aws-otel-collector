// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertFormVersionUpsertParams Concurrency control parameters for the form version upsert operation.
type UpsertFormVersionUpsertParams struct {
	// The ETag of the latest version. Required when `match_policy` is `if_etag_match`.
	Etag datadog.NullableString `json:"etag,omitempty"`
	// If true, only a new version may be inserted; updating the current draft is not allowed.
	InsertOnly *bool `json:"insert_only,omitempty"`
	// The policy for matching the latest form version during an upsert operation.
	MatchPolicy LatestVersionMatchPolicy `json:"match_policy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertFormVersionUpsertParams instantiates a new UpsertFormVersionUpsertParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertFormVersionUpsertParams(matchPolicy LatestVersionMatchPolicy) *UpsertFormVersionUpsertParams {
	this := UpsertFormVersionUpsertParams{}
	this.MatchPolicy = matchPolicy
	return &this
}

// NewUpsertFormVersionUpsertParamsWithDefaults instantiates a new UpsertFormVersionUpsertParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertFormVersionUpsertParamsWithDefaults() *UpsertFormVersionUpsertParams {
	this := UpsertFormVersionUpsertParams{}
	return &this
}

// GetEtag returns the Etag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertFormVersionUpsertParams) GetEtag() string {
	if o == nil || o.Etag.Get() == nil {
		var ret string
		return ret
	}
	return *o.Etag.Get()
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpsertFormVersionUpsertParams) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Etag.Get(), o.Etag.IsSet()
}

// HasEtag returns a boolean if a field has been set.
func (o *UpsertFormVersionUpsertParams) HasEtag() bool {
	return o != nil && o.Etag.IsSet()
}

// SetEtag gets a reference to the given datadog.NullableString and assigns it to the Etag field.
func (o *UpsertFormVersionUpsertParams) SetEtag(v string) {
	o.Etag.Set(&v)
}

// SetEtagNil sets the value for Etag to be an explicit nil.
func (o *UpsertFormVersionUpsertParams) SetEtagNil() {
	o.Etag.Set(nil)
}

// UnsetEtag ensures that no value is present for Etag, not even an explicit nil.
func (o *UpsertFormVersionUpsertParams) UnsetEtag() {
	o.Etag.Unset()
}

// GetInsertOnly returns the InsertOnly field value if set, zero value otherwise.
func (o *UpsertFormVersionUpsertParams) GetInsertOnly() bool {
	if o == nil || o.InsertOnly == nil {
		var ret bool
		return ret
	}
	return *o.InsertOnly
}

// GetInsertOnlyOk returns a tuple with the InsertOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionUpsertParams) GetInsertOnlyOk() (*bool, bool) {
	if o == nil || o.InsertOnly == nil {
		return nil, false
	}
	return o.InsertOnly, true
}

// HasInsertOnly returns a boolean if a field has been set.
func (o *UpsertFormVersionUpsertParams) HasInsertOnly() bool {
	return o != nil && o.InsertOnly != nil
}

// SetInsertOnly gets a reference to the given bool and assigns it to the InsertOnly field.
func (o *UpsertFormVersionUpsertParams) SetInsertOnly(v bool) {
	o.InsertOnly = &v
}

// GetMatchPolicy returns the MatchPolicy field value.
func (o *UpsertFormVersionUpsertParams) GetMatchPolicy() LatestVersionMatchPolicy {
	if o == nil {
		var ret LatestVersionMatchPolicy
		return ret
	}
	return o.MatchPolicy
}

// GetMatchPolicyOk returns a tuple with the MatchPolicy field value
// and a boolean to check if the value has been set.
func (o *UpsertFormVersionUpsertParams) GetMatchPolicyOk() (*LatestVersionMatchPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchPolicy, true
}

// SetMatchPolicy sets field value.
func (o *UpsertFormVersionUpsertParams) SetMatchPolicy(v LatestVersionMatchPolicy) {
	o.MatchPolicy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertFormVersionUpsertParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Etag.IsSet() {
		toSerialize["etag"] = o.Etag.Get()
	}
	if o.InsertOnly != nil {
		toSerialize["insert_only"] = o.InsertOnly
	}
	toSerialize["match_policy"] = o.MatchPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertFormVersionUpsertParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Etag        datadog.NullableString    `json:"etag,omitempty"`
		InsertOnly  *bool                     `json:"insert_only,omitempty"`
		MatchPolicy *LatestVersionMatchPolicy `json:"match_policy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MatchPolicy == nil {
		return fmt.Errorf("required field match_policy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"etag", "insert_only", "match_policy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Etag = all.Etag
	o.InsertOnly = all.InsertOnly
	if !all.MatchPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.MatchPolicy = *all.MatchPolicy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
