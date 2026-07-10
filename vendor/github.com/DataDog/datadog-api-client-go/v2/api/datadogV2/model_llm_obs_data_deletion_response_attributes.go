// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDataDeletionResponseAttributes Attributes of a submitted LLM Observability data deletion request.
type LLMObsDataDeletionResponseAttributes struct {
	// Timestamp when the deletion request was created.
	CreatedAt time.Time `json:"created_at"`
	// UUID of the user who created the deletion request.
	CreatedBy string `json:"created_by"`
	// Start of the deletion time range in milliseconds since Unix epoch.
	FromTime int64 `json:"from_time"`
	// ID of the organization that submitted the deletion request.
	OrgId int64 `json:"org_id"`
	// Product name for the deletion request.
	Product string `json:"product"`
	// The query string used to select data for deletion.
	Query string `json:"query"`
	// End of the deletion time range in milliseconds since Unix epoch.
	ToTime int64 `json:"to_time"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDataDeletionResponseAttributes instantiates a new LLMObsDataDeletionResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDataDeletionResponseAttributes(createdAt time.Time, createdBy string, fromTime int64, orgId int64, product string, query string, toTime int64) *LLMObsDataDeletionResponseAttributes {
	this := LLMObsDataDeletionResponseAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.FromTime = fromTime
	this.OrgId = orgId
	this.Product = product
	this.Query = query
	this.ToTime = toTime
	return &this
}

// NewLLMObsDataDeletionResponseAttributesWithDefaults instantiates a new LLMObsDataDeletionResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDataDeletionResponseAttributesWithDefaults() *LLMObsDataDeletionResponseAttributes {
	this := LLMObsDataDeletionResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsDataDeletionResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *LLMObsDataDeletionResponseAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetFromTime returns the FromTime field value.
func (o *LLMObsDataDeletionResponseAttributes) GetFromTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FromTime
}

// GetFromTimeOk returns a tuple with the FromTime field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetFromTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTime, true
}

// SetFromTime sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetFromTime(v int64) {
	o.FromTime = v
}

// GetOrgId returns the OrgId field value.
func (o *LLMObsDataDeletionResponseAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetProduct returns the Product field value.
func (o *LLMObsDataDeletionResponseAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetProduct(v string) {
	o.Product = v
}

// GetQuery returns the Query field value.
func (o *LLMObsDataDeletionResponseAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetQuery(v string) {
	o.Query = v
}

// GetToTime returns the ToTime field value.
func (o *LLMObsDataDeletionResponseAttributes) GetToTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ToTime
}

// GetToTimeOk returns a tuple with the ToTime field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionResponseAttributes) GetToTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTime, true
}

// SetToTime sets field value.
func (o *LLMObsDataDeletionResponseAttributes) SetToTime(v int64) {
	o.ToTime = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDataDeletionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["from_time"] = o.FromTime
	toSerialize["org_id"] = o.OrgId
	toSerialize["product"] = o.Product
	toSerialize["query"] = o.Query
	toSerialize["to_time"] = o.ToTime

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDataDeletionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time `json:"created_at"`
		CreatedBy *string    `json:"created_by"`
		FromTime  *int64     `json:"from_time"`
		OrgId     *int64     `json:"org_id"`
		Product   *string    `json:"product"`
		Query     *string    `json:"query"`
		ToTime    *int64     `json:"to_time"`
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
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.ToTime == nil {
		return fmt.Errorf("required field to_time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "from_time", "org_id", "product", "query", "to_time"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.FromTime = *all.FromTime
	o.OrgId = *all.OrgId
	o.Product = *all.Product
	o.Query = *all.Query
	o.ToTime = *all.ToTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
