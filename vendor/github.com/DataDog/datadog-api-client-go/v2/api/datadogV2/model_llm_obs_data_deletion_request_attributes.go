// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDataDeletionRequestAttributes Attributes for an LLM Observability data deletion request.
type LLMObsDataDeletionRequestAttributes struct {
	// Optional delay in seconds before the deletion is executed.
	Delay *int64 `json:"delay,omitempty"`
	// Start of the deletion time range in milliseconds since Unix epoch.
	From int64 `json:"from"`
	// Query filters selecting the data to delete. Must include a `query` key with an `@trace_id` filter.
	Query map[string]string `json:"query"`
	// End of the deletion time range in milliseconds since Unix epoch.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDataDeletionRequestAttributes instantiates a new LLMObsDataDeletionRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDataDeletionRequestAttributes(from int64, query map[string]string, to int64) *LLMObsDataDeletionRequestAttributes {
	this := LLMObsDataDeletionRequestAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewLLMObsDataDeletionRequestAttributesWithDefaults instantiates a new LLMObsDataDeletionRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDataDeletionRequestAttributesWithDefaults() *LLMObsDataDeletionRequestAttributes {
	this := LLMObsDataDeletionRequestAttributes{}
	return &this
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *LLMObsDataDeletionRequestAttributes) GetDelay() int64 {
	if o == nil || o.Delay == nil {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionRequestAttributes) GetDelayOk() (*int64, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *LLMObsDataDeletionRequestAttributes) HasDelay() bool {
	return o != nil && o.Delay != nil
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *LLMObsDataDeletionRequestAttributes) SetDelay(v int64) {
	o.Delay = &v
}

// GetFrom returns the From field value.
func (o *LLMObsDataDeletionRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *LLMObsDataDeletionRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *LLMObsDataDeletionRequestAttributes) GetQuery() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionRequestAttributes) GetQueryOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *LLMObsDataDeletionRequestAttributes) SetQuery(v map[string]string) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *LLMObsDataDeletionRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *LLMObsDataDeletionRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *LLMObsDataDeletionRequestAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDataDeletionRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDataDeletionRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Delay *int64             `json:"delay,omitempty"`
		From  *int64             `json:"from"`
		Query *map[string]string `json:"query"`
		To    *int64             `json:"to"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delay", "from", "query", "to"})
	} else {
		return err
	}
	o.Delay = all.Delay
	o.From = *all.From
	o.Query = *all.Query
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
