// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataDeletionResponseItemAttributes Deletion attribute for data deletion response.
type DataDeletionResponseItemAttributes struct {
	// Creation time of the deletion request.
	CreatedAt string `json:"created_at"`
	// User who created the deletion request.
	CreatedBy string `json:"created_by"`
	// Start of requested time window, milliseconds since Unix epoch.
	FromTime int64 `json:"from_time"`
	// List of indexes for the search. If not provided, the search is performed in all indexes.
	Indexes []string `json:"indexes,omitempty"`
	// Whether the deletion request is fully created or not. It can take several minutes to fully create a deletion request depending on the target query and timeframe.
	IsCreated bool `json:"is_created"`
	// Organization ID.
	OrgId int64 `json:"org_id"`
	// Product name.
	Product string `json:"product"`
	// Query for creating a data deletion request.
	Query string `json:"query"`
	// Starting time of the process to delete the requested data.
	StartingAt string `json:"starting_at"`
	// Status of the deletion request.
	Status string `json:"status"`
	// End of requested time window, milliseconds since Unix epoch.
	ToTime int64 `json:"to_time"`
	// Total number of elements to be deleted. Only the data accessible to the current user that matches the query and timeframe provided will be deleted.
	TotalUnrestricted int64 `json:"total_unrestricted"`
	// Update time of the deletion request.
	UpdatedAt string `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataDeletionResponseItemAttributes instantiates a new DataDeletionResponseItemAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataDeletionResponseItemAttributes(createdAt string, createdBy string, fromTime int64, isCreated bool, orgId int64, product string, query string, startingAt string, status string, toTime int64, totalUnrestricted int64, updatedAt string) *DataDeletionResponseItemAttributes {
	this := DataDeletionResponseItemAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.FromTime = fromTime
	this.IsCreated = isCreated
	this.OrgId = orgId
	this.Product = product
	this.Query = query
	this.StartingAt = startingAt
	this.Status = status
	this.ToTime = toTime
	this.TotalUnrestricted = totalUnrestricted
	this.UpdatedAt = updatedAt
	return &this
}

// NewDataDeletionResponseItemAttributesWithDefaults instantiates a new DataDeletionResponseItemAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataDeletionResponseItemAttributesWithDefaults() *DataDeletionResponseItemAttributes {
	this := DataDeletionResponseItemAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DataDeletionResponseItemAttributes) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DataDeletionResponseItemAttributes) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *DataDeletionResponseItemAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *DataDeletionResponseItemAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetFromTime returns the FromTime field value.
func (o *DataDeletionResponseItemAttributes) GetFromTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromTime
}

// GetFromTimeOk returns a tuple with the FromTime field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetFromTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTime, true
}

// SetFromTime sets field value.
func (o *DataDeletionResponseItemAttributes) SetFromTime(v int64) {
	o.FromTime = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *DataDeletionResponseItemAttributes) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *DataDeletionResponseItemAttributes) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *DataDeletionResponseItemAttributes) SetIndexes(v []string) {
	o.Indexes = v
}

// GetIsCreated returns the IsCreated field value.
func (o *DataDeletionResponseItemAttributes) GetIsCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsCreated
}

// GetIsCreatedOk returns a tuple with the IsCreated field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetIsCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCreated, true
}

// SetIsCreated sets field value.
func (o *DataDeletionResponseItemAttributes) SetIsCreated(v bool) {
	o.IsCreated = v
}

// GetOrgId returns the OrgId field value.
func (o *DataDeletionResponseItemAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *DataDeletionResponseItemAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetProduct returns the Product field value.
func (o *DataDeletionResponseItemAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *DataDeletionResponseItemAttributes) SetProduct(v string) {
	o.Product = v
}

// GetQuery returns the Query field value.
func (o *DataDeletionResponseItemAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DataDeletionResponseItemAttributes) SetQuery(v string) {
	o.Query = v
}

// GetStartingAt returns the StartingAt field value.
func (o *DataDeletionResponseItemAttributes) GetStartingAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StartingAt
}

// GetStartingAtOk returns a tuple with the StartingAt field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetStartingAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartingAt, true
}

// SetStartingAt sets field value.
func (o *DataDeletionResponseItemAttributes) SetStartingAt(v string) {
	o.StartingAt = v
}

// GetStatus returns the Status field value.
func (o *DataDeletionResponseItemAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DataDeletionResponseItemAttributes) SetStatus(v string) {
	o.Status = v
}

// GetToTime returns the ToTime field value.
func (o *DataDeletionResponseItemAttributes) GetToTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToTime
}

// GetToTimeOk returns a tuple with the ToTime field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetToTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTime, true
}

// SetToTime sets field value.
func (o *DataDeletionResponseItemAttributes) SetToTime(v int64) {
	o.ToTime = v
}

// GetTotalUnrestricted returns the TotalUnrestricted field value.
func (o *DataDeletionResponseItemAttributes) GetTotalUnrestricted() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalUnrestricted
}

// GetTotalUnrestrictedOk returns a tuple with the TotalUnrestricted field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetTotalUnrestrictedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalUnrestricted, true
}

// SetTotalUnrestricted sets field value.
func (o *DataDeletionResponseItemAttributes) SetTotalUnrestricted(v int64) {
	o.TotalUnrestricted = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *DataDeletionResponseItemAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DataDeletionResponseItemAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *DataDeletionResponseItemAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataDeletionResponseItemAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["from_time"] = o.FromTime
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["is_created"] = o.IsCreated
	toSerialize["org_id"] = o.OrgId
	toSerialize["product"] = o.Product
	toSerialize["query"] = o.Query
	toSerialize["starting_at"] = o.StartingAt
	toSerialize["status"] = o.Status
	toSerialize["to_time"] = o.ToTime
	toSerialize["total_unrestricted"] = o.TotalUnrestricted
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataDeletionResponseItemAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *string  `json:"created_at"`
		CreatedBy         *string  `json:"created_by"`
		FromTime          *int64   `json:"from_time"`
		Indexes           []string `json:"indexes,omitempty"`
		IsCreated         *bool    `json:"is_created"`
		OrgId             *int64   `json:"org_id"`
		Product           *string  `json:"product"`
		Query             *string  `json:"query"`
		StartingAt        *string  `json:"starting_at"`
		Status            *string  `json:"status"`
		ToTime            *int64   `json:"to_time"`
		TotalUnrestricted *int64   `json:"total_unrestricted"`
		UpdatedAt         *string  `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.FromTime == nil {
		return fmt.Errorf("required field from_time missing")
	}
	if all.IsCreated == nil {
		return fmt.Errorf("required field is_created missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.StartingAt == nil {
		return fmt.Errorf("required field starting_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.ToTime == nil {
		return fmt.Errorf("required field to_time missing")
	}
	if all.TotalUnrestricted == nil {
		return fmt.Errorf("required field total_unrestricted missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "from_time", "indexes", "is_created", "org_id", "product", "query", "starting_at", "status", "to_time", "total_unrestricted", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.FromTime = *all.FromTime
	o.Indexes = all.Indexes
	o.IsCreated = *all.IsCreated
	o.OrgId = *all.OrgId
	o.Product = *all.Product
	o.Query = *all.Query
	o.StartingAt = *all.StartingAt
	o.Status = *all.Status
	o.ToTime = *all.ToTime
	o.TotalUnrestricted = *all.TotalUnrestricted
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
