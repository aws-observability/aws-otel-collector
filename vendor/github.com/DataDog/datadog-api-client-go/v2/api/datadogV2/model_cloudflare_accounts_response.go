// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareAccountsResponse The expected response schema when getting Cloudflare accounts.
type CloudflareAccountsResponse struct {
	// The JSON:API data schema.
	Data []CloudflareAccountResponseData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudflareAccountsResponse instantiates a new CloudflareAccountsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudflareAccountsResponse() *CloudflareAccountsResponse {
	this := CloudflareAccountsResponse{}
	return &this
}

// NewCloudflareAccountsResponseWithDefaults instantiates a new CloudflareAccountsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudflareAccountsResponseWithDefaults() *CloudflareAccountsResponse {
	this := CloudflareAccountsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudflareAccountsResponse) GetData() []CloudflareAccountResponseData {
	if o == nil || o.Data == nil {
		var ret []CloudflareAccountResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountsResponse) GetDataOk() (*[]CloudflareAccountResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudflareAccountsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []CloudflareAccountResponseData and assigns it to the Data field.
func (o *CloudflareAccountsResponse) SetData(v []CloudflareAccountResponseData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudflareAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudflareAccountsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []CloudflareAccountResponseData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
