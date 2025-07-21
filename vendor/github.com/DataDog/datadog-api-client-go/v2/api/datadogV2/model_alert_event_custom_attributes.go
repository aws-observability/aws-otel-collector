// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertEventCustomAttributes Alert event attributes.
type AlertEventCustomAttributes struct {
	// Free form JSON object for arbitrary data. Supports up to 100 properties per object and a maximum nesting depth of 10 levels.
	Custom map[string]interface{} `json:"custom,omitempty"`
	// The links related to the event. Maximum of 20 links allowed.
	Links []AlertEventCustomAttributesLinksItems `json:"links,omitempty"`
	// The priority of the alert.
	Priority *AlertEventCustomAttributesPriority `json:"priority,omitempty"`
	// The status of the alert.
	Status AlertEventCustomAttributesStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAlertEventCustomAttributes instantiates a new AlertEventCustomAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertEventCustomAttributes(status AlertEventCustomAttributesStatus) *AlertEventCustomAttributes {
	this := AlertEventCustomAttributes{}
	var priority AlertEventCustomAttributesPriority = ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_FIVE
	this.Priority = &priority
	this.Status = status
	return &this
}

// NewAlertEventCustomAttributesWithDefaults instantiates a new AlertEventCustomAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertEventCustomAttributesWithDefaults() *AlertEventCustomAttributes {
	this := AlertEventCustomAttributes{}
	var priority AlertEventCustomAttributesPriority = ALERTEVENTCUSTOMATTRIBUTESPRIORITY_PRIORITY_FIVE
	this.Priority = &priority
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *AlertEventCustomAttributes) GetCustom() map[string]interface{} {
	if o == nil || o.Custom == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributes) GetCustomOk() (*map[string]interface{}, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return &o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *AlertEventCustomAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *AlertEventCustomAttributes) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlertEventCustomAttributes) GetLinks() []AlertEventCustomAttributesLinksItems {
	if o == nil || o.Links == nil {
		var ret []AlertEventCustomAttributesLinksItems
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributes) GetLinksOk() (*[]AlertEventCustomAttributesLinksItems, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlertEventCustomAttributes) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []AlertEventCustomAttributesLinksItems and assigns it to the Links field.
func (o *AlertEventCustomAttributes) SetLinks(v []AlertEventCustomAttributesLinksItems) {
	o.Links = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AlertEventCustomAttributes) GetPriority() AlertEventCustomAttributesPriority {
	if o == nil || o.Priority == nil {
		var ret AlertEventCustomAttributesPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributes) GetPriorityOk() (*AlertEventCustomAttributesPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AlertEventCustomAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given AlertEventCustomAttributesPriority and assigns it to the Priority field.
func (o *AlertEventCustomAttributes) SetPriority(v AlertEventCustomAttributesPriority) {
	o.Priority = &v
}

// GetStatus returns the Status field value.
func (o *AlertEventCustomAttributes) GetStatus() AlertEventCustomAttributesStatus {
	if o == nil {
		var ret AlertEventCustomAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AlertEventCustomAttributes) GetStatusOk() (*AlertEventCustomAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AlertEventCustomAttributes) SetStatus(v AlertEventCustomAttributesStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertEventCustomAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["status"] = o.Status
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertEventCustomAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Custom   map[string]interface{}                 `json:"custom,omitempty"`
		Links    []AlertEventCustomAttributesLinksItems `json:"links,omitempty"`
		Priority *AlertEventCustomAttributesPriority    `json:"priority,omitempty"`
		Status   *AlertEventCustomAttributesStatus      `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}

	hasInvalidField := false
	o.Custom = all.Custom
	o.Links = all.Links
	if all.Priority != nil && !all.Priority.IsValid() {
		hasInvalidField = true
	} else {
		o.Priority = all.Priority
	}
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
