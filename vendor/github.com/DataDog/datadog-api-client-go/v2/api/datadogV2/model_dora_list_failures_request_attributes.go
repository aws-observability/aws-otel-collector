// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAListFailuresRequestAttributes Attributes to get a list of failures.
type DORAListFailuresRequestAttributes struct {
	// Minimum timestamp for requested events.
	From *time.Time `json:"from,omitempty"`
	// Maximum number of events in the response.
	Limit *int32 `json:"limit,omitempty"`
	// Search query with event platform syntax.
	Query *string `json:"query,omitempty"`
	// Sort order (prefixed with `-` for descending).
	Sort *string `json:"sort,omitempty"`
	// Maximum timestamp for requested events.
	To *time.Time `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORAListFailuresRequestAttributes instantiates a new DORAListFailuresRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAListFailuresRequestAttributes() *DORAListFailuresRequestAttributes {
	this := DORAListFailuresRequestAttributes{}
	var limit int32 = 10
	this.Limit = &limit
	return &this
}

// NewDORAListFailuresRequestAttributesWithDefaults instantiates a new DORAListFailuresRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAListFailuresRequestAttributesWithDefaults() *DORAListFailuresRequestAttributes {
	this := DORAListFailuresRequestAttributes{}
	var limit int32 = 10
	this.Limit = &limit
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *DORAListFailuresRequestAttributes) GetFrom() time.Time {
	if o == nil || o.From == nil {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestAttributes) GetFromOk() (*time.Time, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *DORAListFailuresRequestAttributes) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *DORAListFailuresRequestAttributes) SetFrom(v time.Time) {
	o.From = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DORAListFailuresRequestAttributes) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestAttributes) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DORAListFailuresRequestAttributes) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *DORAListFailuresRequestAttributes) SetLimit(v int32) {
	o.Limit = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DORAListFailuresRequestAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DORAListFailuresRequestAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DORAListFailuresRequestAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *DORAListFailuresRequestAttributes) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestAttributes) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *DORAListFailuresRequestAttributes) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *DORAListFailuresRequestAttributes) SetSort(v string) {
	o.Sort = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *DORAListFailuresRequestAttributes) GetTo() time.Time {
	if o == nil || o.To == nil {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORAListFailuresRequestAttributes) GetToOk() (*time.Time, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *DORAListFailuresRequestAttributes) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *DORAListFailuresRequestAttributes) SetTo(v time.Time) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAListFailuresRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		if o.From.Nanosecond() == 0 {
			toSerialize["from"] = o.From.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["from"] = o.From.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.To != nil {
		if o.To.Nanosecond() == 0 {
			toSerialize["to"] = o.To.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["to"] = o.To.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORAListFailuresRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From  *time.Time `json:"from,omitempty"`
		Limit *int32     `json:"limit,omitempty"`
		Query *string    `json:"query,omitempty"`
		Sort  *string    `json:"sort,omitempty"`
		To    *time.Time `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "limit", "query", "sort", "to"})
	} else {
		return err
	}
	o.From = all.From
	o.Limit = all.Limit
	o.Query = all.Query
	o.Sort = all.Sort
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
