// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SingleAggregatedConnectionResponseDataAttributes Attributes for an aggregated connection.
type SingleAggregatedConnectionResponseDataAttributes struct {
	// The total number of bytes sent by the client over the given period.
	BytesSentByClient *int64 `json:"bytes_sent_by_client,omitempty"`
	// The total number of bytes sent by the server over the given period.
	BytesSentByServer *int64 `json:"bytes_sent_by_server,omitempty"`
	// The key, value pairs for each group by.
	GroupBys map[string][]string `json:"group_bys,omitempty"`
	// The total number of packets sent by the client over the given period.
	PacketsSentByClient *int64 `json:"packets_sent_by_client,omitempty"`
	// The total number of packets sent by the server over the given period.
	PacketsSentByServer *int64 `json:"packets_sent_by_server,omitempty"`
	// Measured as TCP smoothed round trip time in microseconds (the time between a TCP frame being sent and acknowledged).
	RttMicroSeconds *int64 `json:"rtt_micro_seconds,omitempty"`
	// The number of TCP connections in a closed state. Measured in connections per second from the client.
	TcpClosedConnections *int64 `json:"tcp_closed_connections,omitempty"`
	// The number of TCP connections in an established state. Measured in connections per second from the client.
	TcpEstablishedConnections *int64 `json:"tcp_established_connections,omitempty"`
	// The number of TCP connections that were refused by the server. Typically this indicates an attempt to connect to an IP/port that is not receiving connections, or a firewall/security misconfiguration.
	TcpRefusals *int64 `json:"tcp_refusals,omitempty"`
	// The number of TCP connections that were reset by the server.
	TcpResets *int64 `json:"tcp_resets,omitempty"`
	// TCP Retransmits represent detected failures that are retransmitted to ensure delivery. Measured in count of retransmits from the client.
	TcpRetransmits *int64 `json:"tcp_retransmits,omitempty"`
	// The number of TCP connections that timed out from the perspective of the operating system. This can indicate general connectivity and latency issues.
	TcpTimeouts *int64 `json:"tcp_timeouts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSingleAggregatedConnectionResponseDataAttributes instantiates a new SingleAggregatedConnectionResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSingleAggregatedConnectionResponseDataAttributes() *SingleAggregatedConnectionResponseDataAttributes {
	this := SingleAggregatedConnectionResponseDataAttributes{}
	return &this
}

// NewSingleAggregatedConnectionResponseDataAttributesWithDefaults instantiates a new SingleAggregatedConnectionResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSingleAggregatedConnectionResponseDataAttributesWithDefaults() *SingleAggregatedConnectionResponseDataAttributes {
	this := SingleAggregatedConnectionResponseDataAttributes{}
	return &this
}

// GetBytesSentByClient returns the BytesSentByClient field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetBytesSentByClient() int64 {
	if o == nil || o.BytesSentByClient == nil {
		var ret int64
		return ret
	}
	return *o.BytesSentByClient
}

// GetBytesSentByClientOk returns a tuple with the BytesSentByClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetBytesSentByClientOk() (*int64, bool) {
	if o == nil || o.BytesSentByClient == nil {
		return nil, false
	}
	return o.BytesSentByClient, true
}

// HasBytesSentByClient returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasBytesSentByClient() bool {
	return o != nil && o.BytesSentByClient != nil
}

// SetBytesSentByClient gets a reference to the given int64 and assigns it to the BytesSentByClient field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetBytesSentByClient(v int64) {
	o.BytesSentByClient = &v
}

// GetBytesSentByServer returns the BytesSentByServer field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetBytesSentByServer() int64 {
	if o == nil || o.BytesSentByServer == nil {
		var ret int64
		return ret
	}
	return *o.BytesSentByServer
}

// GetBytesSentByServerOk returns a tuple with the BytesSentByServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetBytesSentByServerOk() (*int64, bool) {
	if o == nil || o.BytesSentByServer == nil {
		return nil, false
	}
	return o.BytesSentByServer, true
}

// HasBytesSentByServer returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasBytesSentByServer() bool {
	return o != nil && o.BytesSentByServer != nil
}

// SetBytesSentByServer gets a reference to the given int64 and assigns it to the BytesSentByServer field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetBytesSentByServer(v int64) {
	o.BytesSentByServer = &v
}

// GetGroupBys returns the GroupBys field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetGroupBys() map[string][]string {
	if o == nil || o.GroupBys == nil {
		var ret map[string][]string
		return ret
	}
	return o.GroupBys
}

// GetGroupBysOk returns a tuple with the GroupBys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetGroupBysOk() (*map[string][]string, bool) {
	if o == nil || o.GroupBys == nil {
		return nil, false
	}
	return &o.GroupBys, true
}

// HasGroupBys returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasGroupBys() bool {
	return o != nil && o.GroupBys != nil
}

// SetGroupBys gets a reference to the given map[string][]string and assigns it to the GroupBys field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetGroupBys(v map[string][]string) {
	o.GroupBys = v
}

// GetPacketsSentByClient returns the PacketsSentByClient field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetPacketsSentByClient() int64 {
	if o == nil || o.PacketsSentByClient == nil {
		var ret int64
		return ret
	}
	return *o.PacketsSentByClient
}

// GetPacketsSentByClientOk returns a tuple with the PacketsSentByClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetPacketsSentByClientOk() (*int64, bool) {
	if o == nil || o.PacketsSentByClient == nil {
		return nil, false
	}
	return o.PacketsSentByClient, true
}

// HasPacketsSentByClient returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasPacketsSentByClient() bool {
	return o != nil && o.PacketsSentByClient != nil
}

// SetPacketsSentByClient gets a reference to the given int64 and assigns it to the PacketsSentByClient field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetPacketsSentByClient(v int64) {
	o.PacketsSentByClient = &v
}

// GetPacketsSentByServer returns the PacketsSentByServer field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetPacketsSentByServer() int64 {
	if o == nil || o.PacketsSentByServer == nil {
		var ret int64
		return ret
	}
	return *o.PacketsSentByServer
}

// GetPacketsSentByServerOk returns a tuple with the PacketsSentByServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetPacketsSentByServerOk() (*int64, bool) {
	if o == nil || o.PacketsSentByServer == nil {
		return nil, false
	}
	return o.PacketsSentByServer, true
}

// HasPacketsSentByServer returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasPacketsSentByServer() bool {
	return o != nil && o.PacketsSentByServer != nil
}

// SetPacketsSentByServer gets a reference to the given int64 and assigns it to the PacketsSentByServer field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetPacketsSentByServer(v int64) {
	o.PacketsSentByServer = &v
}

// GetRttMicroSeconds returns the RttMicroSeconds field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetRttMicroSeconds() int64 {
	if o == nil || o.RttMicroSeconds == nil {
		var ret int64
		return ret
	}
	return *o.RttMicroSeconds
}

// GetRttMicroSecondsOk returns a tuple with the RttMicroSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetRttMicroSecondsOk() (*int64, bool) {
	if o == nil || o.RttMicroSeconds == nil {
		return nil, false
	}
	return o.RttMicroSeconds, true
}

// HasRttMicroSeconds returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasRttMicroSeconds() bool {
	return o != nil && o.RttMicroSeconds != nil
}

// SetRttMicroSeconds gets a reference to the given int64 and assigns it to the RttMicroSeconds field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetRttMicroSeconds(v int64) {
	o.RttMicroSeconds = &v
}

// GetTcpClosedConnections returns the TcpClosedConnections field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpClosedConnections() int64 {
	if o == nil || o.TcpClosedConnections == nil {
		var ret int64
		return ret
	}
	return *o.TcpClosedConnections
}

// GetTcpClosedConnectionsOk returns a tuple with the TcpClosedConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpClosedConnectionsOk() (*int64, bool) {
	if o == nil || o.TcpClosedConnections == nil {
		return nil, false
	}
	return o.TcpClosedConnections, true
}

// HasTcpClosedConnections returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasTcpClosedConnections() bool {
	return o != nil && o.TcpClosedConnections != nil
}

// SetTcpClosedConnections gets a reference to the given int64 and assigns it to the TcpClosedConnections field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetTcpClosedConnections(v int64) {
	o.TcpClosedConnections = &v
}

// GetTcpEstablishedConnections returns the TcpEstablishedConnections field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpEstablishedConnections() int64 {
	if o == nil || o.TcpEstablishedConnections == nil {
		var ret int64
		return ret
	}
	return *o.TcpEstablishedConnections
}

// GetTcpEstablishedConnectionsOk returns a tuple with the TcpEstablishedConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpEstablishedConnectionsOk() (*int64, bool) {
	if o == nil || o.TcpEstablishedConnections == nil {
		return nil, false
	}
	return o.TcpEstablishedConnections, true
}

// HasTcpEstablishedConnections returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasTcpEstablishedConnections() bool {
	return o != nil && o.TcpEstablishedConnections != nil
}

// SetTcpEstablishedConnections gets a reference to the given int64 and assigns it to the TcpEstablishedConnections field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetTcpEstablishedConnections(v int64) {
	o.TcpEstablishedConnections = &v
}

// GetTcpRefusals returns the TcpRefusals field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpRefusals() int64 {
	if o == nil || o.TcpRefusals == nil {
		var ret int64
		return ret
	}
	return *o.TcpRefusals
}

// GetTcpRefusalsOk returns a tuple with the TcpRefusals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpRefusalsOk() (*int64, bool) {
	if o == nil || o.TcpRefusals == nil {
		return nil, false
	}
	return o.TcpRefusals, true
}

// HasTcpRefusals returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasTcpRefusals() bool {
	return o != nil && o.TcpRefusals != nil
}

// SetTcpRefusals gets a reference to the given int64 and assigns it to the TcpRefusals field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetTcpRefusals(v int64) {
	o.TcpRefusals = &v
}

// GetTcpResets returns the TcpResets field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpResets() int64 {
	if o == nil || o.TcpResets == nil {
		var ret int64
		return ret
	}
	return *o.TcpResets
}

// GetTcpResetsOk returns a tuple with the TcpResets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpResetsOk() (*int64, bool) {
	if o == nil || o.TcpResets == nil {
		return nil, false
	}
	return o.TcpResets, true
}

// HasTcpResets returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasTcpResets() bool {
	return o != nil && o.TcpResets != nil
}

// SetTcpResets gets a reference to the given int64 and assigns it to the TcpResets field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetTcpResets(v int64) {
	o.TcpResets = &v
}

// GetTcpRetransmits returns the TcpRetransmits field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpRetransmits() int64 {
	if o == nil || o.TcpRetransmits == nil {
		var ret int64
		return ret
	}
	return *o.TcpRetransmits
}

// GetTcpRetransmitsOk returns a tuple with the TcpRetransmits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpRetransmitsOk() (*int64, bool) {
	if o == nil || o.TcpRetransmits == nil {
		return nil, false
	}
	return o.TcpRetransmits, true
}

// HasTcpRetransmits returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasTcpRetransmits() bool {
	return o != nil && o.TcpRetransmits != nil
}

// SetTcpRetransmits gets a reference to the given int64 and assigns it to the TcpRetransmits field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetTcpRetransmits(v int64) {
	o.TcpRetransmits = &v
}

// GetTcpTimeouts returns the TcpTimeouts field value if set, zero value otherwise.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpTimeouts() int64 {
	if o == nil || o.TcpTimeouts == nil {
		var ret int64
		return ret
	}
	return *o.TcpTimeouts
}

// GetTcpTimeoutsOk returns a tuple with the TcpTimeouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) GetTcpTimeoutsOk() (*int64, bool) {
	if o == nil || o.TcpTimeouts == nil {
		return nil, false
	}
	return o.TcpTimeouts, true
}

// HasTcpTimeouts returns a boolean if a field has been set.
func (o *SingleAggregatedConnectionResponseDataAttributes) HasTcpTimeouts() bool {
	return o != nil && o.TcpTimeouts != nil
}

// SetTcpTimeouts gets a reference to the given int64 and assigns it to the TcpTimeouts field.
func (o *SingleAggregatedConnectionResponseDataAttributes) SetTcpTimeouts(v int64) {
	o.TcpTimeouts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SingleAggregatedConnectionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BytesSentByClient != nil {
		toSerialize["bytes_sent_by_client"] = o.BytesSentByClient
	}
	if o.BytesSentByServer != nil {
		toSerialize["bytes_sent_by_server"] = o.BytesSentByServer
	}
	if o.GroupBys != nil {
		toSerialize["group_bys"] = o.GroupBys
	}
	if o.PacketsSentByClient != nil {
		toSerialize["packets_sent_by_client"] = o.PacketsSentByClient
	}
	if o.PacketsSentByServer != nil {
		toSerialize["packets_sent_by_server"] = o.PacketsSentByServer
	}
	if o.RttMicroSeconds != nil {
		toSerialize["rtt_micro_seconds"] = o.RttMicroSeconds
	}
	if o.TcpClosedConnections != nil {
		toSerialize["tcp_closed_connections"] = o.TcpClosedConnections
	}
	if o.TcpEstablishedConnections != nil {
		toSerialize["tcp_established_connections"] = o.TcpEstablishedConnections
	}
	if o.TcpRefusals != nil {
		toSerialize["tcp_refusals"] = o.TcpRefusals
	}
	if o.TcpResets != nil {
		toSerialize["tcp_resets"] = o.TcpResets
	}
	if o.TcpRetransmits != nil {
		toSerialize["tcp_retransmits"] = o.TcpRetransmits
	}
	if o.TcpTimeouts != nil {
		toSerialize["tcp_timeouts"] = o.TcpTimeouts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SingleAggregatedConnectionResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BytesSentByClient         *int64              `json:"bytes_sent_by_client,omitempty"`
		BytesSentByServer         *int64              `json:"bytes_sent_by_server,omitempty"`
		GroupBys                  map[string][]string `json:"group_bys,omitempty"`
		PacketsSentByClient       *int64              `json:"packets_sent_by_client,omitempty"`
		PacketsSentByServer       *int64              `json:"packets_sent_by_server,omitempty"`
		RttMicroSeconds           *int64              `json:"rtt_micro_seconds,omitempty"`
		TcpClosedConnections      *int64              `json:"tcp_closed_connections,omitempty"`
		TcpEstablishedConnections *int64              `json:"tcp_established_connections,omitempty"`
		TcpRefusals               *int64              `json:"tcp_refusals,omitempty"`
		TcpResets                 *int64              `json:"tcp_resets,omitempty"`
		TcpRetransmits            *int64              `json:"tcp_retransmits,omitempty"`
		TcpTimeouts               *int64              `json:"tcp_timeouts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bytes_sent_by_client", "bytes_sent_by_server", "group_bys", "packets_sent_by_client", "packets_sent_by_server", "rtt_micro_seconds", "tcp_closed_connections", "tcp_established_connections", "tcp_refusals", "tcp_resets", "tcp_retransmits", "tcp_timeouts"})
	} else {
		return err
	}
	o.BytesSentByClient = all.BytesSentByClient
	o.BytesSentByServer = all.BytesSentByServer
	o.GroupBys = all.GroupBys
	o.PacketsSentByClient = all.PacketsSentByClient
	o.PacketsSentByServer = all.PacketsSentByServer
	o.RttMicroSeconds = all.RttMicroSeconds
	o.TcpClosedConnections = all.TcpClosedConnections
	o.TcpEstablishedConnections = all.TcpEstablishedConnections
	o.TcpRefusals = all.TcpRefusals
	o.TcpResets = all.TcpResets
	o.TcpRetransmits = all.TcpRetransmits
	o.TcpTimeouts = all.TcpTimeouts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
