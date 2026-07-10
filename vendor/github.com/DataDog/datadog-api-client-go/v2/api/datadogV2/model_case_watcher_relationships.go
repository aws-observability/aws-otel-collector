// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseWatcherRelationships Relationships for a case watcher, linking to the underlying user resource.
type CaseWatcherRelationships struct {
	// The user relationship for a case watcher.
	User CaseWatcherUserRelationship `json:"user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseWatcherRelationships instantiates a new CaseWatcherRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseWatcherRelationships(user CaseWatcherUserRelationship) *CaseWatcherRelationships {
	this := CaseWatcherRelationships{}
	this.User = user
	return &this
}

// NewCaseWatcherRelationshipsWithDefaults instantiates a new CaseWatcherRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseWatcherRelationshipsWithDefaults() *CaseWatcherRelationships {
	this := CaseWatcherRelationships{}
	return &this
}

// GetUser returns the User field value.
func (o *CaseWatcherRelationships) GetUser() CaseWatcherUserRelationship {
	if o == nil {
		var ret CaseWatcherUserRelationship
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *CaseWatcherRelationships) GetUserOk() (*CaseWatcherUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *CaseWatcherRelationships) SetUser(v CaseWatcherUserRelationship) {
	o.User = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseWatcherRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseWatcherRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		User *CaseWatcherUserRelationship `json:"user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = *all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
