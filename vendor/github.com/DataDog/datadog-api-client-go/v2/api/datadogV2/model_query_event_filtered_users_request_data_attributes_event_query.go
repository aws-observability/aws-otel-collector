// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryEventFilteredUsersRequestDataAttributesEventQuery
type QueryEventFilteredUsersRequestDataAttributesEventQuery struct {
	//
	Query *string `json:"query,omitempty"`
	//
	TimeFrame *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame `json:"time_frame,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueryEventFilteredUsersRequestDataAttributesEventQuery instantiates a new QueryEventFilteredUsersRequestDataAttributesEventQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryEventFilteredUsersRequestDataAttributesEventQuery() *QueryEventFilteredUsersRequestDataAttributesEventQuery {
	this := QueryEventFilteredUsersRequestDataAttributesEventQuery{}
	return &this
}

// NewQueryEventFilteredUsersRequestDataAttributesEventQueryWithDefaults instantiates a new QueryEventFilteredUsersRequestDataAttributesEventQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryEventFilteredUsersRequestDataAttributesEventQueryWithDefaults() *QueryEventFilteredUsersRequestDataAttributesEventQuery {
	this := QueryEventFilteredUsersRequestDataAttributesEventQuery{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) SetQuery(v string) {
	o.Query = &v
}

// GetTimeFrame returns the TimeFrame field value if set, zero value otherwise.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) GetTimeFrame() QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame {
	if o == nil || o.TimeFrame == nil {
		var ret QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame
		return ret
	}
	return *o.TimeFrame
}

// GetTimeFrameOk returns a tuple with the TimeFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) GetTimeFrameOk() (*QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame, bool) {
	if o == nil || o.TimeFrame == nil {
		return nil, false
	}
	return o.TimeFrame, true
}

// HasTimeFrame returns a boolean if a field has been set.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) HasTimeFrame() bool {
	return o != nil && o.TimeFrame != nil
}

// SetTimeFrame gets a reference to the given QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame and assigns it to the TimeFrame field.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) SetTimeFrame(v QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame) {
	o.TimeFrame = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryEventFilteredUsersRequestDataAttributesEventQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.TimeFrame != nil {
		toSerialize["time_frame"] = o.TimeFrame
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryEventFilteredUsersRequestDataAttributesEventQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query     *string                                                          `json:"query,omitempty"`
		TimeFrame *QueryEventFilteredUsersRequestDataAttributesEventQueryTimeFrame `json:"time_frame,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query", "time_frame"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Query = all.Query
	if all.TimeFrame != nil && all.TimeFrame.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeFrame = all.TimeFrame

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
