// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityContextPage Pagination metadata for the entity context response.
type EntityContextPage struct {
	// An opaque token to pass as `page_token` in a subsequent request to retrieve the next page of results. Empty when there are no more results.
	NextToken string `json:"next_token"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityContextPage instantiates a new EntityContextPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityContextPage(nextToken string) *EntityContextPage {
	this := EntityContextPage{}
	this.NextToken = nextToken
	return &this
}

// NewEntityContextPageWithDefaults instantiates a new EntityContextPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityContextPageWithDefaults() *EntityContextPage {
	this := EntityContextPage{}
	return &this
}

// GetNextToken returns the NextToken field value.
func (o *EntityContextPage) GetNextToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value
// and a boolean to check if the value has been set.
func (o *EntityContextPage) GetNextTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextToken, true
}

// SetNextToken sets field value.
func (o *EntityContextPage) SetNextToken(v string) {
	o.NextToken = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityContextPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["next_token"] = o.NextToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityContextPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextToken *string `json:"next_token"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NextToken == nil {
		return fmt.Errorf("required field next_token missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_token"})
	} else {
		return err
	}
	o.NextToken = *all.NextToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
