// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPostRequestAttributes Attributes for Azure config Post Request.
type AzureUCConfigPostRequestAttributes struct {
	// The tenant ID of the Azure account.
	AccountId string `json:"account_id"`
	// Bill config.
	ActualBillConfig BillConfig `json:"actual_bill_config"`
	// Bill config.
	AmortizedBillConfig BillConfig `json:"amortized_bill_config"`
	// The client ID of the Azure account.
	ClientId string `json:"client_id"`
	// The scope of your observed subscription.
	Scope string `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureUCConfigPostRequestAttributes instantiates a new AzureUCConfigPostRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureUCConfigPostRequestAttributes(accountId string, actualBillConfig BillConfig, amortizedBillConfig BillConfig, clientId string, scope string) *AzureUCConfigPostRequestAttributes {
	this := AzureUCConfigPostRequestAttributes{}
	this.AccountId = accountId
	this.ActualBillConfig = actualBillConfig
	this.AmortizedBillConfig = amortizedBillConfig
	this.ClientId = clientId
	this.Scope = scope
	return &this
}

// NewAzureUCConfigPostRequestAttributesWithDefaults instantiates a new AzureUCConfigPostRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureUCConfigPostRequestAttributesWithDefaults() *AzureUCConfigPostRequestAttributes {
	this := AzureUCConfigPostRequestAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AzureUCConfigPostRequestAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPostRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AzureUCConfigPostRequestAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetActualBillConfig returns the ActualBillConfig field value.
func (o *AzureUCConfigPostRequestAttributes) GetActualBillConfig() BillConfig {
	if o == nil {
		var ret BillConfig
		return ret
	}
	return o.ActualBillConfig
}

// GetActualBillConfigOk returns a tuple with the ActualBillConfig field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPostRequestAttributes) GetActualBillConfigOk() (*BillConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualBillConfig, true
}

// SetActualBillConfig sets field value.
func (o *AzureUCConfigPostRequestAttributes) SetActualBillConfig(v BillConfig) {
	o.ActualBillConfig = v
}

// GetAmortizedBillConfig returns the AmortizedBillConfig field value.
func (o *AzureUCConfigPostRequestAttributes) GetAmortizedBillConfig() BillConfig {
	if o == nil {
		var ret BillConfig
		return ret
	}
	return o.AmortizedBillConfig
}

// GetAmortizedBillConfigOk returns a tuple with the AmortizedBillConfig field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPostRequestAttributes) GetAmortizedBillConfigOk() (*BillConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmortizedBillConfig, true
}

// SetAmortizedBillConfig sets field value.
func (o *AzureUCConfigPostRequestAttributes) SetAmortizedBillConfig(v BillConfig) {
	o.AmortizedBillConfig = v
}

// GetClientId returns the ClientId field value.
func (o *AzureUCConfigPostRequestAttributes) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPostRequestAttributes) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *AzureUCConfigPostRequestAttributes) SetClientId(v string) {
	o.ClientId = v
}

// GetScope returns the Scope field value.
func (o *AzureUCConfigPostRequestAttributes) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPostRequestAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AzureUCConfigPostRequestAttributes) SetScope(v string) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureUCConfigPostRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["actual_bill_config"] = o.ActualBillConfig
	toSerialize["amortized_bill_config"] = o.AmortizedBillConfig
	toSerialize["client_id"] = o.ClientId
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureUCConfigPostRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId           *string     `json:"account_id"`
		ActualBillConfig    *BillConfig `json:"actual_bill_config"`
		AmortizedBillConfig *BillConfig `json:"amortized_bill_config"`
		ClientId            *string     `json:"client_id"`
		Scope               *string     `json:"scope"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.ActualBillConfig == nil {
		return fmt.Errorf("required field actual_bill_config missing")
	}
	if all.AmortizedBillConfig == nil {
		return fmt.Errorf("required field amortized_bill_config missing")
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "actual_bill_config", "amortized_bill_config", "client_id", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = *all.AccountId
	if all.ActualBillConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActualBillConfig = *all.ActualBillConfig
	if all.AmortizedBillConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AmortizedBillConfig = *all.AmortizedBillConfig
	o.ClientId = *all.ClientId
	o.Scope = *all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
