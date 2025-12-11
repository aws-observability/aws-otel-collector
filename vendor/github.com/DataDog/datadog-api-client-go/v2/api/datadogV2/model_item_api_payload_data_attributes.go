// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadDataAttributes Metadata and content of a datastore item.
type ItemApiPayloadDataAttributes struct {
	// Timestamp when the item was first created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp when the item was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The ID of the organization that owns this item.
	OrgId *int64 `json:"org_id,omitempty"`
	// The name of the primary key column for this datastore. Primary column names:
	//   - Must abide by both [PostgreSQL naming conventions](https://www.postgresql.org/docs/7.0/syntax525.htm)
	//   - Cannot exceed 63 characters
	PrimaryColumnName *string `json:"primary_column_name,omitempty"`
	// A unique signature identifying this item version.
	Signature *string `json:"signature,omitempty"`
	// The unique identifier of the datastore containing this item.
	StoreId *string `json:"store_id,omitempty"`
	// The data content (as key-value pairs) of a datastore item.
	Value map[string]interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewItemApiPayloadDataAttributes instantiates a new ItemApiPayloadDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewItemApiPayloadDataAttributes() *ItemApiPayloadDataAttributes {
	this := ItemApiPayloadDataAttributes{}
	return &this
}

// NewItemApiPayloadDataAttributesWithDefaults instantiates a new ItemApiPayloadDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewItemApiPayloadDataAttributesWithDefaults() *ItemApiPayloadDataAttributes {
	this := ItemApiPayloadDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ItemApiPayloadDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ItemApiPayloadDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *ItemApiPayloadDataAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetPrimaryColumnName returns the PrimaryColumnName field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetPrimaryColumnName() string {
	if o == nil || o.PrimaryColumnName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryColumnName
}

// GetPrimaryColumnNameOk returns a tuple with the PrimaryColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetPrimaryColumnNameOk() (*string, bool) {
	if o == nil || o.PrimaryColumnName == nil {
		return nil, false
	}
	return o.PrimaryColumnName, true
}

// HasPrimaryColumnName returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasPrimaryColumnName() bool {
	return o != nil && o.PrimaryColumnName != nil
}

// SetPrimaryColumnName gets a reference to the given string and assigns it to the PrimaryColumnName field.
func (o *ItemApiPayloadDataAttributes) SetPrimaryColumnName(v string) {
	o.PrimaryColumnName = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasSignature() bool {
	return o != nil && o.Signature != nil
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ItemApiPayloadDataAttributes) SetSignature(v string) {
	o.Signature = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetStoreId() string {
	if o == nil || o.StoreId == nil {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetStoreIdOk() (*string, bool) {
	if o == nil || o.StoreId == nil {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasStoreId() bool {
	return o != nil && o.StoreId != nil
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *ItemApiPayloadDataAttributes) SetStoreId(v string) {
	o.StoreId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ItemApiPayloadDataAttributes) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadDataAttributes) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ItemApiPayloadDataAttributes) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *ItemApiPayloadDataAttributes) SetValue(v map[string]interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ItemApiPayloadDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.PrimaryColumnName != nil {
		toSerialize["primary_column_name"] = o.PrimaryColumnName
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.StoreId != nil {
		toSerialize["store_id"] = o.StoreId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ItemApiPayloadDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *time.Time             `json:"created_at,omitempty"`
		ModifiedAt        *time.Time             `json:"modified_at,omitempty"`
		OrgId             *int64                 `json:"org_id,omitempty"`
		PrimaryColumnName *string                `json:"primary_column_name,omitempty"`
		Signature         *string                `json:"signature,omitempty"`
		StoreId           *string                `json:"store_id,omitempty"`
		Value             map[string]interface{} `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "org_id", "primary_column_name", "signature", "store_id", "value"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.OrgId = all.OrgId
	o.PrimaryColumnName = all.PrimaryColumnName
	o.Signature = all.Signature
	o.StoreId = all.StoreId
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
