// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertAndPublishFormVersionUpsertParams Concurrency control parameters for the upsert and publish operation.
type UpsertAndPublishFormVersionUpsertParams struct {
	// The ETag of the latest version used for optimistic concurrency control.
	Etag string `json:"etag"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertAndPublishFormVersionUpsertParams instantiates a new UpsertAndPublishFormVersionUpsertParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertAndPublishFormVersionUpsertParams(etag string) *UpsertAndPublishFormVersionUpsertParams {
	this := UpsertAndPublishFormVersionUpsertParams{}
	this.Etag = etag
	return &this
}

// NewUpsertAndPublishFormVersionUpsertParamsWithDefaults instantiates a new UpsertAndPublishFormVersionUpsertParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertAndPublishFormVersionUpsertParamsWithDefaults() *UpsertAndPublishFormVersionUpsertParams {
	this := UpsertAndPublishFormVersionUpsertParams{}
	return &this
}

// GetEtag returns the Etag field value.
func (o *UpsertAndPublishFormVersionUpsertParams) GetEtag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Etag
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
func (o *UpsertAndPublishFormVersionUpsertParams) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etag, true
}

// SetEtag sets field value.
func (o *UpsertAndPublishFormVersionUpsertParams) SetEtag(v string) {
	o.Etag = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertAndPublishFormVersionUpsertParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["etag"] = o.Etag

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertAndPublishFormVersionUpsertParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Etag *string `json:"etag"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Etag == nil {
		return fmt.Errorf("required field etag missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"etag"})
	} else {
		return err
	}
	o.Etag = *all.Etag

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
