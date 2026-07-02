// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertCloudInventorySyncConfigRequestData Storage Management configuration data for the create or update request.
type UpsertCloudInventorySyncConfigRequestData struct {
	// Settings for the cloud provider specified in `data.id`. Include only the matching provider object (`aws`, `gcp`, or `azure`).
	Attributes UpsertCloudInventorySyncConfigRequestAttributes `json:"attributes"`
	// Cloud provider for this sync configuration (`aws`, `gcp`, or `azure`). For requests, must match the provider block supplied under `attributes`.
	Id CloudInventoryCloudProviderId `json:"id"`
	// Always `cloud_provider`.
	Type CloudInventoryCloudProviderRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertCloudInventorySyncConfigRequestData instantiates a new UpsertCloudInventorySyncConfigRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertCloudInventorySyncConfigRequestData(attributes UpsertCloudInventorySyncConfigRequestAttributes, id CloudInventoryCloudProviderId, typeVar CloudInventoryCloudProviderRequestType) *UpsertCloudInventorySyncConfigRequestData {
	this := UpsertCloudInventorySyncConfigRequestData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUpsertCloudInventorySyncConfigRequestDataWithDefaults instantiates a new UpsertCloudInventorySyncConfigRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertCloudInventorySyncConfigRequestDataWithDefaults() *UpsertCloudInventorySyncConfigRequestData {
	this := UpsertCloudInventorySyncConfigRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *UpsertCloudInventorySyncConfigRequestData) GetAttributes() UpsertCloudInventorySyncConfigRequestAttributes {
	if o == nil {
		var ret UpsertCloudInventorySyncConfigRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpsertCloudInventorySyncConfigRequestData) GetAttributesOk() (*UpsertCloudInventorySyncConfigRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *UpsertCloudInventorySyncConfigRequestData) SetAttributes(v UpsertCloudInventorySyncConfigRequestAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *UpsertCloudInventorySyncConfigRequestData) GetId() CloudInventoryCloudProviderId {
	if o == nil {
		var ret CloudInventoryCloudProviderId
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpsertCloudInventorySyncConfigRequestData) GetIdOk() (*CloudInventoryCloudProviderId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UpsertCloudInventorySyncConfigRequestData) SetId(v CloudInventoryCloudProviderId) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *UpsertCloudInventorySyncConfigRequestData) GetType() CloudInventoryCloudProviderRequestType {
	if o == nil {
		var ret CloudInventoryCloudProviderRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpsertCloudInventorySyncConfigRequestData) GetTypeOk() (*CloudInventoryCloudProviderRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpsertCloudInventorySyncConfigRequestData) SetType(v CloudInventoryCloudProviderRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertCloudInventorySyncConfigRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertCloudInventorySyncConfigRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpsertCloudInventorySyncConfigRequestAttributes `json:"attributes"`
		Id         *CloudInventoryCloudProviderId                   `json:"id"`
		Type       *CloudInventoryCloudProviderRequestType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Id.IsValid() {
		hasInvalidField = true
	} else {
		o.Id = *all.Id
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
