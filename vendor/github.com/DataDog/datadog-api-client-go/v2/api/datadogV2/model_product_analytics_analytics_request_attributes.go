// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAnalyticsRequestAttributes Attributes for an analytics request.
type ProductAnalyticsAnalyticsRequestAttributes struct {
	// Override the query execution strategy.
	EnforcedExecutionType *ProductAnalyticsExecutionType `json:"enforced_execution_type,omitempty"`
	// Start time in epoch milliseconds. Must be less than `to`.
	From int64 `json:"from"`
	// The analytics query definition containing a base query, compute rule, and optional grouping.
	Query ProductAnalyticsAnalyticsQuery `json:"query"`
	// Optional request ID for multi-step query continuation.
	RequestId *string `json:"request_id,omitempty"`
	// End time in epoch milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAnalyticsRequestAttributes instantiates a new ProductAnalyticsAnalyticsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAnalyticsRequestAttributes(from int64, query ProductAnalyticsAnalyticsQuery, to int64) *ProductAnalyticsAnalyticsRequestAttributes {
	this := ProductAnalyticsAnalyticsRequestAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewProductAnalyticsAnalyticsRequestAttributesWithDefaults instantiates a new ProductAnalyticsAnalyticsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAnalyticsRequestAttributesWithDefaults() *ProductAnalyticsAnalyticsRequestAttributes {
	this := ProductAnalyticsAnalyticsRequestAttributes{}
	return &this
}

// GetEnforcedExecutionType returns the EnforcedExecutionType field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetEnforcedExecutionType() ProductAnalyticsExecutionType {
	if o == nil || o.EnforcedExecutionType == nil {
		var ret ProductAnalyticsExecutionType
		return ret
	}
	return *o.EnforcedExecutionType
}

// GetEnforcedExecutionTypeOk returns a tuple with the EnforcedExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetEnforcedExecutionTypeOk() (*ProductAnalyticsExecutionType, bool) {
	if o == nil || o.EnforcedExecutionType == nil {
		return nil, false
	}
	return o.EnforcedExecutionType, true
}

// HasEnforcedExecutionType returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) HasEnforcedExecutionType() bool {
	return o != nil && o.EnforcedExecutionType != nil
}

// SetEnforcedExecutionType gets a reference to the given ProductAnalyticsExecutionType and assigns it to the EnforcedExecutionType field.
func (o *ProductAnalyticsAnalyticsRequestAttributes) SetEnforcedExecutionType(v ProductAnalyticsExecutionType) {
	o.EnforcedExecutionType = &v
}

// GetFrom returns the From field value.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ProductAnalyticsAnalyticsRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetQuery() ProductAnalyticsAnalyticsQuery {
	if o == nil {
		var ret ProductAnalyticsAnalyticsQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetQueryOk() (*ProductAnalyticsAnalyticsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsAnalyticsRequestAttributes) SetQuery(v ProductAnalyticsAnalyticsQuery) {
	o.Query = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ProductAnalyticsAnalyticsRequestAttributes) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTo returns the To field value.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ProductAnalyticsAnalyticsRequestAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAnalyticsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnforcedExecutionType != nil {
		toSerialize["enforced_execution_type"] = o.EnforcedExecutionType
	}
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAnalyticsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnforcedExecutionType *ProductAnalyticsExecutionType  `json:"enforced_execution_type,omitempty"`
		From                  *int64                          `json:"from"`
		Query                 *ProductAnalyticsAnalyticsQuery `json:"query"`
		RequestId             *string                         `json:"request_id,omitempty"`
		To                    *int64                          `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforced_execution_type", "from", "query", "request_id", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EnforcedExecutionType != nil && !all.EnforcedExecutionType.IsValid() {
		hasInvalidField = true
	} else {
		o.EnforcedExecutionType = all.EnforcedExecutionType
	}
	o.From = *all.From
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	o.RequestId = all.RequestId
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
