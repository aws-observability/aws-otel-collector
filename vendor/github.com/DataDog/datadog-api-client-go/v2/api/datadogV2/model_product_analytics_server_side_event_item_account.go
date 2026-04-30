// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsServerSideEventItemAccount The account linked to your event.
type ProductAnalyticsServerSideEventItemAccount struct {
	// The account ID used in Datadog.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsServerSideEventItemAccount instantiates a new ProductAnalyticsServerSideEventItemAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsServerSideEventItemAccount(id string) *ProductAnalyticsServerSideEventItemAccount {
	this := ProductAnalyticsServerSideEventItemAccount{}
	this.Id = id
	return &this
}

// NewProductAnalyticsServerSideEventItemAccountWithDefaults instantiates a new ProductAnalyticsServerSideEventItemAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsServerSideEventItemAccountWithDefaults() *ProductAnalyticsServerSideEventItemAccount {
	this := ProductAnalyticsServerSideEventItemAccount{}
	return &this
}

// GetId returns the Id field value.
func (o *ProductAnalyticsServerSideEventItemAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItemAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsServerSideEventItemAccount) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsServerSideEventItemAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsServerSideEventItemAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id"})
	} else {
		return err
	}
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
