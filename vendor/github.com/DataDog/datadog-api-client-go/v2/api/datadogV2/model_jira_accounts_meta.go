// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraAccountsMeta Metadata for Jira accounts response
type JiraAccountsMeta struct {
	// Public key for the Jira integration
	PublicKey *string `json:"public_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraAccountsMeta instantiates a new JiraAccountsMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraAccountsMeta() *JiraAccountsMeta {
	this := JiraAccountsMeta{}
	return &this
}

// NewJiraAccountsMetaWithDefaults instantiates a new JiraAccountsMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraAccountsMetaWithDefaults() *JiraAccountsMeta {
	this := JiraAccountsMeta{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *JiraAccountsMeta) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraAccountsMeta) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *JiraAccountsMeta) HasPublicKey() bool {
	return o != nil && o.PublicKey != nil
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *JiraAccountsMeta) SetPublicKey(v string) {
	o.PublicKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraAccountsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraAccountsMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PublicKey *string `json:"public_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"public_key"})
	} else {
		return err
	}
	o.PublicKey = all.PublicKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
