// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRumRequest Sankey widget with RUM data source.
type SankeyRumRequest struct {
	// Sankey widget with RUM data source query.
	Query SankeyRumQuery `json:"query"`
	// Type of the Sankey widget.
	RequestType SankeyWidgetDefinitionType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyRumRequest instantiates a new SankeyRumRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyRumRequest(query SankeyRumQuery, requestType SankeyWidgetDefinitionType) *SankeyRumRequest {
	this := SankeyRumRequest{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewSankeyRumRequestWithDefaults instantiates a new SankeyRumRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyRumRequestWithDefaults() *SankeyRumRequest {
	this := SankeyRumRequest{}
	var requestType SankeyWidgetDefinitionType = SANKEYWIDGETDEFINITIONTYPE_SANKEY
	this.RequestType = requestType
	return &this
}

// GetQuery returns the Query field value.
func (o *SankeyRumRequest) GetQuery() SankeyRumQuery {
	if o == nil {
		var ret SankeyRumQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SankeyRumRequest) GetQueryOk() (*SankeyRumQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SankeyRumRequest) SetQuery(v SankeyRumQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *SankeyRumRequest) GetRequestType() SankeyWidgetDefinitionType {
	if o == nil {
		var ret SankeyWidgetDefinitionType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *SankeyRumRequest) GetRequestTypeOk() (*SankeyWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *SankeyRumRequest) SetRequestType(v SankeyWidgetDefinitionType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyRumRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyRumRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query       *SankeyRumQuery             `json:"query"`
		RequestType *SankeyWidgetDefinitionType `json:"request_type"`
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
