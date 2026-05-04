// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCurveWidgetRequest Retention curve widget request.
type RetentionCurveWidgetRequest struct {
	// Retention query definition.
	Query RetentionQuery `json:"query"`
	// Request type for retention curve widget.
	RequestType RetentionCurveRequestType `json:"request_type"`
	// Style configuration for retention curve.
	Style *RetentionCurveStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionCurveWidgetRequest instantiates a new RetentionCurveWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionCurveWidgetRequest(query RetentionQuery, requestType RetentionCurveRequestType) *RetentionCurveWidgetRequest {
	this := RetentionCurveWidgetRequest{}
	this.Query = query
	this.RequestType = requestType
	return &this
}

// NewRetentionCurveWidgetRequestWithDefaults instantiates a new RetentionCurveWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionCurveWidgetRequestWithDefaults() *RetentionCurveWidgetRequest {
	this := RetentionCurveWidgetRequest{}
	return &this
}

// GetQuery returns the Query field value.
func (o *RetentionCurveWidgetRequest) GetQuery() RetentionQuery {
	if o == nil {
		var ret RetentionQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RetentionCurveWidgetRequest) GetQueryOk() (*RetentionQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *RetentionCurveWidgetRequest) SetQuery(v RetentionQuery) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *RetentionCurveWidgetRequest) GetRequestType() RetentionCurveRequestType {
	if o == nil {
		var ret RetentionCurveRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *RetentionCurveWidgetRequest) GetRequestTypeOk() (*RetentionCurveRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *RetentionCurveWidgetRequest) SetRequestType(v RetentionCurveRequestType) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *RetentionCurveWidgetRequest) GetStyle() RetentionCurveStyle {
	if o == nil || o.Style == nil {
		var ret RetentionCurveStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionCurveWidgetRequest) GetStyleOk() (*RetentionCurveStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *RetentionCurveWidgetRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RetentionCurveStyle and assigns it to the Style field.
func (o *RetentionCurveWidgetRequest) SetStyle(v RetentionCurveStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionCurveWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionCurveWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query       *RetentionQuery            `json:"query"`
		RequestType *RetentionCurveRequestType `json:"request_type"`
		Style       *RetentionCurveStyle       `json:"style,omitempty"`
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
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
