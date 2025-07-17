// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JobDefinitionFromRule Definition of a historical job based on a security monitoring rule.
type JobDefinitionFromRule struct {
	// Starting time of data analyzed by the job.
	From int64 `json:"from"`
	// ID of the detection rule used to create the job.
	Id string `json:"id"`
	// Index used to load the data.
	Index string `json:"index"`
	// Notifications sent when the job is completed.
	Notifications []string `json:"notifications,omitempty"`
	// Ending time of data analyzed by the job.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJobDefinitionFromRule instantiates a new JobDefinitionFromRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJobDefinitionFromRule(from int64, id string, index string, to int64) *JobDefinitionFromRule {
	this := JobDefinitionFromRule{}
	this.From = from
	this.Id = id
	this.Index = index
	this.To = to
	return &this
}

// NewJobDefinitionFromRuleWithDefaults instantiates a new JobDefinitionFromRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJobDefinitionFromRuleWithDefaults() *JobDefinitionFromRule {
	this := JobDefinitionFromRule{}
	return &this
}

// GetFrom returns the From field value.
func (o *JobDefinitionFromRule) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *JobDefinitionFromRule) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *JobDefinitionFromRule) SetFrom(v int64) {
	o.From = v
}

// GetId returns the Id field value.
func (o *JobDefinitionFromRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobDefinitionFromRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *JobDefinitionFromRule) SetId(v string) {
	o.Id = v
}

// GetIndex returns the Index field value.
func (o *JobDefinitionFromRule) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *JobDefinitionFromRule) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *JobDefinitionFromRule) SetIndex(v string) {
	o.Index = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *JobDefinitionFromRule) GetNotifications() []string {
	if o == nil || o.Notifications == nil {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinitionFromRule) GetNotificationsOk() (*[]string, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *JobDefinitionFromRule) HasNotifications() bool {
	return o != nil && o.Notifications != nil
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *JobDefinitionFromRule) SetNotifications(v []string) {
	o.Notifications = v
}

// GetTo returns the To field value.
func (o *JobDefinitionFromRule) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *JobDefinitionFromRule) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *JobDefinitionFromRule) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JobDefinitionFromRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	toSerialize["id"] = o.Id
	toSerialize["index"] = o.Index
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JobDefinitionFromRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From          *int64   `json:"from"`
		Id            *string  `json:"id"`
		Index         *string  `json:"index"`
		Notifications []string `json:"notifications,omitempty"`
		To            *int64   `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Index == nil {
		return fmt.Errorf("required field index missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "id", "index", "notifications", "to"})
	} else {
		return err
	}
	o.From = *all.From
	o.Id = *all.Id
	o.Index = *all.Index
	o.Notifications = all.Notifications
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
