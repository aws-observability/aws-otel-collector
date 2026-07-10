// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseWatcher Represents a user who is subscribed to notifications for a case. Watchers receive updates when the case's status, priority, assignee, or comments change.
type CaseWatcher struct {
	// The primary identifier of the case watcher.
	Id string `json:"id"`
	// Relationships for a case watcher, linking to the underlying user resource.
	Relationships CaseWatcherRelationships `json:"relationships"`
	// JSON:API resource type for case watchers.
	Type CaseWatcherResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseWatcher instantiates a new CaseWatcher object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseWatcher(id string, relationships CaseWatcherRelationships, typeVar CaseWatcherResourceType) *CaseWatcher {
	this := CaseWatcher{}
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewCaseWatcherWithDefaults instantiates a new CaseWatcher object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseWatcherWithDefaults() *CaseWatcher {
	this := CaseWatcher{}
	var typeVar CaseWatcherResourceType = CASEWATCHERRESOURCETYPE_WATCHER
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *CaseWatcher) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaseWatcher) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CaseWatcher) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *CaseWatcher) GetRelationships() CaseWatcherRelationships {
	if o == nil {
		var ret CaseWatcherRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *CaseWatcher) GetRelationshipsOk() (*CaseWatcherRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *CaseWatcher) SetRelationships(v CaseWatcherRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *CaseWatcher) GetType() CaseWatcherResourceType {
	if o == nil {
		var ret CaseWatcherResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseWatcher) GetTypeOk() (*CaseWatcherResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseWatcher) SetType(v CaseWatcherResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseWatcher) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseWatcher) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                   `json:"id"`
		Relationships *CaseWatcherRelationships `json:"relationships"`
		Type          *CaseWatcherResourceType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
