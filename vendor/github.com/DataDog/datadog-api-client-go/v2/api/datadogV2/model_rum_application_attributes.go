// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationAttributes RUM application attributes.
type RUMApplicationAttributes struct {
	// ID of the API key associated with the application.
	ApiKeyId *int32 `json:"api_key_id,omitempty"`
	// ID of the RUM application.
	ApplicationId string `json:"application_id"`
	// Client token of the RUM application.
	ClientToken string `json:"client_token"`
	// Timestamp in ms of the creation date.
	CreatedAt int64 `json:"created_at"`
	// Handle of the creator user.
	CreatedByHandle string `json:"created_by_handle"`
	// Hash of the RUM application. Optional.
	Hash *string `json:"hash,omitempty"`
	// Indicates if the RUM application is active.
	IsActive *bool `json:"is_active,omitempty"`
	// Name of the RUM application.
	Name string `json:"name"`
	// Org ID of the RUM application.
	OrgId int32 `json:"org_id"`
	// Product Scales configuration for the RUM application.
	ProductScales *RUMProductScales `json:"product_scales,omitempty"`
	// Type of the RUM application. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`, `roku`, `electron`, `unity`, `kotlin-multiplatform`.
	Type string `json:"type"`
	// Timestamp in ms of the last update date.
	UpdatedAt int64 `json:"updated_at"`
	// Handle of the updater user.
	UpdatedByHandle string `json:"updated_by_handle"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMApplicationAttributes instantiates a new RUMApplicationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationAttributes(applicationId string, clientToken string, createdAt int64, createdByHandle string, name string, orgId int32, typeVar string, updatedAt int64, updatedByHandle string) *RUMApplicationAttributes {
	this := RUMApplicationAttributes{}
	this.ApplicationId = applicationId
	this.ClientToken = clientToken
	this.CreatedAt = createdAt
	this.CreatedByHandle = createdByHandle
	this.Name = name
	this.OrgId = orgId
	this.Type = typeVar
	this.UpdatedAt = updatedAt
	this.UpdatedByHandle = updatedByHandle
	return &this
}

// NewRUMApplicationAttributesWithDefaults instantiates a new RUMApplicationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationAttributesWithDefaults() *RUMApplicationAttributes {
	this := RUMApplicationAttributes{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *RUMApplicationAttributes) GetApiKeyId() int32 {
	if o == nil || o.ApiKeyId == nil {
		var ret int32
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetApiKeyIdOk() (*int32, bool) {
	if o == nil || o.ApiKeyId == nil {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *RUMApplicationAttributes) HasApiKeyId() bool {
	return o != nil && o.ApiKeyId != nil
}

// SetApiKeyId gets a reference to the given int32 and assigns it to the ApiKeyId field.
func (o *RUMApplicationAttributes) SetApiKeyId(v int32) {
	o.ApiKeyId = &v
}

// GetApplicationId returns the ApplicationId field value.
func (o *RUMApplicationAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *RUMApplicationAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetClientToken returns the ClientToken field value.
func (o *RUMApplicationAttributes) GetClientToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetClientTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientToken, true
}

// SetClientToken sets field value.
func (o *RUMApplicationAttributes) SetClientToken(v string) {
	o.ClientToken = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RUMApplicationAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RUMApplicationAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedByHandle returns the CreatedByHandle field value.
func (o *RUMApplicationAttributes) GetCreatedByHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByHandle, true
}

// SetCreatedByHandle sets field value.
func (o *RUMApplicationAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *RUMApplicationAttributes) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *RUMApplicationAttributes) HasHash() bool {
	return o != nil && o.Hash != nil
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *RUMApplicationAttributes) SetHash(v string) {
	o.Hash = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RUMApplicationAttributes) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RUMApplicationAttributes) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RUMApplicationAttributes) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetName returns the Name field value.
func (o *RUMApplicationAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMApplicationAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *RUMApplicationAttributes) GetOrgId() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetOrgIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *RUMApplicationAttributes) SetOrgId(v int32) {
	o.OrgId = v
}

// GetProductScales returns the ProductScales field value if set, zero value otherwise.
func (o *RUMApplicationAttributes) GetProductScales() RUMProductScales {
	if o == nil || o.ProductScales == nil {
		var ret RUMProductScales
		return ret
	}
	return *o.ProductScales
}

// GetProductScalesOk returns a tuple with the ProductScales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetProductScalesOk() (*RUMProductScales, bool) {
	if o == nil || o.ProductScales == nil {
		return nil, false
	}
	return o.ProductScales, true
}

// HasProductScales returns a boolean if a field has been set.
func (o *RUMApplicationAttributes) HasProductScales() bool {
	return o != nil && o.ProductScales != nil
}

// SetProductScales gets a reference to the given RUMProductScales and assigns it to the ProductScales field.
func (o *RUMApplicationAttributes) SetProductScales(v RUMProductScales) {
	o.ProductScales = &v
}

// GetType returns the Type field value.
func (o *RUMApplicationAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RUMApplicationAttributes) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *RUMApplicationAttributes) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *RUMApplicationAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value.
func (o *RUMApplicationAttributes) GetUpdatedByHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedByHandle
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationAttributes) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByHandle, true
}

// SetUpdatedByHandle sets field value.
func (o *RUMApplicationAttributes) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKeyId != nil {
		toSerialize["api_key_id"] = o.ApiKeyId
	}
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["client_token"] = o.ClientToken
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by_handle"] = o.CreatedByHandle
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	if o.ProductScales != nil {
		toSerialize["product_scales"] = o.ProductScales
	}
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updated_by_handle"] = o.UpdatedByHandle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKeyId        *int32            `json:"api_key_id,omitempty"`
		ApplicationId   *string           `json:"application_id"`
		ClientToken     *string           `json:"client_token"`
		CreatedAt       *int64            `json:"created_at"`
		CreatedByHandle *string           `json:"created_by_handle"`
		Hash            *string           `json:"hash,omitempty"`
		IsActive        *bool             `json:"is_active,omitempty"`
		Name            *string           `json:"name"`
		OrgId           *int32            `json:"org_id"`
		ProductScales   *RUMProductScales `json:"product_scales,omitempty"`
		Type            *string           `json:"type"`
		UpdatedAt       *int64            `json:"updated_at"`
		UpdatedByHandle *string           `json:"updated_by_handle"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.ClientToken == nil {
		return fmt.Errorf("required field client_token missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedByHandle == nil {
		return fmt.Errorf("required field created_by_handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedByHandle == nil {
		return fmt.Errorf("required field updated_by_handle missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key_id", "application_id", "client_token", "created_at", "created_by_handle", "hash", "is_active", "name", "org_id", "product_scales", "type", "updated_at", "updated_by_handle"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKeyId = all.ApiKeyId
	o.ApplicationId = *all.ApplicationId
	o.ClientToken = *all.ClientToken
	o.CreatedAt = *all.CreatedAt
	o.CreatedByHandle = *all.CreatedByHandle
	o.Hash = all.Hash
	o.IsActive = all.IsActive
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	if all.ProductScales != nil && all.ProductScales.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProductScales = all.ProductScales
	o.Type = *all.Type
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedByHandle = *all.UpdatedByHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
