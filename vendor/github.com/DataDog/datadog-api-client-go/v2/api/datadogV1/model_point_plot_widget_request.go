// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotWidgetRequest Request configuration for the point plot widget.
type PointPlotWidgetRequest struct {
	// Maximum number of data points to return.
	Limit *int64 `json:"limit,omitempty"`
	// Projection configuration for the point plot widget.
	Projection PointPlotProjection `json:"projection"`
	// Query configuration for a data projection request.
	Query DataProjectionQuery `json:"query"`
	// Type of a data projection request.
	RequestType DataProjectionRequestType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPointPlotWidgetRequest instantiates a new PointPlotWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotWidgetRequest(projection PointPlotProjection, query DataProjectionQuery, requestType DataProjectionRequestType) *PointPlotWidgetRequest {
	this := PointPlotWidgetRequest{}
	this.Projection = projection
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewPointPlotWidgetRequestWithDefaults instantiates a new PointPlotWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotWidgetRequestWithDefaults() *PointPlotWidgetRequest {
	this := PointPlotWidgetRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PointPlotWidgetRequest) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetRequest) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PointPlotWidgetRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *PointPlotWidgetRequest) SetLimit(v int64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value.
func (o *PointPlotWidgetRequest) GetProjection() PointPlotProjection {
	if o == nil {
		var ret PointPlotProjection
		return ret
	}
	return o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetRequest) GetProjectionOk() (*PointPlotProjection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projection, true
}

// SetProjection sets field value.
func (o *PointPlotWidgetRequest) SetProjection(v PointPlotProjection) {
	o.Projection = v
}

// GetQuery returns the Query field value.
func (o *PointPlotWidgetRequest) GetQuery() DataProjectionQuery {
	if o == nil {
		var ret DataProjectionQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetRequest) GetQueryOk() (*DataProjectionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *PointPlotWidgetRequest) SetQuery(v DataProjectionQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *PointPlotWidgetRequest) GetRequestType() DataProjectionRequestType {
	if o == nil {
		var ret DataProjectionRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetRequest) GetRequestTypeOk() (*DataProjectionRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *PointPlotWidgetRequest) SetRequestType(v DataProjectionRequestType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["projection"] = o.Projection
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PointPlotWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit       *int64                     `json:"limit,omitempty"`
		Projection  *PointPlotProjection       `json:"projection"`
		Query       *DataProjectionQuery       `json:"query"`
		RequestType *DataProjectionRequestType `json:"request_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Projection == nil {
		return fmt.Errorf("required field projection missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit", "projection", "query", "request_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Limit = all.Limit
	if all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = *all.Projection
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
