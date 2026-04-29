// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelRequest User journey funnel widget request.
type ProductAnalyticsFunnelRequest struct {
	// Comparison segments.
	ComparisonSegments []string `json:"comparison_segments,omitempty"`
	// Comparison time configuration for funnel widgets.
	ComparisonTime *FunnelComparisonDuration `json:"comparison_time,omitempty"`
	// User journey funnel query definition.
	Query ProductAnalyticsFunnelQuery `json:"query"`
	// Request type for user journey funnel widget.
	RequestType ProductAnalyticsFunnelRequestType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelRequest instantiates a new ProductAnalyticsFunnelRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelRequest(query ProductAnalyticsFunnelQuery, requestType ProductAnalyticsFunnelRequestType) *ProductAnalyticsFunnelRequest {
	this := ProductAnalyticsFunnelRequest{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewProductAnalyticsFunnelRequestWithDefaults instantiates a new ProductAnalyticsFunnelRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelRequestWithDefaults() *ProductAnalyticsFunnelRequest {
	this := ProductAnalyticsFunnelRequest{}
	return &this
}

// GetComparisonSegments returns the ComparisonSegments field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelRequest) GetComparisonSegments() []string {
	if o == nil || o.ComparisonSegments == nil {
		var ret []string
		return ret
	}
	return o.ComparisonSegments
}

// GetComparisonSegmentsOk returns a tuple with the ComparisonSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelRequest) GetComparisonSegmentsOk() (*[]string, bool) {
	if o == nil || o.ComparisonSegments == nil {
		return nil, false
	}
	return &o.ComparisonSegments, true
}

// HasComparisonSegments returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelRequest) HasComparisonSegments() bool {
	return o != nil && o.ComparisonSegments != nil
}

// SetComparisonSegments gets a reference to the given []string and assigns it to the ComparisonSegments field.
func (o *ProductAnalyticsFunnelRequest) SetComparisonSegments(v []string) {
	o.ComparisonSegments = v
}

// GetComparisonTime returns the ComparisonTime field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelRequest) GetComparisonTime() FunnelComparisonDuration {
	if o == nil || o.ComparisonTime == nil {
		var ret FunnelComparisonDuration
		return ret
	}
	return *o.ComparisonTime
}

// GetComparisonTimeOk returns a tuple with the ComparisonTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelRequest) GetComparisonTimeOk() (*FunnelComparisonDuration, bool) {
	if o == nil || o.ComparisonTime == nil {
		return nil, false
	}
	return o.ComparisonTime, true
}

// HasComparisonTime returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelRequest) HasComparisonTime() bool {
	return o != nil && o.ComparisonTime != nil
}

// SetComparisonTime gets a reference to the given FunnelComparisonDuration and assigns it to the ComparisonTime field.
func (o *ProductAnalyticsFunnelRequest) SetComparisonTime(v FunnelComparisonDuration) {
	o.ComparisonTime = &v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsFunnelRequest) GetQuery() ProductAnalyticsFunnelQuery {
	if o == nil {
		var ret ProductAnalyticsFunnelQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelRequest) GetQueryOk() (*ProductAnalyticsFunnelQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsFunnelRequest) SetQuery(v ProductAnalyticsFunnelQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *ProductAnalyticsFunnelRequest) GetRequestType() ProductAnalyticsFunnelRequestType {
	if o == nil {
		var ret ProductAnalyticsFunnelRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelRequest) GetRequestTypeOk() (*ProductAnalyticsFunnelRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *ProductAnalyticsFunnelRequest) SetRequestType(v ProductAnalyticsFunnelRequestType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComparisonSegments != nil {
		toSerialize["comparison_segments"] = o.ComparisonSegments
	}
	if o.ComparisonTime != nil {
		toSerialize["comparison_time"] = o.ComparisonTime
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComparisonSegments []string                           `json:"comparison_segments,omitempty"`
		ComparisonTime     *FunnelComparisonDuration          `json:"comparison_time,omitempty"`
		Query              *ProductAnalyticsFunnelQuery       `json:"query"`
		RequestType        *ProductAnalyticsFunnelRequestType `json:"request_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}

	hasInvalidField := false
	o.ComparisonSegments = all.ComparisonSegments
	if all.ComparisonTime != nil && all.ComparisonTime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ComparisonTime = all.ComparisonTime
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
