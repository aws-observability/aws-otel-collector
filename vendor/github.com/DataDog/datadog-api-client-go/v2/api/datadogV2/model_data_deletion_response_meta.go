// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataDeletionResponseMeta The metadata of the data deletion response.
type DataDeletionResponseMeta struct {
	// The total deletion requests created by product.
	CountProduct map[string]int64 `json:"count_product,omitempty"`
	// The total deletion requests created by status.
	CountStatus map[string]int64 `json:"count_status,omitempty"`
	// The next page when searching deletion requests created in the current organization.
	NextPage *string `json:"next_page,omitempty"`
	// The product of the deletion request.
	Product *string `json:"product,omitempty"`
	// The status of the executed request.
	RequestStatus *string `json:"request_status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataDeletionResponseMeta instantiates a new DataDeletionResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataDeletionResponseMeta() *DataDeletionResponseMeta {
	this := DataDeletionResponseMeta{}
	return &this
}

// NewDataDeletionResponseMetaWithDefaults instantiates a new DataDeletionResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataDeletionResponseMetaWithDefaults() *DataDeletionResponseMeta {
	this := DataDeletionResponseMeta{}
	return &this
}

// GetCountProduct returns the CountProduct field value if set, zero value otherwise.
func (o *DataDeletionResponseMeta) GetCountProduct() map[string]int64 {
	if o == nil || o.CountProduct == nil {
		var ret map[string]int64
		return ret
	}
	return o.CountProduct
}

// GetCountProductOk returns a tuple with the CountProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseMeta) GetCountProductOk() (*map[string]int64, bool) {
	if o == nil || o.CountProduct == nil {
		return nil, false
	}
	return &o.CountProduct, true
}

// HasCountProduct returns a boolean if a field has been set.
func (o *DataDeletionResponseMeta) HasCountProduct() bool {
	return o != nil && o.CountProduct != nil
}

// SetCountProduct gets a reference to the given map[string]int64 and assigns it to the CountProduct field.
func (o *DataDeletionResponseMeta) SetCountProduct(v map[string]int64) {
	o.CountProduct = v
}

// GetCountStatus returns the CountStatus field value if set, zero value otherwise.
func (o *DataDeletionResponseMeta) GetCountStatus() map[string]int64 {
	if o == nil || o.CountStatus == nil {
		var ret map[string]int64
		return ret
	}
	return o.CountStatus
}

// GetCountStatusOk returns a tuple with the CountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseMeta) GetCountStatusOk() (*map[string]int64, bool) {
	if o == nil || o.CountStatus == nil {
		return nil, false
	}
	return &o.CountStatus, true
}

// HasCountStatus returns a boolean if a field has been set.
func (o *DataDeletionResponseMeta) HasCountStatus() bool {
	return o != nil && o.CountStatus != nil
}

// SetCountStatus gets a reference to the given map[string]int64 and assigns it to the CountStatus field.
func (o *DataDeletionResponseMeta) SetCountStatus(v map[string]int64) {
	o.CountStatus = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *DataDeletionResponseMeta) GetNextPage() string {
	if o == nil || o.NextPage == nil {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseMeta) GetNextPageOk() (*string, bool) {
	if o == nil || o.NextPage == nil {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *DataDeletionResponseMeta) HasNextPage() bool {
	return o != nil && o.NextPage != nil
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *DataDeletionResponseMeta) SetNextPage(v string) {
	o.NextPage = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *DataDeletionResponseMeta) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseMeta) GetProductOk() (*string, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *DataDeletionResponseMeta) HasProduct() bool {
	return o != nil && o.Product != nil
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *DataDeletionResponseMeta) SetProduct(v string) {
	o.Product = &v
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *DataDeletionResponseMeta) GetRequestStatus() string {
	if o == nil || o.RequestStatus == nil {
		var ret string
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseMeta) GetRequestStatusOk() (*string, bool) {
	if o == nil || o.RequestStatus == nil {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *DataDeletionResponseMeta) HasRequestStatus() bool {
	return o != nil && o.RequestStatus != nil
}

// SetRequestStatus gets a reference to the given string and assigns it to the RequestStatus field.
func (o *DataDeletionResponseMeta) SetRequestStatus(v string) {
	o.RequestStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataDeletionResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CountProduct != nil {
		toSerialize["count_product"] = o.CountProduct
	}
	if o.CountStatus != nil {
		toSerialize["count_status"] = o.CountStatus
	}
	if o.NextPage != nil {
		toSerialize["next_page"] = o.NextPage
	}
	if o.Product != nil {
		toSerialize["product"] = o.Product
	}
	if o.RequestStatus != nil {
		toSerialize["request_status"] = o.RequestStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataDeletionResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CountProduct  map[string]int64 `json:"count_product,omitempty"`
		CountStatus   map[string]int64 `json:"count_status,omitempty"`
		NextPage      *string          `json:"next_page,omitempty"`
		Product       *string          `json:"product,omitempty"`
		RequestStatus *string          `json:"request_status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count_product", "count_status", "next_page", "product", "request_status"})
	} else {
		return err
	}
	o.CountProduct = all.CountProduct
	o.CountStatus = all.CountStatus
	o.NextPage = all.NextPage
	o.Product = all.Product
	o.RequestStatus = all.RequestStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
