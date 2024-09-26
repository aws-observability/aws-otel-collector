// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelWidgetRequest Updated funnel widget.
type FunnelWidgetRequest struct {
	// Updated funnel widget.
	Query FunnelQuery `json:"query"`
	// Widget request type.
	RequestType FunnelRequestType `json:"request_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelWidgetRequest instantiates a new FunnelWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelWidgetRequest(query FunnelQuery, requestType FunnelRequestType) *FunnelWidgetRequest {
	this := FunnelWidgetRequest{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewFunnelWidgetRequestWithDefaults instantiates a new FunnelWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelWidgetRequestWithDefaults() *FunnelWidgetRequest {
	this := FunnelWidgetRequest{}
	return &this
}

// GetQuery returns the Query field value.
func (o *FunnelWidgetRequest) GetQuery() FunnelQuery {
	if o == nil {
		var ret FunnelQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FunnelWidgetRequest) GetQueryOk() (*FunnelQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FunnelWidgetRequest) SetQuery(v FunnelQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *FunnelWidgetRequest) GetRequestType() FunnelRequestType {
	if o == nil {
		var ret FunnelRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *FunnelWidgetRequest) GetRequestTypeOk() (*FunnelRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *FunnelWidgetRequest) SetRequestType(v FunnelRequestType) {
	o.RequestType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query       *FunnelQuery       `json:"query"`
		RequestType *FunnelRequestType `json:"request_type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "request_type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
