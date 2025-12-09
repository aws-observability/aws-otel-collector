// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsIndexesOrder Object containing the ordered list of log index names.
type LogsIndexesOrder struct {
	// Array of strings identifying by their name(s) the index(es) of your organization.
	// Logs are tested against the query filter of each index one by one, following the order of the array.
	// Logs are eventually stored in the first matching index.
	IndexNames []string `json:"index_names"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsIndexesOrder instantiates a new LogsIndexesOrder object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsIndexesOrder(indexNames []string) *LogsIndexesOrder {
	this := LogsIndexesOrder{}
	this.IndexNames = indexNames
	return &this
}

// NewLogsIndexesOrderWithDefaults instantiates a new LogsIndexesOrder object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsIndexesOrderWithDefaults() *LogsIndexesOrder {
	this := LogsIndexesOrder{}
	return &this
}

// GetIndexNames returns the IndexNames field value.
func (o *LogsIndexesOrder) GetIndexNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IndexNames
}

// GetIndexNamesOk returns a tuple with the IndexNames field value
// and a boolean to check if the value has been set.
func (o *LogsIndexesOrder) GetIndexNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexNames, true
}

// SetIndexNames sets field value.
func (o *LogsIndexesOrder) SetIndexNames(v []string) {
	o.IndexNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsIndexesOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["index_names"] = o.IndexNames

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsIndexesOrder) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IndexNames *[]string `json:"index_names"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IndexNames == nil {
		return fmt.Errorf("required field index_names missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"index_names"})
	} else {
		return err
	}
	o.IndexNames = *all.IndexNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
