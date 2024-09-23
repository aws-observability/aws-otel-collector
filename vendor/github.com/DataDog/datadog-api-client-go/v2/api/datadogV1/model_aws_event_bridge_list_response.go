// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeListResponse An object describing the EventBridge configuration for multiple accounts.
type AWSEventBridgeListResponse struct {
	// List of accounts with their event sources.
	Accounts []AWSEventBridgeAccountConfiguration `json:"accounts,omitempty"`
	// True if the EventBridge sub-integration is enabled for your organization.
	IsInstalled *bool `json:"isInstalled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeListResponse instantiates a new AWSEventBridgeListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeListResponse() *AWSEventBridgeListResponse {
	this := AWSEventBridgeListResponse{}
	return &this
}

// NewAWSEventBridgeListResponseWithDefaults instantiates a new AWSEventBridgeListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeListResponseWithDefaults() *AWSEventBridgeListResponse {
	this := AWSEventBridgeListResponse{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *AWSEventBridgeListResponse) GetAccounts() []AWSEventBridgeAccountConfiguration {
	if o == nil || o.Accounts == nil {
		var ret []AWSEventBridgeAccountConfiguration
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeListResponse) GetAccountsOk() (*[]AWSEventBridgeAccountConfiguration, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *AWSEventBridgeListResponse) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given []AWSEventBridgeAccountConfiguration and assigns it to the Accounts field.
func (o *AWSEventBridgeListResponse) SetAccounts(v []AWSEventBridgeAccountConfiguration) {
	o.Accounts = v
}

// GetIsInstalled returns the IsInstalled field value if set, zero value otherwise.
func (o *AWSEventBridgeListResponse) GetIsInstalled() bool {
	if o == nil || o.IsInstalled == nil {
		var ret bool
		return ret
	}
	return *o.IsInstalled
}

// GetIsInstalledOk returns a tuple with the IsInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeListResponse) GetIsInstalledOk() (*bool, bool) {
	if o == nil || o.IsInstalled == nil {
		return nil, false
	}
	return o.IsInstalled, true
}

// HasIsInstalled returns a boolean if a field has been set.
func (o *AWSEventBridgeListResponse) HasIsInstalled() bool {
	return o != nil && o.IsInstalled != nil
}

// SetIsInstalled gets a reference to the given bool and assigns it to the IsInstalled field.
func (o *AWSEventBridgeListResponse) SetIsInstalled(v bool) {
	o.IsInstalled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.IsInstalled != nil {
		toSerialize["isInstalled"] = o.IsInstalled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts    []AWSEventBridgeAccountConfiguration `json:"accounts,omitempty"`
		IsInstalled *bool                                `json:"isInstalled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accounts", "isInstalled"})
	} else {
		return err
	}
	o.Accounts = all.Accounts
	o.IsInstalled = all.IsInstalled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
