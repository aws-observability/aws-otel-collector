// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NetworkHealthInsightTrafficVolume Network traffic volume metrics between the client and server services during the query window.
type NetworkHealthInsightTrafficVolume struct {
	// Total bytes read from the server to the client during the query window.
	BytesRead *int64 `json:"bytes_read,omitempty"`
	// Total bytes written from the client to the server during the query window.
	BytesWritten *int64 `json:"bytes_written,omitempty"`
	// Sum of bytes written and bytes read across the query window.
	TotalTraffic *int64 `json:"total_traffic,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkHealthInsightTrafficVolume instantiates a new NetworkHealthInsightTrafficVolume object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkHealthInsightTrafficVolume() *NetworkHealthInsightTrafficVolume {
	this := NetworkHealthInsightTrafficVolume{}
	return &this
}

// NewNetworkHealthInsightTrafficVolumeWithDefaults instantiates a new NetworkHealthInsightTrafficVolume object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkHealthInsightTrafficVolumeWithDefaults() *NetworkHealthInsightTrafficVolume {
	this := NetworkHealthInsightTrafficVolume{}
	return &this
}

// GetBytesRead returns the BytesRead field value if set, zero value otherwise.
func (o *NetworkHealthInsightTrafficVolume) GetBytesRead() int64 {
	if o == nil || o.BytesRead == nil {
		var ret int64
		return ret
	}
	return *o.BytesRead
}

// GetBytesReadOk returns a tuple with the BytesRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightTrafficVolume) GetBytesReadOk() (*int64, bool) {
	if o == nil || o.BytesRead == nil {
		return nil, false
	}
	return o.BytesRead, true
}

// HasBytesRead returns a boolean if a field has been set.
func (o *NetworkHealthInsightTrafficVolume) HasBytesRead() bool {
	return o != nil && o.BytesRead != nil
}

// SetBytesRead gets a reference to the given int64 and assigns it to the BytesRead field.
func (o *NetworkHealthInsightTrafficVolume) SetBytesRead(v int64) {
	o.BytesRead = &v
}

// GetBytesWritten returns the BytesWritten field value if set, zero value otherwise.
func (o *NetworkHealthInsightTrafficVolume) GetBytesWritten() int64 {
	if o == nil || o.BytesWritten == nil {
		var ret int64
		return ret
	}
	return *o.BytesWritten
}

// GetBytesWrittenOk returns a tuple with the BytesWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightTrafficVolume) GetBytesWrittenOk() (*int64, bool) {
	if o == nil || o.BytesWritten == nil {
		return nil, false
	}
	return o.BytesWritten, true
}

// HasBytesWritten returns a boolean if a field has been set.
func (o *NetworkHealthInsightTrafficVolume) HasBytesWritten() bool {
	return o != nil && o.BytesWritten != nil
}

// SetBytesWritten gets a reference to the given int64 and assigns it to the BytesWritten field.
func (o *NetworkHealthInsightTrafficVolume) SetBytesWritten(v int64) {
	o.BytesWritten = &v
}

// GetTotalTraffic returns the TotalTraffic field value if set, zero value otherwise.
func (o *NetworkHealthInsightTrafficVolume) GetTotalTraffic() int64 {
	if o == nil || o.TotalTraffic == nil {
		var ret int64
		return ret
	}
	return *o.TotalTraffic
}

// GetTotalTrafficOk returns a tuple with the TotalTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightTrafficVolume) GetTotalTrafficOk() (*int64, bool) {
	if o == nil || o.TotalTraffic == nil {
		return nil, false
	}
	return o.TotalTraffic, true
}

// HasTotalTraffic returns a boolean if a field has been set.
func (o *NetworkHealthInsightTrafficVolume) HasTotalTraffic() bool {
	return o != nil && o.TotalTraffic != nil
}

// SetTotalTraffic gets a reference to the given int64 and assigns it to the TotalTraffic field.
func (o *NetworkHealthInsightTrafficVolume) SetTotalTraffic(v int64) {
	o.TotalTraffic = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkHealthInsightTrafficVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BytesRead != nil {
		toSerialize["bytes_read"] = o.BytesRead
	}
	if o.BytesWritten != nil {
		toSerialize["bytes_written"] = o.BytesWritten
	}
	if o.TotalTraffic != nil {
		toSerialize["total_traffic"] = o.TotalTraffic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkHealthInsightTrafficVolume) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BytesRead    *int64 `json:"bytes_read,omitempty"`
		BytesWritten *int64 `json:"bytes_written,omitempty"`
		TotalTraffic *int64 `json:"total_traffic,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bytes_read", "bytes_written", "total_traffic"})
	} else {
		return err
	}
	o.BytesRead = all.BytesRead
	o.BytesWritten = all.BytesWritten
	o.TotalTraffic = all.TotalTraffic

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
