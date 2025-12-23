// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateJiraIssueRequestArray List of requests to create Jira issues for security findings.
type CreateJiraIssueRequestArray struct {
	//
	Data []CreateJiraIssueRequestData `json:"data"`
	//
	Included []CreateJiraIssueRequestArrayIncluded `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateJiraIssueRequestArray instantiates a new CreateJiraIssueRequestArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateJiraIssueRequestArray(data []CreateJiraIssueRequestData) *CreateJiraIssueRequestArray {
	this := CreateJiraIssueRequestArray{}
	this.Data = data
	return &this
}

// NewCreateJiraIssueRequestArrayWithDefaults instantiates a new CreateJiraIssueRequestArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateJiraIssueRequestArrayWithDefaults() *CreateJiraIssueRequestArray {
	this := CreateJiraIssueRequestArray{}
	return &this
}

// GetData returns the Data field value.
func (o *CreateJiraIssueRequestArray) GetData() []CreateJiraIssueRequestData {
	if o == nil {
		var ret []CreateJiraIssueRequestData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestArray) GetDataOk() (*[]CreateJiraIssueRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *CreateJiraIssueRequestArray) SetData(v []CreateJiraIssueRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CreateJiraIssueRequestArray) GetIncluded() []CreateJiraIssueRequestArrayIncluded {
	if o == nil || o.Included == nil {
		var ret []CreateJiraIssueRequestArrayIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJiraIssueRequestArray) GetIncludedOk() (*[]CreateJiraIssueRequestArrayIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CreateJiraIssueRequestArray) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []CreateJiraIssueRequestArrayIncluded and assigns it to the Included field.
func (o *CreateJiraIssueRequestArray) SetIncluded(v []CreateJiraIssueRequestArrayIncluded) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateJiraIssueRequestArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateJiraIssueRequestArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]CreateJiraIssueRequestData         `json:"data"`
		Included []CreateJiraIssueRequestArrayIncluded `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
