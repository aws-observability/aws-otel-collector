// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BillingDimensionsMappingBodyItemAttributesEndpointsItems An endpoint's keys mapped to the billing_dimension.
type BillingDimensionsMappingBodyItemAttributesEndpointsItems struct {
	// The URL for the endpoint.
	Id *string `json:"id,omitempty"`
	// The billing dimension.
	Keys []string `json:"keys,omitempty"`
	// Denotes whether mapping keys were available for this endpoint.
	Status *BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBillingDimensionsMappingBodyItemAttributesEndpointsItems instantiates a new BillingDimensionsMappingBodyItemAttributesEndpointsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBillingDimensionsMappingBodyItemAttributesEndpointsItems() *BillingDimensionsMappingBodyItemAttributesEndpointsItems {
	this := BillingDimensionsMappingBodyItemAttributesEndpointsItems{}
	return &this
}

// NewBillingDimensionsMappingBodyItemAttributesEndpointsItemsWithDefaults instantiates a new BillingDimensionsMappingBodyItemAttributesEndpointsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBillingDimensionsMappingBodyItemAttributesEndpointsItemsWithDefaults() *BillingDimensionsMappingBodyItemAttributesEndpointsItems {
	this := BillingDimensionsMappingBodyItemAttributesEndpointsItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) SetId(v string) {
	o.Id = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) GetKeysOk() (*[]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return &o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) HasKeys() bool {
	return o != nil && o.Keys != nil
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) SetKeys(v []string) {
	o.Keys = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) GetStatus() BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus {
	if o == nil || o.Status == nil {
		var ret BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) GetStatusOk() (*BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus and assigns it to the Status field.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) SetStatus(v BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BillingDimensionsMappingBodyItemAttributesEndpointsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BillingDimensionsMappingBodyItemAttributesEndpointsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                                                         `json:"id,omitempty"`
		Keys   []string                                                        `json:"keys,omitempty"`
		Status *BillingDimensionsMappingBodyItemAttributesEndpointsItemsStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "keys", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Keys = all.Keys
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
