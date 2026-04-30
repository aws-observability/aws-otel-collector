// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultNetpath Network Path test result capturing the path between source and destination.
type SyntheticsTestResultNetpath struct {
	// Destination endpoint of a network path measurement.
	Destination *SyntheticsTestResultNetpathDestination `json:"destination,omitempty"`
	// Hops along the network path.
	Hops []SyntheticsTestResultNetpathHop `json:"hops,omitempty"`
	// Origin of the network path (for example, probe source).
	Origin *string `json:"origin,omitempty"`
	// Identifier of the path trace.
	PathtraceId *string `json:"pathtrace_id,omitempty"`
	// Protocol used for the path trace (for example, `tcp`, `udp`, `icmp`).
	Protocol *string `json:"protocol,omitempty"`
	// Source endpoint of a network path measurement.
	Source *SyntheticsTestResultNetpathEndpoint `json:"source,omitempty"`
	// Tags associated with the network path measurement.
	Tags []string `json:"tags,omitempty"`
	// Unix timestamp (ms) of the network path measurement.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultNetpath instantiates a new SyntheticsTestResultNetpath object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultNetpath() *SyntheticsTestResultNetpath {
	this := SyntheticsTestResultNetpath{}
	return &this
}

// NewSyntheticsTestResultNetpathWithDefaults instantiates a new SyntheticsTestResultNetpath object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultNetpathWithDefaults() *SyntheticsTestResultNetpath {
	this := SyntheticsTestResultNetpath{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetDestination() SyntheticsTestResultNetpathDestination {
	if o == nil || o.Destination == nil {
		var ret SyntheticsTestResultNetpathDestination
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetDestinationOk() (*SyntheticsTestResultNetpathDestination, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasDestination() bool {
	return o != nil && o.Destination != nil
}

// SetDestination gets a reference to the given SyntheticsTestResultNetpathDestination and assigns it to the Destination field.
func (o *SyntheticsTestResultNetpath) SetDestination(v SyntheticsTestResultNetpathDestination) {
	o.Destination = &v
}

// GetHops returns the Hops field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetHops() []SyntheticsTestResultNetpathHop {
	if o == nil || o.Hops == nil {
		var ret []SyntheticsTestResultNetpathHop
		return ret
	}
	return o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetHopsOk() (*[]SyntheticsTestResultNetpathHop, bool) {
	if o == nil || o.Hops == nil {
		return nil, false
	}
	return &o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasHops() bool {
	return o != nil && o.Hops != nil
}

// SetHops gets a reference to the given []SyntheticsTestResultNetpathHop and assigns it to the Hops field.
func (o *SyntheticsTestResultNetpath) SetHops(v []SyntheticsTestResultNetpathHop) {
	o.Hops = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasOrigin() bool {
	return o != nil && o.Origin != nil
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *SyntheticsTestResultNetpath) SetOrigin(v string) {
	o.Origin = &v
}

// GetPathtraceId returns the PathtraceId field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetPathtraceId() string {
	if o == nil || o.PathtraceId == nil {
		var ret string
		return ret
	}
	return *o.PathtraceId
}

// GetPathtraceIdOk returns a tuple with the PathtraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetPathtraceIdOk() (*string, bool) {
	if o == nil || o.PathtraceId == nil {
		return nil, false
	}
	return o.PathtraceId, true
}

// HasPathtraceId returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasPathtraceId() bool {
	return o != nil && o.PathtraceId != nil
}

// SetPathtraceId gets a reference to the given string and assigns it to the PathtraceId field.
func (o *SyntheticsTestResultNetpath) SetPathtraceId(v string) {
	o.PathtraceId = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasProtocol() bool {
	return o != nil && o.Protocol != nil
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SyntheticsTestResultNetpath) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetSource() SyntheticsTestResultNetpathEndpoint {
	if o == nil || o.Source == nil {
		var ret SyntheticsTestResultNetpathEndpoint
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetSourceOk() (*SyntheticsTestResultNetpathEndpoint, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given SyntheticsTestResultNetpathEndpoint and assigns it to the Source field.
func (o *SyntheticsTestResultNetpath) SetSource(v SyntheticsTestResultNetpathEndpoint) {
	o.Source = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SyntheticsTestResultNetpath) SetTags(v []string) {
	o.Tags = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetpath) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetpath) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetpath) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *SyntheticsTestResultNetpath) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultNetpath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Hops != nil {
		toSerialize["hops"] = o.Hops
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.PathtraceId != nil {
		toSerialize["pathtrace_id"] = o.PathtraceId
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultNetpath) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destination *SyntheticsTestResultNetpathDestination `json:"destination,omitempty"`
		Hops        []SyntheticsTestResultNetpathHop        `json:"hops,omitempty"`
		Origin      *string                                 `json:"origin,omitempty"`
		PathtraceId *string                                 `json:"pathtrace_id,omitempty"`
		Protocol    *string                                 `json:"protocol,omitempty"`
		Source      *SyntheticsTestResultNetpathEndpoint    `json:"source,omitempty"`
		Tags        []string                                `json:"tags,omitempty"`
		Timestamp   *int64                                  `json:"timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination", "hops", "origin", "pathtrace_id", "protocol", "source", "tags", "timestamp"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Destination != nil && all.Destination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Destination = all.Destination
	o.Hops = all.Hops
	o.Origin = all.Origin
	o.PathtraceId = all.PathtraceId
	o.Protocol = all.Protocol
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	o.Tags = all.Tags
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
