// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultTracerouteHop A network probe result, used for traceroute hops and ping summaries.
type SyntheticsTestResultTracerouteHop struct {
	// Target hostname.
	Host *string `json:"host,omitempty"`
	// Latency statistics for a network probe.
	Latency *SyntheticsTestResultNetworkLatency `json:"latency,omitempty"`
	// Percentage of probe packets lost.
	PacketLossPercentage *float64 `json:"packet_loss_percentage,omitempty"`
	// Size of each probe packet in bytes.
	PacketSize *int64 `json:"packet_size,omitempty"`
	// Number of probe packets received.
	PacketsReceived *int64 `json:"packets_received,omitempty"`
	// Number of probe packets sent.
	PacketsSent *int64 `json:"packets_sent,omitempty"`
	// Resolved IP address for the target.
	ResolvedIp *string `json:"resolved_ip,omitempty"`
	// List of intermediate routers for the traceroute.
	Routers []SyntheticsTestResultRouter `json:"routers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultTracerouteHop instantiates a new SyntheticsTestResultTracerouteHop object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultTracerouteHop() *SyntheticsTestResultTracerouteHop {
	this := SyntheticsTestResultTracerouteHop{}
	return &this
}

// NewSyntheticsTestResultTracerouteHopWithDefaults instantiates a new SyntheticsTestResultTracerouteHop object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultTracerouteHopWithDefaults() *SyntheticsTestResultTracerouteHop {
	this := SyntheticsTestResultTracerouteHop{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SyntheticsTestResultTracerouteHop) SetHost(v string) {
	o.Host = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetLatency() SyntheticsTestResultNetworkLatency {
	if o == nil || o.Latency == nil {
		var ret SyntheticsTestResultNetworkLatency
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetLatencyOk() (*SyntheticsTestResultNetworkLatency, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasLatency() bool {
	return o != nil && o.Latency != nil
}

// SetLatency gets a reference to the given SyntheticsTestResultNetworkLatency and assigns it to the Latency field.
func (o *SyntheticsTestResultTracerouteHop) SetLatency(v SyntheticsTestResultNetworkLatency) {
	o.Latency = &v
}

// GetPacketLossPercentage returns the PacketLossPercentage field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetPacketLossPercentage() float64 {
	if o == nil || o.PacketLossPercentage == nil {
		var ret float64
		return ret
	}
	return *o.PacketLossPercentage
}

// GetPacketLossPercentageOk returns a tuple with the PacketLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetPacketLossPercentageOk() (*float64, bool) {
	if o == nil || o.PacketLossPercentage == nil {
		return nil, false
	}
	return o.PacketLossPercentage, true
}

// HasPacketLossPercentage returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasPacketLossPercentage() bool {
	return o != nil && o.PacketLossPercentage != nil
}

// SetPacketLossPercentage gets a reference to the given float64 and assigns it to the PacketLossPercentage field.
func (o *SyntheticsTestResultTracerouteHop) SetPacketLossPercentage(v float64) {
	o.PacketLossPercentage = &v
}

// GetPacketSize returns the PacketSize field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetPacketSize() int64 {
	if o == nil || o.PacketSize == nil {
		var ret int64
		return ret
	}
	return *o.PacketSize
}

// GetPacketSizeOk returns a tuple with the PacketSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetPacketSizeOk() (*int64, bool) {
	if o == nil || o.PacketSize == nil {
		return nil, false
	}
	return o.PacketSize, true
}

// HasPacketSize returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasPacketSize() bool {
	return o != nil && o.PacketSize != nil
}

// SetPacketSize gets a reference to the given int64 and assigns it to the PacketSize field.
func (o *SyntheticsTestResultTracerouteHop) SetPacketSize(v int64) {
	o.PacketSize = &v
}

// GetPacketsReceived returns the PacketsReceived field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetPacketsReceived() int64 {
	if o == nil || o.PacketsReceived == nil {
		var ret int64
		return ret
	}
	return *o.PacketsReceived
}

// GetPacketsReceivedOk returns a tuple with the PacketsReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetPacketsReceivedOk() (*int64, bool) {
	if o == nil || o.PacketsReceived == nil {
		return nil, false
	}
	return o.PacketsReceived, true
}

// HasPacketsReceived returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasPacketsReceived() bool {
	return o != nil && o.PacketsReceived != nil
}

// SetPacketsReceived gets a reference to the given int64 and assigns it to the PacketsReceived field.
func (o *SyntheticsTestResultTracerouteHop) SetPacketsReceived(v int64) {
	o.PacketsReceived = &v
}

// GetPacketsSent returns the PacketsSent field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetPacketsSent() int64 {
	if o == nil || o.PacketsSent == nil {
		var ret int64
		return ret
	}
	return *o.PacketsSent
}

// GetPacketsSentOk returns a tuple with the PacketsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetPacketsSentOk() (*int64, bool) {
	if o == nil || o.PacketsSent == nil {
		return nil, false
	}
	return o.PacketsSent, true
}

// HasPacketsSent returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasPacketsSent() bool {
	return o != nil && o.PacketsSent != nil
}

// SetPacketsSent gets a reference to the given int64 and assigns it to the PacketsSent field.
func (o *SyntheticsTestResultTracerouteHop) SetPacketsSent(v int64) {
	o.PacketsSent = &v
}

// GetResolvedIp returns the ResolvedIp field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetResolvedIp() string {
	if o == nil || o.ResolvedIp == nil {
		var ret string
		return ret
	}
	return *o.ResolvedIp
}

// GetResolvedIpOk returns a tuple with the ResolvedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetResolvedIpOk() (*string, bool) {
	if o == nil || o.ResolvedIp == nil {
		return nil, false
	}
	return o.ResolvedIp, true
}

// HasResolvedIp returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasResolvedIp() bool {
	return o != nil && o.ResolvedIp != nil
}

// SetResolvedIp gets a reference to the given string and assigns it to the ResolvedIp field.
func (o *SyntheticsTestResultTracerouteHop) SetResolvedIp(v string) {
	o.ResolvedIp = &v
}

// GetRouters returns the Routers field value if set, zero value otherwise.
func (o *SyntheticsTestResultTracerouteHop) GetRouters() []SyntheticsTestResultRouter {
	if o == nil || o.Routers == nil {
		var ret []SyntheticsTestResultRouter
		return ret
	}
	return o.Routers
}

// GetRoutersOk returns a tuple with the Routers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTracerouteHop) GetRoutersOk() (*[]SyntheticsTestResultRouter, bool) {
	if o == nil || o.Routers == nil {
		return nil, false
	}
	return &o.Routers, true
}

// HasRouters returns a boolean if a field has been set.
func (o *SyntheticsTestResultTracerouteHop) HasRouters() bool {
	return o != nil && o.Routers != nil
}

// SetRouters gets a reference to the given []SyntheticsTestResultRouter and assigns it to the Routers field.
func (o *SyntheticsTestResultTracerouteHop) SetRouters(v []SyntheticsTestResultRouter) {
	o.Routers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultTracerouteHop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Latency != nil {
		toSerialize["latency"] = o.Latency
	}
	if o.PacketLossPercentage != nil {
		toSerialize["packet_loss_percentage"] = o.PacketLossPercentage
	}
	if o.PacketSize != nil {
		toSerialize["packet_size"] = o.PacketSize
	}
	if o.PacketsReceived != nil {
		toSerialize["packets_received"] = o.PacketsReceived
	}
	if o.PacketsSent != nil {
		toSerialize["packets_sent"] = o.PacketsSent
	}
	if o.ResolvedIp != nil {
		toSerialize["resolved_ip"] = o.ResolvedIp
	}
	if o.Routers != nil {
		toSerialize["routers"] = o.Routers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultTracerouteHop) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Host                 *string                             `json:"host,omitempty"`
		Latency              *SyntheticsTestResultNetworkLatency `json:"latency,omitempty"`
		PacketLossPercentage *float64                            `json:"packet_loss_percentage,omitempty"`
		PacketSize           *int64                              `json:"packet_size,omitempty"`
		PacketsReceived      *int64                              `json:"packets_received,omitempty"`
		PacketsSent          *int64                              `json:"packets_sent,omitempty"`
		ResolvedIp           *string                             `json:"resolved_ip,omitempty"`
		Routers              []SyntheticsTestResultRouter        `json:"routers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"host", "latency", "packet_loss_percentage", "packet_size", "packets_received", "packets_sent", "resolved_ip", "routers"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Host = all.Host
	if all.Latency != nil && all.Latency.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Latency = all.Latency
	o.PacketLossPercentage = all.PacketLossPercentage
	o.PacketSize = all.PacketSize
	o.PacketsReceived = all.PacketsReceived
	o.PacketsSent = all.PacketsSent
	o.ResolvedIp = all.ResolvedIp
	o.Routers = all.Routers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
