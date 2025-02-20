// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveIntegrationAzure The Azure archive's integration destination.
type LogsArchiveIntegrationAzure struct {
	// A client ID.
	ClientId string `json:"client_id"`
	// A tenant ID.
	TenantId string `json:"tenant_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveIntegrationAzure instantiates a new LogsArchiveIntegrationAzure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveIntegrationAzure(clientId string, tenantId string) *LogsArchiveIntegrationAzure {
	this := LogsArchiveIntegrationAzure{}
	this.ClientId = clientId
	this.TenantId = tenantId
	return &this
}

// NewLogsArchiveIntegrationAzureWithDefaults instantiates a new LogsArchiveIntegrationAzure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveIntegrationAzureWithDefaults() *LogsArchiveIntegrationAzure {
	this := LogsArchiveIntegrationAzure{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *LogsArchiveIntegrationAzure) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationAzure) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *LogsArchiveIntegrationAzure) SetClientId(v string) {
	o.ClientId = v
}

// GetTenantId returns the TenantId field value.
func (o *LogsArchiveIntegrationAzure) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationAzure) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *LogsArchiveIntegrationAzure) SetTenantId(v string) {
	o.TenantId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveIntegrationAzure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["tenant_id"] = o.TenantId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveIntegrationAzure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId *string `json:"client_id"`
		TenantId *string `json:"tenant_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "tenant_id"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.TenantId = *all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
