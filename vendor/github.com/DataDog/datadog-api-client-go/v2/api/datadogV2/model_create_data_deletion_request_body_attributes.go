// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDataDeletionRequestBodyAttributes Attributes for creating a data deletion request.
type CreateDataDeletionRequestBodyAttributes struct {
	// Start of requested time window, milliseconds since Unix epoch.
	From int64 `json:"from"`
	// List of indexes for the search. If not provided, the search is performed in all indexes.
	Indexes []string `json:"indexes,omitempty"`
	// Query for creating a data deletion request.
	Query map[string]string `json:"query"`
	// End of requested time window, milliseconds since Unix epoch.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateDataDeletionRequestBodyAttributes instantiates a new CreateDataDeletionRequestBodyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateDataDeletionRequestBodyAttributes(from int64, query map[string]string, to int64) *CreateDataDeletionRequestBodyAttributes {
	this := CreateDataDeletionRequestBodyAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewCreateDataDeletionRequestBodyAttributesWithDefaults instantiates a new CreateDataDeletionRequestBodyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateDataDeletionRequestBodyAttributesWithDefaults() *CreateDataDeletionRequestBodyAttributes {
	this := CreateDataDeletionRequestBodyAttributes{}
	return &this
}

// GetFrom returns the From field value.
func (o *CreateDataDeletionRequestBodyAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateDataDeletionRequestBodyAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *CreateDataDeletionRequestBodyAttributes) SetFrom(v int64) {
	o.From = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *CreateDataDeletionRequestBodyAttributes) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataDeletionRequestBodyAttributes) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *CreateDataDeletionRequestBodyAttributes) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *CreateDataDeletionRequestBodyAttributes) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQuery returns the Query field value.
func (o *CreateDataDeletionRequestBodyAttributes) GetQuery() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CreateDataDeletionRequestBodyAttributes) GetQueryOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *CreateDataDeletionRequestBodyAttributes) SetQuery(v map[string]string) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *CreateDataDeletionRequestBodyAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *CreateDataDeletionRequestBodyAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *CreateDataDeletionRequestBodyAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateDataDeletionRequestBodyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateDataDeletionRequestBodyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From    *int64             `json:"from"`
		Indexes []string           `json:"indexes,omitempty"`
		Query   *map[string]string `json:"query"`
		To      *int64             `json:"to"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "indexes", "query", "to"})
	} else {
		return err
	}
	o.From = *all.From
	o.Indexes = all.Indexes
	o.Query = *all.Query
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
