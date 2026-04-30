// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultNetstats Aggregated network statistics from the test execution.
type SyntheticsTestResultNetstats struct {
	// Statistics about the number of hops for a network test.
	Hops *SyntheticsTestResultNetstatsHops `json:"hops,omitempty"`
	// Network jitter in milliseconds.
	Jitter *float64 `json:"jitter,omitempty"`
	// Latency statistics for a network probe.
	Latency *SyntheticsTestResultNetworkLatency `json:"latency,omitempty"`
	// Percentage of probe packets lost.
	PacketLossPercentage *float64 `json:"packet_loss_percentage,omitempty"`
	// Number of probe packets received.
	PacketsReceived *int64 `json:"packets_received,omitempty"`
	// Number of probe packets sent.
	PacketsSent *int64 `json:"packets_sent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultNetstats instantiates a new SyntheticsTestResultNetstats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultNetstats() *SyntheticsTestResultNetstats {
	this := SyntheticsTestResultNetstats{}
	return &this
}

// NewSyntheticsTestResultNetstatsWithDefaults instantiates a new SyntheticsTestResultNetstats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultNetstatsWithDefaults() *SyntheticsTestResultNetstats {
	this := SyntheticsTestResultNetstats{}
	return &this
}

// GetHops returns the Hops field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstats) GetHops() SyntheticsTestResultNetstatsHops {
	if o == nil || o.Hops == nil {
		var ret SyntheticsTestResultNetstatsHops
		return ret
	}
	return *o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstats) GetHopsOk() (*SyntheticsTestResultNetstatsHops, bool) {
	if o == nil || o.Hops == nil {
		return nil, false
	}
	return o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstats) HasHops() bool {
	return o != nil && o.Hops != nil
}

// SetHops gets a reference to the given SyntheticsTestResultNetstatsHops and assigns it to the Hops field.
func (o *SyntheticsTestResultNetstats) SetHops(v SyntheticsTestResultNetstatsHops) {
	o.Hops = &v
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstats) GetJitter() float64 {
	if o == nil || o.Jitter == nil {
		var ret float64
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstats) GetJitterOk() (*float64, bool) {
	if o == nil || o.Jitter == nil {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstats) HasJitter() bool {
	return o != nil && o.Jitter != nil
}

// SetJitter gets a reference to the given float64 and assigns it to the Jitter field.
func (o *SyntheticsTestResultNetstats) SetJitter(v float64) {
	o.Jitter = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstats) GetLatency() SyntheticsTestResultNetworkLatency {
	if o == nil || o.Latency == nil {
		var ret SyntheticsTestResultNetworkLatency
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstats) GetLatencyOk() (*SyntheticsTestResultNetworkLatency, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstats) HasLatency() bool {
	return o != nil && o.Latency != nil
}

// SetLatency gets a reference to the given SyntheticsTestResultNetworkLatency and assigns it to the Latency field.
func (o *SyntheticsTestResultNetstats) SetLatency(v SyntheticsTestResultNetworkLatency) {
	o.Latency = &v
}

// GetPacketLossPercentage returns the PacketLossPercentage field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstats) GetPacketLossPercentage() float64 {
	if o == nil || o.PacketLossPercentage == nil {
		var ret float64
		return ret
	}
	return *o.PacketLossPercentage
}

// GetPacketLossPercentageOk returns a tuple with the PacketLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstats) GetPacketLossPercentageOk() (*float64, bool) {
	if o == nil || o.PacketLossPercentage == nil {
		return nil, false
	}
	return o.PacketLossPercentage, true
}

// HasPacketLossPercentage returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstats) HasPacketLossPercentage() bool {
	return o != nil && o.PacketLossPercentage != nil
}

// SetPacketLossPercentage gets a reference to the given float64 and assigns it to the PacketLossPercentage field.
func (o *SyntheticsTestResultNetstats) SetPacketLossPercentage(v float64) {
	o.PacketLossPercentage = &v
}

// GetPacketsReceived returns the PacketsReceived field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstats) GetPacketsReceived() int64 {
	if o == nil || o.PacketsReceived == nil {
		var ret int64
		return ret
	}
	return *o.PacketsReceived
}

// GetPacketsReceivedOk returns a tuple with the PacketsReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstats) GetPacketsReceivedOk() (*int64, bool) {
	if o == nil || o.PacketsReceived == nil {
		return nil, false
	}
	return o.PacketsReceived, true
}

// HasPacketsReceived returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstats) HasPacketsReceived() bool {
	return o != nil && o.PacketsReceived != nil
}

// SetPacketsReceived gets a reference to the given int64 and assigns it to the PacketsReceived field.
func (o *SyntheticsTestResultNetstats) SetPacketsReceived(v int64) {
	o.PacketsReceived = &v
}

// GetPacketsSent returns the PacketsSent field value if set, zero value otherwise.
func (o *SyntheticsTestResultNetstats) GetPacketsSent() int64 {
	if o == nil || o.PacketsSent == nil {
		var ret int64
		return ret
	}
	return *o.PacketsSent
}

// GetPacketsSentOk returns a tuple with the PacketsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultNetstats) GetPacketsSentOk() (*int64, bool) {
	if o == nil || o.PacketsSent == nil {
		return nil, false
	}
	return o.PacketsSent, true
}

// HasPacketsSent returns a boolean if a field has been set.
func (o *SyntheticsTestResultNetstats) HasPacketsSent() bool {
	return o != nil && o.PacketsSent != nil
}

// SetPacketsSent gets a reference to the given int64 and assigns it to the PacketsSent field.
func (o *SyntheticsTestResultNetstats) SetPacketsSent(v int64) {
	o.PacketsSent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultNetstats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hops != nil {
		toSerialize["hops"] = o.Hops
	}
	if o.Jitter != nil {
		toSerialize["jitter"] = o.Jitter
	}
	if o.Latency != nil {
		toSerialize["latency"] = o.Latency
	}
	if o.PacketLossPercentage != nil {
		toSerialize["packet_loss_percentage"] = o.PacketLossPercentage
	}
	if o.PacketsReceived != nil {
		toSerialize["packets_received"] = o.PacketsReceived
	}
	if o.PacketsSent != nil {
		toSerialize["packets_sent"] = o.PacketsSent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultNetstats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hops                 *SyntheticsTestResultNetstatsHops   `json:"hops,omitempty"`
		Jitter               *float64                            `json:"jitter,omitempty"`
		Latency              *SyntheticsTestResultNetworkLatency `json:"latency,omitempty"`
		PacketLossPercentage *float64                            `json:"packet_loss_percentage,omitempty"`
		PacketsReceived      *int64                              `json:"packets_received,omitempty"`
		PacketsSent          *int64                              `json:"packets_sent,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hops", "jitter", "latency", "packet_loss_percentage", "packets_received", "packets_sent"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Hops != nil && all.Hops.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Hops = all.Hops
	o.Jitter = all.Jitter
	if all.Latency != nil && all.Latency.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Latency = all.Latency
	o.PacketLossPercentage = all.PacketLossPercentage
	o.PacketsReceived = all.PacketsReceived
	o.PacketsSent = all.PacketsSent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
