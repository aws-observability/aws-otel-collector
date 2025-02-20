// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeCreateResponse A created EventBridge source.
type AWSEventBridgeCreateResponse struct {
	// The event source name.
	EventSourceName *string `json:"event_source_name,omitempty"`
	// True if the event bus was created in addition to the source.
	HasBus *bool `json:"has_bus,omitempty"`
	// The event source's [AWS region](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints).
	Region *string `json:"region,omitempty"`
	// The event source status "created".
	Status *AWSEventBridgeCreateStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeCreateResponse instantiates a new AWSEventBridgeCreateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeCreateResponse() *AWSEventBridgeCreateResponse {
	this := AWSEventBridgeCreateResponse{}
	return &this
}

// NewAWSEventBridgeCreateResponseWithDefaults instantiates a new AWSEventBridgeCreateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeCreateResponseWithDefaults() *AWSEventBridgeCreateResponse {
	this := AWSEventBridgeCreateResponse{}
	return &this
}

// GetEventSourceName returns the EventSourceName field value if set, zero value otherwise.
func (o *AWSEventBridgeCreateResponse) GetEventSourceName() string {
	if o == nil || o.EventSourceName == nil {
		var ret string
		return ret
	}
	return *o.EventSourceName
}

// GetEventSourceNameOk returns a tuple with the EventSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateResponse) GetEventSourceNameOk() (*string, bool) {
	if o == nil || o.EventSourceName == nil {
		return nil, false
	}
	return o.EventSourceName, true
}

// HasEventSourceName returns a boolean if a field has been set.
func (o *AWSEventBridgeCreateResponse) HasEventSourceName() bool {
	return o != nil && o.EventSourceName != nil
}

// SetEventSourceName gets a reference to the given string and assigns it to the EventSourceName field.
func (o *AWSEventBridgeCreateResponse) SetEventSourceName(v string) {
	o.EventSourceName = &v
}

// GetHasBus returns the HasBus field value if set, zero value otherwise.
func (o *AWSEventBridgeCreateResponse) GetHasBus() bool {
	if o == nil || o.HasBus == nil {
		var ret bool
		return ret
	}
	return *o.HasBus
}

// GetHasBusOk returns a tuple with the HasBus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateResponse) GetHasBusOk() (*bool, bool) {
	if o == nil || o.HasBus == nil {
		return nil, false
	}
	return o.HasBus, true
}

// HasHasBus returns a boolean if a field has been set.
func (o *AWSEventBridgeCreateResponse) HasHasBus() bool {
	return o != nil && o.HasBus != nil
}

// SetHasBus gets a reference to the given bool and assigns it to the HasBus field.
func (o *AWSEventBridgeCreateResponse) SetHasBus(v bool) {
	o.HasBus = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AWSEventBridgeCreateResponse) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateResponse) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AWSEventBridgeCreateResponse) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AWSEventBridgeCreateResponse) SetRegion(v string) {
	o.Region = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AWSEventBridgeCreateResponse) GetStatus() AWSEventBridgeCreateStatus {
	if o == nil || o.Status == nil {
		var ret AWSEventBridgeCreateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateResponse) GetStatusOk() (*AWSEventBridgeCreateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AWSEventBridgeCreateResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AWSEventBridgeCreateStatus and assigns it to the Status field.
func (o *AWSEventBridgeCreateResponse) SetStatus(v AWSEventBridgeCreateStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventSourceName != nil {
		toSerialize["event_source_name"] = o.EventSourceName
	}
	if o.HasBus != nil {
		toSerialize["has_bus"] = o.HasBus
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventSourceName *string                     `json:"event_source_name,omitempty"`
		HasBus          *bool                       `json:"has_bus,omitempty"`
		Region          *string                     `json:"region,omitempty"`
		Status          *AWSEventBridgeCreateStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_source_name", "has_bus", "region", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EventSourceName = all.EventSourceName
	o.HasBus = all.HasBus
	o.Region = all.Region
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
